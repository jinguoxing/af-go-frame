package httpx

import (
    "af-go-frame/core/gin"
    "net/http"
)

type AFHttpService struct {

    engine *gin.Engine
}

// 初始化web引擎服务实例
func NewAFHttpService(params ...interface{}) (interface{}, error) {
    httpEngine := params[0].(*gin.Engine)
    return &AFHttpService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *AFHttpService) HttpEngine() http.Handler {
    return s.engine
}
