package gin

import (
    "af-go-frame/core/container"
)

func(engine *Engine) SetContainer(container container.Container){

    engine.container = container
}



// GetContainer 从Engine中获取container
func (engine *Engine) GetContainer() container.Container {
    return engine.container
}

// engine实现container的绑定封装
func (engine *Engine) Bind(provider container.ServiceProvider) error {
    return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
    return engine.container.IsBind(key)
}

