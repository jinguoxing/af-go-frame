package errors

import (
    "github.com/jinguoxing/af-go-frame/core/errorx/codes"
    "fmt"
)


func New(s string) error {

    return &Error{
        stack: callers(),
        text:  s,
        code:  codes.CodeNil,
    }
}

func Newf(format string, a ...interface{}) error {

    return &Error{
        stack: callers(),
        text:  fmt.Sprintf(format, a...),
        code:  codes.CodeNil,
    }
}





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

