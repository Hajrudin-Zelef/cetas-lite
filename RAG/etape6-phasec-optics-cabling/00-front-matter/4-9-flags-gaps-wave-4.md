---
id: etape6-phasec-optics-cabling/00-front-matter/4-9-flags-gaps-wave-4
title: "4.9 Flags & gaps (Wave 4)"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2017-11", "2018-09", "2026-06", "2026-09-22"]
keywords: ["consumer", "distribution", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [738, 783]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: c7bc3f54d328b7f13b27a3aa6e1b9ea7705c3deae6f27b4a364fb3932deaa12e
---

# 4.9 Flags & gaps (Wave 4)

### 4.9 Flags & gaps (Wave 4)
1. **Task claim correction:** 200G/400G-SR4.2 (BiDi/SWDM2) uses a **12-fiber MPO** (8 fibers active, 2 λ/fiber), not MPO-24 [official — IEEE 802.3 NGMMF, Arista]. MPO-24 is for backbones/breakouts.
2. **OM3 100G-SR4 reach conflict:** 70 m (FS.com) vs 100 m (IEEE 802.3bm / other vendors) — flagged above.
3. **MM return-loss conflict:** ≥20 dB (FS.com) vs ≥30 dB (Holight) for UPC multimode MPO — vendor-claimed, both flagged.
4. **Missing current prices:** Panduit HD Flex cassettes/enclosures (quote-only distribution); Corning EDGE-16 module SKUs/pricing; armored MPO patch cords (no retail figure found) — marked as gaps, not invented.
5. **OS1a specifics** (tight-buffered indoor SMF, ~10 km @ 10GBASE-LX4) are tagged [secondary/unverified] — only partially sourced this pass.
6. **Price comparability:** FS.com regional stores, lengths, jackets, and IL grades differ; Corning list prices are 3–5× FS street prices for comparable 12F modules ($506 vs $110–$139) — compare only within vendor and config.
7. No information in this pass is newer than ~Sept 2026; crawler dates on pricing pages range 2025-2026 — prices should be re-verified at purchase time.

---

## Wave 5 — Copper structured cabling: Cat6A/Cat8 (research date 2026-09-22)

**Compiled 2026-09-22. All facts tagged: [official] [vendor-reported] [independent] [secondary] [unverified]. Prices are USD unless noted; "observed Sept 2026 via web search crawls" with crawl dates where known. Cross-region/cross-tier prices are flagged as non-comparable.**

### 5.1 Category standards

#### 5.1.1 Governing standards bodies and documents
- **ANSI/TIA-568-C.2-1 (2016)**: Addendum 1 to TIA-568-C.2; this is the document that specified Category 8 twisted-pair cabling in North America [secondary — FS.com blog, cablewholesale.com product copy].
- **ANSI/TIA-568.2-D (published September 2018)**: superseded TIA-568-C.2. Formally recognized 28 AWG patch cords as standards-compliant and carried forward Category 8 requirements [vendor-reported — Panduit TSB-184-A guidance doc, panduit.com]. Note: TIA's Category 8 only recognizes the Class I (RJ45) solution, not Class II [secondary — Cabling Installation & Maintenance, quoting CommScope white paper].
- **ISO/IEC 11801-1:2017 (major revision, November 2017)**: unified commercial/home/industrial requirements; defines channel classes: Class EA = up to 500 MHz using Cat6A components; Class I = up to 2000 MHz using Cat8.1; Class II = up to 2000 MHz using Cat8.2 [secondary — Wikipedia ISO/IEC 11801, updated Sept 2026].
- **IEEE 802.3an** = 10GBASE-T (10G over twisted pair, 100m on Cat6A) [secondary — l-p.com FAQ, Jan 2026]; **IEEE 802.3bq** = 25GBASE-T and 40GBASE-T [vendor-reported — FS.com blog 2019; qsfptek]; **IEEE 802.3bz (2016)** = 2.5GBASE-T/5GBASE-T ("NBASE-T") [secondary — Wikipedia].
- Connector detail standards: IEC 60603-7-81 (Cat8.1 RJ45-style interface); IEC 61076-3-104 (Siemon TERA connector); GG45 per IEC 60603-7-71 [secondary — eeworldonline; cablinginstall citing CommScope].

#### 5.1.2 Cat8 definition
- 2000 MHz tested bandwidth; supports 25GBASE-T and 40GBASE-T; **maximum 30 m channel in a 2-connector configuration**; **24 m maximum permanent link** [secondary — TTI Fiber blog (updated ~Sept 2026), Cabling Installation & Maintenance]. The 30 m cap applies only at 25/40G; the same cable carries 10GBASE-T and below over a full 100 m channel [secondary — TTI Fiber blog].
- Shielding is mandatory: unlike Cat5e–Cat6A (UTP or shielded permitted), Category 8 requires shielded construction; most likely construction is 22 AWG S/FTP [secondary — cablinginstall "Category 8 Questions Answered", quoting CommScope's Masood Shariff].
- Application target: data-center switch-to-server interconnects (top-of-rack, and end-of-row/middle-of-row) [secondary — cablinginstall; Siemon Cat8.2 spec sheet]. Designed as a copper alternative between twinax DAC (≤7 m) and fiber (up to 500 m) [secondary — cablinginstall].

#### 5.1.3 Cat8.1 vs Cat8.2 vs Class I vs Class II
- **Cat8.1 / Class I**: traditional RJ45 interface (backward-compatible RJ45 form factor, "vastly improved performance relative to Category 6A connectors," specified to 2000 MHz); fully backward compatible at mated-plug level — a Cat8.1 channel can use Cat6A patch cords until active equipment is upgraded [secondary — cablinginstall/CommScope].
- **Cat8.2 / Class II**: non-RJ45 connectors — GG45, ARJ45, or TERA (all permitted under ISO 11801/ISO 24764 for Cat7A carry-over); three connector families are **not mutually compatible** nor RJ45-compatible, requiring hybrid equipment cords; ISO TR 11801-99-1 does not prescribe a connector type [secondary — cablinginstall citing CommScope white paper]. GG45 jack mates with RJ45 plug (jack-side backward compatible), but GG45 plug does NOT mate with RJ45 jack [secondary — same source]. TERA standardized in IEC 61076-3-104; per Siemon, TERA is used for some CAT-8.1 applications and required for Cat8.2 2 GHz (more RF-friendly, lower insertion loss) [secondary — eeworldonline, 2023].
- CONFLICT FLAG: a consumer guide (baudcom, updated ~June 2026) maps "25 Gbps = Cat8.1 / 40 Gbps = Cat8.2". The ISO standard specifies **both** classes to 2000 MHz and both support 25 and 40GBASE-T [secondary — Wikipedia ISO/IEC 11801]; treat the 8.1=25G/8.2=40G mapping as vendor-simplification, not standards text.

#### 5.1.4 Comparison table — Cat6 vs Cat6A vs Cat8
| Parameter | Cat6 | Cat6A | Cat8 |
|---|---|---|---|
| Bandwidth | 250 MHz [secondary — Eaton/scribd table] | 500 MHz [secondary — ISO/IEC 11801, multiple] | 2000 MHz [official-class — ISO/IEC 11801-1; TIA-568-C.2-1] |
| Max data rate | 1 Gbps (100 m); 10 Gbps to ~55 m [secondary — TTI Fiber blog] | 10 Gbps (100 m) [secondary — multiple] | 25/40 Gbps (30 m) [secondary — multiple] |
| Max channel distance at top speed | ~55 m @10G [secondary]; some sources say 37 m @10G — CONFLICT, see flags | 100 m @10G [secondary] | 30 m @25/40G (24 m permanent link); 100 m @10G and below [secondary — TTI Fiber] |
| Shielding | UTP or shielded permitted [secondary] | UTP or shielded permitted; alien-crosstalk spec is the key delta [secondary] | Shielded mandatory (S/FTP typical) [secondary — cablinginstall] |
| PoE support | 802.3af/at; 802.3bt supported but thermally less ideal in bundles [secondary] | Full 802.3af/at/bt incl. Type 4 90 W; recommended cable for PoE++ [vendor-reported — FS.com, Cisco, VersaTek] | 802.3af/at/bt per vendor spec sheets [vendor-reported — FS.com, Siemon spec sheet lists af/at/bt Type 3] |
| Typical conductor | 23/24 AWG solid (horizontal) [secondary] | 23 AWG solid (horizontal); 26/28 AWG patch [secondary/vendor] | 22–24 AWG solid (Siemon Cat8.2: 22 AWG solid, OD 8.3–8.5 mm) [vendor-reported — Siemon spec sheet] |

### 5.2 Cable construction

