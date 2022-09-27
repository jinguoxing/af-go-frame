package errors_test

import (
    "af-go-frame/core/errorx/codes"
    "af-go-frame/core/errorx/errors"
    "fmt"
)

func ExampleNewCode(){

    err :=  errors.NewCode(codes.New(1000,"",nil,""),"My Error")

    fmt.Println(err.Error())
    fmt.Println(errors.Code(err))
}
