package mapx

import (
    "github.com/jinguoxing/af-go-frame/core/syncx/rwmutex"
)

// map以支付串为key 任意类型为值


type StrAnyMap struct {
    mu   rwmutex.RWMutex
    data map[string]interface{}
}


func NewStrAnyMap(safe ...bool) *StrAnyMap {

    return &StrAnyMap{
        mu:   rwmutex.Create(safe...),
        data: make(map[string]interface{}),
    }

}

func (m *StrAnyMap) Clear() {
    m.mu.Lock()
    m.data = make(map[string]interface{})
    m.mu.Unlock()
}


