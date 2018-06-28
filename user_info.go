package wechat

import (
	"fmt"
)

//通过个人access_token获取个人信息
func (info *WechatInfo)GetUserInfoByPerson(token string, openId string) (*UserInfo, error) {
	url := fmt.Sprintf(Url_Oauth_UserInfo, token, openId)
	var result = new(UserInfo)
	b, err := HttpGetJson(url, &result)
	if err != nil {
		return nil, err
	}
	if result.OpenID == "" {
		return nil, ErrNoData(b)
	}
	return result, nil
}

//获取个人信息
func (info *WechatInfo)GetUserInfo(openId string) (*UserInfo, error) {
	result := new(UserInfo)
	url := fmt.Sprintf(Url_UserInfo, info.AccessToken, openId)
	b, err := HttpGetJson(url, &result)
	if err != nil {
		return nil, err
	}
	if result.OpenID == "" {
		return nil, ErrNoData(b)
	}
	return result, nil
}

