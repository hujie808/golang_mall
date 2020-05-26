package models

type RetailLog struct {
	Id                int     `xorm:"not null pk autoincr comment('id') INT(10)"`
	MallWechatId      int     `xorm:"not null comment('购买商品，微信用户') INT(11)"`
	MallWechatShareId int     `xorm:"not null comment('微信分享人Id') INT(11)"`
	MallOrderId       int     `xorm:"not null comment('订单id') INT(11)"`
	MallOrderPrice    float32 `xorm:"not null comment('订单价格') FLOAT"`
	OrderRetailPrice  float32 `xorm:"not null comment('分佣金额') FLOAT"`
	AddTime           int     `xorm:"not null comment('分佣时间记录') INT(11)"`
}
