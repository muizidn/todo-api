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

func _CallerInfo(level int) logrus.Fields {
	_, file, line, ok := runtime.Caller(level)
	if ok {
		comp := strings.Split(file, "/")
		return logrus.Fields{
			"file": comp[len(comp)-1],
			"line": line,
		}
	}
	return nil
}

type Fields = logrus.Fields

func (l *logger) WithFields(fields Fields) *logger {
	logrusFields := logrus.Fields{}
	for k, v := range fields {
		logrusFields[k] = v
	}
	newLogger := &logger{entry: l.entry.WithFields(logrusFields)}
	return newLogger
}

/// ToDo log interface
func (l *logger) TError(err error, info ...interface{}) error {
	l.WithFields(Fields{
		"err":      err,
		"act_file": _CallerInfo(2),
	}).Error(info...)
	return err
}

/// logrus interface

func (l *logger) Trace(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Trace(args...)
}
func (l *logger) Debug(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Debug(args...)
}
func (l *logger) Print(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Print(args...)
}
func (l *logger) Info(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Info(args...)
}
func (l *logger) Warn(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Warn(args...)
}
func (l *logger) Warning(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Warning(args...)
}
func (l *logger) Error(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Error(args...)
}
func (l *logger) Panic(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Panic(args...)
}
func (l *logger) Fatal(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Fatal(args...)
}
func (l *logger) Tracef(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Tracef(format, args...)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Debugf(format, args...)
}
func (l *logger) Printf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Printf(format, args...)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Infof(format, args...)
}
func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Warnf(format, args...)
}
func (l *logger) Warningf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Warningf(format, args...)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Errorf(format, args...)
}
func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Panicf(format, args...)
}
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Fatalf(format, args...)
}
func (l *logger) Traceln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Traceln(args...)
}
func (l *logger) Debugln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Debugln(args...)
}
func (l *logger) Println(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Println(args...)
}
func (l *logger) Infoln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Infoln(args...)
}
func (l *logger) Warnln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Warnln(args...)
}
func (l *logger) Warningln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Warningln(args...)
}
func (l *logger) Errorln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Errorln(args...)
}
func (l *logger) Panicln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Panicln(args...)
}
func (l *logger) Fatalln(args ...interface{}) {
	l.entry.WithFields(_CallerInfo(2)).Fatalln(args...)
}
