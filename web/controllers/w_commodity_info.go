package controllers

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func (c *WCommodity) PostInfo() {
	result := make(map[string]interface{}, 0)
	nowTime := comm.NowUnix()
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	c_id := int(gjson.Get(data, "commodity_id").Int())
	commodity := c.ServiceCommodity.Get(c_id)
	historySave(c.ServiceMallHistory, user.Id, commodity.Id) //历史记录储存
	if commodity.Id > 0 {
		result["commodity"] = commodity
		result["coupons"] = c.ServiceCoupons.GetOnline(0, 0, nowTime)
		r_orderInfo := make([]models.MallOrderInfo, 0) //有评论的订单
		listStart := []int{}
		orderInfoList := c.ServiceMallOrderInfo.GetCommodity(c_id)
		for _, orderInfo := range orderInfoList {
			if orderInfo.Comment != "" {
				r_orderInfo = append(r_orderInfo, orderInfo)
			}
			if orderInfo.Start > 0 {
				listStart = append(listStart, orderInfo.Start)
			}
		}
		good_numer := sumStart(listStart)
		if good_numer > 0 {
			result["good"] = fmt.Sprintf("好评率 %d%%", good_numer)
		} else {
			result["good"] = fmt.Sprintf("还未有用户评论此商品")
		}
		commentList := make([]map[string]interface{}, 0)
		for _, v := range r_orderInfo {
			commentLists := make(map[string]interface{}, 0)
			commentLists["Nickname"] = v.WechatNickname
			commentLists["WechatHead"] = v.WechatHead
			commentLists["Comment"] = v.Comment
			commentLists["SkuTitle"] = v.SkuTitle
			commentLists["CreateTime"] = comm.FormatFromUnixTimeShort(int64(v.CreateTime))
			commentList = append(commentList, commentLists)
		}
		result["Stock"] = c.ServiceMallSku.GetCommodityStock(commodity.Id)
		result["commentList"] = commentList       //评论
		result["commentCount"] = len(r_orderInfo) //评论数量
		result["commodity_Sku"] = c.ServiceMallSku.GetCommodityId(commodity.Id)
		collection := c.ServiceMallShopping.GetIsCollection(user.Id, commodity.Id)

		if len(collection) == 1 {
			if collection[0].SysStatus == 0 {
				result["is_collection"] = 1
			} else {
				result["is_collection"] = 0
			}

		} else {
			result["is_collection"] = 0
		}
	}

	//TODO 历史记录ridis 记录历史信息
	c.Ctx.JSON(result)
}

//收藏
func (c *WCommodity) PostCollection() {
	result := make(map[string]interface{}, 0)
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	c_id := int(gjson.Get(data, "commodity_id").Int())
	commodity := c.ServiceCommodity.Get(c_id)
	get_collection := gjson.Get(data, "collection").Bool()
	fmt.Printf("%T,%s", get_collection, get_collection)
	collection := c.ServiceMallShopping.GetIsCollection(user.Id, c_id)
	if len(collection) == 1 {
		collection := collection[0]
		fmt.Println(collection.Id)
		if get_collection == false {
			fmt.Println("取消收藏")
			c.ServiceMallShopping.Update(&models.MallShopping{Id: collection.Id, SysStatus: 1}, []string{"sys_status"})
			commodity.Collect = commodity.Collect + 1
			c.ServiceCommodity.Update(commodity, []string{})
		} else {
			collection.SysStatus = 0
			fmt.Println(collection)
			c.ServiceMallShopping.Update(&collection, []string{"sys_status"})
			commodity.Collect = commodity.Collect - 1
			c.ServiceCommodity.Update(commodity, []string{})
		}
	} else {
		shopping := &models.MallShopping{
			MallWechatId: user.Id,
			Type:         2,
			MallSkuId:    c_id,
			SysStatus:    0,
		}
		err := c.ServiceMallShopping.Create(shopping)
		commodity.Collect = commodity.Collect + 1
		c.ServiceCommodity.Update(commodity, []string{})
		log.Println("w_commodit_info ServiceMallShopping.Create err=", err)
	}
	result["code"] = 200
	result["collection"] = collection
	c.Ctx.JSON(result)
}

//查看收藏
func (c *WCommodity) PostCollectionView() {
	result := make(map[string]interface{}, 0)
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	idListCommodity := c.ServiceMallShopping.GetCollection(user.Id)
	commodityList := c.ServiceCommodity.GetCommodityAll(idListCommodity)
	r_commodityList := make([]map[string]interface{}, 0)
	for _, commodity := range commodityList {
		r_comm := make(map[string]interface{}, 0)
		r_comm["Id"] = commodity.Id
		r_comm["Title"] = commodity.Title
		r_comm["Image"] = commodity.Image
		r_comm["PriceNow"] = commodity.PriceNow
		r_comm["Price"] = commodity.Price
		r_commodityList = append(r_commodityList, r_comm)
	}
	result["code"] = 200
	result["data"] = r_commodityList
	c.Ctx.JSON(result)
}

//删除收藏
func (c *WCommodity) PostViewCollectionDelete() {
	result := make(map[string]interface{}, 0)
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		c.ServiceMallShopping.DeleteShopping(id)
		result["code"] = 200
		c.Ctx.JSON(result)
	} else {
		result["code"] = conf.FailedCode
		c.Ctx.JSON(result)
	}
}

func sumStart(intArr []int) int {
	sum := 0
	for _, val := range intArr {
		//累计求和
		sum += val
	}
	//平均值保留到2位小数
	average := float64(sum) / float64(len(intArr))

	return int(average * 2 * 10)
}
