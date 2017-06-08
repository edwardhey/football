package models

import (
	"fmt"
	"testing"
)

// func init() {
// 	// beego.AppConfigPath = ""
// 	// beego.AppConfig("ini", "/opt/local/go/src/edwardhey.com/football/wx/conf/app.conf")
// 	beego.AppConfigPath = "../conf/app.conf"
// 	beego.ParseConfig()
// 	// fmt.Println(beego.AppConfig.String("appname"))
// 	InitManual()
// }

func TestGetByOpenID(t *testing.T) {
	fmt.Println(GetPlayerWithOpenID("oIlcp1ocbxjOhScmOWIn67ViKpXo"))
}
