package controllers

import (
	"github.com/kataras/iris"
	"sort"
	"web_iris/golang_mall/services"
)

type WCategory struct {
	Ctx              iris.Context
	ServiceBanner    services.MallBannerService
	ServiceCommodity services.MallCommodityService
	ServiceCategory  services.MallCategoryService
}

func (c *WCategory) Post() {
	yiji := c.ServiceCategory.GetParent(0)
	c_number := int(c.ServiceCategory.CountAll())
	results := make(map[string]interface{}, 0)
	result := make([]map[string]interface{}, 0)
	all_category := c.ServiceCategory.GetAll(1, c_number)
	categoryIdList := []int{}
	for _, category := range yiji {
		categoryIdList = append(categoryIdList, category.Id)

	}
	erji := c.ServiceCategory.GetParenAll(categoryIdList)
	if len(yiji) == c_number { //一级
		results["category"] = "1"
		for _, category := range yiji {
			d_category := make(map[string]interface{})
			d_category["Id"] = category.Id
			d_category["Title"] = category.Name
			d_category["Image"] = category.Image
			d_category["ParentId"] = category.ParentId
			d_category["level"] = 1
			result = append(result, d_category)
			categoryIdList = append(categoryIdList, category.Id)
		}
		if categoryIdList != nil {
			commodityList := c.ServiceCommodity.GetCategoryAll(categoryIdList)
			for _, comm := range commodityList {
				if comm.ComType == "xn" {
					continue
				}
				d_comm := make(map[string]interface{})
				d_comm["Id"] = comm.Id
				d_comm["Title"] = comm.Title
				d_comm["ParentId"] = comm.CategoryId
				d_comm["Image"] = comm.Image
				d_comm["PriceNow"] = comm.PriceNow
				d_comm["type"] = "commodity"
				d_comm["level"] = 1
				result = append(result, d_comm)
			}
		}

	} else if len(erji) > 0 && len(erji)+len(yiji) == c_number { //二级
		results["category"] = "2"
		erji_commodity := []int{} //二级商品列表
		if erji != nil {
			for _, category := range all_category {
				d_category := make(map[string]interface{})
				d_category["Id"] = category.Id
				d_category["Title"] = category.Name
				d_category["Image"] = category.Image
				d_category["ParentId"] = category.ParentId
				if category.ParentId > 0 {
					d_category["level"] = 2
					erji_commodity = append(erji_commodity, category.Id)
				} else {
					cLevel1Commodity := c.ServiceCommodity.GetCategory(category.Id)
					if len(cLevel1Commodity) >= 1 {
						d_category2 := make(map[string]interface{})
						d_category2["Id"] = category.Id
						d_category2["Title"] = category.Name
						d_category2["Image"] = category.Image
						d_category2["ParentId"] = category.Id
						d_category2["level"] = 2
						result = append(result, d_category2)
					}

					d_category["level"] = 1
				}
				result = append(result, d_category)

			}
		}
		if categoryIdList != nil { //新建二级==一级名字, 再加入商品
			commodityList := c.ServiceCommodity.GetCategoryAll(categoryIdList)
			for _, comm := range commodityList {
				if comm.ComType == "xn" {
					continue
				}
				d_comm := make(map[string]interface{})
				d_comm["Id"] = comm.Id
				d_comm["Title"] = comm.Title
				d_comm["ParentId"] = comm.CategoryId
				d_comm["Image"] = comm.Image
				d_comm["PriceNow"] = comm.PriceNow
				d_comm["type"] = "commodity"
				result = append(result, d_comm)
			}
		}

		if erji_commodity != nil {
			commodityList := c.ServiceCommodity.GetCategoryAll(erji_commodity)
			for _, comm := range commodityList {
				if comm.ComType == "xn" {
					continue
				}
				d_comm := make(map[string]interface{})
				d_comm["Id"] = comm.Id
				d_comm["Title"] = comm.Title
				d_comm["ParentId"] = comm.CategoryId
				d_comm["Image"] = comm.Image
				d_comm["PriceNow"] = comm.PriceNow
				d_comm["type"] = "commodity"
				result = append(result, d_comm)
			}
		}

	} else { //多级
		sort.Ints(categoryIdList)

		results["category"] = "3"
		all_category := all_category
		if all_category != nil {
			for _, category := range all_category {
				d_category := make(map[string]interface{})
				d_category_two := make(map[string]interface{}) //重新包二级分类 父id为自己
				d_category["Id"] = category.Id
				d_category["Title"] = category.Name
				d_category["Image"] = category.Image
				d_category["ParentId"] = category.ParentId
				i := sort.Search(len(categoryIdList),
					func(i int) bool { return categoryIdList[i] >= category.ParentId }) //找到二级分类的index
				if category.ParentId == 0 {                                             //一级分类
					d_category["level"] = 1
				} else if i < len(categoryIdList) && categoryIdList[i] == category.ParentId {
					d_category_two["Id"] = category.Id
					d_category_two["Title"] = category.Name
					d_category_two["Image"] = category.Image
					d_category_two["level"] = 0
					d_category_two["ParentId"] = category.Id
					result = append(result, d_category_two)
					d_category["ParentId"] = category.ParentId
					d_category["level"] = 2
				} else {
					d_category["level"] = 0
				}
				result = append(result, d_category)
			}
		}
	}
	results["data"] = result
	c.Ctx.JSON(results)
}
