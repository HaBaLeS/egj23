GO_VERSION=1.20.3
cp /work/build/go$GO_VERSION.linux-amd64.tar.gz /tmp/
tar xf /tmp/go$GO_VERSION.linux-amd64.tar.gz

export PATH=$PATH:/tmp/go/bin/
go --version

exit
id
ls /go
cd /work
mkdir go
cd go
cp /go/go$GO_VERSION.linux-amd64.tar.gz .
echo unpack
tar xzf go$GO_VERSION.linux-amd64.tar.gz
export PATH=$PATH:/work/go/go/bin
cd /work
cd build

export CGO_CFLAGS=-std=gnu99
go build -o ../bin/steam_egj23
chown 1000:1000 ../bin/steam_egj23