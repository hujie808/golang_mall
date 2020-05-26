package models

type MallOrder struct {
	Id                int     `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallWechatId      int     `xorm:"not null comment('微信用户id') index INT(11)"`
	OrderNumber       string  `xorm:"comment('订单号') VARCHAR(64)"`
	MallOrderInfoId   string  `xorm:"comment('订单详情id [3,4,5]列表') VARCHAR"`
	OrderAddress      string  `xorm:"not null comment('收货地址') VARCHAR(200)"`
	OrderTel          string  `xorm:"not null comment('收货手机号') VARCHAR(11)"`
	Logistics         string  `xorm:"not null comment('物流公司') VARCHAR(16)"`
	LogisticsNumber   string  `xorm:"not null comment('物流号') VARCHAR(128)"`
	ClickTime         int     `xorm:"not null comment('生成订单时间') INT(11)"`
	PayTime           int     `xorm:"comment('确认交易时间') INT(11)"`
	SysStatus         int     `xorm:"not null default 0 comment('状态 0正常，1取消订单') SMALLINT(5)"`
	PayStatus         string  `xorm:"not null default '1' comment('1 代付款,2 代发货,3待收货,4待评价,5 完成订单,6 取消订单') VARCHAR(5)"`
	OriginalPrice     float32 `xorm:"comment('订单原价') FLOAT"`
	OrderPrice        float32 `xorm:"comment('订单总价格') FLOAT"`
	MallUserCouponsId int     `xorm:"comment('使用的优惠券id') INT(11)"`
}
