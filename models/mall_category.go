package models

type MallCategory struct {
	Id          int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	Name        string `xorm:"not null comment('分类名') VARCHAR(500)"`
	Desc        int    `xorm:"not null comment('排序，1最小') INT(11)"`
	Image       string `xorm:"comment('分类图片') VARCHAR(100)"`
	SysStatus   int    `xorm:"not null comment('状态，0 正常，1 删除') SMALLINT(5)"`
	ParentId    int    `xorm:"comment('父分类，id') index INT(11)"`
	CategoryHot string `xorm:"comment('是否推荐，0推荐 1不推荐') VARCHAR(1)"`
}
