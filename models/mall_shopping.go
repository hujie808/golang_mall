package models

type MallShopping struct {
	Id             int `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallWechatId   int `xorm:"not null comment('微信用户id') INT(11)"`
	Type           int `xorm:"not null default 1 comment('类型，1购物车，2，收藏') SMALLINT(5)"`
	MallSkuId      int `xorm:"not null comment('sku_id') INT(11)"`
	SysStatus      int `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
	ShoppingNumber int `xorm:"default 1 comment('购物车默认数量') INT(11)"`
}
