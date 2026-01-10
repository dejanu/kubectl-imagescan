# kubectl-imagescan

Scan container images used by Pods.

Motivation: A Non-invazive way to quickly inspect container images used in cluster.

Scanners: Trivy `--scanners <scanner1,scanner2>` i.e. `--scanners vuln,secret,misconfig` default scanner is vulnerabilities scanner

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
shasum -a 256 ./kubectl-imagescan_v1.0.0_linux_amd64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.0_linux_arm64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.0_darwin_amd64.tar.gz | awk '{print $1}'
shasum -a 256 ./kubectl-imagescan_v1.0.0_darwin_arm64.tar.gz | awk '{print $1}'
```
