package models

type MallBanner struct {
	Id              int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	Title           int    `xorm:"not null comment('1, 轮播图 2, 二维码 3, 分享图片,4广告图 ,567,背景图 8，首页分类图片') SMALLINT(5)"`
	Image           string `xorm:"not null comment('图片地址') VARCHAR(200)"`
	Desc            int    `xorm:"comment('排序，1最小') INT(10)"`
	SysCreated      int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysStatus       int    `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
	MallCommodityId int    `xorm:"comment('轮播图对应商品id') INT(11)"`
	BannerName      string `xorm:"comment('首页二级分类图片名字') VARCHAR(200)"`
}
