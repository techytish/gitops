# 09 - Scaling and Enhancements

## Manual Scaling

Scale replicas:

```bash
kubectl scale deployment student-platform --replicas=4
```

Verify:

```bash
kubectl get pods
```

Kubernetes automatically schedules additional Pods.

---

## Production Enhancements

Future improvements could include:

- Horizontal Pod Autoscaler (HPA)
- CPU and memory resource limits
- Liveness probes
- Readiness probes
- Ingress controller integration
- Multi-environment GitOps (dev/stage/prod)
- Canary deployments
- Image version pinning
- Helm chart packaging

---

## Why Scaling Matters

Kubernetes enables:

- Horizontal scaling
- Self-healing workloads
- High availability
- Declarative rollout strategies

These are core platform engineering capabilities.