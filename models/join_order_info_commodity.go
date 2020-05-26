package models

type JoinOrderInfo struct {
	MallOrderInfo `xorm:"extends"`
	VideoUrl      string
	Detail        string
}

func (JoinOrderInfo) TableName() string {
	return "mall_order_info"
}
