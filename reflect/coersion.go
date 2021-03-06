package reflect

import "fmt"

type toType int

const (
	String toType = iota
	Int
	Bool
)

type typeConverter func(interface{}) (interface{}, error)

var types map[toType]typeConverter = map[toType]typeConverter{
	String: StringConverter,
	Int:    IntConverter,
	Bool:   BoolConverter,
}

func Coerse(v interface{}, t toType) (interface{}, error) {
	if convert, ok := types[t]; ok {
		return convert(v)
	} else {
		return nil, fmt.Errorf("type uncoerciable")
	}
}
