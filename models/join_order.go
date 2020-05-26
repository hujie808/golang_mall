package models

type JionOrder struct {
	MallOrder `xorm:"extends"`
	Nickname  string
	Tel       string
}

func (JionOrder) TableName() string {
	return "mall_order"
}
