package wechat

import (
	"strconv"
	"errors"
	"fmt"
	"time"
	"math/rand"
)

//永久二维码 自定义参数
func (info *WechatInfo)GetLimitQRCode(i ... interface{}) (*QrCode, error){
	param := []byte{}
	if len(i) == 0 {
		i = append(i, rand.Int())
	}
	switch i[0].(type) {
	case int:
		s := strconv.FormatInt(int64(i[0].(int)), 10)
		param = []byte(`{"action_name": "` + QR_Limit_Scene + `", "action_info": {"scene": {"scene_id": "` + s + `"}}}`)
	case string:
		param = []byte(`{"action_name": "` + QR_Limit_Str_Scene + `", "action_info": {"scene": {"scene_str": "` + i[0].(string) + `"}}}`)
	default:
		return nil, errors.New("param type not int or string")
	}

	return info.getQRCode(param)
}

//临时二维码 有效时间 自定义参数
func (info *WechatInfo)GetNoLimitQRCode(t int64, i ...interface{}) (*QrCode, error) {
	param := []byte{}
	if len(i) == 0 {
		i = append(i, int(time.Now().Unix()))
	}
	expire := strconv.FormatInt(t, 10)
	switch i[0].(type) {
	case int:
		s := strconv.FormatInt(int64(i[0].(int)), 10)
		param = []byte(`{"expire_seconds":` + expire + `,"action_name": "` + QR_Scene + `", "action_info": {"scene": {"scene_id": "` + s + `"}}}`)
	case string:
		param = []byte(`{"expire_seconds":` + expire + `,"action_name": "` + QR_Str_Scene + `", "action_info": {"scene": {"scene_str": "` + i[0].(string) + `"}}}`)
	default:
		return nil, errors.New("param type not int or string")
	}
	return info.getQRCode(param)
}

func (info *WechatInfo)getQRCode(param []byte) (*QrCode, error) {
	url := fmt.Sprintf(Url_QRCode, info.AccessToken)
	result := new(QrCode)
	b, err := HttpPostJson(url, param, &result)
	if err != nil {
		return nil, err
	}
	if result.Url == "" {
		return nil, ErrNoData(b)
	}
	return result, nil
}