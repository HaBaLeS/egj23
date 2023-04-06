VERSION=v0.0.1

.PHONY: build deck
build:
	go build -ldflags "-X main.version=$(VERSION) -X main.buildtime=`date +%Y-%m-%d@%H:%M:%S`" -o bin/

clean:
	rm -rf bin/*

deck:
	cd build && ./build_4_deck.sh $(VERSION) && cd ..

run: build
	./bin/egj23
