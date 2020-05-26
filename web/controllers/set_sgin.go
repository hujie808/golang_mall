package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/utils"
)

type SetSignatureController struct {
	Ctx              iris.Context
	ServiceAdminUser services.AdminUserService
}

func (c *SetSignatureController) Post() {
	//TODO 拿取数据
	rs := make(map[string]interface{})
	data := comm.GetBody(c.Ctx.Request().Body)
	username := gjson.Get(data, "username").String()
	password := gjson.Get(data, "password").String()
	adminUser := c.ServiceAdminUser.LoginUser(username, password)
	if adminUser != nil {
		loginUser := &models.ObjLoginuser{
			Uid:      adminUser.Id,
			Username: adminUser.Username,
			Password: adminUser.Password,
			Ip:       comm.ClientIP(c.Ctx.Request()),
			Sign:     "",
		}

		uid, sign := comm.SetLoginUser(loginUser)
		if uid >= 0 {
			rs["code"] = 200
			rs["uid"] = uid
			rs["sign"] = sign
		} else {
			rs["code"] = conf.LogBackInCode
		}
	} else {
		rs["code"] = conf.LogBackInCode
	}

	c.Ctx.JSON(rs)
}

func (c *SetSignatureController) Options() {}

func (c *SetSignatureController) PostSaverich() {
	rs := make(map[string]interface{})
	image_url := utils.FileSave(c.Ctx, "file", "rich_txt", 5)
	rs["code"] = 200
	rs["url"] = conf.Host + image_url
	c.Ctx.JSON(rs)
	return
	//c.Ctx.HTML()
	//return
}
func (c *SetSignatureController) OptionsSaverich() {}

func (c *SetSignatureController) PostSaverichvideo() {
	rs := make(map[string]interface{})
	image_url := utils.FileSaveVideo(c.Ctx, "file", "rich_txt", 5)
	rs["code"] = 200
	rs["url"] = conf.Host + image_url
	c.Ctx.JSON(rs)
	return
	//c.Ctx.HTML()
	//return
}
func (c *SetSignatureController) OptionsSaverichvideo() {}
