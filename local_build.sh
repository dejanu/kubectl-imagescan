#!/usr/bin/env bash

# set vars
VERSION=v1.0.1
BIN=kubectl-imagescan
# build flags: strip symbols to reduce size
LDFLAGS="-s -w"

# clean old artifacts
rm -f ./*_${VERSION}_*.tar.gz || true

# macOS arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
go build -ldflags="$LDFLAGS" -o ${BIN} .

tar czf ${BIN}_${VERSION}_darwin_arm64.tar.gz ${BIN}
rm ${BIN}

# macOS amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
go build -ldflags="$LDFLAGS" -o ${BIN} .

tar czf ${BIN}_${VERSION}_darwin_amd64.tar.gz ${BIN}
rm ${BIN}

# Linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
go build -ldflags="$LDFLAGS" -o ${BIN} .

tar czf ${BIN}_${VERSION}_linux_arm64.tar.gz ${BIN}
rm ${BIN}

# Linux amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -ldflags="$LDFLAGS" -o ${BIN} 

tar czf ${BIN}_${VERSION}_linux_amd64.tar.gz ${BIN}
rm ${BIN}