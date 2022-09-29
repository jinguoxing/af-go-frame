package codes


var (
    // No error code specified.
    CodeNil                      = newLocalCoder("","","","")
    // It is OK.
    CodeOK                       = newLocalCoder("OK", "OK", "","")
    // An error occurred internally.
    CodeInternalError            = newLocalCoder("Public.InternalError",
        "服务 [serviceName] 内部错误。", "","")

    // Unknown error.
    CodeUnknown                  = newLocalCoder("Public.UnknownError",
        "服务 [serviceName] 内部出现未知错误。", "","")

    CodeInvalidParameter         = newLocalCoder("Public.InvalidParameter", "调用服务 [serviceName] 接口 [interfaceName] 参数值 [params] 校验不通过。", "","请使用请求参数构造规范化的请求字符串。详细信息参见产品 API 文档。")

    // The given parameter for current operation is invalid.
    CodeMissingParameter         = newLocalCoder("Public.MissingParameter", "调用服务 [serviceName] 接口 [interfaceName] 缺少必须参数 [params] 。", "","请检查调用时是否填写了此参数，并重试请求。详细信息参见产品 API 文档。")


    CodeUnsupportedHTTPMethod = newLocalCoder("Public.UnsupportedHTTPMethod", "服务 [serviceName] 未提供 HTTP 接口 [interfaceName] 支持。", "","建议查看产品 API 文档。")
    //CodeUnsupportedHTTPMethod

    CodeServiceUnavailable = newLocalCoder("Public.ServiceUnavailableDuringUpgrade", "服务 [serviceName] 暂时不可用。", "系统正在升级，暂时不可用。","请稍后尝试。")

    CodeNotFound                 = newLocalCoder("Public.NotFound", "Not Found", "","")                  // Resource does not exist.

    CodeNotAuthorized            = newLocalCoder("Public.NotAuthorized", "Not Authorized", "","" )            // Not Authorized.


    //
    //CodeValidationFailed         = localCoder{51, "Validation Failed",nil ,""}           // Data validation failed.
    //CodeDbOperationError         = localCoder{52, "Database Operation Error", nil,""}    // Database operation error.
    //               // Parameter for current operation is missing.
    //CodeInvalidOperation         = localCoder{55, "Invalid Operation", nil,""}           // The function cannot be used like this.
    //CodeInvalidConfiguration     = localCoder{56, "Invalid Configuration", nil,""}       // The configuration is invalid for current operation.
    //CodeMissingConfiguration     = localCoder{57, "Missing Configuration", nil,""}       // The configuration is missing for current operation.
    //CodeNotImplemented           = localCoder{58, "Not Implemented", nil,""}             // The operation is not implemented yet.
    //CodeNotSupported             = localCoder{59, "Not Supported", nil,""}               // The operation is not supported yet.
    //CodeOperationFailed          = localCoder{60, "Operation Failed", nil,""}            // I tried, but I cannot give you what you want.

    //CodeSecurityReason           = localCoder{62, "Security Reason", nil,""}             // Security Reason.
    //CodeServerBusy               = localCoder{63, "Server Is Busy", nil,""}              // Server is busy, please try again later.
    //

    //CodeInvalidRequest           = localCoder{66, "Invalid Request", nil,""}             // Invalid request.
    //CodeBusinessValidationFailed = localCoder{300, "Business Validation Failed", nil,""} // Business validation failed.

)
