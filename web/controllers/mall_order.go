package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

//订单
type MallOrderController struct {
	Ctx                      iris.Context
	ServiceMallOrder         services.MallOrderService
	ServiceMallOrderInfo     services.MallOrderInfoService
	ServiceMallWechatService services.MallWechatService
}

//删除订单
func (c *MallOrderController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		err := c.ServiceMallOrder.Delete(id)
		if err != nil {
			rs["code"], rs["msg"] = conf.FailedCode, "Failed"
		} else {
			rs["code"], rs["msg"] = 200, "Succeed"
		}
		c.Ctx.JSONP(rs)
	}

}
func (c *MallOrderController) OptionsDelete() {}

func (c *MallOrderController) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	orderInfo := gjson.Get(data, "order_info").Array()
	for _, info := range orderInfo {
		infoDict := info.Map()
		addInfo := &models.MallOrderInfo{
			Id:     int(infoDict["Id"].Int()),
			IsShow: int(infoDict["IsShow"].Int()),
		}
		err := c.ServiceMallOrderInfo.Update(addInfo, []string{"is_show"})
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "是否显示错误", data
			c.Ctx.JSONP(rs)
			return
		}

	}

	id := int(gjson.Get(data, "Id").Int())
	user := c.ServiceMallOrder.Get(id)
	if user != nil {
		addList := &models.MallOrder{
			Id:              int(gjson.Get(data, "Id").Int()),
			Logistics:       gjson.Get(data, "Logistics").String(),
			LogisticsNumber: gjson.Get(data, "LogisticsNumber").String(),
			PayStatus:       gjson.Get(data, "PayStatus").String(),
			OrderAddress:    gjson.Get(data, "OrderAddress").String(),
			OrderTel:        gjson.Get(data, "OrderTel").String(),
			OrderPrice:      float32(gjson.Get(data, "OrderPrice").Float()),
		}
		if user.OrderPrice != addList.OrderPrice {
			addList.OrderNumber="tea_" + strconv.Itoa(comm.Random(1000)) + strconv.Itoa(comm.NowUnix())[1:]}

		code, msg, adminUser := form.MallOrderForm(addList)
		rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
		if msg != "" {
			c.Ctx.JSONP(rs)
			//return
		}
		err := c.ServiceMallOrder.Update(adminUser, []string{"mall_commodity_id"})
		if err != nil {
			rs["code"], rs["msg"] = conf.FailedCode, msg
			c.Ctx.JSONP(rs)
			return
		}

	}
	rs["code"], rs["msg"] = 200, ""
	c.Ctx.JSONP(rs)
	return
}
func (c *MallOrderController) OptionsUpdate() {}

//		用户id单页
func (c *MallOrderController) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallOrder.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallOrderController) OptionsByid() {}

func (c *MallOrderController) PostByup() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		var result = make(map[string]interface{})
		order := c.ServiceMallOrder.Get(uid)
		orderInfoList := c.ServiceMallOrderInfo.GetOrder(order.Id)
		log.Println(orderInfoList)
		if orderInfoList != nil {
			r, _ := json.Marshal(orderInfoList)
			result = comm.StructMapAdd(*order, "order_info", string(r))
		}
		log.Println(result)
		if result != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", result
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	} else {
		rs["code"] = conf.FailedCode
		c.Ctx.JSONP(rs)
	}
}

func (c *MallOrderController) OptionsByup() {

}

//用户列表页
func (c *MallOrderController) PostPage() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	var orderList []models.JionOrder
	var total int
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	pay_status := gjson.Get(data, "PayStatus").String()
	tel := gjson.Get(data, "Tel").String()
	orderList, total = c.ServiceMallOrder.SearchPageIs(page, size, pay_status, tel)
	if orderList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", orderList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = total
	rs["PayStatus"] = pay_status
	rs["Tel"] = tel
	c.Ctx.JSONP(rs)
}
func (c *MallOrderController) OptionsPage() {}
