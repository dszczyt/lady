package lady

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
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
	v := reflect.MakeFunc(reflect.TypeOf(f), func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{}
	})

	args := []reflect.Value{c.data[reflect.PointerTo(v.Type().In(0))]}

	spew.Dump(v, args)
	res := v.Call(args)
	spew.Dump(res)
}
