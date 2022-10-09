package options

import "time"

type ServerOptions struct {
    RunMode      string
    HttpPort     string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}
