---
id: etape6-phasec-optics-cabling/00-front-matter/9-7-termination-testing
title: "9.7 Termination & testing"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "United States"]
dates: ["2022-08", "2023-03", "2025-03", "2026-05", "2026-09", "2026-09-22"]
keywords: ["asic", "attention", "capex", "cost", "datacenter", "distribution", "ethernet", "full-duplex", "inference", "latency", "optics", "parameters"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1457, 1491]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 54070a5b9ec2d4fad6f01041be357f3a3ef05581810e12fb8d51baf7fe4702e2
---

# 9.7 Termination & testing

- **DAC (twinax)**: factory-terminated point-to-point assembly — twinax cable with SFP/QSFP housings; passive ≤3 m (up to ~7 m active); used server↔ToR within/between adjacent racks [secondary — FS.com, QSFPTEK, network-switch.com]. Pros: lowest capex for the link, near-zero power, lowest latency, simple. Cons: fixed lengths, thick/stiff bundles hurt cable management and airflow at density, vendor coding/compatibility testing required, no patching flexibility, range collapses at higher speeds (800G passive DAC ≈3 m, pushing adoption of ACC/AEC/AOC) [secondary].
- **Structured twisted-pair**: horizontal distribution architecture — 90 m permanent link terminated on patch panels/keystones, finished with patch cords; supports moves/adds/changes without re-pulling; cabling is independent of equipment generations (e.g., install Cat8/Class I jack-to-jack now, use Cat6A cords until active gear upgrades) [secondary]. Different use case from DAC: row/room distribution and OOB/management, not intra-rack equipment interconnect.

### 9.7 Termination & testing

- **Permanent link vs channel**: permanent link = ≤90 m installed solid-conductor cabling between two connection points (tested at the jack/panel); channel = ≤100 m including patch cords and equipment cords [official rule — via TIA-568 explainer, secondary]. Certification is performed against link or channel limits accordingly [secondary].
- **Fluke DSX CableAnalyzer** (DSX-8000 to 2000 MHz for Cat8/Class I/II; DSX-5000 to 1000 MHz for Class FA and below) certifies TIA Cat 3–8 and ISO Class C–I/II [official — Fluke datasheet]. Measured parameters: wire map, length, propagation delay, delay skew, DC loop resistance, pair-to-pair/pair resistance unbalance (PoE-relevant), insertion loss, return loss, NEXT, PSNEXT, FEXT, ACR-N, ACR-F (ELFEXT), PSACR-F, alien crosstalk PSANEXT / PSAACR-F, plus balance metrics TCL, ELTCTL, CDNEXT, CMRL [official — Fluke datasheet]. Autotest times: Cat6A/Class EA ≈8 s, Cat8 ≈16 s (DSX-8000) [official].
- **Key parameters explained**: NEXT/PSNEXT — near-end crosstalk within the cable, the dominant 10G limiter along with alien crosstalk; ACR-F (ELFEXT) — far-end crosstalk normalized to attenuation; return loss — reflections from impedance mismatches (kinks, bad terminations, water ingress) [secondary — FS.com Cat6A FAQ]; PSANEXT/PSAACR-F — alien (bundle-level) crosstalk, the reason Cat6A field certification for 10GBASE-T includes alien testing; TCL/ELTCTL — pair balance vs common-mode noise, increasingly important for shielded systems and PoE [secondary].
- **Standards**: ANSI/TIA-568.2-D (2018) — mechanical/transmission requirements for balanced twisted-pair components, recognizes 28 AWG cords [official — via Panduit/Quabbin secondary]; ISO/IEC 11801-1 — generic cabling classes (EA, FA, I, II) [official — via Fluke datasheet test-standard list].
- Field termination notes: 110-style IDC punch-down or toolless keystone jacks, T568A/B wiring, minimize pair untwist; keystone jacks support 22–26 AWG solid conductors typically (e.g., CablesOnline Cat6A jack 22–26 AWG, 6–10 mm OD [vendor-reported]); Cat8 jacks are toolless/shielded (Primus Cable $11.99 [vendor-reported]).

### 9.8 2026 trends — Cat8 adoption & single-pair Ethernet

- **Cat8 adoption**: growing vendor mindshare (FS.com March 2025 deployment guide; QSFPTEK; Tripp Lite Cat8 patch launch) but real-world adoption remains niche [vendor-reported + secondary]. Structural headwinds: (a) very limited 25GBASE-T/40GBASE-T switch and NIC ecosystem versus SFP28/QSFP+ optics and DAC — no evidence found of broad 25/40G BASE-T port availability in 2026; (b) 30 m/2-connector limit confines it to in-room/ToR; (c) hyperscale server attach already moved to DAC/AOC/fiber. Credible use cases: enterprise DC ToR wanting RJ45 continuity with a 25G upgrade path, and cost-vs-fiber under 30 m where BASE-T ports exist [vendor-reported — FS.com, QSFPTEK]. Vendor claim that Cat8 is cheaper than fiber (<30 m) and than twinax (<5 m) [vendor-reported — QSFPTEK/FS.com] should be treated as marketing: a 3 m 10G DAC at US$20 vs a Cat8 cord + port-power premium is not a clean comparison [unverified head-to-head]. Net assessment: Cat8 is a specialty play, not a broad 2026 trend [secondary analysis].
- **Cat7/Cat7A**: effectively a dead branch for new structured installs; vendor attention shifted to Cat8 for short-reach 25/40G [secondary].
- **Single-pair Ethernet**: essentially no datacenter-server relevance; momentum is automotive and industrial. May 2026: Microchip launched LAN878x/LAN888x SPE PHYs (100BASE-T1/1000BASE-T1, MACsec, TSN, ASIL-B) for software-defined vehicles and industrial [secondary — GlobeNewswire]. Industrial variants: 10BASE-T1S (15 m multidrop) and 10BASE-T1L (up to 1000 m) per IEEE 802.3cg, with PoDL power (802.3bu); connectors per IEC 63171-1 (building) / 63171-6 (industrial) [secondary — EE Times]. Datacenter relevance is limited to possible edge/building-automation sensor use; not a server-networking technology [unverified — inference, no DC-SPE source found].

### 9.9 Gaps & conflicts summary (Wave 9)
1. **Missing prices** (not found, not invented): Cat6A bulk 305 m box / per-meter price; Cat8 bulk cable price; Cat8 patch-cord single-unit price; 10GBASE-T RJ45 SFP+ module price; 25/40GBASE-T port pricing. FS.com sells these lines but list prices did not surface in fetched pages.
2. **Cat6 10G distance**: 55 m vs 37 m — condition-dependent (alien crosstalk environment), not a single standards value [conflict flagged in §9.1].
3. **Type 4 PoE wattage**: 90 W PSE / 71.3 W PD per IEEE 802.3bt vs vendor "100 W" claims — the 100 W figure is marketing/chipset-max, not the standard [conflict flagged in §9.5].
4. **Cat8 cost claims** ("cheaper than fiber <30 m / twinax <5 m") are vendor marketing without published head-to-head port-power-inclusive comparisons [flagged in §9.8].
5. **EE Times 10GBASE-T vs SFP+ analysis** is from ~2009; the directional conclusions (power, density, cost drove datacenter choice) remain cited by current vendors, but the 6 W-era PHY numbers are dated — modern PHY figures (2–5 W) from 2024–2026 sources are given alongside [recency caveat].
6. **FS.com prices** carry mixed crawl dates (107–361 days old on some pages; some refreshed within weeks). Treat as list-price data points, not live quotes.

---

## Wave 10 — The 800G Switching Ecosystem (as of September 2026)

*Research date: 2026-09-22. Single-writer wave. Throughput convention: switch-ASIC capacity is stated unidirectional (one-way) unless labeled "full-duplex/bidirectional"; 64 × 800G = 51.2 Tbps one-way = 102.4 Tbps full-duplex, and vendors mix the two — non-comparable figures are flagged inline. Dates given are publication/announcement dates found on cited pages. Tags: [official], [vendor-reported], [independent], [secondary], [unverified].*

### 10.1 Switch-ASIC generation matrix (800G-class)

**Broadcom Tomahawk 5 (BCM78900 family).** 51.2 Tbps, monolithic 5 nm Ethernet switch ASIC exposing 512 × 100G PAM4 SerDes, headline radix 64 × 800GbE / 128 × 400GbE / 256 × 200GbE [official] (nasdaq.com press release). Broadcom announced shipment in August 2022; production-volume shipping began March 2023 [vendor-reported]. AI-oriented features claimed: shared buffering, Cognitive Routing, dynamic flowlet steering, congestion control, hardware failover, RoCEv2 support, Clos/non-Clos topology support [vendor-reported] (telecomlead.com). Tomahawk 5 is the merchant silicon inside Arista 7060X6, Juniper/HPE QFX5240, Edgecore AIS800/DCS560, Micas M2-W6940-64OC, and FS.com N9600-64OD (each mapped under its vendor section; per-section tags apply) [independent].

