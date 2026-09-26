---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/4-uec-ultra-ethernet-consortium-super-ethernet
title: "4. UEC (Ultra Ethernet Consortium) / Super Ethernet"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "China", "Google", "Intel", "Meta", "Microsoft", "Nvidia", "Oracle", "TSMC"]
dates: ["2023-07-19", "2025-06-11", "2025-09", "2026-06", "2026-06-09", "2026-07-16", "2026-08-25"]
keywords: ["ethernet", "amd", "asic", "benchmark", "chiplet", "cpo", "dsp", "foundry", "hyperscaler", "intel", "latency", "lpo"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [807, 853]
section: "Data-Center Networking for AI (2026)"
sha256: 8dc986c60d9c65760806871e8573459a5c34d28104b14609e45bc808e18f0178
---

# 4. UEC (Ultra Ethernet Consortium) / Super Ethernet

## 4. UEC (Ultra Ethernet Consortium) / Super Ethernet

### 4.1 Specification status

- **UEC Specification 1.0:** released **June 11, 2025**; 562-page document defining the full stack (physical layer → transport → application API), including the new **Ultra Ethernet Transport (UET)** protocol with unordered packet delivery, receiver-driven congestion control, native multipath/packet spraying, and operation without mandatory lossless-fabric (PFC) configuration. [independent] (LLM-Systems-Wiki, fetched 2026-08-25, primary sources = UEC press releases + spec)
- Maintenance updates: **1.0.1 (September 2025)** [independent] (networkworld), **1.0.2 and 1.0.3** release notes; **1.0.3 (July 16, 2026) is current** [unverified — sourced from a community wiki; verify against ultraethernet.org]. UEC 2.0 development expected/ongoing (expected 2026 per a June-2025-dated secondary blog; no confirmed 2.0 release found as of Sept 2026). [unverified]
- IEEE 802.3dj (200G/400G/800G/1.6T at 200G/lane) on track for completion late 2026; a 400G/lane follow-on project is being prepared. [independent] (networkworld, quoting Ethernet Alliance chair Peter Jones)
- **Why UEC exists (per consortium framing):** RoCEv2's PFC+DCQCN lossless requirements cause head-of-line blocking, congestion spreading, single-hash-path flows (no packet spraying), and ~30% performance loss at scale vs InfiniBand. UET is a clean-slate transport. [secondary] (Medium/Teradata Labs piece — note: vendor-adjacent authorship, treat the 30% figure as consortium-aligned advocacy, not measured benchmark)

### 4.2 Member momentum

- Founded **July 19, 2023** (Linux Foundation Joint Development Foundation project) by AMD, Arista, Broadcom, Cisco, Eviden/Atos, HPE, Intel, Meta, Microsoft. [independent]
- Consortium has grown to **100+ member companies and ~1,500 participants** (2026). [secondary] (Medium/Teradata Labs)
- Compliance/certification programs "already underway" per UEC launch materials; Keysight's OFC 2026 demo shows the test-equipment ecosystem aligning. [official] (Keysight bizwire via financialcontent)

### 4.3 Products shipping with UEC in 2026

- **Broadcom Thor Ultra NIC** (announced Oct 14, 2025): first NIC built to the UEC 1.0 spec. [secondary]
- **Broadcom Tomahawk Ultra Ethernet switch** (UEC-compliant, LLR + CBFC features): used in the Keysight interop demo at **OFC 2026** — "industry's first public interoperability demonstration of Ultra Ethernet Consortium (UEC) specification with Link Layer Retry (LLR) and Credit-Based Flow Control (CBFC) at 800GE line rate," with Keysight's Interconnect and Network Performance Tester. [official] (Keysight press release)
- LLR (local error recovery, lower tail latency) and CBFC (link-layer flow control) identified as the UEC link-layer capabilities hyperscalers now require for large AI clusters. [official] (Keysight)
- Arista announced support for UEC across its **Etherlink** portfolio. [secondary] (Medium/Teradata Labs — [unverified] detail; no dated Arista press release located)
- **Reality check:** the UEC endgame pieces framing (Medium, June 2026) also notes a Broadcom "UE-compatible Tomahawk switch with 250 nanosecond latency" — [unverified], single-source advocacy; do not quote as fact.
- **Uncertainty flag:** no hyperscaler production deployment of UEC transport (UET) was confirmed in any source; 2026 activity is interop demos, NIC/switch silicon, and test equipment. UEC adoption in production AI fabrics is best described as "pre-commercial."

---

## 5. Optical interconnect

### 5.1 Co-packaged optics (CPO) status 2026

- CPO **entered the market in 2026** through early products from **NVIDIA and Broadcom** (early AI-networking deployments). [secondary] (eetimes, citing analyst Vladimir Kozlov)
- Realistic assessment: "**it's going to take another two years before CPO is really shipping in volumes**" — large-scale CPO deployment realistically **2028–2030** (Yole Group forecast cited in the github research note); current deployments are pilot/proprietary, focused on scale-up AI networks. [independent]
- June 9, 2026: SemiAnalysis report flagged CPO packaging yield risk (~19.4% system yield in a 0.95^32 scenario), triggering optical-stock selloffs (AAOI −17%, COHR −11%, LITE −8%); NVIDIA CTO Gilad Shainer publicly disputed the thesis, stating CPO is already shipping and ramping H2 2026. [secondary] [unverified detail]
- Broadcom Tomahawk 6 supports CPO options; Cisco has CPO prototypes; TSMC COUPE is the foundry vehicle. [official] [secondary]
- Intel's OCI chiplet claim of 5 pJ/bit vs 15 pJ/bit legacy systems is a real physics datapoint (vendor claim). Indium/InP supply concentration (China) is a physical chokepoint; external light sources for CPO are an emerging ~$1B+/yr sub-market. [secondary] (github research note — [unverified] figures)

### 5.2 Linear pluggable optics (LPO)

- LPO removes the DSP from the module; 200G LPO targets: ~10 W/module vs 23–25 W retimed DSP modules and ~16 W linear-receive (RTLR) modules; a 512-port Tomahawk-6 switch saves ~500 W at module level (nearly 1 kW with cooling) moving from retimed to LPO. 200G LPO target reach 500 m (3 dB optical budget); up to ~2 km in low-dispersion DR with sufficient budget. [independent] (itbrief.co.nz citing Semtech)
- Deployment status: LPO "already in use by four or five companies, including Oracle"; Google has started deploying **1.6T linear receive optics** and is preparing volume rollouts in 2026; millions of LPO modules expected to ship in 2026, tripling from 2025 to 2030. [independent] (eetimes, citing analyst Vladimir Kozlov)
- LPO market (Dataintelo research report, 2026): $2.8B in 2025 → $14.7B by 2034, ~20.2% CAGR 2026–2034. (Market-research firm estimate — treat methodology as [secondary], unverifiable.) [secondary]
- Trade-off: lower module power with narrower operating margins; the engineering burden shifts to the host ASIC and board design; signal-integrity at 224G SerDes is the key scaling challenge; standards/interop for LPO are still evolving. [secondary] (eedesignit)
- Players exploring LPO: Broadcom, Marvell, Microsoft, Meta, Google. [secondary]
- Kozlov (eetimes): "The adoption of LPO and CPO in scale-up networks is expected to accelerate in 2026–2027 and reach high volumes by 2028."

### 5.3 800G / 1.6T optical transceiver market (2026)

