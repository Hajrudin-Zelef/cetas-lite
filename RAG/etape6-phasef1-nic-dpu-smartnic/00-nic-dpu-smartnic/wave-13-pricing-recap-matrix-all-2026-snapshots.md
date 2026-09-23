---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-13-pricing-recap-matrix-all-2026-snapshots
title: "Wave 13 — Pricing recap matrix (all 2026 snapshots)"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["Google", "Intel", "Lambda", "Nvidia", "Qualcomm", "United States"]
dates: ["2025-06", "2026-09-15", "2026-09-22"]
keywords: ["pricing", "acquisition", "apache", "chiplet", "gpu", "intel", "nvidia", "research", "serdes"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [473, 536]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: fa4a2659a17350a097486124f92fd3a4b839efefe0183410d3756577d3d04452
---

# Wave 13 — Pricing recap matrix (all 2026 snapshots)

## Wave 13 — Pricing recap matrix (all 2026 snapshots)

| Product | Snapshot | Condition | Channel | Date | Tag |
|---|---|---|---|---|---|
| BF-2 MBF2M516A-CECOT | $2,299 (MSRP $3,143) | New | SHI | 2026-09 | [secondary] |
| BF-2 BF2M515A | $379.99 | Used | eBay | 2026-09 | [secondary] |
| BF-3 B3210E (900-9D3B4-00CC-EA0) | $3,053 (MSRP $4,686) | New | SHI commercial | 2026-09 | [secondary] |
| BF-3 B3210E (same SKU) | $4,425 | New | SHI public sector | 2026-09 | [secondary] |
| BF-3 B3210E (same SKU) | $4,040.60 | New, OOS | Compsource | 2026-09-15 | [secondary] |
| BF-3 B3240 | $3,999 | Used | LambdaTek/eBay-type | 2026-09 | [secondary] |
| BF-3 B3220 (Dell OEM) | $6,982–$7,229 (MSRP $11,579.78) | New | PC-Canada | 2026-09 | [secondary] |
| ConnectX-8 800G OSFP | $2,519 | New | FS.com US | 2026-09 | [secondary] |
| ConnectX-8 800G OSFP | £2,389.20 | New | FS.com UK | 2026-09 | [secondary] |
| ConnectX-8 2×400G QSFP112 | £2,450.40 | New | FS.com UK | 2026-09 | [secondary] |
| ConnectX-8 C8180/C8240 | $1,779 | New | NADDOD | 2026-09 | [secondary] |
| ConnectX-8 (3rd-party) | ~$2,500 | New (grey-market risk) | ecer.com | 2026-09 | [unverified] |
| BlueField-4 | — | — | — | 2026-09-22 | [gap — not found] |
| Pollara 400 | — | — | — | 2026-09-22 | [gap — not found] |
| Mount Evans / E2000 | — | — | — | 2026-09-22 | [gap — Google internal] |

---
*End of Phase F1. 13 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

## Wave 14 — DOCA release detail, P4 ecosystem, consolidation correction

### 14.1 DOCA 2.9.x release notes (official)

Per NVIDIA DOCA v2.9.0 / v2.9.1 documentation [official] [source: https://docs.nvidia.com/doca/archive/2-9-0/changes+and+new+features/index.html and https://docs.nvidia.com/doca/sdk/doca-sdk-v2-9-1.1.pdf ]:

- **DOCA 2.9.0 LTS branch is the final release supporting ConnectX-4** (support ends Jan 2025); subsequent DOCA-Host releases drop CX-4.
- **BlueField-3 requires DOCA-Host as the host driver — MLNX_OFED does not support BlueField-3** (important migration note for BlueField-2 → 3).
- DOCA 2.9.0 **does not support ConnectX-8** (CX-8 came in later DOCA branches).
- Added: Spectrum-X 1.2 reference architecture support; beta-level Quantum-X800 (XDR) support.
- OVN bare-metal: tenant isolation for E-W and N-S.
- DOCA Telemetry: high-frequency sampling, load-balancing between central observability DBs.
- DOCA Flow: performance-analysis tool, live upgrade, FDB RSS pipe in switch model; pkt_meta changed to big endian.
- GA: OVS-DOCA + local mirroring; DOCA Management Service (DMS).
- HBN: SimX support, HBN LTS, PBR, improved scale.
- Alpha: DOCA RDMA on GPU; BlueField-3 storage-controller target offload; beta SNAP virtio-fs.
- virtio-net: 1K device scale, optimizations, service restart flow for VFE, hotplug/unplug during host power-off; **virtio-net-controller HA** to minimize per-device downtime; virtio-blk improved recovery.
- Firefly service: SyncE, DTS (time sync).
- Storage/crypto: DOCA-Host new **DOCA-ROCE** profile.
- **PLDM firmware update support for BlueField-3**: firmware updates via PLDM type-5 over MCTP over PCIe through the platform BMC (~100 MB image); with this release PLDM flow requires full system power cycle [official].
- App Shield: PoC-level NIM container threat detection.

### 14.2 DOCA Platform Framework (DPF)

- DPF v25.04 / v25.1.1 (2025): provisions and orchestrates BlueField DPUs in Kubernetes — `DPUSet` CRD, `DPUService` CRDs (HBN, DTS telemetry, BlueMan), Helm charts from NVIDIA NGC registry; supports rolling updates with maxUnavailable 10% [official] [source: https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/release-notes/v25.1.1.md ].
- v25.1.1 patch notes pin **HBN 2.4.2-doca2.9.2-32** image for DPUService; documents HBN pod-restart traffic-loss caveat with workaround [official].
- DOCA Pipeline Language (DPL): P4-language match-action tables on BlueField pipeline; Flex Parsing, immediate-execution model (no deparser step), mid-pipeline reparsing, P4Runtime integration for SDN controllers [official] [source: https://docs.nvidia.com/doca/sdk/doca-pipeline-language-services-guide.pdf ].
- DOCA programming model: Kernel, RoCE, DPDK (`rte_flow`), SPDK, P4, P4Runtime workloads can coexist on the same DPU datapath [official] [source: https://developer.nvidia.com/blog/programming-the-entire-data-center-infrastructure-with-the-nvidia-doca-sdk/ ].

### 14.3 IPDK detail

- IPDK is Intel's infrastructure programming model: a vendor-neutral specification where P4 programs describe the forwarding pipeline and the runtime maps them to DPU/IPU hardware via DPDK, SPDK, OVS, and P4Runtime [official].
- IPDK targets any DPU/IPU that exposes an Infrastructure Abstraction Layer; reference implementations exist for Intel IPU E2000 and Mount Morgan-class devices; open-source on GitHub (Apache-2.0) [official/secondary].
- Distinction vs DOCA: IPDK is a portable programming abstraction; DOCA is a full product SDK + orchestrated services bound to BlueField hardware [independent].

### 14.4 Correction — consolidation chain

- **CORRECTION** to earlier working note: **Napatech → Alphawave Semi was NOT the deal**. The verified deal is **Qualcomm acquiring Alphawave Semi for $2.4B** (announced June 2025, completed ahead of schedule, ~Q4 2025); Qualcomm pairs Alphawave's high-speed wired connectivity/SerDes/chiplet IP with Oryon CPU and Hexagon NPU for AI data-center infrastructure [secondary] [source: https://www.techradar.com/pro/watch-out-nvidia-qualcomm-acquires-alphawave-semi-in-latest-addition-to-its-ai-data-center-push ].
- Alphawave Semi shareholders paid 183 pence/share (~96% premium) [secondary].
- Napatech (NT200A02 FPGA SmartNIC) post-acquisition status remains **unconfirmed** as of 2026-09-22 [gap — flag for recheck; the Alphawave linkage was erroneous].

