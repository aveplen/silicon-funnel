package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type LivenessProxy struct {
	h        http.Handler
	handlers map[string]http.HandlerFunc
}

func LivenessProbe(h http.Handler) http.Handler {
	logrus.Infof("mounted liveness probe")

	started := time.Now()
	proxy := &LivenessProxy{
		h: h,
		handlers: map[string]http.HandlerFunc{
			"/started": startedHandler(&started),
			"/healthz": healthzHandler(&started),
		},
	}

	return proxy
}

func (l *LivenessProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	handler, ok := l.handlers[r.URL.Path]
	if ok {
		handler(w, r)
		return
	}

	l.h.ServeHTTP(w, r)
}

func startedHandler(started *time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)
		data := (time.Since(*started)).String()
		w.Write([]byte(data))
	}
}

// TODO: reimplement to be real health check
func healthzHandler(started *time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		duration := time.Since(*started)
		if duration.Seconds() < 3 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	}
}
