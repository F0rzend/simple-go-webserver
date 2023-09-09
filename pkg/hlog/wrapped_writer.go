package hlog

import "net/http"

type WrappedWriter struct {
	http.ResponseWriter

	wroteHeader bool

	statusCode int
	body       []byte
}

func wrapWriter(w http.ResponseWriter) *WrappedWriter {
	return &WrappedWriter{ResponseWriter: w}
}

func (ww *WrappedWriter) WriteHeader(statusCode int) {
	if ww.wroteHeader {
		return
	}

	ww.statusCode = statusCode
	ww.wroteHeader = true
	ww.ResponseWriter.WriteHeader(statusCode)
}

func (ww *WrappedWriter) Write(b []byte) (int, error) {
	if !ww.wroteHeader {
		ww.WriteHeader(http.StatusOK)
	}

	ww.body = append(ww.body, b...)

	return ww.ResponseWriter.Write(b)
}

func (ww *WrappedWriter) StatusCode() int {
	return ww.statusCode
}

func (ww *WrappedWriter) Body() []byte {
	return ww.body
}
