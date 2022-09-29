package codes


type  Coder interface {

    GetErrorCode() string
    GetDescription() string
    GetCause() string
    GetSolution() string
    GetErrorDetails() interface{}
    GetErrorLink() string
}


func New(errorCode, description, cause, solution string, detail interface{}, errLink string) Coder {

    return localCoder{
        ErrorCode:    errorCode,
        Description:  description,
        Cause:        cause,
        Solution:     solution,
        ErrorDetails: detail,
        ErrorLink:    errLink,
    }
}


func WithCode(code Coder, detail interface{}) Coder {

    return localCoder{
        ErrorCode:    code.GetErrorCode(),
        Description:  code.GetDescription(),
        Cause:        code.GetCause(),
        Solution:     code.GetSolution(),
        ErrorDetails: detail,
        ErrorLink:    code.GetErrorLink(),
    }
}



func Register() {




}
















