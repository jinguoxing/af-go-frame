package zapx

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//RegisterHook register hook
func (l *zapLogger) RegisterHook(hooks ...func(zapcore.Entry) error) {
	option := zap.Hooks(hooks...)
	l.zapLogger = l.zapLogger.WithOptions(option)
}

func (l *zapLogger) WithHook(hooks ...func(zapcore.Entry) error) Logger {
	l.RegisterHook(hooks...)
	return l
}

//SampleHook demo hook
func SampleHook(entry zapcore.Entry) error {
	bts, _ := json.Marshal(entry)
	fmt.Println("SampleHook: %s", string(bts))
	return nil
}
