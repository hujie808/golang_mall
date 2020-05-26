package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"time"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

//admin 展示注册登录 删除页面
type MallArticleController struct {
	Ctx                iris.Context
	ServiceMallArticle services.MallArticleService
}

func (c *MallArticleController) PostAdd() { //添加用户
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	log.Println(gjson.Get(data, "AType").String())
	addList := &models.MallArticle{
		Title:      gjson.Get(data, "Title").String(),
		Intro:      gjson.Get(data, "Intro").String(),
		AType:      gjson.Get(data, "AType").String(),
		Content:    gjson.Get(data, "Content").String(),
		Image:      gjson.Get(data, "Image").String(),
		AddDate:    int(time.Now().Unix()),
		Order:      int(gjson.Get(data, "Consumption").Int()),
		ALevel:     gjson.Get(data, "ALevel").String(),
		LikeNumber: 0,
	}
	code, msg, adminUser := form.MallArticleForm(addList) //验证表单
	log.Println(addList)
	if msg == "" {
		err := c.ServiceMallArticle.Create(adminUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.MsgCode, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"] = code, msg
	c.Ctx.JSONP(rs)
}
func (c *MallArticleController) OptionsAdd() {}

//删除用户
func (c *MallArticleController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		err := c.ServiceMallArticle.Delete(id)
		if err != nil {
			rs["code"], rs["msg"] = conf.FailedCode, "Failed"
		} else {
			rs["code"], rs["msg"] = 200, "Succeed"
		}
		c.Ctx.JSONP(rs)
	}

}
func (c *MallArticleController) OptionsDelete() {}

func (c *MallArticleController) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")

	id := int(gjson.Get(data, "Id").Int())
	user := c.ServiceMallArticle.Get(id)
	if user != nil {
		addList := &models.MallArticle{
			Id:         user.Id,
			Title:      gjson.Get(data, "Title").String(),
			Intro:      gjson.Get(data, "Intro").String(),
			AType:      gjson.Get(data, "AType").String(),
			Content:    gjson.Get(data, "Content").String(),
			Image:      gjson.Get(data, "Image").String(),
			AddDate:    int(time.Now().Unix()),
			Order:      int(gjson.Get(data, "Consumption").Int()),
			ALevel:     gjson.Get(data, "ALevel").String(),
			LikeNumber: 0,
		}
		code, msg, adminUser := form.MallArticleForm(addList)
		rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
		if msg != "" {
			c.Ctx.JSONP(rs)
			return
		}
		err := c.ServiceMallArticle.Update(adminUser, []string{"mall_commodity_id"})
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

func (c *MallArticleController) OptionsUpdate() {}

//		用户id单页
func (c *MallArticleController) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallArticle.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallArticleController) OptionsByid() {}

//用户列表页
func (c *MallArticleController) PostPage() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	log.Println(data)
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	userList := c.ServiceMallArticle.GetAll(page, size)
	log.Println(page,size)
	log.Println(len(userList))
	log.Println(11111111111)
	if userList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", userList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = c.ServiceMallArticle.CountAll()
	c.Ctx.JSONP(rs)
}
func (c *MallArticleController) OptionsPage() {}
