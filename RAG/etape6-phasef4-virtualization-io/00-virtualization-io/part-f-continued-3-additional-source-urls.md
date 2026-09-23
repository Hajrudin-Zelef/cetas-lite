---
id: etape6-phasef4-virtualization-io/00-virtualization-io/part-f-continued-3-additional-source-urls
title: "Part F (continued 3) — additional source URLs"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Broadcom", "Intel", "Nvidia"]
dates: ["2026-04", "2026-09-22"]
keywords: ["amd", "aws", "benchmarks", "blackwell", "cost", "gpu", "intel", "latency", "memory", "nvidia", "research"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [690, 751]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 2161f9a57e68fdbde5100fa48cc0e8a3c814b1b103e3db94d03cfa1ba26f3cc0
---

# Part F (continued 3) — additional source URLs

## Part F (continued 3) — additional source URLs
- https://aws.amazon.com/about-aws/whats-new/2022/11/elastic-network-adapter-ena-express-amazon-ec2-instances/
- https://aws.amazon.com/about-aws/whats-new/2026/05/ena-express-availability-zones/
- https://aws.amazon.com/about-aws/whats-new/2025/06/ena-express-govcloud-us-regions/
- https://blog.ipSpace.net/2022/12/quick-look-aws-srd/
- http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ena-express.html

---

## Wave 13 — SPDK vhost targets: vhost-blk / vhost-scsi for VMs

### 13.1 What the SPDK vhost target is
- The SPDK `vhost` application is a **vhost-user slave server** exposing Unix-domain sockets; QEMU (or any vhost-user master) connects and gets virtualized storage devices backed by SPDK bdevs — vhost-blk (QEMU ≥ 2.12) and vhost-scsi (QEMU ≥ 2.10); vhost-user-nvme was deprecated in SPDK 21.07 [official].
  Source: https://spdk.io/doc/vhost.html
  Source: https://github.com/microsoft/kata-containers/blob/HEAD/docs/use-cases/using-SPDK-vhostuser-and-kata.md
- Wiring: QEMU needs shared hugepage memory (`-object memory-backend-file,...,share=on`), a `-chardev socket` per vhost controller, and `-device vhost-user-blk-pci` / `vhost-user-scsi-pci`; SPDK side: `build/bin/vhost -S /var/tmp -m 0x3` (poll cores fully occupied), then `rpc.py bdev_malloc_create` / `bdev_nvme_attach_controller` plus `vhost_create_blk_controller`/`vhost_create_scsi_target` to expose bdevs [official].
  Source: https://github.com/spdk/spdk/blob/HEAD/doc/vhost.md
- Guest requirement: virtio-blk or virtio-scsi drivers in the guest (in-box on Linux/FreeBSD; virtio-win drivers needed on Windows); tested with Ubuntu, Fedora, Windows [official].
  Source: https://spdk.io/doc/vhost.html
- This is the storage twin of the vhost-user networking story: steady-state I/O avoids VM exits and the kernel, at the cost of pinned poll cores and shared-memory trust between hypervisor and target [secondary].

### 13.2 Security note (2026 audit)
- A 2026 independent audit of **SPDK v26.05 vhost-user-blk/scsi** documents the TOCTOU reality of vhost-user: the guest writes descriptor rings/indirect tables in shared memory while SPDK's poller parses them without copying to private memory first — a malicious/compromised guest racing a sibling vCPU can change validated fields; impact ceiling is OOB read/write inside the privileged SPDK target process [independent].
  Source: https://github.com/xoreaxeaxeax/schrodingers-toctou/blob/HEAD/observer-effect/audits/audit-spdk-v26.05.md
- Operational takeaway: treat the vhost-user backend as part of the trust boundary — run one SPDK vhost target per tenant or per trust domain where feasible, and track upstream hardening [secondary].

---

## Part F (continued 4) — additional source URLs
- https://spdk.io/doc/vhost.html
- https://github.com/spdk/spdk/blob/HEAD/doc/vhost.md
- https://github.com/microsoft/kata-containers/blob/HEAD/docs/use-cases/using-SPDK-vhostuser-and-kata.md
- https://github.com/xoreaxeaxeax/schrodingers-toctou/blob/HEAD/observer-effect/audits/audit-spdk-v26.05.md

---

## Part J — Delivery notes and known thin spots
- This file is the Phase F4 research deliverable for Step 6 ("Virtualization & I/O"): virtio, SR-IOV, CPU/IOMMU virtualization, SIMD/ISA (AVX-512/AVX10/AMX), DPDK/SPDK, GPU virtualization. Written in English; every factual claim carries one of the five allowed provenance tags.
- Corrections applied during the final QC pass (2026-09-22): removed all non-allowlisted tags (analysis→secondary, gap→unverified, conflict-flagged→official/vendor-reported); fixed a corrupted character in the IOMMU section; corrected the virtio v1.1 cs01 device-table citation to the actual v1.1 document URL.
- Deliberate version-status correction inside the file: Wave 7 supersedes Wave 1's "no formal 1.4 publication" wording — OASIS virtio 1.4 cs01 was published 8 April 2026 (csprd01 9 Dec 2025); 1.3 was the version never formally published.
- Known thin spots for a future refresh: Azure MANA specifics; exact DPDK 25.11 release-day stamp; Diamond Rapids GA date and final SKU list; MIG on Blackwell B200 (community-calculator data only); AMD AVX10/APX adoption plans; gVNIC internals/wire format; virtio-fs DAX stability status.
- Benchmarks cited are era-labeled; do not compare 2018 DPDK PVP numbers with 2026 SPDK/DPDK reports. Vendor claims (ConnectX VF counts, ENA Express latency, SPDK overhead reductions) are tagged as vendor-reported and should be re-validated on target hardware.
- File statistics at delivery: 750+ lines, ~85 KB, 60+ verbatim source URLs in Part F. Research cut-off: 2026-09-22.

---

## Appendix K — DPDK PMD / device quick lookup (all names sourced from DPDK docs)
| PMD / device | Vendor / family | Notes |
|---|---|---|
| mlx5 | NVIDIA ConnectX-4 → ConnectX-9, BlueField-2/3 | PF, VF, SF, eSwitch, representors; 800G on CX-9 in 25.11 HCL |
| ice | Intel E810 | QSFP 100G-class, DPDK 25.11 tested |
| iavf | Intel E810 VF (AVF) | VF PMD for SR-IOV guests |
| ena | Amazon ENA/EC2 | fragment-bypass egress mode in 25.07 |
| bnxt | Broadcom NetXtreme | PF + VF, flow bifurcation, vfio-pci/uio |
| rnp | Mucse 10GbE | new in DPDK 25.07 |
| virtio | virtio-net PMD | guest-side DPDK endpoint, vhost-user peer |
| vhost | vhost-user lib/PMD | host-side backend for virtio guests |
| crypto_qat | Intel QAT | thread-safe queue-pairs, IPsec lookaside |
| compress_mlx5 | BlueField-2 | compressdev offload |
| octeontx2 | Marvell OCTEON TX2 | inline IPsec, crypto lookaside |

*End of file.*
