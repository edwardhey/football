package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:@/db?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:@/football?charset=utf8", 30)
	orm.RegisterModel(new(Player))
	syncDB()

}

func syncDB() {
	orm.RunSyncdb("default", false, true)
}
