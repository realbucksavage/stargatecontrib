package middleware

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/realbucksavage/stargate"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging() stargate.MiddlewareFunc {
	return LoggingWithOutput(os.Stdout)
}

func LoggingWithOutput(dst io.Writer) stargate.MiddlewareFunc {

	logger := log.New(dst, "[stargate.requests] ", log.LstdFlags)
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			lrw := &loggingResponseWriter{ResponseWriter: rw, status: http.StatusOK}
			defer func(begin time.Time) {
				logger.Printf("[%s | %d] %s\t(%v)", r.Method, lrw.status, r.RequestURI, time.Since(begin))
			}(time.Now())

			next.ServeHTTP(lrw, r)
		})
	}
}
