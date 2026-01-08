# kubectl-imagescan

Scan container images used by Pods.

Motivation: A Non-invazive to quickly inspect container images used in cluster.
Scanners : Trivy --scanners <scanner1,scanner2> i.e. --scanners vuln,secret,misconfig default scanner is vulnerabilities scanner

## Installation

```bash
kubectl krew install imagescan
```

## Usage

```bash
kubectl imagescan pod <pod-name>
kubectl imagescan pod -n <namespace> <pod-name> 
```
