package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/glog"
	"reflect"
	"sort"
	"strings"
	"web_iris/golang_mall/conf"
)

const (
	AckSuccess = `<xml><return_code><![CDATA[SUCCESS]]></return_code></xml>`
	AckFail    = `<xml><return_code><![CDATA[FAIL]]></return_code></xml>`
)

//微信 商户Key
var WXPApiKey string = conf.ApiKey

type WXPayNotify struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int64  `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int64  `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	CouponFee     int64  `xml:"coupon_fee"`
	CouponCount   int64  `xml:"coupon_count"`
	CouponID0     string `xml:"coupon_id_0"`
	CouponFee0    int64  `xml:"coupon_fee_0"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

func ProcessWX(wxn WXPayNotify) bool {

	if !WXPayVerify(wxn) {
		glog.Warning("SIGN FAILED")
		return false
	}

	if !(wxn.ReturnCode == "SUCCESS" && wxn.ResultCode == "SUCCESS") {
		glog.Warning("INVALID STATUS", wxn)
		return false
	}

	/*
		业务逻辑 start

		.
		.
		.

		业务逻辑 end
	*/
	return true
}

func WXPayVerify(data WXPayNotify) bool {
	glog.Info(data)
	sign := WXmd5Sign(data)
	if data.Sign == sign {
		return true
	} else {
		glog.V(8).Info(data.Sign, "  !=  ", sign)
		glog.Warning("WEIXIN PAY VERIFY FAIL")
	}
	return false
}

func WXmd5Sign(data interface{}) (sign string) {
	val := make(map[string]string)
	datavalue := reflect.ValueOf(data)
	if datavalue.Kind() != reflect.Struct {
		glog.Warning("NOT A STRUCT ", data)
		return ""
	}
	var keys []string
	for i := 0; i < datavalue.NumField(); i++ {
		k := datavalue.Type().Field(i)
		kl := k.Tag.Get("xml")
		v := fmt.Sprintf("%v", datavalue.Field(i))

		if v != "" && v != "0" && kl != "sign" {
			val[kl] = v
			keys = append(keys, kl)
		}
	}
	sort.Strings(keys)
	var stra string
	for _, v := range keys {
		stra = stra + v + "=" + val[v] + "&"
	}
	strb := stra + "key=" + WXPApiKey
	glog.V(8).Info("SIGN STRING ", strb)
	hstr := md5.Sum([]byte(strb))

	sum := fmt.Sprintf("%x", hstr)
	sign = strings.ToUpper(sum)
	return sign
}
