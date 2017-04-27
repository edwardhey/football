package models

type Player struct {
	ID           uint64 `orm:"column(ID)"`
	OpenID       string `orm:"column(openID);size(100)"` //微信openID
	Name         string `orm:"size(100)"`
	HeadImg      string `orm:"column(headImg);size(255)"`
	Balance      uint32 `orm:"default(0)"`
	LastLogon    uint32 `orm:"column(lastLogon);default(0)"` //上次登陆时间戳
	AccessToken  string `orm:"column(accessToken)"`
	RefreshToken string `orm:"column(refreshToken)"`
	TokenExpired uint32 `orm:"column(tokenExpred)"`
	Ctime        uint32 `orm:"column(ctime);default(0)"` //创建时间
	Mtime        uint32 `orm:"column(mtime);default(0)"` //更改时间
}
