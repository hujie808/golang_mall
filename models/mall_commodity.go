package models

type MallCommodity struct {
	Id           int     `xorm:"not null pk autoincr comment('id') INT(11)"`
	Title        string  `xorm:"not null comment('商品名称') VARCHAR(200)"`
	Desc         int     `xorm:"comment('排序，1最小') INT(10)"`
	Intro        string  `xorm:"comment('简介') TEXT"`
	Number       string  `xorm:"comment('编号') VARCHAR(32)"`
	CategoryId   int     `xorm:"comment('分类id') INT(11)"`
	Detail       string  `xorm:"not null comment('商品详情') TEXT"`
	Price        float32 `xorm:"comment('展示商品原价价格') FLOAT"`
	PriceNow     float32 `xorm:"comment('商品打折价格') FLOAT"`
	Sales        int     `xorm:"not null default 00000000000 comment('商品销量') INT(11)"`
	Views        int     `xorm:"not null default 00000000000 comment('浏览量') INT(11)"`
	Collect      int     `xorm:"not null comment('收藏') INT(10)"`
	Image        string  `xorm:"not null comment('封面图片') VARCHAR(200)"`
	ImageOne     string  `xorm:"comment('轮播图片一') VARCHAR(200)"`
	ImageTwo     string  `xorm:"comment('轮播图片二') VARCHAR(200)"`
	ImageThree   string  `xorm:"comment('轮播图片三') VARCHAR(200)"`
	ShareXiaji   int     `xorm:"comment('下级分佣比例') SMALLINT(5)"`
	ShareZhitui  int     `xorm:"comment('直推分佣比例') SMALLINT(5)"`
	SysCreated   int     `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated   int     `xorm:"default 0 comment('更新时间') INT(10)"`
	SysStatus    int     `xorm:"not null default 0 comment('状态，0正常，1作废') SMALLINT(5)"`
	CommodityHot string  `xorm:"comment('是否推荐，0推荐 1不推荐') VARCHAR(1)"`
	VideoUrl     string  `xorm:"comment('视频url') VARCHAR(200)"`
	ComType      string  `xorm:"not null default '' comment('sw实物，xn虚拟') VARCHAR(200)"`
}
