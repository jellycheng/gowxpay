package gowxpay

import (
	"encoding/json"
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
		SplitLine("-", 18)

	privateKey, err := LoadPrivateKeyWithPath("./apiclient_key.pem")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	fmt.Println(sign)
		SplitLine("-", 18)

	payCfg := SimpleIni2Map("cjs.ini")
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	authorizationHeader := PinAuthorizationHeaderVal(mchid, nonce, timestamp, serialNo, sign)
	fmt.Println(authorizationHeader)

}

//go test -run="TestJsapiPrepayV3"
func TestJsapiPrepayV3(t *testing.T) {
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	openid := payCfg["openid"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile}
	prepayDto := PrepayReqV3Dto{Appid: String(appid),
								Mchid: String(mchid),
								Description: String("购买cjs商品"),
								OutTradeNo: String("2021"+gosupport.GetRandString(10)), // 订单号
								NotifyUrl: String("https://www.weixin.qq.com/wxpay/pay.php"),
								Amount: &AmountReqV3Dto{ // 订单金额
									Currency: String(FeeTypeCNY),
									Total:    Int64(100),
								},
								Payer: &PayerReqV3Dto{
									Openid: String(openid),
								},
				}
	if res, allHeaders, err := JsapiPrepayV3(prepayDto, accountV3Obj);err == nil {
		fmt.Println(gosupport.ToJson(allHeaders))
		fmt.Println(res)

		var prepayRespDtoObj = PrepayRespV3Dto{}
		json.Unmarshal([]byte(res), &prepayRespDtoObj)
		fmt.Println(*prepayRespDtoObj.PrepayId)
		// 验证签名
		//if er:= CheckSignV3(allHeaders, []byte(res));er==nil{
		//	fmt.Println("签名通过")
		//} else {
		//	fmt.Println(er.Error())
		//}
	} else {
		fmt.Println(err.Error())
	}

}

//go test -run="TestGetCertificatesV3"
func TestGetCertificatesV3(t *testing.T) {
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile,ApiV3Key: apiv3key}

	if res, allHeaders, err := GetCertificatesV3(accountV3Obj);err == nil {
		fmt.Println(gosupport.ToJson(allHeaders))
			SplitLine("-", 18)
		fmt.Println(res)
			SplitLine("-", 18)

		var certificatesRespDtoObj = new(CertificatesRespDto)
		json.Unmarshal([]byte(res), certificatesRespDtoObj)
		fmt.Printf("%#v \r\n", certificatesRespDtoObj)
		SplitLine("-", 18)
		// 解析
		associatedData:=*certificatesRespDtoObj.Data[0].EncryptCertificate.AssociatedData
		nonce := *certificatesRespDtoObj.Data[0].EncryptCertificate.Nonce
		ciphertext := *certificatesRespDtoObj.Data[0].EncryptCertificate.Ciphertext
		if certificateData,e := DecryptAES256GCM(apiv3key, associatedData, nonce, ciphertext);e== nil{
			fmt.Println(certificateData)
			if certificateObj, err := LoadCertificate(certificateData);err == nil{
				// 验证签名
				if er:= CheckSignV3(allHeaders, []byte(res), certificateObj);er==nil{
					fmt.Println("签名通过")
				} else {
					fmt.Println(er.Error())
				}
			}
		} else {
			fmt.Println(e.Error())
		}

	} else {
		fmt.Println(err.Error())
	}
}

//go test -run="TestNotifiesReturnV3"
func TestNotifiesReturnV3(t *testing.T)  {
	notify := NotifiesReturnV3{}
	fmt.Println(notify.OK())

	fmt.Println(notify.Fail("处理失败"))
}

// go test -run="TestGetCertificateSerialNumber"
func TestGetCertificateSerialNumber(t *testing.T) {
	// 从证书中获取序列号
	if payCertificate, err := LoadCertificateWithPath("./apiclient_cert.pem");err == nil {
		s := GetCertificateSerialNumber(*payCertificate)
		fmt.Println(s)
	}

}