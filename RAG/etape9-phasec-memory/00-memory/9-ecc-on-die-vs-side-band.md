---
id: etape9-phasec-memory/00-memory/9-ecc-on-die-vs-side-band
title: "9. ECC: on-die vs side-band"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["AMD", "Intel"]
dates: ["2023-05"]
keywords: ["amd", "cost", "dram", "hbm", "hbm3", "hbm4", "intel", "memory", "nand", "research"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [167, 225]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: 389ce635ed13ac5f8de761b91f30bf2597ea4cbaef10493ce2e79982b9ffadca
---

# 9. ECC: on-die vs side-band

## 9. ECC: on-die vs side-band

- **Every DDR5 chip includes on-die ECC by JEDEC mandate**: 8 bits of ECC storage per 128 bits of data,
  correcting single-bit errors *inside the DRAM die* before data leaves the chip. Its purpose is to
  preserve manufacturing yields as cells shrink — it is not a substitute for system-level ECC
  [secondary].
- **Side-band (DIMM-wide) ECC** protects data in transit across the whole link, DRAM chip → module →
  memory controller. Traditional server ECC is SEC-DED (single-error correct, double-error detect)
  [secondary].
- Physical difference: ECC modules are wider. DDR2–DDR4 ECC used a 72-bit bus (e.g., nine x8 chips);
  a **DDR5 RDIMM uses an 80-bit bus** — two independent 40-bit sub-channels, each 32 data + 8 ECC bits
  (2×(32+8)) — and registered DIMMs are always side-band ECC [secondary].
- On-die ECC operates silently: it does not report error metrics to the OS or BMC, which is why servers,
  workstations and safety-critical systems still require side-band ECC DIMMs [secondary].
- Common myth to avoid: "DDR5 has ECC built in, so servers don't need ECC DIMMs." False — on-die ECC
  and side-band ECC are complementary layers [secondary].
- ZFS angle: ECC remains recommended for ZFS because bit-flips in dirty buffers (before checksums are
  computed) can be checksummed-corrupt and written permanently; on-die ECC alone does not close this
  window [independent].

## 10. Advanced ECC and RAS: Chipkill, SDDC, ADDDC, mirroring

- **Chipkill** (AMD-originated term): ECC scheme that survives the complete failure of one DRAM device
  (x4-based), correcting multi-bit errors spanning a failed chip; the server standard for mission-critical
  memory [secondary].
- Intel server platforms implement **SDDC** (Single Device Data Correction, Intel's chipkill-class mode)
  and **ADDDC** (Adaptive Double Device Data Correction), which can map out a failing device and adapt
  correction strength at runtime [secondary].
- DDR5's wider 80-bit sub-channel structure and on-die ECC combine with controller-side ECC (often
  referred to as EC4/EC8 schemes in technical discussion) to implement these device-failure-tolerant modes;
  exact code strength is controller-implemented and varies by vendor [secondary].
- Additional server RAS features: memory mirroring (channel mirroring, 50% capacity cost), rank sparing,
  patrol/demand scrubbing, memory address parity, and machine-check architecture (MCA) recovery flows
  [secondary].
- Practical note: enabling Chipkill/SDDC-class modes typically requires x4 DRAM-based RDIMMs and
  population rules per the platform vendor's validated lists [secondary].

## 11. HBM generations: the specification table

| Generation | Introduced | Data rate / pin | Interface width | Max BW / stack | Max capacity / stack | Dies per stack |
|---|---|---|---|---|---|---|
| HBM1 | 2015 | — | 1024-bit | 128 GB/s | 4 GB | 4 |
| HBM2 | 2018 | — | 1024-bit | 256 GB/s | 8 GB | 8 |
| HBM2e | 2020 | — | 1024-bit | ~460 GB/s | 16–24 GB | 8–12 |
| HBM3 | Jan 2022 (JEDEC) | 6.4 Gb/s | 1024-bit | 819.2 GB/s | 24 GB | 12 |
| HBM3e | May 2023 (vendors) | up to 12.4 Gb/s | 1024-bit | up to ~1.33 TB/s | 24–36 GB | 8–16 |
| HBM4 | Apr 16, 2025 (JESD270-4) | 8 Gb/s baseline (13+ Gb/s demonstrated) | **2048-bit** | 2.0–2.56 TB/s | 36–64 GB | 8–16 |
| HBM4e | previewed Mar 2026 | 16 Gb/s | 2048-bit | 4.0 TB/s | — | — |
[secondary]

- The big HBM4 change: the interface width doubles from 1024-bit to 2048-bit with 32 channels, so
  bandwidth scales via width rather than ever-faster per-pin signaling; HBM4 also introduces an optional
  logic layer at the base of the stack for test, repair and thermal management [secondary].
- Rambus outlined an HBM4 memory controller at up to 10 Gb/s speeds, 2.56 TB/s bandwidth and 64GB
  capacities per stack [vendor-reported].
- Forward roadmap (KAIST research roadmap via Tom's Hardware): HBM5 ~2029 keeps HBM4's data rate but
  doubles I/O to 4096-bit for 4 TB/s and 80 GB/stack with stacked decoupling capacitors and 3D cache;
  HBM8 (2038 projection): 16,384-bit interface, up to 64 TB/s, embedded NAND [secondary].

