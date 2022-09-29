package codes


type localCoder struct {

    // Error code, usually an integer.
    errorCode int `json:"errorcode"`

    // RPC/Http status that should be used for the associated error code.
   // StatusCode int

    // Brief message for this error code.
    errorMessage string

    // As type of interface, it is mainly designed as an extension field for error code.
    errorDetail interface{}

    // Ref specify the reference document.
    errorRef string

}



func(c localCoder) ErrorCode() int {

    return c.errorCode
}


//func(c localCoder) Status() int{
//
//    return c.StatusCode
//}

func(c localCoder)Message() string{

    return c.errorMessage
}


func (c localCoder) Detail() interface{} {
    return c.errorDetail
}


func (c localCoder) Reference() string {
    return c.errorRef
}







