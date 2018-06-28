package wechat

import "encoding/xml"

func (info *WechatInfo) AnalyzeAll(str []byte, v interface{}) (*PushAll, error) {
	result := new(PushAll)
	if err := xml.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	if result.FromUserName == "" {
		return nil, ErrNoData(str)
	}
	return result, nil
}

func (info *WechatInfo) AnalyzeMessage(str []byte) (*PushMessage, error) {
	result := new(PushMessage)
	if err := xml.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	if result.FromUserName == "" {
		return nil, ErrNoData(str)
	}
	return result, nil
}

func (info *WechatInfo) AnalyzeEvent(str []byte) (*PushEvent, error) {
	result := new(PushEvent)
	if err := xml.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	if result.FromUserName == "" {
		return nil, ErrNoData(str)
	}
	return result, nil
}
