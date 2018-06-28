package wechat

import (
	"sort"
	"strings"
	"fmt"
	"crypto/md5"
	"reflect"
	"encoding/xml"
)

func Sign(param map[string]string, token string) string {
	newMap := make(map[string]string)
	for k, v := range param {
		if k == "sign" {
			continue
		}
		if v == "" {
			continue
		}
		newMap[k] = v
	}
	signStr := SortAndConcat(newMap)
	signWithKey := signStr + "&key=" + token
	return fmt.Sprintf("%X", md5.Sum([]byte(signWithKey)))
}

func SortAndConcat(param map[string]string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	var sortedParam []string
	sort.Strings(keys)
	for _, k := range keys {
		sortedParam = append(sortedParam, k+"="+param[k])
	}
	return strings.Join(sortedParam, "&")
}

func MapToXmlString(param map[string]string) string {
	x := "<xml>"
	for k, v := range param {
		x = x + fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	x = x + "</xml>"

	return x
}

func ToMap(in interface{}) (map[string]string, error) {
	out := make(map[string]string)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagv := fi.Tag.Get("xml"); tagv != "" && tagv != "xml" {
			out[tagv] = v.Field(i).String()
		}
	}
	return out, nil
}

func ParseOrderResult(resp []byte) (*OrderResult, error) {
	result := new(OrderResult)
	err := xml.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
