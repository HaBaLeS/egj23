GO_VERSION=1.20.3
STEAM_RT=registry.gitlab.steamos.cloud/steamrt/scout/sdk:latest

if test -f "go$GO_VERSION.linux-amd64.tar.gz"; then
  echo "No downloading GO"
else
  wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz
fi

WORK_DIR=$(pwd)/../

docker run --rm  -it --volume $WORK_DIR:/work   --user $(id -u):$(id -g)  $STEAM_RT  /bin/bash -c "
GO_VERSION=$GO_VERSION
cd /tmp
tar xf /work/build/go$GO_VERSION.linux-amd64.tar.gz
cd /work
export PATH=$PATH:/tmp/go/bin/
export GOCACHE=/tmp/
export GOPATH=/tmp/go
export CGO_CFLAGS=-std=gnu99
go env
go build -o bin/steam_egj23
"