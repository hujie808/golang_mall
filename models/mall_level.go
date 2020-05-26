package models

type MallLevel struct {
	Id             int    `xorm:"not null pk autoincr comment('id') INT(10)"`
	Levle          string `xorm:"not null comment('等级设置') VARCHAR(32)"`
	Consumption    int    `xorm:"comment('总消费，满元成为会员') INT(11)"`
	ShareZhitui    int    `xorm:"comment('分享-直推分佣比例%') SMALLINT(3)"`
	ShareXiaji     int    `xorm:"comment('分享-下级分佣比例%') SMALLINT(3)"`
	CommodityPrice int    `xorm:"comment('用户打折价格%') SMALLINT(3)"`
	SysStatus      int    `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
}
