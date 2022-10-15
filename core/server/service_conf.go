package server

import (
    "time"
)

type ServiceConfig struct {
    Name            string
    RunMode         string `json:",default=pro,options=dev|test|rt|pre|pro"`
    Healthz         bool
    Middlewares     []string
    EnableProfiling bool
    EnableMetrics   bool
    MetricsUrl      string `json:",optional"`

    Jwt *JwtInfo
}

// CertKey contains configuration items related to certificate.
type CertKey struct {
    // CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
    CertFile string
    // KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
    KeyFile string
}

// SecureServingInfo holds configuration of the TLS server.
type SecureServingInfo struct {
    BindAddress string
    BindPort    int
    CertKey     CertKey
}

// InsecureServingInfo holds configuration of the insecure http server.
type InsecureServingInfo struct {
    Address string
}

// JwtInfo defines jwt fields used to create jwt authentication middleware.
type JwtInfo struct {
    // defaults to "iam jwt"
    Realm string
    // defaults to empty
    Key string
    // defaults to one hour
    Timeout time.Duration
    // defaults to zero
    MaxRefresh time.Duration
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *ServiceConfig {
    return &ServiceConfig{
        Healthz:         true,
        RunMode:         ProMode,
        Middlewares:     []string{},
        EnableProfiling: true,
        EnableMetrics:   true,
        Jwt: &JwtInfo{
            Realm:      "afGo jwt",
            Timeout:    1 * time.Hour,
            MaxRefresh: 1 * time.Hour,
        },
    }
}

//func (sc ServiceConfig) InitMode() {
//    switch sc.RunMode {
//    case DevMode, TestMode:
//
//    }
//}
