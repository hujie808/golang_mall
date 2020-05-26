package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/services"
)

type WMyPage struct {
	Ctx              iris.Context
	ServiceCommodity services.MallCommodityService
	//ServiceCategory      services.MallCategoryService
	ServiceCoupons       services.MallCouponsService
	ServiceMallOrderInfo services.MallOrderInfoService
	//ServiceMallSku       services.MallSkuService
	ServiceMallWechat services.MallWechatService
	ServiceMallLevel  services.MallLevelService
	//ServiceMallShopping  services.MallShoppingService
	ServiceMallHistory     services.MallHistoryService
	ServiceMallUserCoupons services.MallUserCouponsService
}

func (c *WMyPage) Post() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	r_user := make(map[string]interface{})
	r_user["Headimgurl"] = user.Headimgurl
	r_user["Nickname"] = user.Nickname
	r_user["Integral"] = user.Integral
	r_user["Sum"] = user.Sum
	r_user["Tel"] = user.Tel
	r_user["MyDistributorId"] = user.MyDistributorId
	r_user["ShareId"] = user.ShareId
	r_user["UserType"] = user.UserType       //是否是研发师
	r_user["MinPrice"] = conf.MinPrice       //最小提现金额
	r_user["RetailPrice"] = user.RetailPrice //可提现金额
	userLeveId := user.LevelId
	if userLeveId > 0 {
		r_user["LevelName"] = c.ServiceMallLevel.Get(userLeveId).Levle
	} else {
		r_user["LevelName"] = "普通会员"
	}
	//未使用的个人购物券
	r_user["all_Video_count"] = len(c.ServiceMallOrderInfo.GetVideoAll(user.Id)) //购买视频数量
	r_user["CouponsNumber"] = len(c.ServiceMallUserCoupons.GetUserCoupons(user.Id,comm.NowUnix()))
	r_user["History"] = historyGetFive(c.ServiceMallHistory, c.ServiceCommodity, user.Id)
	result["code"] = 200
	result["data"] = r_user
	c.Ctx.JSON(result)
}

func (c *WMyPage) PostHistory() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	history := historyGetAll(c.ServiceMallHistory, c.ServiceCommodity, user.Id)
	result["data"] = history
	result["code"] = 200
	c.Ctx.JSON(result)
}

func (c *WMyPage) PostVideo() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	all_Video := c.ServiceMallOrderInfo.GetVideoAll(user.Id)
	result["data"] = all_Video
	result["code"] = 200
	c.Ctx.JSON(result)
}
