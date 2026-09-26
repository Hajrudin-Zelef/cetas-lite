---
id: etape6-phasec-optics-cabling/00-front-matter/11-26-structured-cabling-9-216-strands-commscope-propel-mpo-
title: "11.26 Structured cabling: 9,216 strands, CommScope Propel, MPO-16"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Broadcom", "Crusoe", "IREN", "Intel", "Lambda", "Meta", "Microsoft", "Nebius", "Nvidia", "Oracle", "United States"]
dates: ["2022-03-31", "2026-01", "2026-04-20", "2026-09"]
keywords: ["alignment", "amd", "compute", "cost", "cpo", "dsp", "ethernet", "gpu", "hyperscaler", "intel", "latency", "lpo"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1877, 1932]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 8575b74b83e9ab0e27cb13a2ac399090e7f6ce60e33c298552bc52c6f3aff31d
---

# 11.26 Structured cabling: 9,216 strands, CommScope Propel, MPO-16

### 11.26 Structured cabling: 9,216 strands, CommScope Propel, MPO-16

- A structured-cabling vendor page describes an **8-rack, 576-GPU** DGX GB200 scalable unit with **9,216 active fiber strands total (4,608 compute-fabric strands)**, 1,152 fibers per rack at server-to-leaf level, 400G/800G native port speeds [vendor-reported] (americas.scalefibre.com).
- The internal 72-GPU NVLink fabric is passive copper with zero optical fiber; fiber serves scale-out compute and storage/management fabrics [vendor-reported].
- **CommScope Propel** (launched March 31, 2022) is an MPO-16-fiber structured cabling platform supporting 400G/800G and emerging 1.6T; CommScope claims it was the industry's first 16-fiber MPO platform [vendor-reported] (nasdaq.com press release).
- Propel includes 1U/2U/4U panels, LC/SN/MPO adapter packs, 8/12/16/24-fiber modules (single-mode and multimode), and MPO-8/12/16/24 plus SN/LC uniboot cable assemblies, all ultra-low-loss with QR-code performance tracking and SYSTIMAX warranty [vendor-reported].
- Propel coverage: lightwaveonline.com [independent].
- Corning EDGE-16 details were **not** verified in this research round — flag as a gap; do not substitute CommScope details [unverified].
- 800G DR8 optics use 8 fibers (MPO-8/APC); 800G SR8/VR8 use 16 fibers (MPO-16); 400G DR4 uses MPO-12 (8 used) — connector choice follows the optic's fiber count, not the NIC alone [independent].

### 11.27 Optics market (LightCounting / TrendForce)

- LightCounting January 2026 "Optics for AI" report: Ethernet optical-transceiver and CPO sales for AI scale-up/scale-out plus related categories estimated at **US$16.5 billion in 2025**, forecast **US$26 billion in 2026** (~60% growth) [secondary].
- LightCounting: VCSEL constraints eased by mid-2024 after NVIDIA shifted from SR8 toward DR8 800G; InP laser shortages expected to ease by mid-2026; possible flat quarter or two in late 2026 as supply/demand rebalance [secondary].
- TrendForce April 20, 2026: same US$16.5B→US$26B estimate, over 57% YoY growth, rising demand for 800G-and-above AI-cluster optics, with EML, CW laser, alignment, power, and thermal bottlenecks [secondary].
- Comparability caveat: LightCounting's category includes Ethernet optics starting at 100G plus CPO for AI scale-up/scale-out — **do not equate the full US$26B to 800G transceivers alone** [independent].

### 11.28 1.6T roadmap, LPO/CPO outlook

- IEEE P802.3dj defines 200G/400G/800G/1.6T interfaces using **200G-per-lane** technology; completion expected late/H2 2026, with early 200G-per-lane products expected during 2026 [independent].
- NVIDIA's photonics roadmap (GTC 2025): Spectrum-X Photonics and Quantum-X Photonics up to 1.6 Tb/s per port; reported configs 128×800G or 512×200G for 100T-class and larger 400T systems — announced platforms, production wording needs care [vendor-reported].
- Broadcom OFC 2026: Taurus 400G/lane optical DSP plus 400G EML and photodiodes to enable 1.6T modules and future 3.2T; 200G/lane retimers/AECs; Tomahawk 6 (102.4T) described by Broadcom as **shipping in production volume** [vendor-reported].
- **LPO (Linear Pluggable Optics)** removes the module DSP and relies on host SerDes: potential lower power, lower latency, potentially lower cost [independent].
- LPO risks: less equalization margin, greater channel sensitivity, interoperability and operational qualification burden [independent].
- LPO evidence: Broadcom BCM57608 explicitly advertises LPO support; Oracle Acceleron uses LPO/LRO at 400G/800G; Thor Ultra claims linear-drive optics support [vendor-reported].
- Do not claim broad hyperscaler LPO production without stronger deployment evidence [independent].
- **CPO (Co-packaged optics)**: pursued for 1.6T+ power reduction (e.g., NVIDIA Spectrum-X Photonics roadmap, Broadcom Tomahawk 6 CPO-class products); no broad merchant-NIC CPO deployment verified in this research [independent].

### 11.29 Vendor 800G product matrix (as of 22 September 2026)

- NVIDIA ConnectX-8: 800G, PCIe Gen6 x16, OSFP/QSFP112, IB + Ethernet — **shipping/channel available** [secondary].
- NVIDIA ConnectX-9: 1.6T, PCIe Gen6 — **announced Oct 2025, Rubin timeframe** [vendor-reported].
- NVIDIA BlueField-4: 800G DPU, Grace + ConnectX-9 — **announced Oct 2025; partner STX systems guided H2 2026; no standalone retail price found** [vendor-reported].
- Broadcom Thor Ultra: 800G, PCIe Gen6 x16, UEC-oriented — **sampling to select customers** [secondary].
- Broadcom Thor 2: **400G only** — not an 800G product [independent].
- Broadcom Stingray PS1100R: **100G storage adapter** — not an 800G product [official].
- AMD Pensando Vulcano: 800G, PCIe Gen6, UEC — **2026 target per TechRadar / 2027 transition per roadmap interpretation; conflicting** [secondary].
- AMD Pensando Pollara: **400G only** — not an 800G product [vendor-reported].
- Intel E830: **200G ceiling** — not an 800G product [official].
- Marvell: **no verified 800G NIC/DPU found** [independent].
- Fungible: **acquired by Microsoft (Jan 2023)** — no independent 800G product [vendor-reported].

### 11.30 Open research gaps for a follow-up wave

- Primary-source GA/shipping confirmation for ConnectX-8 (beyond firmware manuals and channel listings) [unverified].
- Standalone BlueField-4 card availability and pricing [unverified].
- Thor Ultra production ramp date and pricing [unverified].
- Vulcano launch timing (2026 vs 2027) reconciliation with an AMD primary source [unverified].
- Meta primary-source 800G deployment confirmation [unverified].
- Corning EDGE-16 product details [unverified].
- Arista 800G AI switching specifics and deployment counts [unverified].
- Lambda, Crusoe, Nebius, IREN, Vultr compute-fabric (non-storage) 800G NIC details [unverified].
- Official OSFP1600/QSFP-DD1600 MSA roadmaps for the 1.6T NIC generation [unverified].

### 11.31 Source list (verbatim URLs)

