package form

import (
	"fmt"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func MallSkuForm(sku *models.MallSku) (int, string, *models.MallLevel) {
	if len(sku.Title) < 1 {
		fmt.Println(sku.Title)
		return conf.MsgCode, "Sku名称不得小于1", nil
	}
	if sku.MallCommodityId < 1 {
		return conf.MsgCode, "suk相关商品Id不正确", nil
	}
	if sku.Pice < 0&&sku.Pice==0 {
		return conf.MsgCode, "suk价格不得小于或者等于0", nil
	}
	//TODO照片限制
	//if len(sku.Images)<1{
	//	return conf.MsgCode, "sku照片需上传", nil
	//}
	if sku.Stock < 1 {
		return conf.MsgCode, "suk库存必须大于1", nil
	}
	return 200, "", nil
}
