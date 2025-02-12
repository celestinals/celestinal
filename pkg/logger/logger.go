/*
 * Copyright 2024 The Tickex Authors.
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
	"github.com/tickexvn/tickex/pkg/utils"
	"go.uber.org/zap"
)

// New creates a new logger.
func New() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

// Info logs an info message.
func Info(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Info(message...)
}

// Infof logs an info message with a format.
func Infof(template string, message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Infof(template, message...)
}

// Debug logs a debug message.
func Debug(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Debug(message...)
}

// Error logs an error message.
func Error(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Error(message...)
}

// Fatal logs a fatal message.
func Fatal(message ...interface{}) {
	if len(removeNil(message)) == 0 {
		return
	}

	logger := New()
	defer utils.CallBack(logger.Sync)

	sugar := logger.Sugar()
	sugar.Fatal(message...)
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
