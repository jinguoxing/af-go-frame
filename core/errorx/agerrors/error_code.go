// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package agerrors

import (
    "github.com/jinguoxing/af-go-frame/core/errorx/agcodes"
)

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() agcodes.Coder {
    if err == nil {
        return agcodes.CodeNil
    }
    if err.code == agcodes.CodeNil {
        return Code(err.Unwrap())
    }
    return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code agcodes.Coder) {
    if err == nil {
        return
    }
    err.code = code
}
