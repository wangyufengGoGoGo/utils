package errorx

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	err := fmt.Errorf("%s", "error test 400")
	newError := NewError()
	newError.WithCode(400).WithErr(err.Error())
	fmt.Println(newError)
}

func TestNewErrorWithOptions(t *testing.T) {
	err := fmt.Errorf("%s", "error test 400")
	errorWithOptions := NewErrorWithOptions(400, err)
	fmt.Println(errorWithOptions)
}

func TestNewErrorBadRequest(t *testing.T) {
	err := fmt.Errorf("%s", "error test 400")
	badRequest := NewErrorBadRequest(err)
	fmt.Println(badRequest)
}

func TestNewErrorStrWithOptions(t *testing.T) {
	err := fmt.Errorf("%s", "error test 400")
	badRequest := NewErrorStrWithOptions(400, err.Error())
	fmt.Println(badRequest)
}

func TestNewErrorUnauthorized(t *testing.T) {
	err := fmt.Errorf("%s", "error test 401")
	unauthorized := NewErrorUnauthorized(err)
	fmt.Println(unauthorized)
}

func TestExitError(t *testing.T) {
	err := fmt.Errorf("%s", "error test 502")
	ExitError(err)
}
