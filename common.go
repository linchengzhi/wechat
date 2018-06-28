package wechat

import (
	"fmt"
	"errors"
)

func ErrNoData(b []byte) error {
	tips := fmt.Sprintf("no data result=%s", string(b))
	return errors.New(tips)
}