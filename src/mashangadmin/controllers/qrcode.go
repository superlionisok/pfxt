package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"models"
	_ "models"
)

type QrCodeController struct {
	BaseController
}

func (c *QrCodeController) IndexGet() {

	c.DoInit()
	var u = c.SessionLogin

	var payChannels []models.SysPayChannel
	o := orm.NewOrm()
	aa, err := o.QueryTable("SysPayChannel").Filter("IsDel", 0).All(&payChannels, "Id", "Title")
	if err != nil {
		panic(err)
		c.Abort("404")
	}
	beego.Notice(aa)
	c.Data["payChannels"] = payChannels

	var qrcodes []models.MsQrCode
	o.QueryTable("MsQrCode").Filter("IsDel", 0).Filter("MsUserID", u.ID).All(&qrcodes)
	c.Data["qrcodes"] = qrcodes

	c.TplName = "qrcode/index.html"

}
