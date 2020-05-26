package models

type MallAddress struct {
	Id              int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallWechatId    int    `xorm:"not null comment('微信用户id') INT(11)"`
	RealName        string `xorm:"not null comment('收货真实姓名') VARCHAR(12)"`
	LocationAddress string `xorm:"not null comment('位置') VARCHAR(128)"`
	DetailAddress   string `xorm:"not null comment('详细地址') TEXT"`
	Tel             string `xorm:"not null comment('收货手机号') VARCHAR(11)"`
	IdDefautl       int    `xorm:"default 2 comment('默认地址，1是，2不是') SMALLINT(5)"`
	SysStatus       int    `xorm:"not null comment('删除，0是，1不是') SMALLINT(5)"`
}
