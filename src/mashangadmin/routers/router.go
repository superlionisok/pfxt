package routers

import (
	"github.com/astaxie/beego"
	"mashangadmin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{}, "get:LoginGet;post:LoginPost")
	beego.Router("/user/imagecode", &controllers.UserController{}, "get:ImageCode")
	beego.Router("/grab/index", &controllers.GrabController{}, "get:IndexGet")
	beego.Router("/grab/test", &controllers.GrabController{}, "get:TestGet")
	beego.Router("/grab/GetQrCodeByCodeType", &controllers.GrabController{}, "get:GetQrCodeByCodeType")
	beego.Router("/grab/GetQrCodeDetailById", &controllers.GrabController{}, "get:GetQrCodeDetailById")
	beego.Router("/grab/AddTempOrder", &controllers.GrabController{}, "get:AddTempOrder")
	beego.Router("/grab/DelTempOrder", &controllers.GrabController{}, "get:DelTempOrder")
	beego.Router("/qrcode/index", &controllers.QrCodeController{}, "get:IndexGet")
	beego.Router("/qrcode/index", &controllers.QrCodeController{}, "get:IndexGet")
	beego.Router("/qrcode/create", &controllers.QrCodeController{}, "get:CreateGet;post:CreatePost")
	beego.Router("/files/upfile", &controllers.FilesController{}, "post:UpFile")

}
