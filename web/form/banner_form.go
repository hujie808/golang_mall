package form

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func BannerForm(banner *models.MallBanner) (int, string, *models.MallBanner) {
	if banner.Title > 10 {
		return conf.MsgCode, "标题类型不正确", nil
	}
	if banner.Image == "" {
		return conf.MsgCode, "图片地址不正确", nil
	}
	if banner.Desc < 1 {
		return conf.MsgCode, "顺序不能少于1", nil
	}
	if banner.Title == 8 {
		if banner.BannerName == "" {
			return conf.MsgCode, "需上传首页分类名称", nil
		}
	}

	return 200, "", banner
}
