package redis

import (
    "context"
    "fmt"
    stdRedis "github.com/go-redis/redis/v8"
    "time"
)


const (
    // ClusterType means redis cluster.
    ClusterType = "cluster"
    // NodeType means redis node.
    NodeType = "node"
    // Nil is an alias of redis.Nil.
    Nil = stdRedis.Nil

    blockingQueryTimeout = 5 * time.Second
    readWriteTimeout     = 2 * time.Second
    defaultSlowThreshold = time.Millisecond * 100
)

type (

    // Redis defines a redis node/cluster. It is thread-safe.
    Redis struct {
        Addr string
        Type string
        Pass string
        tls  bool
    }

    // Option defines the method to customize a Redis.
    Option func(r *Redis)

    RedisNode interface {
        stdRedis.Cmdable
    }

)

func New(addr string,opts ...Option) *Redis {
    r := &Redis{
        Addr:addr,
        Type:NodeType,
    }

    for _,opt := range opts {
        opt(r)
    }

    return r
}

// Cluster customizes the given Redis as a cluster.
func Cluster() Option {

    return func(r *Redis){
        r.Type = ClusterType
    }
}

func WithPass(pass string)Option {
    return func(r *Redis){
        r.Pass = pass
    }
}

// WithTLS customizes the given Redis with TLS enabled.
func WithTLS() Option {
    return func(r *Redis) {
        r.tls = true
    }
}


func getRedis(r *Redis) (RedisNode, error) {
    switch r.Type {

    case NodeType:
        return getClient(r)
    default:
        return nil, fmt.Errorf("redis type '%s' is not supported", r.Type)
    }
}

func (s *Redis) Decr(key string) (int64, error) {

    return s.DecrCtx(context.Background(), key)
}

func (s *Redis) DecrCtx(ctx context.Context, key string) (value int64, err error) {

    conn, err := getRedis(s)
    if err != nil {
        return
    }

   return  conn.Decr(ctx, key).Result()

}



