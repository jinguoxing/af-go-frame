package zapx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

//ZapWriter
//"github.com/zeromicro/go-zero/core/logx"
// the blow method  implement the logx.Writer
//Writer interface {
//		Alert(v interface{})
//		Close() error
//		Error(v interface{}, fields ...LogField)
//		Info(v interface{}, fields ...LogField)
//		Severe(v interface{})
//		Slow(v interface{}, fields ...LogField)
//		Stack(v interface{})
//		Stat(v interface{}, fields ...LogField)
//	}
type ZapWriter zapLogger

func (z *ZapWriter) Alert(v interface{}) {
	z.zapLogger.Error(fmt.Sprint(v))
}

func (z *ZapWriter) Close() error {
	return z.zapLogger.Sync()
}

func (z *ZapWriter) Error(v interface{}, fields ...logx.LogField) {
	z.zapLogger.Error(fmt.Sprint(v), toZapFields(fields...)...)
}

func (z *ZapWriter) Info(v interface{}, fields ...logx.LogField) {
	z.zapLogger.Info(fmt.Sprint(v), toZapFields(fields...)...)
}

func (z *ZapWriter) Severe(v interface{}) {
	z.zapLogger.Fatal(fmt.Sprint(v))
}

func (z *ZapWriter) Slow(v interface{}, fields ...logx.LogField) {
	z.zapLogger.Warn(fmt.Sprint(v), toZapFields(fields...)...)
}

func (z *ZapWriter) Stack(v interface{}) {
	z.zapLogger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (z *ZapWriter) Stat(v interface{}, fields ...logx.LogField) {
	z.zapLogger.Info(fmt.Sprint(v), toZapFields(fields...)...)
}

func toZapFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
