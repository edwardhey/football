package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "fb.edwardhey.com"
	c.Data["Email"] = "fb@edwardhey.com"
	c.TplName = "index.tpl"
}
