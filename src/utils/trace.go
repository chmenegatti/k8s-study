package utils

import (
	"fmt"
	"runtime"
)

func Trace() (info string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	info = fmt.Sprintf("%s:%v", frame.Function, frame.Line)
	return
}
