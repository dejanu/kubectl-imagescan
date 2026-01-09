package scanner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type TrivyScanner struct {
	cacheDir string
}

// NewTrivyScanner creates a new Trivy scanner instance
func NewTrivyScanner() *TrivyScanner {
	homeDir, _ := os.UserHomeDir()
	cacheDir := filepath.Join(homeDir, ".cache", "trivy")
	return &TrivyScanner{
		cacheDir: cacheDir,
	}
}

// CheckDockerRunning verifies if Docker daemon is running
func CheckDockerRunning() error {
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker daemon is not running")
	}
	return nil
}

// UpdateDatabase updates the Trivy vulnerability database
func (t *TrivyScanner) UpdateDatabase() error {
	// Create cache directory
	if err := os.MkdirAll(t.cacheDir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/root/.cache/", t.cacheDir),
		"aquasec/trivy:latest",
		"image", "--download-db-only")

	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update Trivy database: %w", err)
	}

	return nil
}

// ScanImage scans a container image using Trivy
func (t *TrivyScanner) ScanImage(image string) error {
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/root/.cache/", t.cacheDir),
		"aquasec/trivy:latest",
		"image",
		"--skip-db-update",
		"--severity", "HIGH,CRITICAL",
		"--format", "table",
		"--quiet",
		"--timeout", "10m",
		image)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	return nil
}
