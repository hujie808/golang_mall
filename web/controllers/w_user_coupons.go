package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

type WUserCoupons struct {
	Ctx                iris.Context
	ServiceCommodity   services.MallCommodityService
	ServiceCategory    services.MallCategoryService
	ServiceCoupons     services.MallCouponsService
	ServiceUserCoupons services.MallUserCouponsService
	ServiceMallWechat  services.MallWechatService
}

func (c *WUserCoupons) Post() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	couponse := c.ServiceCoupons.GetOnline(0, 0, comm.NowUnix())
	log.Println(c.ServiceCoupons.Get(2))
	userCoupons := c.ServiceUserCoupons.GetWechatId(user.Id)
	r_coupons_list := make([]map[string]interface{}, 0)
	for _, coupon := range couponse {
		r_coupons := make(map[string]interface{})
		r_coupons["user_get"] = true
		r_coupons["ID"] = coupon.MallCoupons.Id
		r_coupons["Title"] = coupon.MallCoupons.Title
		r_coupons["MinMoney"] = coupon.MinMoney
		r_coupons["MaxMoney"] = coupon.MaxMoney
		r_coupons["NowNumber"] = coupon.NowNumber
		r_coupons["EndTime"] = comm.FormatFromUnixTimeShort(int64(coupon.EndTime))
		for _, user_c := range userCoupons {
			if coupon.NowNumber > 0 {
				if coupon.Id == user_c.MallCouponsId {
					r_coupons["user_get"] = false
					r_coupons_list = append(r_coupons_list, r_coupons)
					break
				} else {
					r_coupons["user_get"] = true
				}
			}
		}
		if r_coupons["user_get"] == true {
			log.Println(r_coupons_list, r_coupons, "外循环添加")
			r_coupons_list = append(r_coupons_list, r_coupons)
		}
	}
	result["data"] = r_coupons_list
	c.Ctx.JSON(result)
}

func (c *WUserCoupons) PostMy() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	nowTime := comm.NowUnix()
	userCoupons := c.ServiceUserCoupons.GetUserCoupons(user.Id, nowTime)
	result["data"] = userCoupons
	c.Ctx.JSON(result)
}

func (c *WUserCoupons) PostAdd() {
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	coupons_id := int(gjson.Get(data, "coupons_id").Int())
	user := c.ServiceMallWechat.GetOpenid(openid)
	user_coupons := c.ServiceUserCoupons.GetCouponsId(user.Id, coupons_id)
	if coupons_id > 0 {
		if user_coupons.Id > 0 {
			result["code"] = conf.MsgCode
			result["msg"] = "您已领取过此优惠券~"
			c.Ctx.JSON(result)
			return
		} else {
			coupons := c.ServiceCoupons.Get(coupons_id)
			if coupons.NowNumber > 0 {
				cCoupons := &models.MallUserCoupons{
					MallWechatId:  user.Id,
					MallCouponsId: coupons_id,
					IsEmploy:      1,
				}
				err := c.ServiceUserCoupons.Create(cCoupons)
				if err != nil {
					log.Println("w_user_coupons.go PostAdd err=", err)
					return
				}
				coupons.NowNumber--
				c_err := c.ServiceCoupons.Update(coupons, []string{})
				if c_err != nil {
					log.Println("w_user_coupons.go ServiceCoupons.Update err=", err)
					return
				}
			} else {
				result["code"] = 300
				result["msg"] = "优惠券已经领取完毕,请勿非法领取"
				c.Ctx.JSON(result)
				return
			}

		}
	}
	result["code"] = 200
	c.Ctx.JSON(result)
}
