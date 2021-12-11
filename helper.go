package gowxpay

import (
	"bufio"
	"crypto/tls"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/ini"
	"golang.org/x/crypto/pkcs12"
	"io"
	"io/ioutil"
	"regexp"
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

// GetCertData 获取证书内容
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
			content := string(token)
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

func SimpleIni2Map(fileName string) map[string]string {
	ret := map[string]string{}
	if gosupport.FileExists(fileName) {
		con,_ := gosupport.FileGetContents(fileName)
		reader := bufio.NewReader(strings.NewReader(con))
		for {
			l, err := reader.ReadString('\n')
			if err!=nil && err!=io.EOF {
				break
			}
			line := strings.TrimSpace(l)
			if len(line) == 0 {
				if err==io.EOF {
					break
				} else {
					continue
				}
			}
			if IsCommentLine(line) {
				continue
			}
			line = string(ini.GetCleanComment([]byte(line)))
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(string(ini.GetCleanComment([]byte(line[i+1 : len(line)]))))
			ret[strings.TrimSpace(line[0:i])] = value
		}
	}

	return ret
}

// IsCommentLine 判断是否注释字符串，以 #;开头的字符就算
func IsCommentLine(str string) bool {
	isMatch,_ := regexp.MatchString("^\\s*[#;]+", str)
	if isMatch {
		return true
	} else {
		return false
	}
}

func PinAuthorizationHeaderVal(mchid string, nonceStr string, timestamp int64, serialNo string, sign string) string {
	str := fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",timestamp="%d",serial_no="%s",signature="%s"`,
		mchid, nonceStr, timestamp, serialNo, sign)

	return str
}

// PinReqMessage 拼接请求签名原文格式: HTTP请求方法\n支付接口URL Path\n请求时间戳\n请求随机串\n请求报文主体\n
func PinReqMessage(method string, urlPath string, timestamp int64, nonce string, body string) string {
	// 签名原文格式
	str := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n",
		method, urlPath, timestamp, nonce, body)

	return str
}

// PinRespMessage 拼接响应客户端签名串，prepayId为预单号,如 prepay_id=xxx
func PinRespMessage(appid string, timeStamp int64, nonceStr string, prepayId string) string {
	str := fmt.Sprintf("%s\n%d\n%s\n%s\n", appid, timeStamp, nonceStr, prepayId)
	return str
}

// SplitLine 打印分割线
func SplitLine(s string, l int)  {
	sl := strings.Repeat(s, l)
	fmt.Println(sl)
}
