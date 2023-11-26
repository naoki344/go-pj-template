package monitoring

import (
	"log/slog"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

const sentryTimeout = 2 * time.Second

func SetupSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("EN_SENTRY_DSN"),
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		slog.Error("sentry.Init: %s", err)
	}
}
