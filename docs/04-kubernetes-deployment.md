# 04 – Deploying to Kubernetes

## Purpose

We now define the desired state of the application using YAML.

This includes:

- Deployment (manages Pods)
- Service (network abstraction)

---

## Apply Manifests

```bash
kubectl apply -f k8s/
```

Kubernetes now:
- Creates a Deployment 
- Creates a ReplicaSet 
- Schedules Pods 
- Exposes a Service

## Verify Deployment

```bash
kubectl get deployments
kubectl get pods
kubectl get svc
```

Pods should be in Running state.

## Test Application

Visit:
```bash
http://localhost:30007
```

This confirms:
- Container is running 
- Service routing works 
- Networking is configured correctly

## Why Deployments?

Deployments provide:
- Rolling updates 
- Replica management 
- Self-healing if Pods fail