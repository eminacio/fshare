package fshare

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Config struct {
	Port      string
	Directory string
}

func MiddlewareLogger(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		logger.Info(msg)
		next.ServeHTTP(w, r)
	})
}

func NewServer(cfg Config, logger *slog.Logger) (*http.Server, error) {
	handler := http.FileServer(http.Dir(cfg.Directory))

	mux := http.NewServeMux()
	mux.Handle("GET /", MiddlewareLogger(logger, handler))

	srv := &http.Server{
		Addr:     ":" + cfg.Port,
		Handler:  mux,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	return srv, nil
}
