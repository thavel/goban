package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger       *zap.SugaredLogger
	level        = zap.NewAtomicLevel()
	levelMapping = map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
)

func init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	l := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		level,
	))
	defer l.Sync()
	logger = l.Sugar()
}

// SetLevel sets the log level.
func SetLevel(l string) {
	v, ok := levelMapping[l]
	if !ok {
		v = levelMapping["info"]
	}
	level.SetLevel(v)
}

// Debug wrapper.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Error wrapper.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal wrapper.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Info wrapper.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Panic wrapper.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Warn wrapper.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Debugf wrapper.
func Debugf(t string, args ...interface{}) {
	logger.Debugf(t, args...)
}

// Errorf wrapper.
func Errorf(t string, args ...interface{}) {
	logger.Errorf(t, args...)
}

// Fatalf wrapper.
func Fatalf(t string, args ...interface{}) {
	logger.Fatalf(t, args...)
}

// Infof wrapper.
func Infof(t string, args ...interface{}) {
	logger.Infof(t, args...)
}

// Panicf wrapper.
func Panicf(t string, args ...interface{}) {
	logger.Panicf(t, args...)
}

// Warnf wrapper.
func Warnf(t string, args ...interface{}) {
	logger.Warnf(t, args...)
}
