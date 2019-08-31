package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "mashangadmin/routers"
)

func main() {

	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	//beego.SetLogFuncCall(true)

	logs.Async()
	//logs.SetLogger(logs.AdapterFile,`{"filename":"test.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.SetLevel(logs.LevelDebug)
	//logs.SetLogFuncCall(true)
	logs.EnableFuncCallDepth(true)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.Emergency("Emergency")
	logs.Alert("Alert")
	logs.Critical("Critical")
	logs.Error("Error")
	logs.Warning("Warning")
	logs.Notice("Notice")
	logs.Informational("Informational")
	logs.Debug("Debug")

	//log := logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
	//// 设置配置文件
	//log.SetLevel(logs.LevelNotice)     // 设置日志写入缓冲区的等级
	//log.EnableFuncCallDepth(true)     // 输出log时能显示输出文件名和行号（非必须）
	//
	//log.Emergency("Emergency")
	//log.Alert("Alert")
	//log.Critical("Critical")
	//log.Error("Error")
	//log.Warning("Warning")
	//log.Notice("Notice")
	//log.Informational("Informational")
	//log.Debug("Debug")
	//
	//log.Flush() // 将日志从缓冲区读出，写入到文件
	//log.Close()
	//

	beego.Run()

}
