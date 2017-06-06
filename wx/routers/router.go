package routers

import (
	"edwardhey.com/football/wx/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/api/", &controllers.ApiController{}, "*")
	beego.AutoRouter(&controllers.WxController{})
}
