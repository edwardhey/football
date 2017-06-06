package controllers

import "fmt"

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.IsJSON = false
	c.BaseController.Prepare()
}

func (c *MainController) Get() {
	c.JsonData["data"] = "123"
	fmt.Println(c.Data)
	c.Data["Website"] = "fb.edwardhey.com"
	c.Data["Email"] = "fb@edwardhey.com"
	c.TplName = "index.tpl"
}
