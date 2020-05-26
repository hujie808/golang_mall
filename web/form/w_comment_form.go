package form

import (
	"web_iris/golang_mall/models"
)

func WCommoentForm(mallOrderInfo *models.MallOrderInfo) (int, string, *models.MallOrderInfo) {
	if len(mallOrderInfo.Comment) < 1 {
		mallOrderInfo.Comment = "用户未评论"
		//return conf.MsgCode,"请勿提交空评论",nil
	}
	return 200, "", mallOrderInfo
}
