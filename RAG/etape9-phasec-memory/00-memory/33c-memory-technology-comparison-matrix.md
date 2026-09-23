---
id: etape9-phasec-memory/00-memory/33c-memory-technology-comparison-matrix
title: "33C. Memory technology comparison matrix"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["AMD", "Falcon", "Intel", "Meta", "Microsoft", "Nvidia", "Samsung", "United States"]
dates: ["2026-09-22"]
keywords: ["memory", "accelerator", "amd", "attention", "compute", "cost", "dram", "energy", "gpu", "gpus", "hbm", "hbm4"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [591, 674]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: b6c763da3c670c734125a87272f2c041f5d9aa14d4f309c50a108db737772cf1
---

# 33C. Memory technology comparison matrix

## 33C. Memory technology comparison matrix

| Attribute | DDR5 (server RDIMM) | HBM3e | HBM4 | GDDR7 | LPDDR5X |
|---|---|---|---|---|---|
| Interface | 2× 40-bit (80-bit ECC) per DIMM | 1024-bit/stack | 2048-bit/stack | 32-bit/device, 4 ch | 128-bit (LPCAMM2) |
| Per-device BW | ~51 GB/s @6400 | ~1.2 TB/s/stack | ~2 TB/s/stack | 128 GB/s @32Gb/s | ~77 GB/s @9600 |
| Latency class | 80–90 ns | 100–150 ns | 100–150 ns | ~100 ns | ~100 ns |
| Capacity/GPU | TBs (socket) | 141–288 GB | 288–432 GB | 128 GB (CPX) | N/A (host) |
| Power | 1.1V DIMMs | High (stack) | ~75W/stack proj. | 1.2V devices | Lowest |
| Role | Host capacity tier | Accelerator BW tier | Accelerator BW tier | Cost-efficient BW | Power-efficient host |

## 33D. Server memory power and efficiency

- DDR5's 1.1V operation plus on-module PMIC improves pJ/bit vs DDR4; Micron claims up to 24% better
  energy efficiency (pJ/bit) on its 1β-based 128GB RDIMMs vs competitive 3DS TSV products
  [vendor-reported].
- LPCAMM2 (LPDDR5X): up to 58% lower active power and 80% lower standby power vs DDR5 SODIMM subsystems
  — the efficiency argument for LPDDR in power-constrained edge/AI-PC and (per Nvidia's pivot) possibly
  accelerator-adjacent designs [vendor-reported].
- HBM power is now a rack-design variable: B200-class 1000W GPUs and B300-class 1400W parts include
  significant HBM-stack power; direct-liquid cooling is effectively mandatory at these densities
  [secondary].
- CXL pooling's efficiency claim: up to 50% higher effective memory utilization reduces provisioned
  DRAM per workload — a TCO/power argument, not just capacity [secondary].

## 34. What to watch after 2026-09-22

- DDR6 final spec ratification and first platform validation samples (expected through 2027).
- HBM4 volume ramp on Rubin (H2 2026) and MI400/MI455X (2026); HBM4e qualification for Rubin Ultra (H2 2027).
- CXL 3.1/4.0 platform enablement: switch silicon (Atlas 3, Marvell, Montage) moving from sampling to
  production; OS/hypervisor fabric-memory management maturity.
- Whether DRAM supply normalizes by 2028 or the shortage extends to 2030 (analyst split).
- NVIDIA's LPDDR pivot scale — potential new demand sink for low-power DRAM.
- Samsung vs SK hynix HBM4 share evolution and its effect on accelerator pricing.
- MRDIMM 8800 MT/s platform qualification results (Q1 2027 Xeon 6900P systems) and whether 12,800 MT/s
  Gen2 stays on track.
- First JEDEC DDR6 module spec drafts and CAMM2-for-servers platform announcements.
- HBM failure-rate data from large fleets — whether memory RAS for accelerators becomes a productized
  feature (on-die ECC in HBM is already standard; stack-level sparing/repair is the next frontier)
  [secondary].
- Outcome of the US DRAM class-action suit and any regulatory response to 2026 memory pricing.

## 35. Conflicts, gaps and non-comparable figures

1. **B200 memory**: 180GB vs 192GB per GPU across sources — DGX-derived 192GB used here; 180GB figure
   [unverified].
2. **DDR6 timeline**: one article claims 2026 debut/2027 mass adoption; the weight of 2026 reporting +
   JEDEC status points to 2028–2029 commercial — outlier flagged [unverified].
3. **DDR6 on CAMM2**: reported as industry direction, not a finalized JEDEC module standard — treat
   connector claims as [secondary], not settled.
4. **HBM4 pricing**: $500 (Samsung→Nvidia), ~$550 (SK hynix), ~$600+/unit (Korean press), $3,500 spot for
   16-layer pre-production — different products, contract vs spot; not directly comparable.
5. **DRAM $/GB**: retail ($18.44), secondary-market ($8.11–$18.10), Framework pass-through ($10), contract
   ($2.94/Gb ≈ $23.5/GB equivalent basis) — different channels; do not blend.
6. **Rubin per-GPU HBM4 bandwidth**: 13 TB/s (NVIDIA slides via press) vs one source's 22 TB/s mention
   for Rubin-class — 13 TB/s used; the 22 TB/s figure is [unverified].
7. **Venice EPYC 1.6 TB/s memory bandwidth**: single weak source — [unverified].
8. **MI325X**: originally announced 288GB, shipped 256GB — both figures documented; use 256GB for the
   shipping product.
9. **Gaudi 3 networking**: 24× 200GbE (4.8 Tbps aggregate) vs one source's 8-port/1.6 Tb/s phrasing —
   24-port figure from Intel/IBM used; discrepancy noted [unverified] on the 8-port claim.
10. Gaps: exact 2026 server-RDIMM contract $/GB; NVDIMM 2026 product news; Meta MTIA / Microsoft Maia
    2026 memory configs; Intel Falcon Shores memory config; HBM power-per-stack vendor figures
    (roadmap projections only); hard 2026 job-board counts for memory-adjacent roles; per-vendor
    DDR5-8000 RDIMM street pricing.

## 36. Glossary

- **RDIMM/LRDIMM/3DS**: registered / load-reduced / 3D-stacked DIMM variants (Section 4).
- **MRDIMM**: Multiplexed Rank DIMM — multiplexes two ranks per channel for higher effective bandwidth.
- **RCD**: Registering Clock Driver — buffers command/address on RDIMMs (Gen4/Gen5/Gen6 generations).
- **On-die ECC vs side-band ECC**: in-chip correction (yield) vs DIMM-wide end-to-end protection (Section 9).
- **SDDC/ADDDC/Chipkill**: device-failure-tolerant ECC modes (Intel/AMD).
- **HBM**: High Bandwidth Memory — stacked DRAM on interposer, 1024/2048-bit interface.
- **HBM2e/3/3e/4/4e**: successive HBM generations (Section 11).
- **Base die**: logic die at the bottom of an HBM stack (I/O, test, increasingly controllers).
- **GDDR**: Graphics DDR — discrete DRAM chips for GPUs; GDDR7 uses PAM3.
- **CXL**: Compute Express Link — PCIe-based coherent interconnect; Type 3 = memory device.
- **GFAM**: Global Fabric Attached Memory (CXL 3.1).
- **KV cache**: per-session attention key/value storage in LLM inference; scales with context × batch.
- **DWPD/TBW**: endurance metrics — covered in Phase A/B; referenced here only for completeness.
- **PLP**: power-loss protection (capacitors on SSDs/controllers).
- **1DPC/2DPC**: DIMMs per channel — population density affecting max speed.

