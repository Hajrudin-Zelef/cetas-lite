---
id: etape6-tracka-cisco-juniper/01-part-1-cisco/8-400g-800g-portfolio-optics-cpo
title: "8. 400G/800G portfolio, optics, CPO"
domain: part-1-cisco
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Nvidia"]
dates: ["2026-05-04", "2026-05-13", "2026-06-23"]
keywords: ["cpo", "optics", "agentic", "amd", "benchmark", "benchmarks", "dci", "dsp", "ethernet", "gpu", "hbm", "hyperscaler"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [173, 225]
section: "PART 1 — CISCO"
sha256: af96ce1717f3a586259f7db8f3a5dfa05babd62a93962ac5790cdd4bc271d7d9
---

# 8. 400G/800G portfolio, optics, CPO

## 8. 400G/800G portfolio, optics, CPO

### Pluggable optics launched Feb 2026 (with G300)
- **1.6T OSFP** pluggable optics: ultra-high-bandwidth for AI scale-out; targets 1.6T switch-to-NIC links and 1.6T/800G/400G/200G switch-to-server links; high performance + reliability [official]. Both ship **this calendar year** per Cisco (CRN) [independent].
- **800G Linear Pluggable Optics (LPO)**: **50% lower optical-module power vs. retimed modules**; with LPO-supporting N9000/8000 systems, customers can cut **overall switch power by 30%** [official]. Built with Cisco silicon photonics technology [secondary via Zacks].
- AI POD 2-8-9-400 BOM shows: QSFP-400G-DR4 (400G QSFP112), QDD-8X100G-FR dual-port 800G, QSFP-200G-SR4-S in use [official via CVD].

### Optical transport (DCI/backbone, context)
- **Cisco Open Transport 3000 Series** multi-rail open line system (2026): Cisco claims 75% power reduction and 80% rack-space reduction per rail vs. prior single-rail systems (C-band and C&L-band) [independent via Light Reading]. Primarily hyperscalers/SPs; "14× capacity" framing vs. legacy DCI [independent].
- **NCS 1014** upgraded with 800G pluggable transponder line card: 12.8T capacity (16× 800G in one card), 50% less rack space, 38% less power vs. prior 400G generation [independent via Light Reading].
- **New 100ZR coherent pluggable** using Cisco's own silicon photonics + DSP; 100–150 km reach; extends Routed Optical Networking into access/edge [independent via Light Reading].
- Acacia (Cisco's coherent unit): record Q3 FY2026 — **>750,000 400G + 40,000 800G coherent pluggables shipped**; five new hyperscaler design wins (incl. Silicon One P200/G200 systems); CEO: Acacia to **grow 200% in FY2026** [independent via Lightwave]. "Commanding share" of the ZR/ZR+ market [independent].

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

