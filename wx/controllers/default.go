package controllers

import (
	"fmt"

	"edwardhey.com/football/wx/models"
)

type MainController struct {
	AuthorizedController
}

func (c *MainController) Prepare() {
	c.IsJSON = false
	c.AuthorizedController.Prepare()
	// c.StartSession()
}

func (c *MainController) Get() {
	fmt.Println(c.GetSession("openID"))
	c.Data["title"] = "活动列表"
	c.Data["activities"] = models.GetActivities()
	// for a, b := range models.GetActivities() {
	// 	fmt.Println(a, b.ID)
	// }
	// fmt.Println(c.Data)
	// c.JsonData["data"] = "123"
	// c.StartSession()
	// fmt.Println(c.Data, c.GetSession("openID"))
	// c.Data["Website"] = "fb.edwardhey.com"
	// c.Data["Email"] = "fb@edwardhey.com"
	c.TplName = "index.html"
}
