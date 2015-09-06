package reflect

import (
	"fmt"
	"strconv"
)

func IntConverter(v interface{}) (interface{}, error) {
	switch t := v.(type) {
	case int:
		return t, nil
	case string:
		return strconv.Atoi(t)
	default:
		return nil, fmt.Errorf("could not convert \"%v\" into type int", v)
	}
}
