package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"helper"
	"html/template"
	"models"
	_ "models"
)

type BaseController struct {
	beego.Controller
	SessionLogin models.MsUser
}

func (c *BaseController) DoInit() {

	// page start time

	cooValue, res := c.GetSecureCookie(cookiekey, "msuser")
	if !res {
		beego.Error("get cookie err")
		c.Ctx.Redirect(302, "/user/login")
		return
	}
	ustr, err := helper.RedisGetAndResetTime(cooValue)
	if err != nil {
		//	panic("get redis value err"+err.Error())

		beego.Error("get redis err")
		c.Ctx.Redirect(302, "/user/login")
		return
	}
	var u models.MsUser
	json.Unmarshal([]byte(ustr), &u)
	c.SessionLogin = u

	c.Layout = "layout.html"
	c.Data["LoginUser"] = u.LoginName
	c.Data["MenuList"] = template.HTML(GetMenuList())
}

func GetMenuList() string {

	var str = "<ul class=\"layui-nav layui-nav-tree\" lay-shrink=\"all\" id=\"LAY-system-side-menu\" lay-filter=\"layadmin-system-side-menu\">" +
		"<li data-name=\"\" data-jump=\"\" class=\"layui-nav-item lion-abiankang lion-abiankang-top layui-nav-itemed\"><a href=\"javascript:;\"" +
		" lay-tips=\"主页\" lay-direction=\"2\"><i class=\"layui-icon layui-icon-home\"></i> <cite>主页</cite> <span class=\"layui-nav-more\">" +
		"</span></a><dl class=\"layui-nav-child\">" +
		"<dd data-name=\"\" data-jump=\"home\" class=\"lion-dd layui-this\">" +
		"<a href=\"/home/index\" lay-href=\"/home/index\">控制台</a>" +
		"</dd></dl></li>" +

		"<li data-name=\"user\" data-jump=\"\" class=\"layui-nav-item lion-abiankang\">" +
		"<a href=\"javascript:;\" lay-tips=\"用户1\" lay-direction=\"2\">" +
		"<i class=\"layui-icon layui-icon-user\"></i>" +
		"<cite>管理人员</cite>" +
		"<span class=\"layui-nav-more\"></span>" +
		"</a>" +
		"<dl class=\"layui-nav-child\">" +
		"<dd data-name=\"user\" data-jump=\"muser\" class=\"lion-dd\"><a href=\"/muser/index\" lay-href=\"\" style=\"left: 14px;\">管理员列表</a></dd></dl>" +
		"</li>" +

		"<span class=\"layui-nav-bar\" style=\"top: 28px; height: 0px; opacity: 0;\"></span>" +
		"</ul>"
	return str

}
