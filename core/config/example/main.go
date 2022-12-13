package main

import (
	"fmt"
	"github.com/jinguoxing/af-go-frame/core/config"
	"github.com/jinguoxing/af-go-frame/core/config/sources/env"
	"github.com/jinguoxing/af-go-frame/core/config/sources/file"
	"os"
	"path"
	"runtime"
)

var pwd string = "."

func init() {
	os.Setenv(config.ProjectEnvKey, "DEV")

	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		pwd = abPath
	}
}

type CoreConfig struct {
	Destination string `json:"destination"` //log output destination
	LogLevel    string `json:"log_level"`   //lowest log level in this core
}

type Options struct {
	Name        string       `json:"name"`
	CoreConfigs []CoreConfig `json:"cores"`
}

type LogConfigs struct {
	Logs []Options `json:"logs"`
}

func InitSources(paths ...string) {
	sources := make([]config.Source, len(paths)+1)
	sources[0] = env.NewSource()
	for i, path := range paths {
		sources[i+1] = file.NewSource(path)
	}
	config.Init(sources...)
}

func main() {
	InitSources("config.yaml")
	logs := config.Scan[LogConfigs]()

	fmt.Printf("logs config %v", logs)

}
