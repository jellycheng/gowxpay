# gowxpay
```
封装微信支付api sdk

直连商户V2接口文档：
    https://pay.weixin.qq.com/wiki/doc/api/index.html
服务商V2接口文档：
    https://pay.weixin.qq.com/wiki/doc/api/sl.html
    
直连商户V3文档：
    https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml
    https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pages/index.shtml
服务商V3接口文档：
    https://pay.weixin.qq.com/wiki/doc/apiv3_partner/index.shtml

名词表：接入模式、支付产品、参数等说明
    https://pay.weixin.qq.com/wiki/doc/apiv3_partner/terms_definition/chapter1_1.shtml

```
[微信支付商户平台：https://pay.weixin.qq.com/](https://pay.weixin.qq.com/) <br>
[https://wechatpay-api.gitbook.io/wechatpay-api-v3/](https://wechatpay-api.gitbook.io/wechatpay-api-v3/) <br>

## 下载与更新依赖
```
go get -u github.com/jellycheng/gowxpay
    或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gowxpay

```

## 微信支付v2与v3版本差异
```
v2版本：
    使用XML作为数据交互格式

v3版本：
    遵循统一的RESTful的设计风格
    使用JSON作为数据交互的格式，不再使用XML
    使用基于非对称密钥的SHA256-RSA的数字签名算法，不再使用MD5或HMAC-SHA256
    不再要求HTTPS客户端证书
    使用AES-256-GCM加密，对回调中的关键信息进行加密保护
    仅支持UTF-8字符编码

```
规则差异   | V2    | V3   | 备注
------------|-----------|-----------|-----------
参数格式| XML |  JSON | 
提交方式|POST | POST、GET、PUT、PATCH、DELETE  | 
回调加密|无需加密 | AES-256-GCM加密  | 
敏感加密| RSA加密|  RSA加密 |     
编码方式|UTF-8 | UTF-8  | 
签名方式| MD5或HMAC-SHA256| 非对称密钥SHA256-RSA  | 


## 注意事项
```
1. 仅有JSAPI支付和Native支付需要在微信支付后台配置支付域名（产品中心->开发配置->支付配置,域名支持http和https的协议），
    APP支付、付款码支付无需配置域名
2. 所有使用JSAPI方式发起支付请求的链接地址，都必须是在微信支付后台配置的支付授权目录之下。
    下单前需要调用【网页授权获取用户信息】接口获取到用户的openid
3. 微信平台接到Native支付请求时，会回调用在微信支付后台配置的支付回调链接，用于传递订单信息
    其它方式的通知地址通过接口入参notify_url设置（支持http、https协议）

JSAPI支付：公众号中拉起支付、小程序支付，需要在商户后台配置JSAPI支付授权目录（最多可添加5个域名地址）
Native支付: 用户扫码支付，即生成二维码提供用户扫码支付，需要在商户后台配置Native支付回调链接，一般用于扫pc网站二维码支付、收银台二维码场景
APP支付： 使用ios、android app发起支付
H5支付：在移动浏览器中的支付
付款码支付：扫用户的"付款码"支付，适用于商超、便利店、餐饮等系统
微信刷脸支付： 用户刷脸支付
小程序支付：使用JSAPI下单场景，申请入口在小程序后台申请支付（https://mp.weixin.qq.com/），关联支付商户号
        小程序后台-》功能-》微信支付

```
[支付场景与用途说明：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/terms_definition/chapter1_1_0.shtml](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/terms_definition/chapter1_1_0.shtml) <br>

## 密钥相关
```
生成公钥： openssl x509 -in apiclient_cert.pem -pubkey -noout > apiclient_pub.pem

```

## 示例
```
调用示例请参考以 _test.go 结尾的文件

```

