package rest

type (
    ResetConfig struct {
        Host     string `json:",default=0.0.0.0"`
        Port     int
        CertFile string `json:",optional"`
        KeyFile  string `json:",optional"`
        Verbose  bool   `json:",optional"`
        MaxConns int    `json:",default=10000"`
        MaxBytes int64  `json:",default=1048576"`
        // milliseconds
        Timeout int64 `json:",default=3000"`
    }
)
