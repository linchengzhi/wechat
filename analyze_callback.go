package wechat

import "encoding/xml"

//不知道返回的是什么消息
func AnalyzeAll(str []byte) (*PushAll, error) {
	result := new(PushAll)
	if err := xml.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	if result.FromUserName == "" {
		return nil, ErrNoData(str)
	}
	return result, nil
}

//普通消息
func AnalyzeMessage(str []byte) (*PushMessage, error) {
	result := new(PushMessage)
	if err := xml.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	if result.FromUserName == "" {
		return nil, ErrNoData(str)
	}
	return result, nil
}
//事件消息
func AnalyzeEvent(str []byte) (*PushEvent, error) {
	result := new(PushEvent)
	if err := xml.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	if result.FromUserName == "" {
		return nil, ErrNoData(str)
	}
	return result, nil
}
