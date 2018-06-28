package wechat

import "encoding/xml"

const (
	Url_Oauth_AccessToken = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	Url_Oauth_UserInfo    = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"

	Url_Pay_Order = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	Url_Pay_Query = "https://api.mch.weixin.qq.com/pay/orderquery"

	Url_AccessToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	Url_QRCode      = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	Url_UserInfo    = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

const (
	QR_Scene           = "QR_SCENE"
	QR_Str_Scene       = "QR_STR_SCENE"
	QR_Limit_Scene     = "QR_LIMIT_SCENE"
	QR_Limit_Str_Scene = "QR_LIMIT_STR_SCENE"
)

const (
	Pay_Type_App = "APP"
	Pay_Type_JS = "JSAPI"
	Pay_Type_QRCode = "NATIVE"
)

type WechatInfo struct {
	AppId             string
	AppSecret         string
	AccessToken       string
	AccessTokenExpire int64
	Pay               *PayConfig
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type QrCode struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	Url           string `json:"url"`
}

type UserAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

type UserInfo struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        uint8    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

type PushBase struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string	//开发者微信号
	FromUserName string	//发送方帐号（一个OpenID）
	CreateTime   int64	//消息创建时间
	MsgType      string	//消息类型
}

//事件
type PushEvent struct {
	PushBase
	Event        string
	EventKey     string
	Ticket       string
	Latitude	 float64
	Longitude	 float64
	Precision    int
}

//消息
type PushMessage struct {
	PushBase
	MsgId        string	//	消息id，64位整型
	Content      string
	PicUrl       string
	MediaId      string
	Format       string
	Recognition  string
	ThumbMediaId string
	Location_X   float64
	Location_Y   float64
	Scale        int
	Label        string
	Title		 string
	Description  string
	Url          string
}

type PushAll struct {
	PushBase
	Event        string
	EventKey     string
	Ticket       string
	Latitude	 float64
	Longitude	 float64
	Precision    int
	MsgId        string	//	消息id，64位整型
	Content      string
	PicUrl       string
	MediaId      string
	Format       string
	Recognition  string
	ThumbMediaId string
	Location_X   float64
	Location_Y   float64
	Scale        int
	Label        string
	Title		 string
	Description  string
	Url          string
}

//支付结构
type PayConfig struct {
	AppId         string
	Token         string
	MchId         string
	NotifyUrl     string
	PlaceOrderUrl string
	QueryOrderUrl string
	TradeType     string
}

//订单参数
type Order struct {
	XMLName        xml.Name `xml:"xml"`
	AppId          string   `xml:"appid"`                       //应用ID
	MchId          string   `xml:"mch_id"`                      //商户号
	DeviceInfo     string   `xml:"device_info, omitempty"`      //设备号
	NonceStr       string   `xml:"nonce_str, omitempty"`        //随机字符串
	Sign           string   `xml:"sign, omitempty"`             //签名
	SignType       string   `xml:"sign_type, omitempty"`        //签名类型
	Body           string   `xml:"body, omitempty"`             //商品描述
	Detail         string   `xml:"detail, omitempty"`           //商品详情
	Attach         string   `xml:"attach, omitempty"`           //附加数据 携带订单的自定义数据
	OutTradeNo     string   `xml:"out_trade_no, omitempty"`     //商户系统内部订单号
	FeeType        string   `xml:"fee_type, omitempty"`         //货币类型  CNY
	TotalFee       string   `xml:"total_fee, omitempty"`        //订单总金额
	SpbillCreateIp string   `xml:"spbill_create_ip, omitempty"` //终端IP
	TimeStart      string   `xml:"time_start, omitempty"`       //交易起始时间
	TimeExpire     string   `xml:"time_expire, omitempty"`      //交易起始时间
	GoodsTag       string   `xml:"goods_tag, omitempty"`        //交易结束时间
	NotifyUrl      string   `xml:"notify_url, omitempty"`       //通知地址
	TradeType      string   `xml:"trade_type, omitempty"`       //交易类型 JSAPI--公众号支付、NATIVE--原生扫码支付、APP--app支付
	ProductId      string   `xml:"product_id, omitempty"`       //商品ID
	Openid         string   `xml:"openid, omitempty"`           //用户标识
	LimitPay       string   `xml:"limit_pay, omitempty"`        //指定支付方式
	SceneInfo      string   `json:"scene_info, omitempty"`      //场景信息
}

//下单返回结果
type OrderResult struct {
	XMLName       xml.Name `xml:"xml"`
	ReturnCode    string   `xml:"return_code"`
	ReturnMsg     string   `xml:"return_msg"`
	ErrCode       string   `xml:"err_code"`
	ErrCodeDesc   string   `xml:"err_code_des"`
	ResultCode    string   `xml:"result_code"`
	AppId         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	DeviceInfo    string   `xml:"device_info"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	TradeType     string   `xml:"trade_type"`
	PrepayId      string   `xml:"prepay_id"`
	IsSubscribe   string   `xml:"is_subscribe"`
	CodeUrl   	  string   `xml:"code_url"`	//二维码链接

	Attach        string   `xml:"attach"`
	CashFee       string   `xml:"cash_fee"`
	FeeType       string   `xml:"fee_type"`
	SignType      string   `xml:"sign_type"`
	OpenId        string   `xml:"openid"`
	TotalFee      string   `xml:"total_fee"`
	BankType      string   `xml:"bank_type"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	TimeEnd       string   `xml:"time_end"`
	MwebUrl       string   `xml:"mweb_url"`
}

//查询返回结果
type QueryResult struct {
	XMLName        xml.Name `xml:"xml"`
	AppId          string   `xml:"appid"`
	Attach         string   `xml:"attach"`
	BankType       string   `xml:"bank_type"`
	CashFee        string   `xml:"cash_fee"`
	CashFeeType    string   `xml:"cash_fee_type"`
	CouponFee      string   `xml:"coupon_fee"`
	FeeType        string   `xml:"fee_type"`
	MchId          string   `xml:"mch_id"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	ResultCode     string   `xml:"result_code"`
	ErrCode        string   `xml:"err_code"`
	ErrCodeDesc    string   `xml:"err_code_des"`
	DeviceInfo     string   `xml:"device_info"`
	OpenId         string   `xml:"openid"`
	IsSubscribe    string   `xml:"is_subscribe"`
	TradeType      string   `xml:"trade_type"`
	TradeState     string   `xml:"trade_state"`
	TradeStateDesc string   `xml:"trade_state_desc"`
	TotalFee       string   `xml:"total_fee"`
	CouponCount    string   `xml:"coupon_count"`
	TransactionId  string   `xml:"transaction_id"`
	OrderId        string   `xml:"out_trade_no"`
	ReturnCode     string   `xml:"return_code"`
	ReturnMsg      string   `xml:"return_msg"`
	TimeEnd        string   `xml:"time_end"`
}

