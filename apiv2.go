package gowxpay

//封装v2版本接口

// 统一下单
func UnifiedOrderV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/unifiedorder"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

//订单查询
func OrderQueryV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/orderquery"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// 退款
func RefundV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/secapi/pay/refund"
	xmlStr, err := c.Post4Cert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// 退款查询
func RefundQueryV2(c PayClient, params MapParams) (MapParams, error) {
	var urlStr = PayDomainUrl + "/pay/refundquery"
	xmlStr, err := c.Post4NotCert(urlStr, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}


