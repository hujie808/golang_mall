package form

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func MallArticleForm(article *models.MallArticle) (int, string, *models.MallArticle) {
	if article.Title == "" {
		return conf.MsgCode, "标题类型不正确", nil
	}
	if len(article.AType) > 3 {
		return conf.MsgCode, "类型不正确", nil
	}
	if len(article.Content) == 0 {
		return conf.MsgCode, "内容不得为空", nil
	}
	if len(article.Image) > 200 {
		return conf.MsgCode, "图片地址不正确", nil
	}
	if article.Order == 0 {
		article.Order=100
	}
	return 200, "", article
}
