package gowxpay

import "time"

type PayV3Err struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// APIErrDetail 微信支付API v3 版本错误结构
type APIErrDetail struct {
	StatusCode int         // 应答报文的 HTTP 状态码
	Header     map[string]string // 应答报文的 Header 信息
	Body       string      // 应答报文的 Body 原文
	Code       string      `json:"code"`             // 应答报文的 Body 解析后的错误码信息，仅不符合预期/发生系统错误时存在
	Message    string      `json:"message"`          // 应答报文的 Body 解析后的文字说明信息，仅不符合预期/发生系统错误时存在
	Detail     interface{} `json:"detail,omitempty"` // 应答报文的 Body 解析后的详细信息，仅不符合预期/发生系统错误时存在
}

// WechatPayHeader 微信支付接口响应的请求头
type WechatPayHeader struct {
	RequestID string
	Serial    string
	Signature string
	Nonce     string
	Timestamp int64
}

// PayerReqV3Dto 支付者信息，入参
type PayerReqV3Dto struct {
	// 用户在商户appid下的唯一标识。
	Openid *string `json:"openid,omitempty"`
}

// PayerRespV3Dto 支付者信息，出参
type PayerRespV3Dto struct {
	// 用户在商户appid下的唯一标识。
	Openid *string `json:"openid,omitempty"`
}

// AmountReqV3Dto 订单金额信息，入参
type AmountReqV3Dto struct {
	// 订单总金额，单位为分
	Total *int64 `json:"total"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency *string `json:"currency,omitempty"`
}

// AmountRespV3Dto 订单金额信息，出参
type AmountRespV3Dto struct {
	// 订单总金额，单位为分
	Total *int64 `json:"total,omitempty"`
	// 用户实付金额，单位分，订单总金额Total - 优惠券金额
	PayerTotal *int64  `json:"payer_total,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency *string `json:"currency,omitempty"`
	// 用户支付币种 CNY：人民币
	PayerCurrency *string `json:"payer_currency,omitempty"`
}

// GoodsDetailReqV3Dto 单品列表信息
type GoodsDetailReqV3Dto struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称
	GoodsName *string `json:"goods_name,omitempty"`
	// 用户购买的数量
	Quantity *int64 `json:"quantity"`
	// 商品单价，单位为分
	UnitPrice *int64 `json:"unit_price"`
}

// PromotionGoodsDetailRespV3Dto 单品列表，出参
type PromotionGoodsDetailRespV3Dto struct {
	// 商品编码
	GoodsId *string `json:"goods_id"`
	// 商品数量
	Quantity *int64 `json:"quantity"`
	// 商品价格
	UnitPrice *int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount *int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark *string `json:"goods_remark,omitempty"`
}

// DetailReqV3Dto 优惠功能，入参
type DetailReqV3Dto struct {
	// 订单原价
	CostPrice *int64 `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   *string       `json:"invoice_id,omitempty"`
	// 单品列表信息
	GoodsDetail []GoodsDetailReqV3Dto `json:"goods_detail,omitempty"`
}

// PromotionDetailRespV3Dto 优惠功能，出参
type PromotionDetailRespV3Dto struct {
	// 券ID
	CouponId *string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name *string `json:"name,omitempty"`
	// 优惠范围，GLOBAL：全场代金券；SINGLE：单品优惠
	Scope *string `json:"scope,omitempty"`
	// 优惠类型，CASH：充值；NOCASH：预充值。
	Type *string `json:"type,omitempty"`
	// 优惠券面额
	Amount *int64 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockId *string `json:"stock_id,omitempty"`
	// 微信出资，单位为分
	WechatpayContribute *int64 `json:"wechatpay_contribute,omitempty"`
	// 商户出资，单位为分
	MerchantContribute *int64 `json:"merchant_contribute,omitempty"`
	// 其他出资，单位为分
	OtherContribute *int64 `json:"other_contribute,omitempty"`
	// 优惠币种，CNY：人民币，境内商户号仅支持人民币。
	Currency    *string                `json:"currency,omitempty"`
	// 单品列表
	GoodsDetail []PromotionGoodsDetailRespV3Dto `json:"goods_detail,omitempty"`
}

// StoreInfoReqV3Dto 商户门店信息
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

// H5InfoReqV3Dto H5场景信息
type H5InfoReqV3Dto struct {
	Type string `json:"type"` //场景类型
	AppName string `json:"app_name,omitempty"` //应用名称
	AppUrl string `json:"app_url,omitempty"` //网站URL
	BundleId string `json:"bundle_id,omitempty"` //iOS平台BundleID
	PackageName string `json:"package_name,omitempty"` //Android平台PackageName
}

// SceneInfoReqV3Dto 支付场景描述,入参
type SceneInfoReqV3Dto struct {
	// 用户终端IP
	PayerClientIp *string `json:"payer_client_ip"`
	// 商户端设备号
	DeviceId  *string    `json:"device_id,omitempty"`
	StoreInfo *StoreInfoReqV3Dto `json:"store_info,omitempty"`
	H5Info *H5InfoReqV3Dto `json:"h5_info,omitempty"`
}

// SceneInfoRespV3Dto 支付场景描述,出参
type SceneInfoRespV3Dto struct {
	// 商户端设备号
	DeviceId  *string    `json:"device_id,omitempty"`
}

// SettleInfoReqV3Dto 结算信息
type SettleInfoReqV3Dto struct {
	// 是否指定分账
	ProfitSharing *bool `json:"profit_sharing,omitempty"`
}

// PrepayReqV3Dto 预下单请求参数
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
	Amount        *AmountReqV3Dto   `json:"amount"`
	// 支付者信息
	Payer         *PayerReqV3Dto   `json:"payer,omitempty"`
	// 优惠功能
	Detail        *DetailReqV3Dto     `json:"detail,omitempty"`
	// 场景信息
	SceneInfo     *SceneInfoReqV3Dto  `json:"scene_info,omitempty"`
	// 结算信息
	SettleInfo    *SettleInfoReqV3Dto `json:"settle_info,omitempty"`
}

// PrepayRespV3Dto 预下单返回的内容
type PrepayRespV3Dto struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"`
}

// PrepayWithRequestPaymentRespV3Dto 响应小程序拉起支付的参数
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
