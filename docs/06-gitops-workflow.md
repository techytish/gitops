# 06 - Configuring GitOps

## Purpose

We now configure Argo CD to monitor our Git repository and automatically deploy changes.

Argo CD continuously:

- Pulls manifests from Git
- Compares desired vs actual cluster state
- Applies changes if drift is detected

---

## Create Application in Argo CD

Inside the Argo CD UI:

1. Click **New App**
2. Configure:

- Application Name: student-platform
- Project: default
- Repository URL: <your-github-repo-url>
- Path: k8s
- Cluster URL: https://kubernetes.default.svc
- Namespace: default
- Sync Policy: Automatic

Click **Create**.

---

## Verify Sync

Argo CD should show:

- Application status: Healthy
- Sync status: Synced

This confirms the cluster matches Git state.

---

## Test GitOps Reconciliation

Edit `k8s/deployment.yaml`:

```yaml
replicas: 3
```

Commit and push:

```bash
git add .
git commit -m "feat: scale replicas to 3"
git push
```

Argo CD will automatically:

- Detect change
- Reconcile desired state
- Perform rolling update

---

## Why GitOps?

GitOps enables:

- Auditability (all changes via Git)
- Controlled change management
- Easy rollback
- Drift visibility
- Environment promotion strategies

Git becomes the single source of truth.