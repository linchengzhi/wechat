package wechat

var Info *WechatInfo

func NewWechatInfo(appId, secret string) *WechatInfo {
	info := new(WechatInfo)
	info.AppId = appId
	info.AppSecret = secret
	return info
}

func InitWechatInfo(appId, secret string) {
	Info = NewWechatInfo(appId, secret)
}

//appId为空，则为默认
func NewPay(appId, token, mchId, notifyUrl, tradeType string) *PayConfig {
	Pay := &PayConfig{
		AppId:         appId,
		Token:         token,
		MchId:         mchId,
		NotifyUrl:     notifyUrl,
		PlaceOrderUrl: Url_Pay_Order,
		QueryOrderUrl: Url_Pay_Query,
		TradeType:     tradeType,
	}
	return Pay
}
