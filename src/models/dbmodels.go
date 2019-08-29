package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type User struct {
	Id         int
	Name       string `orm:"size(100)"`
	CreateTime time.Time
	Sort       int
}

func init() {

	var conn = beego.AppConfig.String("conn")
	// set default database
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}
