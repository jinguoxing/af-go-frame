package httpx

import (
    "af-go-frame/core/container"
    "af-go-frame/core/contract"
    "af-go-frame/core/gin"
)

type AFKernelProvider struct {
    HttpEngine *gin.Engine
}

// Register 注册服务提供者
func (provider *AFKernelProvider) Register(c container.Container) container.NewInstance {
    return NewAFHttpService
}

// Boot 启动的时候判断是否由外界注入了Engine，如果注入的化，用注入的，如果没有，重新实例化
func (provider *AFKernelProvider) Boot(c container.Container) error {
    if provider.HttpEngine == nil {
        provider.HttpEngine = gin.Default()
    }
    provider.HttpEngine.SetContainer(c)
    return nil
}

// IsDefer 引擎的初始化我们希望开始就进行初始化
func (provider *AFKernelProvider) IsDefer() bool {
    return false
}

// Params 参数就是一个HttpEngine
func (provider *AFKernelProvider) Params(c container.Container) []interface{} {
    return []interface{}{provider.HttpEngine}
}

// Name 提供凭证
func (provider *AFKernelProvider) Name() string {
    return contract.HttpEngineKey
}

