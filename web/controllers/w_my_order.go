package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

type WMyOrder struct {
	Ctx                  iris.Context
	ServiceMallOrder     services.MallOrderService
	ServiceMallOrderInfo services.MallOrderInfoService
	ServiceMallWechat    services.MallWechatService
	ServiceMallAddress   services.MallAddressService
	ServiceMallSku       services.MallSkuService
	ServiceUserCoupons   services.MallUserCouponsService
	ServiceCoupons       services.MallCouponsService
	ServiceMallLevel     services.MallLevelService
	ServiceMallShopping  services.MallShoppingService
	ServiceMallCommodity services.MallCommodityService
}

func (c *WMyOrder) PostSearch() { //订单查询
	result := make(map[string]interface{})
	var orderList []models.MallOrder
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	order_type := int(gjson.Get(data, "order_type").Int())
	orderInfoId := int(gjson.Get(data, "orderInfoId").Int())
	user := c.ServiceMallWechat.GetOpenid(openid)
	if orderInfoId > 0 {
		orderList = c.ServiceMallOrder.SearchUserOrderId(order_type, user.Id, orderInfoId)
	} else {
		orderList = c.ServiceMallOrder.SearchUserOrder(order_type, user.Id)
	}

	orderInfoMap := orderAllInfo(orderList, c.ServiceMallOrderInfo)
	rOrderList := make([]map[string]interface{}, 0)
	for _, order := range orderList {
		rOrder := make(map[string]interface{})
		rOrder["Id"] = order.Id
		rOrder["ClickTime"] = comm.FormatFromUnixTimeShort(int64(order.ClickTime))
		rOrder["PayStatus"] = order.PayStatus
		rOrder["OrderPrice"] = order.OrderPrice
		rOrder["OrderNumber"] = order.OrderNumber
		rOrder["OrderInfoList"] = orderInfoMap[order.Id]
		rOrderList = append(rOrderList, rOrder)
	}
	log.Println()
	result["orderList"] = rOrderList
	c.Ctx.JSON(result)
}

func (c *WMyOrder) PostDelete() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	orderId := int(gjson.Get(data, "Id").Int())
	user := c.ServiceMallWechat.GetOpenid(openid)
	order := c.ServiceMallOrder.Get(orderId)
	err := c.ServiceMallOrderInfo.ReallyDelete(order.Id, user.Id)
	if err != nil {
		log.Println("w_my_order.go PostDelete MallOrderInfo.ReallyDelete err=", err)
	}
	oErr := c.ServiceMallOrder.ReallyDelete(order.Id, user.Id)
	if oErr != nil {
		log.Println("w_my_order.go PostDelete MallOrder.ReallyDelete err=", err)
	}
	result["code"] = 200
	c.Ctx.JSON(result)
}

func (c *WMyOrder) PostCancel() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	orderId := int(gjson.Get(data, "Id").Int())
	order := c.ServiceMallOrder.Get(orderId)
	c.ServiceMallOrder.CancelDelete(order.Id)
	result["code"] = 200
	c.Ctx.JSON(result)
}

func (c *WMyOrder) PostView() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	result["address"] = c.ServiceMallAddress.GetUserDefault(user.Id)
	sumPrice := float32(0)
	sku_id := gjson.Get(data, "sku_id").Array()
	num := gjson.Get(data, "number").Array()
	r_view := make([]map[string]interface{}, 0)
	for index, sku := range sku_id {
		firstSku := c.ServiceMallSku.GetTitle(int(sku.Int()))
		rSku := make(map[string]interface{})
		rSku["Id"] = firstSku.Id
		rSku["Title"] = firstSku.Title
		rSku["CTitle"] = firstSku.CTitle
		rSku["Pice"] = firstSku.Pice
		rSku["MallCommodityId"] = firstSku.MallCommodityId
		rSku["Stock"] = firstSku.Stock
		rSku["number"] = num[index].Int()
		sumPrice = sumPrice + firstSku.Pice*float32(num[index].Int())
		r_view = append(r_view, rSku)
	}
	result["commodity"] = r_view
	result["user_coupons"] = userCouponsMap(c.ServiceUserCoupons.GetUserCoupons(user.Id, comm.NowUnix()))
	result["sumPrice"] = sumPrice //原价多少
	_, result["levelPrice"] = userLevelPrice(user, c.ServiceMallLevel, sumPrice)
	c.Ctx.JSON(result)
}

func (c *WMyOrder) PostAdd() { //生成订单
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	//生成订单 记录分享人
	share_openid := gjson.Get(data, "share_openid").String()
	if len(share_openid) > 10 {
		log.Println("这里是share",share_openid)
		openid := gjson.Get(data, "openid").String()
		user := c.ServiceMallWechat.GetOpenid(openid)
		share_user := c.ServiceMallWechat.GetOpenid(share_openid)
		if user.Id >= 1 &&  share_user.Id >= 1 && share_user.UserType == 2 &&  user.Id != share_user.Id { //没有分享人,分享人是研发师
			user.ShareId = share_user.Id
			c.ServiceMallWechat.Update(user, []string{"share_id"})
		}
	}

	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	address_id := int(gjson.Get(data, "address_id").Int())           //收货地址
	user_coupons_id := int(gjson.Get(data, "user_coupons_id").Int()) //用户所使用的优惠券id
	sumPrice := float32(gjson.Get(data, "sumPrice").Float())         //原价
	sku_id := gjson.Get(data, "sku_id").Array()                      //sku
	num := gjson.Get(data, "number").Array()                         //购买数量
	leave_msg := gjson.Get(data, "leave_msg").Array()                //留言
	address := c.ServiceMallAddress.Get(address_id)
	// 是否使用优惠券
	sumPriceCoupons, msg := userPriceEmploy(user, c.ServiceUserCoupons, c.ServiceCoupons, sumPrice, user_coupons_id) //是否使用优惠券
	if msg != "" {
		result["code"] = conf.MsgCode
		result["msg"] = msg
		c.Ctx.JSON(result)
		return
	}
	//会员打折
	levelPrice, _ := userLevelPrice(user, c.ServiceMallLevel, sumPriceCoupons)
	//创建订单
	order := &models.MallOrder{ //创建订单
		MallWechatId:      user.Id,
		OrderNumber:       "tea_" + strconv.Itoa(comm.Random(1000)) + strconv.Itoa(comm.NowUnix())[1:], //TODO 支付订单 微信版本
		MallOrderInfoId:   "",
		OrderAddress:      address.LocationAddress + address.DetailAddress,
		OrderTel:          address.Tel,
		ClickTime:         comm.NowUnix(),
		SysStatus:         0,
		PayStatus:         "1",
		OriginalPrice:     sumPrice,
		OrderPrice:        levelPrice,
		MallUserCouponsId: user_coupons_id,
	}
	orderId, o_err := c.ServiceMallOrder.CreateId(order)
	if o_err != nil || orderId < 1 {
		log.Println("w_my_order.go PostAdd ServiceMallOrder.Create err=", o_err, "orderId=", orderId)
	}
	for index, sku := range sku_id {
		var isVirtual int
		commoditySku := c.ServiceMallSku.GetTitle(int(sku.Int()))
		if commoditySku.ComType == "xn" {
			isVirtual = 1
		} else {
			isVirtual = 0
		}

		orderInfo := &models.MallOrderInfo{ //创建订单详情
			MallWecahtId:   user.Id,
			WechatHead:     user.Headimgurl,
			WechatNickname: user.Nickname,
			MallOrderId:    orderId,
			Number:         int(num[index].Int()),
			LeaveMsg:       leave_msg[index].String(),
			CommodityId:    commoditySku.MallCommodityId,
			SkuTitle:       commoditySku.Title,
			CommodityTitle: commoditySku.CTitle,
			CommodityImage: commoditySku.CImage,
			SkuPrice:       commoditySku.Pice,
			PayPrice:       commoditySku.Pice,
			CreateTime:     comm.NowUnix(),
			SysStatus:      0,
			IsVirtual:      isVirtual,
		}
		or_err := c.ServiceMallOrderInfo.Create(orderInfo)
		commdity := c.ServiceMallCommodity.Get(commoditySku.MallCommodityId)
		commdity.Sales = commdity.Sales + int(num[index].Int())
		c_err := c.ServiceMallCommodity.Update(commdity, []string{"sales"})
		if c_err != nil {
			log.Println("w_my_order.go PostAdd ServiceMallOrderInfo.Update field err=", c_err)
		}
		if or_err != nil {
			log.Println("w_my_order.go PostAdd ServiceMallOrderInfo.Create err=", or_err)
		}
		err := c.ServiceMallShopping.DeleteShoppingAll(user.Id, commoditySku.Id) //删除购物车
		if err != nil {
			log.Println("删除购物车失败err=:", err)
		}

		stock := commoditySku.Stock - int(num[index].Int())
		err = c.ServiceMallSku.Update(&models.MallSku{Id: commoditySku.Id, Stock: stock}, []string{"stock"}) //减库存
		if err != nil {
			log.Println("更新库存失败err=:", err)
		}

	}
	//TODO 清空购物车 redis

	result["code"] = 200
	result["orderId"] = orderId
	c.Ctx.JSON(result)
}

func (c *WMyOrder) GetComment() { //用户提交评论
	c.Ctx.JSON(`{"ni":"nihao"}`)
}

func (c *WMyOrder) PostComment() { //用户提交评论
	result := make(map[string]interface{})
	var mallOrderId int

	data := c.Ctx.Values().GetString("data")
	orderInfoList := gjson.Get(data, "data").Array()
	for _, orderInfo := range orderInfoList {
		var start int
		infoDict := orderInfo.Map()
		start = int(infoDict["Start"].Int())
		if start == 0 {
			start = 5
		}
		mallOrderId = int(infoDict["MallOrderId"].Int())
		r_orderInfo := &models.MallOrderInfo{
			Id:      int(infoDict["Id"].Int()),
			Comment: infoDict["Comment"].String(),
			Start:   start,
			IsShow:  0,
		}
		code, msg, result_orderInfo := form.WCommoentForm(r_orderInfo)
		if msg != "" {
			c.Ctx.StatusCode(code)
			c.Ctx.WriteString(msg)
			return
		}
		c.ServiceMallOrderInfo.Update(result_orderInfo, []string{"Comment", "Start", "IsShow"})
	}
	//更新订单为已完成
	c.ServiceMallOrder.Update(&models.MallOrder{Id: mallOrderId, PayStatus: "5"}, []string{"SysStatus"})

	result["code"] = 200
	result["msg"] = "后台审核通过后,将会展示评论"
	c.Ctx.JSON(result)
}

func userPriceEmploy(user *models.MallWechat, sUserCoupons services.MallUserCouponsService, sCoupons services.MallCouponsService, sumPrice float32, user_coupons_id int) (float32, string) {
	var rPrice float32
	userCoupons := sUserCoupons.GetUserCoupons(user.Id, comm.NowUnix()) //找到用户的优惠券
	log.Println(userCoupons)
	if user_coupons_id > 0 {
		for _, coupons := range userCoupons {
			if float32(coupons.MaxMoney) <= sumPrice && float32(coupons.IsEmploy) == 1 && coupons.MallCouponsId == user_coupons_id {
				rPrice = sumPrice - float32(coupons.MinMoney)
				log.Printf("要删除优惠券id=%d", coupons.MallUserCouponsId)
				err := sUserCoupons.Update(&models.MallUserCoupons{Id: coupons.MallUserCouponsId, IsEmploy: 0}, []string{"is_employ"})
				if err != nil {
					log.Println("w_my_order.go userPriceEmploy sUserCoupons.Update err=", err)
				}
				return rPrice, ""
			} else {
				rPrice = sumPrice
			}
		}
		return rPrice, "没有可用优惠券"
	} else {
		return sumPrice, ""
	}

}

func userLevelPrice(user *models.MallWechat, sLevel services.MallLevelService, Price float32) (float32, int) {
	var discount int
	if user.LevelId > 0 {
		discount = sLevel.Get(user.LevelId).CommodityPrice
	} else {
		discount = 100
	}
	result := Price * float32(discount) / 100
	return result, discount
}

func userCouponsMap(list []models.JionUserCoupons) []map[string]interface{} { //晒选出未使用
	result := make([]map[string]interface{}, 0)
	for _, v := range list {
		if v.IsEmploy == 1 {
			coupons := make(map[string]interface{})
			coupons["Id"] = v.Id
			coupons["Title"] = v.Title
			coupons["MaxMoney"] = v.MaxMoney
			coupons["MinMoney"] = v.MinMoney
			coupons["EndTime"] = comm.FormatFromUnixTimeShort(int64(v.EndTime))
			result = append(result, coupons)
		}

	}
	return result
}

func orderAllInfo(list []models.MallOrder, sOrderInfo services.MallOrderInfoService) map[int][]models.MallOrderInfo { //晒选出未使用
	//1拿到所有orderId
	//2用所有orderId 找到所属orderinfo
	//3用key是orderId 形成列表
	result := make(map[int][]models.MallOrderInfo, 0)
	orderIdList := func(l []models.MallOrder) []int {
		r := make([]int, 0)
		for _, v := range l {
			r = append(r, v.Id)
		}
		return r
	}(list)
	allOrderInfo := sOrderInfo.GetOrderListInfoId(orderIdList)
	for _, v := range allOrderInfo {
		result[v.MallOrderId] = append(result[v.MallOrderId], v)
	}
	return result
}

func makeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(getDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数

	pow := math.Pow(10, float64(sum)) - 1
	//fmt.Println("sum=>", sum)
	//fmt.Println("pow=>", pow)
	result := strconv.Itoa(rand.Intn(int(pow)))
	count = len(result)
	//fmt.Println("result=>", result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}
func getDaysInYearByThisYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1

	} else {
		total = total + d
	}
	return total;
}
