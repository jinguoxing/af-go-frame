package codes


type localCoder struct {

    // Error code, usually an integer.
    ErrorCode string `json:"errorcode"`
    // Brief message for this error code.
    Description string `json:"description"`

    // error cause
    Cause string  `json:"cause"`

    // error solution
    Solution string

    // As type of interface, it is mainly designed as an extension field for error code.
    ErrorDetails interface{}

    // Ref specify the reference document.
    ErrorLink string

}


func newLocalCoder(code,description,cause,solution string) localCoder {

    return localCoder{
        ErrorCode:code,
        Description:description,
        Cause:cause,
        Solution:solution,
    }
}


func(c localCoder) GetErrorCode() string {

    return c.ErrorCode
}



func(c localCoder)GetDescription() string{

    return c.Description
}

func(c localCoder)GetCause() string{

    return c.Cause
}

func(c localCoder)GetSolution() string{

    return c.Solution
}


func(c localCoder) GetErrorDetails() interface{}{

    return c.ErrorDetails
}


func(c localCoder)GetErrorLink() string{

    return c.ErrorLink
}






