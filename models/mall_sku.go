package models

type MallSku struct {
	Id              int     `xorm:"not null pk autoincr comment('id') INT(11)"`
	MallCommodityId int     `xorm:"not null comment('商品id') INT(11)"`
	Title           string  `xorm:"not null comment('名称') VARCHAR(64)"`
	Pice            float32 `xorm:"not null comment('价格') FLOAT(11)"`
	Images          string  `xorm:"not null comment('图片') VARCHAR(200)"`
	Stock           int     `xorm:"not null comment('库存') INT(10)"`
	SysStatus       int     `xorm:"not null comment('状态，0正常，1废除') SMALLINT(5)"`
}
