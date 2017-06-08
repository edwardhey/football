package controllers

import (
	"fmt"
	"time"

	"edwardhey.com/football/wx/models"
	"github.com/astaxie/beego"
	wx_oauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	"github.com/chanxuehong/wechat.v2/oauth2"
)

const (
	oauth2RedirectURI = "http://fb.innomix.com/wx/callback" // 填上自己的参数
	oauth2AuthURI     = "http://fb.innomix.com/wx/auth"     //授权url
	oauth2Scope       = "snsapi_userinfo"                   // 填上自己的参数
)

var (
	wxAppID     = beego.AppConfig.String("wx.appID")
	wxAppSecret = beego.AppConfig.String("wx.appSecret")
	// sessionStorage                 = session.New(20*60, 60*60)
	oauth2Endpoint oauth2.Endpoint = wx_oauth2.NewEndpoint(wxAppID, wxAppSecret)
)

type WxController struct {
	BaseController
}

func (c *WxController) Prepare() {
	// c.IsJSON = false
	c.BaseController.Prepare()
}

func (c *WxController) Auth() {
	url := wx_oauth2.AuthCodeURL(wxAppID, "http://fb.innomix.com.cn/wx/callback", "snsapi_userinfo", "")
	c.Redirect(url, 302)
	//fmt.Println(url)
}

func (c *WxController) Callback() {
	// var err error
	fmt.Println(c.GetString("code"))
	code := c.GetString("code")
	if code == "" {
		// c.Ctx.WriteString("用户禁止授权")
		c.ThrowErr("用户禁止授权")
	}
	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	fmt.Println(oauth2Client)
	// token, err := oauth2Client.ExchangeToken(code)
	// if err != nil {
	// 	// io.WriteString(w, err.Error())
	// 	// log.Println(err)
	// 	c.ThrowErr(err.Error())
	// 	// return
	// }
	token := &oauth2.Token{
		AccessToken:  "IanLd6tnFFpYr2uqLYaDMmtebSLU8KC2o2-RfCjfJUN4kdGDjn0-AwQAwZiiRMJtvLQOXGviiXc4CVrcPo75Zg-SpGlblohTYXacS8VzEVw",
		CreatedAt:    1496555586,
		ExpiresIn:    6000,
		RefreshToken: "FE0WuLpaeljVxfsJwQqy8hVseerDM7J0pCdTSHrYAuhbnNSYZksT9jtXSfnyHpDU6AXi1EbrCY9XpZLLqU4M5Yi0R0IuGfvPJXQh4cWAFm0",
		OpenId:       "oIlcp1ocbxjOhScmOWIn67ViKpXo",
		Scope:        "snsapi_userinfo",
	}
	// log.Printf("token: %+v\r\n", token)

	// userinfo, err := wx_oauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	// if err != nil {
	// 	// io.WriteString(w, err.Error())
	// 	c.ThrowErr(err.Error())
	// 	// return
	// }

	// json.NewEncoder(w).Encode(userinfo)
	userinfo := wx_oauth2.UserInfo{
		OpenId:       "oIlcp1ocbxjOhScmOWIn67ViKpXo",
		Nickname:     "Edward",
		Sex:          1,
		City:         "深圳",
		Province:     "广东",
		Country:      "中国",
		HeadImageURL: "http://wx.qlogo.cn/mmopen/gd2cuSbic0WspIBdQEGyTusX4ezM0aibhqR3UUAHbHLo1Eo42YBibnibFmH0QzPUagSIiaia3aaXDD6TQZ3X9ictoT5v3P2wrbpUWkI/0",
	}
	// log.Printf("userinfo: %+v\r\n", userinfo)

	now := uint32(time.Now().Unix())

	player := models.GetPlayerWithOpenID(token.OpenId)
	// fmt.Println(player.IsExists())
	// fmt.Println(models.Get(models.TPlayer, uint64(51399136452280320)))
	// return
	if player.IsNew() {
		player.ID, _ = models.IDGenr.Next()
	}

	player.TokenExpired = uint32(token.ExpiresIn) + now
	player.OpenID = token.OpenId
	player.Name = userinfo.Nickname
	player.HeadImg = userinfo.HeadImageURL
	player.LastLogon = now
	player.AccessToken = token.AccessToken
	player.RefreshToken = token.RefreshToken
	models.Save(player)

	c.StartSession()
	c.SetSession("player", player)
	c.Redirect("/", 302)
	// c.CustomAbort(200, "aaa")
}
