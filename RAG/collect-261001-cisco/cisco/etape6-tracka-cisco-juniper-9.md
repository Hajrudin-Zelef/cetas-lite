---
id: collect-261001-cisco/cisco/etape6-tracka-cisco-juniper-9
title: "Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Broadcom", "CoreWeave", "Meta", "Nvidia", "Oracle", "TSMC", "United States"]
dates: ["2026-02-10", "2026-02-25", "2026-03-17", "2026-03-24", "2026-05-04", "2026-05-06", "2026-05-08", "2026-06-01", "2026-06-02", "2026-06-15", "2026-08-14", "2026-08-25", "2026-08-31", "2026-09", "2026-09-15", "2026-09-17", "2026-09-20", "2026-09-22"]
keywords: ["acquisition", "amd", "asic", "backlog", "benchmarks", "cpo", "dci", "ethernet", "helios", "hyperscaler", "lpo", "nvidia"]
source: docs/RAG/collect-261001-cisco/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [431, 503]
sha256: 82488c469f069b02a3334c8d8398f486d4528248974e32f86913c762fe149cd2
---

# Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026

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

### 9.4 Master timeline (January → September 2026)

| Date | Event |
|---|---|
| 2026-02-10 | Cisco announces Silicon One G300 (102.4T) at Cisco Live EMEA Amsterdam [official] |
| ~2026-02-25 | HPE announces PTX12000/PTX10002 Express 5 routers pre-MWC (MWC: Mar 2–5) [official/secondary] |
| 2026-03-17 | HPE AI Grid with NVIDIA announced at NVIDIA GTC (PTX/MX + 800G ZR+, Spectrum-X AI Factory) [official/secondary] |
| 2026-03-24–26 | Juniper SRX400 firewalls launched at RSA Conference [official] |
| ~2026-04 | Cisco acquires Galileo (AI observability, terms undisclosed) [secondary] |
| 2026-05-04 | Cisco announces Astrix Security acquisition (press-reported ~$400M, officially undisclosed) [secondary/unverified price] |
| 2026-05-06 | HPE ships 723H Wi-Fi 7 AP — first unified Juniper×Aruba product (dual Mist/Central) [secondary] |
| 2026-05-08 | HPE "Self-Driving Networks" autonomous workflows launch [secondary] |
| 2026-06-02–04 | Cisco Live Las Vegas: Cisco Cloud Control / AgenticOps + AI Canvas [official/secondary] |
| 2026-06-15–18 | HPE Discover: refreshed AI switch lineup (QFX5252 liquid-cooled scale-up, QFX5250, QFX5140), quantum-safe SRX4700, SASE Orchestrator [secondary] |
| 2026-06-01 | CoreWeave first Vera Rubin rack (context) [from step-5 track A] |
| ~2026-07 | CPO switches (NVIDIA Spectrum-X, Broadcom Bailly) in volume production per TrendForce [secondary] |
| 2026-08-14 | Judge Casey Pitts dismisses states' challenge to HPE–Juniper DOJ settlement (final legal clearance) [secondary] |
| 2026-08-25 | Cisco + Teleport identity partnership [secondary] |
| 2026-08-31 | AMD/Cisco/HUMAIN Saudi AI zone operational on MI355X (Luma AI anchor) [secondary — Cisco draft] |
| 2026-09-15 | Cisco AI POD for Splunk announced at Splunk.conf Denver [official/secondary] |
| 2026-09-17 | Cisco joins Verizon 6G Forum [secondary] |
| 2026-09-20 | Next Platform: IDC Q2 2026 Ethernet data ($18.9B, +43.4%) [independent] |
| 2026-09-22 | Report date |

### 9.5 Master verification log (combined)

1. Cisco earnings-sourced figures ($9.3B orders, $7.5B guide) rest on consistent trade-press coverage, not the IR release text itself — Cisco draft §11.
2. Silicon One G300 on-chip buffer size undisclosed.
3. No Cisco G500 / 204.8T announcement found in 2026.
4. No Cisco CPO product found in 2026 — stance is pluggables.
5. Cisco Astrix price ($400M) press-reported, officially undisclosed.
6. Cisco claims AI Nexus orders +85% sequentially and "market share gains"; one displacement anecdote unverified.
7. Juniper UEC membership unconfirmed.
8. No 2026 Juniper CPO product; strategy rides Broadcom's roadmap.
9. No discrete 2026 Apstra release found.
10. No new named Junos OS version for 2026 found (latest AI-ML docs 24.4-era).
11. No 2026 ACX announcements found.
12. No 2026 Juniper DCQCN/RoCEv2 announcement found.
13. IDC Q2 2026 NVIDIA figure discrepancy: $2.5B (kad8) vs $3.86B (Next Platform) — scope unresolved.
14. HPE–Juniper deal value cited as $14B and $13.4B (likely gross vs net).
15. Instant On divestiture details rest on a single buyer-guide source (Fortinet $5–15M offer, Extreme declined) — buyer approval status unknown.
16. TrendForce CPO "mass production" claims are market-research reporting, not vendor production disclosures — treat as secondary.
17. thecodew market-share figures (NVIDIA 21.5% Q1 2026 etc.) are approximate/secondary.
18. Luma AI / HUMAIN Saudi JV (100 MW → 1 GW) operational status per vendor/PR reporting — secondary.
19. Cisco vs Arista vs HPE-Juniper: no 2026 head-to-head AI-fabric benchmarks located — comparative performance claims are vendor-reported only.

### 9.6 Collection metadata

