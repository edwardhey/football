package routers

import (
	"football/wx/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.AutoRouter(&controllers.ObjectController{})
}
