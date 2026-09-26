---
id: etape6-phasef4-virtualization-io/00-virtualization-io/6-2-amd-mxgpu-and-intel-gpu-virtualization
title: "6.2 AMD MxGPU and Intel GPU virtualization"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: hardware
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2026-01-28"]
keywords: ["amd", "gpu", "intel", "blackwell", "claude", "compute", "gpus", "inference", "memory", "nvidia", "nvlink", "training"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [284, 305]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 22c5592ba22a5d0f02f264015afc54a2009be4786ea3c417f65308b9ed265ee5
---

# 6.2 AMD MxGPU and Intel GPU virtualization

### 6.2 AMD MxGPU and Intel GPU virtualization
- AMD **MxGPU** (launched on FirePro S7150/S7150 x2, up to 16/32 users per card): the industry's first **SR-IOV-based GPU sharing** — hardware GPU scheduling with QoS plus hardware-enforced memory isolation between VMs; exposes full DirectX/OpenGL/OpenCL to guests [vendor-reported].
  Source: http://hexus.net/tech/items/graphics/90230-amd-reveals-worlds-first-hardware-virtualized-gpu-product-line/
  Source: https://www.networkworld.com/article/968586/what-is-sr-iov-and-why-is-it-the-gold-standard-for-gpu-sharing.html
- Intel **GVT-g** (mediated passthrough, Broadwell→Comet Lake via kvmgt/mdev): **officially discontinued** — Intel archived the repo with a "known security escapes" notice and ceased all development; Gen11+ (Ice Lake and later) never got GVT-g, only GVT-d full passthrough [official].
  Source: https://wiki.archlinux.org/title/Intel_GVT-g
- Intel's stated replacement for Xe (Gen12) and newer: **SR-IOV-based GPU virtualization** — VFs created on the iGPU/dGPU and assigned to VMs, requiring recent kernel/QEMU with IOMMU enabled [secondary].
  Source: https://wiki.archlinux.org/title/Intel_GVT-g
- BIOS prerequisites common to all GPU-sharing modes (VT-x/VT-d or AMD-V/AMD-Vi, IOMMU, SR-IOV enabled) are documented by hypervisor vendors and OpenStack Kayobe playbooks alike [secondary].
  Source: https://github.com/verge-io/docs-vergeos/blob/HEAD/learn/06-virtual-machines/03-gpu-passthrough.md
  Source: https://github.com/openstack/kayobe/blob/HEAD/doc/source/configuration/reference/vgpu.rst

### 6.3 AI-cluster relevance
- Compute virtualization (vCS/MIG/time-sliced C profiles) lets inference fleets right-size GPU slices per model instead of whole-GPU allocation; H100/H200/B200 GB200/GB300 appear in cloud as full-GPU instances with NVLink/NVSwitch fabrics rather than vGPU slices at the largest scales [secondary].
  Source: https://github.com/rochacbruno/gcd-claude-skill/blob/HEAD/docs-raw-s3ns/compute_gpus.md
- Storage+network convergence on AI clusters: SPDK NVMe-oF targets increasingly run **on DPUs** (BlueField-3 validated at SPDK 26.01), freeing host CPUs for training — DPDK and SPDK roadmaps track each other (SPDK v26.01 pins DPDK 25.11) [official].
  Source: https://review.spdk.io/download/performance-reports/SPDK_rdma_nvda_perf_report_2601.pdf
  Source: https://github.com/spdk/spdk.github.io/blob/HEAD/_posts/2026-01-28-v26.01_release.md
- [unverified] Whether NVIDIA will extend MIG to Blackwell B200-class GPUs in 2026/27 was not confirmed in sources collected here — check current vGPU release notes before sizing a MIG-based inference fleet.

---

