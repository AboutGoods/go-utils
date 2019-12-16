// Package log provides global context to logrus logging library
package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var context = logrus.Fields{}
var mux = sync.Mutex{}

func Init(env string) {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	if env == "prod" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func contextLog() *logrus.Entry {
	return logrus.WithFields(context)
}

// AddToContext add elements to a global context
// It will add the key-value pair when calling the other functions available
func AddToContext(key string, value interface{}) {
	mux.Lock()
	defer mux.Unlock()
	context[key] = value
}

// RemoveFromContext remove element(s) from the global context
// It will remove on next call from the global context the key(s) passed as argument
func RemoveFromContext(keys ...string) {
	mux.Lock()
	defer mux.Unlock()
	for _, k := range keys {
		delete(context, k)
	}
}

// ResetContext will reset the global context of the logger
func ResetContext() {
	mux.Lock()
	defer mux.Unlock()
	context = logrus.Fields{}
}

// PanicOnError stops the program and logs it if the error passed as the first argument is not nil.
func PanicOnError(err error, args ...interface{}) {
	if err != nil {
		data := append([]interface{}{err}, args...)
		Fatal(data...)
	}
}

// NoticeOnError if an error occurs, logs it and return true, otherwise false.
func NoticeOnError(err error, args ...interface{}) bool {
	if err != nil {
		data := append([]interface{}{err}, args...)
		Info(data...)
		return true
	}
	return false
}

// Fatal is the global Contextualisation of logrus' eponym function
func Fatal(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Fatal(args...)
}

// Warning is the global Contextualisation of logrus' eponym function
func Warning(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Warning(args)
}

// Error is the global Contextualisation of logrus' eponym function
func Error(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Error(args...)
}

// Panic is the global Contextualisation of logrus' eponym function
func Panic(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Panic(args...)
}

// Info is the global Contextualisation of logrus' eponym function
func Info(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Info(args...)
}

// Debug is the global Contextualisation of logrus' eponym function
func Debug(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Debug(args...)
}

// Printf is the global Contextualisation of logrus' eponym function
func Printf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Printf(format, args)
}

// Trace is the global Contextualisation of logrus' eponym function
func Trace(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Trace(args...)
}

// Print is the global Contextualisation of logrus' eponym function.
func Print(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Print(args...)
}

// Warn is the global Contextualisation of logrus' eponym function.
func Warn(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Warn(args...)
}

// Tracef is the global Contextualisation of logrus' eponym function.
func Tracef(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Tracef(format, args...)
}

// Debugf is the global Contextualisation of logrus' eponym function.
func Debugf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Debugf(format, args...)
}

// Infof is the global Contextualisation of logrus' eponym function.
func Infof(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Infof(format, args...)
}

// Warnf is the global Contextualisation of logrus' eponym function.
func Warnf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Warnf(format, args...)
}

// Warningf is the global Contextualisation of logrus' eponym function.
func Warningf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Warningf(format, args...)
}

// Errorf is the global Contextualisation of logrus' eponym function.
func Errorf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Errorf(format, args...)
}

// Panicf is the global Contextualisation of logrus' eponym function.
func Panicf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Panicf(format, args...)
}

// Fatalf is the global Contextualisation of logrus' eponym function
func Fatalf(format string, args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Fatalf(format, args...)
}

// Traceln is the global Contextualisation of logrus' eponym function.
func Traceln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Traceln(args...)
}

// Debugln is the global Contextualisation of logrus' eponym function.
func Debugln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Debugln(args...)
}

// Println is the global Contextualisation of logrus' eponym function.
func Println(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Println(args...)
}

// Infoln is the global Contextualisation of logrus' eponym function.
func Infoln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Infoln(args...)
}

// Warnln is the global Contextualisation of logrus' eponym function.
func Warnln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Warnln(args...)
}

// Warningln is the global Contextualisation of logrus' eponym function.
func Warningln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Warningln(args...)
}

// Errorln is the global Contextualisation of logrus' eponym function.
func Errorln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Errorln(args...)
}

// Panicln is the global Contextualisation of logrus' eponym function.
func Panicln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Panicln(args...)
}

// Fatalln is the global Contextualisation of logrus' eponym function
func Fatalln(args ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	contextLog().Fatalln(args...)
}
