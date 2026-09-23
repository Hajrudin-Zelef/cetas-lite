---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-11-dpdk-security-offloads-and-nvme-of-transport-selecti
title: "Wave 11 — DPDK security offloads and NVMe-oF transport selection"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AWS", "Intel", "United States"]
dates: ["2025-06", "2026-05", "2026-07"]
keywords: ["accelerator", "aws", "cost", "ethernet", "hyperscaler", "intel", "latency", "memory", "research", "throughput"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [622, 689]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 155ba88b41735dbbf082a36de79d09e7db93474e416f9c64de37b10ae1f0e5e4
---

# Wave 11 — DPDK security offloads and NVMe-oF transport selection

## Wave 11 — DPDK security offloads and NVMe-oF transport selection

### 11.1 rte_security / cryptodev / compressdev
- DPDK's `rte_security` framework models three offload styles for IPsec: **lookaside-protocol** (whole SA offloaded to a crypto accelerator via cryptodev), **inline-protocol** (offload on the Ethernet device itself), and **inline-crypto** (crypto processing inline on the ethdev during TX); the ipsec-secgw sample app supports all four modes including `no-offload` [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/prog_guide/rte_security.rst
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/sample_app_ug/ipsec_secgw.rst
- Inline path mechanics: egress SADB lookup marks the packet, the NIC PMD sets the hardware crypto context, and the NIC adds ESP/tunnel headers + encryption/authentication inline; post-encryption TSO is supported on capable devices; a `fallback lookaside-none` session catches packets the inline engine failed to process (with widened anti-replay windows to absorb inline/lookaside reordering) [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/prog_guide/rte_security.rst
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/sample_app_ug/ipsec_secgw.rst
- Hardware coverage notes: Marvell OCTEON TX2 PMD added full inline IPsec (SAD lookup + decrypt/transform in hardware, plain-packet submission on egress); Intel QAT PMD queue-pairs are thread-safe on Intel CPUs (TX enqueue and RX dequeue may use different threads) [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_20_02.rst
- **compressdev**: DPDK 21.02 added an **mlx5 compress PMD for BlueField-2** adapters — compression offload on the DPU, relevant to storage/VNF pipelines that compress before NVMe-oF or WAN transport [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_21_02.rst
- Relevance to I/O virtualization: inline IPsec + inline crypto on the NIC/DPU is the same offload family as SR-IOV/VF RSS/flow-steering — it moves per-packet transforms off the vCPU, which is why 5G UPF and SD-WAN VNF designs pair DPDK PMDs with inline-capable NICs [secondary].

### 11.2 NVMe-oF transport selection: RDMA vs TCP
- SPDK performance reports (24.05, ConnectX-5, QD=1, null block device): SPDK NVMe-oF **RDMA** initiator avg latency ~4.06–4.72 µs vs kernel initiator ~11.1–12.1 µs (SPDK cuts initiator overhead ~50%); SPDK **TCP** target+initiator cut average latency up to 7.38 µs vs kernel NVMe/TCP — eliminating up to 61% of NVMe-oF software overhead [official].
  Source: https://olo.0l0.workers.dev:443/https/review.spdk.io/download/performance-reports/SPDK_rdma_mlx_perf_report_2405.pdf
- SPDK 23.01 (RoCEv2): SPDK RDMA target −2.23 µs avg round-trip vs kernel target (~11.3% overhead cut); SPDK initiator −4.96 µs (~28% overhead cut) [official].
  Source: https://olo.0l0.workers.dev:443/https/review.spdk.io/download/performance-reports/SPDK_rdma_cvl_roce_perf_report_2301.pdf
- Why TCP exists anyway (Lightbits/StackConf 2025): RDMA needs special networks/hardware; TCP is ubiquitous, scales to large topologies and long distances, and NVMe/TCP can approach direct-attached NVMe latency/throughput — in clouds, network bandwidth often exceeds provider-attached storage bandwidth [vendor-reported].
  Source: https://www.slideshare.net/slideshow/stackconf-2025-how-nvme-over-tcp-runs-postgresql-in-quicksilver-mode-by-sagy-volkov-pdf/278786270
- TCP caveats (Blocks&Files): NVMe/TCP latency runs a few microseconds above RDMA; head-of-line blocking and incast add latency; no hardware acceleration — fine for iSCSI migratees, questionable for the most latency-sensitive workloads [secondary].
  Source: https://www.blocksandfiles.com/block/2019/02/05/nvme/tcp-needs-good-tcp-network-design/1601407
- Independent 2026 test harness (ioutgt, July 2026): SPDK NVMe-oF target vs in-kernel `nvmet` on the same two-NIC wire, TCP and RDMA branches — TCP 466k IOPS loopback, RDMA 35k IOPS over rxe in their VM env; notes that `intel_iommu=on` blocks SPDK/DPDK DMA memory in VMs without VFIO (runner disables IOMMU for that env) [independent].
  Source: https://github.com/io-target/ioutgt/commit/b4e6fcae7eb815a1c09fd16c3de4a57448dd2e1e
- Selection rule: RDMA/RoCE for the lowest latency on lossless fabrics; NVMe/TCP for ubiquity, cloud, and long-distance; SPDK on both sides beats kernel initiator/target on software overhead in SPDK's own reports — corroborate with your own fio runs [secondary].

---

## Part F (continued 2) — additional source URLs
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/prog_guide/rte_security.rst
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/sample_app_ug/ipsec_secgw.rst
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_20_02.rst
- https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_21_02.rst
- https://olo.0l0.workers.dev:443/https/review.spdk.io/download/performance-reports/SPDK_rdma_mlx_perf_report_2405.pdf
- https://olo.0l0.workers.dev:443/https/review.spdk.io/download/performance-reports/SPDK_rdma_cvl_roce_perf_report_2301.pdf
- https://www.slideshare.net/slideshow/stackconf-2025-how-nvme-over-tcp-runs-postgresql-in-quicksilver-mode-by-sagy-volkov-pdf/278786270
- https://www.blocksandfiles.com/block/2019/02/05/nvme/tcp-needs-good-tcp-network-design/1601407
- https://github.com/io-target/ioutgt/commit/b4e6fcae7eb815a1c09fd16c3de4a57448dd2e1e
- https://intelligentvisibility.com/blog/ethernet-storage-protocols-nvme-nfs-iscsi

---

## Wave 12 — Hyperscaler NIC programs (2026) and virtio-net feature reference

### 12.1 AWS ENA Express / SRD (2026 update)
- **ENA Express** layers AWS's proprietary **Scalable Reliable Datagram (SRD)** transport under ENA: single-flow bandwidth rises from 5 Gbps to **25 Gbps**, with up to **85% P99.9 latency improvement** for high-throughput workloads; congestion control, multi-pathing and packet reordering run on the Nitro card, transparent to guest TCP/UDP [vendor-reported].
  Source: https://aws.amazon.com/about-aws/whats-new/2022/11/elastic-network-adapter-ena-express-amazon-ec2-instances/
- May 2026: ENA Express now works **between Availability Zones** in a Region (previously same-AZ only) — 25 Gbps single-flow cross-AZ; EBS io2 Block Express and EFA (HPC/ML) also ride on SRD [vendor-reported].
  Source: https://aws.amazon.com/about-aws/whats-new/2026/05/ena-express-availability-zones/
- June 2025: ENA Express extended to AWS GovCloud (US) regions at no additional cost [vendor-reported].
  Source: https://aws.amazon.com/about-aws/whats-new/2025/06/ena-express-govcloud-us-regions/
- How SRD works (ipSpace analysis): reliable datagram transport like UDP but reliable, unlike TCP it **permits out-of-order delivery** — dropping in-order delivery lets AWS spray packets across parallel paths, killing head-of-line blocking and link congestion; the Nitro card reorders on receive before handing to the guest stack [secondary].
  Source: https://blog.ipSpace.net/2022/12/quick-look-aws-srd/
- Buyer note: for pure PPS/latency-at-low-load workloads, plain ENA enhanced networking may still fit better; ENA Express cannot be used in a Local Zone [vendor-reported].
  Source: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ena-express.html
- DPDK-side tie-in: DPDK 25.07's ENA PMD gained fragment-bypass mode for egress (bypassing EC2's per-ENI PPS cap on fragmented packets) — the PMD tracks Nitro behavior closely [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/rel_notes/release_25_07.rst

### 12.2 virtio-net feature bits quick reference (spec 1.x)
- Core negotiated features: `VIRTIO_NET_F_CSUM` (partial checksum offload), `VIRTIO_NET_F_GUEST_CSUM/TSO/ECN/UFO`, `VIRTIO_NET_F_MAC`, `VIRTIO_NET_F_GUEST_ANNOUNCE`, `VIRTIO_NET_F_MQ` (multiqueue, queue pairs = 2N), `VIRTIO_NET_F_CTRL_VQ` (the third control queue for MAC/VLAN/RX-mode programming), `VIRTIO_NET_F_CTRL_RX/VLAN/MAC_ADDR`, `VIRTIO_NET_F_MRG_RXBUF` (mergeable buffers), `VIRTIO_NET_F_STATUS`, `VIRTIO_NET_F_SPEED_DUPLEX`, `VIRTIO_NET_F_RSS`, `VIRTIO_NET_F_HASH_REPORT`, `VIRTIO_NET_F_NOTIF_COAL` [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- vDPA/virtio-pmd implication: a DPDK virtio PMD or hardware vDPA NIC negotiates the same feature bits — feature parity between the paravirtual guest driver and the backend is what keeps live migration safe across hypervisor upgrades [secondary].
- [unverified] Azure MANA specifics were not captured in this research round — treat the hyperscaler-NIC comparison as AWS+GCP-weighted; a MANA pass belongs in a future refresh.

---

