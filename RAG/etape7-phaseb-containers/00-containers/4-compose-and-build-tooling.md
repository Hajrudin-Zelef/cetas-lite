---
id: etape7-phaseb-containers/00-containers/4-compose-and-build-tooling
title: "4. Compose and build tooling"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2025-08", "2025-12-17", "2026-01-27", "2026-04-22", "2026-06-02", "2026-06-23", "2026-07-30", "2026-09", "2026-09-22", "2026-10-27", "2027-02-28", "2027-06-28", "2027-08-02", "2028-08-02"]
keywords: ["accelerator", "agent", "aws", "gpu", "mcp", "research"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [165, 227]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 8fee519bbbf18e87fcd056831ee7b83b9a53399fff885b20510bf3d7f6b1fe77
---

# 4. Compose and build tooling

## 4. Compose and build tooling

- **Docker Compose v5.5.1** ships in Docker Desktop 4.91.0 (September 2026) [official].
- **Buildx v0.37.0** in Desktop 4.91.0; the `version:` key in compose files is deprecated
  and ignored [official/secondary].
- Docker Offload v0.6.17; Docker MCP gateway v0.42.2 and `docker pass` v0.1.2 appear in
  2026 Desktop releases — Docker is layering AI-agent tooling (MCP, Gordon) on Desktop [official].
- **[unverified]** Exact Compose v5.x breaking-change list vs v2 line was not fully
  enumerated in this pass.

## 5. Kubernetes — upstream releases

### 5.1 Version state as of 2026-09-22

| Minor | Latest patch (observed 2026-07-30) | Upstream state | EOL |
| --- | --- | --- | --- |
| 1.34 | 1.34.10 | Active | 2026-10-27 |
| 1.35 | 1.35.7 | Active | 2027-02-28 |
| 1.36 | 1.36.3 | Active | 2027-06-28 |

- Source: https://github.com/linkmaq/kube-accelerator-sim/blob/HEAD/docs/research/kubernetes-compatibility-policy.md
- 1.33 and earlier are upstream-EOL as of mid-2026 [secondary].

### 5.2 Kubernetes 1.34 "Of Wind & Will" (released August 2025)

- **58 enhancements:** 23 stable, 22 beta, 13 alpha [independent].
  - Source: https://medium.com/@harshalsant0/kubernetes-1-34-of-wind-will-deep-dive-whats-new-and-why-it-matters-dd50659da480
- **Headline: Dynamic Resource Allocation (DRA) GA.** Apps can prioritize device
  alternatives; DRA usage reportable via PodResources API; resource health status,
  binding conditions, consumable capacity. Standard mechanism for GPU/FPGA/NIC
  allocation for ML/AI/HPC [independent].
  - Source: https://www.i-programmer.info/news/240-devops/18359-kubernetes-134-adds-dynamic-resource-allocation.html
- **VolumeAttributesClass (VAC) GA:** graduated from `storage.k8s.io/v1beta1` to
  `storage.k8s.io/v1`; enables volume modification rollback [official].
  - Source: https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions-standard.html
- **AppArmor deprecated in 1.34** (AWS guidance: migrate to seccomp / Pod Security
  Standards) [official].
- External JWT signer for service-account tokens promoted to Beta [official].
- containerd updated to **2.1** for EKS 1.34 launch; no EKS-optimized Amazon Linux 2 AMI
  for 1.34 — AWS pushes migration to Amazon Linux 2023 [official].

### 5.3 Kubernetes 1.35 (released December 17, 2025)

- Upstream GA **2025-12-17**; EKS release **2026-01-27** [official].
  - Source: https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions
- ImageVolume GA (used by DocumentDB operator, which requires K8s 1.35+) [secondary].
  - Source: https://github.com/documentdb/documentdb-kubernetes-operator/blob/HEAD/docs/operator-public-documentation/preview/getting-started/quickstart-k3s.md
- Cluster-autoscaler and cert-manager ecosystems lagged 1.36 adoption in mid-2026
  (no `v1.36.x` cluster-autoscaler image; cert-manager 1.21 only in alpha as of 2026-06-23),
  making 1.35 the conservative production target at that time [secondary].
  - Source: https://github.com/aws-samples/amazon-eks-machine-learning-with-terraform-and-kubeflow/blob/HEAD/docs/upgrade-1.35.md

### 5.4 Kubernetes 1.36 (released April 22, 2026)

- Upstream GA **2026-04-22**; EKS GA **2026-06-02** [official].
- Standard support through **2027-08-02** on EKS (extended to 2028-08-02) [official].
- Ecosystem note: K3s shipped v1.36.4+k3s1 with etcd 3.6, Nix snapshotter, containerd 2.3.x
  (see §6.1) [vendor-reported].
- **[unverified]** The exact upstream 1.36 headline-feature list (beyond DRA/VAC lineage)
  was not exhaustively re-verified in this pass.

---

