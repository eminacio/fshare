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

func main() {
	type Config struct {
		Port      string
		Directory string
	}
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

	http.Handle("GET /", MiddlewareLogger(logger, handler))

	msg := fmt.Sprintf("Serving HTTP on :: port %s", cfg.Port)
	logger.Info(msg)

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
}
