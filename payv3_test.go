package gowxpay

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"net/http"
	"testing"
)

// v3版本接口测试

//go test -run="TestPinAuthorizationHeaderVal"
func TestPinAuthorizationHeaderVal(t *testing.T) {
	urlPath := "/v3/certificates"
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)
	body := ``

	reqStr := PinReqMessage(http.MethodGet, urlPath, timestamp, nonce, body)
	fmt.Println(reqStr)
	fmt.Println("--------------------------")

	privateKey, err := LoadPrivateKeyWithPath("./apiclient_key.pem")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	fmt.Println(sign)
	fmt.Println("--------------------------")

	payCfg := SimpleIni2Map("cjs.ini")
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	authorizationHeader := PinAuthorizationHeaderVal(mchid, nonce, timestamp, serialNo, sign)
	fmt.Println(authorizationHeader)

}

