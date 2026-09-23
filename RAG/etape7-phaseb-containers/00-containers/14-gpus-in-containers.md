---
id: etape7-phaseb-containers/00-containers/14-gpus-in-containers
title: "14. GPUs in containers"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: hardware
actors: ["AMD", "EU", "Intel", "Nvidia"]
dates: ["2026-09"]
keywords: ["gpu", "gpus", "acquisition", "amd", "blackwell", "cost", "distribution", "intel", "nvidia", "omni", "packaging"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [611, 726]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 01704bf2f6fe117621aeafbc23cb031d77dd5b63211c894b3e7ed0f452c49d76
---

# 14. GPUs in containers

## 14. GPUs in containers

### 14.1 NVIDIA Container Toolkit

- **Latest:** **v1.19.0 / v1.19.1** (lifecycle policy, September 2026); **v1.18.1** widely
  deployed via Cloud Native Stack [official].
  - Source: https://github.com/nvidia/cloud-native-docs/blob/HEAD/gpu-operator/life-cycle-policy.rst
- **1.18.0 headline:** default runtime mode switched from `legacy` to just-in-time
  **CDI** spec generation; systemd unit auto-generates CDI specs; legacy mode deprecated
  but still supported; CRI-O OCI-hook config deprecated in favor of drop-in files [official].
  - Source: https://github.com/nvidia/cloud-native-docs/blob/HEAD/container-toolkit/release-notes.md
- Driver versions observed: 580.105.08 (CNS 25.12.0), 595.71.05/610.43.02 (community,
  CachyOS) [vendor-reported/secondary].

### 14.2 NVIDIA GPU Operator

- Component matrix (recent): Container Toolkit 1.17.x–1.19.x, Kubernetes Device Plugin
  0.17.0–0.19.3, DCGM Exporter 3.3.8–4.5.x, MIG Manager 0.10.0–0.14.2, NFD 0.16.6–0.18.3,
  vGPU Device Manager 0.2.8–0.4.2, KubeVirt GPU plugin 1.2.10–1.5.0, Kata Manager 0.2.2 [official].
  - Source: https://github.com/nvidia/cloud-native-docs/blob/HEAD/gpu-operator/release-notes.rst
- **MIG profiles added:** GH200 NVL2 144GB HBM3e (`1g.18gb`, `1g.36gb`, `2g.36gb`,
  `3g.72gb`, `4g.72gb`, `7g.144gb`, `all-balanced`) [official].
- Precompiled driver containers GA; operator/NFD CRDs auto-upgrade default true [official].
- K8s DRA (GA in 1.34) is the long-term successor to the device-plugin model for GPU
  allocation; NVIDIA's DRA driver work ongoing through 2026 [independent].

### 14.3 MIG and fractional GPU in Kubernetes

- **MIG (Multi-Instance GPU):** A100/H100/H200/B200/GH200 partition into isolated GPU
  instances; exposed in K8s via `nvidia.com/mig-1g.10gb` style resources or DRA [official].
- NVIDIA KAI Scheduler 0.10.2 (CNS) adds gang-scheduling/queueing for AI workloads [vendor-reported].
- AMD alternative: ROCm device plugin / AMD GPU Operator; Intel: GPU plugin for
  Data Center GPU Flex/Max series [secondary].
- **[unverified]** Exact MIG-on-B200 community calculator data is third-party only;
  treat profile tables for Blackwell as provisional.

### 14.4 Podman/Docker GPU

- Docker: `--gpus all` via nvidia-container-runtime/CDI; compose `runtime: nvidia`
  (Ollama GPU template is the canonical homelab example) [secondary].
- Podman 6.0: `--gpus` with **AMD GPU** support; CDI-based [independent].

---

## 15. Ecosystem, adoption, and operations

### 15.1 CNCF and community state 2026

- Kubernetes remains the control-plane standard; graduation pace steady; KubeCon EU 2026
  (London) and NA 2025 (Atlanta, Helm 4 launch) as anchor events [independent].
- **Platform engineering:** Backstage, Argo CD/Workflows, Crossplane as the IDP stack;
  Talos+Omni and Cluster API as cluster lifecycle engines [independent].
- GitOps: Flux and Argo CD; Helm OCI charts as the packaging unit [independent].

### 15.2 Cost and operations

- Docker Desktop licensing (see §2.1) is the main 2026 cost debate for dev fleets;
  Podman/Rancher/OrbStack/Colima are the $0 alternatives [independent].
- K8s cost tooling: OpenCost/Kubecost, Karpenter consolidation, VPA/HPA/KEDA autoscaling [secondary].
- **KEDA** for event-driven autoscaling; **Descheduler**, **Goldilocks/VPA recommender**
  for rightsizing [secondary].
- Observability: Prometheus Operator/kube-prometheus-stack, Grafana, Loki, Tempo,
  Hubble (Cilium), Tetragon (eBPF runtime security) [secondary].

### 15.3 Security posture

- Pod Security Standards (restricted) replacing PSP (removed in 1.25); AppArmor
  deprecated in K8s 1.34 — seccomp default profiles via `SeccompDefault` [official].
- Image supply chain: signed images (cosign), SBOMs, admission policy (Kyverno/Gatekeeper),
  runtime threat detection (Falco, Tetragon) [independent].
- CVE-2025-52881 (Podman), CVE-2026-3288 (ingress-nginx), GHSA-fmg6-v47x-52wr (Kata/CoCo),
  CVE-2026-77176 — patched in current 2026 releases; pin versions [independent/secondary].

---

## 16. Conflicts, gaps, and unverified claims

- **C1 — Calico latest:** 3.31.3 seen in NVIDIA CNS (Jan 2026) but September-2026 latest
  not confirmed — treat 3.31.3 as a floor, not current [gap].
- **C2 — Linkerd version:** no verbatim 2026 stable number captured [gap].
- **C3 — gVisor:** 2026 release numbers not captured [gap].
- **C4 — RKE2:** 2026 build numbers not captured verbatim [gap].
- **C5 — OpenShift 4.19/4.20:** September-2026 exact versions not captured [gap].
- **C6 — Harbor:** latest 2.x release number not captured verbatim [gap].
- **C7 — Firecracker v1.18.0-dev:** referenced only in community docs [unverified].
- **C8 — CRI-O 1.35/1.36:** release dates not independently confirmed [gap].
- **C9 — Compose v5.x:** breaking changes vs v2 not fully enumerated [unverified].
- **C10 — MIG on Blackwell B200:** community-calculator data only [unverified].
- **C11 — Canonical LXD 2026:** exact version not captured; community momentum with
  Incus is the observable fact [gap].
- **C12 — AKS default version:** September-2026 default not captured verbatim [gap].
- **C13 — Sidero/Yardi acquisition:** single secondary source; terms undisclosed [unverified].
- **C14 — Talos Hypervisor GA (Dec 2026):** vendor-reported target, not yet shipped [vendor-reported].
- **Non-comparable:** Docker Hub "unlimited" pulls are fair-use unlimited, not
  contractual unlimited; EKS standard vs extended support windows differ from upstream
  EOL — do not mix the two calendars.

---

## 17. Glossary

- **CDI (Container Device Interface):** standard for advertising host devices (GPUs) to runtimes.
- **CNI:** Container Network Interface — pod networking plugin API.
- **CRI:** Container Runtime Interface — kubelet ↔ runtime API.
- **CSI:** Container Storage Interface — volume plugin API.
- **DRA:** Dynamic Resource Allocation — K8s-native device scheduling (GA 1.34).
- **VAC:** VolumeAttributesClass — storage parameter mutation (GA 1.34).
- **OCI:** Open Container Initiative — image/runtime/distribution specs.
- **MIG:** Multi-Instance GPU — hardware GPU partitioning (NVIDIA).
- **eBPF:** in-kernel programmable packet/security engine (Cilium, Tetragon).
- **SBOM:** Software Bill of Materials.
- **Quadlet:** Podman systemd unit generator for containers.
- **microVM:** minimal KVM virtual machine (Firecracker).

---

