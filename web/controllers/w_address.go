package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

type WAddress struct {
	Ctx                iris.Context
	ServiceMallWechat  services.MallWechatService
	ServiceMallAddress services.MallAddressService
}

func (c *WAddress) Post() { //查看
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	address := c.ServiceMallAddress.GetWechatId(user.Id)
	result["data"] = address
	result["code"] = 200
	c.Ctx.JSON(result)
}

func (c *WAddress) PostAdd() { //添加
	result := make(map[string]interface{})
	var idDefautl int
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	detailAddress := gjson.Get(data, "DetailAddress").String()
	tel := gjson.Get(data, "Tel").String()
	realName := gjson.Get(data, "RealName").String()
	locationAddress := gjson.Get(data, "LocationAddress").String()
	idDefautls := gjson.Get(data, "IdDefautl").Bool()
	if idDefautls {
		idDefautl = 1
	} else {
		idDefautl = 2
	}
	user := c.ServiceMallWechat.GetOpenid(openid)

	address := &models.MallAddress{
		MallWechatId:    user.Id,
		RealName:        realName,
		LocationAddress: locationAddress,
		DetailAddress:   detailAddress,
		Tel:             tel,
		IdDefautl:       idDefautl,
		SysStatus:       0,
	}
	code, msg, r_address := form.MallAddressForm(address, c.ServiceMallAddress)
	if msg == "" {
		err := c.ServiceMallAddress.Create(r_address)
		if err != nil {
			log.Println("w_address.go PostAdd err=", err)
			result["code"] = 301
			c.Ctx.JSON(result)
			return
		}
		result["code"] = code
		result["msg"] = msg
		c.Ctx.JSON(result)

	} else {
		result["code"] = code
		result["msg"] = msg
		c.Ctx.JSON(result)
	}

}

func (c *WAddress) PostUpdate() { //修改
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	openid := gjson.Get(data, "openid").String()
	user := c.ServiceMallWechat.GetOpenid(openid)
	detailAddress := gjson.Get(data, "DetailAddress").String()
	tel := gjson.Get(data, "Tel").String()
	var idDefautl int
	idDefautls := gjson.Get(data, "IdDefautl").Bool()
	if idDefautls {
		idDefautl = 1
	} else {
		idDefautl = 2
	}
	realName := gjson.Get(data, "RealName").String()
	locationAddress := gjson.Get(data, "LocationAddress").String()
	a_id := int(gjson.Get(data, "Id").Int())
	if a_id > 0 {
		address := &models.MallAddress{
			Id:              a_id,
			MallWechatId:    user.Id,
			RealName:        realName,
			LocationAddress: locationAddress,
			DetailAddress:   detailAddress,
			Tel:             tel,
			IdDefautl:       idDefautl,
			SysStatus:       0,
		}
		code, msg, r_address := form.MallAddressForm(address, c.ServiceMallAddress)
		if msg == "" {
			err := c.ServiceMallAddress.Update(r_address, []string{})
			if err != nil {
				log.Println("w_address.go PostAdd err=", err)
				result["code"] = 301
				c.Ctx.JSON(result)
				return
			}
		} else {
			result["code"] = code
			result["msg"] = msg
			c.Ctx.JSON(result)
		}

	}

}

func (c *WAddress) PostDelete() { //删除
	result := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	//openid := gjson.Get(data, "openid").String()
	//user := c.ServiceMallWechat.GetOpenid(openid)
	a_id := int(gjson.Get(data, "Id").Int())
	err := c.ServiceMallAddress.Delete(a_id)
	if err != nil {
		result["code"] = 301
		c.Ctx.JSON(result)
	} else {
		result["code"] = 200
		c.Ctx.JSON(result)
	}

}
