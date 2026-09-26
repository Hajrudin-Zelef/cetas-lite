---
id: etape6-phasef4-virtualization-io/00-virtualization-io/1-5-virtio-in-the-public-clouds-aws-gcp-and-the-proprietary-
title: "1.5 virtio in the public clouds — AWS, GCP, and the proprietary-NIC trend"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AWS", "Google", "Intel"]
dates: []
keywords: ["aws", "benchmark", "compute", "ethernet", "gpu", "intel", "latency"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [51, 75]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 1569d06d7e858926974ac65abbb9b54aaddc39c7156d3189bec39b0a10d2cc89
---

# 1.5 virtio in the public clouds — AWS, GCP, and the proprietary-NIC trend

### 1.5 virtio in the public clouds — AWS, GCP, and the proprietary-NIC trend
- **AWS:** EC2 enhanced networking is based on the proprietary Elastic Network Adapter (ENA), "not compatible with the Virtio specification"; ENA itself is SR-IOV-based, up to 100 Gbps on supported instance types, no extra charge [official].
  Source: https://docs.aws.amazon.com/en_us/AWSEC2/latest/UserGuide/enhanced-networking.html
- Consequence noted by Red Hat: a container/VM image built directly against the ENA PMD is coupled to AWS; a "translator" mediator layer (virtio ring data plane + vhost-user control plane) is needed to present a standard virtio interface on ENA-backed hosts [secondary].
  Source: https://www.redhat.com/it/blog/making-high-performance-networking-applications-work-hybrid-clouds
- For HPC/AI, AWS offers the Elastic Fabric Adapter (EFA) with OS-bypass (libfabric) for MPI workloads — ENA for general traffic, EFA for low-latency collectives [official].
  Source: https://docs.aws.amazon.com/en_us/AWSEC2/latest/UserGuide/enhanced-networking.html
- **GCP:** Google Virtual NIC (gVNIC) is "an alternative to the virtIO-based ethernet driver" and "the next generation network interface which succeeds VirtIO"; it is the only supported network interface for Generation 3+ machine types, required for 200 Gbps Tier_1 networking [official].
  Source: https://docs.cloud.google.com/compute/docs/networking/using-gvnic
- GCP benchmark (c2-standard-60, HPC VM image): gVNIC at 100 Gbps delivered ~57% higher MPI bandwidth on average vs VirtIO; WRFv3 +51%, ANSYS Fluent +13%, LS-DYNA +11% vs Virtio-Net [vendor-reported].
  Source: https://cloudsteak.com/gcp-accelerating-mpi-applications-using-google-virtual-nic-gvnic/
- GCP bare-metal instances use IDPF (Intel's Infrastructure Data Path Function); H4D-class machines optionally use Cloud RDMA (IRDMA) alongside gVNIC, with RDMA traffic prioritized [official].
  Source: https://docs.cloud.google.com/compute/docs/networking/network-overview
- Trend reading: all three hyperscalers now push proprietary/guest-visible NIC interfaces (ENA, gVNIC, Azure MANA) ahead of plain virtio-net for performance tiers; virtio remains the portability baseline and the lingua franca for private cloud / KVM / edge [secondary].

### 1.6 virtio ecosystem notes
- Google's crosvm implements a full virtio device set: net, block, console, fs (FUSE), gpu, input, iommu, p9, pmem, rng, scsi, snd, tpm (vTPM-backed), video, vsock, wayland, plus vhost-user offload to another process [secondary].
  Source: https://github.com/google/crosvm/blob/HEAD/docs/book/src/devices/index.md
- virtio-scsi (device ID 8) is the standard paravirtual SCSI HBA path; virtio-blk (ID 2) now documents runtime capacity change with config-change notification (behavior in Linux/QEMU since 2011) [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.3/virtio-v1.3.html
- Emerging device: virtio IOMMU (ID 23) — paravirtual IOMMU for nested/assigned-device DMA control inside guests [official].
  Source: http://docs.oasis-open.org/virtio/virtio/v1.0/cs02/virtio-v1.0-cs02.pdf
- [unverified] Whether virtio 1.4 will reach Committee Specification / OASIS Standard status in 2026 is unknown; the TC's own status thread describes voting and mailing-list infrastructure as broken with no ETA.

---
