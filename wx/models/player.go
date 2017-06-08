package models

type Player struct {
	ID           uint64 `gorm:"column:ID;primary_key"`
	OpenID       string `gorm:"column:openID;type:varchar(60);unique_index"` //微信openID
	Name         string `gorm:"type:varchar(100)"`
	HeadImg      string `gorm:"column:headImg;type:char(255)"`
	Balance      uint32 `gorm:"default:0"`
	LastLogon    uint32 `gorm:"column:lastLogon;default:0"` //上次登陆时间戳
	AccessToken  string `gorm:"column:accessToken"`
	RefreshToken string `gorm:"column:refreshToken"`
	TokenExpired uint32 `gorm:"column:tokenExpred"`
	Base
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
	})
}

func getPlayerByID(id uint64) (l *Player, err error) {
	l = new(Player)
	err = DB.Where("id=?", id).First(l).Error
	l.ID = id
	return
}

func GetPlayerWithOpenID(openID string) *Player {

	var id struct {
		ID uint64
	}

	if err := DB.Table("players").First(&id, "openID=?", openID).Error; err == nil {
		return Get(TPlayer, id.ID).(*Player)
	}
	// fmt.Println(err)
	return &Player{
		OpenID: openID,
	}
}
