package rest

import (
    "github.com/gin-gonic/gin"
    "github.com/jinguoxing/af-go-frame/core/errorx/codes"
    "github.com/jinguoxing/af-go-frame/core/errorx/errors"

    "net/http"
)


type HttpError struct {

    Code string `json:"code"`
    Description string  `json:"description"`
    Solution string  `json:"solution"`
    Cause string  `json:"cause"`
    Detail interface{}  `json:"detail,omitempty"`
    Data interface{}    `json:"data"`
}





// success Json Response
func ResOKJson(c *gin.Context, data interface{}) {

    c.JSON(http.StatusOK, data)
}

// failed Json Response
func ResErrJson(c *gin.Context, err error) {
    var (

        code = errors.Code(err)
    )
    if err != nil {
        if code == codes.CodeNil {
            code = codes.CodeInternalError
        }
    } else if c.Writer.Status() > 0 && c.Writer.Status() != http.StatusOK {
        switch c.Writer.Status() {
        case http.StatusNotFound:
            code = codes.CodeNotFound
        case http.StatusForbidden:
            code = codes.CodeNotAuthorized
        default:
            code = codes.CodeInternalError
        }
    } else {
        code = codes.CodeOK
    }

    c.JSON(http.StatusOK, HttpError{
        Code:        code.GetErrorCode(),
        Description: code.GetDescription(),
        Solution:    code.GetSolution(),
        Cause:       code.GetCause(),
    })
}




