package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	clientset *kubernetes.Clientset
}

// NewClient creates a new Kubernetes client using the default kubeconfig
func NewClient() (*Client, error) {
	// Use the current context from kubeconfig
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	return &Client{clientset: clientset}, nil
}

// ValidateNamespace checks if a namespace exists
func (c *Client) ValidateNamespace(namespace string) error {
	_, err := c.clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("namespace %s does not exist", namespace)
	}
	return nil
}

// ListNamespaces returns a list of all namespaces
func (c *Client) ListNamespaces() ([]string, error) {
	nsList, err := c.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	namespaces := make([]string, len(nsList.Items))
	for i, ns := range nsList.Items {
		namespaces[i] = ns.Name
	}
	return namespaces, nil
}

// ListPods returns a list of all pods in a namespace
func (c *Client) ListPods(namespace string) ([]string, error) {
	podList, err := c.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pods := make([]string, len(podList.Items))
	for i, pod := range podList.Items {
		pods[i] = pod.Name
	}
	return pods, nil
}

// GetPodImages returns all container images used by a pod
func (c *Client) GetPodImages(namespace, podName string) ([]string, error) {
	pod, err := c.clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod: %w", err)
	}

	var images []string
	for _, container := range pod.Spec.Containers {
		images = append(images, container.Image)
	}

	return images, nil
}
