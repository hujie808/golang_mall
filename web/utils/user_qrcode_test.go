package utils

import (
	"testing"
	"web_iris/golang_mall/conf"
)

func TestQr(T *testing.T) {
	QrUser(conf.Host, "woaini")
}
