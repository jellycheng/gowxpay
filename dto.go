package gowxpay

import "time"

type PayV3Err struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// 微信支付API v3 版本错误结构
type APIErrDetail struct {
	StatusCode int         // 应答报文的 HTTP 状态码
	Header     map[string]string // 应答报文的 Header 信息
	Body       string      // 应答报文的 Body 原文
	Code       string      `json:"code"`             // 应答报文的 Body 解析后的错误码信息，仅不符合预期/发生系统错误时存在
	Message    string      `json:"message"`          // 应答报文的 Body 解析后的文字说明信息，仅不符合预期/发生系统错误时存在
	Detail     interface{} `json:"detail,omitempty"` // 应答报文的 Body 解析后的详细信息，仅不符合预期/发生系统错误时存在
}

// 微信支付接口响应的请求头
type WechatPayHeader struct {
	RequestID string
	Serial    string
	Signature string
	Nonce     string
	Timestamp int64
}

// 支付者信息，入参
type PayerReqV3Dto struct {
	// 用户在商户appid下的唯一标识。
	Openid *string `json:"openid,omitempty"`
}

// 订单金额信息
type AmountReqV3Dto struct {
	// 订单总金额，单位为分
	Total *int64 `json:"total"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency *string `json:"currency,omitempty"`
}

// 单品列表信息
type GoodsDetailReqV3Dto struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）。
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity *int64 `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice *int64 `json:"unit_price"`
}

// 优惠功能
type DetailReqV3Dto struct {
	// 订单原价
	CostPrice *int64 `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   *string       `json:"invoice_id,omitempty"`
	// 单品列表信息
	GoodsDetail []GoodsDetailReqV3Dto `json:"goods_detail,omitempty"`
}

// 商户门店信息
type StoreInfoReqV3Dto struct {
	// 商户侧门店编号
	Id *string `json:"id"`
	// 商户侧门店名称
	Name *string `json:"name,omitempty"`
	// 地区编码，详细请见微信支付提供的文档
	AreaCode *string `json:"area_code,omitempty"`
	// 详细的商户门店地址
	Address *string `json:"address,omitempty"`
}

// 支付场景描述
type SceneInfoReqV3Dto struct {
	// 用户终端IP
	PayerClientIp *string `json:"payer_client_ip"`
	// 商户端设备号
	DeviceId  *string    `json:"device_id,omitempty"`
	StoreInfo *StoreInfoReqV3Dto `json:"store_info,omitempty"`
}
// 结算信息
type SettleInfoReqV3Dto struct {
	// 是否指定分账
	ProfitSharing *bool `json:"profit_sharing,omitempty"`
}

// 预下单请求参数
type PrepayReqV3Dto struct {
	// 公众号ID、小程序ID
	Appid *string `json:"appid"`
	// 直连商户号
	Mchid *string `json:"mchid"`
	// 商品描述
	Description *string `json:"description"`
	// 商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 订单失效时间，格式为rfc3339格式，交易结束时间
	TimeExpire *time.Time `json:"time_expire,omitempty"`
	// 附加数据
	Attach *string `json:"attach,omitempty"`
	// 通知地址， 有效性：1. HTTPS协议；2. 不允许携带查询串即不能带参数
	NotifyUrl *string `json:"notify_url"`
	// 订单优惠标记，商品标记，代金券或立减优惠功能的参数。
	GoodsTag *string `json:"goods_tag,omitempty"`
	// 订单金额信息
	Amount        *AmountReqV3Dto    `json:"amount"`
	// 支付者信息
	Payer         *PayerReqV3Dto   `json:"payer"`
	// 优惠功能
	Detail        *DetailReqV3Dto     `json:"detail,omitempty"`
	// 场景信息
	SceneInfo     *SceneInfoReqV3Dto  `json:"scene_info,omitempty"`
	// 结算信息
	SettleInfo    *SettleInfoReqV3Dto `json:"settle_info,omitempty"`
}

// 预下单返回的内容
type PrepayRespV3Dto struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"`
}

// 响应小程序拉起支付的参数
type PrepayWithRequestPaymentRespV3Dto struct {
	// 预支付交易会话标识，预单号
	PrepayId *string `json:"prepayId"`
	// 商户号
	PartnerId *string `json:"partnerId"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名
	Sign *string `json:"sign"`
}
