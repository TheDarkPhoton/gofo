package gofo

type Function struct {
	args []interface{}
	overloads []Overload
}

func CreateFunction(args []interface{}) Function {
	return Function{ args: args }
}

func (f *Function) Run() {
	for _, o := range f.overloads {
		mapped_args := o.args.mapped(f.args)
		if len(mapped_args) != 0 {
			o.callback(mapped_args)
			return
		}
	}
}