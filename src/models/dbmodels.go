package models
import (

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)
// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
	CreateTime time.Time
}

func init() {
	// set default database
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:!QAZxsw2@tcp(10.10.15.154:3306)/pfxt?charset=utf8", maxIdle, maxConn)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}
