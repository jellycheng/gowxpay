package gowxpay

//封装v2版本接口

// 统一下单
func UnifiedOrderV2(c PayClient, params MapParams) (MapParams, error) {
	var url = PayDomainUrl + "/pay/unifiedorder"
	xmlStr, err := c.Post4NotCert(url, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

//订单查询
func OrderQueryV2(c PayClient, params MapParams) (MapParams, error) {
	var url = PayDomainUrl + "/pay/orderquery"
	xmlStr, err := c.Post4NotCert(url, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// 退款
func RefundV2(c PayClient, params MapParams) (MapParams, error) {
	var url = PayDomainUrl + "/secapi/pay/refund"
	xmlStr, err := c.Pos4Cert(url, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}

// 退款查询
func RefundQueryV2(c PayClient, params MapParams) (MapParams, error) {
	var url string = PayDomainUrl + "/pay/refundquery"
	xmlStr, err := c.Post4NotCert(url, params)
	if err != nil {
		return nil, err
	}
	return c.ProcessResponseXml(xmlStr)
}


