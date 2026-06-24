package fshare_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/eminacio/fshare"
)

// Integration test

func TestServer(t *testing.T) {
	dir := t.TempDir()

	os.WriteFile(
		filepath.Join(dir, "index.html"),
		[]byte("hi"),
		0644,
	)

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	srv, _ := fshare.NewServer(
		fshare.Config{
			Port:      "0",
			Directory: dir,
		},
		logger,
	)

	tsrv := httptest.NewServer(srv.Handler)
	defer tsrv.Close()

	resp, err := http.Get(tsrv.URL + "/index.html")
	if err != nil {
		t.Fatal(err)
	}

	body, _ := io.ReadAll(resp.Body)
	if got, want := string(body), "hi"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
