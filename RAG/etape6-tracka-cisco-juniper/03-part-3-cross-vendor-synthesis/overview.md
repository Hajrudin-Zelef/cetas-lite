---
id: etape6-tracka-cisco-juniper/03-part-3-cross-vendor-synthesis/overview
title: "PART 3 — CROSS-VENDOR SYNTHESIS"
domain: part-3-cross-vendor-synthesis
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "China", "Huawei", "Meta", "Nvidia", "Oracle", "TSMC"]
dates: ["2026-09-22"]
keywords: ["asic", "backlog", "cpo", "dci", "ethernet", "helios", "hyperscaler", "lpo", "nvidia", "optics", "packaging", "research"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [411, 455]
section: "PART 3 — CROSS-VENDOR SYNTHESIS"
sha256: dd9d90d64e852ddf0dd6183e2ecce096d2267004d98de6173f283e2d25f0c83e
---

# PART 3 — CROSS-VENDOR SYNTHESIS

*(Assembler-added section; research as of September 22, 2026.)*

---

## 9. Vendor positioning in AI data-center networking (2026)

### 9.1 Market snapshot — Ethernet switch revenue, Q2 2026

| Vendor | Q2 2026 switch revenue | Global Ethernet share | Data-center share | Notable 2026 move |
|---|---|---|---|---|
| Cisco | $5.43B (+36.2% y/y) | 28.7% (#1) | ~$2.2B (~18%) | Silicon One G300 102.4T; FY26 $9.3B hyperscaler AI infra orders |
| NVIDIA | ~$2.5B (+181% y/y; Next Platform cites $3.86B — scope discrepancy flagged) | — | 20.4% (#1 branded DC, IDC) | Spectrum-X CPO + Spectrum-X Ethernet Photonics in production with Vera Rubin |
| Arista | $2.53B (+37.6% y/y) | 13.4% | 18.7% | Tomahawk 6-based 7060XE7 up to 1.6T; full-year 2026 revenue guided >$10B |
| HPE (incl. Juniper) | $1.15B (+6.1%) | — | ~#4 | Oracle gigawatt-scale win (QFX scale-out + PTX scale-across); Networks-for-AI $2.2B cumulative orders |
| Huawei | $1.55B (+28.6%) | 8.2% | — | SP strength, China/emerging markets |

[independent — IDC via Next Platform (Sept 20, 2026) and kad8; NVIDIA figure discrepancy flagged in both drafts]

- Dell'Oro Q2 2026 AI back-end report: Ethernet AI back-end sales dominated by Celestica (white box) then NVIDIA; Arista third; **Cisco gained the most share, ranked fourth**. 800G was the vast majority of shipments; 1.6T began sampling, ramping H2 2026. AI back-end switch sales surpassed front-end sales for the first time in Q2 2026. [independent — Dell'Oro]
- Total Ethernet switch market: $18.9B in Q2 2026 (+43.4% y/y); DC segment $12.3B (+64.5%). [independent — IDC]
- thecodew analysis (Aug 2026): NVIDIA's DC Ethernet share rose from <4% to ~21.5% in Q1 2026, ahead of Arista (~19%) and Cisco (~14%). [secondary — treat figures as approximate]

### 9.2 Strategic comparison: Cisco vs HPE-Juniper

| Dimension | Cisco | HPE-Juniper |
|---|---|---|
| Core AI silicon | Silicon One G300 102.4T (in-house), plus Spectrum-X silicon support in same chassis | Merchant: Broadcom Tomahawk 6 (incl. CPO variant) for scale-up; Express 5 custom ASIC for routing |
| Fabric scale taxonomy | Scale-out focus; "Intelligent Collective Networking" | Scale-out (QFX), scale-up (TH6 over UALoE for Helios), scale-across/DCI (PTX12000) |
| Unified management | Nexus One + Unified Fabric (Feb 2026); Cloud Control / AgenticOps AI Canvas | HPE Mist AIOps + Aruba Central; "build once, deploy twice" (rollout Q1 2026) |
| UEC role | Member | Endorses UEC-compliant silicon; formal membership **unconfirmed** [unverified] |
| CPO posture | No 2026 CPO product; pluggable bet (1.6T OSFP, 800G LPO) | No Juniper CPO product; rides Broadcom TH6-Davisson CPO roadmap |
| Named 2026 hyperscaler wins | None named publicly (orders in aggregate) | Oracle gigawatt-scale AI cloud (QFX + PTX) |
| FY2026 AI networking figures | $9.3B hyperscaler AI infra orders (4.5× FY25); ~$4B revenue; FY27 guide $7.5B | Networks-for-AI: $2.2B cumulative orders Q3; FY26 target $2.5–3.0B; combined AI backlog $7.6B |

### 9.3 Optics & co-packaged optics — 2026 context

- **CPO entered volume production in 2026** (NVIDIA + Broadcom), though optical-engine/package supply remains a constraint risk. TrendForce: NVIDIA Spectrum-X CPO switch (co-developed with TSMC, COUPE packaging) and Broadcom **Bailly CPO** switches are in mass production, expanding late 2026. Spectrum-X CPO delivers up to 400 Tb/s; Bailly CPO cuts power up to 70% vs pluggables and completed deployment validation with hyperscalers including Meta. [secondary — TrendForce via electronics360]
- **Broadcom Tomahawk 6 – Davisson (TH6-Davisson)**: third-gen CPO Ethernet switch, industry's first 102.4 Tbps optically enabled switching, 16× 6.4T optical engines on TSMC COUPE; 70% lower optical-interconnect power (>3.5× vs pluggables); improved link-flap stability. [official — Broadcom]
- **NVIDIA**: Spectrum-X Ethernet Photonics in production alongside Vera Rubin generation; Quantum-X800 CPO InfiniBand platform at 115.2 Tbps; per-port power 30W→9W vs pluggables. [secondary]
- **Cisco's 2026 stance**: 1.6T OSFP + 800G LPO pluggable modules launched with G300 (50% lower module power, −30% switch power); Acacia shipped >750k 400G + 40k 800G coherent pluggables (record quarter; CEO says 200% FY2026 growth); new Open Transport 3000 multi-rail line system and NCS 1014 800G transponder card. **No Cisco CPO product found in 2026** — Cisco is betting on pluggables while NVIDIA/Broadcom lead CPO. [official/secondary — Cisco draft]
- **Juniper's 2026 stance**: no Juniper CPO product; CPO exposure via Broadcom Tomahawk 6; optics moves in 2026 were **800G ZR+ coherent** on PTX/MX for DCI (HPE AI Grid). [official/secondary — Juniper draft]
- Trend toward 1.6T: 1.6T optical modules ramping in 2026; thermal management of 1.6T pluggable systems pushing adoption of SiPh/CPO designs. [secondary]

