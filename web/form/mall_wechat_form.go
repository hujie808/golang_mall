package form

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

func MallWechatForm(wechat *models.MallWechat, s services.MallWechatService) (int, string, *models.MallWechat) {
	if wechat.Id < 1 {
		return conf.MsgCode, "修改Id不得小于1", nil
	}
	//if wechat.Integral < 1 {
	//	return conf.MsgCode, "虚拟积分不得小于1", nil
	//}
	if wechat.LevelId >= 1 {
		s := services.NewMallLevelService()
		if s.Get(wechat.LevelId).Id < 1 {
			return conf.MsgCode, "无此级别", nil
		}
	}
	//if len(maydistributor) == 11 {
	//	s_tel := s.SearchNickname(1, 2, "", maydistributor)
	//	if len(s_tel) == 1 {
	//		wechat.MyDistributorId = s_tel[0].Id
	//
	//	} else {
	//		return conf.MsgCode, "无此经销商", nil
	//	}
	//} else {
	//	return conf.MsgCode, "手机长度不够", nil
	//}

	if wechat.ShareId > 0 {
		user_Shaer := s.Get(wechat.ShareId)
		if user_Shaer.Id < 1 {
			return conf.MsgCode, "无此分享人", nil
		}
		if user_Shaer.UserType != 2 {
			return conf.MsgCode, "此分享人不是研发师", nil
		}
	}

	//if len(shareid) == 11 {
	//	s_tel := s.SearchNickname(1, 2, "", maydistributor)
	//	if len(s_tel) == 1 {
	//		wechat.ShareId = s_tel[0].Id
	//	} else {
	//		return conf.MsgCode, "无此分享人", nil
	//	}
	//} else {
	//	return conf.MsgCode, "手机长度不够", nil
	//}
	return 200, "", wechat
}
