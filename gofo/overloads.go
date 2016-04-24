package gofo

type Overload struct {
	args Arguments
	callback func(map[string]interface{})
}

func (f *Function) Overload(args Arguments, callback func(map[string]interface{})){
	f.overloads = append(f.overloads, Overload{ args, callback })
}