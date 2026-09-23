---
id: etape6-phasec-optics-cabling/00-front-matter/4-9-flags-gaps-wave-4
title: "4.9 Flags & gaps (Wave 4)"
domain: front-matter
role: reference
task: reference
actors: ["EU", "United States"]
dates: ["2017-11", "2018-09", "2026-06", "2026-09-22"]
keywords: ["consumer", "distribution", "pricing", "regulation", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [738, 818]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 5adc408b882038f6c53dcaa24ac1b87b33a4370f1e6683efc9ef87c20ed2a038
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

#### 5.2.1 Shielding nomenclature (ISO/IEC 11801 code: XX/YYY — XX = overall cable shield, YYY = per-pair shield)
- **U/UTP**: unshielded/unshielded; "global default for Cat5e and Cat6" [independent — network-data-cabling.co.uk, ACCL].
- **F/UTP**: overall foil, unshielded pairs (aliases: FTP, STP, ScTP) — "the modern minimum for Cat6A in Europe" [independent — ACCL]. Class I Cat8.1 minimum construction per early ISO drafts was F/UTP [secondary — CommScope white paper via technologydecisions.com.au].
- **U/FTP**: no overall shield, foil per pair — "common in slim Cat6A where diameter matters" [independent — ACCL].
- **S/UTP**: overall braid, unshielded pairs [independent — accesscomms.com.au].
- **SF/UTP**: overall braid + foil, unshielded pairs — industrial environments [independent — accesscomms; comx-computers].
- **S/FTP**: overall braid + foil per pair — "heavy-duty option for Cat7A & Cat8"; aliases SSTP/SFTP/PiMF [independent — ACCL; accesscomms]. Siemon Cat8.2 cable is S/FTP with bonded-polyester individual pair foils + high-coverage outer braid [vendor-reported — Siemon spec sheet].
- **F/FTP**: overall foil + foil per pair — "common in 10GBaseT applications" [independent — accesscomms].
- **SF/FTP**: braid + foil overall + foil per pair — "highest level of protection available" [independent — accesscomms; c-techsolutions].
- Code letters: U = unshielded/unscreened, F = foil, S = braided shielding, TP = twisted pair [independent — github.com/barnumbirr wiki].

#### 5.2.2 Solid vs stranded
- **Solid**: single solid wire per conductor; for fixed installations — in walls, ceilings, patch panels [independent — comx-computers.co.za].
- **Stranded**: multiple thin strands per conductor; flexible, longer flex-life; for patch cables and areas with frequent movement [independent — comx-computers; FS.com Cat6A patch cord copy].
- Standard channel model: 90 m solid horizontal + 10 m stranded patch [secondary — Quabbin Wire technical brief].

#### 5.2.3 Conductor gauge (AWG)
- Horizontal cabling: 22–24 AWG conductors permitted; patch cords may be 26 or 28 AWG [secondary — TTI Fiber blog summarizing TIA-568 installation rules].
- 28 AWG patch cords formally recognized as standards-compliant by ANSI/TIA-568.2-D (Sept 2018); TIA working on TSB-184-A-1 addendum for 28 AWG power-delivery guidance [vendor-reported — Panduit].
- Real examples: Cat6A U/UTP bulk = 23 AWG solid (OD 7.30 mm) [vendor-reported — RTC-Electronics, crawled Sept 2026]; FS.com Cat6A S/FTP bulk = 23 AWG solid [vendor-reported — FS.com]; Siemon Cat8.2 = 22 AWG solid [vendor-reported — Siemon]; FS.com Cat8 patch cords = 24 AWG class/28 AWG per blog table (FS blog table says "28" — CONFLICT with typical 24–26 AWG patch construction; see flags); FS Cat6A patch = 26 AWG stranded [vendor-reported — FS.com].
- Cisco guidance for 802.3bt Type 4: 22 AWG conductors recommended where maximum power efficiency and ≤1 Gbps data are the goals; Cat6A with 23 AWG or larger for higher data speeds [vendor-reported — Cisco UPOE/802.3bt doc via scribd].

#### 5.2.4 Fire/jacket ratings
- **CM/CMG** (general purpose): UL-1581; will burn, partially self-extinguishing; workstation/patch-cord use; not for walls/ceilings of commercial buildings [independent — cablesplususa.com].
- **CMR** (riser): UL-1666; vertical tray/floor-to-floor; must self-extinguish and prevent flame traveling up the cable [independent — cablesplususa].
- **CMP** (plenum): NFPA-262 / UL-910 (Steiner Tunnel Test); only cable allowed in air-handling plenum spaces (drop ceilings, raised floors); must self-extinguish and not reignite; typically FEP (fluoropolymer) jacket; still emits some toxic fumes but low smoke [independent — cablesplususa; tektel.com]. Hierarchy: CMP may substitute for CMR; CMR may never go in a plenum [secondary — FS.com blog; Medium/Davina, Mar 2026].
- **CMX**: limited/outdoor use; UV-resistant LLDPE jacket; NOT fire-rated for indoor riser/plenum; cannot replace CM/CMR/CMP indoors [secondary — FS.com blog].
- **LSZH/LSOH** (Low Smoke Zero Halogen): halogen-free thermoplastic compounds; emits minimal toxic/corrosive gas when burned; prioritized by EU Construction Products Regulation (CPR) and required in much of Europe/Asia for public buildings, transport, ships, tunnels [secondary — FS.com blog; tektel.com]. **Critical limitation**: LSZH cables generally do NOT pass the US UL-910 Steiner Tunnel plenum test — do not use in a US plenum ceiling unless dual-rated CMP-LSZH [independent/secondary — Medium/Davina Mar 2026; cablesplususa]. Independent fire-test reporting (1xtechnologies, citing BRE/FRS-style tests) found LSZH burned readily with higher heat release than CMP in plenum-ladder tests — treat as [secondary/independent, vendor-hosted (1xtechnologies is a cable vendor)].
- CCA (copper-clad aluminum) conductors are NOT permitted under the electrical code, "however the cable tests at the factory" [secondary — TTI Fiber blog]. Pure bare copper required (e.g., Monoprice cites UL 444 / NEC TIA-568-C.2 requiring pure copper) [vendor-reported — Monoprice/Newegg listing].

#### 5.2.5 Alien crosstalk — why Cat6A/Cat8 need shielding or bulk
- Alien crosstalk (crosstalk between adjacent cables in a bundle) is the defining new parameter for Cat6A vs Cat6; unshielded Cat6A manages it via larger diameter/increased pair separation, shielded Cat6A via foil/braid [secondary — TTI Fiber blog; vendor — FS.com patch cord copy "superior alien crosstalk (AXT) suppression"].
- Evidence of bulk: Cat6A 23 AWG U/UTP bulk OD ≈ 7.3 mm [vendor-reported — RTC-Electronics] vs Cat8.2 22 AWG S/FTP OD 8.3–8.5 mm [vendor-reported — Siemon].
- Mitigations for UTP Cat6A: staggered patch-panel port layouts "eliminate alien crosstalk" [vendor-reported — D-Link NPP-6A1BLK241 copy]; U/FTP construction for slim Cat6A [independent — ACCL].

