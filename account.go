package gowxpay

import "io/ioutil"

//支付商户号，appid和mch_id两者之间需要具备绑定关系，直连模式：appid和mch_id之间的绑定关系可以是多对多
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

// 设置证书
func (a *Account) SetCertData(certPath string) (*Account,error) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return a,err
	}
	a.certData = certData
	return a,nil
}

//创建微信支付账号，支付商户号
func NewAccount(appID string, mchID string, apiKey string) *Account {
	return &Account{
		appID:     appID,
		mchID:     mchID,
		apiKey:    apiKey,
		isSandbox: false,
	}
}

