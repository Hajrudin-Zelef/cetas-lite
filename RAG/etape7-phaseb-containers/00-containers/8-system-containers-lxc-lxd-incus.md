---
id: etape7-phaseb-containers/00-containers/8-system-containers-lxc-lxd-incus
title: "8. System containers: LXC, LXD, Incus"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: []
dates: ["2026-05-01"]
keywords: ["agent", "governance"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [362, 406]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 322cd3a17ea4f095ed51061f784561c8371682a56d67e52d291d544cf0bf3eda
---

# 8. System containers: LXC, LXD, Incus

## 8. System containers: LXC, LXD, Incus

### 8.1 The fork

- **Incus** is the community fork of LXD, led by original LXD lead Stéphane Graber,
  created after Canonical took LXD in-house (2023) and changed governance [independent].
  - Source: https://www.Theregister.Com/2024/04/10/lxc_6_and_incus_6/?td=keepreading
- **LXC 6.0 LTS** (2024): dropped Upstart, single multi-call binary option, IPv6 by default;
  support until **2029** [independent].

### 8.2 Incus 7.0 LTS (released 2026-05-01)

- **Incus 7.0 LTS** shipped **2026-05-01** alongside **LXC 7.0 LTS** and **LXCFS 7.0 LTS** [independent].
  - Source: https://linuxiac.com/incus-7-0-lts-container-virtual-machine-manager-released/
- Highlights: OCI image support (from 6.3) — application containers from OCI images with
  resource limits and syscall interception; LINSTOR remote storage (DRBD replication);
  **TrueNAS storage driver** (remote TrueNAS pool via API + iSCSI); network address sets
  for ACLs; cluster-group CPU baselines for mixed hardware [independent].
- 6.22-era features now in LTS: Windows VM agent over vsock; direct backup streaming;
  disk-only snapshot restore; QCOW2 defaults; ACME multi-domain certs; SR-IOV NIC
  `security.trusted`; instance boot-time metrics [independent].
  - Source: https://linuxiac.com/incus-6-22-container-and-virtual-machine-manager-released/
- Install path of record on Ubuntu: Zabbly stable repo (ships 7.0); Ubuntu 24.04/26.04
  archives ship the 6.0 LTS branch [secondary].
  - Source: https://computingforgeeks.com/install-lxc-incus-ubuntu/
- **Positioning:** system containers (full userland, VM-like) + real QEMU VMs from one CLI;
  heavier than Docker, lighter than KVM-only; sweet spot = homelab/small-fleet hosting [secondary].

### 8.3 LXD (Canonical)

- Canonical LXD continues as the Ubuntu-integrated path (Ubuntu Pro images, OVN, MAAS
  integrations); community momentum is with Incus [independent].
- **[gap]** Exact Canonical LXD 2026 version numbers not captured verbatim this pass.

### 8.4 LXC vs Docker (use-case split)

- **LXC/Incus:** full OS userland per container, persistent state, systemd inside —
  "lightweight VMs"; multi-tenant hosting, dev environments, legacy app lift.
- **Docker/Podman:** single-process application containers, immutable images, ephemeral —
  microservices, CI, K8s pods.
- Incus 7's OCI support narrows the gap: run Docker/OCI images as Incus application
  containers with Incus networking/storage [independent].

---

