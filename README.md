# kubectl-imagescan

Scan container images used by Pods.

Motivation: A Non-invazive way to quickly inspect container images used in cluster.

Scanners: Trivy --scanners <scanner1,scanner2> i.e. --scanners vuln,secret,misconfig default scanner is vulnerabilities scanner

✅ Interactive scanning: Prompts before scanning each image
✅ Namespace/pod validation: Helps users find the right resources
✅ Severity filtering: Only shows HIGH and CRITICAL vulnerabilities

## Installation

```bash
kubectl krew install imagescan
```

## Usage

```bash
# local test
chmod +x kubectl-imagescan
export PATH="$PATH:kubectl-imagescan"

kubectl imagescan pod <namespace> <pod-name> 
```

## Krew

```bash
# local test
kubectl krew install --manifest=krew/imagescan.yaml
```

## Test stuff

```bash

# compute sha

curl -L --fail --silent 'https://raw.githubusercontent.com/dejanu/kubectl-imagescan/v1.0.0/kubectl-imagescan' \
  | shasum -a 256 | awk '{print $1}'    # macOS
  
curl -L --fail -o /tmp/kubectl-imagescan.raw 'https://raw.githubusercontent.com/dejanu/kubectl-imagescan/v1.0.0/kubectl-imagescan'

# mac
shasum -a 256 /tmp/kubectl-imagescan.raw 

# Linux
sha256sum /tmp/kubectl-imagescan.raw 

file /tmp/kubectl-imagescan.raw
```