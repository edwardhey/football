package models

import (
	"fmt"
	"testing"
	"time"

	"net/url"

	"edwardhey.com/asyncscheduler/job"
)

func init() {
	// InitManual()
	// DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Activity{})
	// 	// beego.AppConfigPath = ""
	// 	// beego.AppConfig("ini", "/opt/local/go/src/edwardhey.com/football/wx/conf/app.conf")
	// 	beego.AppConfigPath = "../conf/app.conf"
	// 	beego.ParseConfig()
	// 	// fmt.Println(beego.AppConfig.String("appname"))
	// 	InitManual()
}

func _TestCreate(t *testing.T) {

	// fmt.Println(_t)
	act := NewActivaty(GetNextWeekday(time.Wednesday))
	act.EndTime = act.OpenTime + 3600
	Save(act)
	// err := Save(act)
	j := &job.Job{
		ID:              act.IDString() + "active",
		Payload:         url.Values{"ID": {act.IDString()}, "act": {"active"}},
		TTR:             act.OpenTime,
		TTL:             act.EndTime,
		Priority:        10,
		AttemptInterval: 10,
		MaxAttempts:     200,
		URL:             "http://localhost:8080/asyncscheduler/activeactivity",
		Method:          job.MethodPost,
		IgnoreResponse:  true,
	}
	SetAsyncscheduler(j)

	//下架活动
	j.ID = act.IDString() + "close"
	j.TTR = act.EndTime
	j.URL = "http://localhost:8080/asyncscheduler/deactiveactivity"
	j.Payload = url.Values{"ID": {act.IDString()}, "act": {"end"}}
	SetAsyncscheduler(j)
	// jobActivite.URL = "http://localhost:8080/asyncscheduler/"

	fmt.Println()

	// fmt.Println(getPlayerByID(52174304581779456))
}

func TestGetActivities(t *testing.T) {
	GetNewestActivities(10)
}
