package ginx

import (
    "context"
    "crypto/tls"
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinguoxing/af-go-frame/core/internal/endpoint"
    "github.com/jinguoxing/af-go-frame/core/internal/host"
    "github.com/jinguoxing/af-go-frame/core/middleware/ginMiddleWare"
    "github.com/jinguoxing/af-go-frame/core/middleware/ginMiddleWare/pprof"
    "github.com/jinguoxing/af-go-frame/core/options"
    "github.com/jinguoxing/af-go-frame/core/transport"
    "github.com/zeromicro/go-zero/core/logx"
    "net"
    "net/http"
    "net/url"
    "time"
)

var (
    _ transport.Server = (*GenericAPIServer)(nil)
)

// ServerOption is an HTTP server option.
type ServerOption func(options *options.Options)

// Address with server address.
func Address(addr string) ServerOption {
    return func(s *GenericAPIServer) {
        s.address = addr
    }
}

type GenericAPIServer struct {
    *http.Server
    *gin.Engine

    lis      net.Listener
    tlsConf  *tls.Config
    endpoint *url.URL
    err      error
    network  string
    address  string
    timeout  time.Duration

    // ShutdownTimeout is the timeout used for server shutdown. This specifies the timeout before server
    // gracefully shutdown returns.
    ShutdownTimeout time.Duration

    healthz         bool
    enableMetrics   bool
    enableProfiling bool

    middlewares []string

    strictSlash bool
}

func initGenericAPIServer(s *GenericAPIServer) {
    s.Setup()
    s.InstallMiddlewares()
    s.InstallAPIs()
}

//InstallAPIs
func (s *GenericAPIServer) InstallAPIs() {

    if s.enableProfiling {
        pprof.Register(s.Engine)
    }

}

// Setup do some setup work for gin engine.
func (s *GenericAPIServer) Setup() {
    gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
        logx.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
    }
}

func (s *GenericAPIServer) InstallMiddlewares() {

    for _, m := range s.middlewares {
        mw, ok := ginMiddleWare.GinMiddlewares[m]
        if !ok {
            continue
        }
        logx.Infof("install middleware: %s", m)
        s.Use(mw)
    }
}

// NewServer creates an HTTP server by options.
func NewServer(opts ...ServerOption) *GenericAPIServer {

    // gin.SetMode

    srv := &GenericAPIServer{
        network:     "tcp",
        address:     ":0",
        timeout:     1 * time.Second,
        Engine:      gin.New(),
        strictSlash: true,
    }
    for _, o := range opts {
        o(srv)
    }
    // router := gin.Default()
    //srv.router = mux.NewRouter().StrictSlash(srv.strictSlash)
    //srv.router.NotFoundHandler = http.DefaultServeMux
    //srv.router.MethodNotAllowedHandler = http.DefaultServeMux
    //srv.router.Use(srv.filter())
    srv.Server = &http.Server{
        Handler:   srv,
        TLSConfig: srv.tlsConf,
    }

    initGenericAPIServer(srv)
    return srv
}

//
//// CompletedConfig is the completed configuration for GenericAPIServer.
//type CompletedConfig struct {
//    *server.ServiceConfig
//}
//
//// Complete fills in any fields not set that are required to have valid data and can be derived
//// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
//func (c *server.ServiceConfig) Complete() CompletedConfig {
//    return CompletedConfig{c}
//}
//
//// New returns a new instance of GenericAPIServer from the given config.
//func (c CompletedConfig) New() (*GenericAPIServer, error) {
//    // setMode before gin.New()
//    gin.SetMode(c.RunMode)
//
//    s := &GenericAPIServer{
//        healthz:         c.Healthz,
//        enableMetrics:   c.EnableMetrics,
//        enableProfiling: c.EnableProfiling,
//        middlewares:     c.Middlewares,
//        Engine:          gin.New(),
//    }
//
//    initGenericAPIServer(s)
//
//    return s, nil
//}

//Start start the HTTP server.
func (s *GenericAPIServer) Start(ctx context.Context) error {
    if err := s.listenAndEndpoint(); err != nil {
        return err
    }
    s.BaseContext = func(net.Listener) context.Context {
        return ctx
    }

    fmt.Printf("[HTTP] server listening on: %s", s.lis.Addr().String())
    var err error
    if s.tlsConf != nil {
        err = s.ServeTLS(s.lis, "", "")
    } else {
        err = s.Serve(s.lis)

    }

    if !errors.Is(err, http.ErrServerClosed) {
        return err
    }
    return nil
}

// Stop  Close graceful shutdown the api server.
func (s *GenericAPIServer) Stop(ctx context.Context) error {
    fmt.Print("[HTTP] server stopping")
    return s.Shutdown(ctx)
}

func (s *GenericAPIServer) listenAndEndpoint() error {
    if s.lis == nil {
        lis, err := net.Listen(s.network, s.address)
        if err != nil {
            s.err = err
            return err
        }
        s.lis = lis
    }
    if s.endpoint == nil {
        addr, err := host.Extract(s.address, s.lis)
        if err != nil {
            s.err = err
            return err
        }
        s.endpoint = endpoint.NewEndpoint(endpoint.Scheme("http", s.tlsConf != nil), addr)
    }
    return s.err
}
