package form

import (
	"fmt"
	"strconv"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

func CommodityForm(commodity *models.MallCommodity, categoryService services.MallCategoryService) (int, string, *models.MallCommodity) {
	if len(commodity.Title) < 3 {
		return conf.MsgCode, "标题类型不正确", nil
	}
	if commodity.Desc < 1 {
		commodity.Desc=100
		//return conf.MsgCode, "排序不正确", nil
	}
	if commodity.CategoryId > 0 {
		c_id := categoryService.Get(commodity.CategoryId)
		if c_id.Id != commodity.CategoryId {
			return conf.MsgCode, "分类Id找不到", nil
		} else {
			parent := categoryService.GetParent(commodity.CategoryId)
			if len(parent) >= 1 {
				return conf.MsgCode, "分类错误:请把商品上传至最下级分类", nil
			}
		}
	} else {
		return conf.MsgCode, "分类Id不正确", nil
	}
	commodity.Price = floatRound(commodity.Price, 2)
	if commodity.Price < 0.01 {
		return conf.MsgCode, "原价格不正确", nil
	}
	commodity.PriceNow = floatRound(commodity.PriceNow, 2)
	if commodity.PriceNow < 0.01 {
		return conf.MsgCode, "现价格不得小于0.01", nil
	}
	if len(commodity.Image) < 1 {
		return conf.MsgCode, "封面图片没有上传", nil
	}
	if commodity.ShareZhitui >= 100 {
		return conf.MsgCode, "直推分佣不得大于100%", nil
	}
	if commodity.ShareXiaji >= 100 {
		return conf.MsgCode, "下级分佣不得大于100%", nil
	}
	if commodity.ShareXiaji+commodity.ShareZhitui >= 100 {
		return conf.MsgCode, "下级价值退比例不得大于100%", nil
	}

	if commodity.ComType == "xn" && commodity.VideoUrl == "" {
		fmt.Println(commodity.ComType)
		return conf.MsgCode, fmt.Sprintf("虚拟物品需上传 付费后观看的视频%s", commodity.ComType), nil
	}

	if commodity.ComType == "" {
		commodity.ComType = "sw"
	}

	return 200, "", commodity
}

func floatRound(f float32, n int) float32 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 32)
	return float32(res)
}
