---
id: collect-261001-cisco/cisco/etape6-tracka-cisco-juniper-4
title: "Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia"]
dates: ["2026-05-04", "2026-05-13", "2026-06-23", "2026-07-08"]
keywords: ["agentic", "amd", "benchmark", "benchmarks", "cpo", "ethernet", "gpu", "hbm", "hyperscaler", "lpo", "neocloud", "nvidia"]
source: docs/RAG/collect-261001-cisco/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [186, 245]
sha256: 567ad255c6e0f8d00bbb0c967ad97d866356a7007d94636ca5bc353610d05a10
---

# Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026

### Co-packaged optics (CPO)
- **No Cisco CPO product announcement found in 2026 research.** Cisco's 2026 posture is pro-**pluggable**: 1.6T OSFP + 800G LPO (headlined "Cisco bets big on pluggables" [independent via Light Reading]). Cisco's CPO work appears limited to groundwork ("laying the groundwork for solutions that will enable seamless and scalable AI at unprecedented densities" [secondary via JR Sekwele]). CPO commercial debut in 2026 is industry-wide (NVIDIA Quantum-X/Spectrum-X Photonics cited as leading; Broadcom Bailly) [secondary]. Cisco's near-term efficiency play = **LPO**, not CPO. Flag as a competitive gap vs. NVIDIA/Broadcom CPO roadmaps.

### 800G/400G switch portfolio status
- 400G: Nexus 9364D-GX2A (64p 400G), 9332D-H2R (32p deep-buffer 400G, HBM), 9364D-GX2A-class 25.6T boxes; breakouts down to 10G [official].
- 800G: N9364E-SG2-Q (64× 800G QSFP-DD), N9364E-SP2R-X (P200, 51.2T), N9364F-SG3 (G300, 102.4T), 800G fixed (XF3-class DCN licensing) [official/vendor-reported/independent].

---

## 9. Cisco vs. competitors — AI networking positioning (2026)

- **Market position**: Cisco remains the overall Ethernet switch/router market leader but **lost early AI-cluster share to Arista and NVIDIA** in the hyperscaler buildout [secondary via Motley Fool]. 2026 narrative: Cisco "finally winning over hyperscalers" — $9.3B FY2026 hyperscaler AI orders as evidence [secondary].
- **Cisco's claimed differentiation**: (a) enterprise + SP + hyperscaler breadth vs. Arista's hyperscaler focus; (b) multi-vendor interoperability vs. NVIDIA Spectrum-X's stack lock-in; (c) lifecycle support, integrated security (Hypershield/AI Defense), cross-domain policy; (d) largest on-chip buffers, Intelligent Packet Flow [vendor-reported/independent].
- **Named competitive displacements**: Q4 FY2026 earnings coverage notes Cisco said it is **"taking customers from rivals"** (one anecdote), and projected "market share gains" into FY2027 — but the sources caution one anecdote ≠ trend [secondary via Benzinga cloudfront mirror]. **Named displaced competitor / customer not found** — flag as gap.
- **Benchmarks**: No head-to-head 2026 benchmarks (e.g., vs. Arista 800G Etherlink or NVIDIA Spectrum-X) were found in the covered sources. Claims of "boost GPU utilization / improve job completion times" are Cisco marketing without published comparative numbers [official claims; independent verification absent]. Flag as gap.
- **Analyst framing**: Cisco's AI pitch targets the "regulated middle" — enterprise, neocloud, sovereign clouds — where lifecycle support and security matter [secondary via CryptoDaily]. McKinsey-cited figures (Cisco VP Jake Katz blog): $4.7T global DC IT-equipment spend 2025–2030; hyperscalers >60%; neoclouds ~17% growing to >30% in ten years [independent via Network World]. Dell'Oro's Sian Morgan and ZK's Zeus Kerravala are skeptical of Arista's campus-enterprise push — indirect tailwind framing for Cisco in enterprise [independent].
- **Arista counter-moves (2026 context)**: acquired VeloCloud (SD-WAN) from Broadcom; targeting $1.25B campus revenue next year [independent].
- **NVIDIA counter-position**: Spectrum-X Photonics (up to 400 Tbps total, 2,048 ports) and Quantum-X Photonics (144× 800G InfiniBand, CPO) positioned as the AI-native benchmark [secondary].

---

## 10. Event timeline (2026)

| Date | Event | Source |
|---|---|---|
| Feb 10, 2026 | Cisco Live EMEA, Amsterdam: Silicon One G300 (102.4T), new Nexus 9000/8000 systems, 1.6T OSFP + 800G LPO optics, Nexus One + Unified Fabric, native Splunk integration | [official/independent] |
| ~Apr 2026 | Intent to acquire Galileo Technologies (undisclosed) | [secondary] |
| May 4, 2026 | Intent to acquire Astrix Security (~$400M per press; official terms undisclosed) | [independent/secondary] |
| May 13, 2026 | Q3 FY2026 earnings: $15.84B rev; $5.3B AI orders YTD; FY target raised to $9B; ~4,000 job cuts | [independent] |
| June 2–4, 2026 | Cisco Live, Las Vegas: Cisco Cloud Control + AgenticOps/AI Canvas (CA June 2, GA planned July); Agentic Actions beta (Meraki); Digital Twin alpha July; Multicloud Fabric | [independent] |
| June 23, 2026 | CEO: Acacia to grow 200% in FY2026; record 400G/800G coherent quarter | [independent] |
| Aug 12, 2026 | Q4 FY2026 earnings: $17.3B rev (+18%); $4B Q4 AI orders; $9.3B FY AI orders; $4B FY AI revenue; FY27 guide $72.2–73.4B; $7.5B FY27 AI-infra revenue | [independent] |
| Aug 25, 2026 | Cisco + Teleport "Infrastructure Identity" partnership; Cisco largest strategic investor | [secondary] |
| Aug 31, 2026 | AMD/Cisco/HUMAIN JV first segment operational (MI355X) | [independent] |
| Sept 15, 2026 | Cisco AI POD for Splunk announced at Splunk.conf, Denver | [secondary] |
| Sept 17, 2026 | Cisco joins Verizon 6G Innovation Forum | [secondary] |
| ~Sept 21, 2026 | Internet2/NRP: two AI PODs (8× H200 total) deployed | [official via PRN] |

---

## 11. Gaps and uncertainties (explicit flags)

1. **G300 on-chip buffer size**: "industry's largest on-chip buffer" claimed but no MB/GB figure found. Not verified against independent spec.
2. **G500**: no 2026 announcement found. Only G300 confirmed new for 2026.
3. **204.8T**: no Cisco 2026 announcement found; top public density is 102.4T.
4. **CPO**: no Cisco CPO product in 2026 found; Cisco is betting on pluggables (1.6T OSFP, 800G LPO) while NVIDIA/Broadcom lead CPO. Any Cisco CPO roadmap timing is unknown from these sources.
5. **Named hyperscaler customer wins**: Cisco reports AI orders in aggregate; no named hyperscaler logos or named displacement wins verified.
6. **Head-to-head benchmarks** vs. Arista/NVIDIA: none found.
7. **Earnings figures**: sourced from multiple consistent trade-press reports, but the actual Cisco IR release text was not directly read in this pass — treat numbers as independently-reported-from-earnings.
8. **AI POD for Splunk** details: sourced from secondary outlets; original Cisco announcement not directly verified.
9. **Astrix price ($400M)**: press-reported (Calcalist/CTech); Cisco did not disclose. Talks-stage $250–350M also reported. Use "$250–400M range per press; officially undisclosed."
10. **Nexus Dashboard 4.2 AI job-monitoring details**: from a secondary blog; not verified against Cisco docs.
11. **List prices**: only third-party retailer pricing found ($128,731.99 for N9K-C9332D-H2R via Zones [secondary]); no Cisco list prices found.
12. **Internet2/NRP campus names**: not captured in this pass.
13. **R&A partnership** (July 8, 2026, Official Network Supplier for The Open etc.) is real [official] but a sports-hospitality deal, not data-center — excluded from main narrative.

---

## 12. Key sources (verbatim URLs)

