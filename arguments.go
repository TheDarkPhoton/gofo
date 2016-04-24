package gofo

import "reflect"

type Arg struct {
	Name string
	Kind reflect.Kind
}

type Arguments struct {
	pairs []Arg
}

func CreateArguments(pairs ...Arg) Arguments {
	return Arguments{ pairs }
}

func (a *Arguments) mapped(args []interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	if len(args) != len(a.pairs) { return result }
	for i, arg := range args {
		if reflect.TypeOf(arg).Kind() != a.pairs[i].Kind {
			return make(map[string]interface{})
		} else {
			result[a.pairs[i].Name] = arg
		}
	}

	return result
}