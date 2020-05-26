package controllers

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"strings"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
	"web_iris/golang_mall/web/utils"
)

type MallCommodityController struct {
	Ctx                  iris.Context
	ServiceMallCommodity services.MallCommodityService
	ServiceMallCategory  services.MallCategoryService
	ServiceMallSku       services.MallSkuService
}

func (c *MallCommodityController) PostAdd() { //添加
	//c.Ctx.FormFile()
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	Image := strings.Replace(gjson.Get(data, "Image").String(), conf.Host, "", 1)
	ImageOne := strings.Replace(gjson.Get(data, "ImageOne").String(), conf.Host, "", 1)
	ImageTwo := strings.Replace(gjson.Get(data, "ImageTwo").String(), conf.Host, "", 1)
	ImageThree := strings.Replace(gjson.Get(data, "ImageThree").String(), conf.Host, "", 1)
	addUser := &models.MallCommodity{
		Title:        gjson.Get(data, "Title").String(),
		Desc:         int(gjson.Get(data, "Desc").Int()),
		Intro:        gjson.Get(data, "Intro").String(),
		Number:       gjson.Get(data, "Number").String(),
		CategoryId:   int(gjson.Get(data, "CategoryId").Int()),
		Detail:       gjson.Get(data, "Detail").String(),
		Price:        float32(gjson.Get(data, "Price").Float()),
		PriceNow:     float32(gjson.Get(data, "PriceNow").Float()),
		Sales:        int(gjson.Get(data, "Sales").Int()),
		Views:        int(gjson.Get(data, "Views").Int()),
		Collect:      int(gjson.Get(data, "Collect").Int()),
		Image:        Image,
		ImageOne:     ImageOne,
		ImageTwo:     ImageTwo,
		ImageThree:   ImageThree,
		ShareXiaji:   int(gjson.Get(data, "ShareXiaji").Int()),
		ShareZhitui:  int(gjson.Get(data, "ShareZhitui").Int()),
		SysCreated:   comm.NowUnix(),
		SysUpdated:   comm.NowUnix(),
		SysStatus:    0,
		CommodityHot: gjson.Get(data, "CommodityHot").String(),
		VideoUrl:     gjson.Get(data, "VideoUrl").String(),
		ComType:      gjson.Get(data, "ComType").String(),
	}
	code, msg, adminUser := form.CommodityForm(addUser, c.ServiceMallCategory) //验证表单
	if msg == "" {
		num, err := c.ServiceMallCommodity.CreateId(adminUser)
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = code, "创建失败", ""
			c.Ctx.JSON(rs)
			return
		} else {
			log.Println(num, "is num")
			s_code, s_msg := utils.SaveSku(gjson.Get(data, "SkuList").Array(), num, c.ServiceMallSku)
			if s_msg != "" {
				rs["code"], rs["msg"] = s_code, s_msg
				c.Ctx.JSON(rs)
				return
			}
		}

	}

	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	c.Ctx.JSON(rs)
}
func (c *MallCommodityController) OptionsAdd() {}

func (c *MallCommodityController) PostByup() { //添加
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	id := int(gjson.Get(data, "Id").Int())
	if id > 0 {
		rs["data"] = c.ServiceMallCommodity.Get(id)
		rs["SkuList"] = c.ServiceMallSku.GetCommodityId(id)
		rs["Category"] = utils.GetAllCategory(c.ServiceMallCategory)
	} else {
		rs["Category"] = utils.GetAllPageCategory(c.ServiceMallCategory)
	}
	rs["code"] = 200
	c.Ctx.JSONP(rs)
	return

}

func (c *MallCommodityController) OptionsByup() {}

func (c *MallCommodityController) PostSaveimage() { //添加图片
	//rs := make(map[string]interface{})
	//s :=c.Ctx.Request().Body
	//log.Println(c.Ctx.FormFile("uploadFile"))
	filetype := c.Ctx.FormValue("filetype")
	image_url := utils.FileSave(c.Ctx, "uploadFile", filetype, 5)
	c.Ctx.HTML(image_url)
	return
}
func (c *MallCommodityController) OptionsSaveimage() {}

//上传视频单独接口
func (c *MallCommodityController) PostSavevideo() { //添加图片
	//rs := make(map[string]interface{})
	//s :=c.Ctx.Request().Body
	//log.Println(c.Ctx.FormFile("uploadFile"))
	filetype := c.Ctx.FormValue("filetype")
	image_url := utils.FileSaveVideo(c.Ctx, "uploadFile", filetype, 5)
	c.Ctx.HTML(image_url)
	return
}
func (c *MallCommodityController) OptionsSavevideo() {}

//
//删除
func (c *MallCommodityController) PostDelete() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")

	id := int(gjson.Get(data, "Id").Int())
	if id==0{//如果没有id 等于批量删除
		idlist:=gjson.Get(data, "IdList").Array()
		for _,i :=range idlist{
			oId:=int(i.Int())
			c.ServiceMallCommodity.Delete(oId)
		}
		rs["code"], rs["msg"] = 200, "Succeed"
		return
	}
	err := c.ServiceMallCommodity.Delete(id)
	if err != nil {
		rs["code"], rs["msg"] = conf.FailedCode, "Failed"
	} else {
		rs["code"], rs["msg"] = 200, "Succeed"
	}
}
func (c *MallCommodityController) OptionsDelete() {}

//修改用户
func (c *MallCommodityController) PostUpdate() { //修改
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	Image := strings.Replace(gjson.Get(data, "Image").String(), conf.Host, "", 1)
	ImageOne := strings.Replace(gjson.Get(data, "ImageOne").String(), conf.Host, "", 1)
	ImageTwo := strings.Replace(gjson.Get(data, "ImageTwo").String(), conf.Host, "", 1)
	ImageThree := strings.Replace(gjson.Get(data, "ImageThree").String(), conf.Host, "", 1)
	addUser := &models.MallCommodity{
		Id:           int(gjson.Get(data, "Id").Int()),
		Title:        gjson.Get(data, "Title").String(),
		Desc:         int(gjson.Get(data, "Desc").Int()),
		Intro:        gjson.Get(data, "Intro").String(),
		Number:       gjson.Get(data, "Number").String(),
		CategoryId:   int(gjson.Get(data, "CategoryId").Int()),
		Detail:       gjson.Get(data, "Detail").String(),
		Price:        float32(gjson.Get(data, "Price").Float()),
		PriceNow:     float32(gjson.Get(data, "PriceNow").Float()),
		Sales:        int(gjson.Get(data, "Sales").Int()),
		Views:        int(gjson.Get(data, "Views").Int()),
		Collect:      int(gjson.Get(data, "Collect").Int()),
		Image:        strings.Replace(Image, conf.Host, "", 1),
		ImageOne:     strings.Replace(ImageOne, conf.Host, "", 1),
		ImageTwo:     strings.Replace(ImageTwo, conf.Host, "", 1),
		ImageThree:   strings.Replace(ImageThree, conf.Host, "", 1),
		ShareXiaji:   int(gjson.Get(data, "ShareXiaji").Int()),
		ShareZhitui:  int(gjson.Get(data, "ShareZhitui").Int()),
		SysCreated:   comm.NowUnix(),
		SysUpdated:   comm.NowUnix(),
		SysStatus:    0,
		CommodityHot: gjson.Get(data, "CommodityHot").String(),
		VideoUrl:     gjson.Get(data, "VideoUrl").String(),
		ComType:      gjson.Get(data, "ComType").String(),
	}
	code, msg, adminUser := form.CommodityForm(addUser, c.ServiceMallCategory) //验证表单
	if msg == "" {
		err := c.ServiceMallCommodity.Update(addUser, []string{})
		if err != nil {
			rs["code"], rs["msg"], rs["data"] = code, "创建失败", ""
			c.Ctx.JSON(rs)
			return
		} else {
			s_code, s_msg := utils.SaveSku(gjson.Get(data, "SkuList").Array(), addUser.Id, c.ServiceMallSku)
			if s_msg != "" {
				rs["code"], rs["msg"] = s_code, s_msg
				c.Ctx.JSON(rs)
				return
			}
		}

	}

	rs["code"], rs["msg"], rs["data"] = code, msg, adminUser
	c.Ctx.JSON(rs)
}
func (c *MallCommodityController) OptionsUpdate() {}

//		用户id单页
func (c *MallCommodityController) PostByid() {
	rs := make(map[string]interface{})
	data := c.Ctx.Values().GetString("data")
	uid := int(gjson.Get(data, "Id").Int())
	if uid > 0 {
		user := c.ServiceMallCommodity.Get(uid)
		if user != nil {
			rs["code"], rs["msg"], rs["data"] = 200, "Succeed", user
		} else {
			rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
		}
		c.Ctx.JSONP(rs)
	}
}
func (c *MallCommodityController) OptionsByid() {}

//用户列表页
func (c *MallCommodityController) PostPage() {
	var number int
	rs := make(map[string]interface{})
	var userList []models.MallCommodity
	data := c.Ctx.Values().GetString("data")
	page := int(gjson.Get(data, "page").Int())
	size := int(gjson.Get(data, "size").Int())
	searchTitle := gjson.Get(data, "searchTitle").String()
	if searchTitle == "" {
		userList = c.ServiceMallCommodity.GetAll(page, size)
		number=int(c.ServiceMallCommodity.CountAll())
	} else {
		number,userList = c.ServiceMallCommodity.GetSearchPage(page, size, searchTitle)
	}

	if userList != nil {
		rs["code"], rs["msg"], rs["data"] = 200, "Succeed", userList
	} else {
		rs["code"], rs["msg"], rs["data"] = conf.FailedCode, "Failed", ""
	}
	rs["total"] = number
	rs["random_commodity"] = c.ServiceMallCommodity.GetRandomHot()
	c.Ctx.JSONP(rs)
}
func (c *MallCommodityController) OptionsPage() {}
