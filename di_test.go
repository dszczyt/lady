package lady_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/dszczyt/lady"
	"github.com/stretchr/testify/mock"
)

type DemoIface interface {
	Demo()
}

type DemoStruct struct {
	mock.Mock
}

func (d *DemoStruct) Demo() {
	fmt.Printf("CALLED FUNC %+v\n", d)
	d.Called()
}

func TestDI(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		// assert := tassert.New(t)

		container := lady.New()

		obj := DemoStruct{}

		obj.On("Demo").Return()

		container.Bind(new(DemoIface), &obj)

		test := func(d DemoIface) {
			spew.Dump("CALLING", d)
			d.Demo()
		}

		lady.Call(container, test)

		obj.AssertExpectations(t)
	})
}
