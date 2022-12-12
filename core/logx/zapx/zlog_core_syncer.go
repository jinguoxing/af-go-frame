package zapx

import (
	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
)

var syncMap map[string]zapcore.WriteSyncer

func init() {
	syncMap = make(map[string]zapcore.WriteSyncer)
}

//StdoutWriteSyncer write log to stdout
func (c CoreConfig) StdoutWriteSyncer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

//FileWriteSyncer write log to zapcore.WriteSyncer
func (c CoreConfig) FileWriteSyncer() zapcore.WriteSyncer {
	// check file path, create file if not exists
	checkDir(c.Destination)
	// set filename format
	fileNameFormat := DefaultFileNameFormat
	if c.FileNameFormat != "" {
		fileNameFormat = c.FileNameFormat
	}
	hook, _ := rotatelogs.New(
		strings.Replace(c.Destination, ".log", "", -1)+fileNameFormat,
		rotatelogs.WithLinkName(c.Destination),
		rotatelogs.WithRotationSize(parseSize(c.RotateSize)),
	)
	return zapcore.AddSync(hook)
}

func RegisterWriteSyncer(coreType string, syncer io.Writer) {
	syncMap[coreType] = zapcore.AddSync(syncer)
}

//GetWriteSyncer get writeSyncer by different
func (c CoreConfig) GetWriteSyncer(coreType string) zapcore.WriteSyncer {
	if coreType == StdoutCore {
		return c.StdoutWriteSyncer()
	}
	if coreType == LocalFileCore {
		return c.FileWriteSyncer()
	}
	syncer, ok := syncMap[coreType]
	if ok {
		return syncer
	}
	log.Warnf("WriteSyncer %s not found, use stdout writeSyncer", coreType)
	return zapcore.Lock(os.Stdout)
}
