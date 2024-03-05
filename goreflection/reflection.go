package goreflection

import "reflect"

//Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming.

func walk(x interface{}, fn func(input string)) {

	// get the value from getValue
	val := getValue(x)
	numOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Slice, reflect.Array:
		numOfValues = val.Len()
		getField = val.Index
		for i := 0; i < numOfValues; i++ {
			walk(getField(i).Interface(), fn)
		}

	case reflect.Struct:
		numOfValues = val.NumField()
		getField = val.Field
		for i := 0; i < numOfValues; i++ {
			walk(getField(i).Interface(), fn)
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}

	case reflect.Func:
		fnResp := val.Call(nil)
		for _, value := range fnResp {
			walk(value.Interface(), fn)
		}
	}

	// reafactored the below code above
	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// iterate for the struct properties
	// for i := 0; i < val.NumField(); i++ {
	// 	ip := val.Field(i)
	// 	switch ip.Kind() {
	// 	case reflect.String:
	// 		fn(ip.String())
	// 	case reflect.Struct:
	// 		walk(ip.Interface(), fn)
	// 	}
	// }
}

// Refactored with switch case
// if ip.Kind() == reflect.String {
// 	fn(ip.String())
// }

// if ip.Kind() == reflect.Struct {
// 	walk(ip.Interface(), fn)
// }

// refactor the code move reflect pointer scenario below
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
