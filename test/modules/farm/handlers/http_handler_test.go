package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func DoSomething(number int) int {
	return number + 1
}

type HandlerTestObject struct {
	mock.Mock
}

func (o *HandlerTestObject) DoSomething(number int) int {
	args := o.Called(number)
	return args.Int(number + 1)
}

func TestSomething(t *testing.T) {
	// create an instance of our test object
	testObj := new(HandlerTestObject)

	// set up expectations
	testObj.On("DoSomething", 5).Return(6)

	// assert that the expectations were met
	ok := DoSomething(5)

	assert.Equal(ok, 6, 6)
}
