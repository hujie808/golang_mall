package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"strings"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
	"web_iris/golang_mall/web/utils"
)

//admin 展示注册登录 删除页面
type MallCategoryController struct {
	Ctx                     iris.Context
	ServiceMallCategory     services.MallCategoryService
	ServiceCommodityService services.MallCommodityService
}

func (c *MallCategoryController) PostAdd() {
	var code int
	var msg string
	var adminUser *models.MallCategory
	rs := make(map[string]interface{})
	image_url := utils.FileSave(c.Ctx, "files[]", "category", 3)
	addUser := &models.MallCategory{
		Name:        c.Ctx.FormValue("Name"),
		Desc:        comm.StringGetInt(c.Ctx.FormValue("Desc")),
		Image:       image_url,
		SysStatus:   0,
		ParentId:    comm.StringGetInt(c.Ctx.FormValue("ParentId")),
		CategoryHot: c.Ctx.FormValue("CategoryHot"),
	}
	forced := c.Ctx.FormValue("forced_updating")
	if forced == "true" {
		code, msg, adminUser = 200, "", &models.MallCategory{}
	} else {
		code, msg, adminUser = form.MallCategoryForm(addUser, c.ServiceMallCategory, c.ServiceCommodityService) //验证表单
	}

	if msg == "" {
		err := c.ServiceMallCategory.Create(addUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.MsgCode, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	c.Ctx.JSON(rs)
}
func (c *MallCategoryController) OptionsAdd() {}

func (c *MallCategoryController) PostByup() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	categoryName := gjson.Get(data, "name").String()
	if id > 0 {
		rs["data"] = c.ServiceMallCategory.Get(id)
		if categoryName != "" {
			categoryList := c.ServiceMallCategory.GetName(categoryName)
			r_category := make([]map[string]interface{}, 0)
			for _, c := range categoryList {
				category := make(map[string]interface{})
				category["key"] = c.Id
				category["title"] = c.Name
				category["children"] = "[]"
				category["Image"] = conf.Host + c.Image
				r_category = append(r_category, category)
			}
			rs["Category"] = r_category
		} else {
			rs["Category"] = utils.GetAllCategory(c.ServiceMallCategory)
		}

	} else {
		if categoryName != "" {
			rs["Category"] = c.ServiceMallCategory.GetName(categoryName)
			categoryList := c.ServiceMallCategory.GetName(categoryName)
			r_category := make([]map[string]interface{}, 0)
			for _, c := range categoryList {
				category := make(map[string]interface{})
				category["key"] = c.Id
				category["title"] = c.Name
				category["children"] = "[]"
				category["Image"] = conf.Host + c.Image
				r_category = append(r_category, category)
			}
			rs["Category"] = r_category
		} else {
			rs["Category"] = utils.GetAllCategory(c.ServiceMallCategory)
		}
	}
	rs["code"] = 200
	c.Ctx.JSONP(rs)
	return
}

func (c *MallCategoryController) OptionsByup() {}

func (c *MallCategoryController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	category := c.ServiceMallCategory.Get(id)
	rs["code"], rs["msg"] = form.MallCategoryDeleteForm(category, c.ServiceMallCategory, c.ServiceCommodityService)
	if rs["code"] == 200 {
		if id > 0 {
			err := c.ServiceMallCategory.Delete(id)
			if err != nil {
				rs["code"], rs["msg"] = conf.FailedCode, "Failed"
			} else {
			}
			c.Ctx.JSONP(rs)
			return
		}
	}
	c.Ctx.JSONP(rs)
	return

}
func (c *MallCategoryController) OptionsDelete() {}

func (c *MallCategoryController) PostUpdate() {
	var code int
	var msg string
	var adminUser *models.MallCategory
	rs := make(map[string]interface{})
	addUser := &models.MallCategory{
		Id:          comm.StringGetInt(c.Ctx.FormValue("Id")),
		Name:        c.Ctx.FormValue("Name"),
		Desc:        comm.StringGetInt(c.Ctx.FormValue("Desc")),
		SysStatus:   0,
		ParentId:    comm.StringGetInt(c.Ctx.FormValue("ParentId")),
		CategoryHot: c.Ctx.FormValue("CategoryHot"),
	}
	image :=c.Ctx.FormValue("Image")
	if len(image)>1{
		imageList:=strings.Split(image,"public")
		addUser.Image="/public"+imageList[len(imageList)-1]
	}else{
		image_url := utils.FileSave(c.Ctx, "files[]", "category", 3)
		addUser.Image=image_url
	}

	forced := c.Ctx.FormValue("forced_updating")
	if forced == "true" {
		code, msg, adminUser = 200, "", &models.MallCategory{}
	} else {
		code, msg, adminUser = form.MallCategoryForm(addUser, c.ServiceMallCategory, c.ServiceCommodityService) //验证表单
	}
	if msg == "" {
		err := c.ServiceMallCategory.Update(adminUser,[]string{})
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = conf.MsgCode, "创建失败", ""
			c.Ctx.JSONP(rs)
			return
		}
	}
	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	c.Ctx.JSONP(rs)

	return

}
func (c *MallCategoryController) OptionsUpdate() {}
