package wechat

//import (
//	"testing"
//	. "github.com/smartystreets/goconvey/convey"
//	"fmt"
//)
//
//func TestGetAccessToken(t *testing.T) {
//	info := NewWechatInfo(AppId, AppSecret)
//	code := "011RfeyT0zFrtY16WBuT0iRWxT0RfeyS"
//	Convey("获取access_token", t, func() {
//		result, err := info.GetAccessToken()
//		fmt.Println(result, err)
//		So(err, ShouldEqual, nil)
//	})
//
//	Convey("获取user access_token", t, func() {
//		result, err := info.GetUserAccessToken(code)
//		fmt.Println(result, err)
//		So(err, ShouldNotEqual, nil)
//	})
//}
