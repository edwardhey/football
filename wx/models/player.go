package models

type Player struct {
	ID           uint64 `gorm:"column:ID;primary_key"`
	OpenID       string `gorm:"column:openID;type:char(64);unique_index"` //微信openID
	Name         string `gorm:"type:char(100)"`
	HeadImg      string `gorm:"column:headImg;type:char(255)"`
	Balance      uint32 `gorm:"default:0"`
	LastLogon    uint32 `gorm:"column:lastLogon;default:0"` //上次登陆时间戳
	AccessToken  string `gorm:"column:accessToken;type:char(255)"`
	RefreshToken string `gorm:"column:refreshToken;type:char(255)"`
	TokenExpired uint32 `gorm:"column:tokenExpred"`
	Base
}

func (u *Player) TableName() string {
	return initConfigMap[TPlayer].TableName
}

func init() {
	cacheExpired := int64(600)
	cacheEnable := false
	if cacheExpired > 0 {
		cacheEnable = true
	}
	InitConfig(&Config{
		Type:         TPlayer,
		EnableCache:  cacheEnable,
		CachePrefix:  "Player",
		CacheExpired: cacheExpired,
		GetByIdFunc:  getPlayerByID,
		TableName:    "players",
	})
}

func getPlayerByID(id uint64) (l *Player, err error) {
	l = new(Player)
	// ORM.QueryTable(l).Filter("id", id)
	err = DB.Where("id=?", id).First(l).Error
	l.ID = id
	return
}

func (p *Player) IsJoinActivity(a *Activity) bool {
	// return ORM.QueryTable("players").Filter("playerID", p.ID).Filter("activityID", a.ID).Exist()
	var rowNums uint

	DB.Table(initConfigMap[TActor].TableName).Where("playerID=? and activityID=?", p.ID, a.ID).Count(&rowNums)
	if rowNums > 0 {
		return true
	}
	return false
}

func (p *Player) JoinActivity(a *Activity, nums uint8) *Actor {
	id, _ := IDGenr.Next()
	price := int32(nums) * a.Price
	return &Actor{
		ID:         id,
		PlayerID:   p.ID,
		ActivityID: a.ID,
		Price:      price,
		Nums:       nums,
		Status:     ActorStatusPayed,
	}
}

func GetPlayerWithOpenID(openID string) *Player {
	var id uint64
	row := DB.Table(initConfigMap[TPlayer].TableName).Where("openID = ?", openID).Select("id").Row() // (*sql.Row)
	err := row.Scan(&id)
	if err == nil {
		return Get(TPlayer, id).(*Player)
	}
	return &Player{
		OpenID: openID,
	}
}
