package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/utils"
)

type WRetail struct {
	Ctx                iris.Context
	ServiceCommodity   services.MallCommodityService
	ServiceCategory    services.MallCategoryService
	ServiceCoupons     services.MallCouponsService
	ServiceUserCoupons services.MallUserCouponsService
	ServiceMallWechat  services.MallWechatService
	ServiceRedLog      services.RedLogService
	ServiceRetailLog   services.RetailLogService
}

func (c *WRetail) Post() { //展示页面
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	result["RetailPrice"] = user.RetailPrice
	result["MinPrice"] = conf.MinPrice
	c.Ctx.JSON(result)
}

func (c *WRetail) Put() { //提现函数
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	price := float32(gjson.Get(data, "price").Int()) //提现金额
	log.Printf("这里是提现金额%f,这里是最小提现金额%f",price,float32(conf.MinPrice))
	if price >= float32(conf.MinPrice) {
		log.Printf("这里是账号余额金额%f",user.RetailPrice)
		if price <= user.RetailPrice {
			user.RetailPrice = user.RetailPrice - price //TODO 这里是提现函数
			c.ServiceMallWechat.Update(user, []string{"retail_price"})
			c.ServiceRedLog.Create(&models.RedLog{ //建立提现红包
				MallWechatId: user.Id,
				MallPrice:    price,
				MallRedOrder: "",
				MallRedBool:  0,
				AddTime:      int(utils.GetTimestamp()),
			})
			result["msg"] = "提现成功,后台审核通过将会给您提现"
			result["RetailPrice"] = user.RetailPrice
			result["MinPrice"] = conf.MinPrice
			c.Ctx.JSON(result)
			return
		} else {
			result["msg"] = "您的提现金额不够"
			c.Ctx.StatusCode(400)
		}

	} else {
		result["msg"] = "不符合最低提现规则"
		c.Ctx.StatusCode(400)
	}

	result["RetailPrice"] = user.RetailPrice
	result["MinPrice"] = conf.MinPrice

	c.Ctx.JSON(result)
}

func (c *WRetail) PostQr() { //二维码页面
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	if user.UserType == 2 {
		result["img_url"] = utils.QrUser(conf.Host, user.Openid)
	} else {
		result["msg"] = "该用户不是研发师"
		result["code"] = 2000
	}

	c.Ctx.JSON(result)
}
