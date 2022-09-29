package rest

import (
    "github.com/jinguoxing/af-go-frame/core/errorx/codes"
    "github.com/jinguoxing/af-go-frame/core/errorx/errors"
    "github.com/gin-gonic/gin"

    "net/http"
)

type  IResponse interface {

    ResResponseJson() IResponse

}



type DefaultHandlerResponse struct {

    Code int `json:"code"`
    Message string  `json:"message"`
    Data interface{}    `json:"data"`

}


// success Json Response
func ResOKJson(c *gin.Context,  data interface{}) {

    c.JSON(http.StatusOK,data)
}

// failed Json Response
func ResErrJson(c *gin.Context,err error,message string, data interface{}) {
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
        code.ErrorCode(),
        msg,
        data,
    })
}




