package models

type Actor struct {
	ID         uint64 `gorm:"column:ID;primary_key"`
	PlayerID   uint64 `gorm:"column:playerID;not null"`   //用户ID
	ActivityID uint64 `gorm:"column:activityID;not null"` //活动ID
	Nums       uint8  `gorm:"default:1"`                  //参与人数
	Price      int32  `gorm:"defualt:0"`                  //价格
	Status     int8   `gorm:"default:0"`                  //状态
	Base
}

type ActorStatus int8

const (
	// ActorStatusCancel = -1 //已取消
	ActorStatusSignUp = 0 //已报名
	ActorStatusPayed  = 1 //已支付
)

func GetPlayersWithActivity(activity *Activity) (out []*Player) {
	out = make([]*Player, 0)
	rows, err := DB.Table(initConfigMap[TActor].TableName).Where("activityID=?", activity.ID).Select("playerID").Order("ctime asc").Limit(10).Rows()
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id uint64
			rows.Scan(&id)
			out = append(out, Get(TPlayer, id).(*Player))
		}
	}
	return
}

func init() {
	cacheExpired := int64(600)
	cacheEnable := false
	if cacheExpired > 0 {
		cacheEnable = true
	}
	InitConfig(&Config{
		Type:         TActor,
		EnableCache:  cacheEnable,
		CachePrefix:  "Actor",
		CacheExpired: cacheExpired,
		GetByIdFunc:  getActorByID,
		TableName:    "actors",
	})
}

func getActorByID(id uint64) (l *Actor, err error) {
	l = new(Actor)
	err = DB.Where("id=?", id).First(l).Error
	l.ID = id
	return
}
