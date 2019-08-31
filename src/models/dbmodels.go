package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

//region 系统表

type SysPayChannel struct {
	Id         int
	Title      string
	CreateTime time.Time
	Sort       int
	IsDel      int
}

//endregion

//region 码商表

//码商二维码
type MsQrCode struct {
	ID              int
	SysPayChannelID int
	QrImageUrl      string
	MinMoney        float32
	MaxMoney        float32
	Title           string
	LimitMoney      float32
	CreateTime      time.Time
}

//码商用户
type MsUser struct {
	ID         int
	LoginName  string
	LoginPwd   string
	NickName   string
	IsDel      int
	Email      string
	CreateTime time.Time
	QQ         string
	WX         string
	Phone      string
	HeadImage  string
	Gender     int
	RoleID     int
}

//码商临时提交抢单表
type MsTempOrder struct {
	ID              int
	MsUserID        int
	MsQrCodeID      int
	SysPayChannelID int
	MinMoney        float32
	MaxMoney        float32
	CreateTime      time.Time
}

//endregion

//region 商户表
//商户用户
type MerMerchant struct {
	ID         int
	LoginName  string
	LoginPwd   string
	NickName   string
	IsDel      int
	Email      string
	CreateTime time.Time
	QQ         string
	WX         string
	Phone      string
	HeadImage  string
	Gender     int
	RoleID     int
}

// 商户订单表
type MerOrder struct {
	ID int
	//码商ID
	MsUserID int
	//支付渠道ID
	SysPayChannelID int
	//码商二维码ID
	MsQrCodeID int
	//创建时间
	CreateTime time.Time
	//商户ID
	MerMerchantID int
	//订单编号
	OrderNo string
	//订单金额
	Price float32
	//实际支付金额
	TruePrice float32
	//回调地址
	CallBackUrl string
	// 支付完成跳转地址
	ReturnUrl string
	//0支付中 1支付成功 2支付超时 3订单取消 4支付金额不符
	OrderState int
	//回调状态 0未回调 1回调一次 2回调2次 3回调3次  4回调成功
	CallBackState int
	//商品名称
	GoodsName string
	//订单备注
	Remark string
	//订单提交IP
	MerMerchantIP string
}

//endregion

func init() {

	var conn = beego.AppConfig.String("conn")
	// set default database
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)

	// register model
	orm.RegisterModel(new(SysPayChannel), new(MsQrCode), new(MsUser), new(MsTempOrder), new(MerMerchant), new(MerOrder))

	// create table
	orm.RunSyncdb("default", false, true)
}
