---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/part-7
title: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) (part 7)"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "AWS", "Broadcom", "Google", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "intel", "nvidia", "aws", "ethernet", "gpu", "research", "revenue", "rubin"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [265, 271]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 6fdbd31ab34106e2b60fa459d644b06c0b5926430e0b733b1cf1b7e91998012e
---

# Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) (part 7)

- **AI networking bifurcation**: (a) NVIDIA vertical stack (ConnectX-8 SuperNIC / BlueField-4 + Spectrum-X + InfiniBand) vs (b) open Ethernet (AMD Pollara/Vulcano + UEC + Broadcom Tomahawk 6 switching + Arista 7800R4). Competition centers on software programmability, congestion control, transport, TCO — not raw bandwidth [secondary/analysis].
- **BlueField-4 topology detail** (SDxCentral, from NVIDIA director): 800 Gb/s total, typically **two 400 Gb/s bidirectional links**; Ethernet at the front end for north–south; GPU east–west handled by **multiple ConnectX-9 chips packaged together delivering 1.6 Tb/s per GPU** starting with Rubin; BlueField-4's 800G seen as sufficient headroom for the Rubin generation [vendor-reported/secondary].
- **BlueField-5** already on roadmap: slated **2028** with NVIDIA's **Feynman** architecture, per NVIDIA's "annual rhythm" cadence [vendor-reported].
- DPU market forecast: $554M (2021) → **$5.5B by 2031** at 27% CAGR (Allied Market Research via DataCenterKnowledge) [secondary]; 2026 realized revenue not located [gap].
- "Third socket" narrative: DPU as infrastructure root-of-trust, first line of defense, primary data orchestrator — freeing 100% of host CPU for applications [independent/analysis] [source: https://github.com/jorgeadev/the-architects-blueprint/blob/HEAD/web/src/content/blog/2026/04/28/beyond-the-cpu-architecting-hyperscale-analytics-.md ].
- Vendor landscape summary (merchant + in-house): NVIDIA BlueField, AMD Pensando, Intel IPU, Marvell OCTEON, Broadcom Stingray, Fungible (→Microsoft), Napatech (→Alphawave Semi), Xsight Labs, Kalray, Ethernity, Chelsio, Silicom, Achronix, Netronome, Asterfusion; in-house: AWS Nitro, Azure SmartNIC, Google IPU [secondary] [source: https://medium.com/@lixian_58397/the-most-comprehensive-dpu-smartnic-vendors-with-its-product-line-summary-db7899a725c8 ].

