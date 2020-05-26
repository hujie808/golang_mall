package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/golang/glog"
	"github.com/kataras/iris"
	"github.com/objcoding/wxpay"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/models"
	"io/ioutil"
	"log"
	"strconv"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/services/wservice"

	"web_iris/golang_mall/web/utils"
)

type WechatPay struct {
	Ctx                     iris.Context
	ServiceMallOrder        services.MallOrderService
	ServiceCommodityService services.MallCommodityService
	ServiceMallOrderInfo    services.MallOrderInfoService
	ServiceMallWechat       services.MallWechatService
	ServiceMallLevel        services.MallLevelService
	//ServiceRetailLog        services.RetailLogService
}

const (
	AckSuccess = `<xml><return_code><![CDATA[SUCCESS]]></return_code></xml>`
	AckFail    = `<xml><return_code><![CDATA[FAIL]]></return_code></xml>`
)

func (c WechatPay) PostTo_pay() {
	result := make(map[string]interface{})
	r_params := make(wxpay.Params)
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	orderId := int(gjson.Get(data, "orderId").Int())
	order := c.ServiceMallOrder.Get(orderId)
	payinfo := wservice.PayInfo{
		OpenId:     openid,
		Body:       "order NO: " + order.OrderNumber,
		OutTradeNo: order.OrderNumber,
		TotalFee:   int64(order.OrderPrice * 100),
	}

	params, err := wservice.CreateUnifiedOrder(payinfo)
	params["timeStamp"] = strconv.Itoa(comm.NowUnix())
	params["signType"] = "MD5"
	params["orderInfo"] = order.OrderNumber
	result["data"] = params
	if err != nil {
		result["code"] = conf.MsgCode
		result["msg"] = "微信支付失败"
		c.Ctx.StatusCode(400)
		c.Ctx.JSON(result)

	} else {
		if params["result_code"] == "SUCCESS" {
			appid := conf.Appid
			r_params.SetString("appId", appid).
				SetString("timeStamp", strconv.Itoa(comm.NowUnix())).
				SetString("nonceStr", params["nonce_str"]).
				SetString("package", "prepay_id="+params["prepay_id"]).
				SetString("signType", "MD5")
			mchid := conf.Mch_id
			apikey := conf.ApiKey
			account := wxpay.NewAccount(appid, mchid, apikey, false)
			client := wxpay.NewClient(account)
			sign := client.Sign(r_params)
			r_params.SetString("paySign", sign)
			result["data"] = r_params
			c.Ctx.JSON(result)
			//}
		} else {
			result["code"] = conf.MsgCode
			result["msg"] = "微信支付失败"
			c.Ctx.StatusCode(400)
			c.Ctx.JSON(result)
		}

	}
}

type CheckPay struct {
	Ctx                     iris.Context
	ServiceMallOrder        services.MallOrderService
	ServiceCommodityService services.MallCommodityService
	ServiceMallOrderInfo    services.MallOrderInfoService
	ServiceMallWechat       services.MallWechatService
	ServiceMallLevel        services.MallLevelService
	ServiceRedLog           services.RedLogService
	ServiceRetailLog        services.RetailLogService
}

//微信回调参数
func (c CheckPay) Any() {
	log.Println("调用支付")
	req := c.Ctx.Request()
	glog.V(8).Info(req)
	w := c.Ctx.ResponseWriter()

	if req.Method != "POST" {
		fmt.Fprint(w, AckFail)
		return
	}

	bodydata, err := ioutil.ReadAll(req.Body)
	if err != nil {
		glog.Error(err)
		fmt.Fprint(w, AckFail)
		return
	}
	log.Println("这里是支付请求参数")
	glog.Info(string(bodydata))
	var wxn utils.WXPayNotify
	err = xml.Unmarshal(bodydata, &wxn)
	if err != nil {
		glog.Error(err)
		fmt.Fprint(w, AckFail)
		return
	}
	if utils.ProcessWX(wxn) {
		//支付成功后回调
		fmt.Println(wxn.OutTradeNo)
		r_order := c.ServiceMallOrder.GetOutTradeNo(wxn.OutTradeNo) //订单
		if r_order.PayStatus!="1"{
			return
		}
		r_order.PayStatus = "2"
		r_order.PayTime = comm.NowUnix()
		r_wechat := c.ServiceMallWechat.Get(r_order.MallWechatId) //购买商品的微信号
		r_wechat.Sum = r_wechat.Sum + float32(r_order.OrderPrice) //增加总消费
		allLevelList := c.ServiceMallLevel.GetAll(0, 30)
		//TODO 如果订单等于 2 已付款 pass

		level_min := float32(0)
		for _, one_level := range allLevelList { //根据总消费提升用户等级
			if (r_wechat.Sum >= float32(one_level.Consumption)) && (float32(one_level.Consumption) > level_min) {
				r_wechat.LevelId = one_level.Id
				level_min = float32(one_level.Consumption)
			}
		}
		//增加分佣
		if r_wechat.UserType == 0 && r_wechat.ShareId >= 1 && float32(r_order.OrderPrice) > 0.01 { //普通用户,有分享人,价格大于1
			share_wechat := c.ServiceMallWechat.Get(r_wechat.ShareId) //微信分享人
			if share_wechat.UserType == 2 {                           //分享人是研发师
				log.Println("已进入分佣系统")
				var r_retail float32                                      //分佣最后的钱
				all_info := c.ServiceMallOrderInfo.GetOrder(r_order.Id)   //根据订单拿到订单详情
				r_order_sub := r_order.OriginalPrice - r_order.OrderPrice //商品总共便宜了多少钱
				order_number := 0                                         //共有多少个商品
				for _, order_info := range all_info {
					order_number = order_number + order_info.Number
				}
				one_commodity_sub := float32(r_order_sub) / float32(order_number) //每个商品便宜了多少钱
				log.Printf("共便宜了%d元", one_commodity_sub)

				for _, order_info := range all_info {
					commodity := c.ServiceCommodityService.Get(order_info.CommodityId)
					log.Println(commodity)
					log.Printf("计算一个商品应该分佣多少:(订单现价%d-(每一个便宜了多少钱%d * 订单数量%d)) * 订单直推分佣比例%d/100", order_info.PayPrice, one_commodity_sub, order_info.Number, commodity.ShareZhitui)
					r_retail = r_retail + ((float32(order_info.PayPrice)-float32(one_commodity_sub))*float32(order_info.Number))*float32(commodity.ShareZhitui)/100 //（（商品价格 -单个便宜的价格） * 个数 ）*分佣比例
				}
				if r_retail<=0{
					return
				}
				share_wechat.RetailPrice = share_wechat.RetailPrice + r_retail
				c.ServiceMallWechat.Update(share_wechat, []string{}) //增加分佣金额
				log.Println(r_retail)
				//增加分佣记录
				retail_err := c.ServiceRetailLog.Create(&models.RetailLog{
					MallWechatId:      r_wechat.Id,
					MallWechatShareId: share_wechat.Id,
					MallOrderId:       r_order.Id,
					MallOrderPrice:    r_order.OriginalPrice,
					OrderRetailPrice:  r_retail,
					AddTime:           int(utils.GetTimestamp()),
				})
				r_wechat.RetailPrice=r_wechat.RetailPrice+r_retail
				if retail_err != nil {
					log.Println("wechat_pay.go PostCheck_wxpay c.ServiceRetailLog.Create err=", retail_err)
				}
			}

		}
		o_err := c.ServiceMallOrder.Update(r_order, []string{})
		if o_err != nil {
			log.Println("wechat_pay.go  PostCheck_wxpay() ServiceMallOrder.Update err=", o_err)
		}

		log.Println("这里是更新后的微信用户", r_wechat)
		w_err := c.ServiceMallWechat.Update(r_wechat, []string{})
		if w_err != nil {
			log.Println("wechat_pay.go PostCheck_wxpay() ServiceMallOrder.Update err=", w_err)
		}

		glog.Info("PROCESSWX SUCCESS")
		fmt.Fprint(w, AckSuccess)
		return

	}

	fmt.Fprint(w, AckFail)
	return
}
