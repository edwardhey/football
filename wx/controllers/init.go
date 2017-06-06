package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type BaseController struct {
	beego.Controller `json:"-"`
	IsJSON           bool                   `json:"-"`
	Code             int                    `json:"code"`
	Err              string                 `json:"err,omiempty"`
	JsonData         map[string]interface{} `json:"data"`
}

func (c *BaseController) Prepare() {
	if c.IsJSON {
		c.JsonData = make(map[string]interface{}, 20)
	}
}

func (c *BaseController) Render() error {
	if c.IsJSON {
		c.Data["json"] = c
		c.Ctx.Output.SetStatus(200)
		c.ServeJSON()
	}
	return nil
}

func init() {
	tk := toolbox.NewTask("createActivity", "0 19 0 * * 3", func() error {
		//创建报名活动，并推送到微信用户
		return nil
	})
	tk.Run()
}
