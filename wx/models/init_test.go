package models

import (
	"fmt"
	"testing"

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

func TestBoxInitRecords(t *testing.T) {
	p := GetPlayerWithOpenID("oIlcp1ocbxjOhScmOWIn67ViKpXo")
	fmt.Println(p)
	// Save(p)

	// fmt.Println(getPlayerByID(52174304581779456))
}
