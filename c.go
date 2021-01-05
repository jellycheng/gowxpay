package gowxpay

const (
	Fail    = "FAIL"
	Success = "SUCCESS"
	//支付接口域名,不用写以/结尾哦
	PayDomainUrl = "https://api.mch.weixin.qq.com"

	//货币类型-人民币
	FeeTypeCNY = "CNY"

	//交易类型，不同trade_type决定了调起支付的方式，请根据支付产品正确上传
	TradeTypeJSAPI    = "JSAPI"    //JSAPI支付、小程序支付、公众号内支付
	TradeTypeNATIVE   = "NATIVE"   //Native支付,微信扫一扫支付
	TradeTypeAPP      = "APP"      //app支付
	TradeTypeMWEB     = "MWEB"     //H5支付,微信之外的浏览器中支付
	TradeTypeMICROPAY = "MICROPAY" //付款码支付，付款码支付有单独的支付接口，所以接口不需要上传，该字段在对账单中会出现

)
