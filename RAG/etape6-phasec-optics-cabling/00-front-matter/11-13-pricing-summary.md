---
id: etape6-phasec-optics-cabling/00-front-matter/11-13-pricing-summary
title: "11.13 Pricing summary"
domain: front-matter
role: reference
task: pricing
actors: ["Broadcom", "CoreWeave", "Intel", "Nvidia", "United States", "xAI"]
dates: ["2024-10", "2025-02", "2025-03-14", "2026-01-28"]
keywords: ["pricing", "benchmarks", "blackwell", "compute", "ethernet", "gpu", "gpus", "inference", "intel", "latency", "memory", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1774, 1821]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: bc400be9142daef8df3d309da1f6912f322df733868e6a1b370c577fabd2cc08
---

# 11.13 Pricing summary

### 11.13 Pricing summary

- ConnectX-8 channel/retail prices found: US$2,519 (FS.com, one-port 800G OSFP and dual-port 400G QSFP112), US$1,779 (NADDOD C8240/C8180) [secondary].
- No public unit prices were found for BlueField-4, Thor Ultra, or Pensando Vulcano [unverified].
- Intel E830 four-port 25GbE HPE option ~US$7,000 is an integrated server option and not comparable with bare 800G NIC prices [secondary].
- Resale listings (e.g., US$2,000 asking for new single-port 800G cards) are anecdotal only [secondary].

### 11.14 Scale-up vs scale-out primer

- **Scale-up**: tightly couples accelerators within a server, tray, rack, or supernode for memory-semantic or very-high-bandwidth GPU/XPU collectives (e.g., NVIDIA NVLink/NVSwitch in GB200 NVL72) [official].
- **Scale-out**: connects servers or rack-scale systems across racks using InfiniBand or Ethernet/RoCE/UET; 400G and 800G NICs and switch ports are primarily scale-out components [independent].
- In NVIDIA NVL72, 800G is principally part of the external compute/storage fabrics, not the internal NVLink domain [independent].
- Broadcom's Thor Ultra positioning makes this explicit: NVLink connects 72–256 XPUs in rack-scale domains; Thor Ultra targets rack-to-rack scale-out [independent].

### 11.15 NVIDIA DGX GB200 SuperPOD reference architecture

- One official scalable unit: **8 DGX GB200 NVL72 compute racks**, each with 36 Grace CPUs and 72 Blackwell GPUs (576 GPUs per scalable unit) [official].
- Internal scale-up: NVLink 5 with integrated NVLink switch trays [official].
- Compute scale-out fabric in the reference design: Quantum QM9700 InfiniBand with eight NDR400 connections per system [official].
- Storage and in-band management: Spectrum-4 SN5600 Ethernet switches with 64 ports of 800Gb/s [official].
- Note: the reference design shows **NDR400** QM9700 links for cross-rack compute — do not imply every GB200 SuperPOD uses 800G InfiniBand [official].
- Reference components page: docs.nvidia.com DGX SuperPOD GB200 scalable infrastructure [official]. Reference architecture PDF: docs.nvidia.com RA11338001-DSPGB200 [official].
- Terminology caveat: one NVIDIA table snippet called NVL72 a "scale-out GPU interconnect" while the same architecture uses NVLink internally and InfiniBand externally; use functional definitions (§11.14), not the inconsistent label [official].

### 11.16 NVIDIA Spectrum-X800 / Quantum-X

- Spectrum-X800 platform: SN5600 switch, 64 OSFP ports at up to 800GbE, 51.2 Tb/s switching, paired with BlueField-3 SuperNICs in current deployments [vendor-reported].
- Spectrum-X features: adaptive routing, congestion control, direct data placement, telemetry, performance isolation [vendor-reported].
- Quantum-X Photonics series (announced, GTC 2025): reported 144 ports at 800 Gb/s InfiniBand; treat as announced platform/roadmap, production availability needs careful wording [vendor-reported].
- Photonics roadmap coverage: tomshardware.com [independent].
- GTC 2025 secondary reporting also described Spectrum-X Photonics configurations at 128×800G or 512×200G for 100T-class systems [secondary].

### 11.17 xAI Colossus

- NVIDIA/xAI reporting (October 2024): initial cluster 100,000 Hopper GPUs, expansion target 200,000 H100/H200 GPUs, built in 122 days [vendor-reported].
- Colossus used **Spectrum-X Ethernet, not InfiniBand**, for the RDMA network [independent].
- Reported endpoint design: one 400GbE BlueField-3 SuperNIC per GPU; SN5600 switches provide 64×800G ports [independent].
- This is an 800G-switch deployment but primarily **400G host-facing NIC** connectivity — an 800G switch port may break out to two 400G endpoints [independent].
- NVIDIA's claims of 95% data throughput with no latency degradation or packet loss from flow collisions are vendor claims, not independent neutral benchmarks [vendor-reported].
- Coverage: sdxcentral.com [independent]; tomshardware.com [independent]; globenewswire.com announcement [vendor-reported].

### 11.18 CoreWeave

- March 14, 2025: CoreWeave/Bulk announced a large GB200 NVL72 cluster in Norway with Quantum-2 InfiniBand, expected operational summer 2025 [vendor-reported].
- January 28, 2026: CoreWeave announced NVIDIA Exemplar validation for GB200 NVL72 inference and referenced prior training validation [vendor-reported].
- Secondary reporting: GB200 NVL72 instances became available February 2025, scale stated up to 110,000 GPUs, BlueField-3 DPUs used for cloud networking and data access [secondary].
- **No explicit 800G host-NIC deployment count was found** for CoreWeave [unverified].

