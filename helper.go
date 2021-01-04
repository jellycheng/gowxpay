package gowxpay

import (
	"crypto/tls"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"strings"
)

// 将Pkcs12转成Pem
func Pkcs12ToPem(p12 []byte, password string) (tls.Certificate,error) {
	blocks, err := pkcs12.ToPEM(p12, password)

	if err != nil {
		var cert tls.Certificate
		return cert, err
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	return cert,err
}

//获取证书内容
func GetCertData(certPath string) ([]byte, error){
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return certData, errors.New("读取证书失败:" + certPath)
	}
	return certData, nil
}

func XmlToMap(xmlStr string) MapParams {
	params := make(MapParams)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	var (
		key   string
		value string
	)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			//处理开始标签
			key = token.Name.Local
		case xml.CharData:
			//处理标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
	}

	return params
}
