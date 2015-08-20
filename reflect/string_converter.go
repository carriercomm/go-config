package reflect

import "fmt"

func StringConverter(v interface{}) (interface{}, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	default:
		return nil, fmt.Errorf("unable to coerce %v into type string", v)
	}
}
