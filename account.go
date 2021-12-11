package gowxpay

import "io/ioutil"

// Account 支付商户号，appid和mch_id两者之间需要具备绑定关系，直连模式：appid和mch_id之间的绑定关系可以是多对多
type Account struct {
	appID     string //小程序appid
	mchID     string //支付商户号
	apiKey    string //API密钥
	certData  []byte //p12证书内容
	isSandbox bool
}

func (a *Account) SetIsSandbox(isSandbox bool) *Account {
	a.isSandbox = isSandbox
	return a
}

func (a *Account) GetIsSandbox() bool {
	return a.isSandbox
}

// SetCertData 设置证书
func (a *Account) SetCertData(certPath string) (*Account,error) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return a,err
	}
	a.certData = certData
	return a,nil
}

// EditCertData 修改证书内容
func (a *Account) EditCertData(certCon []byte) *Account {
	a.certData = certCon
	return a
}

// NewAccount 创建微信支付账号，支付商户号
func NewAccount(appID string, mchID string, apiKey string) *Account {
	return &Account{
		appID:     appID,
		mchID:     mchID,
		apiKey:    apiKey,
		isSandbox: false,
	}
}


type AccountV3 struct {
	AppID    string // 小程序等appid
	MchID    string // 支付商户号
	SerialNo string // 支付商户api证书序列号
	ApiClientKeyPemFile string // apiclient_key.pem文件
	ApiClientKeyCertFile string // apiclient_cert.pem文件
	ApiV3Key string // api v3 key
}
