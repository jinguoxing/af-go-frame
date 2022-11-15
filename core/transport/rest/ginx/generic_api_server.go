package ginx

import (
    "crypto/tls"
    "github.com/gin-gonic/gin"
    "github.com/jinguoxing/af-go-frame/core/middleware/ginMiddleWare/pprof"
    "github.com/jinguoxing/af-go-frame/core/options"
    "github.com/zeromicro/go-zero/core/logx"
    "net"
    "net/http"
    "net/url"
    "time"
)

//var (
//    _ transport.Server = (*GenericAPIServer)(nil)
//)

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
