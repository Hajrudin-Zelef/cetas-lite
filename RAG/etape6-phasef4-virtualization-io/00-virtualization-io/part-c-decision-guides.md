---
id: etape6-phasef4-virtualization-io/00-virtualization-io/part-c-decision-guides
title: "Part C — Decision guides"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Google", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "aws", "benchmark", "blackwell", "compute", "fp8", "gpu", "inference", "intel", "license", "nvidia"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [306, 358]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 0906e9cb57f6d5766cf7c78ae5c550e1c6d12e1a6e301e67285ba3a355e45dea
---

# Part C — Decision guides

## Part C — Decision guides

### C.1 Choosing a virtual NIC path
| Workload | Recommended path | Why |
|---|---|---|
| Cloud VM (general) | virtio-net + vhost-net (kernel) | default, mature, live-migration friendly |
| NFV/vSwitch fast path | virtio-net + vhost-user + DPDK PMD | 2–4× vs kernel path, costs poll cores |
| AWS | ENA (SR-IOV) | native, up to 100 Gbit/s documented |
| GCP Gen3+ | gVNIC | required for 200 Gbit/s tiers |
| VM needing hardware QoS | SR-IOV VF passthrough | line-rate, but no live migration |

### C.2 Choosing an IOMMU/CPU virtualization posture
- Always enable VT-d/AMD-Vi in BIOS for any passthrough or vhost-user/DMA workload (documented requirement across DPDK, OpenStack, NVIDIA vGPU) [secondary].
- Prefer posted-interrupt capable platforms for high-interrupt-rate guests (networking, NVMe) [secondary].
- For CXL-attached accelerators, verify ATS/PRI/PASID support in the device *and* the IOMMU driver — PASID without PRI risks I/O page faults stalling the device [secondary].

### C.3 SIMD/ISA selection for data-plane and AI-on-CPU
| CPU | Vector ISA | Matrix ISA | Note |
|---|---|---|---|
| Ice Lake-SP | AVX-512 (no BF16/FP16) | — | legacy vector baseline |
| Sapphire Rapids | AVX-512 + BF16 | AMX INT8/BF16 | first AMX |
| Emerald Rapids | same as SPR | same as SPR | network SKUs (-N) for 5G/edge |
| Granite Rapids | AVX-512 + BF16 | AMX + FP16 | Xeon 6 |
| Diamond Rapids (Xeon 7) | **AVX10.2** + APX | AMX + FP8 | disclosed Hot Chips 2026, no SMT |
- Treat Skylake-SP AVX-512 throttling anecdotes as historical; validate frequency behavior on the target SKU stepping [secondary].

### C.4 DPDK vs kernel vs SPDK
| Need | Pick |
|---|---|
| Packet processing / vSwitch / 5G UPF | DPDK 25.11 LTS + matching PMD |
| Block storage target / NVMe-oF | SPDK v26.01 LTS (pins DPDK 25.11) |
| Container CNF data plane | DPDK in container, hugepages + CPU pinning |
| DPU-offloaded storage | SPDK NVMe-oF target on BlueField-3 |

### C.5 GPU sharing selection
| Need | Pick |
|---|---|
| AI inference slicing | NVIDIA vCS/MIG (C profiles) via AI Enterprise |
| VDI / graphics workstation | vWS/vPC (graphics-capable GPU + separate license) |
| Open-standards GPU sharing | AMD MxGPU SR-IOV line |
| Intel client/iGPU sharing | SR-IOV VFs on Xe+ (GVT-g is dead) |

---

## Part D — Conflicts, gaps, and open items
1. **ConnectX VF counts**: 127 VFs/port (docs) vs 1K VFs/port (ConnectX-6 Dx brief) — generational difference; confirm per-SKU in the product brief before sizing [vendor-reported].
2. **Diamond Rapids availability**: disclosed at Hot Chips Aug 2026; GA date and final SKUs not confirmed in collected sources — tag all DMR specs as pre-GA [unverified].
3. **virtio 1.3**: no formal publication found; 1.4 cs01 diffs compute from 1.2 — spec-version citations should say "1.2/1.4" not "1.3" [official].
4. **MIG on Blackwell**: not confirmed in collected sources [unverified].
5. **gVNIC internals**: Google documents behavior, not the wire format — no independent benchmark of gVNIC vs ENA found [unverified].
6. **DPDK 25.11 exact release date**: documented as LTS and packaged, but the release-day stamp was not captured [unverified].
7. AMD AVX10/APX adoption: no announcement found [unverified].

