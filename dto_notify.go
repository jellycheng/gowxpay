package gowxpay

import (
	"time"
)

// EncryptedResourceDto 通知资源数据：微信支付通知请求中的内容
type EncryptedResourceDto struct {
	Algorithm      string `json:"algorithm"` //加密算法类型
	Ciphertext     string `json:"ciphertext"` //数据密文
	AssociatedData string `json:"associated_data"` //附加数据
	OriginalType   string `json:"original_type"` //原始类型
	Nonce          string `json:"nonce"` //随机串
	Plaintext      string // 解密后内容,正向
}

// NotifyDto 微信支付通知结果结构,正向
type NotifyDto struct {
	ID           string             `json:"id"` //通知ID
	CreateTime   *time.Time         `json:"create_time"` //通知创建时间
	EventType    string             `json:"event_type"`  //通知类型
	ResourceType string             `json:"resource_type"` //通知数据类型
	Resource     *EncryptedResourceDto `json:"resource"` // 通知资源数据
	Summary      string             `json:"summary"` //回调摘要
}

// RefundEncryptedResourceDto 微信退款通知资源数据
type RefundEncryptedResourceDto struct {
	Algorithm      string `json:"algorithm"` //加密算法类型，AEAD_AES_256_GCM
	Ciphertext     string `json:"ciphertext"` //数据密文
	AssociatedData string `json:"associated_data"` //附加数据，refund
	OriginalType   string `json:"original_type"` //原始类型，refund
	Nonce          string `json:"nonce"` //随机串
	Plaintext      string // 解密后内容
}

// RefundNotifyDto 微信退款通知结果结构
type RefundNotifyDto struct {
	ID           string             `json:"id"` //通知ID,唯一
	CreateTime   *time.Time         `json:"create_time"` //通知创建时间,2021-12-13T18:08:31+08:00
	EventType    string             `json:"event_type"`  //通知类型 REFUND.SUCCESS
	ResourceType string             `json:"resource_type"` //通知数据类型 encrypt-resource
	Resource     *RefundEncryptedResourceDto `json:"resource"` // 通知资源数据
	Summary      string             `json:"summary"` //回调摘要 退款成功
}

// RefundNotifyResourceAmountDto 退款通知接口，resource字段解析出来amount退款金额信息
type RefundNotifyResourceAmountDto struct {
	Total       int64 `json:"total"` // 订单总金额，单位为分
	Refund      int64 `json:"refund"` // 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额，如果有使用券，后台会按比例退
	PayerTotal  int64 `json:"payer_total"` // 用户实际支付金额，单位为分，只能为整数
	PayerRefund int64 `json:"payer_refund"` // 退款给用户的金额，不包含所有优惠券金额
}

// RefundNotifyResourceDto 退款通知接口，resource字段解析出来的内容
type RefundNotifyResourceDto struct {
	Mchid         string    `json:"mchid"` // 直连商户号
	OutTradeNo    string    `json:"out_trade_no"` //商户订单号
	TransactionId string    `json:"transaction_id"` //微信支付订单号，正向单号
	OutRefundNo   string    `json:"out_refund_no"`  //商户退款单号
	RefundId      string    `json:"refund_id"`     // 微信支付退款单号
	RefundStatus  string    `json:"refund_status"` // 退款状态，对应常量值 RefundStatusSuccess
	SuccessTime   *time.Time `json:"success_time"`  // 退款成功时间
	Amount        *RefundNotifyResourceAmountDto `json:"amount"` // 退款金额信息
	UserReceivedAccount string `json:"user_received_account"`  //退款入账账户
}
