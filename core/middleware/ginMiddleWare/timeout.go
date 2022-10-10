package ginMiddleWare

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

func TimeOut(d time.Duration) gin.HandlerFunc {

    return func(c *gin.Context) {

        finish := make(chan struct{}, 1)
        panicChan := make(chan interface{}, 1)
        // 执行业务逻辑前预操作：初始化超时context
        durationCtx, cancel := context.WithTimeout(c.Request.Context(), d)
        defer cancel()

        go func() {
            defer func() {
                if p := recover(); p != nil {
                    panicChan <- p
                }
            }()
            // 使用next执行具体的业务逻辑
            c.Next()

            finish <- struct{}{}
        }()

        select{

        case p := <-panicChan:
            c.JSON(http.StatusInternalServerError,gin.H{})
            log.Println(p)
        case <-finish:
            fmt.Println("finish")
         case <-durationCtx.Done():
             c.JSON(http.StatusBadGateway,gin.H{})
        }
    }
}
