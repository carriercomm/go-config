package reflect

import (
	"fmt"
	"strings"
)

func BoolConverter(v interface{}) (interface{}, error) {
	switch t := v.(type) {
	case int:
		return t == 1, nil
	case string:
		truthy := []string{
			"true", "yes", "ok", "Y", "on",
		}

		for _, test := range truthy {
			if strings.ToLower(t) == test {
				return true, nil
			}
		}

		return false, fmt.Errorf("unable to understand value of \"%s\", expected one of %v", t, truthy)
	default:
		return nil, fmt.Errorf("could not convert \"%v\" into type int", v)
	}
}
