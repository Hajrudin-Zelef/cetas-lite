---
id: etape6-phasef4-virtualization-io/00-virtualization-io/part-e-glossary-key-terms
title: "Part E — Glossary (key terms)"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["Intel", "Nvidia"]
dates: ["2026-01-28", "2026-05-28", "2026-09-22"]
keywords: ["amd", "aws", "claude", "compute", "datacenter", "fp8", "gpu", "gpus", "intel", "licenses", "nvidia", "research"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [359, 428]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 48d0f733b28037fe7ca6552b6b8e513e7179337906b9d418aa15657db3c3689d
---

# Part E — Glossary (key terms)

## Part E — Glossary (key terms)
- **eSwitch**: embedded switch in the NIC that forwards between PF, VFs, and wire; exposed via switchdev representors.
- **PASID**: Process Address Space ID — lets a device share a process's address space through the IOMMU.
- **Posted interrupts**: IOMMU/CPU mechanism delivering device interrupts directly to a running vCPU, skipping hypervisor exits.
- **MIG**: Multi-Instance GPU — hardware partitioning of an NVIDIA datacenter GPU.
- **vhost-user**: userspace virtio backend over a Unix socket; the fast path for DPDK-based hypervisors.
- **Representor**: netdev standing in for a VF/PF on the host for control-plane programming.
- **AMX**: Advanced Matrix Extensions — Intel tile-based matrix engine (INT8/BF16/FP16/FP8 by generation).
- **AVX10.2**: converged 128/256/512-bit vector ISA running on both P- and E-cores.
- **APX**: Advanced Performance Extensions — 32 GPRs, 3-operand forms, binary-compatible.
- **PMD**: Poll Mode Driver — DPDK's userspace NIC driver model.
- **NVMe-oF**: NVMe over Fabrics — remote NVMe over RDMA/TCP transports.

## Part F — Source index (verbatim URLs)
- https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- https://groups.oasis-open.org/discussion/current-state-of-the-infrastructure-for-the-virtio-tc
- https://docs.oasis-open.org/virtio/virtio/v1.3/virtio-v1.3.html
- http://commondatastorage.googleapis.com/chromeos-localmirror/distfiles/virtio-v1.1-cs01.pdf
- https://www.redhat.com/ja/blog/how-vhost-user-came-being-virtio-networking-and-dpdk
- https://docs.aws.amazon.com/en_us/AWSEC2/latest/UserGuide/enhanced-networking.html
- https://docs.cloud.google.com/compute/docs/networking/using-gvnic
- https://docs.cloud.google.com/compute/docs/networking/network-overview
- https://docs.nvidia.com/networking/display/mlnxenv24040660/single+root+io+virtualization+(sr-iov)
- https://solutions.asbis.com/api/uploads/files/40/pb-connectx-6-dx-en-card.pdf
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/howto/lm_bond_virtio_sriov.rst
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/nics/mlx5.rst
- https://projectacrn.github.io/2.5/developer-guides/hld/hv-vt-d.html
- https://cdrdv2-public.intel.com/868911/D51397-018-vt-directed-io-spec.pdf
- https://github.com/microsoft/openvmm/pull/3697
- https://github.com/microsoft/openvmm/blob/HEAD/Guide/src/reference/openvmm/management/cli.md
- https://www.tomshardware.com/pc-components/cpus/intel-emerald-rapids-5th-gen-xeon-platinum-8592-review-64-cores-320mb-of-l3-and-350w-tdp
- https://github.com/sunstoneinstitute/horndb/blob/HEAD/crates/simd/simd-research-findings.md
- https://github.com/adaworldapi/ndarray/blob/HEAD/.claude/knowledge/td-simd-cpu-dispatch-matrix.md
- https://www.theregister.com/hpc/2026/08/25/intel-diamond-rapids-xeon-7-cpu-deep-dive/5292427
- https://www.eweek.com/news/intel-diamond-rapids-256-core-xeon-amd-venice/
- https://www.techtimes.com/articles/325660/20260826/diamond-rapids-disclosed-intel-xeon-7-packs-256-cores-gigabyte-cache-drops-smt.htm
- https://videocardz.com/newz/intel-documents-confirm-avx10-support-on-next-gen-nova-lake
- https://www.cryptopolitan.com/intels-5th-gen-xeon-avx-512-boosts/
- https://en.wikipedia.org/wiki/Emerald_Rapids
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_25_11.rst
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_25_07.rst
- https://security-tracker.debian.org/tracker/source-package/dpdk
- https://packages.debian.org/sid/dpdk
- https://rpmfind.net/linux/rpm2html/search.php?query=dpdk
- https://github.com/keysight/cyperf/blob/HEAD/deployment/containers/dpdk/README.md
- https://github.com/spdk/spdk.github.io/blob/HEAD/_posts/2026-01-28-v26.01_release.md
- https://github.com/spdk/spdk.github.io/blob/HEAD/_posts/2026-05-28-v26.05_release.md
- https://review.spdk.io/download/performance-reports/SPDK_rdma_nvda_perf_report_2601.pdf
- https://medium.com/@jerome.tollet/spdk-inside-vpp-accelerating-nvme-tcp-on-bluefield-3-bc0419048b1e
- https://app.opencve.io/cve/?vendor=spdk
- https://github.com/spdk/spdk.github.io/blob/HEAD/index.md
- https://docs.cloud.google.com/compute/docs/gpus/grid-drivers-table
- https://forums.developer.nvidia.com/t/does-nvidia-ai-enterprise-include-vws-vpc-or-vapps-licenses-or-only-vcs-vgpu-for-compute/346037
- https://github.com/microsoft/physical-ai-toolchain/blob/HEAD/docs/reference/gpu-configuration.md
- https://github.com/nvidia/cloud-native-docs/blob/HEAD/confidential-containers/supported-platforms.rst
- https://docs.nvidia.com/vgpu/19.0/pdf/grid-vgpu-release-notes-microsoft-windows-server.pdf
- http://hexus.net/tech/items/graphics/90230-amd-reveals-worlds-first-hardware-virtualized-gpu-product-line/
- https://www.networkworld.com/article/968586/what-is-sr-iov-and-why-is-it-the-gold-standard-for-gpu-sharing.html
- https://www.eetimes.com/amd-says-more-than-one-way-to-do-virtualization/
- https://github.com/verge-io/docs-vergeos/blob/HEAD/learn/06-virtual-machines/03-gpu-passthrough.md
- http://documentation.suse.com/sles/15-SP7/html/SLES-all/article-nvidia-vgpu.html
- https://github.com/openstack/kayobe/blob/HEAD/doc/source/configuration/reference/vgpu.rst
- https://wiki.archlinux.org/title/Intel_GVT-g
- https://github.com/rochacbruno/gcd-claude-skill/blob/HEAD/docs-raw-s3ns/compute_gpus.md

---
*Research completed 2026-09-22. virtio/SR-IOV/IOMMU sections reflect sources current to that date; AVX10/Diamond Rapids based on Hot Chips Aug 2026 disclosures; DPDK/SPDK on 2026 LTS lines.*

---

