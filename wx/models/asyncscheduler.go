package models

import (
	"fmt"

	// "golib/link/httplink"

	"edwardhey.com/asyncscheduler/interfaces"
	"edwardhey.com/asyncscheduler/job"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

func SetAsyncscheduler(job *job.Job) {
	// json := []byte(`{"read_write_timeout": "3","id": "111","connect_timeout": "3","payload": {"aa": "bb"},"max_attempts": "2","attempt_interval": "10","callback": "http://chargerlink.com","url": "http://charger.com"}`)
	// job := &job.Job{}
	// resp := interfaces.Resp{}
	var resp interfaces.Resp
	req, _ := httplib.Post(beego.AppConfig.String("asyncscheduler") + "/job/set").JSONBody(job)
	err := req.ToJSON(&resp)
	if err != nil {
		fmt.Println(err)
	}
	// body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp, err)
}
