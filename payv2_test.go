package gowxpay

import (
	"encoding/json"
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/sign"
	"strconv"
	"testing"
)

//go test -run="TestNotifiesReturnV2"
func TestNotifiesReturnV2(t *testing.T) {
	nrObj := NotifiesReturnV2{}
	fmt.Println(nrObj.OK())

	fmt.Println(nrObj.Fail("签名失败"))

}

//go test -run="TestPkcs12ToPem"
func TestPkcs12ToPem(t *testing.T) {
	res,_ := GetCertData("./apiclient_cert.p12")
	mchid := "mchid123"
	//将Pkcs12转成Pem
	if s, e :=Pkcs12ToPem(res, mchid);e == nil{
		fmt.Println(s)
	} else {
		fmt.Println(e.Error())
	}
}

//统一下单： go test -run="TestUnifiedOrder"
func TestUnifiedOrder(t *testing.T) {
	payCfg := SimpleIni2Map("cjs.ini")
	//fmt.Println(payCfg)
	appid := payCfg["appid"]
	mchid := payCfg["mchid"]
	apikey := payCfg["apikey"]
	openid := payCfg["openid"]
	notifyUrl := payCfg["wxpay_notify_url"]
	client := NewPayClient(NewAccount(appid, mchid, apikey))
	params := make(MapParams)
	params.SetString("body", "购买商品").
		SetString("out_trade_no", "orderID" + gosupport.GetRandString(6)).
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", notifyUrl).
		SetString("trade_type", TradeTypeJSAPI).SetString("openid", openid)
	if res, err := UnifiedOrderV2(*client, params);err == nil{//统一下单
		fmt.Println(res)
		var retWxRequestPayment = make(map[string]string)
		if tradeType,ok:=res["trade_type"];ok && tradeType == TradeTypeJSAPI { //小程序支付
			retWxRequestPayment["appId"] = appid
			retWxRequestPayment["package"] = "prepay_id=" + res["prepay_id"]
			retWxRequestPayment["timeStamp"] =  strconv.FormatInt(gosupport.Time(), 10)
			retWxRequestPayment["nonceStr"] = gosupport.GetRandString(16)
			retWxRequestPayment["signType"] = "MD5"
			paySign,srcStr,_ := sign.WxPaySign(retWxRequestPayment, retWxRequestPayment["signType"], apikey)
			retWxRequestPayment["paySign"] = paySign
			fmt.Println(srcStr)
			retByte, _ := json.Marshal(retWxRequestPayment)
			fmt.Println(string(retByte))
		}
	} else {
		fmt.Println(err.Error())
	}


}


// 付款码支付 go test -run=TestMicropayV2
func TestMicropayV2(t *testing.T) {
	payCfg := SimpleIni2Map("cjs.ini")
	//fmt.Println(payCfg)
	appid := payCfg["appid"]
	mchid := payCfg["mchid"]
	apikey := payCfg["apikey"]
	auth_code_demo := payCfg["auth_code_demo"]
	client := NewPayClient(NewAccount(appid, mchid, apikey))
	params := make(MapParams)
	params.SetString("body", "购买商品").
		SetString("out_trade_no", "so" + gosupport.GetRandString(6)). //商户订单号
		SetInt64("total_fee", 1). //订单金额
		SetString("spbill_create_ip", "127.0.0.1"). //终端IP
		SetString("auth_code", auth_code_demo) // 付款码
	if res, err := MicropayV2(*client, params);err == nil{
		fmt.Println(res) //result_code:FAIL return_code:SUCCESS return_msg:OK sign:27B6D2411B8AC994CC15A50333E39E36]
		if res["return_code"] == "SUCCESS" && res["result_code"] == "SUCCESS" { //付款成功
			fmt.Println(fmt.Sprintf("付款成功:微信单号%s", res["transaction_id"]))
		} else {
			fmt.Println("付款失败")
			if res["err_code"] == "USERPAYING" {
				fmt.Println("需要用户输入支付密码,此时再等待5秒，通过out_trade_no查询订单支付结果")
			}
		}
	} else {
		fmt.Println(err.Error())
	}

}
