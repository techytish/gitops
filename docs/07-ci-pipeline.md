# 07 - CI Pipeline (GitHub Actions)

## Purpose

Automate container image builds when application code changes.

Pipeline flow:

1. Developer pushes code
2. GitHub Actions builds Docker image
3. Image is pushed to container registry
4. Argo CD reconciles deployment

---

## GitHub Actions Workflow

Create:

.github/workflows/docker-build.yml

```yaml
name: Build and Push Docker Image

on:
  push:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Login to DockerHub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build and Push Image
        run: |
          docker build -t <your-username>/student-platform:latest .
          docker push <your-username>/student-platform:latest
```

---

## Required GitHub Secrets

In repository settings → Secrets:

- DOCKER_USERNAME
- DOCKER_PASSWORD (or Docker access token)

---

## Why CI?

CI provides:

- Automated builds
- Reproducible artifacts
- Immutable container images
- Reduced manual error
- Fast feedback loop

Combined with GitOps, this forms a complete CI/CD pipeline.