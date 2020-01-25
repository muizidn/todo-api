package app

import (
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logger

func _CreateLogger(module string) *logrus.Entry {
	logrus.SetOutput(os.Stdout)
	entry := logrus.WithFields(logrus.Fields{
		"module": module,
	})
	return entry
}

func setupLogger() {
	entry := _CreateLogger("app")
	log = &logger{entry: entry}
}

type logger struct {
	entry *logrus.Entry
}

func _CallerInfo() logrus.Fields {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		comp := strings.Split(file, "/")
		return logrus.Fields{
			"file": comp[len(comp)-1],
			"line": line,
		}
	}
	return nil
}

func (l *logger) WithFields(fields map[string]interface{}) *logger {
	logrusFields := logrus.Fields{}
	for k, v := range fields {
		logrusFields[k] = v
	}
	newLogger := &logger{entry: l.entry.WithFields(logrusFields)}
	return newLogger
}

/// ToDo log interface
func (l *logger) TError(err error) error {
	l.Error(err)
	return err
}

/// logrus interface

func (l *logger) Trace(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Trace(args...)
}
func (l *logger) Debug(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Debug(args...)
}
func (l *logger) Print(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Print(args...)
}
func (l *logger) Info(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Info(args...)
}
func (l *logger) Warn(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Warn(args...)
}
func (l *logger) Warning(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Warning(args...)
}
func (l *logger) Error(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Error(args...)
}
func (l *logger) Panic(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Panic(args...)
}
func (l *logger) Fatal(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Fatal(args...)
}
func (l *logger) Tracef(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Tracef(format, args...)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Debugf(format, args...)
}
func (l *logger) Printf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Printf(format, args...)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Infof(format, args...)
}
func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Warnf(format, args...)
}
func (l *logger) Warningf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Warningf(format, args...)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Errorf(format, args...)
}
func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Panicf(format, args...)
}
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Fatalf(format, args...)
}
func (l *logger) Traceln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Traceln(args...)
}
func (l *logger) Debugln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Debugln(args...)
}
func (l *logger) Println(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Println(args...)
}
func (l *logger) Infoln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Infoln(args...)
}
func (l *logger) Warnln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Warnln(args...)
}
func (l *logger) Warningln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Warningln(args...)
}
func (l *logger) Errorln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Errorln(args...)
}
func (l *logger) Panicln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Panicln(args...)
}
func (l *logger) Fatalln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo()).Fatalln(args...)
}
