package gowxpay

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/curl"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//封装v3版本接口

// JsapiPrepayV3 jsapi下单-小程序下单
func JsapiPrepayV3(dto PrepayReqV3Dto, acc AccountV3) (string, map[string]string, error) {
	var (
		urlStr = PayDomainUrl + "/v3/pay/transactions/jsapi"
		respContent = ""
		allHeaders = map[string]string{}
	)
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	rawPostBodyData := gosupport.ToJson(dto)
	reqStr := PinReqMessage(http.MethodPost, urlPath, timestamp, nonce, rawPostBodyData)
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
					"Accept": "*/*",
					"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
					"Authorization": authorizationHeader,
				}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).SetPostType("json").SetRawPostData(rawPostBodyData).Post()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()

	// 返回结果
	respContent = resp.GetBody()
	return respContent, allHeaders, nil
}

// GetCertificatesV3 获取平台证书列表: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
func GetCertificatesV3(acc AccountV3)(string, map[string]string, error) {
	var (
		urlStr = PayDomainUrl + "/v3/certificates"
		respContent = ""
		allHeaders = map[string]string{}
	)
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	reqStr := PinReqMessage(http.MethodGet, urlPath, timestamp, nonce, "")
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
		"Accept": "*/*",
		"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
		"Authorization": authorizationHeader,
	}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).Get()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()

	// 返回结果
	respContent = resp.GetBody()
	return respContent, allHeaders, nil

}

func GetWechatPayHeaderV3(allheaders map[string]string) (WechatPayHeader, error) {
	requestID := strings.TrimSpace(allheaders[RequestID])
	getHeaderString := func(key string) (string, error) {
		val := strings.TrimSpace(allheaders[key])
		if val == "" {
			return "", fmt.Errorf("`%s` is empty, Request-Id=[%s]", key, requestID)
		}
		return val, nil
	}

	getHeaderInt64 := func(key string) (int64, error) {
		val, err := getHeaderString(key)
		if err != nil {
			return 0, nil
		}
		ret, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid `%s` in header, Request-Id=[%s], err:%w", key, requestID, err)
		}
		return ret, nil
	}

	ret := WechatPayHeader{
		RequestID: requestID,
	}
	var err error

	if ret.Serial, err = getHeaderString(WechatPaySerial); err != nil {
		return ret, err
	}

	if ret.Timestamp, err = getHeaderInt64(WechatPayTimestamp); err != nil {
		return ret, err
	}

	if ret.Nonce, err = getHeaderString(WechatPayNonce); err != nil {
		return ret, err
	}

	if ret.Signature, err = getHeaderString(WechatPaySignature); err != nil {
		return ret, err
	}

	return ret, nil

}

// CheckWechatPayHeader 检查请求头的Timestamp 与当前时间之差不得超过 FiveMinute
func CheckWechatPayHeader(args WechatPayHeader) error {
	if math.Abs(float64(time.Now().Unix()-args.Timestamp)) >= FiveMinute {
		return fmt.Errorf("timestamp=[%d] expires, Request-Id=[%s]", args.Timestamp, args.RequestID)
	}
	return nil
}

// QueryOrder4TransactionId 微信支付订单号查询
func QueryOrder4TransactionId(q QueryOrderReqDto, acc AccountV3) (string, map[string]string, error) {
	var (
		urlStr = fmt.Sprintf("%s/v3/pay/transactions/id/%s?mchid=%s", PayDomainUrl, *q.TransactionId, *q.Mchid)
		respContent = ""
		allHeaders = map[string]string{}
	)
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	reqStr := PinReqMessage(http.MethodGet, urlPath, timestamp, nonce, "")
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
		"Accept": "*/*",
		"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
		"Authorization": authorizationHeader,
	}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).Get()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()

	// 返回结果
	respContent = resp.GetBody()
	return respContent, allHeaders, nil
}

// QueryOrder4OutTradeNo 商户订单号查询
func QueryOrder4OutTradeNo(q QueryOrderReqDto, acc AccountV3) (string, map[string]string, error) {
	var (
		urlStr = fmt.Sprintf("%s/v3/pay/transactions/out-trade-no/%s?mchid=%s", PayDomainUrl, *q.OutTradeNo, *q.Mchid)
		respContent = ""
		allHeaders = map[string]string{}
	)
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	reqStr := PinReqMessage(http.MethodGet, urlPath, timestamp, nonce, "")
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
		"Accept": "*/*",
		"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
		"Authorization": authorizationHeader,
	}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).Get()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()

	// 返回结果
	respContent = resp.GetBody()
	return respContent, allHeaders, nil
}

// CloseOrder 关闭订单
func CloseOrder(q CloseOrderReqDto, acc AccountV3) (bool, map[string]string, error) {
	var (
		urlStr = fmt.Sprintf("%s/v3/pay/transactions/out-trade-no/%s/close", PayDomainUrl, *q.OutTradeNo)
		respContent = false
		allHeaders = map[string]string{}
	)
	if q.OutTradeNo == nil {
		return respContent, allHeaders, fmt.Errorf("`OutTradeNo` is required")
	}
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	dto := ClosePostBodyReqDto{
		Mchid:q.Mchid,
	}
	rawPostBodyData := gosupport.ToJson(dto)

	reqStr := PinReqMessage(http.MethodPost, urlPath, timestamp, nonce, rawPostBodyData)
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
		"Accept": "*/*",
		"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
		"Authorization": authorizationHeader,
	}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).SetPostType("json").SetRawPostData(rawPostBodyData).Post()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()
	if resp.GetRaw().StatusCode >= 200 && resp.GetRaw().StatusCode <= 299 {
		respContent = true
	}
	return respContent, allHeaders, nil

}

// RefundOrder 退款
func RefundOrder(reqDto RefundReqV3Dto, acc AccountV3) (string, map[string]string, error) {
	var (
		urlStr = PayDomainUrl + "/v3/refund/domestic/refunds"
		respContent = ""
		allHeaders = map[string]string{}
	)
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	rawPostBodyData := gosupport.ToJson(reqDto)

	reqStr := PinReqMessage(http.MethodPost, urlPath, timestamp, nonce, rawPostBodyData)
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
		"Accept": "*/*",
		"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
		"Authorization": authorizationHeader,
	}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).SetPostType("json").SetRawPostData(rawPostBodyData).Post()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()
	respContent = resp.GetBody()
	return respContent, allHeaders, nil
}

// RefundQuery 退款查询
func RefundQuery(reqDto QueryByOutRefundNoReqV3Dto, acc AccountV3) (string, map[string]string, error) {
	var (
		urlStr = fmt.Sprintf("%s/v3/refund/domestic/refunds/%s", PayDomainUrl, *reqDto.OutRefundNo)
		respContent = ""
		allHeaders = map[string]string{}
	)
	varUrl,_ := url.Parse(urlStr)
	urlPath := varUrl.RequestURI()
	timestamp := gosupport.Time()
	nonce := gosupport.GetRandString(8)

	rawPostBodyData := ""

	reqStr := PinReqMessage(http.MethodGet, urlPath, timestamp, nonce, rawPostBodyData)
	privateKey, err := LoadPrivateKeyWithPath(acc.ApiClientKeyPemFile)
	if err != nil {
		return respContent, allHeaders, err
	}
	sign, _ := SignSHA256WithRSA(reqStr, privateKey)
	authorizationHeader := PinAuthorizationHeaderVal(acc.MchID, nonce, timestamp, acc.SerialNo, sign)

	headers := map[string]string{
		"Accept": "*/*",
		"User-Agent": gosupport.GenerateUserAgent(PaySdkName, PaySdkVersion),
		"Authorization": authorizationHeader,
	}
	reqObj := curl.NewHttpRequest()
	resp, err := reqObj.SetUrl(urlStr).SetTimeout(int64(DefaultTimeout)).SetHeaders(headers).Get()
	if err != nil{
		return respContent, allHeaders, err
	}
	// 获取响应头
	allHeaders = resp.GetHeaders()
	respContent = resp.GetBody()
	return respContent, allHeaders, nil

}
