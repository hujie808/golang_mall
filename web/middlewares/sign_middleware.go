package identity

import (
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"web_iris/golang_mall/comm"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/models"
)

//登录测试

func SignMiddleware(ctx iris.Context) {
	rs := make(map[string]interface{})
	sign := ctx.FormValue("sign")
	var uid int
	if sign != "" {
		i_sign, err := strconv.Atoi(ctx.FormValue("sign_uid"))
		uid = int(i_sign)
		if err != nil {
			log.Println("sign_middleware.go SignMiddleware err=", err)
			return
		}
	} else {
		data := comm.GetBody(ctx.Request().Body)
		ctx.Request().Body = ctx.Request().Body
		log.Println(data)
		ctx.Values().Set("data", data)
		sign = gjson.Get(data, "sign").String()
		uid = int(gjson.Get(data, "sign_uid").Int())
	}

	if sign == "" {
		rs["code"] = conf.LogBackInCode
		rs["msg"] = "log back in,Signature is nil"
		ctx.JSON(rs)
	} else {
		adminUser := &models.ObjLoginuser{
			Uid:      uid,
			Username: "",
			Password: "",
			Ip:       comm.ClientIP(ctx.Request()),
			Sign:     sign,
		}
		is_sign := comm.GetLoginUser(adminUser)
		log.Println(adminUser)
		if is_sign {
			ctx.Values().Set("user_id", uid)
			ctx.Next() // 执行下一个处理器。
		} else {
			rs["code"] = conf.LogBackInCode
			rs["msg"] = "log back in,signature is past due"
			ctx.JSON(rs)
		}
	}
}

//openid 验证
func OpenidMiddleware(ctx iris.Context) {
	if ctx.Method() == "GET" {
		ctx.Next()
		return
	}
	rs := make(map[string]interface{})
	openid := ctx.FormValue("openid")

	if openid != "" && len(openid) > 25 {
		ctx.Next() // 执行下一个处理器。
	} else {
		data := comm.GetBody(ctx.Request().Body)
		ctx.Request().Body = ctx.Request().Body
		log.Println(data)
		ctx.Values().Set("data", data)
		openid = gjson.Get(data, "openid").String()
		if openid != "" && len(openid) > 25 {
			ctx.Next() // 执行下一个处理器。
		} else {
			rs["code"] = conf.LogBackInCode
			rs["msg"] = "not opneid "
			ctx.JSON(rs)
		}
	}
}

//openid 验证
func NoOpenidMiddleware(ctx iris.Context) {
	if ctx.Method() == "GET" {
		ctx.Next()
		return
	}
	data := comm.GetBody(ctx.Request().Body)
	ctx.Request().Body = ctx.Request().Body
	log.Println(data)
	ctx.Values().Set("data", data)
	ctx.Next() // 执行下一个处理器。

}
