package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

//admin 展示注册登录 删除页面
type MallLevelController struct {
	Ctx              iris.Context
	ServiceMallLevel services.MallLevelService
}

func (c *MallLevelController) PostAdd() { //添加用户
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	addList := &models.MallLevel{
		Levle:          gjson.Get(data, "Levle").String(),
		Consumption:    int(gjson.Get(data, "Consumption").Int()),
		ShareZhitui:    int(gjson.Get(data, "ShareZhitui").Int()),
		ShareXiaji:     int(gjson.Get(data, "ShareXiaji").Int()),
		CommodityPrice: int(gjson.Get(data, "CommodityPrice").Int()),
		SysStatus:      0,
	}
	code, msg, adminUser := form.MallLevelForm(addList) //验证表单

	if msg == "" {
		err := c.ServiceMallLevel.Create(adminUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.MsgCode, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"] = code, msg
	c.Ctx.JSONP(rs)
}
func (c *MallLevelController) OptionsAdd() {}

//删除用户
func (c *MallLevelController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		err := c.ServiceMallLevel.Delete(id)
		if err != nil {
			rs["code"], rs["msg"] = conf.FailedCode, "Failed"
		} else {
			rs["code"], rs["msg"] = 200, "Succeed"
		}
		c.Ctx.JSONP(rs)
	}

}
func (c *MallLevelController) OptionsDelete() {}

func (c *MallLevelController) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")

	id := int(gjson.Get(data, "Id").Int())
	user := c.ServiceMallLevel.Get(id)
	if user != nil {
		addList := &models.MallLevel{
			Id:             id,
			Levle:          gjson.Get(data, "Levle").String(),
			Consumption:    int(gjson.Get(data, "Consumption").Int()),
			ShareZhitui:    int(gjson.Get(data, "ShareZhitui").Int()),
			ShareXiaji:     int(gjson.Get(data, "ShareXiaji").Int()),
			CommodityPrice: int(gjson.Get(data, "CommodityPrice").Int()),
			SysStatus:      0,
		}
		code, msg, adminUser := form.MallLevelForm(addList)
		rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
		if msg != "" {
			c.Ctx.JSONP(rs)
			return
		}
		err := c.ServiceMallLevel.Update(adminUser, []string{"mall_commodity_id"})
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, msg, data
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"] = 200, ""
	c.Ctx.JSONP(rs)
	return
}
func (c *MallLevelController) OptionsUpdate() {}

//		用户id单页
func (c *MallLevelController) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallLevel.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallLevelController) OptionsByid() {}

//用户列表页
func (c *MallLevelController) PostPage() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	userList := c.ServiceMallLevel.GetAll(page, size)
	if userList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", userList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = c.ServiceMallLevel.CountAll()
	c.Ctx.JSONP(rs)
}
func (c *MallLevelController) OptionsPage() {}
