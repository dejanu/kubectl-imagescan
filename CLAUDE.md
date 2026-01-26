# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

kubectl-imagescan is a kubectl plugin that scans container images used by Kubernetes Pods for vulnerabilities using Trivy. It provides an interactive, non-invasive way to inspect container images in a cluster without modifying workloads.

## Build and Development Commands

### Local Development
```bash
# Build for local testing
go build -o kubectl-imagescan .
chmod +x kubectl-imagescan

# Test locally (without installing)
./kubectl-imagescan <namespace> <pod-name>

# Install as kubectl plugin (add to PATH)
export PATH="$PATH:$(pwd)"
kubectl imagescan <namespace> <pod-name>
```

### Testing with Krew (local)
```bash
# Test krew installation locally for your platform
kubectl krew install --manifest=krew/imagescan.yaml

# Test cross-platform installations
KREW_OS=darwin KREW_ARCH=arm64 kubectl krew install --manifest=krew/imagescan.yaml
KREW_OS=linux KREW_ARCH=amd64 kubectl krew install --manifest=krew/imagescan.yaml
```

### Building Release Artifacts
```bash
# Build multi-platform binaries and compute SHA256 checksums
./local_build.sh

# Manually compute SHA for a tarball
shasum -a 256 kubectl-imagescan.tar.gz
```

The build script (`local_build.sh`) creates binaries for:
- darwin/amd64, darwin/arm64
- linux/amd64, linux/arm64

Build flags include `-s -w` to strip symbols and reduce binary size.

### Dependencies
```bash
# Install/update dependencies
go mod tidy
go mod download
```

## Architecture

The codebase follows a clean separation of concerns with three main layers:

### 1. Entry Point (`main.go`)
Simple entry point that delegates to the CLI package and handles top-level errors.

### 2. CLI Layer (`cmd/kubectl-imagescan/root.go`)
- **Framework**: Uses `spf13/cobra` for CLI structure
- **Interactive prompts**: Uses `AlecAivazis/survey/v2` for user confirmations
- **Flow**:
  1. Validate namespace exists (lists available if invalid)
  2. Validate pod exists in namespace (lists available if invalid)
  3. Check Docker daemon is running
  4. Extract container images from pod spec
  5. Update Trivy database
  6. Prompt user to confirm scan for each image
  7. Scan images that user approves

### 3. Business Logic (`pkg/`)

**`pkg/k8s/client.go`**: Kubernetes operations
- Creates clientset from default kubeconfig (respects current context)
- Methods: `ValidateNamespace()`, `ListNamespaces()`, `ListPods()`, `GetPodImages()`
- Extracts images from `pod.Spec.Containers[].Image`

**`pkg/scanner/trivy.go`**: Trivy scanning operations
- Runs Trivy via Docker (uses `aquasec/trivy:latest` image)
- Cache directory: `~/.cache/trivy`
- `CheckDockerRunning()`: Validates Docker daemon is accessible
- `UpdateDatabase()`: Downloads vulnerability DB before scanning
- `ScanImage()`: Scans with severity filtering (HIGH, CRITICAL only)
- Scan parameters: `--skip-db-update`, `--timeout 10m`, `--format table`, `--quiet`

## Key Design Decisions

### Trivy Execution Model
The scanner runs Trivy as a Docker container (not a local binary), which:
- Ensures consistent Trivy version across environments
- Requires Docker daemon to be running
- Mounts `~/.cache/trivy` for persistent vulnerability database

### Interactive UX
The tool prompts before scanning each image to give users control over:
- Which images to scan (some pods may have many containers)
- Time commitment (scans can be slow for large images)

### Validation-First Approach
The CLI validates namespace and pod existence before attempting scans, providing helpful error messages with available options rather than cryptic API errors.

## Krew Plugin Packaging

The `krew/imagescan.yaml` manifest defines the kubectl plugin for distribution via Krew:
- Version managed separately from code version (currently v1.0.1)
- Requires SHA256 checksums for each platform tarball
- When updating version: rebuild all platform binaries, compute new SHAs, update manifest

## Current Version

Code version: `2.0.0` (defined in `cmd/kubectl-imagescan/root.go:13`)
Krew version: `v1.0.1` (defined in `krew/imagescan.yaml` and `local_build.sh`)

Note: Version numbers are currently out of sync between code and release artifacts.
