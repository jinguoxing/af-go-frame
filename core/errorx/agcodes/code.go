package agcodes

type Coder interface {
    GetErrorCode() string
    GetDescription() string
    GetCause() string
    GetSolution() string
    GetErrorDetails() interface{}
    GetErrorLink() string
}

// New creates and returns an error code.
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

// WithCode creates and returns a new error code based on given Code.
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
