package models

import (
	"testing"
	"time"

	"net/url"

	"edwardhey.com/asyncscheduler/job"
)

// func init() {
// 	// beego.AppConfigPath = ""
// 	// beego.AppConfig("ini", "/opt/local/go/src/edwardhey.com/football/wx/conf/app.conf")
// 	beego.AppConfigPath = "../conf/app.conf"
// 	beego.ParseConfig()
// 	// fmt.Println(beego.AppConfig.String("appname"))
// 	InitManual()
// }

func _TestCreate(t *testing.T) {

	// fmt.Println(_t)
	act := NewActivaty(GetNextWeekday(time.Wednesday))
	// err := Save(act)
	j := &job.Job{
		ID:              act.IDString(),
		Payload:         url.Values{"ID": {act.IDString()}, "act": {"active"}},
		TTR:             act.OpenTime,
		TTL:             act.EndTime,
		Priority:        10,
		AttemptInterval: 10,
		MaxAttempts:     200,
		URL:             "http://localhost:8080/activity/active",
		Method:          job.MethodPost,
		IgnoreResponse:  true,
	}
	Save(act)
	SetAsyncscheduler(j)
	// fmt.Println(j)

	// fmt.Println(getPlayerByID(52174304581779456))
}

func _TestGetActivities(t *testing.T) {
	GetActivities()
}
