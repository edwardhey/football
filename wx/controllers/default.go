package controllers

import (
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
	// fmt.Println(c.GetSession("openID"))
	// fmt.Println(GetRequestURIWithUrlFor("AsyncschedulerController.ActiveActivity"))
	c.Data["title"] = "活动列表"
	c.Data["activities"] = models.GetNewestActivities(5)
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
