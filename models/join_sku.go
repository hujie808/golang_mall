package models

type JionSku struct {
	MallSku `xorm:"extends"`
	CTitle  string
	CImage  string
	ComType string //TODO com_type
}

func (JionSku) TableName() string {
	return "mall_sku"
}
