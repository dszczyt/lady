package lady_test

import (
	"testing"

	"github.com/dszczyt/lady"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DemoIface interface {
	Demo()
}

type DemoIface2 interface {
	Demo()
}

type DemoStruct struct {
	mock.Mock
}

func (d *DemoStruct) Demo() {
	d.Called()
}

type DemoStruct2 struct {
	mock.Mock
}

func (d *DemoStruct2) Demo() {
	d.Called()
}

func TestDI(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		container := lady.New()

		obj := DemoStruct{}

		obj.On("Demo").Return()

		container.Bind(new(DemoIface), &obj)

		test := func(d DemoIface) {
			d.Demo()
		}

		lady.Call1(container, test)

		obj.AssertExpectations(t)
	})

	t.Run("simple with arg", func(t *testing.T) {
		assert := tassert.New(t)

		container := lady.New()

		obj := DemoStruct{}

		obj.On("Demo").Return()

		container.Bind(new(DemoIface), &obj)
		container.Bind(new(DemoIface), &obj)

		test := func(d DemoIface, s string) {
			assert.Equal("good", s)
			d.Demo()
		}

		lady.Call2(container, test, "good")

		obj.AssertExpectations(t)
	})

	t.Run("simple with arg inversed", func(t *testing.T) {
		assert := tassert.New(t)

		container := lady.New()

		obj := DemoStruct{}

		obj.On("Demo").Return()

		container.Bind(new(DemoIface), &obj)

		test := func(s string, d DemoIface) {
			assert.Equal("good", s)
			d.Demo()
		}

		lady.Call2(container, test, "good")

		obj.AssertExpectations(t)
	})

	t.Run("2 injections", func(t *testing.T) {
		container := lady.New()

		obj := DemoStruct{}
		obj2 := DemoStruct2{}

		obj.On("Demo").Return()
		obj2.On("Demo").Return()

		container.Bind(new(DemoIface), &obj)
		container.Bind(new(DemoIface2), &obj)

		test := func(d DemoIface, d2 DemoIface2) {
			d.Demo()
			d2.Demo()
		}

		lady.Call2(container, test)

		obj.AssertExpectations(t)
	})
}
