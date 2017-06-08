package models

import (
	"github.com/astaxie/beego"
)

func init() {
	// beego.AppConfigPath = ""
	// beego.AppConfig("ini", "/opt/local/go/src/edwardhey.com/football/wx/conf/app.conf")
	beego.AppConfigPath = "../conf/app.conf"
	beego.ParseConfig()
	// fmt.Println(beego.AppConfig.String("appname"))
	InitManual()
}
