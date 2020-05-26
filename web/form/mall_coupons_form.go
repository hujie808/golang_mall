package form

import (
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func MallCouponsForm(coupons *models.MallCoupons) (int, string, *models.MallCoupons) {
	if len(coupons.Title) < 1 {
		return conf.MsgCode, "标题不得小于1", nil
	}
	if coupons.EndTime < comm.NowUnix() {
		return conf.MsgCode, "过期时间不得小于现在时间", nil
	}
	if coupons.Number < 1 {
		return conf.MsgCode, "发放优惠券不得小于1", nil
	}
	if coupons.MinMoney < 1 {
		return conf.MsgCode, "减价不得小于0", nil
	}
	if coupons.MaxMoney < coupons.MinMoney {
		return conf.MsgCode, "满多少元不得大于减多少元", nil
	}
	if coupons.MallCommodityId < 1 {
		coupons.MallCommodityId = 1 //TODO 有会员不能指定商品 默认是1
		//return conf.MsgCode,"商品id不正确",nil
	}

	return 200, "", coupons
}
