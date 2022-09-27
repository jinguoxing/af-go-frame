package codes


type  Coder interface {

    ErrorCode() int
   // StatusCode() int
    Message() string
    Reference() string
    Detail() interface{}
}


var (
    // No error code specified.
    CodeNil                      = localCoder{-1, "", nil,""}
    // It is OK.
    CodeOK                       = localCoder{0, "OK", nil,""}
    CodeInternalError            = localCoder{50, "Internal Error", nil,""}              // An error occurred internally.
    CodeValidationFailed         = localCoder{51, "Validation Failed",nil ,""}           // Data validation failed.
    CodeDbOperationError         = localCoder{52, "Database Operation Error", nil,""}    // Database operation error.
    CodeInvalidParameter         = localCoder{53, "Invalid Parameter", nil,""}           // The given parameter for current operation is invalid.
    CodeMissingParameter         = localCoder{54, "Missing Parameter", nil,""}           // Parameter for current operation is missing.
    CodeInvalidOperation         = localCoder{55, "Invalid Operation", nil,""}           // The function cannot be used like this.
    CodeInvalidConfiguration     = localCoder{56, "Invalid Configuration", nil,""}       // The configuration is invalid for current operation.
    CodeMissingConfiguration     = localCoder{57, "Missing Configuration", nil,""}       // The configuration is missing for current operation.
    CodeNotImplemented           = localCoder{58, "Not Implemented", nil,""}             // The operation is not implemented yet.
    CodeNotSupported             = localCoder{59, "Not Supported", nil,""}               // The operation is not supported yet.
    CodeOperationFailed          = localCoder{60, "Operation Failed", nil,""}            // I tried, but I cannot give you what you want.
    CodeNotAuthorized            = localCoder{61, "Not Authorized", nil,""}              // Not Authorized.
    CodeSecurityReason           = localCoder{62, "Security Reason", nil,""}             // Security Reason.
    CodeServerBusy               = localCoder{63, "Server Is Busy", nil,""}              // Server is busy, please try again later.
    CodeUnknown                  = localCoder{64, "Unknown Error", nil,""}               // Unknown error.
    CodeNotFound                 = localCoder{65, "Not Found", nil,""}                   // Resource does not exist.
    CodeInvalidRequest           = localCoder{66, "Invalid Request", nil,""}             // Invalid request.
    CodeBusinessValidationFailed = localCoder{300, "Business Validation Failed", nil,""} // Business validation failed.


)

func New(code int, message string, detail interface{}, ref string, ) Coder {

    return localCoder{
        errorCode:    code,
        errorMessage: message,
        errorDetail:  detail,
        errorRef:     ref,
    }
}


func WithCode(code Coder, detail interface{}) Coder {

    return localCoder{
        errorCode:    code.ErrorCode(),
        errorMessage: code.Message(),
        errorDetail:  detail,
        errorRef:     code.Reference(),
    }
}



func Register() {




}
















