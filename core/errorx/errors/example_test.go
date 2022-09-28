package errors_test

import (
    "github.com/jinguoxing/af-go-frame/core/errorx/codes"
    "github.com/jinguoxing/af-go-frame/core/errorx/errors"
    "fmt"
)

func ExampleNewCode(){

    err :=  errors.NewCode(codes.New(1000,"",nil,""),"My Error")

    fmt.Println(err.Error())
    fmt.Println(errors.Code(err))
}
