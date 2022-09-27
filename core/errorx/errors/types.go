package errors

import "af-go-frame/core/errorx/codes"


type ICode interface {

    Error() string
    Code() codes.Coder
}


// IUnwrap is the interface for Unwrap feature.
type IUnwrap interface {
    Error() string
    Unwrap() error
}

