package options

import (
    "github.com/jinguoxing/af-go-frame/core/server"
)

type ServerOptions struct {
    RunMode     string   `json:"run_mode"        mapstructure:"run_mode"`
    Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
    Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters.
func NewServerOptions() *ServerOptions {
    defaults := server.NewConfig()

    return &ServerOptions{
        RunMode:     defaults.RunMode,
        Healthz:     defaults.Healthz,
        Middlewares: defaults.Middlewares,
    }
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerOptions) ApplyTo(c *server.ServiceConfig) error {
    c.RunMode = s.RunMode
    c.Healthz = s.Healthz
    c.Middlewares = s.Middlewares

    return nil
}

// Validate checks validation of ServerRunOptions.
func (s *ServerOptions) Validate() []error {
    errors := []error{}

    return errors
}
