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

func Call[T any, F func(T)](c *Container, f F) {
	v := reflect.ValueOf(f)

	args := []reflect.Value{c.data[reflect.PointerTo(v.Type().In(0))]}

	v.Call(args)
}
