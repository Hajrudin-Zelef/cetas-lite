---
id: etape7-phaseb-containers/00-containers/2-desktop-and-local-development-tooling
title: "2. Desktop and local development tooling"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AMD"]
dates: ["2024-12-10", "2026-08", "2026-09", "2026-09-14"]
keywords: ["amd", "apache", "cost", "gpu", "gpus", "open source", "pricing", "revenue"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [88, 164]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: cde46cd67dbd6e0421454658d201a64bdf4376b12dbf0a082987c27362f24a19
---

# 2. Desktop and local development tooling

## 2. Desktop and local development tooling

### 2.1 Docker Desktop

- **Latest release:** **4.91.0**, released **2026-09-14**, bundles Docker Engine v29.8.0,
  containerd v2.3.4, Compose v5.5.1, Buildx v0.37.0, Docker Offload v0.6.17 [official].
  - Source: https://github.com/rashworld786/docs-2/blob/HEAD/content/manuals/desktop/release-notes.md
- **Pricing 2026** (verified against docker.com/pricing, September 2026) [independent]:
  - Personal: **$0** — individuals at companies with <250 employees AND <$10M annual revenue.
  - Pro: **$9/mo annual ($108/yr)** or **$11/mo monthly**.
  - Team: **$15/user/mo annual ($180/yr)** or **$16 monthly**; caps at **100 seats**.
  - Business: **$24/user/mo** ($288/yr), no annual discount; unlimited seats; SSO/SCIM,
    hardened images, audit logs, VDI support.
  - Prices raised **December 10, 2024** (Pro +80% from $5, Team +67% from $9) [independent].
  - Sources: https://devtoolhub.com/docker-desktop-licensing-cost/,
    https://procurementvms.com/vendors/docker-pricing-guide.html
- Build Cloud minutes and Testcontainers Cloud bill separately beyond included
  allowances; 200/500/1,500 Build Cloud minutes by tier reported [independent].
- Bundled Kubernetes (kubeadm-based single node) still ships; Desktop docs also describe
  kind-based multi-node options [independent].
  - Source: https://sliplane.io/blog/orbstack-vs-docker

### 2.2 Alternatives

- **Podman Desktop:** free, open source (Apache 2.0); rootless Podman security model;
  enterprise-ready extensions, registry management, air-gapped options; no paid admin
  suite equivalent to Docker Business [independent].
  - Source: https://blog.apps.deals/docker-desktop-vs-orbstack-podman-desktop-rancher-desktop-mac
- **Rancher Desktop:** free, open source (Apache 2.0); bundles K3s, nerdctl, kubectl, helm,
  selectable container engines [independent].
- **OrbStack (macOS):** free for personal/non-commercial; **Pro $8/user/mo billed annually**
  for commercial use; ships Linux Machines (full Linux VMs, e.g. `orb create ubuntu:24.04`),
  native Swift app, Kubernetes support [independent].
  - Source: https://sliplane.io/blog/orbstack-vs-docker
- **Colima / Lima:** open-source VM-based container runtimes on macOS/Linux; Lima v2 has
  external drivers; Lima underpins several Podman Machine providers [secondary].

---

## 3. Podman

### 3.1 Releases 2026

- **Podman 6.1.0** is the latest feature release (announced ~August 2026) [independent].
  - Source: https://linuxiac.com/podman-6-1-adds-volume-renaming-machine-restart/
- **Podman 6.0.0** introduced breaking changes: `podman volume prune` now removes only unused
  anonymous volumes by default (Docker parity; `--all` restores old behavior); Podman Machine
  VMs mount host volumes via systemd on Linux (existing machines must be recreated);
  macOS default machine provider changed to **libkrun** [independent].
  - Source: https://linuxiac.com/podman-6-0-lands-with-breaking-changes-amd-gpus-support/
- Podman 6.0 added **AMD GPU** support to `--gpus` for `podman create`/`run`;
  multiple static IPs via repeated `--net ip=`; blackhole/unreachable/prohibit routes
  in `podman network create` [independent].
- Podman 6.1 highlights: `podman volume rename`; `podman machine restart`;
  Quadlet `.container` units gain `ImageVolume=`; `podman generate kube` emits healthchecks
  as `livenessProbe`; Pesto rootless port-forwarder gains IPv6; `force_port_listen`
  in containers.conf for WSL port forwarding [official].
  - Source: https://github.com/podman-container-tools/podman/blob/HEAD/RELEASE_NOTES.md
- Podman 5.7 (late 2025) added full TLS/mTLS for remote connections and fixed
  **CVE-2025-52881** (container escape via arbitrary write gadgets / procfs redirects) [independent].
- Podman 5.8: multi-file Quadlet install, SQLite state migration [secondary].
- Networking stack: **Netavark** + **Aardvark-DNS** replaced CNI plugins in the 5.x line;
  rootless networking is considered production-reliable from Podman 5 onward [independent].
  - Source: https://medium.com/@developeryusuf/docker-vs-podman-in-2026-i-migrated-our-servers-and-kept-my-laptop-heres-the-split-that-4cc8d13ad3cb

### 3.2 Quadlet

- Quadlet automates running containers as systemd services; `.container`, `.volume`,
  `.network`, `.kube`, `.image`, `.build`, `.pod`, `.artifact` unit types across 5.7–6.1 [official].
- `podman quadlet install` supports multiple files (5.8+); new search paths help distros
  package Quadlets (6.0) [independent].
- Operational note (independent, 2026): teams migrating 6 production hosts from Docker to
  Podman report ~4 engineer-days of work; Docker BuildKit retains ~15% cache-hit advantage
  in CI; common hybrid pattern = Docker for CI builds, Podman for runtime [independent].

---

