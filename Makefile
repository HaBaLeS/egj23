VERSION=v0.0.1

.PHONY: build deck win
build:
	go build -ldflags "-X main.version=$(VERSION) -X main.buildtime=`date +%Y-%m-%d@%H:%M:%S`" -o bin/

clean:
	rm -rf bin/*

deck:
	cd build && ./build_4_deck.sh $(VERSION) && cd ..

run: build
	./bin/egj23

win:
	env GOOS=windows GOARCH=amd64   go build -ldflags "-X main.version=$(VERSION) -X main.buildtime=`date +%Y-%m-%d@%H:%M:%S`" -o bin/

upload: deck
	scp bin/steam_egj23 deck@192.168.1.87:~/Downloads/