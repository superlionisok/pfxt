package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	logs.Error("wo 就是 err test")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
