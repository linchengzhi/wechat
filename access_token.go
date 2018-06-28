package wechat

import (
	"fmt"
	"time"
)

//赋值access_token
func (info *WechatInfo) UpdateAccessToken() error {
	token, err := info.GetAccessToken()
	if err != nil {
		return err
	}
	info.AccessToken = token.AccessToken
	info.AccessTokenExpire = time.Now().Unix() + token.ExpiresIn - 100
	return nil
}

//获取公众号的access_token
func (info *WechatInfo) GetAccessToken() (*AccessTokenRes, error) {
	url := fmt.Sprintf(Url_AccessToken, info.AppId, info.AppSecret)
	var result = new(AccessTokenRes)
	b, err := HttpGetJson(url, &result)
	if err != nil {
		return nil, err
	}
	if result.AccessToken == "" {
		return nil, ErrNoData(b)
	}
	return result, nil
}

//获取个人的access_token
func (info *WechatInfo)GetUserAccessToken(code string) (*UserAccessTokenRes, error) {
	url := fmt.Sprintf(Url_Oauth_AccessToken, info.AppId, info.AppSecret, code)
	var result = new(UserAccessTokenRes)
	b, err := HttpGetJson(url, &result)
	if err != nil {
		return nil, err
	}
	if result.AccessToken == "" {
		return nil, ErrNoData(b)
	}
	return result, nil
}