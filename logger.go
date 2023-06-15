package hedera

/*-
 *
 * Hedera Go SDK
 *
 * Copyright (C) 2020 - 2022 Hedera Hashgraph, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type LogLevel string

const (
	LoggerLevelDebug    LogLevel = "DEBUG"
	LoggerLevelTrace    LogLevel = "TRACE"
	LoggerLevelInfo     LogLevel = "INFO"
	LoggerLevelDisabled LogLevel = "DISABLED"
)

type LoggerEvent struct {
	event *zerolog.Event
}

func (le *LoggerEvent) Str(key string, value string) *LoggerEvent {
	le.event = le.event.Str(key, value)
	return le
}

func (le *LoggerEvent) Msg(msg string) {
	le.event.Msg(msg)
}

func (le *LoggerEvent) Dur(key string, d time.Duration) *LoggerEvent {
	le.event = le.event.Dur(key, d)
	return le
}

func (le *LoggerEvent) Int64(key string, value int64) *LoggerEvent {
	le.event = le.event.Int64(key, value)
	return le
}

type Logger struct {
	logger *zerolog.Logger
	level  LogLevel
}

func NewLogger(component string, level LogLevel) *Logger {
	var logger zerolog.Logger
	logger = zerolog.New(os.Stdout).With().Str("module", component).Timestamp().Logger()

	if os.Getenv("HEDERA_SDK_GO_LOG_PRETTY") != "" {
		// Pretty logging
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger = zerolog.New(output).With().Str("module", component).Timestamp().Logger()
	}

	inst := &Logger{
		logger: &logger,
		level:  level,
	}

	inst.SetLevel(level)
	return inst
}

func (l *Logger) SetSilent(isSilent bool) {
	if isSilent {
		logger := l.logger.Level(zerolog.Disabled)
		l.logger = &logger
	} else {
		l.SetLevel(l.level)
	}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
	var logger zerolog.Logger
	switch level {
	case LoggerLevelDebug:
		logger = l.logger.Level(zerolog.DebugLevel)
	case LoggerLevelTrace:
		logger = l.logger.Level(zerolog.TraceLevel)
	case LoggerLevelInfo:
		logger = l.logger.Level(zerolog.InfoLevel)
	default:
		logger = l.logger.Level(zerolog.Disabled)
	}
	l.logger = &logger
}

func (l *Logger) Debug() *LoggerEvent {
	return &LoggerEvent{l.logger.Debug()}
}

func (l *Logger) Info() *LoggerEvent {
	return &LoggerEvent{l.logger.Info()}
}

func (l *Logger) Warn() *LoggerEvent {
	return &LoggerEvent{l.logger.Warn()}
}

func (l *Logger) Error() *LoggerEvent {
	return &LoggerEvent{l.logger.Error()}
}

func (l *Logger) Trace() *LoggerEvent {
	return &LoggerEvent{l.logger.Trace()}
}
