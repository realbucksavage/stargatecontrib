package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/quipo/statsd"
	"github.com/realbucksavage/stargate"
)

const (
	statsdStatusCodeFmt   = "http_status_%d"
	statsdResponseTimeTag = "response_times"
)

func StatsdMiddleware(client *statsd.StatsdClient) stargate.MiddlewareFunc {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			lrw := &loggingResponseWriter{ResponseWriter: rw, status: http.StatusOK}
			defer func(begin time.Time) {

				if err := client.Timing(statsdResponseTimeTag, time.Since(begin).Milliseconds()); err != nil {
					stargate.Log.Error("%s not recorded to statsd: %v", statsdResponseTimeTag, err)
				}

				if err := client.Incr(fmt.Sprintf(statsdStatusCodeFmt, lrw.status), int64(1)); err != nil {
					stargate.Log.Error("increment of http satus %d not recorded t statsd: %v", lrw.status, err)
				}

			}(time.Now())

			next.ServeHTTP(rw, r)
		})
	}
}
