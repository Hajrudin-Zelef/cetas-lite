---
id: etape5-trackd-servers/01-server-vendors-wave-1-supermicro-dell-research-report-as-of-/1-2-new-launches-with-dates-2026
title: "1.2 New launches with dates (2026)"
domain: server-vendors-wave-1-supermicro-dell-research-report-as-of-
role: deep-dive
task: reference
actors: ["CoreWeave", "Intel", "Lambda", "Nscale", "Nvidia", "xAI"]
dates: ["2025-03", "2025-03-19", "2025-05", "2026-10"]
keywords: ["blackwell", "compute", "gpu", "gpus", "intel", "liquid cooling", "memory", "nvidia", "pricing", "rubin", "vera rubin"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [85, 134]
section: "Server Vendors — Wave 1: Supermicro + Dell (Research Report, as of 2026-09-22)"
sha256: bac0f682c2fb5bfed32e70bb797fe31d90967ac0ec06c3c8287e9b37ceb5d48b
---

# 1.2 New launches with dates (2026)

### 1.2 New launches with dates (2026)

| Date | Launch | Key specs | Source |
|---|---|---|---|
| Jan 2026 (reported Jan 2) | SuperBlade SBI-622BA-1NE12-LCC | Dual Intel Xeon 6900-series, liquid-cooled blade, up to 3 TB memory, 93% cabling reduction **[secondary]** | Barchart via greencitylivestock.com |
| Mar 2026 | Blackwell Ultra portfolio: HGX B300 NVL16 + GB300 NVL72, air- and liquid-cooled | 288 GB HBM3e per GPU; up to 800 Gb/s direct-to-GPU networking; 40°C warm-water (8-node rack) or 35°C warm-water (16-node double-density rack) operation with new CDUs; **up to 40% power reduction** and water conservation claimed **[official]** | Supermicro press release via europeansuntimes.com / worldnewsnetwork.net (article slug dated 2025-03-19 — note: slug suggests the underlying announcement may date to NVIDIA GTC March 2025 rather than 2026; TechPowerUp index lists it among Supermicro news) **[unverified date]** |
| Mar 2026 (MWC, Barcelona) | Pre-fabricated modular AI data centers with **SK Telecom + Schneider Electric**: pre-fab modules combining AI compute servers + power + cooling in one ready-to-deploy unit; SKT = operations, Supermicro = GPU servers, Schneider = power/infrastructure **[secondary]** | Ready-to-deploy building-block DC modules **[secondary]** | Barchart via business.dailytimesleader.com / markets.financialcontent.com |
| ~Oct–Nov 2025 (still 2026-relevant lineup) | HGX B300 4U + 2OU DLC-2 systems | DLC-2 captures **up to 98%** of system heat via liquid cooling; 45°C warm-water operation, no chilled water/compressors; up to 40% power savings; delivered as fully validated L11/L12 rack solutions **[official]** | stocktitan.net SMCI news page |
| ~Oct 2025 (NVIDIA GTC Washington, D.C.) | Expanded NVIDIA collaboration: 2OU HGX B300 (144 GPUs/rack), **Vera Rubin NVL144 + NVL144 CPX plans for 2026**, U.S.-manufactured **TAA-compliant** systems from San Jose, Super AI Station (GB300), GB200 NVL4 HPC; full-stack NVIDIA AI Factory for Government reference design **[official]** | Buy American Act-capable, TAA-compliant **[official]** | TechPowerUp / barchart.com / stocktitan.net (press-release text) |
| Aug 25–26, 2026 | **Cisco strategic collaboration**: Supermicro liquid- and air-cooled AI systems embedded in the **Cisco Secure AI Factory with NVIDIA**, validated and sold through Cisco and authorized channel partners; availability scheduled **from October 2026**; no disclosed financial transaction or equity stake **[official via vendor blog + independent (Axios)]** | Enterprise + AI-cloud channel access **[official]** | learn-more.supermicro.com; equityswarm.com citing Axios Aug 25, 2026 |

- **DLC-2 liquid cooling (2026 spec claims, [official])**: directly cools GPUs, CPUs,
  and memory with warm water up to 45°C inlet; up to 40% power-consumption reduction vs
  air cooling; up to 20% TCO reduction claimed in the DataVolt partnership context
  **[secondary]**; **250 kW per-rack heat-dissipation** capability; packs 8 NVIDIA
  Blackwell GPUs + 2 Intel Xeon CPUs into a single 4U server **[secondary (ainvest.com,
  May 2025)]**; captures up to 98% of heat through liquid **[official]**.

### 1.3 Pricing

- **No published per-server list prices** for Supermicro's 2026 Blackwell/Blackwell Ultra
  systems were found in collected sources; Supermicro sells via quote/channel
  **[unverified]** — pricing information is not publicly disclosed. Flagged as a gap.
- Deal-level figures are vendor/press-reported rather than list prices (see 1.4).

### 1.4 Major customer wins / partnerships (2026 and relevant 2025 context)

- **Cisco Secure AI Factory with NVIDIA** (Aug 2026): Supermicro AI systems sold via
  Cisco's enterprise and AI-cloud partners from Oct 2026 — no disclosed dollar value
  **[official/independent]** (see 1.2).
- **SK Telecom + Schneider Electric** (MWC, Mar 2026): modular AI data-center alliance
  **[secondary]**.
- **Lambda + ECL** (announced Sept 2025, still the reference 2026 GB300 deployment):
  first hydrogen-powered production-grade **NVIDIA GB300 NVL72** systems at ECL's
  Mountain View MV1 off-grid campus — **Supermicro-built, 142 kW systems**, ~4,000 lbs
  each, direct-to-chip liquid cooling, integrated into production in under two hours;
  Lambda doubled its footprint to 100% of MV1 **[independent]** (convergedigest.com;
  globalhydrogenreview.com). Supermicro SVP Technology & AI Vik Malyala confirmed the
  systems were delivered by Supermicro **[official quote]**.
- **DataVolt** (announced May 2025): **$20 billion partnership** to build hyperscale AI
  campuses in Saudi Arabia using Supermicro DLC-2 liquid-cooled systems
  **[secondary (ainvest.com)]** — the headline figure comes from market commentary, not
  a verified vendor filing; treat as **[unverified]** pending primary-source
  confirmation.
- **xAI (Colossus, Memphis)**: 2024–2025 context — Colossus supercomputer built with
  both **Dell and Supermicro** servers, ~100,000 NVIDIA GPUs, expanded toward 200,000
  **[secondary (SDxCentral)]**.
- CoreWeave/Nscale: **no specific 2026 Supermicro–CoreWeave win found** in collected
  sources — flagged **[unverified]**; not claimed.

