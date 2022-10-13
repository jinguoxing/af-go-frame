package agcodes

import "fmt"

type localCoder struct {

    // Error code, usually an integer.
    ErrorCode string `json:"errorcode"`
    // Brief message for this error code.
    Description string `json:"description"`

    // error cause
    Cause string `json:"cause"`

    // error solution
    Solution string `json:"solution"`

    // As type of interface, it is mainly designed as an extension field for error code.
    ErrorDetails interface{} `json:"error_details"`

    // Ref specify the reference document.
    ErrorLink string `json:"error_link"`
}

func newLocalCoder(code, description, cause, solution string) localCoder {

    return localCoder{
        ErrorCode:   code,
        Description: description,
        Cause:       cause,
        Solution:    solution,
    }
}

func (c localCoder) GetErrorCode() string {

    return c.ErrorCode
}

func (c localCoder) GetDescription() string {

    return c.Description
}

func (c localCoder) GetCause() string {

    return c.Cause
}

func (c localCoder) GetSolution() string {

    return c.Solution
}

func (c localCoder) GetErrorDetails() interface{} {

    return c.ErrorDetails
}

func (c localCoder) GetErrorLink() string {

    return c.ErrorLink
}

// String returns current error code as a string.
func (c localCoder) String() string {
    if c.Description != "" {
        return fmt.Sprintf(`%s:%s`, c.ErrorCode, c.Description)
    }
    if c.ErrorDetails != nil {
        return fmt.Sprintf(`%s:%s %v`, c.ErrorCode, c.Description, c.ErrorDetails)
    }

    return fmt.Sprintf(`%s`, c.ErrorCode)
}
