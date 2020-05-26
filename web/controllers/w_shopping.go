package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

type WShopping struct {
	Ctx                 iris.Context
	ServiceCommodity    services.MallCommodityService
	ServiceMallWechat   services.MallWechatService
	ServiceMallShopping services.MallShoppingService
	ServiceMallSku      services.MallSkuService
}

func (c *WShopping) Post() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	if user.Id > 0 {
		r_commodity_list := make([]map[string]interface{}, 0)
		sku_list := c.ServiceMallShopping.GetShopping(user.Id)
		shopping_list := c.ServiceMallShopping.GetShoppingList(user.Id)
		rCommodity := c.ServiceMallSku.GetIdAllCommodityTitle(sku_list)
		for _, v := range shopping_list {
			for _, commodity := range rCommodity {
				if v.MallSkuId == commodity.Id {
					r_commodity := make(map[string]interface{})
					r_commodity["Id"] = v.Id
					r_commodity["Stock"] = commodity.Stock
					r_commodity["Title"] = commodity.Title
					r_commodity["MallCommodityId"] = commodity.MallCommodityId
					r_commodity["SkuId"] = commodity.MallSku.Id
					r_commodity["CTitle"] = commodity.CTitle
					r_commodity["Images"] = commodity.Images
					r_commodity["Pice"] = commodity.Pice
					r_commodity["Number"] = v.ShoppingNumber
					r_commodity_list = append(r_commodity_list, r_commodity)
				}
			}
		}
		result["data"] = r_commodity_list
	}
	//TODO 缓存购物车 订单生成清空购物车

	result["code"] = 200
	c.Ctx.JSON(result)
}

func (c *WShopping) PostAdd() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	sku_id := int(gjson.Get(data, "sku_id").Int())
	number := int(gjson.Get(data, "number").Int())
	user := c.ServiceMallWechat.GetOpenid(openid)
	shopping_sku := c.ServiceMallShopping.GetIsShopping(user.Id, sku_id) //用户 sku的购物车
	sku := c.ServiceMallSku.Get(sku_id)
	if number <= sku.Stock {
		if shopping_sku.Id > 0 { //如果购物车里有
			if (shopping_sku.ShoppingNumber + number) <= sku.Stock {
				shopping_sku.ShoppingNumber = shopping_sku.ShoppingNumber + number
				c.ServiceMallShopping.Update(shopping_sku, []string{})
				result["code"] = 200
				c.Ctx.JSON(result)
				return
			} else {
				c.Ctx.StatusCode(400)
				result["code"] = 400
				c.Ctx.JSON(result)
				return
			}

		} else { //没有则创建
			shopping := &models.MallShopping{
				MallWechatId:   user.Id,
				Type:           1,
				MallSkuId:      sku_id,
				ShoppingNumber: number,
			}
			c.ServiceMallShopping.Create(shopping)
		}
		result["code"] = 200
		c.Ctx.JSON(result)
		return
	} else {
		c.Ctx.StatusCode(400)
		result["code"] = 400
		c.Ctx.JSON(result)
	}

}

func (c *WShopping) PostDelete() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	shopping_id := int(gjson.Get(data, "shopping_id").Int())
	if shopping_id == 0 {
		shopping_id = 1
	}
	detele_all := int(gjson.Get(data, "detele_all").Int())
	user := c.ServiceMallWechat.GetOpenid(openid)
	if shopping_id > 0 {
		c.ServiceMallShopping.DeleteShopping(shopping_id)
	}
	if detele_all == 1 {
		c.ServiceMallShopping.DeleteShopping(user.Id)
	}
	result["code"] = 200
	c.Ctx.JSON(result)
}
