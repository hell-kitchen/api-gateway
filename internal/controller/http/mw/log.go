package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"sync"
)

var (
	logConfig = middleware.RequestLoggerConfig{
		LogLatency:    true,
		LogRemoteIP:   true,
		LogHost:       true,
		LogMethod:     true,
		LogURI:        true,
		LogRequestID:  true,
		LogStatus:     true,
		LogError:      true,
		LogValuesFunc: logValuesFunc,
	}
	_loggerMu sync.RWMutex
	_logger   = zap.L()
)

// LogMiddleware is middleware which logs all necessary data.
func LogMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	_loggerMu.Lock()
	_logger = logger
	_loggerMu.Unlock()
	return middleware.RequestLoggerWithConfig(logConfig)
}

func logValuesFunc(_ echo.Context, v middleware.RequestLoggerValues) error {
	_loggerMu.RLock()
	_logger.Info(
		"request",
		zap.String("uri", v.URI),
		zap.Duration("latency", v.Latency),
		zap.String("remote-ip", v.RemoteIP),
		zap.String("host", v.Host),
		zap.String("method", v.Method),
		zap.String("request-id", v.RequestID),
		zap.Int("status", v.Status),
		zap.Error(v.Error),
	)
	_loggerMu.RUnlock()
	return nil
}
