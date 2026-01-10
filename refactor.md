Preserve original feature but separated concerns (k8s client, scanner, CLI)

# Initialize Go module
go mod init github.com/dejanu/kubectl-imagescan

# Create main.go file
touch main.go

# Create project structure
# CLI logic isolated in cmd/
# Business logic in pkg/
mkdir -p cmd/kubectl-imagescan && touch cmd/kubectl-imagescan/root.go
mkdir -p pkg/{scanner,k8s} && touch pkg/scanner/trivy.go && touch pkg/k8s/client.go

kubectl-imagescan/
├── main.go                 # Entry point
├── cmd/
│   └── kubectl-imagescan/
│       └── root.go        # CLI commands
├── pkg/
│   ├── scanner/
│   │   └── trivy.go       # Trivy scanning logic
│   └── k8s/
│       └── client.go      # Kubernetes client logic
├── go.mod
├── go.sum
├── README.md
└── krew/
    └── imagescan.yaml

---- 

Core:
github.com/spf13/cobra - CLI framework
github.com/AlecAivazis/survey/v2 - Interactive prompts
k8s.io/client-go - Kubernetes API client
k8s.io/apimachinery - Kubernetes types

# Add Kubernetes client-go
go get k8s.io/client-go@latest
go get k8s.io/apimachinery@latest

# Add CLI framework (cobra for better CLI handling)
go get github.com/spf13/cobra@latest

# Add survey for interactive prompts
go get github.com/AlecAivazis/survey/v2@latest

# Clean up and re-download everything
go mod tidy
go mod download

* BUILD

```bash
# set vars
VERSION=v1.0.0
BIN=kubectl-imagescan

# build flags: strip symbols to reduce size
LDFLAGS="-s -w"

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

chmod +x kubectl-imagescan
./kubectl-imagescan

# compute sha
shasum -a 256 ./kubectl-imagescan_v1.0.0_linux_amd64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.0_linux_arm64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.0_darwin_amd64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.0_darwin_arm64.tar.gz | awk '{print $1}'
```