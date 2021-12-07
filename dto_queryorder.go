package gowxpay

// 查询订单条件
type QueryOrderReqDto struct {
	Mchid         *string // 直连商户号
	TransactionId *string // 微信支付订单号
	OutTradeNo    *string // 商户订单号
}

// 查询订单返回结果
type QueryOrderRespDto struct {

}


