package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

//admin 展示注册登录 删除页面
type MallCouponsController struct {
	Ctx                  iris.Context
	ServiceMallCoupons   services.MallCouponsService
	ServiceMallCommodity services.MallCommodityService
}

func (c *MallCouponsController) PostAdd() { //添加用户
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	//id :=int(gjson.Get(data,"id").Int())
	addList := &models.MallCoupons{
		MallCommodityId: int(gjson.Get(data, "MallCommodityId").Int()),
		Title:           gjson.Get(data, "Title").String(),
		MaxMoney:        int(gjson.Get(data, "MaxMoney").Int()),
		MinMoney:        int(gjson.Get(data, "MinMoney").Int()),
		StartTime:       int(gjson.Get(data, "StartTime").Int()) / 1000,
		EndTime:         int(gjson.Get(data, "EndTime").Int()) / 1000,
		Number:          int(gjson.Get(data, "Number").Int()),
		NowNumber:       int(gjson.Get(data, "Number").Int()),
		SysStatus:       0,
	}
	code, msg, adminUser := form.MallCouponsForm(addList) //验证表单
	if msg == "" {
		err := c.ServiceMallCoupons.Create(adminUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.MsgCode, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"] = code, msg
	c.Ctx.JSONP(rs)
}
func (c *MallCouponsController) OptionsAdd() {}

//删除用户
func (c *MallCouponsController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		err := c.ServiceMallCoupons.Delete(id)
		if err != nil {
			rs["code"], rs["msg"] = conf.FailedCode, "Failed"
		} else {
			rs["code"], rs["msg"] = 200, "Succeed"
		}
		c.Ctx.JSONP(rs)
	}

}
func (c *MallCouponsController) OptionsDelete() {}

func (c *MallCouponsController) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")

	id := int(gjson.Get(data, "Id").Int())

	addList := &models.MallCoupons{
		Id:              id,
		Title:           gjson.Get(data, "Title").String(),
		MallCommodityId: int(gjson.Get(data, "MallCommodityId").Int()),
		MaxMoney:        int(gjson.Get(data, "MaxMoney").Int()),
		MinMoney:        int(gjson.Get(data, "MinMoney").Int()),
		StartTime:       int(gjson.Get(data, "StartTime").Int()) / 1000,
		EndTime:         int(gjson.Get(data, "EndTime").Int()) / 1000,
		Number:          int(gjson.Get(data, "Number").Int()),
		NowNumber:       int(gjson.Get(data, "Number").Int()),
	}
	code, msg, adminUser := form.MallCouponsForm(addList) //验证表单
	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	if msg != "" {
		c.Ctx.JSONP(rs)
		return
	}
	err := c.ServiceMallCoupons.Update(adminUser, []string{})
	if err != nil {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, msg, data
		c.Ctx.JSONP(rs)
		return
	}
	rs["code"], rs["msg"] = 200, ""
	c.Ctx.JSONP(rs)
	return
}

func (c *MallCouponsController) OptionsUpdate() {}

//		用户id单页
func (c *MallCouponsController) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallCoupons.Get(uid)
		rs["commodity"] = c.ServiceMallCommodity.Get(user.MallCommodityId).Title
		rs["endtime"] = comm.FormatFromUnixTime(int64(user.EndTime))
		rs["starttime"] = comm.FormatFromUnixTime(int64(user.StartTime))
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallCouponsController) OptionsByid() {}

//		用户id单页
func (c *MallCouponsController) PostByup() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallCoupons.Get(uid)
		rs["commodity"] = c.ServiceMallCommodity.Get(user.MallCommodityId).Title
		rs["endtime"] = comm.FormatFromUnixTime(int64(user.EndTime))
		rs["starttime"] = comm.FormatFromUnixTime(int64(user.StartTime))
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallCouponsController) OptionsByup() {}

//用户列表页
func (c *MallCouponsController) PostPage() {
	rs := make(map[string]interface{})
	var userList []models.JionCoupons
	nowTime := comm.NowUnix()
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	search := gjson.Get(data, "searchTime").String()
	switch search {
	case "start":
		userList = c.ServiceMallCoupons.GetStart(page, size, nowTime)
		rs["total"] = len(c.ServiceMallCoupons.GetStart(0, 0, nowTime))
	case "end":
		userList = c.ServiceMallCoupons.GetEnd(page, size, nowTime)
		rs["total"] = len(c.ServiceMallCoupons.GetEnd(0, 0, nowTime))
	case "online":
		userList = c.ServiceMallCoupons.GetOnline(page, size, nowTime)
		rs["total"] = len(c.ServiceMallCoupons.GetOnline(0, 0, nowTime))
	default:
		userList = c.ServiceMallCoupons.GetTest(page, size)
		rs["total"] = c.ServiceMallCoupons.CountAll()
	}

	if userList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", userList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}

	rs["searchTime"] = search
	c.Ctx.JSONP(rs)
}
func (c *MallCouponsController) OptionsPage() {}
