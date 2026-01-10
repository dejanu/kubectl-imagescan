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

chmod +x kubectl-imagescan
./kubectl-imagescan

# compute sha
shasum -a 256 ./kubectl-imagescan_v1.0.1_linux_amd64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.1_linux_arm64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.1_darwin_amd64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.1_darwin_arm64.tar.gz | awk '{print $1}'


# local test
KREW_OS=darwin KREW_ARCH=amd64 kubectl krew install  --manifest=krew/imagescan.yaml

KREW_OS=darwin KREW_ARCH=arm64 kubectl krew install  --manifest=krew/imagescan.yaml

KREW_OS=linux KREW_ARCH=amd64 kubectl krew install  --manifest=krew/imagescan.yaml

KREW_OS=linux KREW_ARCH=arm64 kubectl krew install  --manifest=krew/imagescan.yaml