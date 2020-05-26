package models

type MallUserCoupons struct {
	Id            int `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallWechatId  int `xorm:"not null comment('微信id') index INT(11)"`
	MallCouponsId int `xorm:"not null comment('优惠券id') index INT(11)"`
	IsEmploy      int `xorm:"default 1 comment('0使用，1，未使用') SMALLINT(5)"`
}
