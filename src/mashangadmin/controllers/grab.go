package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"helper"
	"models"
	_ "models"
	"strconv"
	"time"
)

type GrabController struct {
	BaseController
}

func (c *GrabController) IndexGet() {
	c.DoInit()

	var payChannels []models.SysPayChannel
	o := orm.NewOrm()
	o.QueryTable("SysPayChannel").Filter("IsDel", 0).All(&payChannels, "Id", "Title")

	c.Data["payChannels"] = payChannels

	c.TplName = "grab/index.html"
}

func (c *GrabController) TestGet() {

	var payc models.SysPayChannel
	payc.Id = 1
	payc.Title = "111"

	str, _ := json.Marshal(payc)
	c.Ctx.WriteString(string(str))
	c.Ctx.ResponseWriter.Write([]byte(str))
}

func (c *GrabController) GetQrCodeByCodeType() {

	var id = c.GetString("id")

	var res models.ResultModel

	o := orm.NewOrm()
	var qrCode []models.MsQrCode
	_, err := o.QueryTable("MsQrCode").Filter("SysPayChannelID", id).All(&qrCode)
	if err == orm.ErrNoRows {
		// 没有找到记录
		res.Status = 0
		res.Msg = "没有记录"
		res.Data = ""
	}
	res.Status = 1
	res.Msg = "ok"
	res.Count = len(qrCode)
	res.Data = qrCode

	jsonstr, err := json.Marshal(res)
	if err != nil {
		panic("转换json失败")
		return
	}

	fmt.Println(string(jsonstr))
	c.Ctx.WriteString(string(jsonstr))

}

func (c *GrabController) GetQrCodeDetailById() {

	var id = c.GetString("id")

	var res models.ResultModel

	o := orm.NewOrm()
	var qrCode models.MsQrCode
	err := o.QueryTable("MsQrCode").Filter("ID", id).One(&qrCode)
	if err == orm.ErrNoRows {
		// 没有找到记录
		res.Status = 0
		res.Msg = "没有记录"
		res.Data = ""
	}
	res.Status = 1
	res.Msg = "ok"
	res.Count = 1
	res.Data = qrCode

	jsonstr, err := json.Marshal(res)
	if err != nil {
		panic("转换json失败")
		return
	}

	fmt.Println(string(jsonstr))
	c.Ctx.WriteString(string(jsonstr))

}

func (c *GrabController) AddTempOrder() {
	var res models.ResultModel
	var id = c.GetString("id")
	var idChannel = c.GetString("idChannel")
	var maxMoneystr = c.GetString("maxMoney")
	var minMoneystr = c.GetString("minMoney")

	qrid, err := strconv.Atoi(id)
	maxMoney, err := strconv.ParseFloat(maxMoneystr, 64)
	minMoney, err := strconv.ParseFloat(minMoneystr, 64)

	if err != nil {
		res.Status = 0
		res.Msg = "参数qrid不对"
		res.Count = 0
		res.Data = ""
		jsonstr, err := json.Marshal(res)
		if err != nil {
			panic("转换json失败")
			return
		}

		fmt.Println(string(jsonstr))
		c.Ctx.WriteString(string(jsonstr))
	}
	cid, _ := strconv.Atoi(idChannel)

	cooValue, r := c.GetSecureCookie(cookiekey, "msuser")
	if !r {
		beego.Error("get cookie err")
		c.Ctx.Redirect(302, "/user/login")
		return
	}
	ustr, err := helper.RedisGet(cooValue)
	if err != nil {
		panic("get redis value err")
	}
	var u models.MsUser
	json.Unmarshal([]byte(ustr), &u)

	o := orm.NewOrm()
	var havaone models.MsTempOrder
	err = o.QueryTable("MsTempOrder").Filter("MsUserID", u.ID).One(&havaone)

	if err != orm.ErrNoRows {
		// 没有找到记录
		res.Status = 1
		res.Msg = "已经有记录"
		res.Data = havaone
		res.Count = 1
		jsonstr, err := json.Marshal(res)
		if err != nil {
			panic("转换json失败")
			return
		}
		c.Ctx.WriteString(string(jsonstr))
		return
	}

	var add models.MsTempOrder
	add.MsUserID = u.ID
	add.MsQrCodeID = qrid
	add.SysPayChannelID = cid
	add.MinMoney = minMoney
	add.MaxMoney = maxMoney
	add.CreateTime = time.Now().In(time.Local)

	idadd, err := o.Insert(&add)

	if err != nil {
		// 没有找到记录
		res.Status = 0
		res.Msg = err.Error()
		res.Data = ""
	} else {

		res.Status = 1
		res.Msg = "ok"
		res.Count = 0
		res.Data = idadd
	}

	jsonstr, err := json.Marshal(res)
	if err != nil {
		panic("转换json失败")
		return
	}

	c.Ctx.WriteString(string(jsonstr))

}

func (c *GrabController) DelTempOrder() {
	var res models.ResultModel
	c.DoInit()
	var u = c.SessionLogin

	if u.ID == 0 {

		res.Status = 0
		res.Msg = "登录过期。请重新登录"
		res.Data = ""
		return

	} else {
		o := orm.NewOrm()
		num, err := o.QueryTable("MsTempOrder").Filter("MsUserID", u.ID).Delete()

		if err != nil {
			// 没有找到记录
			res.Status = 0
			res.Msg = err.Error()
			res.Data = ""
		} else {

			res.Status = 1
			res.Msg = "ok"
			res.Count = 1
			res.Data = num
		}

	}

	jsonstr, err := json.Marshal(res)
	if err != nil {
		panic("转换json失败")
		return
	}

	c.Ctx.WriteString(string(jsonstr))

}
