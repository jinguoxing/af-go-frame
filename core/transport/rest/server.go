package rest

import (
    "github.com/gin-gonic/gin"
    "github.com/jinguoxing/af-go-frame/core/internal/endpoint"
    "github.com/jinguoxing/af-go-frame/core/internal/host"
    "github.com/jinguoxing/af-go-frame/core/transport"
    "context"
    "crypto/tls"
    "errors"
    "fmt"
    "net"
    "net/http"
    "net/url"
    "time"

)

var (
    _ transport.Server     = (*Server)(nil)
)


// ServerOption is an HTTP server option.
type ServerOption func(*Server)



// Address with server address.
func Address(addr string) ServerOption {
    return func(s *Server) {
        s.address = addr
    }
}

type Server struct {

    *http.Server
    lis         net.Listener
    tlsConf     *tls.Config
    endpoint    *url.URL
    err         error
    network     string
    address     string
    timeout     time.Duration
    //filters     []FilterFunc
    //middleware  matcher.Matcher
    //dec         DecodeRequestFunc
    //enc         EncodeResponseFunc
    //ene         EncodeErrorFunc
    strictSlash bool
}



// NewServer creates an HTTP server by options.
func NewServer(r *gin.Engine,opts ...ServerOption) *Server {
    srv := &Server{
        network:     "tcp",
        address:     ":0",
        timeout:     1 * time.Second,
        //middleware:  matcher.New(),
        //dec:         DefaultRequestDecoder,
        //enc:         DefaultResponseEncoder,
        //ene:         DefaultErrorEncoder,
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
        Handler:   r,
        TLSConfig: srv.tlsConf,
    }
    return srv
}


// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
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

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
    fmt.Print("[HTTP] server stopping")
    return s.Shutdown(ctx)
}

func (s *Server) listenAndEndpoint() error {
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
