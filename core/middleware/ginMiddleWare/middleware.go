package ginMiddleWare

import "github.com/gin-gonic/gin"

var GinMiddlewares = defaultGinMiddlewares()

func defaultGinMiddlewares() map[string]gin.HandlerFunc {

    return map[string]gin.HandlerFunc{

        "recovery": gin.Recovery(),
    }
}
