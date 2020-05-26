package form

import (
	"log"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

func MallAddressForm(address *models.MallAddress, s services.MallAddressService) (int, string, *models.MallAddress) {
	if len(address.Tel) <= 10 {
		return conf.MsgCode, "手机号不得小于11位", nil
	}
	if len(address.DetailAddress) < 1 {
		return conf.MsgCode, "详细地址必须填写", nil
	}
	userDefault := s.GetUserDefault(address.MallWechatId)
	if userDefault.Id > 0 {
		if userDefault.IdDefautl == 1 && address.IdDefautl == 1 {
			userDefault.IdDefautl = 2
			err := s.Update(userDefault, []string{"id_defautl"})
			if err != nil {
				log.Println("mall_address_form.go MallAddressForm err=", err)
			}
		}
	} else {
		address.IdDefautl = 1
	}
	if len(address.LocationAddress) < 1 {
		return conf.MsgCode, "位置必须填写", nil
	}
	if len(address.RealName) < 1 {
		return conf.MsgCode, "必须填写姓名", nil
	}

	return 200, "", address
}
