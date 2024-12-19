package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsValue)

	return []byte(quotedJSONValue), nil
}
