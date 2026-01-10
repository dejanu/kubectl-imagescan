#!/usr/bin/env bash

# set vars
VERSION=v1.0.1
BIN=kubectl-imagescan
# build flags: strip symbols to reduce size
LDFLAGS="-s -w"

# clean old artifacts
rm -f ./*-${VERSION}-*.tar.gz || true

# macOS arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
go build -ldflags="$LDFLAGS" -o ${BIN} .

tar czf ${BIN}-${VERSION}-darwin-arm64.tar.gz ${BIN}
rm ${BIN}

# macOS amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
go build -ldflags="$LDFLAGS" -o ${BIN} .

tar czf ${BIN}-${VERSION}-darwin-amd64.tar.gz ${BIN}
rm ${BIN}

# Linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
go build -ldflags="$LDFLAGS" -o ${BIN} .

tar czf ${BIN}-${VERSION}-linux-arm64.tar.gz ${BIN}
rm ${BIN}

# Linux amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -ldflags="$LDFLAGS" -o ${BIN} 

tar czf ${BIN}-${VERSION}-linux-amd64.tar.gz ${BIN}
rm ${BIN}

# compute sha for each artifact
for file in ./*-${VERSION}-*.tar.gz; do
  shasum -a 256 $file 
done