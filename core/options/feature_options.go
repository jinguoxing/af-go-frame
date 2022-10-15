package options

import "github.com/jinguoxing/af-go-frame/core/server"

// FeatureOptions contains configuration items related to API server features.
type FeatureOptions struct {
    EnableProfiling bool `json:"profiling"      mapstructure:"profiling"`
    EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

// NewFeatureOptions creates a FeatureOptions object with default parameters.
func NewFeatureOptions() *FeatureOptions {
    defaults := server.NewConfig()

    return &FeatureOptions{
        EnableMetrics:   defaults.EnableMetrics,
        EnableProfiling: defaults.EnableProfiling,
    }
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *FeatureOptions) ApplyTo(c *server.ServiceConfig) error {
    c.EnableProfiling = o.EnableProfiling
    c.EnableMetrics = o.EnableMetrics

    return nil
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *FeatureOptions) Validate() []error {
    return []error{}
}
