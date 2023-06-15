package hedera

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger("test", LoggerLevelDebug)
	require.NotNil(t, logger)
	assert.Equal(t, LoggerLevelDebug, logger.level)
}

func TestLogger_SetLevel(t *testing.T) {
	logger := NewLogger("test", LoggerLevelDebug)
	logger.SetLevel(LoggerLevelInfo)
	assert.Equal(t, LoggerLevelInfo, logger.level)
}

func TestLogger_SetSilent(t *testing.T) {
	var buf bytes.Buffer
	writer := zerolog.ConsoleWriter{Out: &buf, NoColor: true, TimeFormat: time.RFC3339}

	logger := zerolog.New(writer).With().Str("module", "test").Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	l := &Logger{
		logger: &logger,
		level:  LoggerLevelTrace,
	}

	l.SetLevel(LoggerLevelTrace)
	l.Debug().Msg("debug message")
	l.Info().Msg("info message")
	l.Warn().Msg("warn message")
	l.Error().Msg("error message")
	l.Trace().Msg("trace message")
	assert.Contains(t, buf.String(), "debug message")
	assert.Contains(t, buf.String(), "info message")
	assert.Contains(t, buf.String(), "warn message")
	assert.Contains(t, buf.String(), "error message")
	assert.Contains(t, buf.String(), "trace message")

	buf.Reset()

	l.SetSilent(true)
	l.Debug().Msg("silent debug message")
	l.Info().Msg("silent info message")
	l.Warn().Msg("silent warn message")
	l.Error().Msg("silent error message")
	l.Trace().Msg("silent trace message")
	assert.NotContains(t, buf.String(), "silent debug message")
	assert.NotContains(t, buf.String(), "silent info message")
	assert.NotContains(t, buf.String(), "silent warn message")
	assert.NotContains(t, buf.String(), "silent error message")
	assert.NotContains(t, buf.String(), "silent trace message")

	buf.Reset()

	l.SetSilent(false)
	l.Debug().Msg("debug message")
	l.Info().Msg("info message")
	l.Warn().Msg("warn message")
	l.Error().Msg("error message")
	l.Trace().Msg("trace message")
	assert.Contains(t, buf.String(), "debug message")
	assert.Contains(t, buf.String(), "info message")
	assert.Contains(t, buf.String(), "warn message")
	assert.Contains(t, buf.String(), "error message")
	assert.Contains(t, buf.String(), "trace message")
}
func TestLoggerEvent_Str(t *testing.T) {
	logger := NewLogger("test", LoggerLevelDebug)
	event := logger.Debug().Str("key", "value")
	require.NotNil(t, event)
}

func TestLoggerEvent_Msg(t *testing.T) {
	logger := NewLogger("test", LoggerLevelDebug)
	event := logger.Debug()
	event.Msg("test message")
}

func TestLoggerEvent_Dur(t *testing.T) {
	logger := NewLogger("test", LoggerLevelDebug)
	event := logger.Debug().Dur("key", time.Duration(100))
	require.NotNil(t, event)
}

func TestLoggerEvent_Int64(t *testing.T) {
	logger := NewLogger("test", LoggerLevelDebug)
	event := logger.Debug().Int64("key", 100)
	require.NotNil(t, event)
}

func TestNewLoggerWithEnvironmentVariableSet(t *testing.T) {
	os.Setenv("HEDERA_SDK_GO_LOG_PRETTY", "1")
	logger := NewLogger("test", LoggerLevelDebug)
	require.NotNil(t, logger)
	assert.Equal(t, LoggerLevelDebug, logger.level)
	os.Unsetenv("HEDERA_SDK_GO_LOG_PRETTY")
}
