package comm

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"web_iris/golang_mall/conf"
	"web_iris/golang_mall/datasource"
	"web_iris/golang_mall/models"
	"web_iris/golang_mall/services"
)

//拿取登录IP
func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

//跳转方法
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

// 从cookie中得到当前登录的用户
func GetLoginUser(loginuser *models.ObjLoginuser) bool {
	cacheObj := datasource.InstanceCache()
	key := fmt.Sprintf("user_id_%d", loginuser.Uid)
	adminUser, err := cacheObj.Do("HGET", key, loginuser.Ip)
	if err != nil {
		log.Println("func_web.go  SetLoginUser EXPIRE err=", err, "key, loginuser.Ip", key, loginuser.Ip)
		return false
	}
	if adminUser != nil {
		sign := createLoginuserSign(loginuser)
		if sign == loginuser.Sign {
			return true
		} else {
			return false
		}

	} else {
		return false
	}

}

// 将登录的用户信息设置到redis中
//redis储存两个 1,id2,sign //get会生成sign  然后对比 判断 sign相同就可以访问

func SetLoginUser(loginuser *models.ObjLoginuser) (int, string) {
	key := fmt.Sprintf("user_id_%d", loginuser.Uid)
	if loginuser == nil || loginuser.Uid < 1 {
		return 0, ""
	}
	adminuser := services.NewAdminUserService().LoginUser(loginuser.Username, loginuser.Password)
	if adminuser != nil {
		sign := createLoginuserSign(loginuser)
		cacheObj := datasource.InstanceCache()
		exist, hseterr := cacheObj.Do("HSET", key, loginuser.Ip, sign)
		if hseterr != nil {
			log.Println("func_web.go  SetLoginUser HSET err=", hseterr, "key, loginuser.Ip, sign=", key, loginuser.Ip, sign)
		}
		if exist == nil {
			_, exerr := cacheObj.Do("EXPIRE", key, 86400*30)
			if exerr != nil {
				log.Println("func_web.go  SetLoginUser EXPIRE err=", hseterr, "key=", key)
			}
		} else {
		}

		return loginuser.Uid, sign
	} else {
		return 0, ""
	}

}

// 根据登录用户信息生成加密字符串
func createLoginuserSign(loginuser *models.ObjLoginuser) string {
	str := fmt.Sprintf("uid=%d&ip=%s&secret=%s", loginuser.Uid, loginuser.Ip, conf.CookieSecret)
	return CreateSign(str)
}

//请求数据
func GetBody(bodys io.ReadCloser) string {

	body, err := ioutil.ReadAll(bodys)
	if err != nil {
		log.Println("func_web.go GetBody is err=", err)
		return ""
	}
	return string(body)
}
