package wechat

//import (
//	"testing"
//	"fmt"
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func TestGetQRCode(t *testing.T) {
//	info := NewWechatTest()
//	Convey("get limit qr code", t, func() {
//		result, err := info.GetLimitQRCode(1)
//		fmt.Println(result)
//		So(err, ShouldEqual, nil)
//
//		result, err = info.GetLimitQRCode("str")
//		fmt.Println(result)
//		So(err, ShouldEqual, nil)
//	})
//
//	Convey("get no limit qr code", t, func() {
//		result, err := info.GetNoLimitQRCode(1)
//		fmt.Println(result, err)
//		So(err, ShouldEqual, nil)
//
//		result, err = info.GetNoLimitQRCode(100, "str")
//		fmt.Println(result, err)
//		So(err, ShouldEqual, nil)
//	})
//}
