package main

import (
	"fmt"
	_ "football/wx/routers"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println(111)
	beego.Run()
}
