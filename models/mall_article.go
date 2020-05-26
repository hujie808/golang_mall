package models

type MallArticle struct {
	Id         int    `xorm:"not null pk autoincr comment('id') INT(10)"`
	Title      string `xorm:"not null comment('标题') VARCHAR(200)"`
	Intro      string `xorm:"not null comment('文章简介') VARCHAR(255)"`
	AType      string `xorm:"not null comment('文章类型 yfs 研发师 wz 文章') VARCHAR(200)"`
	Content    string `xorm:"not null comment('文章内容') TEXT"`
	Image      string `xorm:"not null comment('图片') VARCHAR(200)"`
	AddDate    int    `xorm:"comment('添加时间') INT(10)"`
	Order      int    `xorm:"comment('排序') INT(10)"`
	ALevel     string `xorm:"comment('研发师级别') VARCHAR(15)"`
	LikeNumber int    `xorm:"comment('点赞数量') INT(10)"`
}
