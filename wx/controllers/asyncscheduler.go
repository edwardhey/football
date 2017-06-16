package controllers

import "edwardhey.com/football/wx/models"

type AsyncschedulerController struct {
	BaseController
}

func (c *AsyncschedulerController) DeactiveActivity() {
	c.IsJSON = true
	// safe := &io.LimitedReader{R: c.Ctx..Context.Request.Body, N: 100000000}
	_id, err := c.GetInt64("ID")
	if err != nil {
		c.ThrowErr("ID 不合法")
	}
	id := uint64(_id)
	activity := models.Get(models.TActivity, id).(*models.Activity)
	if activity.IsNew() {
		return
	}
	activity.Status = models.ActivityStatusComplete
	models.Save(activity)
}

func (c *AsyncschedulerController) ActiveActivity() {
	c.IsJSON = true
	// safe := &io.LimitedReader{R: c.Ctx..Context.Request.Body, N: 100000000}
	_id, err := c.GetInt64("ID")
	if err != nil {
		c.ThrowErr("ID 不合法")
	}
	id := uint64(_id)
	activity := models.Get(models.TActivity, id).(*models.Activity)
	if activity.IsNew() {
		c.ThrowErr("活动已下线或不存在")
	}
	activity.Status = models.ActivityStatusActivated
	// fmt.Println(activity)
	models.Save(activity)
	// fmt.Println(id, aa, string(c.Ctx.Input.RequestBody))
	// // c.JsonData["data"] = "123"
	// c.StartSession()
	// fmt.Println(c.Data, c.GetSession("openID"))
	// c.Data["Website"] = "fb.edwardhey.com"
	// c.Data["Email"] = "fb@edwardhey.com"
	// c.TplName = "index.tpl"
}
