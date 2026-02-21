# 02 - Creating a Local Kubernetes Cluster

## Purpose

We need a Kubernetes control plane and worker node to deploy our application.

Instead of using a cloud provider, we use kind to:

- Simulate a production-like cluster
- Test deployments locally
- Practise Kubernetes fundamentals

---

## Create Cluster

```bash
kind create cluster --name student-platform
```

This command:
- Creates a Kubernetes control plane container 
- Configures networking 
- Sets kubectl context automatically

## Verify Cluster

```bash
kubectl get nodes
```

Expected result:
- 1 node in Ready state

This confirms:
- API server is running 
- kubectl is configured correctly 
- The cluster is operational

## Why This Matters

Kubernetes works on a declarative model.

The cluster continuously reconciles:

Desired State → Defined in YAML
Actual State → Running containers

Creating the cluster allows us to begin deploying workloads declaratively.
