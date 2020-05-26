package models

type RedLog struct {
	Id           int     `xorm:"not null pk autoincr comment('id') INT(10)"`
	MallWechatId int     `xorm:"not null comment('微信提现提现用户id') INT(11)"`
	MallPrice    float32 `xorm:"not null comment('提现金额') FLOAT"`
	MallRedOrder string  `xorm:"comment('红包订单号') VARCHAR(64)"`
	MallRedBool  int     `xorm:"not null comment('红包是否到账 0，未到账 1到账') INT(10)"`
	AddTime      int     `xorm:"not null comment('提现时间') INT(11)"`
}
