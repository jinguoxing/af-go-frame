package rest

import (
    "af-go-frame/core/errorx/codes"
    "af-go-frame/core/errorx/errors"
    "github.com/gin-gonic/gin"

    "net/http"
)

type  IResponse interface {

    ResResponseJson() IResponse

}



type DefaultHandlerResponse struct {

    Code codes.Coder `json:"code"`
    Message string  `json:"message"`
    Data interface{}    `json:"data"`

}


// success Json Response
func ResponseOkJson(message string, data interface{}, c *gin.Context) {

    var msg string
    code := codes.CodeOK

    if message == "" {
        msg = "success"
    } else {
        msg = message
    }

    c.JSON(http.StatusOK, DefaultHandlerResponse{
        code,
        msg,
        data,
    })
}

// failed Json Response
func ResponseJson(err error,message string, data interface{}, c *gin.Context) {
    var (
        msg  string
        code = errors.Code(err)
    )
    if err != nil {
        if code == codes.CodeNil {
            code = codes.CodeInternalError
        }
        msg = err.Error()
    } else if c.Writer.Status() > 0 && c.Writer.Status() > 0 {
        msg = http.StatusText(c.Writer.Status())
        switch c.Writer.Status() {
        case http.StatusNotFound:
            code = codes.CodeNotFound
        case http.StatusForbidden:
            code = codes.CodeNotAuthorized
        default:
            code = codes.CodeUnknown
        }
    } else {
        code = codes.CodeOK
    }

    c.JSON(http.StatusOK, DefaultHandlerResponse{
        code,
        msg,
        data,
    })
}




