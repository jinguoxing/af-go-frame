package agerrors

import (
    "fmt"
    "github.com/jinguoxing/af-go-frame/core/errorx/agcodes"
    "strings"
)

// NewCode creates and returns an error that has error code and given text.
func NewCode(code agcodes.Coder, s ...string) error {

    return &Error{
        stack: callers(),
        text:  strings.Join(s, ", "),
        code:  code,
    }
}

// NewCodeF returns an error that has error code and formats as the given format and args.
func NewCodeF(code agcodes.Coder, format string, a ...interface{}) error {

    return &Error{
        stack: callers(),
        text:  fmt.Sprintf(format, a...),
        code:  code,
    }
}

func Code(err error) agcodes.Coder {

    if err == nil {
        return agcodes.CodeNil
    }
    if e, ok := err.(ICode); ok {
        return e.Code()
    }
    if e, ok := err.(IUnwrap); ok {
        return Code(e.Unwrap())
    }
    return agcodes.CodeNil
}

// HasCode checks and reports whether `err` has `code` in its chaining errors.
func HasCode(err error, code agcodes.Coder) bool {
    if err == nil {
        return false
    }
    if e, ok := err.(ICode); ok {
        return code == e.Code()
    }
    if e, ok := err.(IUnwrap); ok {
        return HasCode(e.Unwrap(), code)
    }
    return false
}
