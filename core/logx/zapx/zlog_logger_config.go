package zapx

import (
	"go.uber.org/zap"
)

//SetDefault set current logger as default logger
func (l *zapLogger) SetDefault() {
	std = l
}

//Caller enable output log caller
func (l *zapLogger) Caller(flag bool) {
	l.zapLogger.WithOptions(zap.WithCaller(flag))
}

//Development set zap logger development true
func (l *zapLogger) Development() {
	l.zapLogger = l.zapLogger.WithOptions(zap.Development())
}
