package container

import (
    "fmt"
    "sync"

    "github.com/pkg/errors"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
    // Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
    Bind(provider ServiceProvider) error
    // IsBind 关键字凭证是否已经绑定服务提供者
    IsBind(key string) bool

    // Make 根据关键字凭证获取一个服务，
    Make(key string) (interface{}, error)
    // MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
    // 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
    MustMake(key string) interface{}
    // MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
    // 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
    // 这个函数在需要为不同参数启动不同实例的时候非常有用
    MakeNew(key string, params []interface{}) (interface{}, error)
}

// HadeContainer 是服务容器的具体实现
type AFContainer struct {
    Container
    // providers 存储注册的服务提供者，key为字符串凭证
    providers map[string]ServiceProvider
    // instance 存储具体的实例，key为字符串凭证
    instances map[string]interface{}
    // lock 用于锁住对容器的变更操作
    lock sync.RWMutex
}

// NewAFContainer 创建一个服务容器
func NewAFContainer() *AFContainer {
    return &AFContainer{
        providers: map[string]ServiceProvider{},
        instances: map[string]interface{}{},
        lock:      sync.RWMutex{},
    }
}

// PrintProviders 输出服务容器中注册的关键字
func (af *AFContainer) PrintProviders() []string {
    ret := []string{}
    for _, provider := range af.providers {
        name := provider.Name()

        line := fmt.Sprint(name)
        ret = append(ret, line)
    }
    return ret
}

// Bind 将服务容器和关键字做了绑定
func (af *AFContainer) Bind(provider ServiceProvider) error {
    af.lock.Lock()
    key := provider.Name()

    af.providers[key] = provider
    af.lock.Unlock()

    // if provider is not defer
    if provider.IsDefer() == false {
        if err := provider.Boot(af); err != nil {
            return err
        }
        // 实例化方法
        params := provider.Params(af)
        method := provider.Register(af)
        instance, err := method(params...)
        if err != nil {
            fmt.Println("bind service provider ", key, " error: ", err)
            return errors.New(err.Error())
        }
        af.instances[key] = instance
    }
    return nil
}

func (af *AFContainer) IsBind(key string) bool {
    return af.findServiceProvider(key) != nil
}

func (af *AFContainer) findServiceProvider(key string) ServiceProvider {
    af.lock.RLock()
    defer af.lock.RUnlock()
    if sp, ok := af.providers[key]; ok {
        return sp
    }
    return nil
}

func (af *AFContainer) Make(key string) (interface{}, error) {
    return af.make(key, nil, false)
}

func (af *AFContainer) MustMake(key string) interface{} {
    serv, err := af.make(key, nil, false)
    if err != nil {
        panic("container not contain key " + key)
    }
    return serv
}

func (af *AFContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
    return af.make(key, params, true)
}

func (af *AFContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
    // force new a
    if err := sp.Boot(af); err != nil {
        return nil, err
    }
    if params == nil {
        params = sp.Params(af)
    }
    method := sp.Register(af)
    ins, err := method(params...)
    if err != nil {
        return nil, errors.New(err.Error())
    }
    return ins, err
}

// 真正的实例化一个服务
func (af *AFContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
    af.lock.RLock()
    defer af.lock.RUnlock()
    // 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
    sp := af.findServiceProvider(key)
    if sp == nil {
        return nil, errors.New("contract " + key + " have not register")
    }

    if forceNew {
        return af.newInstance(sp, params)
    }

    // 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
    if ins, ok := af.instances[key]; ok {
        return ins, nil
    }

    // 容器中还未实例化，则进行一次实例化
    inst, err := af.newInstance(sp, nil)
    if err != nil {
        return nil, err
    }

    af.instances[key] = inst
    return inst, nil
}

// NameList 列出容器中所有服务提供者的字符串凭证
func (af *AFContainer) NameList() []string {
    ret := []string{}
    for _, provider := range af.providers {
        name := provider.Name()
        ret = append(ret, name)
    }
    return ret
}
