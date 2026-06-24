package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/eminacio/fshare"
)

func main() {
	var port string
	var dir string
	flag.StringVar(&port, "port", "8000", "server port")
	flag.StringVar(&dir, "directory", "", "path to your folder")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	wd, err := os.Getwd()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	cfg := fshare.Config{
		Port:      port,
		Directory: filepath.Join(wd, dir),
	}

	srv, err := fshare.NewServer(cfg, logger)

	msg := fmt.Sprintf("Serving HTTP on :: port %s", cfg.Port)
	logger.Info(msg)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
