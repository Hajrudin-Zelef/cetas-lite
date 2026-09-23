---
id: etape7-phaseb-containers/00-containers/overview
title: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AMD", "Nvidia"]
dates: ["2025-11-10", "2026-01-02", "2026-09", "2026-09-03", "2026-09-15", "2026-09-22"]
keywords: ["sandbox", "amd", "aws", "benchmark", "distribution", "gpu", "gpus", "nvidia", "packaging", "research"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [1, 87]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: d33ca9e9fe5be232f74504d617bc4db175d9ca7a7c6ccdbff2b0afd34247bb75
---

# Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes

**Scope:** Container runtimes (Docker, containerd, Podman, CRI-O), desktop/local tooling,
Kubernetes releases and distributions (k3s, RKE2, Talos, OpenShift, EKS/AKS/GKE),
packaging (Helm, Kustomize, operators), system containers (LXC/LXD/Incus),
sandbox/microVM runtimes (Firecracker, gVisor, Kata), registries and image supply-chain
security, CNI/CSI/service mesh, GPU-in-container stack.

**Observation date / research cutoff:** 2026-09-22.
**Method:** Read-only web research (search + page fetch), observed 2026-09-22.
**Language:** English (per project rule).

## Provenance legend

- `[official]` — vendor/maintainer documentation or release notes.
- `[vendor-reported]` — vendor blog/announcement of own product.
- `[independent]` — third-party reporting, benchmark, or analysis.
- `[secondary]` — community docs, wikis, forums, aggregators.
- `[unverified]` — single weak source or unconfirmed claim; flagged, not asserted as fact.

> Rules: never invent SKUs, prices, URLs, or versions. Gaps, conflicts, and
> non-comparable claims are flagged inline.

---

## 1. Container runtimes

### 1.1 Docker Engine 29

- **Latest release:** Docker Engine **29.8.1**, released **2026-09-15** [official].
  - Source: https://DOCS.DOCKER.COM/engine/release-notes/29/
- Docker Engine 29.0 GA was **November 10, 2025**; security support is active as of 2026-09-22 [secondary].
  - Source: http://eosl.date/eol/product/docker-engine/
- **29.8.1 changes:** containerd image store fix for `docker load` leaving dangling images
  (moby/moby#53595); user-namespace detection fix on OpenVZ; Windows hard-link preservation
  on commit; bundled containerd static binaries updated to **v2.3.5** [official].
- **29.8.0 (2026-09-03):** new `--umask <octal>` flag for `docker create`/`docker run`;
  awslogs logging driver can attach service names/environments/custom CloudWatch entity
  attributes; daemon support for a default AppArmor profile template [official].
- **Architectural shift:** with Engine 29, the **containerd image store became the default**
  for new installations; aligns Docker with containerd-based tooling and unlocks
  snapshotter optimizations, lazy pulling, and P2P image distribution roadmaps [independent].
  - Source: https://linuxiac.com/docker-engine-29-containerd-becomes-default-experimental-nftables-support/
- **Experimental nftables backend:** `dockerd --firewall-backend=nftables` creates nftables
  rules directly instead of iptables translation; still experimental in v29, not recommended
  for production; Swarm support planned [independent].
- **Minimum API version raised to 1.44** (Moby v25); Docker clients v25 and below cannot talk
  to the v29 engine by default [independent].
- Moby migrated to Go modules in the v29 cycle [independent].

### 1.2 containerd

- Docker Engine 29.8.x bundles **containerd v2.3.5 / v2.3.4** static binaries [official].
- EKS 1.34 notes reference **containerd 2.1** as the launch runtime for that version [official].
  - Source: https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions-standard.html
- NVIDIA Cloud Native Stack batch 25.12.0 (released **2026-01-02**) pins **containerd 2.2.0**
  with NVIDIA Container Toolkit 1.18.1 [vendor-reported].
  - Source: https://github.com/NVIDIA/cloud-native-stack/blob/master/README.md
- Community homelab references show **containerd 2.1.5 → 2.3.4** across K3s upgrades in
  September 2026 [secondary].
  - Source: https://github.com/toddpillars/homelab/blob/HEAD/docs/operations/k3s-upgrade.md
- Containerd 2.x is now the de-facto K8s CRI runtime; Docker's v29 default to the containerd
  image store completes the convergence that began with the 1.x/2.x line [independent].

### 1.3 CRI-O

- NVIDIA Cloud Native Stack 25.12.0 pins **CRI-O 1.34.2** (paired with K8s 1.34.2) and
  CRI-O 1.33.6 / 1.32.10 on older branches [vendor-reported].
- CRI-O continues as the Red Hat/OpenShift-preferred CRI runtime; 1.x.y tracks Kubernetes
  minor versions 1:1 [secondary].
- NVIDIA Container Toolkit 1.18.0 deprecated the OCI `hook`-based config mode for CRI-O;
  drop-in config files are now the recommended configuration mechanism [official].
  - Source: https://github.com/nvidia/cloud-native-docs/blob/HEAD/container-toolkit/release-notes.md
- **[gap]** Exact CRI-O 1.35/1.36 release dates were not independently confirmed in this
  research pass.

### 1.4 OCI low-level runtimes

- **runc:** `runc 1.3.0-rc.2` observed in NVIDIA debug containers (September 2026),
  spec 1.2.1, libseccomp 2.5.6 [secondary].
  - Source: https://github.com/ctrsploit/dqd/blob/HEAD/nvidia-container-toolkit/v1.17.6-runc-v1.3.0-rc.2/README.md
- **crun** remains the Podman default on most distros; the NVIDIA toolkit deprecated the
  `chmod` CDI hook that existed as a crun workaround (issue since resolved upstream) [official].
- Podman's `--gpus` option gained **AMD GPU** compatibility in Podman 6.0 (see §3) [independent].

---

