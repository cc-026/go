package util

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

//go:generate stringer -type eLogType -linecomment
type eLogType int //LogType
const (
	eLogType_Debug  eLogType = iota // [DEBUG]
	eLogType_Info                   // [INFO]
	eLogType_Waring                 // [WARN]
	eLogType_Error                  // [ERROR]
)

type Logger interface {
	LogDebug(a ...any)
	LogInfo(a ...any)
	LogWaring(a ...any)
	LogError(a ...any)
	depth() int
}

type logger struct {
}

func (l *logger) LogDebug(a ...any)  { doLog(2, eLogType_Debug, a...) }
func (l *logger) LogInfo(a ...any)   { doLog(2, eLogType_Info, a...) }
func (l *logger) LogWaring(a ...any) { doLog(2, eLogType_Waring, a...) }
func (l *logger) LogError(a ...any)  { doLog(2, eLogType_Error, a...) }
func (l *logger) depth() int         { return 0 }

type tagLogger struct {
	tag    string
	dep    int
	parent Logger
}

func (l *tagLogger) LogDebug(a ...any)  { doLog(2, eLogType_Debug, l.append(a...)...) }
func (l *tagLogger) LogInfo(a ...any)   { doLog(2, eLogType_Info, l.append(a...)...) }
func (l *tagLogger) LogWaring(a ...any) { doLog(2, eLogType_Waring, l.append(a...)...) }
func (l *tagLogger) LogError(a ...any)  { doLog(2, eLogType_Error, l.append(a...)...) }
func (l *tagLogger) depth() int         { return l.dep }
func (l *tagLogger) append(a ...any) []any {
	a = append(make([]any, l.dep, l.dep+len(a)), a...)
	for tl := l; tl != nil; {
		a[tl.depth()-1] = tl.tag
		tl, _ = tl.parent.(*tagLogger)
	}

	return a
}

var (
	loggerInstance Logger
	logInitOnce    sync.Once
)

func Log() Logger {
	logInitOnce.Do(func() {
		loggerInstance = new(logger)
	})
	return loggerInstance
}

func LogWithTag(tag string) Logger {
	return SubLogWithTag(nil, tag)
}

func SubLogWithTag(parent Logger, tag string) Logger {
	if parent == nil {
		parent = Log()
	}

	return &tagLogger{
		parent: parent,
		dep:    parent.depth() + 1,
		tag:    tag,
	}
}

func doLog(skip int, logType eLogType, a ...any) {
	_, file, line, ok := runtime.Caller(skip)
	if false == ok {
		file = "UNKNOWN"
		line = -1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	a = append([]any{time.Now().Format("2006-01-02 15:04:05 MST"), "\t", logType, "\t", file, ":", line, "\t"}, a...)
	fmt.Println(a...)
	if eLogType_Waring == logType || eLogType_Error == logType {
		fmt.Println(string(debug.Stack()))
	}
}
