package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

type MallWechatControlers struct {
	Ctx               iris.Context
	ServiceMallWechat services.MallWechatService
	ServiceMallLevel  services.MallLevelService
}

//删除用户
func (c *MallWechatControlers) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	err := c.ServiceMallWechat.Delete(id)
	if err != nil {
		rs["code"], rs["msg"] = conf.FailedCode, "Failed"
	} else {
		rs["code"], rs["msg"] = 200, "Succeed"
	}
	c.Ctx.JSONP(rs)
}
func (c *MallWechatControlers) OptionsDelete() {}

//修改用户
func (c *MallWechatControlers) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	user := c.ServiceMallWechat.Get(id)
	if user != nil {
		upUser := &models.MallWechat{
			Id:       user.Id,
			Openid:   user.Openid,
			Integral: int(gjson.Get(data, "Integral").Int()),
			Sum:      float32(gjson.Get(data, "Sum").Float()),
			LevelId:  int(gjson.Get(data, "LevelId").Int()),
			UserType: int(gjson.Get(data, "UserType").Int()),
			ShareId:  int(gjson.Get(data, "ShareId").Int()),
			//RetailPrice:  float32(gjson.Get(data, "RetailPrice").Float()),
		}
		code, msg, wechat := form.MallWechatForm(upUser, c.ServiceMallWechat)
		rs["code"], rs["msg"] = code, msg
		if msg != "" {
			c.Ctx.JSONP(rs)
			return
		}
		err := c.ServiceMallWechat.Update(wechat, []string{"user_type"})

		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, msg, data
			log.Println("mall_wechat.go PostUpdate ServiceMallWechat.Update err=", err)
			c.Ctx.JSONP(rs)
			return
		}
	}

	rs["code"], rs["msg"] = 200, ""
	c.Ctx.JSONP(rs)
}
func (c *MallWechatControlers) OptionsUpdate() {}

//		用户id单页
func (c *MallWechatControlers) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallWechat.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallWechatControlers) OptionsByid() {}

//		用户id单页
func (c *MallWechatControlers) PostByupdate() {
	rs := make(map[string]interface{})
	map_user := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallWechat.Get(uid)
		map_user = comm.StructMapAdd(*user, "", "")
		if user.MyDistributorId > 0 {
			map_user["MyDistributorId"] = c.ServiceMallWechat.Get(user.MyDistributorId).Tel
		}
		//if user.ShareId > 0 {
		//	map_user["ShareId"] = c.ServiceMallWechat.Get(user.ShareId).Tel
		//}
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", map_user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
	}
	all_count := c.ServiceMallLevel.CountAll()
	rs["level"] = c.ServiceMallLevel.GetAll(1, int(all_count))
	c.Ctx.JSONP(rs)
}
func (c *MallWechatControlers) OptionsByupdate() {}

//用户列表页
func (c *MallWechatControlers) PostPage() {
	rs := make(map[string]interface{})
	var userList []models.MallWechat
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	nickname := gjson.Get(data, "searchNickname").String()
	tel := gjson.Get(data, "searchTel").String()
	if nickname == "" && tel == "" {
		userList = c.ServiceMallWechat.GetAll(page, size)
	} else {
		userList = c.ServiceMallWechat.SearchNickname(page, size, nickname, tel)
	}
	result := make([]map[string]interface{}, 0)
	for _, all := range userList { //循环all
		mallWechat := make(map[string]interface{})
		get_id := all.LevelId //拿到所需关联id
		if get_id > 0 {
			levelService := c.ServiceMallLevel.Get(get_id)
			if levelService.Id == get_id {
				mallWechat = comm.StructMapAdd(all, "Level_title", levelService.Levle)
			} else {
				mallWechat = comm.StructMapAdd(all, "Level_title", "")
			}
		} else {
			mallWechat = comm.StructMapAdd(all, "Level_title", "")
		}
		result = append(result, mallWechat)
	}

	if result != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", result
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = c.ServiceMallWechat.CountAll(nickname, tel)
	rs["searchNickname"] = nickname
	rs["searchTel"] = tel
	c.Ctx.JSONP(rs)
}
func (c *MallWechatControlers) OptionsPage() {}
