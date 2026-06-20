package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func MiddlewareLogger(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		logger.Info(msg)
		next.ServeHTTP(w, r)
	})
}

type Config struct {
	Port      string
	Directory string
}

func main() {
	var cfg Config

	flag.StringVar(&cfg.Port, "port", "8000", "server port")
	flag.StringVar(&cfg.Directory, "directory", "", "path to your folder")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	wd, err := os.Getwd()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	dirpath := filepath.Join(wd, cfg.Directory)
	handler := http.FileServer(http.Dir(dirpath))

	mux := http.NewServeMux()
	mux.Handle("GET /", MiddlewareLogger(logger, handler))

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%s", cfg.Port),
		Handler:  mux,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info(fmt.Sprintf("Serving HTTP on :: port %s", cfg.Port))

	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
