package gowxpay

// TradeType 交易类型
type TradeType string

func (m TradeType) Ptr() *TradeType {
	return &m
}

const (
	TradeTypeJsapi TradeType = "JSAPI" //：公众号支付、小程序支付
	TradeTypeNative TradeType = "NATIVE" // 扫码支付
	TradeTypeApp TradeType = "APP" //APP支付
	TradeTypeMicropay TradeType = "MICROPAY" //付款码支付
	TradeTypeMweb TradeType = "MWEB" //H5支付
	TradeTypeFacepay TradeType = "FACEPAY" //刷脸支付
)

// TradeState 交易状态
type TradeState string

func (m TradeState) Ptr() *TradeState {
	return &m
}

const (
	TradeStateSuccess TradeState = "SUCCESS" //支付成功
	TradeStateRefund TradeState = "REFUND"  //转入退款
	TradeStateNotpay TradeState = "NOTPAY" //未支付
	TradeStateClosed TradeState = "CLOSED" //已关闭
	TradeStateRevoked TradeState = "REVOKED" //已撤销（付款码支付）
	TradeStateUserpaying TradeState = "USERPAYING" //用户支付中（付款码支付）
	TradeStatePayerror TradeState = "PAYERROR"  //"支付失败(其他原因，如银行返回失败)
)

// RefundStatus 退款状态
type RefundStatus string

func (m RefundStatus) Ptr() *RefundStatus {
	return &m
}

const (
	RefundStatusSuccess    RefundStatus = "SUCCESS"    //退款成功
	RefundStatusClosed     RefundStatus = "CLOSED"     //退款关闭
	RefundStatusProcessing RefundStatus = "PROCESSING" //退款处理中
	RefundStatusAbnormal   RefundStatus = "ABNORMAL"   //退款异常
)

// RefundChannel 退款渠道类型
type RefundChannel string

func (m RefundChannel) Ptr() *RefundChannel {
	return &m
}

const (
	RefundChannelOriginal      RefundChannel = "ORIGINAL"       //原路退款
	RefundChannelBalance       RefundChannel = "BALANCE"        //退回到余额
	RefundChannelOtherBalance  RefundChannel = "OTHER_BALANCE"  //原账户异常退到其他余额账户
	RefundChannelOtherBankcard RefundChannel = "OTHER_BANKCARD" //原银行卡异常退到其他银行卡
)

// RefundAccount 退款资金账户类型
type RefundAccount string

func (m RefundAccount) Ptr() *RefundAccount {
	return &m
}

const (
	RefundAccountUnsettled RefundAccount = "UNSETTLED" //未结算资金
	RefundAccountAvailable RefundAccount = "AVAILABLE"   // 可用余额
	RefundAccountUnavailable  RefundAccount = "UNAVAILABLE" // 不可用余额
	RefundAccountOperation RefundAccount = "OPERATION" //运营户
	RefundAccountBasic RefundAccount = "BASIC" //基本账户（含可用余额和不可用余额）

)

// PromotionScope 优惠范围
type PromotionScope string

func (m PromotionScope) Ptr() *PromotionScope {
	return &m
}

const (
	PromotionScopeGlobal PromotionScope = "GLOBAL" //全场代金券, 全场优惠类型
	PromotionScopeSingle PromotionScope = "SINGLE" //单品优惠, 单品优惠类型
)

// PromotionType 优惠类型
type PromotionType string

func (m PromotionType) Ptr() *PromotionType {
	return &m
}

const (
	PromotionTypeCoupon   PromotionType = "COUPON" //代金券类型，需要走结算资金的充值型代金券
	PromotionTypeDiscount PromotionType = "DISCOUNT" //优惠券类型，不走结算资金的免充值型优惠券
)

// EventType 通知类型
type EventType string

func (m EventType) Ptr() *EventType {
	return &m
}

const (
	EventTypeTransactionSucc = "TRANSACTION.SUCCESS" // 正向，支付成功通知
	EventTypeRefundSucc = "REFUND.SUCCESS" //退款成功通知
	EventTypeRefundAbnormal = "REFUND.ABNORMAL" //退款异常通知
	EventTypeRefundClosed = "REFUND.CLOSED" //退款关闭通知
)

