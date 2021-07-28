package gowxpay

import (
	"crypto/tls"
	"errors"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/sign"
	"io/ioutil"
	"net/http"
	"strings"
)

// 封装支付api v2版本的请求
type PayClient struct {
	account              *Account //支付商户号配置
	signType             string   //签名类型
	httpConnectTimeoutMs int      //连接超时时间
	httpReadTimeoutMs    int      //读取超时时间
}

// 创建微信支付客户端
func NewPayClient(account *Account) *PayClient {
	return &PayClient{
		account:              account,
		signType:             gosupport.SignTypeMD5,
		httpConnectTimeoutMs: 2000,
		httpReadTimeoutMs:    1000,
	}
}

func (c *PayClient) SetHttpConnectTimeoutMs(ms int) *PayClient {
	c.httpConnectTimeoutMs = ms
	return c
}

func (c *PayClient) SetHttpReadTimeoutMs(ms int) *PayClient {
	c.httpReadTimeoutMs = ms
	return c
}

func (c *PayClient) SetSignType(signType string) *PayClient {
	c.signType = signType
	return c
}

func (c *PayClient) SetAccount(account *Account) *PayClient {
	c.account = account
	return c
}

// 向 params 中添加 appid、mch_id、nonce_str、sign_type、sign 公共参数
func (c *PayClient) AppendRequestData(params MapParams) MapParams {
	params["appid"] = c.account.appID
	params["mch_id"] = c.account.mchID
	//随机字符串，长度要求在32位以内
	params["nonce_str"] = gosupport.GetRandString(10)
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)
	return params
}

//签名
func (c *PayClient) Sign(params MapParams) string {
	signStr, _, _ := sign.WxPaySign(params, c.signType, c.account.apiKey)
	return signStr
}

//验证签名
func (c *PayClient) ValidSign(params MapParams) bool {
	if !params.ContainsKey("sign") {
		return false
	}
	return params.GetString("sign") == c.Sign(params)
}


// https no cert post
func (c *PayClient) Post4NotCert(url string, params MapParams) (string, error) {
	h := &http.Client{}
	p := c.AppendRequestData(params)
	response, err := h.Post(url, "application/xml; charset=utf-8", strings.NewReader(gosupport.Map2XMLV2(p)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// https need cert post
func (c *PayClient) Post4Cert(url string, params MapParams) (string, error) {
	if c.account.certData == nil {
		return "", errors.New("证书数据为空")
	}
	//将pkcs12证书转成pem
	cert,_ := Pkcs12ToPem(c.account.certData, c.account.mchID)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	transport := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	h := &http.Client{Transport: transport}
	p := c.AppendRequestData(params)
	response, err := h.Post(url, "application/xml; charset=utf-8", strings.NewReader(gosupport.Map2XMLV2(p)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}


// 处理 HTTPS API返回数据，转换成Map对象。return_code为SUCCESS时，验证签名。
func (c *PayClient) ProcessResponseXml(xmlStr string) (MapParams, error) {
	var returnCode string
	params := XmlToMap(xmlStr)
	if params.ContainsKey("return_code") {
		returnCode = params.GetString("return_code")
	} else {
		return nil, errors.New("no return_code in XML")
	}

	if returnCode == Fail {
		tmpMsg := "失败"
		if params.ContainsKey("return_msg") {
			tmpMsg = params.GetString("return_msg")
		}
		return params, errors.New(tmpMsg)
	} else if returnCode == Success {
		if c.ValidSign(params) {
			return params, nil
		} else {
			return nil, errors.New("invalid sign value in XML")
		}
	} else {
		return nil, errors.New("return_code value is invalid in XML")
	}
}
