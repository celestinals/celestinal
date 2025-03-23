/*
 * Copyright 2025 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package txlog provides the logger for the service.
package txlog

import (
	"os"
	"sync"
	"time"

	"github.com/tickexvn/tickex/pkg/txlog/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

var once sync.Once

var logcore = NewTxLog()

var _ TxLog = (*internal.TxLogCore)(nil)

// TxLog define default logger for txlog
type TxLog interface {
	Info(args ...any)
	Infof(template string, args ...any)
	Infoln(args ...any)

	Debug(args ...any)
	Debugf(template string, args ...any)
	Debugln(args ...any)

	Warning(args ...any)
	Warningf(template string, args ...any)
	Warningln(args ...any)

	Error(args ...any)
	Errorf(template string, args ...any)
	Errorln(args ...any)

	Fatal(args ...any)
	Fatalf(template string, args ...any)
	Fatalln(args ...any)

	V(l int) bool
	Sync() error
}

// TxSystemLog define system logger, include grpclog wrapped
type TxSystemLog interface{ TxLog }

// NewTxLog creates a new txlog instance.
func NewTxLog() TxLog {
	logger := newLogger().Sugar()
	return internal.NewTxLogCore(logger, internal.LevelDebug)
}

// NewTxSystemLog init all system log,
// like txlog global variable, grpclog global variable
func NewTxSystemLog() TxSystemLog {
	once.Do(func() {
		logger := newLogger().Sugar()
		logcore = internal.NewTxLogCore(logger, internal.LevelDebug)

		grpclog.SetLoggerV2(internal.NewTxLogCore(logger, internal.LevelWarning))
	})

	return logcore
}

// newLogger creates a new logger.
func newLogger() *zap.Logger {
	config := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(time.DateTime))
		},
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Console output
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger
}

// Info logs an info message.
func Info(message ...any) {
	logcore.Info(message...)
}

// Infof logs an info message with a format.
func Infof(template string, message ...any) {
	logcore.Infof(template, message...)
}

// Debug logs a debug message.
func Debug(message ...any) {
	logcore.Debug(message...)
}

// Debugf logs a debug message.
func Debugf(template string, message ...any) {
	logcore.Debugf(template, message...)
}

// Error logs an error message.
func Error(message ...any) {
	logcore.Error(message...)
}

// Errorf logs an error message with a format.
func Errorf(template string, message ...any) {
	logcore.Errorf(template, message...)
}

// Warn logs an warn message.
func Warn(message ...any) {
	logcore.Warning(message...)
}

// Warnf logs an error message with a format.
func Warnf(template string, message ...any) {
	logcore.Warningf(template, message...)
}

// Fatal logs a fatal message.
func Fatal(message ...any) {
	logcore.Fatal(message...)
}

// Fatalf logs a fatal message.
func Fatalf(template string, message ...any) {
	logcore.Fatalf(template, message...)
}
