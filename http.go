package wechat

import (
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
)

func HttpGetJson(url string, v interface{}) ([]byte, error) {
	resp, err := Get(url).Response()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b,_ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(b, v); err != nil {
		return b, err
	}
	return b, nil
}

func HttpPostJson(url string, param []byte, v interface{}) ([]byte, error) {
	req := Post(url)
	req.Body(param)
	resp, err := req.Response()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b,_ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(b, v); err != nil {
		return b, err
	}
	return b, nil
}

func HttpPostXml(url string, param []byte, v interface{}) ([]byte, error) {
	req := Post(url)
	req.Body(param)
	resp, err := req.Response()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b,_ := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(b, v); err != nil {
		return b, err
	}
	return b, nil
}