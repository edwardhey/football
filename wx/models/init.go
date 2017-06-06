package models

import (
	"errors"
	"fmt"
	"golib/cache"
	"golib/log"
	"reflect"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ModelType string

type Config struct {
	EnableCache  bool
	CachePrefix  string
	CacheExpired int64
	Type         ModelType
	GetByIdFunc  interface{}
}

const (
	TPlayer = ModelType("Player")
	MC_KEY  = "%s_%v"

	BusinessTypeID       = 0
	BusinessTypeActivity = 0
)

var (
	DB     *gorm.DB
	MC     *cache.MemoryCache
	IDGenr *SnowFlake

	//启用cache的对象列表
	modelsLegth   int                   = 10
	initConfigMap map[ModelType]*Config = make(map[ModelType]*Config, modelsLegth)
)

func InitManual() {
	MC = cache.NewMemoryCache()
	MC.StartAndGC(300)
	IDGenr, _ = NewSnowFlake(BusinessTypeID)
	log.Debug("init db")
	var err error
	DB, err = gorm.Open("mysql", beego.AppConfig.String("db.dsn"))
	if err != nil {
		panic(err)
	}

	err = DB.DB().Ping()
	if err != nil {
		panic(fmt.Sprintf("db connect failued!%s\r\n", err))
	}

	DB.SingularTable(true)
}

func InitConfig(c *Config) {
	initConfigMap[c.Type] = c
}

/**
 * 保存对象
 * @param {[type]} obj interface{}) (err error [description]
 */
func SaveDb(objList ...interface{}) (err error) {
	if len(objList) == 0 {
		return nil
	} else if len(objList) == 1 {
		obj := objList[0]
		elem := reflect.ValueOf(obj).Elem()
		if elem.FieldByName("isDeleted").Bool() == true {
			return nil
		}
		if elem.FieldByName("isExists").Bool() == true {
			err = DB.Save(obj).Error
		} else {
			err = DB.Create(obj).Error
		}
		if err != nil {
			log.Error("save db error", err)
			return
		}
	} else {
		tx := DB.Begin()
		defer func() {
			if err == nil {
				err = tx.Commit().Error
			} else {
				tx.Rollback()
			}
		}()
		for _, obj := range objList {
			elem := reflect.ValueOf(obj).Elem()
			if elem.FieldByName("isDeleted").Bool() == true {
				return fmt.Errorf("%s %v is already deleted", elem.Type().Name(), elem.FieldByName("Id"))
			}
			if elem.FieldByName("isExists").Bool() == true {
				err = tx.Save(obj).Error
			} else {
				err = tx.Create(obj).Error
			}
			if err != nil {
				return
			}
		}
	}
	return
}

func getMcKeyByObjName(objName string) (prefix string, expired int64) {
	return getMcKeyByType(ModelType(objName))
}

func getMcKeyByType(t ModelType) (prefix string, expired int64) {
	if c, ok := initConfigMap[t]; ok {
		if c.EnableCache && c.CacheExpired > 0 {
			return c.CachePrefix, c.CacheExpired
		}
	}
	return "", 0
}

/**
 * 获取对象
 */
func Get(t ModelType, id interface{}) interface{} {
	var obj interface{}

	if config, ok := initConfigMap[t]; ok && config.EnableCache {
		key := fmt.Sprintf(MC_KEY, config.CachePrefix, id)
		if obj = MC.Get(key); obj != nil {
			return obj
		}
	}

	obj = GetDb(t, id)

	elem := reflect.ValueOf(obj).Elem()
	if !elem.FieldByName("isExists").Bool() {
		return obj
	}

	if config, ok := initConfigMap[t]; ok && config.EnableCache {
		key := fmt.Sprintf(MC_KEY, config.CachePrefix, id)
		MC.Put(key, obj, config.CacheExpired)
	}
	return obj
}

func SaveMc(objList ...interface{}) (err error) {
	for _, obj := range objList {
		ref := reflect.ValueOf(obj)
		elem := ref.Elem()
		objType := ModelType(reflect.TypeOf(obj).Elem().Name())

		// elem.MethodByName("SetExists").Call(nil)
		reflect.ValueOf(obj).MethodByName("SetExists").Call(nil)

		if config, ok := initConfigMap[objType]; ok && config.EnableCache {
			key := fmt.Sprintf(MC_KEY, config.CachePrefix, elem.FieldByName("Id"))
			MC.Put(key, obj, config.CacheExpired)
		}
	}
	return nil
}

/**
 * 保存对象(会导致panic)
 * @param {[type]} obj interface{}  需要保存的对象的指针
 * @return  error
 */
func Save(objList ...interface{}) (err error) {
	if len(objList) == 0 {
		return nil
	}
	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%v", msg)
		}
	}()

	// fmt.Println(objList)
	err = SaveDb(objList...)
	/*
		if err != nil {
			return
		}*/
	SaveMc(objList)
	return
}

/**
 * 删除对象(会导致panic)
 * @param {[type]} obj interface{}  需要删除的对象的指针
 * @return  error
 */
func Del(objs ...interface{}) (err error) {
	if len(objs) == 0 {
		return nil
	}
	//从缓存中删除
	for _, obj := range objs {
		ref := reflect.ValueOf(obj)
		elem := ref.Elem()
		objType := ModelType(reflect.TypeOf(obj).Elem().Name())
		objId := elem.FieldByName("Id")
		if reflect.Zero(objId.Type()).Interface() == objId.Interface() {
			return errors.New(fmt.Sprintf("error object to delete %v", obj))
		}

		//先删cache，后删db
		if config, ok := initConfigMap[objType]; ok && config.EnableCache {
			key := fmt.Sprintf(MC_KEY, config.CachePrefix, objId)
			MC.Delete(key)
		}
	}
	if len(objs) == 1 {
		err = DB.Delete(objs[0]).Error
	} else {
		tx := DB.Begin()
		for _, obj := range objs {
			tx.Delete(obj)
		}
		err = tx.Commit().Error
	}
	if err != nil {
		return
	}
	for _, obj := range objs {
		ref := reflect.ValueOf(obj)
		if resetFunc := ref.MethodByName("Reset"); resetFunc.IsValid() {
			resetFunc.Call(nil)
		} else {
			reset(obj)
		}
	}
	return
}

/**
 * 重置对象为空对象(会导致panic)
 * @param  {[type]} obj interface{} [description]
 */
func reset(obj interface{}) {
	elem := reflect.ValueOf(obj).Elem()
	elem.Set(reflect.Zero(reflect.TypeOf(elem.Interface())))
}

func GetDb(t ModelType, id interface{}) interface{} {
	fn := reflect.ValueOf(initConfigMap[t].GetByIdFunc)
	params := []reflect.Value{reflect.ValueOf(id)}
	res := fn.Call(params)
	return res[0].Interface()
}
