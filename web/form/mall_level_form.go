package form

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func MallLevelForm(level *models.MallLevel) (int, string, *models.MallLevel) {
	if level.Levle == "" {
		return conf.MsgCode, "标题类型不正确", nil
	}
	if level.ShareZhitui > 100 || level.ShareXiaji > 100 {
		return conf.MsgCode, "直推和下级不能大于百分百", nil
	}
	if level.ShareZhitui+level.ShareXiaji > 100 {
		return conf.MsgCode, "直推加下级不得大于100", nil
	}
	if level.CommodityPrice < 1 {
		return conf.MsgCode, "打折力度不能小于1折", nil
	}
	return 200, "", level
}
