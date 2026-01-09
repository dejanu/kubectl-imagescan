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

# Then build
go build -o kubectl-imagescan .
chmod +x kubectl-imagescan
./kubectl-imagescan