package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/agungdwiprasetyo/reverse-proxy/helper"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader implement http.ResponseWriter
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger for log all request in this gateway
func Logger(wrap http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}
		wrap.ServeHTTP(rw, r)

		end := time.Now()
		statusColor := helper.ColorForStatus(rw.statusCode)
		methodColor := helper.ColorForMethod(r.Method)

		fmt.Fprintf(os.Stdout, "%s[GATEWAY]%s : %v | %s %3d %s | %13v | %15s | %s %-7s %s %s\n",
			helper.White, helper.Reset,
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusColor, rw.statusCode, helper.Reset,
			end.Sub(start),
			r.RemoteAddr,
			methodColor, r.Method, helper.Reset,
			r.RequestURI,
		)
	}
}
