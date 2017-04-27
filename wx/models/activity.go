package models

type Activity struct {
	ID           uint64 `orm:"column(ID)"`
	Name         string `orm:"size(100)"`                       //活动名称
	Address      string `orm:"size(255)"`                       //活动地址
	BeginTime    uint32 `orm:"column(beginTime);default(0)"`    //开始时间
	EndTime      uint32 `orm:"column(endTime);default(0)"`      //结束时间
	Status       int8   `orm:"default(0)"`                      //状态
	OpenTime     uint32 `orm:"column(openTime);default(0)"`     //活动报名时间
	CloseTime    uint32 `orm:"column(closeTime);default(0)"`    //活动截止报名时间
	MinActorNums uint32 `orm:"column(minActorNums);default(0)"` //最小参与者人数
	MaxActorNums uint32 `orm:"column(maxActorNums);default(0)"` //最大参与者人数
	Ctime        uint32 `orm:"column(ctime);default(0)"`        //创建时间
	Mtime        uint32 `orm:"column(mtime);default(0)"`        //更改时间
}

type ActivityStatus int8

const (
	ActivityStatusCancel      = ActivityStatus(-1)  //活动已取消
	ActivityStatusUnactivated = ActivityStatus(0)   //活动未激活，未到报名开始时间
	ActivityStatusActivated   = ActivityStatus(10)  //活动已激活，可以报名
	ActivityStatusInEffect    = ActivityStatus(30)  //活动已生效
	ActivityStatusComplete    = ActivityStatus(100) //活动已完成报名
)

var idGenr *SnowFlake

func init() {
	idGenr, _ = GetSnowFlake(BusinessTypeActivity)
}

func NewActivatyID() (uint64, error) {
	return idGenr.Next()
}
