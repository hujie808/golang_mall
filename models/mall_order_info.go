package models

type MallOrderInfo struct {
	Id             int     `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallWecahtId   int     `xorm:"not null comment('微信ID') INT(11)"`
	WechatHead     string  `xorm:"comment('微信头像') VARCHAR(200)"`
	WechatNickname string  `xorm:"comment('微信名字') VARCHAR(200)"`
	MallOrderId    int     `xorm:"not null comment('对应订单') index INT(11)"`
	Number         int     `xorm:"not null default 1 comment('商品数量') INT(11)"`
	Comment        string  `xorm:"comment('评论') TEXT"`
	LeaveMsg       string  `xorm:"comment('购买留言') VARCHAR(200)"`
	Start          int     `xorm:"not null comment('星级评分，最高五分') SMALLINT(5)"`
	IsShow         int     `xorm:"comment('状态，0显示，1不显示') SMALLINT(5)"`
	SysStatus      int     `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
	CommodityId    int     `xorm:"comment('商品id') index INT(11)"`
	SkuTitle       string  `xorm:"comment('购买sku的名字') VARCHAR(200)"`
	CommodityTitle string  `xorm:"comment('购买商品名字') VARCHAR(200)"`
	CommodityImage string  `xorm:"comment('购买商品的图片') VARCHAR(200)"`
	SkuPrice       float32 `xorm:"comment('商品价格') FLOAT(11)"`
	PayPrice       float32 `xorm:"comment('购买时价格 优惠券使用之后的价格') FLOAT(11)"`
	CreateTime     int     `xorm:"comment('创建时间') INT(11)"`
	IsVirtual      int     `xorm:"comment('1 虚拟') INT(10)"`
}
