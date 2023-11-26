package monitoring

import (
	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
)

func ErrLoggingWithSentry(err error) {
	defer sentry.Flush(sentryTimeout)
	sentry.CaptureException(err)
	slog.Error(
		err.Error(),
		slog.Any("trace", errors.GetAllSafeDetails(err)),
		slog.Any("hints", errors.GetAllHints(err)),
	)
}
