package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"

	"fmt"
	"log"
	"strconv"
	"web_iris/golang_mall/web/utils"
)

type MallBannerController struct {
	Ctx               iris.Context
	ServiceMallBanner services.MallBannerService
}

func (c *MallBannerController) Get() {
	c.Ctx.HTML("这是轮播图")
}

func (c *MallBannerController) PostAdd() { //添加
	image_url := utils.FileSave(c.Ctx, "files[]", "banner", 5)
	rs := make(map[string]interface{})
	//data := c.Ctx.Values().GetString("data")

	title_id := c.Ctx.FormValue("title")
	title, err := strconv.Atoi(title_id)
	if err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}

	desc_id := c.Ctx.FormValue("desc")
	desc, err := strconv.Atoi(desc_id)
	if err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}

	mall_commodity_id := c.Ctx.FormValue("mall_commodity_id")
	banner_name := c.Ctx.FormValue("banner_name")
	mall_commodity, err := strconv.Atoi(mall_commodity_id)
	if err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}

	addUser := &models.MallBanner{
		Title:           title,
		Image:           image_url,
		Desc:            desc,
		SysCreated:      comm.NowUnix(),
		SysStatus:       0,
		MallCommodityId: mall_commodity,
		BannerName:      banner_name,
	}
	code, msg, adminUser := form.BannerForm(addUser) //验证表单
	if msg == "" {
		err := c.ServiceMallBanner.Create(adminUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = code, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	c.Ctx.JSONP(rs)
}
func (c *MallBannerController) OptionsAdd() {}

//删除
func (c *MallBannerController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")

	id := int(gjson.Get(data, "Id").Int())
	err := c.ServiceMallBanner.Delete(id)
	if err != nil {
		rs["code"], rs["msg"] = conf.FailedCode, "Failed"
	} else {
		rs["code"], rs["msg"] = 200, "Succeed"
	}
}
func (c *MallBannerController) OptionsDelete() {}

//修改用户
func (c *MallBannerController) PostUpdate() {
	rs := make(map[string]interface{})
	c.Ctx.FormFile("files[]")
	image_url := utils.FileSave(c.Ctx, "files[]", "banner", 5)
	//data := c.Ctx.Values().GetString("data")

	title_id := c.Ctx.FormValue("title")
	title, err := strconv.Atoi(title_id)
	if err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}

	desc_id := c.Ctx.FormValue("desc")
	desc, err := strconv.Atoi(desc_id)
	if err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}
	id, i_err := strconv.Atoi(c.Ctx.FormValue("Id"))
	if i_err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}

	user := c.ServiceMallBanner.Get(id)
	mall_commodity_id := c.Ctx.FormValue("mall_commodity_id")
	mall_commodity, err := strconv.Atoi(mall_commodity_id)
	if err != nil {
		log.Println("admin_banner.gp PostAdd strconv.Atoi err=", err, title)
	}
	banner_name := c.Ctx.FormValue("banner_name")


	if user != nil {
		upUser := &models.MallBanner{
			Id:              user.Id,
			Title:           title,
			Desc:            desc,
			SysCreated:      comm.NowUnix(),
			SysStatus:       0,
			MallCommodityId: mall_commodity,
			BannerName:      banner_name,
		}
		if upUser.Image==""{
			upUser.Image=user.Image
		}else {
			upUser.Image=image_url
		}
		fmt.Println(user, upUser)
		code, msg, adminUser := form.BannerForm(upUser)
		rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
		err := c.ServiceMallBanner.Update(upUser, []string{})
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = code, "更新失败", ""
			c.Ctx.JSONP(rs)
			return
		}
		if msg != "" {
			c.Ctx.JSONP(rs)
			return
		}
		//err := c.ServiceMallBanner.Update(upUser, []string{})
		//if err != nil {
		//	rs["code"], rs["msg"], rs["data"] = conf.FailedCode, msg, data
		//	c.Ctx.JSONP(rs)
		//	return
		//}
	}
}
func (c *MallBannerController) OptionsUpdate() {}

//		用户id单页
func (c *MallBannerController) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallBanner.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallBannerController) OptionsByid() {}

//用户列表页
func (c *MallBannerController) PostPage() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	userList := c.ServiceMallBanner.GetAll(page, size)
	if userList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", userList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = c.ServiceMallBanner.CountAll()
	c.Ctx.JSONP(rs)
}
func (c *MallBannerController) OptionsPage() {}
