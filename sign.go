package gowxpay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// 通过私钥对字符串以 SHA256WithRSA 算法生成签名信息
func SignSHA256WithRSA(sourceStr string, privateKey *rsa.PrivateKey) (string, error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	h := crypto.Hash.New(crypto.SHA256)
	_, err := h.Write([]byte(sourceStr))
	if err != nil {
		return "", err
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}


// 根据微信支付签名格式构造验签原文： https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func BuildMessage(headerArgs WechatPayHeader, body []byte) string {
	return fmt.Sprintf("%d\n%s\n%s\n", headerArgs.Timestamp, headerArgs.Nonce, string(body))
}

// 验证签名
func CheckSignV3(allHeaders map[string]string, body []byte, certificate *x509.Certificate) error {
	headerArgs, err := GetWechatPayHeaderV3(allHeaders)
	if err != nil {
		return err
	}
	if err := CheckWechatPayHeader(headerArgs);err!=nil{
		return err
	}
	message := BuildMessage(headerArgs, body)
	sigBytes, err := base64.StdEncoding.DecodeString(headerArgs.Signature)
	if err != nil {
		return fmt.Errorf("verify failed: signature not base64 encoded")
	}
	//certificate,_ := LoadCertificateWithPath("./aaa.pem")
	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(certificate.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sigBytes)
	if err != nil {
		return fmt.Errorf("verifty signature with public key err:%s", err.Error())
	}

	return nil
}

