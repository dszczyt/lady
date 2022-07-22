package lady

import (
	"reflect"
)

type Container struct {
	data map[reflect.Type]reflect.Value
}

func New() *Container {
	return &Container{
		data: map[reflect.Type]reflect.Value{},
	}
}

func (c *Container) Bind(src interface{}, dst interface{}) {
	c.data[reflect.TypeOf(src)] = reflect.ValueOf(dst)
}

func Call1[T any, F func(T)](c *Container, f F) {
	v := reflect.ValueOf(f)

	args := []reflect.Value{c.data[reflect.PointerTo(v.Type().In(0))]}

	v.Call(args)
}

func Call2[T any, U any, F func(T, U)](c *Container, f F, others ...any) {
	v := reflect.ValueOf(f)

	args := []reflect.Value{}

	for i := 0; i < 2; i++ {
		if arg, ok := c.data[reflect.PointerTo(v.Type().In(i))]; ok {
			args = append(args, arg)
		} else {
			args = append(args, reflect.ValueOf(others[0]))
			others = others[1:]
		}
	}

	v.Call(args)
}
