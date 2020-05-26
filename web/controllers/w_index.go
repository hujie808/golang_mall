package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
	"web_iris/golang_mall/services"
)

type WIndex struct {
	Ctx               iris.Context
	ServiceBanner     services.MallBannerService
	ServiceCommodity  services.MallCommodityService
	ServiceCategory   services.MallCategoryService
	ServiceMallWechat services.MallWechatService
}

func (c *WIndex) Post() {
	rtnInfo := make(map[string]interface{})

	//data := c.Ctx.Values().GetString("data")//TODO 首页记录分享人 移动至订单生成页面

	rtnInfo["banner"] = c.ServiceBanner.GetTitle(1, 3, 1)
	//rtnInfo["index_category"] = c.ServiceBanner.GetTitle(1, 5, 8)
	rtnInfo["index_category"] = c.ServiceCategory.GetName("视频专区")[0].Id
	s_indexOne := c.ServiceBanner.GetTitle(1, 1, 4)
	if len(s_indexOne) > 0 {
		rtnInfo["indexOne"] = s_indexOne[0].Image
	}
	s_indexTwo := c.ServiceBanner.GetTitle(1, 1, 5)
	if len(s_indexTwo) > 0 {
		rtnInfo["indexTwo"] = s_indexTwo[0].Image
	}
	s_indexThree := c.ServiceBanner.GetTitle(1, 1, 6)
	if len(s_indexThree) > 0 {
		rtnInfo["indexThree"] = s_indexThree[0].Image
	}
	s_indexFour := c.ServiceBanner.GetTitle(1, 1, 7)
	if len(s_indexFour) > 0 {
		rtnInfo["indexFour"] = s_indexFour[0].Image
	}
	categorys := c.ServiceCategory.GetHot(1, 3)
	if len(categorys) < 3 {
		categorys = c.ServiceCategory.GetAll(1, 3)
	}
	for index, category := range categorys {
		category_name := fmt.Sprintf("category_%d", index+1)
		commodity_name := fmt.Sprintf("category_%d_commodity", index+1)
		rtnInfo[category_name] = category
		commList:=c.ServiceCommodity.GetCategory(category.Id)
			if len(commList) <=6 {
				rtnInfo[commodity_name] = commList
			}else {

				rtnInfo[commodity_name] =commList[:6]
			}
	}
	commdity := c.ServiceCommodity.GetRandomHot()
	if len(commdity) >6 {
		rtnInfo["r_commdity"] = commdity
	} else {
		rtnInfo["r_commdity"] = c.ServiceCommodity.GetAllIndex(1, 6)
	}
	log.Println(rtnInfo["r_commdity"])
	c.Ctx.JSON(rtnInfo)
}
