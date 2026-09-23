---
id: etape7-phaseb-containers/00-containers/9-sandbox-and-microvm-runtimes
title: "9. Sandbox and microVM runtimes"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Google", "Intel", "Lambda"]
dates: ["2026-06-03", "2026-08-21", "2026-09", "2026-09-07", "2026-09-10", "2026-12-03", "2027-03-10"]
keywords: ["sandbox", "agent", "agents", "amd", "apache", "aws", "intel"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [407, 464]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: a5ec06472511b26105be58746312ac0d41ccab04f15117385043d9bc617d96c2
---

# 9. Sandbox and microVM runtimes

## 9. Sandbox and microVM runtimes

### 9.1 Firecracker

- **Latest:** **v1.17.0** released **2026-09-10**; support policy: v1.17 supported to
  2027-03-10; v1.16 (2026-06-03) supported to 2026-12-03; v1.15 EOL'd on v1.17 release [official].
  - Source: https://github.com/firecracker-microvm/firecracker/commit/a5a45f68b862689efa31ea0847695c108d872121
- KVM-based microVM VMM (Rust), minimalist device model; powers AWS Lambda/Fargate;
  Apache 2.0; ~35,820 GitHub stars [official/secondary].
  - Source: https://lib.rs/gh/firecracker-microvm/firecracker/firecracker
- **firecracker-containerd:** control plugin + `aws.firecracker` runtime shim v2 +
  in-guest agent (runc) + devmapper snapshotter; OCI container unit of work, microVM
  unit of tenancy [secondary].
  - Source: https://github.com/dymoo/microvm/blob/HEAD/docs/firecracker-containerd-evaluation.md
- Kata 4.1.0 bundles **Firecracker 1.12.1** as one of its hypervisor backends [secondary].
  - Source: https://github.com/proompteng/lab/blob/HEAD/docs/runbooks/talos-latest-upgrade-plan.md
- Hardening practice: jailer chroot + `CAP_SYS_ADMIN` via file caps, virtio-block/net/vsock
  only, no SSH into guests [secondary].
- **[unverified]** `v1.18.0-dev` referenced in community docs (2026-09-07); contents not
  verified against official release notes.

### 9.2 gVisor

- Google's application-kernel sandbox (`runsc`): user-space kernel intercepting syscalls;
  used by GKE Sandbox / Autopilot restricted workloads [secondary].
- **[gap]** Exact 2026 gVisor release numbers not captured verbatim this pass.

### 9.3 Kata Containers

- **Kata 4.1.0** released **2026-08-21** (first monthly snapshot after 4.0.0); **4.0.0** made
  the Rust runtime (**runtime-rs**) the default [secondary].
  - Source: https://github.com/katexochen/kata-containers/blob/HEAD/docs/releases/4.1.0.md
- 4.1.0 highlights: new hypervisor backend, new network model, VM templates on Cloud
  Hypervisor and Dragonball; release tarball split by runtime; `kata-deploy` job mode
  drops node-selection keys and host privileges; runtime-rs default on all arches [secondary].
- **Security:** GHSA-fmg6-v47x-52wr / CVE-2026-77176 — genpolicy allowed attacker-chosen
  guest paths to be mounted in Confidential Containers; fixed in 4.1.0, upgrade +
  policy regeneration required [independent].
  - Source: https://securityonline.info/cve-2026-77176-kata-containers-guest-rootfs/
- **Kata 4.2.0** (September 2026): job-mode dispatcher moved to independent
  `k8s-job-dispatcher` project (`ghcr.io/kata-containers/k8s-job-dispatcher:0.3.0`);
  old `quay.io/kata-containers/kata-deploy-job-dispatcher` no longer published —
  air-gapped mirrors must repoint [secondary].
- Requirements: hardware virtualization (VT-x/AMD-V/ARM Hyp/Power/IBM Z SIE), bare metal
  or nested virt; hypervisors: QEMU, Cloud Hypervisor, Firecracker, Dragonball [official].
- **Agent sandboxing:** OpenInfra positions Kata 4.x as the runtime for the Kubernetes
  SIG Apps "Agent Sandbox" project — VM-level isolation for nondeterministic AI agents
  (Ant Group cited as production user) [independent].
  - Source: https://cloudnativenow.com/features/rust-rewrite-readies-kata-containers-for-agent-sandboxing/
- Talos ships a signed Kata system extension exposing QEMU/Cloud Hypervisor/Firecracker/
  Dragonball without custom controllers [secondary].

### 9.4 Confidential Containers

- CoCo builds on Kata (TEE-backed pods: AMD SEV-SNP, Intel TDX); the 4.1.0 policy CVE
  above is CoCo-specific [secondary].
- See also Phase F3 (platform security) for TDX/SEV-SNP hardware coverage.

