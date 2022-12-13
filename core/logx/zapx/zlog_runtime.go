package zapx

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	goRootForFilter  = runtime.GOROOT() // goRootForFilter is used for stack filtering purpose.
	binaryVersion    = ""               // The version of current running binary(uint64 hex).
	binaryVersionMd5 = ""               // The version of current running binary(MD5).
	selfPath         = ""               // Current running binary absolute path.
)

const (
	maxCallerDepth = 1000
	stackFilterKey = "/debug/gdebug/gdebug"
	pathFilterKey  = "/os/glog/glog"
)

// StackWithFilters returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
//
// The parameter `filters` is a slice of strings, which are used to filter the path of the
// caller.
//
// TODO Improve the performance using debug.Stack.
func StackWithFilters(filters []string, skip ...int) string {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}
	var (
		name                  string
		space                 = "  "
		index                 = 1
		buffer                = bytes.NewBuffer(nil)
		ok                    = true
		pc, file, line, start = callerFromIndex(filters)
	)
	for i := start + number; i < maxCallerDepth; i++ {
		if i != start {
			pc, file, line, ok = runtime.Caller(i)
		}
		if ok {
			if filterFileByFilters(file, filters) {
				continue
			}
			if fn := runtime.FuncForPC(pc); fn == nil {
				name = "unknown"
			} else {
				name = fn.Name()
			}
			if index > 9 {
				space = " "
			}
			buffer.WriteString(fmt.Sprintf("%d.%s%s\n    %s:%d\n", index, space, name, file, line))
			index++
		} else {
			break
		}
	}
	return buffer.String()
}

func filterFileByFilters(file string, filters []string) (filtered bool) {
	// Filter empty file.
	if file == "" {
		return true
	}
	// Filter gdebug package callings.
	if strings.Contains(file, stackFilterKey) {
		return true
	}
	for _, filter := range filters {
		if filter != "" && strings.Contains(file, filter) {
			return true
		}
	}
	// GOROOT filter.
	if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
		// https://github.com/gogf/gf/issues/2047
		fileSeparator := file[len(goRootForFilter)]
		if fileSeparator == filepath.Separator || fileSeparator == '\\' || fileSeparator == '/' {
			return true
		}
	}
	return false
}

// callerFromIndex returns the caller position and according information exclusive of the
// debug package.
//
// VERY NOTE THAT, the returned index value should be `index - 1` as the caller's start point.
func callerFromIndex(filters []string) (pc uintptr, file string, line int, index int) {
	var ok bool
	for index = 0; index < maxCallerDepth; index++ {
		if pc, file, line, ok = runtime.Caller(index); ok {
			if filterFileByFilters(file, filters) {
				continue
			}
			if index > 0 {
				index--
			}
			return
		}
	}
	return 0, "", -1, -1
}
