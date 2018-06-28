package example

import (
	"log"
	"github.com/linchengzhi/wechat"
	"strconv"
)

var Pay *wechat.PayConfig

func InitPay(appId, token, mchId, notifyUrl string) {
	Pay = wechat.NewPay(appId, token, mchId, notifyUrl, wechat.Pay_Type_QRCode)
}

func WeixinPay(orderId string, amount int64, desc string, ip string) (*wechat.OrderResult, error) {
	order, err := Pay.SubmitOrder(orderId, strconv.FormatInt(amount, 10), desc, ip)
	if err != nil || order.CodeUrl == "" {
		log.Printf("微信支付发起失败", err)
		return nil, err
	}
	return order, nil
}

func Sign(m map[string]string) string {
	return wechat.Sign(m, Pay.Token)
}

