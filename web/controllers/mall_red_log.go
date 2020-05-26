package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/utils"
)

//admin 展示注册登录 删除页面
type MallRedLogController struct {
	Ctx                  iris.Context
	ServiceRedLogService services.RedLogService
}

func (c *MallRedLogController) PostPage() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	number:=int(c.ServiceRedLogService.CountAll())
	rs["data"] = c.ServiceRedLogService.GetAll(page, size)
	rs["total"] = number
	c.Ctx.JSON(rs)
}

func (c *MallRedLogController) OptionsPage() {}

func (c *MallRedLogController) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	fmt.Println(data)
	fmt.Println(int(gjson.Get(data, "MallRedBool").Int()))
	fmt.Println(gjson.Get(data, "MallRedBool").Bool())
	form_strcut := &models.RedLog{
		Id:           int(gjson.Get(data, "Id").Int()),
		MallRedOrder: "",
		MallRedBool:  int(gjson.Get(data, "MallRedBool").Int()),
		AddTime:      int(utils.GetTimestamp()),
	}
	err:=c.ServiceRedLogService.Update(form_strcut,[]string{"mall_red_bool"})
	if err !=nil{
		log.Println("mall_red_log.go PostUpdate() ServiceRedLogService.Update err=",err)
	}
	rs["data"] = form_strcut
	c.Ctx.JSONP(rs)
	return
}
func (c *MallRedLogController ) OptionsUpdate() {}
