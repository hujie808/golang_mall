package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/services"
)

type WArticle struct {
	Ctx                iris.Context
	ServiceMallArticle services.MallArticleService
}

func (c *WArticle) PostWz() {
	result := make(map[string]interface{})
	//data := c.Ctx.Values().GetString("data")
	rList := c.ServiceMallArticle.GetAllWz(1, 20)
	jsonList := comm.StructMapHold(rList, "Id", "Title", "Intro", "AType", "Image", "AddDate", "Order","LikeNumber")
	result["data"] = jsonList
	c.Ctx.JSON(result)
}

func (c *WArticle) PostYfs() {
	result := make(map[string]interface{})
	//data := c.Ctx.Values().GetString("data")
	rList := c.ServiceMallArticle.GetAllYfs(1, 20)
	jsonList := comm.StructMapHold(rList, "Id", "Title", "Intro", "AType", "Image", "AddDate", "Order", "ALevel","LikeNumber")
	result["data"] = jsonList
	c.Ctx.JSON(result)
}

func (c *WArticle) PostByid() {
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


func (c *WArticle) PostLike() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallArticle.Get(uid)
		user.LikeNumber=user.LikeNumber+1
		err:=c.ServiceMallArticle.Update(user,[]string{"like_number"})
		if err !=nil{
			log.Println("w_article.go PostLike Update err=",err)
		}
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
