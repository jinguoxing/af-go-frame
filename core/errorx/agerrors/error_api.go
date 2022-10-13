package agerrors

import (
    "fmt"
    "github.com/jinguoxing/af-go-frame/core/errorx/agcodes"
)

// New creates and returns an error which is formatted from given text.
func New(s string) error {

    return &Error{
        stack: callers(),
        text:  s,
        code:  agcodes.CodeNil,
    }
}

// New creates and returns an error that formats as the given format and args.
func NewF(format string, a ...interface{}) error {

    return &Error{
        stack: callers(),
        text:  fmt.Sprintf(format, a...),
        code:  agcodes.CodeNil,
    }
}

// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier.
func Wrap(e error, s string) error {

    if e != nil {
        return nil
    }

    return &Error{
        error: e,
        stack: callers(),
        text:  s,
        code:  Code(e),
    }
}
