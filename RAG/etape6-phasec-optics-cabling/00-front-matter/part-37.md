---
id: etape6-phasec-optics-cabling/00-front-matter/part-37
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 37)"
domain: front-matter
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1350, 1355]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 80cf453f829d4cc1338d13ed11e1c678348291b32cc32f7d9490597a9a803e10
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 37)

- **Tier 1 (OLTS)**: Optical Loss Test Set = calibrated light source + power meter; measures end-to-end insertion loss and length and verifies polarity (duplex and MPO). The 1-jumper, 2-jumper, 3-jumper reference methods are specified in TIA-568 and ISO/IEC 14763-3; TIA prefers 1-jumper, ISO/IEC 14763-3 prefers 3-jumper [official: ISO/IEC 14763-3, TIA-568-C.3; secondary: cablinginstall.com].
- **Tier 2 (OTDR)**: Optical Time Domain Reflectometer injects short laser pulses and plots backscatter vs time — locates connectors, splices, breaks and quantifies each event's loss; required for troubleshooting and documented acceptance of installed plant. MM tested at 850/1300 nm, SM at 1310/1550 nm [secondary: softing.com, cablinginstall.com]. New IEC 61280-4-5 (per 2026 trade press) covers OTDR attenuation methods for installed cabling [secondary]. ISO/IEC 14763-3:2017 was updated (2024 edition discussed in trade press) [secondary].
- **Cleaning/inspection**: inspect-clean-inspect per IEC 61300-3-35 (pass/fail criteria by endface zones); automated digital inspection scopes give pass/fail; one-push MPO/MPO-12/24 cleaners ~US$90–99 at FS.com [official: IEC 61300-3-35; vendor-reported: FS.com; secondary: cablinginstall.com].
- **Loss budgets**: TIA-568-C.3 / 568.3-D maximum per mated connector pair: **0.75 dB** (MM and SM); return loss minimum 20 dB (MM) / 26 dB (SM); splice loss maximum 0.3 dB [official: ANSI/TIA-568-C.3 / 568.3-D; secondary: ttifiber.com]. Worked example: 500 m SM link with 2 connectors ≈ 1.3–1.7 dB max depending on method [secondary]. Vendor ultra-low-loss MPO cassettes are specified at ~1.0 dB per cassette [vendor-reported: TE/CommScope-style handbook].
- **Relevant standards list**: ANSI/TIA-568.3-D (optical fiber cabling and components), TIA-568-C.0 (generic telecommunications cabling, polarity methods A/B/C), TIA-604 series FOCIS intermateability documents (FOCIS-3 LC, FOCIS-5 MPO-12, FOCIS-10, FOCIS-18 MPO-16), IEC 61754-7 (MPO connector interface), ITU-T G.652/G.657 (SM fiber), ITU-T G.651.1 (MM fiber), ISO/IEC 11801 (generic premises cabling), ISO/IEC 14763-3 (testing of installed optical-fiber cabling), IEC 61300-3-35 (endface cleanliness), IEEE 802.3 (Ethernet optical PHY reaches: 802.3ba 40/100G, 802.3bs 200/400G, 802.3df 800G/1.6T) [official standards references].

