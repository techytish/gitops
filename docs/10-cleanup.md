# 10 - Cleanup

To remove the local Kubernetes cluster:

```bash
kind delete cluster --name student-platform
```

To remove Argo CD namespace:

```bash
kubectl delete namespace argocd
```

To remove Prometheus (if installed):

```bash
helm uninstall prometheus
```

---

## Final Notes

This project demonstrates:

- Containerisation with Go
- Kubernetes fundamentals
- GitOps deployment with Argo CD
- CI pipeline automation
- Observability integration
- Horizontal scalability

It is intentionally lightweight but models real-world cloud-native deployment patterns.