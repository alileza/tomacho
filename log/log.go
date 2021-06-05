package log

import (
	"encoding/json"
	"log"
	"os"
	"tomato/feature"

	"github.com/cucumber/godog/colors"
)

var Base Logger

func init() {
	Base = Logger{
		l: log.New(os.Stdout, "", 0),
	}
}

type Logger struct {
	l *log.Logger
}

func New() *Logger {
	return &Logger{
		l: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.l.Printf("🟡 %+v", colors.Yellow(args))
}

func (l *Logger) Info(args ...interface{}) {
	l.l.Printf("🔵 %+v", colors.Cyan(args))
}

func (l *Logger) Success(args ...interface{}) {
	l.l.Printf("🟢 %+v", colors.Cyan(args))
}

func (l *Logger) Error(args ...interface{}) {
	l.l.Printf("🔴 %+v", colors.Red(args))
}

func (l *Logger) PrintStep(st feature.Step, err error) {
	if err != nil {
		l.Error(st.String(), err)
	} else {
		l.Success(st.String(), err)
	}
}

func PrintStep(st feature.Step, err error) {
	if err != nil {
		Base.Error(st.String(), err)
	} else {
		Base.Success(st.String(), err)
	}
}

func Debug(args ...interface{}) {
	Base.Debug(args...)
}

func Info(args ...interface{}) {
	Base.Info(args...)
}
func Dump(args ...interface{}) {
	b, err := json.MarshalIndent(args, "", "\t")
	if err != nil {
		Base.Error(err)
	} else {
		Base.Debug(string(b))
	}

}

func Error(args ...interface{}) {
	Base.Error(args...)
}
