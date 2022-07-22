package lady_test

import (
	"fmt"
	"testing"

	"github.com/dszczyt/lady"
	tassert "github.com/stretchr/testify/assert"
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

		test := func(d DemoIface, s string) {
			assert.Equal("good", s)
			d.Demo()
		}

		lady.Call2(container, test, "good")

		obj.AssertExpectations(t)
	})
}
