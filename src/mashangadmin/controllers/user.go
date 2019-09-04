package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/mojocn/base64Captcha"
	"helper"
	"models"
	_ "models"
	"strings"
)

var (
	vcodestrbase = ""
)

const cookiekey = "lioncookiekey"

type UserController struct {
	beego.Controller
}

func (c *UserController) LoginGet() {

	//var str = "123456"
	//var md5str = helper.PassWordMd5(str)
	//fmt.Println("md5=", md5str)
	//
	//var base64str = helper.PassWordBase64Ecode(str)
	//
	//fmt.Println("base64=", base64str)
	//
	//var str2 = helper.PassWordBase64Decode(base64str)
	//fmt.Println("ebase64=", str2)
	//
	//var aesbyte = []byte("12345678")
	////key := []byte("2fa6c1e9")
	//aesstr, err := helper.PassWordDesEncrypt(str, aesbyte)
	//if err != nil {
	//	fmt.Println("aes 加密错误", err)
	//} else {
	//
	//	fmt.Println("aes=", aesstr)
	//}
	//
	//aesstr2, err :=helper.PassWordDesDecrypt(aesstr,aesbyte)
	//if err != nil {
	//	fmt.Println("aes 解密错误", err)
	//} else {
	//
	//	fmt.Println("aes解密=", aesstr2)
	//}
	//

	c.TplName = "user/login.html"
}
func (c *UserController) LoginPost() {

	vcodeimg := c.GetString("Vcode")

	vcoderes := verfiyCaptcha(vcodestrbase, vcodeimg)
	if !vcoderes {
		c.Data["msg"] = "验证码不正确!"
		c.TplName = "user/login.html"
		return
	}

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
	beego.Debug("登录成功")
	s, err := json.Marshal(user)
	if err != nil {
		beego.Notice("user to json  err")
		return
	}

	c.SetSecureCookie(cookiekey, "msuser", "ms-"+loginName)
	helper.RedisSetByTime("ms-"+loginName, string(s), 1200)
	c.Ctx.Redirect(302, "/")

	//c.TplName = "user/login.html"
}

func (c *UserController) ImageCode() {

	//c.TplName = "user/login.html"
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 40,
		Width:  120,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idkey := "1234"
	idKeyC, capC := base64Captcha.GenerateCaptcha(idkey, configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	base64stringC = strings.Replace(base64stringC, "data:image/png;base64,", "", -1)
	//c.Ctx.ResponseWriter.Write( []byte(base64stringC) )
	ddd, err := base64.StdEncoding.DecodeString(base64stringC)
	//return  ddd

	if err != nil {
		c.Ctx.WriteString("base64字符串转byte失败")
		return
	}
	vcodestrbase = idKeyC
	c.Ctx.ResponseWriter.Write([]byte(ddd))

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
		CaptchaLen:         4,
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
func verfiyCaptcha(idkey, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	return verifyResult
}

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 120
	cpt.StdHeight = 40
}
