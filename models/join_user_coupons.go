package models

type JionUserCoupons struct {
	MallUserCoupons `xorm:"extends"`
	//MallCoupons `xorm:"extends"`
	MallCommodityId int    `xorm:"comment('商品id') INT(11)"`
	MallUserCouponsId   int    `xorm:"comment('用户id') INT(11)"`
	MallCouponsId   int    `xorm:"comment('优惠券id') INT(11)"`
	Title           string `xorm:"not null comment('优惠券名称') VARCHAR(32)"`
	MaxMoney        int    `xorm:"not null comment('满多少') INT(11)"`
	MinMoney        int    `xorm:"not null comment('减多少') INT(11)"`
	StartTime       int    `xorm:"not null comment('开始时间') INT(11)"`
	EndTime         int    `xorm:"not null comment('过期时间') INT(11)"`
	Number          int    `xorm:"not null comment('发放数量') INT(11)"`
	NowNumber       int    `xorm:"comment('剩余数量') INT(11)"`
	//MallCommodity_Id string `xorm:"extends"`
	//Titles string
}

func (JionUserCoupons) TableName() string {
	return "mall_user_coupons"
}
