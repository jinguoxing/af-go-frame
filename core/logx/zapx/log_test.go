package zapx_test

import (
    "github.com/jinguoxing/af-go-frame/core/logx/zapx"
    "testing"
)

func Test_WithName(t *testing.T) {
    defer zapx.Flush() // used for record logger printer

    logger := zapx.WithName("test")
    logger.Infow("Hello world!", "foo", "bar") // structed logger
}

func Test_WithValues(t *testing.T) {
    defer zapx.Flush() // used for record logger printer

    logger := zapx.WithValues("key", "value") // used for record context
    logger.Info("Hello world!")
    logger.Info("Hello world!")
}

func Test_V(t *testing.T) {
    defer zapx.Flush() // used for record logger printer

    zapx.V(0).Infow("Hello world!", "key", "value")
    zapx.V(1).Infow("Hello world!", "key", "value")
}
