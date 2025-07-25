package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

func WalkV2(x interface{}, fn func(in string)) {
	fn(x.(string))
}
