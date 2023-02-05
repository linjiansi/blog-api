package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{w, http.StatusOK}
}

func (rsw *resLoggingWriter) WriterHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI, req.Method)
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Println("res: ", rlw.code)
	}
	return http.HandlerFunc(fn)
}
