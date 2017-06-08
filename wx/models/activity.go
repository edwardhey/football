package models

import (
	"fmt"
	"strconv"
	"time"
)

type Activity struct {
	ID           uint64         `gorm:"column:ID;primary_key"`
	Name         string         `gorm:"varchar(100)"`                  //活动名称
	Address      string         `gorm:"varchar(255)"`                  //活动地址
	BeginTime    uint32         `gorm:"column:beginTime;default:0"`    //开始时间
	EndTime      uint32         `gorm:"column:endTime;default:0"`      //结束时间
	Status       ActivityStatus `gorm:"default:0;type:tinyint"`        //状态
	OpenTime     uint32         `gorm:"column:openTime;default:0"`     //活动报名时间
	CloseTime    uint32         `gorm:"column:closeTime;default:0"`    //活动截止报名时间
	MinActorNums uint32         `gorm:"column:minActorNums;default:0"` //最小参与者人数
	MaxActorNums uint32         `gorm:"column:maxActorNums;default:0"` //最大参与者人数
	Lat          uint32         `gorm:"default:0"`                     //纬度
	Lng          uint32         `orm:"default:0"`                      //经度
	Base
}

type ActivityStatus int8

const (
	ActivityStatusCancel      = ActivityStatus(-1)  //活动已取消
	ActivityStatusUnactivated = ActivityStatus(0)   //活动未激活，未到报名开始时间
	ActivityStatusActivated   = ActivityStatus(10)  //活动已激活，可以报名
	ActivityStatusInEffect    = ActivityStatus(30)  //活动已生效
	ActivityStatusComplete    = ActivityStatus(100) //活动已完成报名
)

var StatusString map[ActivityStatus]string = map[ActivityStatus]string{
	ActivityStatusActivated:   "已开始，可以报名",
	ActivityStatusUnactivated: "未开始",
	ActivityStatusInEffect:    "已生效，完成报名，请准时参加",
	ActivityStatusComplete:    "已结束",
	ActivityStatusCancel:      "已取消",
}

func GetActivityStatusString(s ActivityStatus) string {
	return StatusString[s]
}

// var idGenr *SnowFlake

// func init() {
// 	idGenr, _ = GetSnowFlake(BusinessTypeActivity)
// }

func (a *Activity) IDString() string {
	return strconv.FormatUint(a.ID, 10)
}
func NewActivaty(t time.Time) *Activity {
	id, _ := IDGenr.Next()

	_t, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 20:30:00", t.Format("2006-01-02")))
	_tUninxTime := uint32(_t.Unix())
	// begintTime := _t.Format("2006-01-02 15:04")

	activity := &Activity{
		ID:      id,
		Status:  ActivityStatusUnactivated,
		Lat:     22550709,
		Lng:     114084155,
		Address: "广东省深圳市福田区华新村五人足球场",
		//Name:         fmt.Sprintf("华新村 %s 足球运动", begintTime),
		Name:         "华新村足球运动",
		MaxActorNums: 15,
		MinActorNums: 5,
		BeginTime:    _tUninxTime,
		EndTime:      _tUninxTime + 7200,
		OpenTime:     uint32(time.Now().Unix()),
	}
	activity.CloseTime = activity.BeginTime - 1
	// activity.Name =
	return activity
}

func GetActivities() (out []*Activity) {
	var ids []struct {
		ID uint64 `gorm:"column:ID"`
	}
	// var activities []Activity
	out = make([]*Activity, 0)
	if err := DB.Table("activities").Limit(10).Order("ctime desc").Scan(&ids).Error; err == nil {
		for _, id := range ids {
			// out[idx] = Get(TActivity, id.ID).(*Activity)
			out = append(out, Get(TActivity, id.ID).(*Activity))
		}
	}
	return
}

func GetNextWeekday(w time.Weekday) time.Time {
	now := time.Now()
	weekDay := now.Weekday()
	days := 0
	if weekDay < w {
		days = int(w - weekDay)
	} else {
		days = int(time.Weekday(7) - weekDay + w)
	}
	// fmt.Println()
	m, _ := time.ParseDuration(fmt.Sprintf("%dh", days*24))
	return now.Add(m)
	// fmt.Println(days)
}

func init() {
	cacheExpired := int64(600)
	cacheEnable := false
	if cacheExpired > 0 {
		cacheEnable = true
	}
	InitConfig(&Config{
		Type:         TActivity,
		EnableCache:  cacheEnable,
		CachePrefix:  "Activity",
		CacheExpired: cacheExpired,
		GetByIdFunc:  getActivityByID,
	})
}

func getActivityByID(id uint64) (l *Activity, err error) {
	l = new(Activity)
	err = DB.Where("id=?", id).First(l).Error
	l.ID = id
	return
}
