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
type AdminController struct {
	Ctx              iris.Context
	ServiceAdminUser services.AdminUserService
}

func (c *AdminController) Get() { //添加用户
	c.Ctx.HTML("nihao")
}

func (c *AdminController) PostAdd() { //添加用户
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	username := gjson.Get(data, "username").String()
	addUser := &models.AdminUser{
		Username:    username,
		Password:    gjson.Get(data, "password").String(),
		LastLogin:   comm.NowUnix(),
		IsSuperuser: int(gjson.Get(data, "is_superuser").Int()),
		SysStatus:   0,
		SysCreated:  comm.NowUnix(),
	}
	code, msg, adminUser := form.AddAdminUserForm(addUser) //验证表单
	if msg == "" {
		err := c.ServiceAdminUser.Create(adminUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.MsgCode, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	c.Ctx.JSONP(rs)
}
func (c *AdminController) OptionsAdd() {}

//删除用户
func (c *AdminController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "uid").Int())
	err := c.ServiceAdminUser.Delete(id)
	if err != nil {
		rs["code"], rs["msg"] = conf.FailedCode, "Failed"
	} else {
		rs["code"], rs["msg"] = 200, "Succeed"
	}
	c.Ctx.JSONP(rs)
}
func (c *AdminController) OptionsDelete() {}

//修改用户
func (c *AdminController) PostUpdate() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "id").Int())
	user := c.ServiceAdminUser.Get(id)
	if user != nil {
		upUser := &models.AdminUser{
			Id:          user.Id,
			Username:    gjson.Get(data, "username").String(),
			Password:    gjson.Get(data, "password").String(),
			LastLogin:   comm.NowUnix(),
			IsSuperuser: int(gjson.Get(data, "is_superuser").Int()),
			SysStatus:   0,
			SysCreated:  0,
		}
		code, msg, adminUser := form.AddAdminUserForm(upUser)
		rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
		if msg != "" {
			c.Ctx.JSONP(rs)
			return
		}
		err := c.ServiceAdminUser.Update(upUser, []string{"mall_commodity_id"})
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, msg, data
			c.Ctx.JSONP(rs)
			return
		}
	}

}
func (c *AdminController) OptionsUpdate() {}

//用户列表页
func (c *AdminController) PostPage() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	userList := c.ServiceAdminUser.GetAll(page, size)
	if userList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", userList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = c.ServiceAdminUser.CountAll()
	c.Ctx.JSONP(rs)
}
func (c *AdminController) OptionsPage() {}

//		用户id单页
func (c *AdminController) PostUserid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "uid").Int())
	if uid > 0 {
		user := c.ServiceAdminUser.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *AdminController) OptionsUserid() {}
