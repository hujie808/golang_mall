package utils

import (
	"github.com/tidwall/gjson"
	"log"

	"strings"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
	"web_iris/golang_mall/web/form"
)

func SaveSku(list []gjson.Result, commodity_id int, s services.MallSkuService) (int, string) {

	for _, sku := range list {
		skuDict := sku.Map()
		title := skuDict["name"].String()
		pice := skuDict["price"].Float()
		if title == "" {
			title = skuDict["Title"].String()
		}
		if pice ==0 {
			pice = skuDict["Pice"].Float()
		}

		stock := skuDict["Stock"].Int()
		image := strings.Replace(skuDict["src"].String(), "SUCCESS", "", 1)
		id := int(skuDict["Id"].Int())
		if id >= 1 {
			m_sku := &models.MallSku{
				Id:              id,
				MallCommodityId: commodity_id,
				Title:           title,
				Pice:            float32(pice),
				Images:          strings.Replace(image, conf.Host, "", 1),
				Stock:           int(stock),
				SysStatus:       0,
			}
			code, msg, _ := form.MallSkuForm(m_sku)
			if msg != "" {
				return code, msg
			}
			err := s.Update(m_sku, []string{})
			if err != nil {
				return conf.MsgCode, "sku更新失败"
			}
		} else {
			m_sku := &models.MallSku{
				MallCommodityId: commodity_id,
				Title:           title,
				Pice:            float32(pice),
				Images:          strings.Replace(image, conf.Host, "", 1),
				Stock:           int(stock),
				SysStatus:       0,
			}
			code, msg, _ := form.MallSkuForm(m_sku)
			if msg != "" {
				return code, msg
			}
			err := s.Create(m_sku)
			if err != nil {
				log.Println("save_sku.go SaveSku Create err=", err)
				return conf.MsgCode, "sku创建失败"
			}
		}

	}
	return 200, ""
}
