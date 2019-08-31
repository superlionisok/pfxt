package routers

import (
	"github.com/astaxie/beego"
	"mashangadmin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{}, "get:LoginGet;post:LoginPost")
}
