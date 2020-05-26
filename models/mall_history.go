package models

type MallHistory struct {
	Id           int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallWechatId int    `xorm:"not null comment('微信id') INT(11)"`
	ViewHistory  string `xorm:"not null comment('最近浏览历史，json{mall_commodity_id:时间戳}') TEXT"`
	SysStatus    int    `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
}
