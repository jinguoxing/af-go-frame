package agerrors_test

import (
    "fmt"
    "github.com/jinguoxing/af-go-frame/core/errorx/agcodes"
    "github.com/jinguoxing/af-go-frame/core/errorx/agerrors"
)

func ExampleNewCode() {

    err := agerrors.NewCode(
        //agcodes.New("10000", "", "", "", nil, ""),
        agcodes.CodeInvalidParameter,
        "My Error")

    fmt.Println(err.Error())
    fmt.Println(agerrors.Code(err))
    //Output:
    //My Error
    //Public.InvalidParameter:调用服务 [serviceName] 接口 [interfaceName] 参数值 [params] 校验不通过。

}
