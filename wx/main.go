package main

import (
	"edwardhey.com/football/wx/models"
	_ "edwardhey.com/football/wx/routers"
	"github.com/astaxie/beego"
)

func main() {
	// fmt.Println(111)
	models.InitManual()
	beego.Run()
}
