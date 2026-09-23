---
id: etape7-phaseb-containers/00-containers/6-kubernetes-distributions
title: "6. Kubernetes distributions"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AWS", "Google", "Nvidia"]
dates: ["2025-08-27", "2025-10-02", "2025-12-17", "2026-01-27", "2026-04-22", "2026-04-27", "2026-05-27", "2026-06-02", "2026-09", "2026-11-26", "2026-12", "2026-12-02", "2027-03-27", "2027-07-29", "2027-08-02"]
keywords: ["distribution", "acquisition", "aws", "gpu", "nvidia", "omni", "research"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [228, 316]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 42c36f323cb6541b8ede7e1b7006e819977963d69c5b35db918136354f76dff3
---

# 6. Kubernetes distributions

## 6. Kubernetes distributions

### 6.1 K3s (SUSE/Rancher)

- **Current:** `v1.36.4+k3s1` observed September 2026; release line tracks upstream minors
  with `+k3sN` build suffix [vendor-reported].
  - Source: https://docs.k3s.io/blog/2026/05/27/K3s-1.36-release
- **K3s 1.36 highlights (2026-05-27):** bundled **etcd 3.6** (data-directory restructure —
  no clean downgrade to 3.5; snapshot restore procedure required for rollback);
  **Nix snapshotter** plugin (Nix store paths as container layers); containerd config
  drop-in imports [vendor-reported].
- Traefik bundled chart auto-upgraded v37 → v40.1.4 (Traefik 3.7.8) across the 1.35 hop;
  containerd 2.1.5 → 2.3.4 [secondary].
  - Source: https://github.com/toddpillars/homelab/blob/HEAD/docs/operations/k3s-upgrade.md
- Install: `curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION="v1.35.0+k3s1" sh -`
  (Linux only; macOS/Windows need Lima/Multipass/WSL2 or kind) [secondary].
- Typical deploy: single binary, SQLite default datastore (HA via etcd or external DB),
  Flannel CNI + kube-proxy defaults (commonly replaced by Cilium in production) [secondary].

### 6.2 RKE2

- Rancher's hardened distribution; tracks upstream K8s minors with `+rke2rN` suffix;
  CIS-hardened by default, etcd bundled, SELinux policies [secondary].
- 2026 lines track K8s 1.34/1.35/1.36; used as the "production-grade Rancher" option vs
  K3s for edge [secondary].
- **[gap]** Exact latest RKE2 build numbers for September 2026 were not captured verbatim.

### 6.3 Talos Linux (Sidero Labs)

- **Talos 1.13.0** released **2026-04-27**: Clang-built kernel with ThinLTO; Container Device
  Interface (CDI) enabled by default; `talosctl debug` privileged debug container;
  `EnvironmentConfig` document replacing `.machine.env`; virtiofs external volumes
  via `ExternalVolumeConfig` [official].
  - Source: https://github.com/siderolabs/talos/releases
- Community rollouts pin **Talos v1.13.9 + K8s v1.36.4** (September 2026) [secondary].
  - Source: https://github.com/proompteng/lab/blob/HEAD/docs/runbooks/talos-latest-upgrade-plan.md
- Architecture: immutable, API-driven, **no SSH/shell/package manager**; ~12 binaries total;
  management via `talosctl` gRPC API; read-only SquashFS root [secondary].
  - Source: https://github.com/fertile-org/banyan/blob/HEAD/docs/whitepaper/research-talos-omni.md
- **Omni** = fleet management plane (UI-driven provisioning; Proxmox infra provider);
  Omni v1.9 current mid-2026 [secondary].
  - Source: https://github.com/mitchross/sidero-omni-talos-proxmox-starter/blob/HEAD/README.md
- **Acquisition:** Sidero Labs acquired by **Yardi Systems** (announced ~September 2026);
  open-source foundation preserved per announcement [secondary].
  - Source: https://lifestyle.articlerich.com/story/353759/talos-linux-adds-native-hypervisor-and-edge-container-support-see-it-live-at-taloscon-2026/
- **Talos Hypervisor (alpha):** native hypervisor + edge container support — run containers
  directly on Talos without Kubernetes for single-node/edge sites; public demo at
  **TalosCon 2026, Amsterdam, October 15–16**; GA targeted December 2026 [vendor-reported via secondary].
- TalosCon 2026 noted [secondary]; schedule details beyond the October date are [unverified].

### 6.4 Red Hat OpenShift

- OpenShift Container Platform 4.x ships CRI-O, tracks K8s with ~4–6 month lag; OpenShift
  Virtualization built on KubeVirt; NVIDIA vGPU 17.4 support for A30/A100/H100 noted in
  GPU Operator docs (requires NVIDIA AI Enterprise) [official/secondary].
  - Source: https://github.com/nvidia/cloud-native-docs/blob/HEAD/gpu-operator/release-notes.rst
- OpenShift 4.19/4.20 align with K8s 1.32/1.33-era payloads; exact 2026-09 versions
  are **[gap]** — not captured verbatim in this pass.
- Docker launched Testcontainers support on OpenShift (noted in 2025 coverage) [secondary].

### 6.5 Managed Kubernetes: EKS / AKS / GKE

**Amazon EKS** [official] — https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions

| K8s | Upstream | EKS GA | Standard support ends |
| --- | --- | --- | --- |
| 1.36 | 2026-04-22 | 2026-06-02 | 2027-08-02 |
| 1.35 | 2025-12-17 | 2026-01-27 | 2027-03-27 |
| 1.34 | 2025-08-27 | 2025-10-02 | 2026-12-02 |

- 1.33 in extended support (ends 2027-07-29); 1.31 extended ends 2026-11-26 [official].
- EKS Auto Mode, EKS Hybrid Nodes, and EKS Pod Identity continue as the 2026 control-plane
  differentiators; EBS CSI driver with VAC sidecars patched through EKS 1.33 EOL [official].
- `aws eks describe-cluster-versions` exposes the version lifecycle API [official].

**Azure AKS / Google GKE:** support K8s 1.34/1.35/1.36 in 2026 with their own ~monthly patch
cadence; GKE attached (multi-cloud) clusters validated on 1.33/1.34 with CVEs
2025-12817/12818/9230/9232 fixed in `*-gke.1/2` builds [official].
- Source: https://docs.cloud.google.com/kubernetes-engine/multi-cloud/docs/attached/eks/reference/supported-versions
- **[gap]** Exact AKS default version for September 2026 not captured verbatim.

### 6.6 Lightweight/local clusters

- **kind**, **k3d**, **minikube** remain the local-dev standards; Docker Desktop and OrbStack
  bundle managed single-node options [independent].
- **Canonical k8s (microk8s successor line)** and Charmed Kubernetes ship 1.34 channels
  with calico 3.29.3, coredns 1.11.1 charms [secondary].
  - Source: https://github.com/charmed-kubernetes/kubernetes-docs/blob/HEAD/pages/k8s/release-notes.md

