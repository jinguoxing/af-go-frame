package agerrors

import (
    "fmt"
    "io"
    "path"
    "runtime"
    "strconv"
    "strings"
)

// Frame represents a program counter inside a stack frame.
// For historical reasons if Frame is interpreted as a uintptr
// its value represents the program counter + 1.
type Frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { return uintptr(f) - 1 }

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f Frame) file() string {
    fn := runtime.FuncForPC(f.pc())
    if fn == nil {
        return "unknown"
    }
    file, _ := fn.FileLine(f.pc())
    return file
}

// line returns the line number of source code of the
// function for this Frame's pc.
func (f Frame) line() int {
    fn := runtime.FuncForPC(f.pc())
    if fn == nil {
        return 0
    }
    _, line := fn.FileLine(f.pc())
    return line
}

// name returns the name of this function, if known.
func (f Frame) name() string {
    fn := runtime.FuncForPC(f.pc())
    if fn == nil {
        return "unknown"
    }
    return fn.Name()
}

// Format formats the frame according to the fmt.Formatter interface.
//
//    %s    source file
//    %d    source line
//    %n    function name
//    %v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+s   function name and path of source file relative to the compile time
//          GOPATH separated by \n\t (<funcname>\n\t<path>)
//    %+v   equivalent to %+s:%d
func (f Frame) Format(s fmt.State, verb rune) {
    switch verb {
    case 's':
        switch {
        case s.Flag('+'):
            io.WriteString(s, f.name())
            io.WriteString(s, "\n\t")
            io.WriteString(s, f.file())
        default:
            io.WriteString(s, path.Base(f.file()))
        }
    case 'd':
        io.WriteString(s, strconv.Itoa(f.line()))
    case 'n':
        io.WriteString(s, funcname(f.name()))
    case 'v':
        f.Format(s, 's')
        io.WriteString(s, ":")
        f.Format(s, 'd')
    }
}

// MarshalText formats a stacktrace Frame as a text string. The output is the
// same as that of fmt.Sprintf("%+v", f), but without newlines or tabs.
func (f Frame) MarshalText() ([]byte, error) {
    name := f.name()
    if name == "unknown" {
        return []byte(name), nil
    }
    return []byte(fmt.Sprintf("%s %s:%d", name, f.file(), f.line())), nil
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
    i := strings.LastIndex(name, "/")
    name = name[i+1:]
    i = strings.Index(name, ".")
    return name[i+1:]
}

const (
    // maxStackDepth marks the max stack depth for error back traces.
    maxStackDepth = 32
)

type stack []uintptr

// callers returns the stack callers.
// Note that it here just retrieves the caller memory address array not the caller information.
func callers(skip ...int) stack {
    var (
        pcs [maxStackDepth]uintptr
        n   = 3
    )
    if len(skip) > 0 {
        n += skip[0]
    }
    return pcs[:runtime.Callers(n, pcs[:])]
}
