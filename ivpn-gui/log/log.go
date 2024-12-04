package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

// _defaultLevel is package default logging level.
var _defaultLevel = atomic.NewUint32(uint32(InfoLevel))
var CurrentLogs []string
var guiMode bool

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func SetLevel(level Level) {
	_defaultLevel.Store(uint32(level))
}

func Debugf(format string, args ...any) {
	logf(DebugLevel, format, args...)
}

func Infof(format string, args ...any) {
	logf(InfoLevel, format, args...)
}

func Warnf(format string, args ...any) {
	logf(WarnLevel, format, args...)
}

func Errorf(format string, args ...any) {
	logf(ErrorLevel, format, args...)
}

func Fatalf(format string, args ...any) {
	logrus.Fatalf(format, args...)
}
func SetGUIMode(guiMode_l bool) {
	guiMode = guiMode_l
	if guiMode_l {
		CurrentLogs = make([]string, 0)
	}
}
func logf(level Level, format string, args ...any) {
	event := newEvent(level, format, args...)
	if uint32(event.Level) > _defaultLevel.Load() {
		return
	}
	if guiMode {
		CurrentLogs = append(CurrentLogs, fmt.Sprintf("%s %s: %s", level.String(), event.Time.String(), event.Message))
		if len(CurrentLogs) > 20 {
			CurrentLogs = CurrentLogs[len(CurrentLogs)-20:]
		}
	}
	switch level {
	case DebugLevel:
		logrus.WithTime(event.Time).Debugln(event.Message)
	case InfoLevel:
		logrus.WithTime(event.Time).Infoln(event.Message)
	case WarnLevel:
		logrus.WithTime(event.Time).Warnln(event.Message)
	case ErrorLevel:
		logrus.WithTime(event.Time).Errorln(event.Message)
	}
}
