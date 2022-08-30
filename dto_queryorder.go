package gowxpay

// QueryOrderReqDto 查询订单条件
type QueryOrderReqDto struct {
	Mchid         *string // 直连商户号
	TransactionId *string // 微信支付订单号
	OutTradeNo    *string // 商户订单号
}

// QueryOrderRespDto 查询订单返回结果，正向通知resource解析后存入Plaintext的结果
type QueryOrderRespDto struct {
	Appid           *string            `json:"appid,omitempty"` //应用appid
	Mchid           *string            `json:"mchid,omitempty"` //商户ID
	OutTradeNo      *string            `json:"out_trade_no,omitempty"` //商户订单号
	TransactionId   *string            `json:"transaction_id,omitempty"` //微信支付订单号
	TradeType       *string            `json:"trade_type,omitempty"`   //交易类型，如JSAPI、NATIVE
	TradeState      *string            `json:"trade_state,omitempty"`  //交易状态，SUCCESS成功，CLOSED已关闭
	TradeStateDesc  *string            `json:"trade_state_desc,omitempty"` //交易状态描述
	BankType        *string            `json:"bank_type,omitempty"`    //付款银行
	Attach          *string            `json:"attach,omitempty"`       //附加数据
	SuccessTime     *string            `json:"success_time,omitempty"` //支付完成时间
	Payer           *PayerRespV3Dto  `json:"payer,omitempty"`        // 支付者信息
	Amount          *AmountRespV3Dto `json:"amount,omitempty"`       //订单金额信息
	SceneInfo       *SceneInfoRespV3Dto `json:"scene_info,omitempty"` //支付场景描述
	PromotionDetail []PromotionDetailRespV3Dto  `json:"promotion_detail,omitempty"`  //优惠功能信息
}


