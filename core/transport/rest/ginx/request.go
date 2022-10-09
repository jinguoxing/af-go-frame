package ginx

import (
    "github.com/gin-gonic/gin"
    "strconv"
)

// Get body data from context
func GetBodyData(c *gin.Context, ReqBodyKey string) []byte {
    if v, ok := c.Get(ReqBodyKey); ok {
        if b, ok := v.([]byte); ok {
            return b
        }
    }
    return nil
}

// Param returns the value of the URL param
func ParseParamID(c *gin.Context, key string) uint64 {
    val := c.Param(key)
    id, err := strconv.ParseUint(val, 10, 64)
    if err != nil {
        return 0
    }
    return id
}


