package options

import (
    "fmt"
)

// InsecureServingOptions are for creating an unauthenticated, unauthorized, insecure port.
// No one should be using these anymore.
type InsecureServingOptions struct {
    BindAddress string `json:"bind-address" mapstructure:"bind-address"`
    BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

// NewInsecureServingOptions is for creating an unauthenticated, unauthorized, insecure port.
// No one should be using these anymore.
func NewInsecureServingOptions() *InsecureServingOptions {
    return &InsecureServingOptions{
        BindAddress: "127.0.0.1",
        BindPort:    8080,
    }
}

// ApplyTo applies the run options to the method receiver and returns self.
//func (s *InsecureServingOptions) ApplyTo(c *server.ServiceConfig) error {
//    c.InsecureServing = &server.InsecureServingInfo{
//        Address: net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort)),
//    }
//
//    return nil
//}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (s *InsecureServingOptions) Validate() []error {
    var errors []error

    if s.BindPort < 0 || s.BindPort > 65535 {
        errors = append(
            errors,
            fmt.Errorf(
                "--insecure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
                s.BindPort,
            ),
        )
    }

    return errors
}
