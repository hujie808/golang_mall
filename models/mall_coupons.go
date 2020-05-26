package models

type MallCoupons struct {
	Id              int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallCommodityId int    `xorm:"comment('商品id') INT(11)"`
	Title           string `xorm:"not null comment('优惠券名称') VARCHAR(32)"`
	MaxMoney        int    `xorm:"not null comment('满多少') INT(11)"`
	MinMoney        int    `xorm:"not null comment('减多少') INT(11)"`
	StartTime       int    `xorm:"not null comment('开始时间') INT(11)"`
	EndTime         int    `xorm:"not null comment('过期时间') INT(11)"`
	Number          int    `xorm:"not null comment('发放数量') INT(11)"`
	NowNumber       int    `xorm:"comment('剩余数量') INT(11)"`
	SysStatus       int    `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
}
