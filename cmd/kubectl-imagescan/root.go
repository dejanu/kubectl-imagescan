package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dejanu/kubectl-imagescan/pkg/k8s"
	"github.com/dejanu/kubectl-imagescan/pkg/scanner"
	"github.com/spf13/cobra"
)

const version = "2.0.0"

var rootCmd = &cobra.Command{
	Use:   "kubectl-imagescan <namespace> <pod>",
	Short: "Scan container images used by Pods",
	Long: `kubectl-imagescan inspects container images referenced by Pods
and scans them for vulnerabilities using Trivy.`,
	Version: version,
	Args:    cobra.RangeArgs(0, 2),
	RunE:    run,
}

func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) error {
	// Create Kubernetes client first
	k8sClient, err := k8s.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	// Handle case: no arguments provided
	if len(args) == 0 {
		fmt.Println("Please provide a valid namespace as the first argument")
		fmt.Println("\nAvailable namespaces:")
		namespaces, err := k8sClient.ListNamespaces()
		if err != nil {
			return fmt.Errorf("failed to list namespaces: %w", err)
		}
		for _, ns := range namespaces {
			fmt.Printf("  - %s\n", ns)
		}
		return fmt.Errorf("missing required arguments")
	}

	namespace := args[0]

	// Validate namespace exists
	if err := k8sClient.ValidateNamespace(namespace); err != nil {
		fmt.Printf("Namespace '%s' does not exist\n", namespace)
		fmt.Println("\nAvailable namespaces:")
		namespaces, _ := k8sClient.ListNamespaces()
		for _, ns := range namespaces {
			fmt.Printf("  - %s\n", ns)
		}
		return fmt.Errorf("namespace validation failed")
	}

	// Handle case: only namespace provided (missing pod name)
	if len(args) == 1 {
		fmt.Println("Please provide the pod name as the second argument")
		fmt.Printf("\nAvailable pods in namespace '%s':\n", namespace)
		pods, err := k8sClient.ListPods(namespace)
		if err != nil {
			return fmt.Errorf("failed to list pods: %w", err)
		}
		for _, pod := range pods {
			fmt.Printf("  - %s\n", pod)
		}
		return fmt.Errorf("missing pod name argument")
	}

	podName := args[1]

	// Check if Docker is running
	if err := scanner.CheckDockerRunning(); err != nil {
		return fmt.Errorf("Docker daemon is not running: %w", err)
	}

	// Get pod images
	images, err := k8sClient.GetPodImages(namespace, podName)
	if err != nil {
		fmt.Printf("Failed to get pod '%s' in namespace '%s'\n", podName, namespace)
		fmt.Printf("\nAvailable pods in namespace '%s':\n", namespace)
		pods, _ := k8sClient.ListPods(namespace)
		for _, pod := range pods {
			fmt.Printf("  - %s\n", pod)
		}
		return fmt.Errorf("failed to get pod images: %w", err)
	}

	if len(images) == 0 {
		return fmt.Errorf("no images found in pod %s/%s", namespace, podName)
	}

	fmt.Printf("\nImages in %s/%s:\n", namespace, podName)
	for _, img := range images {
		fmt.Printf("  - %s\n", img)
	}
	fmt.Println()

	// Initialize scanner
	trivyScanner := scanner.NewTrivyScanner()

	// Update Trivy database
	fmt.Println("📦 Updating vulnerability database...")
	if err := trivyScanner.UpdateDatabase(); err != nil {
		return fmt.Errorf("failed to update Trivy database: %w", err)
	}

	// Scan each image
	for _, image := range images {
		// Prompt user
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Scan %s?", image),
			Default: true,
		}
		var scan bool
		if err := survey.AskOne(prompt, &scan); err != nil {
			return err
		}

		if scan {
			fmt.Printf("\n🔍 Scanning %s...\n\n", image)
			if err := trivyScanner.ScanImage(image); err != nil {
				fmt.Fprintf(os.Stderr, "Error scanning %s: %v\n", image, err)
			}
		} else {
			fmt.Printf("\n⏭️  Skipping %s...\n\n", image)
		}
	}

	return nil
}
