package errors

import (
    "github.com/jinguoxing/af-go-frame/core/errorx/codes"
    "fmt"
    "strings"
)


// NewCode creates and returns an error that has error code and given text.
func NewCode(code codes.Coder, s ...string) error {

    return &Error{
        stack: callers(),
        text:  strings.Join(s, ", "),
        code:  code,
    }
}


func NewCodef(code codes.Coder,format string,a ...interface{})error{

    return &Error{

        stack: callers(),
        text:  fmt.Sprintf(format,a...),
        code:  code,
    }
}


func Code(err error) codes.Coder {

    if err == nil {
        return codes.CodeNil
    }

    if e, ok := err.(ICode); ok {
        return e.Code()
    }

    if e,ok:= err.(IUnwrap);ok{
        return Code(e.Unwrap())
    }

    return codes.CodeNil
}


