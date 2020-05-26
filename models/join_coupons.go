package models

type JionCoupons struct {
	MallCoupons `xorm:"extends"`
	//MallCommodity `xorm:"extends"`
	CTitle string //mall_coupons.*,mall_commodity.title as c_title
	//MallCommodity_Id string `xorm:"extends"`
	//Titles string
}

func (JionCoupons) TableName() string {
	return "mall_coupons"
}
