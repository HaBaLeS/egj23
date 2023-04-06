VERSION=0.0.1

.PHONY: build
build:
	go build -ldflags "-X main.version=$VERSION -X main.buildtime=`date +%Y-%m-%d@%H:%M:%S`" -o bin/

clean:
	rm -rf bin/*

run: build
	./bin/egj23
