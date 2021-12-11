package gowxpay

// CloseOrderReqDto 关单入参
type CloseOrderReqDto struct {
	OutTradeNo *string `json:"out_trade_no"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// ClosePostBodyReqDto 关单入参，post内容
type ClosePostBodyReqDto struct {
	// 直连商户号
	Mchid *string `json:"mchid"`
}
