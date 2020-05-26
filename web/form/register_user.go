package form

import (
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

func AddAdminUserForm(adminUser *models.AdminUser) (int, string, *models.AdminUser) {

	if len(adminUser.Username) < 3 || len(adminUser.Username) > 16 {
		return conf.MsgCode, "用户长度小于3或大于16位", nil
	}
	if len(adminUser.Password) < 6 || len(adminUser.Password) > 16 {
		return conf.MsgCode, "密码长度小于6或大于16位", nil
	}
	if adminUser.SysStatus > 2 {
		return conf.MsgCode, "状态码有误", nil
	}
	if adminUser.IsSuperuser > 2 {
		return conf.MsgCode, "权限设置有误", nil
	}
	return 200, "", adminUser
}
