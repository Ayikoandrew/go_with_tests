package reflection

import (
	"reflect"
)

// TODO: This class is not finished yet! studying other classes.

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				Walk(val.Index(i).Interface(), fn)
			}
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}

func WalkV2(x interface{}, fn func(in string)) {
	fn(x.(string))
}
