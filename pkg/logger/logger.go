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

// Package logger provides the logger for the service.
package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/tickexvn/tickex/internal/version"
	"github.com/tickexvn/tickex/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new logger.
func New() *zap.Logger {
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

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}

// Info logs an info message.
func Info(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	sugar.Info(appendHeader(message...))
}

// Infof logs an info message with a format.
func Infof(template string, message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	msg := fmt.Sprintf("%s %s", version.Header(), fmt.Sprintf(template, message...))
	sugar.Info(msg)
}

// Debug logs a debug message.
func Debug(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	sugar.Debug(appendHeader(message...))
}

// Error logs an error message.
func Error(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	sugar.Error(appendHeader(message...))
}

// Errorf logs an error message with a format.
func Errorf(template string, message ...any) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	msg := fmt.Sprintf("%s %s", version.Header(), fmt.Sprintf(template, message...))
	sugar.Errorf(msg)
}

// Warnf logs an error message with a format.
func Warnf(template string, message ...any) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	msg := fmt.Sprintf("%s %s", version.Header(), fmt.Sprintf(template, message...))
	sugar.Warnf(msg)
}

// Fatal logs a fatal message.
func Fatal(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	sugar.Fatal(appendHeader(message...))
}

// Fatalf logs a fatal message.
func Fatalf(template string, message ...any) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	msg := fmt.Sprintf("%s %s", version.Header(), fmt.Sprintf(template, message...))
	sugar.Fatalf(msg)
}

func removeNil(input []interface{}) []interface{} {
	var result []interface{}
	for _, item := range input {
		if item != nil {
			result = append(result, item)
		}
	}
	return result
}

func appendHeader(message ...any) string {
	msg := fmt.Sprint(message...)
	return fmt.Sprintf("%s %s", version.Header(), msg)
}
