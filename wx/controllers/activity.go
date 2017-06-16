package controllers

import (
	"fmt"

	"edwardhey.com/football/wx/models"
)

type ActivityController struct {
	AuthorizedController
}

func (c *ActivityController) Info() {

}

//Join 参加活动
func (c *ActivityController) Join() {
	c.IsJSON = true
	// safe := &io.LimitedReader{R: c.Ctx..Context.Request.Body, N: 100000000}
	_id, err := c.GetInt64("ID")
	if err != nil {
		c.ThrowErr("ID 不合法")
	}
	id := uint64(_id)
	activity := models.Get(models.TActivity, id).(*models.Activity)
	if activity.IsNew() {
		c.ThrowErr("活动不存在")
	}

	fmt.Println(models.GetActorsWithActivity(activity))
	p := c.GetSession("player").(*models.Player)
	fmt.Println(p.IsJoinActivity(activity))
	if p.IsJoinActivity(activity) {
		c.ThrowErr("您已报名参加，请勿重复报名")
	}
	actor := p.JoinActivity(activity, 1)
	fmt.Println(actor)
	err = models.Save(actor)
	fmt.Println(err)
}
