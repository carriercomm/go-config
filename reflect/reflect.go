package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Fetch(ns string, source interface{}) (interface{}, error) {
	elem := source
	for _, part := range strings.Split(ns, ".") {
		switch t := elem.(type) {
		case map[string]interface{}:
			if val, ok := t[part]; ok {
				elem = val
			} else {
				err := fmt.Errorf("could not find key matching %s",
					part)
				return nil, err
			}
		case []interface{}:
			if n, err := strconv.ParseInt(part, 10, 0); err == nil {
				elem = t[n]
			} else {
				return nil, err
			}
		default:
			found := reflect.TypeOf(t)
			err := fmt.Errorf("invalid type found, expected object "+
				"or array (found %s)", found.String())
			return nil, err
		}
	}

	return elem, nil
}
