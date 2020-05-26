package models

type MallWechat struct {
	Id              int     `xorm:"not null pk autoincr comment('id') INT(11)"`
	Openid          string  `xorm:"not null comment('openid') unique VARCHAR(50)"`
	Nickname        string  `xorm:"not null comment('微信姓名') VARCHAR(16)"`
	Headimgurl      string  `xorm:"not null comment('微信头像，url') VARCHAR(255)"`
	Sex             string  `xorm:"comment('性别') VARCHAR(50)"`
	City            string  `xorm:"comment('城市') VARCHAR(50)"`
	Country         string  `xorm:"comment('国家') VARCHAR(50)"`
	Province        string  `xorm:"comment('省') VARCHAR(50)"`
	AddTime         int     `xorm:"comment('注册时间,字符串') INT(11)"`
	Integral        int     `xorm:"comment('虚拟积分') INT(11)"`
	Sum             float32 `xorm:"comment('总消费') INT(10)"`
	LevelId         int     `xorm:"comment('级别 id') INT(10)"`
	MyDistributorId int     `xorm:"comment('我的研发师 id') INT(11)"`
	ShareId         int     `xorm:"comment('我的分享人 _id') INT(11)"`
	Tel             string  `xorm:"comment('手机号') VARCHAR(11)"`
	SysStatus       int     `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
	UserType        int     `xorm:"not null comment('0普通用户，2研发者') INT(10)"`
	RetailPrice     float32 `xorm:"not null comment('提现金额') INT(10)"`
}
