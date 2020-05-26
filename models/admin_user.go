package models

type AdminUser struct {
	Id          int    `xorm:"not null pk autoincr comment('id') INT(10)"`
	Username    string `xorm:"not null default '' comment('用户名') unique VARCHAR(16)"`
	Password    string `xorm:"not null comment('密码') VARCHAR(32)"`
	LastLogin   int    `xorm:"comment('注册时间') INT(10)"`
	IsSuperuser int    `xorm:"not null comment('0超级用户，1普通用户') SMALLINT(5)"`
	SysStatus   int    `xorm:"comment('状态，0 正常，1 删除') SMALLINT(5)"`
	SysCreated  int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
}
