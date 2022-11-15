package options

import (
    "github.com/jinguoxing/af-go-frame/core/logx"
    "github.com/jinguoxing/af-go-frame/core/server"
)

type CommonOptions interface {
    Validate() []error
}

// CompleteableOptions abstracts options which can be completed.
type CompleteableOptions interface {
    Complete() error
}

type Options struct {
    GenericServerOptions *ServerOptions  `json:"server"  mapstructure:"server"`
    FeatureOptions       *FeatureOptions `json:"feature" mapstructure:"feature"`
    Log                  *logx.LogConf   `json:"log"  mapstructure:"log"`

    InsecureServing *InsecureServingOptions `json:"insecure" mapstructure:"insecure"`

    MySQLOptions *DBOptions `json:"mysql"    mapstructure:"mysql"`
}

//NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
    o := Options{
        GenericServerOptions: NewServerOptions(),
        InsecureServing:      NewInsecureServingOptions(),
        FeatureOptions:       NewFeatureOptions(),
    }

    return &o
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.ServiceConfig) error {
    return nil
}
