package wechat

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestAnalyzeEvent(t *testing.T) {
	data := "<xml><ToUserName><![CDATA[gh_4d3677ff7031]]></ToUserName>\n<FromUserName><![CDATA[oRq-A0Xp4vOT1KNHViki6vbvwIZw]]>" +
		"</FromUserName>\n<CreateTime>1530083791</CreateTime>\n<MsgType><![CDATA[event]]></MsgType>\n<Event><![CDATA[SCAN]]>" +
		"</Event>\n<EventKey><![CDATA[6334fbc4b09f8113163fc2fd20c42e5a]]></EventKey>\n<Ticket><![CDATA[gQEE8DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFx" +
		"LmNvbS9xLzAyTlB1OEJwNkxlMGgxM09XUGhyY1gAAgTGOTNbAwQsAQAA]]></Ticket>\n</xml>"
	Convey("解析普通消息", t, func() {
		info := NewWechatTest()
		result, err := info.AnalyzeMessage([]byte(data))
		fmt.Println(result, err)
		So(err, ShouldEqual, nil)
	})
}