package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/mojocn/base64Captcha"
	"models"
	_ "models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) LoginGet() {
	demoCodeCaptchaCreate()
	c.TplName = "user/login.html"
}
func (c *UserController) LoginPost() {

	loginName := c.GetString("LoginName")
	loginPwd := c.GetString("LoginPwd")

	o := orm.NewOrm()
	var user models.MsUser
	err := o.QueryTable("MsUser").Filter("LoginName", loginName).Filter("LoginPwd", loginPwd).One(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		logs.Error("有多条匹配，发生异常")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录

		beego.Notice("Not row found")
		c.Data["msg"] = "账户或者用户名不正确!"
		c.TplName = "user/login.html"
		return

	}
	logs.Debug("登录成功" + loginName)
	beego.Debug("登录成功")
	c.Ctx.Redirect(302, "/")

	//c.TplName = "user/login.html"
}

func (c *UserController) ImageCode() {

	//c.TplName = "user/login.html"
}

func demoCodeCaptchaCreate() {
	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	//config struct for audio
	//声音验证码配置
	var configA = base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	fmt.Println(idKeyA, base64stringA, "\n")
	fmt.Println(idKeyC, base64stringC, "\n")
	fmt.Println(idKeyD, base64stringD, "\n")
}
func verfiyCaptcha(idkey, verifyValue string) {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		//success
	} else {
		//fail
	}
}
