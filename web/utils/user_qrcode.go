package utils

import (
	"image/png"
	"os"

	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"log"
	"strings"
	"web_iris/golang_mall/conf"
)

func QrUser(qr_url, openid string) string {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	result_dir := strings.Split(dir, "golang_mall")[0] + "golang_mall/web/public/qr_user/" + fmt.Sprintf("%s.png", openid)
	if IsFileExist(result_dir) { //如果存在直接返回url
		resutl_url := conf.Host + strings.Split(result_dir, "golang_mall/web")[1]

		return resutl_url
	}
	var share_url string
	if strings.Contains("?", qr_url) {
		share_url = qr_url + fmt.Sprintf("&share_openid=%s", openid)
	} else {
		share_url = qr_url + fmt.Sprintf("?share_openid=%s", openid)
	}

	qrCode, _ := qr.Encode(share_url, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create(result_dir)
	defer file.Close()
	png.Encode(file, qrCode)
	resutl_url := conf.Host + strings.Split(result_dir, "golang_mall/web")[1]
	return resutl_url
}

func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	}
	fmt.Println("exists", info.Name(), info.Size(), info.ModTime())
	return true
}
