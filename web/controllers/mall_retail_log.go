package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"golang_mall/conf"
	"golang_mall/services"
)

//admin 展示注册登录 删除页面
type MallRetailController struct {
	Ctx              iris.Context
	ServiceRetailLog services.RetailLogService
}


func (c *MallRetailController) PostPage() {
	rs := make(map[string]interface{})
	var number int
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	searchOrder := gjson.Get(data, "searchOrder").String()
	if searchOrder==""{
		rs["data"] = c.ServiceRetailLog.GetAll(page, size)
		number=int(c.ServiceRetailLog.CountAll())
	}
	rs["total"] = number
	c.Ctx.JSON(rs)
}
func (c *MallRetailController) OptionsPage() {}


func (c *MallRetailController) PostDetele() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		err := c.ServiceRetailLog.Delete(id)
		if err != nil {
			rs["code"], rs["msg"] = conf.FailedCode, "Failed"
		} else {
			rs["code"], rs["msg"] = 200, "Succeed"
		}
		c.Ctx.JSONP(rs)
	}
}

func (c *MallRetailController) OptionsDetele() {}
