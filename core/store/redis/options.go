package redis

import "errors"

var (
    ErrEmptyHost = errors.New("empty redis host")

)

type (
    // A RedisConf is a redis config.
    RedisConf struct {

        Host string
        Type string `json:",default=node,options=node|cluster"`
        Pass string `json:"optional"`
    }

    // A RedisKeyConf is a redis config with key.
    RedisKeyConf struct {
        RedisConf
        Key string `json:",optional"`
    }

)

// Validate validates the RedisConf.
func(rc RedisConf) Validate() error {

    if(len(rc.Host)==0) {
        return ErrEmptyHost
    }


    return nil
}


