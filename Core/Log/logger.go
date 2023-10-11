package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"golang.org/x/exp/constraints"
)

var (
	mutex = &sync.Mutex{}
)

const (
	Info  = 1
	Error = 2
	Fatal = 3
)

var GetStackTrace = func() string {
	pc, fn, line, _ := runtime.Caller(1)
	stackTrace := fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)

	return stackTrace
}

func AppendLine(filepath string, s string, m *sync.Mutex) error {
	m.Lock()
	defer m.Unlock()
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(s + "\n"); err != nil {
		return err
	}

	return nil
}

func GetExplicitTime() string {
	return time.Now().Format("01-02-2006 15:04:05")
}

func LogError(stack string, format string, content ...any) {
	log.Error(fmt.Sprintf(format, content...))
	LogMessages[int, string](Error, fmt.Sprintf(format, content...), stack)
}

func LogInfo(stack string, format string, content ...any) {
	log.Info(fmt.Sprintf(format, content...))
	LogMessages[int, string](Info, fmt.Sprintf(format, content...), stack)
}

func LogPanic(stack string, format string, content ...any) {
	LogMessages[int, string](Fatal, fmt.Sprintf(format, content...), stack)
	log.Fatal(fmt.Sprintf(format, content...))
}

// noinspection GoUnusedConst
func LogMessages[T constraints.Integer, S comparable](level T, message S, stackTrace string) {
	var (
		log string
	)

	switch level {
	case Info:
		log = fmt.Sprintf("%v [INFO] %v -> %v", GetExplicitTime(), stackTrace, message)
	case Error:
		log = fmt.Sprintf("%v [ERROR] %v -> %v", GetExplicitTime(), stackTrace, message)
	case Fatal:
		log = fmt.Sprintf("%v [FATAL] %v -> %v", GetExplicitTime(), stackTrace, message)
	}

	if err := AppendLine("Logs/events.log", log, mutex); err != nil {
		LogError(GetStackTrace(), "Could not log message: %v, Error: %v", message, err.Error())
	}
}
