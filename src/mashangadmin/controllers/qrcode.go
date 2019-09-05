package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"models"
	_ "models"
	"time"
)

type QrCodeController struct {
	BaseController
}

func (c *QrCodeController) IndexGet() {

	c.DoInit()
	var u = c.SessionLogin

	var payChannels []models.SysPayChannel
	o := orm.NewOrm()
	_, err := o.QueryTable("SysPayChannel").Filter("IsDel", 0).All(&payChannels, "Id", "Title")
	if err != nil {
		panic(err)
		c.Abort("404")
	}

	c.Data["payChannels"] = payChannels

	var qrcodes []models.MsQrCode
	o.QueryTable("MsQrCode").Filter("MsUserID", u.ID).All(&qrcodes)
	c.Data["qrcodes"] = qrcodes

	c.TplName = "qrcode/index.html"

}
func (c *QrCodeController) QrcodesGet() {
	c.DoInit()
	var u = c.SessionLogin
	o := orm.NewOrm()
	var qrcodes []models.MsQrCode
	o.QueryTable("MsQrCode").Filter("MsUserID", u.ID).All(&qrcodes)

	//	dicta:=make(map[string]string)
	//dicts:=make(map[string]interface{})

}

func (c *QrCodeController) CreateGet() {
	c.DoInit()

	c.TplName = "qrcode/create.html"

}
func (c *QrCodeController) CreatePost() {
	c.DoInit()
	var u = c.SessionLogin
	var title = c.GetString("Title")
	var imgUrl = c.GetString("QrImageUrl")

	o := orm.NewOrm()
	var addqrcode models.MsQrCode
	addqrcode.ID = 0
	addqrcode.Title = title
	addqrcode.QrImageUrl = imgUrl
	addqrcode.CreateTime = time.Now()
	addqrcode.SysPayChannelID = 1
	addqrcode.MinMoney = 10
	addqrcode.MaxMoney = 100
	addqrcode.LimitMoney = 50000
	addqrcode.MsUserID = u.ID
	id, err := o.Insert(&addqrcode)
	if err == nil {
		fmt.Println(id)
		c.Data["msg"] = "ok"
	} else {
		c.Data["msg"] = "保存失败" + err.Error()

	}

	c.TplName = "qrcode/create.html"
}
