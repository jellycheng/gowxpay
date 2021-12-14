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
	Plaintext      string // 解密后内容
}

// NotifyDto 微信支付通知结果结构
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
