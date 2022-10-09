package options

import "time"



type AppOptions struct {
    DefaultPageSize       int
    MaxPageSize           int
    DefaultContextTimeout time.Duration
    LogSavePath           string
    LogFileName           string
    LogFileExt            string
    UploadSavePath        string
    UploadServerUrl       string
    UploadImageMaxSize    int
    UploadImageAllowExts  []string
}
