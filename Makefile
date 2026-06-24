create_bin_dir:
	mkdir -p bin
build: create_bin_dir
	go build -o bin/fshare ./cmd/fshare