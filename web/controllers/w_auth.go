package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"strconv"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/services/wservice"
	"web_iris/golang_mall/web/utils"
)

type WechatAuth struct {
	Ctx               iris.Context
	ServiceMallWechat services.MallWechatService
}

type AuthLoginBody struct {
	Code     string               `json:"code"`
	UserInfo wservice.ResUserInfo `json:"userInfo"`
}

func (c *WechatAuth) Post() {

	var alb AuthLoginBody
	//clientIP := comm.ClientIP(c.Ctx.Request())
	data, _ := ioutil.ReadAll(c.Ctx.Request().Body)
	log.Println(alb)
	log.Println(string(data))
	err := json.Unmarshal(data, &alb)
	if err != nil {
		log.Println("w_auth.go json.Unmarshal err=", err)
	}

	userInfo := wservice.Login(alb.Code, alb.UserInfo)
	if userInfo == nil {
	}
	var user *models.MallWechat
	user = c.ServiceMallWechat.GetOpenid(userInfo.OpenID)
	if user.Id == 0 {
		log.Println("没有此对象")
		users := &models.MallWechat{
			Openid:     userInfo.OpenID,
			Nickname:   userInfo.NickName,
			Headimgurl: userInfo.AvatarUrl,
			Sex:        strconv.Itoa(userInfo.Gender),
			City:       userInfo.City,
			Country:    userInfo.Country,
			Province:   userInfo.Province,
			AddTime:    comm.NowUnix(),
			SysStatus:  0,
		}
		err := c.ServiceMallWechat.Create(users)
		user = users
		log.Println("user == nil", user)
		if err != nil {
			c.Ctx.HTML("储存错误错误")
			log.Println("w_auth.go ServiceMallWechat.Create err=", err, "user=", user)
		}
	}

	userinfo := make(map[string]interface{})
	userinfo["id"] = user.Id
	userinfo["nickname"] = user.Nickname
	userinfo["gender"] = user.Sex
	userinfo["Headimgurl"] = user.Headimgurl
	sessionKey := wservice.Create(utils.Int2String(user.Id))
	rtnInfo := make(map[string]interface{})
	rtnInfo["token"] = sessionKey
	rtnInfo["userInfo"] = userinfo
	rtnInfo["openid"] = user.Openid

	//存入分享人



	c.Ctx.JSON(rtnInfo)
	return
}
func (c *WechatAuth) Options() {}
