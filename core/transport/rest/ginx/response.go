package ginx

import (
    "github.com/gin-gonic/gin"
    "github.com/jinguoxing/af-go-frame/core/errorx/agcodes"
    "github.com/jinguoxing/af-go-frame/core/errorx/agerrors"
    "net/http"
)

type HttpError struct {
    Code        string      `json:"code"`
    Description string      `json:"description"`
    Solution    string      `json:"solution"`
    Cause       string      `json:"cause"`
    Detail      interface{} `json:"detail,omitempty"`
    Data        interface{} `json:"data"`
}

// success Json Response
func ResOKJson(c *gin.Context, data interface{}) {

    if data == nil {
        data = gin.H{}
    }
    c.JSON(http.StatusOK, data)
}

// list Response
func ResList(c *gin.Context, list interface{}, totalCount int) {

    c.JSON(http.StatusOK, gin.H{
        "entries":     list,
        "total_count": totalCount,
    })

}

// failed Json Response
func ResErrJson(c *gin.Context, err error) {

    var (
        code = agerrors.Code(err)
    )
    if err != nil {
        if code == agcodes.CodeNil {
            code = agcodes.CodeInternalError
        }
    } else if c.Writer.Status() > 0 && c.Writer.Status() != http.StatusOK {
        switch c.Writer.Status() {
        case http.StatusNotFound:
            code = agcodes.CodeNotFound
        case http.StatusForbidden:
            code = agcodes.CodeNotAuthorized
        default:
            code = agcodes.CodeInternalError
        }
    } else {
        code = agcodes.CodeOK
    }

    c.JSON(c.Writer.Status(), HttpError{
        Code:        code.GetErrorCode(),
        Description: code.GetDescription(),
        Solution:    code.GetSolution(),
        Cause:       code.GetCause(),
    })
}
