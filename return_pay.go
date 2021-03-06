package gowxpay

import "github.com/jellycheng/gosupport"

//用于微信支付通知，响应微信支付服务的参数，使用V2接口通知响应
type NotifiesReturnV2 struct{}

//通知成功返回
func (n *NotifiesReturnV2) OK() string {
	var params = map[string]string {
		"return_code":Success,
		"return_msg":"ok",
	}
	return gosupport.Map2XMLV2(params)
}

//通知处理失败返回-不成功
func (n *NotifiesReturnV2) Fail(errMsg string) string {
	var params = map[string]string {
		"return_code":Fail,
		"return_msg":errMsg,
	}
	return gosupport.Map2XMLV2(params)
}


//用于微信支付通知，响应微信支付服务的参数，使用V3接口通知响应
type NotifiesReturnV3 struct{}


