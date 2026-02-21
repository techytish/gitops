# 03 - Building and Pushing the Container

## Purpose

Kubernetes does not run source code directly.

It runs containers.

We must:

1. Compile the Go application
2. Package it into a container image
3. Push it to a container registry

---

## Build Image

```bash
docker build -t <your-username>/student-platform:latest .
```

This:
- Compiles the Go binary 
- Creates a minimal runtime image 
- Tags it for the container registry

## Push Image

```bash
docker push <your-username>/student-platform:latest
```

## Why?

Kubernetes nodes pull images from registries.

If the image is not available in a registry,
the cluster cannot deploy it.

## Verification

```bash
docker images
```