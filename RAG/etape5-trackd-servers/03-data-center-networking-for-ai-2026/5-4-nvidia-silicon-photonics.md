---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/5-4-nvidia-silicon-photonics
title: "5.4 NVIDIA silicon photonics"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: actor-profile
actors: ["Broadcom", "Cohere", "Google", "Meta", "Nvidia", "Oracle"]
dates: ["2026-03", "2026-06", "2026-07"]
keywords: ["nvidia", "coherent optics", "cost", "cpo", "datacenter", "dci", "energy", "ethernet", "gpus", "hyperscaler", "optics", "rack-scale"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [854, 875]
section: "Data-Center Networking for AI (2026)"
sha256: b712ec5e090ea52cccade077001b254fd8e115311bfc958e65124d5f3f660578
---

# 5.4 NVIDIA silicon photonics

- TrendForce: **800G+ shipments ~24M units (2025) → ~63M units (2026)**, 2.6× growth; 1.6T shipments **~2.5M (2025) → 20M+ by end-2026** (per marketminute coverage of OFC 2026 analyst confirmation; 1.6T "scaling faster than any previous generation"). [secondary]
- LightCounting-derived TAM (cited in a July 2026 investor research note): total transceivers $23.8B (2025); datacenter modules ~$22.8B (2026), of which **800G+1.6T ~$14.6B**; 800G+ units ~24M (2025) → ~63M (2026). Note: scope disagreements are wide (MarketsandMarkets pegs DC transceivers at only $9.2B for 2025) — treat single points cautiously. [secondary] (github/roachx92 citing LightCounting)
- GM Insights (market-research firm, 2026): optical transceiver market $13.4B (2025) → $15.4B (2026) → $48.1B (2035), 13.5% CAGR; Coherent led with >22.2% share in 2025; top-5 (Coherent, Cisco, Broadcom, Lumentum, Accelink) held 72.2%. (Estimate — [secondary].) [secondary]
- Vendor landscape: **Innolight (~$5.6B datacenter modules)** and Eoptolink together hold **~60% of NVIDIA's 800G volume**; Coherent and Lumentum are the Western integrated laser suppliers (Lumentum the only volume 200G-EML shipper at ~50–60% share, +40% EML capacity in 2025 and again 2026); Coherent expanding 6-inch InP across 4 sites (~4× die/wafer at ~½ cost, capacity doubling 2026); Cisco (Acacia) took >$1B of optics orders in a single quarter (Q4 FY2026). [secondary] [vendor-reported] (mix: github research note citing J.P. Morgan/LightCounting; Cisco via ainvest)
- **EML bottleneck:** TrendForce reports an upstream laser bottleneck; NVIDIA pre-allocated large EML capacity (investor note claims "$4B lock-up of Lumentum+Coherent EML capacity" with multiyear purchase commitments — [unverified]); NVIDIA's March 2026 $2B direct investment in Lumentum to guarantee photonics capacity (marketminute). Extended EML lead times beyond 2027 reported. [secondary]
- Adopters: "1.6T is shipping now," with **NVIDIA and Google already integrating**; Meta and Oracle slated to follow (per Kozlov via eetimes). [independent]

### 5.4 NVIDIA silicon photonics

- NVIDIA's roadmap includes **Spectrum-X Ethernet Photonics (H2 2026)** and Quantum-X InfiniBand (early 2026) — i.e., CPO/silicon-photonics switch variants (per the github silicon-photonics research note). [unverified] — no dated NVIDIA product announcement of these variants was located in vendor sources; treat H2 2026 as roadmap guidance, not confirmed shipping.
- NVIDIA announced **Quantum-X and Spectrum-X silicon photonics networking switches** (referenced in the Spectrum-XGS release as enabling "AI factories to connect millions of GPUs across sites while reducing energy consumption and operational costs"). [official]
- NVIDIA CTO Gilad Shainer: CPO "already shipping and ramping H2 2026" (response to June 2026 SemiAnalysis yield-risk report). [vendor-reported]
- NVIDIA's slower-than-anticipated silicon-photonics/CPO progress is one reason for continued dependence on pluggable modules and the EML pre-allocation strategy (TrendForce). [secondary]

### 5.5 Data center interconnect (DCI) notes

- 800G digital coherent optics (DCO): $963M market (2025) → $1,713M (2032), 8.7% CAGR (MarketPublishers, Jan 2026 — [secondary] estimate); key players II-VI/Coherent, Lumentum, Innolight, Hisense, Accelink. [secondary]
- Scale-across DCI is emerging as its own category: Cisco's Silicon One **P200** is explicitly positioned for scale-across deployments (three hyperscaler wins in Q4 FY2026 incl. a P200 scale-across win); NVIDIA Spectrum-XGS is the Ethernet answer; Arista's 1.6T platforms target scale-across. [vendor-reported] [official]
- Corning announced a multi-billion-dollar "rack-scale" fiber infrastructure deal with Meta for gigascale AI data centers (March 2026). [secondary] (marketminute)

---

