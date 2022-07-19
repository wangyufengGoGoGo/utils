package errorx

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

var (
	ErrPaddingSize = errors.New("padding size error please check the secret key or iv")
)

type ErrOptions struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewError() *ErrOptions {
	return &ErrOptions{}
}

func (e *ErrOptions) WithCode(code int) *ErrOptions {
	e.Code = code
	return e
}

func (e *ErrOptions) WithErr(err string) *ErrOptions {
	e.Msg = err
	return e
}

func (e *ErrOptions) WithData(data any) *ErrOptions {
	e.Data = data
	return e
}

func NewErrorStrWithOptions(code int, err string) *ErrOptions {
	return &ErrOptions{
		Code: code,
		Msg:  err,
	}
}

func NewErrorWithOptions(code int, err error) *ErrOptions {
	return &ErrOptions{
		Code: code,
		Msg:  err.Error(),
	}
}

func NewErrorBadRequest(err error) *ErrOptions {
	return &ErrOptions{
		Code: http.StatusBadRequest,
		Msg:  err.Error(),
	}
}

func NewErrorUnauthorized(err error) *ErrOptions {
	return &ErrOptions{
		Code: http.StatusUnauthorized,
		Msg:  err.Error(),
	}
}

func ExitError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
