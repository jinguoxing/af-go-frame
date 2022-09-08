package gin

import (
    "af-go-frame/core/contract"
)

// MustMakeApp 从容器中获取App服务
func (c *Context) MustMakeApp() contract.AppInfo {
    return c.MustMake(contract.AppKey).(contract.AppInfo)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Context) MustMakeKernel() contract.HttpEngine {
    return c.MustMake(contract.HttpEngineKey).(contract.HttpEngine)
}

// MustMakeConfig 从容器中获取配置服务
func (c *Context) MustMakeConfig() contract.Config {
    return c.MustMake(contract.ConfigKey).(contract.Config)
}

// MustMakeLog 从容器中获取日志服务
func (c *Context) MustMakeLog() contract.Log {
    return c.MustMake(contract.LogKey).(contract.Log)
}