package form

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func MallOrderForm(order *models.MallOrder) (int, string, *models.MallOrder) {
	if len(order.PayStatus) > 2 {
		return conf.MsgCode, "状态设置错误", nil
	}
	//TODO 这设置添加总消费额度

	//if len(order.Logistics) < 1 {
	//	return conf.MsgCode, "物流公司不得小于1", nil
	//}
	//if len(order.LogisticsNumber) < 1 {
	//	return conf.MsgCode, "物流单号不得小于1", nil
	//}
	return 200, "", order
}
