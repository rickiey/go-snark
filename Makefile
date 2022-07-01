
all: go-snark window-snark

env:
	git submodule update --init --recursive; cd extern/filecoin-ffi; make all; cd  ../..; go mod tidy; mkdir bin;

go-snark: env
	go build -o bin/go-snark cmd/snark-server/server.go

window-snark: env
	go build -o bin/window-snark cmd/windowpost/*.go


clean:
	rm -rf bin