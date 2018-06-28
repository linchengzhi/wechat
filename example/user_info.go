package example

import (
	"encoding/xml"
	"log"
	"github.com/linchengzhi/wechat"
)

const (
	TempQRCode_Limit = 300
)

var Info *wechat.WechatInfo
var Info2 *wechat.WechatInfo

func InitWeixin(appId, secret string) {
	Info = wechat.NewWechatInfo(appId, secret)
	if err := Info.UpdateAccessToken(); err != nil {
		log.Fatal(err)
	}
}

func InitWeixin2(appId, secret string) {
	Info2 = wechat.NewWechatInfo(appId, secret)
	if err := Info2.UpdateAccessToken(); err != nil {
		log.Fatal(err)
	}
}

func GetTempQRCode(sid string) (string, error) {
	if err := Info.UpdateAccessToken(); err != nil {
		log.Printf("获取access_token失败 err=%v", err)
		return "", err
	}
	code, err := Info.GetNoLimitQRCode(TempQRCode_Limit, sid)
	if err != nil {
		log.Printf("获取临时二维码失败 err=%v", err)
		return "", err
	}
	return code.Url, nil
}

func GetUserInfo(openId string) (*wechat.UserInfo, error) {
	if err := Info.UpdateAccessToken(); err != nil {
		log.Printf("获取access_token失败 err=%v", err)
		return nil, err
	}
	Info.UpdateAccessToken()
	userInfo, err := Info.GetUserInfo(openId)
	if err != nil {
		log.Printf("获取人员信息失败")
	}
	return userInfo, nil
}

func GetUserInfoPersonal(code string) (*wechat.UserInfo, error) {
	token, err := Info2.GetUserAccessToken(code)
	if err != nil {
		log.Printf("获取个人token失败 err=%v", err)
		return nil, err
	}
	userInfo, err := Info2.GetUserInfoByPerson(token.AccessToken, token.OpenID)
	if err != nil {
		log.Printf("获取个人信息失败 err=%v", err)
		return nil, err
	}
	return userInfo, nil
}

func AnalyzeEvent(str []byte) (*wechat.PushEvent, error) {
	event := new(wechat.PushEvent)
	if err := xml.Unmarshal(str, &event); err != nil {
		log.Printf("解析公众号消息错误 err:=%v, str=%v", err, string(str))
		return nil, err
	}
	return event, nil
}

