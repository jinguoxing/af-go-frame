package ginMiddleWare

import (
    "github.com/gin-gonic/gin"
    "github.com/zeromicro/go-zero/core/logx"
    "go.opentelemetry.io/otel/trace"
    "net"
    "net/http"
    "net/http/httputil"
    "os"
    "runtime/debug"
    "strings"
    "time"
)

// Config is config setting for Ginzap
type Config struct {
    TimeFormat string
    UTC        bool
    SkipPaths  []string
    TraceID    bool // optionally log Open Telemetry TraceID
}


// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//  1. A time package format string (e.g. time.RFC3339).
//  2. A boolean stating whether to use UTC time zone or local.
func GinZap(w logx.Writer, timeFormat string, utc bool) gin.HandlerFunc {
    return GinZapWithConfig(w, &Config{TimeFormat: timeFormat, UTC: utc})
}


func GinZapWithConfig(w logx.Writer ,conf *Config) gin.HandlerFunc {

    skipPaths := make(map[string]bool, len(conf.SkipPaths))
    for _, path := range conf.SkipPaths {
        skipPaths[path] = true
    }


    return func(c *gin.Context){

            start := time.Now()

            path := c.Request.URL.Path
            query := c.Request.URL.RawQuery
            c.Next()

        if _, ok := skipPaths[path]; !ok {
            end := time.Now()
            latency := end.Sub(start)
            if conf.UTC {
                end = end.UTC()
            }

            if len(c.Errors) > 0 {

                for _, e := range c.Errors.Errors() {
                    w.Error(e)
                }

            } else {

                fields := []logx.LogField{
                    logx.Field("status",c.Writer.Status()),
                    logx.Field("method", c.Request.Method),
                    logx.Field("path", path),
                    logx.Field("query", query),
                    logx.Field("ip", c.ClientIP()),
                    logx.Field("user-agent", c.Request.UserAgent()),
                    logx.Field("latency", latency),
                }

                if conf.TimeFormat != "" {
                    fields = append(fields, logx.Field("time", end.Format(conf.TimeFormat)))
                }
                if conf.TraceID {
                    fields = append(fields, logx.Field("traceID", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
                }
                w.Info(path, fields...)
             }
        }
    }
}


func RecoveryWithZap(w logx.Writer,stack bool) gin.HandlerFunc {

    return func(c *gin.Context){
        defer func() {
            if err:= recover(); err!=nil{
                // Check for a broken connection, as it is not really a
                // condition that warrants a panic stack trace.
                var brokenPipe bool
                if ne, ok := err.(*net.OpError); ok {
                    if se, ok := ne.Err.(*os.SyscallError); ok {
                        if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
                            brokenPipe = true
                        }
                    }
                }

                httpRequest, _ := httputil.DumpRequest(c.Request, false)
                if brokenPipe {
                    w.Error(c.Request.URL.Path,

                        logx.Field("error", err),
                        logx.Field("request", string(httpRequest)),
                    )
                    // If the connection is dead, we can't write a status to it.
                    c.Error(err.(error)) // nolint: errcheck
                    c.Abort()
                    return
                }

                if stack {
                    w.Error("[Recovery from panic]",
                        logx.Field("time", time.Now()),
                        logx.Field("error", err),
                        logx.Field("request", string(httpRequest)),
                        logx.Field("stack", string(debug.Stack())),
                    )
                } else {
                    w.Error("[Recovery from panic]",
                        logx.Field("time", time.Now()),
                        logx.Field("error", err),
                        logx.Field("request", string(httpRequest)),
                    )
                }
                c.AbortWithStatus(http.StatusInternalServerError)
            }

        }()

        c.Next()
    }

}

