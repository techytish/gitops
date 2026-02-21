# 08 - Observability

## Purpose

Expose application metrics to support monitoring and performance visibility.

The Go service exposes Prometheus-compatible metrics at:

/metrics

---

## Verify Metrics Endpoint

Port-forward service if needed:

```bash
kubectl get svc
```

Visit:

```
http://localhost:30007/metrics
```

You should see:

- app_requests_total
- Default Go runtime metrics

---

## Optional: Install Prometheus

Add Helm repository:

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

Install:

```bash
helm install prometheus prometheus-community/prometheus
```

Port forward:

```bash
kubectl port-forward svc/prometheus-server 9090:80
```

Visit:

```
http://localhost:9090
```

Query:

```
app_requests_total
```

---

## Why Observability Matters?

Observability enables:

- Performance monitoring
- Capacity planning
- Alerting
- SLO tracking
- Debugging production issues

Metrics are a foundational platform engineering capability.