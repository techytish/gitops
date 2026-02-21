# 01 - Local Environment Setup

## Purpose

Before deploying to Kubernetes, a local development environment must be prepared.

This ensures:

- Container builds can run
- Kubernetes clusters can be created
- CLI tools are available
- We can interact with the cluster programmatically

---

## Install Required Tooling

### Docker

Docker is required to:
- Build container images
- Run kind (Kubernetes in Docker)

Install via Homebrew:

```bash
brew install --cask docker
```

Start Docker Desktop after installation.

Verification:
```bash
docker version
```

### Go

Go is required to:
- Build the microservice 
- Manage dependencies

Install via Homebrew:
```bash
brew install go
```

Verification:
```bash
go version
```

### kubectl

kubectl is the Kubernetes CLI used to interact with clusters.

```bash
brew install kubectl
```

Verification:
```bash
kubectl version --client
```

### kind

kind (Kubernetes in Docker) allows us to run a local Kubernetes cluster using Docker containers.

```bash
brew install kind
```

Verification:
```bash
kind version
```

### Helm

Helm is a package manager for Kubernetes. It is used later for installing Prometheus.

```bash
brew install helm
```

Verification:
```bash
helm version
```

