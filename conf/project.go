package conf

import "time"

const SysTimeform = "2006-01-02 15:04:05" //系统时间

const SysTimeformShort = "2006-01-02" //系统data时间

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing") //系统时区

var SignSecret = []byte("zhejiushimiyao") //定义秘钥

var CookieSecret = "hellolottert"

// 是否需要启动全局计划任务服务
var RunningCrontabService = true

var LogBackInCode = 401 //重新登录code
var MsgCode = 301       //弹出消息code
var FailedCode = 402    //失败code

var Host = "https://go.cttec.top/"

var MinPrice = 10 //最低提现金额

//微信相关配置
//#小程序 appid
const Appid = "wxda6aef231f450ddb"

//#小程序密钥
const Secret = "cdd99d1c65d00a974e951d6dd9fc4f29"

//#商户帐号ID
const Mch_id = "1562835231"

//#微信支付密钥
const ApiKey = "vPlaDk6diOX6zOTD0pWDGy9OHFJBDOyK"

//#微信异步通知，例：https://www.nideshop.com/api/pay/notify

const CREATE_IP = "39.96.160.74" // 发起支付请求的ip

const NotifyUrl = "https://go.cttec.top/pay/check_wxpay"// 微信支付结果回调接口，需要改为你的服务器上处理结果回调的方法路径

const UFDODER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder" // url是微信下单api
