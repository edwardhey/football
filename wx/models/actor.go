package models

type Actor struct {
	ID         uint64 `orm:"column(ID)"`
	PlayerID   uint64 `orm:"column(playerID)"`   //用户ID
	ActivityID uint64 `orm:"column(activityID)"` //活动ID
	Nums       int8   `orm:"default(1)"`         //参与人数
	Price      uint32 `orm:"defualt(0)"`         //价格
	Status     int8   `orm:"default(0)"`         //状态
}

type ActorStatus int8

const (
	// ActorStatusCancel = -1 //已取消
	ActorStatusSignUp = 0 //已报名
	ActorStatusPayed  = 1 //已支付
)
