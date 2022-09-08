package redis


import (
    "af-go-frame/core/syncx"
    "crypto/tls"
    stdRedis "github.com/go-redis/redis/v8"
    "io"
)

const (
    defaultDatabase = 0
    maxRetries      = 3
    idleConns       = 8
)


var clientManager = syncx.NewResourceManager()

func getClient(r *Redis)(*stdRedis.Client,error){

    val, err := clientManager.GetResource(r.Addr, func() (io.Closer, error) {

        var tlsConfig *tls.Config
        if r.tls {
            tlsConfig = &tls.Config{
                InsecureSkipVerify: true,
            }
        }

        store := stdRedis.NewClient(&stdRedis.Options{
            Addr:         r.Addr,
            Password:     r.Pass,
            DB:           defaultDatabase,
            MaxRetries:   maxRetries,
            MinIdleConns: idleConns,
            TLSConfig:    tlsConfig,
        })

        return store,nil
    })

    if err != nil {
        return nil, err
    }

    return val.(*stdRedis.Client),nil

}