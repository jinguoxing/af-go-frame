package rwmutex

import (
    "sync"
)

type RWMutex struct {
    *sync.RWMutex
}


func New(safe ...bool) *RWMutex {

    mu := Create(safe...)
    return &mu

}


func Create(safe ...bool)RWMutex{

    mu := RWMutex{}
    if len(safe) > 0 && safe[0] {
        mu.RWMutex = new(sync.RWMutex)
    }
    return mu
}

func(mu *RWMutex) IsSafe() bool {

    return mu.RWMutex != nil
}

func(mu *RWMutex) Lock(){

    if mu.RWMutex != nil {
        mu.RWMutex.Lock()
    }
}

func(mu *RWMutex) UnLock(){

    if mu.RWMutex !=nil {
        mu.RWMutex.Unlock()
    }
}

func(mu *RWMutex) RLock(){

    if mu.RWMutex != nil {
        mu.RWMutex.RLock()
    }
}

func(mu *RWMutex) RUnLock(){

    if mu.RWMutex != nil {
        mu.RWMutex.RUnlock()
    }
}





