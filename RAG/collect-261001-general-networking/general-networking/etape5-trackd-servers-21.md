---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-21
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "China", "CoreWeave", "Huawei", "Nvidia", "United States", "xAI"]
dates: ["2026-07", "2026-09-22", "2026-10"]
keywords: ["advisory", "amd", "asic", "backlog", "blackwell", "capex", "datacenter", "energy", "gpu", "gpus", "liquid cooling", "nvidia"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [1328, 1403]
sha256: c885e5f0e7a1d44f9436f6f0e2e557676f3d2cdf46694f94177645d88f1dcf4b
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

### 7.4 Liquid cooling adoption
- **~22% of newly built facilities** using liquid cooling by 2026 (KAD industry overview); cold plate / direct-to-chip ≈ **65% of the liquid-cooling market** in 2026; default design for NVIDIA GB200/GB300. [secondary — KAD](https://www.kad8.com/server/data-center-liquid-cooling-for-ai-workloads-2026/)
- APAC liquid cooling market: **$1.7B (2025) → $2.1B (2026) → $17.5B (2036)**, 23.6% CAGR; direct-to-chip ~55% share in APAC 2026. [secondary — DataNext Research]
- AI datacenter liquid cooling (global): **$6.6B in 2026 → $38.4B by 2033**, 28.7% CAGR (Market Minds Advisory, Feb 2026). [secondary]
- AI datacenter liquid cooling: **$17.83B by 2036** at 16.9% CAGR (Future Market Insights); direct-to-chip ≈ **47% of revenue in 2026**; hyperscale AI datacenters ≈ **55% of deployment demand by 2026**. [secondary]
- Immersion: single-phase at 120–150 MW annual deployment rate in 2026–2027; immersion <2% of global DC cooling (by IT load) in 2025 → 12–18% by 2031–2035 (analyst roadmap). [secondary — Energy Solutions Intelligence]
- Investment signal: Supermicro cites liquid-cooled rack-scale systems (DCBBS) as a competitive differentiator with margin benefit; Foxconn/Quanta GB300 NVL72 racks ship with complete liquid cooling. [secondary]

---

## 8. Conflicting estimates — flagged explicitly

| # | Conflict | Sources |
|---|---|---|
| 1 | **2026 AI infrastructure spend**: $497B (+~56%, IDC AI Infrastructure Tracker July 2026) vs. **$487B** (IDC via Data Center Knowledge, Sept 2026). Likely different tracker vintages; both secondhand. | [secondary] vs [secondary] |
| 2 | **2026 AI server shipment growth**: >20% (TrendForce Oct 2025) → >28% (Jan 2026) → ~28% (Apr 2026) → **nearly 31%** (Aug 2026). Use the latest (Aug 2026) as current. | [official] series |
| 3 | **Total 2026 server shipment growth**: 12.8% (TrendForce Jan 2026) vs. **~13%** (Apr 2026, revised down from ~20% on component constraints) vs. 19.2% (DIGITIMES-adjacent compilation, unverified) vs. 12.8% cited again in invest-bud compilation. The 19.2% figure is an outlier — treat as [unverified]. | [official]/[unverified] |
| 4 | **ASIC vs. GPU in 2026**: TrendForce Oct 2025 says ASIC shipments "expected to surpass those of GPUs" by 2026; TrendForce Jan 2026 says GPUs = **69.7% of AI server shipments**. Possible scope mismatch (AI-chip units vs. AI-server shipments) — do not merge. | [official] vs [official] |
| 5 | **Dell FY2026 Q4 AI figures**: Data Center Knowledge: $9B AI server revenue (+342% YoY), $43B backlog; July-2026 compilation: +39% total sales with AI server revenue "expected to double next year." Different framings of the same quarter — the DCK numbers are more specific and consistent with the FY27 trajectory ($74B guide). | [secondary] vs [secondary] |
| 6 | **AI server market $ value 2026**: IDC AI-infrastructure $487–497B vs. syndicated AI-server estimates $157–224B vs. TrendForce revenue-implied growth. These measure different scopes (full AI infrastructure incl. storage/network vs. servers only vs. AI-server systems). Not directly comparable. | various |
| 7 | **Next Platform's IDC-derived unit math** ($140K avg × 840K units ≠ stated $137.35B total) is internally inconsistent; use IDC's official $87.4B/$27.5B revenue splits and treat unit/ASP derivations as estimates. | [independent] |
| 8 | **Foxconn "40% share"**: company claim ("global AI server market") vs. brokerages ("NVIDIA AI server racks") — different denominators; the latter projects 50% in 2027. | [vendor-reported] vs [secondary] |

---

## 9. Gaps and follow-ups
- No open-source **AI-server-only vendor ranking** (Dell/Supermicro/HPE/Lenovo/Giga Computing/ASUS/Inspur shares of AI servers specifically) found; IDC/TrendForce/DIGITIMES carry this in paywalled trackers. Giga Computing, ASUS, and Inspur 2026 AI-server figures not captured in open sources.
- No clean analyst figure for **8-GPU server unit volumes** specifically; available proxies: AMD GPU units (UBS), IDC GPU-accelerated server revenue/units, NVIDIA rack estimates (brokerages), pricing/residual-value trackers.
- **2027 forecasts**: CSP CapEx >$850B–$1T+ (TrendForce/industry); AI servers ~19% of shipments by 2027 (TrendForce); NVIDIA rack shipments 100–110K (brokerages); TPU racks 105K (brokerages). No full 2027 AI-server revenue forecast found in open sources.
- ServeTheHome carried no 2026 AI-server market-size figure in the sources surveyed; Data Center Dynamics coverage was earnings/analyst-commentary rather than primary market sizing.
- Figures in this draft reflect research conducted **September 22, 2026**; IDC Q3 2026 tracker (expected ~Dec 2026) and Gartner's October 2026 IT-spending update will supersede several numbers above.

---

## §5 Open verification items (consolidated)

All items below are explicitly unresolved as of September 22, 2026 and must not be treated as facts.

### Vendors — Supermicro & Dell
1. **No published list prices** from either vendor for AI servers — all sales are quote-based.
2. **xAI ~$5B GB200-server deal** (Bloomberg/Reuters, Feb 2025 report) — finalization never publicly confirmed in 2026 sources.
3. Several announcement dates flagged **[unverified date]** in the draft (B300 4U/2OU platform, Blackwell Ultra PR, Vera Rubin GTC item, XE9680L re-announcement).
4. **$20B DataVolt/Saudi campus** figure — market commentary only, not a contracted figure.
5. Corporate-level DOJ/SEC exposure and securities class-action outcomes (Supermicro) still open.
6. Dell Q4 FY26 AI revenue cited as **$9.0B and $9.5B** across sources — conflict.
7. **No 2026 CoreWeave deal found for either vendor** — absence of evidence, not evidence of absence.

### Vendors — HPE, Giga Computing, ASUS, MSI, xFusion
8. **No public pricing** for any AI server (quote-based across the industry).
9. "Cray XD675" and ASUS "ESC N8A-E12" **not corroborated as current 2026 SKUs**.
10. MSI customer/revenue detail — not located.
11. Giga Computing named customers beyond RIKEN — not located.
12. **xFusion product launches** — thin coverage in English; needs Chinese-language follow-up. No source confirms xFusion itself is on the BIS Entity List (Huawei is, since 2019).

### Networking
13. **Arista FY2026 guide conflict**: $11.5B (official) vs $12.6B (single secondary outlet, Sep 19, 2026) — flagged unverified.
14. UEC 1.0.3 currency as of July 2026 — community wiki, not verified against ultraethernet.org.
15. No Quantum-3 product evidence found (roadmap appears to skip to Quantum-X1600).
16. Spectrum-XGS exact announcement date approximate.
17. NVIDIA silicon-photonics H2 2026 shipping is a vendor roadmap claim.
18. Transceiver TAM estimates conflict widely across firms.
19. No SKU-level switch shipment volumes exist publicly.

### Market figures
20. **IDC 2026 vintage mismatch**: $487B vs $497B across publications.
21. TrendForce's ASIC-surpassing-GPU claim vs its own 69.7% GPU shipment figure — scope ambiguity.
22. Dell FY26-Q4 framing discrepancies across sources.
23. Foxconn's "40% of global AI server market" — denominator ambiguity.
24. Next Platform unit math internally inconsistent — use with caution.
25. Syndicated market-size estimates ($157–224B) vs IDC/TrendForce scope — not comparable; scope mismatch.
26. **No open-source AI-server-only vendor ranking** (Giga Computing/ASUS/Inspur figures absent — paywalled in IDC/TrendForce/DIGITIMES trackers).
27. No clean 8-GPU unit volume figure found; no full 2027 AI-server revenue forecast found.

---

*Collection metadata: 4 parallel research waves completed September 22, 2026. All claims carry provenance tags; all working drafts retained alongside this file. Read-only web research; nothing sent externally.*
