package options

import (
    "github.com/jinguoxing/af-go-frame/core/server"
    "time"
)

type ServerOptions struct {
    RunMode      string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

func NewServerOptions() *ServerOptions {

    defaults := server.NewConfig()

    return &ServerOptions{
        RunMode: defaults.RunMode,
    }

}