package gowxpay

import "time"

// 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
type EncryptCertificateDto struct {
	// 加密所使用的算法，目前可能取值仅为 AEAD_AES_256_GCM
	Algorithm *string `json:"algorithm"`
	// 加密所使用的随机字符串
	Nonce *string `json:"nonce"`
	// 附加数据包（可能为空）
	AssociatedData *string `json:"associated_data"`
	// 证书内容密文，解密后会获得证书完整内容,Base64编码后的密文
	Ciphertext *string `json:"ciphertext"`
}

// 微信支付平台证书信息
type CertificateDto struct {
	// 证书序列号
	SerialNo *string `json:"serial_no"`
	// 证书有效期开始时间
	EffectiveTime *time.Time `json:"effective_time"`
	// 证书过期时间
	ExpireTime *time.Time `json:"expire_time"`
	// 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	EncryptCertificate *EncryptCertificateDto `json:"encrypt_certificate"`
}

// 平台证书列表
type CertificatesRespDto struct {
	Data []CertificateDto `json:"data,omitempty"`
}
