---
id: etape6-phasec-optics-cabling/00-front-matter/part-22
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 22)"
domain: front-matter
role: reference
task: reference
actors: ["EU", "United States"]
dates: []
keywords: ["regulation"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [784, 818]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: f9ffd02df719eab2479d85bfb8169226bb9f0fc9c4c51731e2834fb2743c0e91
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 22)

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

