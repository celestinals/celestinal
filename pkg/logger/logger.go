// Copyright 2025 The Celestinal Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logger provides the logger for the service.
package logger

import (
	"os"
	"sync"
	"time"

	"github.com/celestinals/celestinal/pkg/logger/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

var once sync.Once

var logcore = NewLogger()

var _ Logger = (*internal.Core)(nil)

// Logger define default logger for logger
type Logger interface {
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

// SystemLog define system logger, include grpclog wrapped
type SystemLog interface{ Logger }

// NewLogger creates a new logger instance.
func NewLogger() Logger {
	logger := newLogger().Sugar()
	return internal.NewCore(logger, internal.LevelDebug)
}

// NewSystemLog init all system log,
// like logger global variable, grpclog global variable
func NewSystemLog() SystemLog {
	once.Do(func() {
		logger := newLogger().Sugar()
		logcore = internal.NewCore(logger, internal.LevelDebug)

		grpclog.SetLoggerV2(internal.NewCore(logger, internal.LevelWarning))
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
