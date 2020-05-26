package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

type WCommodity struct {
	Ctx                  iris.Context
	ServiceCommodity     services.MallCommodityService
	ServiceCategory      services.MallCategoryService
	ServiceCoupons       services.MallCouponsService
	ServiceMallOrderInfo services.MallOrderInfoService
	ServiceMallSku       services.MallSkuService
	ServiceMallWechat    services.MallWechatService
	ServiceMallShopping  services.MallShoppingService
	ServiceMallHistory   services.MallHistoryService
}

func (c *WCommodity) Post() {
	//销量 价格排序 综合排序 搜索 分类
	data := c.Ctx.Values().GetString("data")
	commodity_list := make([]models.MallCommodity, 0)
	//sort :=gjson.Get(data,"sort")//1,综合,2,销量,3,价格
	search := gjson.Get(data, "search").String()       //搜索
	category := int(gjson.Get(data, "category").Int()) //分类
	category_list := []int{}//所有分类
	if category > 0 {
		category_list = append(category_list, category)//搜搜分类放到list
		s_category := c.ServiceCategory.Get(category)//拿到搜索分类
		category_parent_list := c.ServiceCategory.GetParentId(s_category.Id)//找到分类的父id
		if category_parent_list == nil {
			commodity_list = c.ServiceCommodity.GetCategory(category)
		} else {
			category_list = append(category_list, category_parent_list...)
			commodity_list = c.ServiceCommodity.GetCategoryAll(category_list)
		}
	}
	if search != "" {
		_,commodity_list = c.ServiceCommodity.GetSearchPage(0, 0, search)
	}
	c.Ctx.JSON(commodity_list)
}
