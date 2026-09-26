---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-5-dpdk-spdk-2026-releases-pmd-coverage-and-where-they-r
title: "Wave 5 — DPDK & SPDK: 2026 releases, PMD coverage, and where they run"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Nvidia", "Samsung"]
dates: ["2026-01-28", "2026-05", "2026-05-28", "2026-07"]
keywords: ["agent", "amd", "blackwell", "compute", "gpu", "gpus", "latency", "license", "licenses", "nvidia", "packaging", "throughput"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [235, 283]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: d3af80b4d23b5df1e213af2d4803fac654e177d3003858ac62ac1c6e23c8e84c
---

# Wave 5 — DPDK & SPDK: 2026 releases, PMD coverage, and where they run

## Wave 5 — DPDK & SPDK: 2026 releases, PMD coverage, and where they run

### 5.1 DPDK 2026 release train
- **DPDK 25.11 is the current LTS** (Long Term Support) release, in wide packaging: Debian sid/forky ships 25.11.3-1; openSUSE Tumbleweed/Leap 16.1 ships 25.11 across x86_64/aarch64/ppc64le; RHEL 10 Stream ships 25.11 [secondary].
  Source: https://packages.debian.org/sid/dpdk
  Source: https://rpmfind.net/linux/rpm2html/search.php?query=dpdk
- DPDK 25.11 release notes list test coverage on NICs from ConnectX-6 Dx up through **ConnectX-9 SuperNIC 800G** (15b3:1025, PCIe 6.0 x16, FW 82.47.0366+) and BlueField-2/BlueField-3 DPUs (Ubuntu 24.04 + MLNX_OFED 25.10 + bf-bundle 3.2.0) — the 2026 HCL reaches 800 GbE per port [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_25_11.rst
- **DPDK 25.07** highlights: full `-l` corelist-to-CPU mapping parity with `--lcores`; new **PMU library** (self-monitoring without `perf`); Amazon ENA PMD updates (fragment-bypass mode for egress PPS limits); new Mucse rnp 10GbE PMD; `RTE_ETH_RSS_IB_BTH` RSS for RoCE v2; burst-mode query on e1000/ixgbe/iavf [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_25_07.rst
- Security: CVE-2026-86564 (lib/vhost missing length validation) marked vulnerable across Debian DPDK versions with no DSA yet at search time; NVIDIA Mellanox PMD CVE-2025-23259 fixed in current releases [secondary].
  Source: https://security-tracker.debian.org/tracker/source-package/dpdk
- Known migration limit (still current): direct live migration of a VF-backed VM is not possible; DPDK documents the virtio+VF bond failover pattern (bond_virtio_sriov howto) for migrating between VF and virtio paths [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/howto/lm_bond_virtio_sriov.rst
- The CyPerf DPDK traffic-generator container (Keysight) pins **DPDK 25.11 LTS** for its 2026 agent images on x86_64 and aarch64 — a sign of LTS adoption in test tooling [secondary].
  Source: https://github.com/keysight/cyperf/blob/HEAD/deployment/containers/dpdk/README.md

### 5.2 SPDK 2026 releases
- **SPDK v26.01 LTS** (announced Jan 2026, supported 1 year): NVMe 2.0 features on the NVMe-oF target, RDMA transport interrupt support, preempt-and-abort reservations, expanded target passthrough (FW/security/sanitize/get-set-features), halved locking in bdev/bdev_nvme thread sync, hot-remove for uring+aio bdevs, write-uncorrectable command support, AMD AE4DMA DMA engine, CUDA accel module (XOR/FILL/COPY offload to GPUs), and a **DPDK 25.11** submodule [official].
  Source: https://github.com/spdk/spdk.github.io/blob/HEAD/_posts/2026-01-28-v26.01_release.md
- **SPDK v26.05** (May 2026): NVMe Key-Value command-set support in the driver, JSON-RPC schema generation as source of truth, JSON-RPC client batching, NVMe-oF duplicate-host policy (I/O fencing), global default multipath config, per-I/O-size buffer-pool selection on the NVMe-oF transport (bandwidth/latency win for large I/O), DSM/Write-Zeroes spec-compliance fixes, and a NUMA-aware vfio-user transport [official].
  Source: https://github.com/spdk/spdk.github.io/blob/HEAD/_posts/2026-05-28-v26.05_release.md
- SPDK performance report 26.01: NVMe-oF RDMA target benchmarked on **BlueField-3** (16× Arm A78, 2×200 GbE ConnectX-7 subsystem, 16× Samsung E3.S Gen5 SSDs, Spectrum-4 800 GbE switch) — DPU-hosted storage target validated at 26.01 release [official].
  Source: https://review.spdk.io/download/performance-reports/SPDK_rdma_nvda_perf_report_2601.pdf
- NVMe/TCP on BlueField-3 (July 2026, experimental): an SPDK-NVMe-oF-target-as-VPP-plugin (polled by VPP workers, no second process) beat SPDK/POSIX, SPDK/io_uring and SPDK/XLIO user-space backends on six matched workloads — 1.32× geometric-mean throughput vs XLIO, 1.96×/2.25× throughput-per-core vs POSIX/io_uring; **not production-ready**, evaluation only [secondary].
  Source: https://medium.com/@jerome.tollet/spdk-inside-vpp-accelerating-nvme-tcp-on-bluefield-3-bc0419048b1e
- Security note: CVE-2025-57275 (buffer overflow in SPDK 25.05 NVMe-oF target lib/nvmf, CVSS 5.5, disclosed Apr 2026) — users should be on v26.x LTS line [secondary].
  Source: https://app.opencve.io/cve/?vendor=spdk
- Positioning recap: SPDK = user-space, polled, lockless NVMe driver + block stack + NVMe-oF/iSCSI/vhost servers, up to an order of magnitude more CPU-efficient than kernel equivalents; standard Linux initiators and QEMU interoperate [official].
  Source: https://github.com/spdk/spdk.github.io/blob/HEAD/index.md

---
## Wave 6 — GPU virtualization: vGPU, MIG, SR-IOV graphics, and AI relevance

### 6.1 NVIDIA vGPU / MIG (2026 state)
- NVIDIA vGPU software train has reached **19.x** (19.0–19.6 documented; GCP publishes GRID drivers up to 19.4 = driver 580.126.09) — license tiers: **vWS** (RTX Virtual Workstation), **vPC** (Virtual PC), **vApps** (Virtual Applications) for graphics; **vCS / Virtual Compute Server** for compute, licensed exclusively through NVIDIA AI Enterprise [vendor-reported] [secondary].
  Source: https://docs.cloud.google.com/compute/docs/gpus/grid-drivers-table
  Source: https://forums.developer.nvidia.com/t/does-nvidia-ai-enterprise-include-vws-vpc-or-vapps-licenses-or-only-vcs-vgpu-for-compute/346037
- NVIDIA AI Enterprise covers the **compute (C-series) vGPU profiles only** — vWS/vPC/vApps for graphics/VDI require separate licenses and generally graphics-capable GPUs; H200 NVL + 5-year AI Enterprise bundle is a common compute-virtualization vehicle [secondary].
  Source: https://forums.developer.nvidia.com/t/does-nvidia-ai-enterprise-include-vws-vpc-or-vapps-licenses-or-only-vcs-vgpu-for-compute/346037
- Release 19.0–19.6 supports Ada Lovelace SR-IOV GPUs (RTX 6000 Ada, RTX 5880 Ada, RTX 5000 Ada), L40S/L40/L20/L4/L2, and RTX PRO 6000 Blackwell Server Edition with vWS/vPC/vApps profiles in time-sliced and (where noted) DDA modes [vendor-reported].
  Source: https://docs.nvidia.com/vgpu/19.0/pdf/grid-vgpu-release-notes-microsoft-windows-server.pdf
- **MIG (Multi-Instance GPU)**: hardware partitioning on A100/H100-class GPUs; a guest VM **cannot create/destroy/reconfigure MIG instances** — the vGPU host manages them, guest `nvidia-smi -mig 0` returns "Insufficient Permissions"; K8s GPU Operator sets `migManager.enabled=false` on such hosts [secondary].
  Source: https://github.com/microsoft/physical-ai-toolchain/blob/HEAD/docs/reference/gpu-configuration.md
- MIG-backed vGPU profiles (1:1 MIG-backed = one vGPU per GPU instance) coexist with time-sliced vGPU modes on supported GPUs [vendor-reported].
  Source: https://docs.nvidia.com/vgpu/19.0/pdf/grid-vgpu-release-notes-microsoft-windows-server.pdf
- Confidential-computing angle: NVIDIA's confidential-containers reference architecture supports H100/H200 (incl. Protected PCIe) and B200 in passthrough for confidential AI — vGPU/MIG interplay with confidential GPU mode is a buyer question to verify per SKU [vendor-reported].
  Source: https://github.com/nvidia/cloud-native-docs/blob/HEAD/confidential-containers/supported-platforms.rst

