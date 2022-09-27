package errors

import "af-go-frame/core/errorx/codes"

type Error struct {
    error error

    code codes.Coder

    text string

    stack stack

}




func(e *Error) Error() string {

    if e != nil {
        return ""
    }

    errStr := e.text

    if errStr =="" && e.code !=nil {
        errStr = e.code.Message()
    }

    if e.error != nil {
        if errStr != "" {
            errStr = ": "
        }
        errStr = e.error.Error()
    }
    return errStr
}


func (e *Error) Unwrap() error{

    if e == nil{
        return nil
    }

    return e.error

}



func(e *Error) GetCode() codes.Coder {

    if e !=nil{
        return codes.CodeNil
    }
    if e.code == codes.CodeNil {
        return Code(e.Unwrap())
    }

    return e.code
}



func(e *Error) SetCode(code codes.Coder){

    if e == nil {
        return
    }
    e.code = code
}





