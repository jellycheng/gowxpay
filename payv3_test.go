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

// 下单： go test -run="TestJsapiPrepayV3"
func TestJsapiPrepayV3(t *testing.T) {
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiClientKeyCertFile := payCfg["apiclient_cert_pem_file"]
	openid := payCfg["openid"]
	payNotifyUrl := payCfg["wxpay_notify_url"]
	accountV3Obj := AccountV3{AppID:appid,
							 MchID: mchid,
							 SerialNo: serialNo,
							 ApiClientKeyPemFile: apiclientKeyPemFile,
							 ApiClientKeyCertFile: apiClientKeyCertFile}
	prepayDto := PrepayReqV3Dto{Appid: StringPtr(appid),
								Mchid: StringPtr(mchid),
								Description: StringPtr("购买商品"),
								OutTradeNo: StringPtr("2021"+gosupport.GetRandString(10)), // 我的订单号，唯一
								NotifyUrl: StringPtr(payNotifyUrl), // 通知地址
								Amount: &AmountReqV3Dto{ // 订单金额
									Currency: StringPtr(FeeTypeCNY),
									Total:    Int64Ptr(1), //订单金额，单位分
								},
								Payer: &PayerReqV3Dto{
									Openid: StringPtr(openid),
								},
				}
	if res, allHeaders, err := JsapiPrepayV3(prepayDto, accountV3Obj);err == nil {
		fmt.Println("我的单号：", *prepayDto.OutTradeNo)
		fmt.Println(gosupport.ToJson(allHeaders))
		fmt.Println(res)

		var prepayRespDtoObj = PrepayRespV3Dto{}
		_ = json.Unmarshal([]byte(res), &prepayRespDtoObj)
		fmt.Println("prepay_id", *prepayRespDtoObj.PrepayId)
		// 验证签名
		if payCertificate, err := LoadCertificateWithPath(accountV3Obj.ApiClientKeyCertFile);err == nil {
			if er:= CheckSignV3(allHeaders, []byte(res), payCertificate);er==nil{
				fmt.Println("签名通过")
			} else {
				fmt.Println("签名失败：",er.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}

}

// go test -run="TestGetCertificatesV3"
func TestGetCertificatesV3(t *testing.T) {
	// 获取 -----BEGIN CERTIFICATE----- 内容
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
		_ = json.Unmarshal([]byte(res), certificatesRespDtoObj)
		fmt.Printf("%#v \r\n", certificatesRespDtoObj)
		SplitLine("-", 18)
		// 解析
		associatedData:=*certificatesRespDtoObj.Data[0].EncryptCertificate.AssociatedData
		nonce := *certificatesRespDtoObj.Data[0].EncryptCertificate.Nonce
		ciphertext := *certificatesRespDtoObj.Data[0].EncryptCertificate.Ciphertext
		if certificateData,e := DecryptAES256GCM(apiv3key, associatedData, nonce, ciphertext);e== nil{
			fmt.Println(certificateData) //打印apiclient_cert.pem文件内容 返回类似-----BEGIN CERTIFICATE-----这样的内容
			if certificateObj, err := LoadCertificate(certificateData);err == nil{
				// 验证签名
				if er:= CheckSignV3(allHeaders, []byte(res), certificateObj);er==nil{
					fmt.Println("签名通过")
				} else {
					fmt.Println("签名失败：",er.Error())
				}
			}
		} else {
			fmt.Println(e.Error())
		}

	} else {
		fmt.Println(err.Error())
	}
}

// 响应通知结果： go test -run="TestNotifiesReturnV3"
func TestNotifiesReturnV3(t *testing.T)  {
	notify := NotifiesReturnV3{}
	fmt.Println("成功示例：", notify.OK())
	fmt.Println("失败示例：", notify.Fail("处理失败"))
}

// go test -run="TestGetCertificateSerialNumber"
func TestGetCertificateSerialNumber(t *testing.T) {
	payCfg := SimpleIni2Map("cjs.ini")
	apiclientCertPemFile := payCfg["apiclient_cert_pem_file"]
	// 从apiclient_cert.pem证书中获取序列号
	if payCertificate, err := LoadCertificateWithPath(apiclientCertPemFile);err == nil {
		s := GetCertificateSerialNumber(*payCertificate)
		fmt.Println(s)
	}

}

// go test -run="TestCloseOrder"
func TestCloseOrder(t *testing.T) {
	// 关单
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile,ApiV3Key: apiv3key}
	reqDto := CloseOrderReqDto{
		OutTradeNo: StringPtr("OutTradeNo_example123"),
		Mchid:      StringPtr(mchid),
	}
	if isOk, allHeaders, err := CloseOrder(reqDto, accountV3Obj);err == nil{
		fmt.Println(allHeaders)
		fmt.Println("关单结果：", isOk)
	} else {
		fmt.Println(err.Error())
	}

}


// go test -run="TestQueryOrder4OutTradeNo"
func TestQueryOrder4OutTradeNo(t *testing.T) {
	// 通过商户订单号查询订单
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile,ApiV3Key: apiv3key}
	reqDto := QueryOrderReqDto{
		OutTradeNo: StringPtr(payCfg["out_trade_no_demo"]),
		Mchid:      StringPtr(mchid),
	}
	if payOrderInfo, allHeaders, err := QueryOrder4OutTradeNo(reqDto, accountV3Obj);err == nil{
		fmt.Println(allHeaders)
		fmt.Println("支付结果：", payOrderInfo)
		payOrderInfoObj := new(QueryOrderRespDto)
		_ = JsonUnmarshal(payOrderInfo, payOrderInfoObj)
		if payOrderInfoObj.Appid != nil {
			if *payOrderInfoObj.TradeState == string(TradeStateSuccess) {
				// 支付成功
				fmt.Println(fmt.Sprintf("支付成功:%+v", payOrderInfoObj))
			} else if *payOrderInfoObj.TradeState == string(TradeStateRefund) {
				// 发生退款
				fmt.Println(fmt.Sprintf("发生退款: %+v", payOrderInfoObj))
			} else {
				fmt.Println(fmt.Sprintf("其它状态: %+v", payOrderInfoObj))
			}

		} else {
			// 查询失败： {"code":"PARAM_ERROR","message":"微信订单号非法"}
			fmt.Println("查询失败：", payOrderInfo)
		}

	} else {
		fmt.Println(err.Error())
	}
}

// go test -run="TestQueryOrder4TransactionId"
func TestQueryOrder4TransactionId(t *testing.T) {
	// 通过微信支付单号查询订单
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile,ApiV3Key: apiv3key}
	reqDto := QueryOrderReqDto{
		TransactionId: StringPtr(payCfg["transaction_id_demo"]),
		Mchid:      StringPtr(mchid),
	}
	if payOrderInfo, allHeaders, err := QueryOrder4TransactionId(reqDto, accountV3Obj);err == nil{
		fmt.Println(allHeaders)
		fmt.Println("支付结果：", payOrderInfo)
		payOrderInfoObj := new(QueryOrderRespDto)
		if err2 := JsonUnmarshal(payOrderInfo, payOrderInfoObj);err2 == nil {
			if payOrderInfoObj.Appid != nil {
				if *payOrderInfoObj.TradeState == string(TradeStateSuccess) {
					// 支付成功
					fmt.Println(fmt.Sprintf("支付成功:%+v", payOrderInfoObj))
				} else if *payOrderInfoObj.TradeState == string(TradeStateRefund) {
					// 发生退款
					fmt.Println(fmt.Sprintf("发生退款: %+v", payOrderInfoObj))
				} else {
					fmt.Println(fmt.Sprintf("其它状态: %+v", payOrderInfoObj))
				}
			} else {
				// 查询失败： {"code":"PARAM_ERROR","message":"微信订单号非法"}
				fmt.Println("查询失败：", payOrderInfo)
			}
		} else {

		}


	} else {
		fmt.Println(err.Error())
	}
}

// go test -run="TestRefundOrder"
func TestRefundOrder(t *testing.T) {
	// 申请退款
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiv3key := payCfg["apiv3key"]
	refundNotifyUrl := payCfg["wxrefund_notify_url"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile,ApiV3Key: apiv3key}
	outrefundNo := "2021refund_" + gosupport.GetRandString(10)
	reqDto := RefundReqV3Dto{
		TransactionId: StringPtr(payCfg["transaction_id_demo"]),
		OutRefundNo:      StringPtr(outrefundNo), //商户退款单号
		NotifyUrl: StringPtr(refundNotifyUrl), // 退款结果回调url
		Amount: &RefundAmountReqV3Dto{
			Refund: Int64Ptr(1), //本次退款金额，单位分
			Total: Int64Ptr(1), //总金额，单位分
			Currency: StringPtr(FeeTypeCNY),
		},
	}
	if refundOrderInfo, allHeaders, err := RefundOrder(reqDto, accountV3Obj);err == nil{
		fmt.Println(allHeaders)
		fmt.Println("退款结果：", refundOrderInfo)
		refundOrderInfoObj := new(RefundRespV3Dto)
		if err2 := JsonUnmarshal(refundOrderInfo, refundOrderInfoObj);err2 == nil {
			if refundOrderInfoObj.RefundId != nil && gosupport.StrInSlice(*refundOrderInfoObj.Status,[]string{string(RefundStatusSuccess), string(RefundStatusProcessing)}) {
				// 退款申请成功
				fmt.Println(fmt.Sprintf("%+v", refundOrderInfoObj))
			} else {
				// 退款失败： {"code":"NOT_ENOUGH","message":"基本账户余额不足，请充值后重新发起"}
				//  {"code":"INVALID_REQUEST","message":"订单已全额退款"}
				//  {"code":"INVALID_REQUEST","message":"订单金额或退款金额与之前请求不一致，请核实后再试"}
				fmt.Println("退款失败：", refundOrderInfo)
			}
		} else {

		}


	} else {
		fmt.Println(err.Error())
	}

}

// go test -run="TestRefundQuery"
func TestRefundQuery(t *testing.T) {
	// 查询退款
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid, MchID: mchid, SerialNo: serialNo,ApiClientKeyPemFile: apiclientKeyPemFile,ApiV3Key: apiv3key}
	reqDto := QueryByOutRefundNoReqV3Dto{
		OutRefundNo: StringPtr(payCfg["out_refund_no_demo"]), //商户退款单号
	}
	if refundOrderInfo, allHeaders, err := RefundQuery(reqDto, accountV3Obj);err == nil{
		fmt.Println(allHeaders)
		fmt.Println("查询退款结果：", refundOrderInfo)
		refundOrderInfoObj := new(RefundRespV3Dto)
		if err2 := JsonUnmarshal(refundOrderInfo, refundOrderInfoObj);err2 == nil {
			if refundOrderInfoObj.RefundId != nil && gosupport.StrInSlice(*refundOrderInfoObj.Status,[]string{string(RefundStatusSuccess), string(RefundStatusProcessing)}) {
				// 退款申请成功
				fmt.Println(fmt.Sprintf("%+v", refundOrderInfoObj))
			} else {
				// 退款查询失败：{"code":"RESOURCE_NOT_EXISTS","message":"退款单不存在"}
				fmt.Println("退款查询失败：", refundOrderInfo)
			}
		} else {

		}

	} else {
		fmt.Println(err.Error())
	}

}


// go test -run="TestRefundNotifyParse"
func TestRefundNotifyParse(t *testing.T) {
	// 退款通知解析
	allHeaders := map[string]string{
		"Content-Type":        "application/json",
		"Wechatpay-Nonce":     "IcWDN8T6pbXxeaFrFLnmHth821K4l3bd",
		"Wechatpay-Timestamp": "1661832052",
		"Wechatpay-Serial":    "2B08154B9E99549F56A2B71EB5F22BD5907F7C52",
		"Wechatpay-Signature": "签名结果",
	}

	postBody := `
{
	"id": "a9a0bc6d-20e3-52ab-88b5-c3ca88db149f",
	"create_time": "2022-08-30T12:00:46+08:00",
	"resource_type": "encrypt-resource",
	"event_type": "REFUND.SUCCESS",
	"summary": "退款成功",
	"resource": {
		"original_type": "refund",
		"algorithm": "AEAD_AES_256_GCM",
		"ciphertext": "Z4+/KUuETT9Q0PnIXyZVDv07uQfiNLlH7d7wkG+J16ScaiA8MxCxKXvswlVye+PDg5xWF0OdYPNSBS5nIBfAlOSiugGJVjZyck0Eka/NDhbqdEcd2rQGKCqoFqRapF6kCwdIjhCmFD2JBtZCahTnzVHZzYTiAA7k9kc1LVPgZRh4gbOysjJsPGRygZEn766AEfowE71/kwJkEMqYRKLPW8hy39/1d2p+wLpOXMePeVxofAUkbTnc7yOQTgOjR3iR20Mja6yPDuP11biSOa5Y8Noe3VM6lMFzmGxSPiZbyFO3fyUaxc+Q1dSSF5fcF9vi2qiejzE0Upir33l4BnwxxEuLtH7uM1sh0WebrzH3H+C3yN+R+FivWkxM7loZKMRFjlrQZtoy+QajghM6tjyleBXgzahDFNuAZlqVf74+m2H0I7bMu8R4OtwFcBv1T6s/yIIR0XjnB/4kOJeZ+SZRhv5DwJyVnHP6ZDmkAj2dU56kItcTX5yG",
		"associated_data": "refund",
		"nonce": "PeIVqUnoKOlc"
	}
}
`
	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiclientCertPemFile := "" //payCfg["apiclient_cert_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid,
							MchID: mchid,
							SerialNo: serialNo,
							ApiClientKeyPemFile: apiclientKeyPemFile,
							ApiClientKeyCertFile: apiclientCertPemFile,
							ApiV3Key: apiv3key}
	CloseCheckTime = true // 关闭校验时间
	isSkipSign := true //跳过签名验证
	if notifyDto, err := RefundNotifyParse(postBody, allHeaders, accountV3Obj, isSkipSign);err == nil {
		fmt.Println(fmt.Sprintf("%+v", notifyDto))
		// {"mchid":"1628039785","out_trade_no":"soUFMywj","transaction_id":"4200067664202208302723736693","out_refund_no":"2021refund_ExUZGGrrDd","refund_id":"50302302932022083024350423805","refund_status":"SUCCESS","success_time":"2022-08-30T12:00:46+08:00","amount":{"total":1,"refund":1,"payer_total":1,"payer_refund":1},"user_received_account":"支付用户零钱"}
		fmt.Println(notifyDto.Resource.Plaintext)
		obj := new(RefundNotifyResourceDto)
		_ = JsonUnmarshal(notifyDto.Resource.Plaintext, obj)
		fmt.Println(fmt.Sprintf("%+v", obj))

	} else {
		fmt.Println(err.Error())
	}

}

// go test -run="TestRefundNotifyResourceDto"
func TestRefundNotifyResourceDto(t *testing.T) {
	str := `{
    "mchid":"1604928906",
    "out_trade_no":"20216oBoepskSO",
    "transaction_id":"4200001322202112137430682123",
    "out_refund_no":"2021refund_f6QyyWlAQM",
    "refund_id":"503011002020211213153447",
    "refund_status":"SUCCESS",
    "success_time":"2021-12-13T18:08:31+08:00",
    "amount":{
        "total":10,
        "refund":3,
        "payer_total":10,
        "payer_refund":3
    },
    "user_received_account":"支付用户零钱"
}`
	obj := new(RefundNotifyResourceDto)
	_ = JsonUnmarshal(str, obj)
	fmt.Println(fmt.Sprintf("%+v", obj))
	fmt.Println(fmt.Sprintf("amount: %+v", *obj.Amount))
	fmt.Println((*obj.SuccessTime).Format("2006-01-02 15:04:05"))

}

// go test -run="TestPayNotifyParse"
func TestPayNotifyParse(t *testing.T) {
	// 正常支付通知内容解析
	allHeaders := map[string]string{
		"Content-Type":        "application/json",
		"Wechatpay-Nonce":     "j9VYUQwmBfTi8rQovzVa62gN99jV8rYS",
		"Wechatpay-Timestamp": "1639301204",
		"Wechatpay-Serial":    "5CDB363A77BE5818B8F12462C36ED5A2892AEC36",
		"Wechatpay-Signature": "签名结果",
	}

	postBody := `{
    "id":"87ff2da3-a165-5c79-b225-cfc1ec2ea7b6",
    "create_time":"2021-12-12T17:26:44+08:00",
    "resource_type":"encrypt-resource",
    "event_type":"TRANSACTION.SUCCESS",
    "summary":"支付成功",
    "resource":{
        "original_type":"transaction",
        "algorithm":"AEAD_AES_256_GCM",
        "ciphertext":"rM2bn6vS9def2ydrAv21DbMGj8XNC+LwrmBQfCGlHL+KBpJpRm94pKHYDl3Ega638QxbGsesFFH2isZPk0HdLii1yLF9v8trIEMJyQ6AacYXKvXvIqTNBUSFx9zKPaQ9rGxLSVkrCx0Ii4MCRoQt7JsrgU3BT1v6AW6Y4eEF0WzcoKXyJDuKhI6Zwwxl4KJwpuUNOtdPW6wTHGLjpI/CX/Hi7RfEI4tmrfoE6hJxmuL+krfH2mB8Enlu4QbW4ukfnvPXazSR+A9lf+EFtQe1CPPjKHlGVOaojUBKMutX8Q3i7QikR5iajWLJlDiZ2lT/JvqhAD8be6lRu4v92ryT4s+sLYf1l8CuzSO/56Jef/l06+PZ5PmMSsL5xYQvimf4FA9TVxvbVGa7Jvuu+mfGXlKcSqWYJcY778TwabCKn+fU+EDNNaLPYlwTh7q6jwTp+aXH/GJ+efPlrU25H5hUFoctxVXVm/RV1pfj4M5h+zTVVI+SZeJUYoBqVa5D7HU/4o/w2TUbJ6Cd094pr+AXxbkW4zkTIWrP5/DUH7HzgocMqE4yubBJ9HI3aIdv",
        "associated_data":"transaction",
        "nonce":"aJUWDm2xmnaD"
    }
}`

	payCfg := SimpleIni2Map("cjs.ini")
	appid := payCfg["appid"]
	mchid := payCfg["mchid"] // 支付商户号
	serialNo := payCfg["serialno"] // 证书序列号
	apiclientKeyPemFile := payCfg["apiclient_key_pem_file"]
	apiclientCertPemFile := "" //payCfg["apiclient_cert_pem_file"]
	apiv3key := payCfg["apiv3key"]
	accountV3Obj := AccountV3{AppID:appid,
							MchID: mchid,
							SerialNo: serialNo,
							ApiClientKeyPemFile: apiclientKeyPemFile,
							ApiClientKeyCertFile: apiclientCertPemFile,
							ApiV3Key: apiv3key}

	CloseCheckTime = true // 关闭校验时间
	isSkipSign := true //跳过签名验证
	if notifyDto, err := PayNotifyParse(postBody, allHeaders, accountV3Obj,isSkipSign);err == nil {
		fmt.Println(fmt.Sprintf("%#v", notifyDto))
	} else {
		fmt.Println(err.Error())
	}

}
