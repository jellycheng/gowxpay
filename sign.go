package gowxpay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
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

