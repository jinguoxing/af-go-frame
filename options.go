package af_go_frame

import (
    "context"
    baseOptions "github.com/jinguoxing/af-go-frame/core/options"
    "github.com/jinguoxing/af-go-frame/core/registry"
    "github.com/jinguoxing/af-go-frame/core/transport"
    "log"
    "net/url"
    "os"
    "time"
)

type options struct {
    id   string
    name string

    endpoints []*url.URL

    ctx  context.Context
    sigs []os.Signal

    baseOptions baseOptions.Options

    logger           log.Logger
    registrar        registry.Registrar
    registrarTimeout time.Duration
    stopTimeout      time.Duration
    servers          []transport.Server
}

type Option func(o *options)

//ID  设置Server的ID
func ID(id string) Option {

    return func(o *options) {
        o.id = id
    }
}

// 设置Server的Name
func Name(name string) Option {

    return func(o *options) {
        o.name = name
    }
}

// Endpoint with service endpoint.
func Endpoint(endpoints ...*url.URL) Option {
    return func(o *options) { o.endpoints = endpoints }
}

// Context with service context.
func Context(ctx context.Context) Option {
    return func(o *options) { o.ctx = ctx }
}

// Logger with service logger.
func Logger(logger log.Logger) Option {
    return func(o *options) { o.logger = logger }
}

func BaseOptions(b baseOptions.Options) Option {

    return func(o *options) {
        o.baseOptions = b
    }
}

// Server with transport servers.
func Server(srv ...transport.Server) Option {
    return func(o *options) { o.servers = srv }
}

// Signal with exit signals.
func Signal(sigs ...os.Signal) Option {
    return func(o *options) { o.sigs = sigs }
}

// Registrar with service registry.
func Registrar(r registry.Registrar) Option {
    return func(o *options) { o.registrar = r }
}

// RegistrarTimeout with registrar timeout.
func RegistrarTimeout(t time.Duration) Option {
    return func(o *options) { o.registrarTimeout = t }
}

// StopTimeout with app stop timeout.
func StopTimeout(t time.Duration) Option {
    return func(o *options) { o.stopTimeout = t }
}
