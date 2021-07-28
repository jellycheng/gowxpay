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
	notifyUrl := payCfg["notify_url"]
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

