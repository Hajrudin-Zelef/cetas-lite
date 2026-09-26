---
id: etape6-phasef4-virtualization-io/00-virtualization-io/overview
title: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Intel", "Nvidia"]
dates: ["2026-04", "2026-09-22"]
keywords: ["amd", "aws", "cost", "gpu", "intel", "memory", "nvidia", "research", "throughput"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [1, 50]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: aaac209314094a8e7a7186ab4d8e4be94eaf77a7c3e5ef56426e89062e09c6fa
---

# Phase F4 — I/O Virtualization & CPU Acceleration Extensions

**Scope:** virtio (net/blk/scsi, spec 1.2/1.3, vhost), SR-IOV, Intel VT-x/VT-d & AMD-V/AMD-Vi (IOMMU, interrupt remapping, ATS/PRI/PASID), AVX-512 / AVX10 / AMX, DPDK & SPDK, GPU virtualization (NVIDIA vGPU, MIG, AMD MxGPU, Intel).
**Date / cutoff:** 2026-09-22. **Method:** read-only web research (browser_search, browser_open); no live-browser visits; nothing sent externally.
**Provenance legend:** [official] = vendor/standards-body primary source · [vendor-reported] = vendor claim not independently verified · [independent] = third-party measurement/review · [secondary] = press/analyst/blog reporting · [unverified] = single-source or uncorroborated claim.
**Conventions:** no SKUs, URLs, or version numbers are guessed — every identifier is verbatim from a cited source. Gaps, conflicts, and unverified claims are flagged inline and consolidated in the Wave 7 open-items log.

---

## Wave 1 — virtio: spec, transports, and cloud reality

### 1.1 What virtio is and where it sits
- virtio is the OASIS-standardized paravirtualized I/O device interface: instead of emulating real hardware, the guest uses a purpose-built "virtual device" front-end driver (e.g. virtio-net) that talks to a host-side back-end over shared-memory virtqueues [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.3/virtio-v1.3.html
- Design goal per the spec: "virtual environments and guests should have a straightforward, efficient, standard and extensible mechanism for virtual devices, rather than boutique per-environment or per-OS mechanisms" [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.3/virtio-v1.3.html
- Core mechanics: descriptor rings (split virtqueue, plus optional packed virtqueue), feature negotiation (VIRTIO_F_* bits), config space, notifications/used-buffer signaling, indirect descriptors, event suppression [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.3/virtio-v1.3.html
- Transports standardized: virtio over PCI (modern PCI transport, vendor ID 0x1AF4, non-transitional device IDs 0x1040–0x107F = 0x1040 + virtio device ID), virtio over MMIO, virtio over channel I/O (s390) [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.3/virtio-v1.3.html

### 1.2 Specification status as of 2026-09-22 — the 1.3/1.4 situation
- **Virtio 1.3 was never formally released.** The 1.4 working draft (virtiov1.4-cs01, dated 8 April 2026) explicitly states: "Since 1.3 was never released, keep the last version of 1.2 for diff generation" [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- A v1.3 CSD01 draft document exists on docs.oasis-open.org, but per the OASIS VIRTIO TC mailing list (Cornelia Huck, Red Hat), OASIS infrastructure migration problems (dead mailing lists, broken voting scripts) stalled any formal 1.3 release and delayed 1.4 progress [official].
  Source: https://groups.oasis-open.org/discussion/current-state-of-the-infrastructure-for-the-virtio-tc
- **Virtio 1.4 cs01** (Standards Track Work Product, 30 March / 8 April 2026, 440–444 pages) is the latest published TC output; changelog entries run through Nov 2025 and are largely editorial (label fixes, chair/editor updates) [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- Notable 1.4-era content: virtio-net IPsec extension ("virtionet: extend virtio_net_hdr for IPsec support", fd15f89a870f) — IPsec offload SA resource objects in the net header [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- Practical consequence: the industry still implements "virtio 1.0/1.1/1.2 + selected 1.3/1.4 draft features"; there is no shipping "virtio 1.4 final" standard as of the cutoff [unverified — inferred from TC state; no formal 1.4 release announced].

### 1.3 Device types (device IDs, v1.1 cs01 table)
- 1 = network card, 2 = block device, 3 = console, 4 = entropy source, 5/13 = memory ballooning, 8 = SCSI host, 9 = 9P transport, 16 = GPU device, 18 = input device, 19 = socket (vsock), 20 = crypto device, 23 = IOMMU device, 24 = memory device, 25 = audio, 26 = filesystem (virtio-fs), 27 = PMEM, 33 = NitroSecureModule (AWS Nitro Enclaves) [official].
  Source: http://commondatastorage.googleapis.com/chromeos-localmirror/distfiles/virtio-v1.1-cs01.pdf (v1.1 cs01 device table, page 71–72)
- virtio-net is described in the spec as "the most complex of the devices supported so far": RX/TX virtqueues plus a third control queue for advanced filtering (MAC/VLAN/RX-mode programming) [official].
  Source: http://docs.oasis-open.org/virtio/virtio/v1.0/cs02/virtio-v1.0-cs02.pdf

### 1.4 Back-end architectures: vhost-net vs vhost-user
- **vhost-net (kernel back-end):** the host kernel handles the data plane; guest uses the in-kernel virtio-net driver and gets a standard Linux netdev. Simple to operate, kernel in the path [secondary].
  Source: https://www.redhat.com/ja/blog/how-vhost-user-came-being-virtio-networking-and-dpdk
- **vhost-user (userspace back-end):** the back-end runs in host userspace (e.g. OVS-DPDK), communicating with the guest front-end over a Unix socket + shared memory. Guest side can use virtio-pmd (DPDK poll-mode driver) for a fully userspace path [secondary].
  Source: https://www.redhat.com/ja/blog/how-vhost-user-came-being-virtio-networking-and-dpdk
- Red Hat's analysis: moving both host and guest data planes to userspace (vhost-user + virtio-pmd) improves throughput "by a factor of 2 to 4" versus kernel vhost-net/virtio-net, at the cost of usability (guest apps must link the DPDK virtio-pmd; dedicated poll cores required) [vendor-reported].
  Source: https://www.redhat.com/ja/blog/how-vhost-user-came-being-virtio-networking-and-dpdk
- Historical DPDK PVP numbers (Release 18.02/18.05, Intel testpmd, 1 core vhost-user + 1 core virtio): 64-byte packets 6.12–6.93 Mpps; 1518-byte 2.10–2.64 Mpps — old silicon/QEMU 2.x, not comparable to 2026 platforms [independent — dated 2018, kept for historical reference only].
  Source: http://fast.dpdk.org/doc/perf/DPDK_18_05_Intel_virtio_performance_report.pdf
- Current DPDK still maintains vhost/virtio PVP multi-path single-core test plans (mergeable, non-mergeable, in-order, vectorized Rx paths) in its DTS suite — evidence the vhost-user path remains the performance reference for NFV [official].
  Source: http://doc.dpdk.org/dts/test_plans/pvp_multi_paths_vhost_single_core_performance_test_plan.html

