package gowxpay

//封装v2版本接口

// UnifiedOrderV2 统一下单
func UnifiedOrderV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/unifiedorder"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// OrderQueryV2 订单查询
func OrderQueryV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/orderquery"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// RefundV2 申请退款
func RefundV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/secapi/pay/refund"
	xmlStr, err := c.Post4Cert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// RefundQueryV2 退款查询
func RefundQueryV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/refundquery"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// MicropayV2 付款码支付
func MicropayV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/micropay"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}
