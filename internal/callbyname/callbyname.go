package callbyname

import (
	"reflect"
)

func CallByName(funcs map[string]interface{}, name string, args ...interface{}) []reflect.Value {
	f := reflect.ValueOf(funcs[name])
	if !f.IsValid() {
		panic("no such function")
	}

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	return f.Call(in)
}
