package mw

import (
	"github.com/getsentry/sentry-go"
	"github.com/hell-kitchen/api-gateway/internal/config"
)

func StartUpSentry(cfg *config.Server) bool {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDNS,
		TracesSampleRate: 1.0,
	})
	return err == nil
}
