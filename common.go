package wechat

import (
	"fmt"
	"errors"
	"time"
	"strconv"
	"crypto/md5"
)

func ErrNoData(b []byte) error {
	tips := fmt.Sprintf("no data result=%s", string(b))
	return errors.New(tips)
}

func newRandStr() string {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 36)
	return fmt.Sprintf("%x", md5.Sum([]byte(nonce)))
}

func NewWechatTest() *WechatInfo {
	info := NewWechatInfo(AppId, AppSecret)
	//info.UpdateAccessToken()
	info.AccessToken = "11_3tTe-swUSvJwomj8RPGf4S7HL7JOxD4myRwB5SXzTs1kB9etlv4mNQSIRhB29jCm0TUTG3ispa1GjVTrLbFz2Urx5fiTz-nvwSuaRKiVXKON11Q9jMNpKQa7JNLhtNWA_U1XHBMMIuJEFBieXCWeADAKPJ"
	return info
}
