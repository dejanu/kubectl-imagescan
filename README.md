# kubectl-imagescan

Scan container images used by Pods.

Motivation: A non-invasive, on-the-spot inspection of container images used in a by workloads in a K8S cluster.

Scanners: Trivy `--scanners <scanner1,scanner2>` i.e. `--scanners vuln,secret,misconfig` default scanner is vulnerabilities scanner

✅ Interactive scanning: Prompts before scanning each image

✅ Namespace/pod validation: Helps users find the right resources

✅ Severity filtering: Only shows HIGH and CRITICAL vulnerabilities

## Krew Installation

```bash
# kubectl plugin list
kubectl krew list

kubectl krew install imagescan
kubectl krew uninstall imagescan
```

## Usage

```bash
# local test
chmod +x kubectl-imagescan
export PATH="$PATH:kubectl-imagescan"

kubectl imagescan pod <namespace> <pod-name> 
```


