package wechat

import "fmt"

//提交订单
func (pay *PayConfig) SubmitOrder(orderId, amount, desc, clientIp string) (*OrderResult, error) {
	param := pay.SetOrderParam(orderId, amount, desc, clientIp, pay.TradeType)
	signStr := Sign(param, pay.Token)
	param["sign"] = signStr
	xmlParam := MapToXmlString(param)
	order, err := pay.SendOrder(xmlParam)
	if err != nil {
		return nil, err
	}
	resultInMap, err := ToMap(order)
	if err != nil {
		return nil, err
	}
	wantSign := Sign(resultInMap, pay.Token)
	gotSign := resultInMap["sign"]
	if wantSign != gotSign {
		return nil, fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}
	if order.ReturnCode != "SUCCESS" {
		return nil, fmt.Errorf("return code:%s, return desc:%s", order.ReturnCode, order.ReturnMsg)
	}
	if order.ResultCode != "SUCCESS" {
		return nil, fmt.Errorf("resutl code:%s, result desc:%s", order.ErrCode, order.ErrCodeDesc)
	}
	return order, nil
}

//查询
func (pay *PayConfig) Query(orderId string) (*QueryResult, error) {
	param := pay.SetQueryParam(orderId)
	signStr := Sign(param, pay.Token)
	param["sign"] = signStr
	xmlParam := MapToXmlString(param)
	order, err := pay.SendQuery(xmlParam)
	if err != nil {
		return nil, err
	}
	resultInMap, err := ToMap(order)
	if err != nil {
		return nil, err
	}
	wantSign := Sign(resultInMap, pay.Token)
	gotSign := resultInMap["sign"]
	if wantSign != gotSign {
		return nil, fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}
	return order, nil
}

//发送订单
func (pay *PayConfig) SendOrder(param string) (*OrderResult, error) {
	url := fmt.Sprintf(Url_Pay_Order)
	result := new(OrderResult)
	b, err := HttpPostXml(url, []byte(param), &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode == "FAIL" {
		return nil, ErrNoData([]byte(b))
	}
	return result, nil
}

//发送查询
func (pay *PayConfig) SendQuery(param string) (*QueryResult, error) {
	url := fmt.Sprintf(Url_Pay_Query)
	result := new(QueryResult)
	b, err := HttpPostXml(url, []byte(param), &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode == "FAIL" {
		return nil, ErrNoData([]byte(b))
	}
	return result, nil
}

//设置订单参数
func (pay *PayConfig) SetOrderParam(orderId, amount, desc, clientIp, tradeType string) map[string]string {
	param := make(map[string]string)
	param["appid"] = pay.AppId
	param["mch_id"] = pay.MchId
	param["nonce_str"] = newRandStr()
	param["body"] = desc
	param["out_trade_no"] = orderId
	param["spbill_create_ip"] = clientIp
	param["notify_url"] = pay.NotifyUrl
	param["total_fee"] = amount
	param["trade_type"] = tradeType
	return param
}

//设置查询参数
func (pay *PayConfig) SetQueryParam(orderId string) map[string]string {
	param := make(map[string]string)
	param["appid"] = pay.AppId
	param["mch_id"] = pay.MchId
	param["out_trade_no "] = orderId
	param["nonce_str"] = newRandStr()
	return param
}