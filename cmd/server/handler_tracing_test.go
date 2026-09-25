package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ory/x/reqlog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetTracingLogger(t *testing.T) {
	testCases := []struct {
		name        string
		traceparent string
		traceID     string
		spanID      string
	}{
		{
			name:        "valid trace context",
			traceparent: "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01",
			traceID:     "4bf92f3577b34da6a3ce929d0e0e4736",
			spanID:      "00f067aa0ba902b7",
		},
		{
			name:        "invalid trace context",
			traceparent: "invalid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := reqlog.NewMiddleware()
			logger.Before = func(entry *logrus.Entry, _ *http.Request, _ string) *logrus.Entry {
				return entry.WithField("existing", "value")
			}
			setTracingLogger(logger)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request.Header.Set("traceparent", tc.traceparent)
			entry := logger.Before(logrus.NewEntry(logrus.New()), request, "")

			assert.Equal(t, "value", entry.Data["existing"])
			if tc.traceID == "" {
				assert.NotContains(t, entry.Data, "trace_id")
				assert.NotContains(t, entry.Data, "span_id")
			} else {
				assert.Equal(t, tc.traceID, entry.Data["trace_id"])
				assert.Equal(t, tc.spanID, entry.Data["span_id"])
			}
		})
	}
}
