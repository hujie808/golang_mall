package form

import (
	"fmt"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

func MallCategoryForm(m *models.MallCategory, s services.MallCategoryService, c services.MallCommodityService) (int, string, *models.MallCategory) {
	//count := int(s.CountAll())
	//all := s.GetAll(1, count, )
	//alls := all
	//all_category:=utils.Allcategory{all}
	if m.Desc < 0 {
		return conf.MsgCode, "排序不正确", nil
	}
	if m.Name == "" {
		return conf.MsgCode, "标题不得为空", nil
	}
	if m.Id== m.ParentId{
		return conf.MsgCode, "上级分类不能设置自己", nil
	}
	if m.ParentId > 0 {
		commoditys := c.GetCategory(m.ParentId)
		if len(commoditys) > 0 {
			c_msg := ""
			for _, commodity := range commoditys {
				c_msg += commodity.Title + ";"
			}
			return conf.MsgCode, fmt.Sprintf("您正在创建新的分类,继续操作后请把商品|%s|移动到此分类", c_msg), nil
		}

	}

	return 200, "", m
}

func MallCategoryDeleteForm(m *models.MallCategory, s services.MallCategoryService, c services.MallCommodityService) (int, string) {
	all_parent := s.GetParent(m.Id)
	if len(all_parent) != 0 {
		return conf.MsgCode, "此分类下还有子分类 请先删除子分类"
	}
	all_commodity := c.GetCategory(m.ParentId)
	if len(all_parent) > 0 {
		c_msg := ""
		for _, commodity := range all_commodity {
			c_msg += commodity.Title + ";"
		}
		return conf.MsgCode, fmt.Sprintf("请先删除相关商品,|%s|", c_msg)
	}
	return 200, ""
}
