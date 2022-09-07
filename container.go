package lady

import (
	"reflect"

	"github.com/samber/lo"
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

func Call1[T any, F func(T)](c *Container, f F, others ...any) {
	v := reflect.ValueOf(f)

	args := []reflect.Value{}

	if arg, ok := c.data[reflect.PointerTo(v.Type().In(0))]; ok {
		args = append(args, arg)
	} else {
		args = append(args, reflect.ValueOf(others[0]))
		others = others[1:]
	}

	args = append(
		args,
		lo.Map(
			others,
			func(value interface{}, _ int) reflect.Value { return reflect.ValueOf(value) },
		)...,
	)

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

func Call3[T1 any, T2 any, T3 any, F func(T1, T2, T3)](c *Container, f F, others ...any) {
	v := reflect.ValueOf(f)

	args := []reflect.Value{}

	for i := 0; i < 3; i++ {
		if arg, ok := c.data[reflect.PointerTo(v.Type().In(i))]; ok {
			args = append(args, arg)
		} else {
			args = append(args, reflect.ValueOf(others[0]))
			others = others[1:]
		}
	}

	v.Call(args)
}
