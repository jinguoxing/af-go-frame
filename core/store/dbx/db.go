package dbx

import (
    "github.com/jinguoxing/af-go-frame/core/container/mapx"
    "time"
)

const(

    defaultCharset          = `utf8`
    defaultMaxIdleConnCount = 10               // Max idle connection count in pool.
    defaultMaxOpenConnCount = 10                // Max open connection count in pool. O is no limit.
    defaultMaxConnLifeTime  = 30 * time.Second // Max lifetime for per connection in pool in seconds.
)


var (
    instances = mapx.NewStrAnyMap(true)

)