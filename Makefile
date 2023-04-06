build:
	go build -o bin/

clean:
	rm -rf bin/*

run: build
	./bin/egj23
