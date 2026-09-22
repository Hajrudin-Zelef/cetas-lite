---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/4-uec-ultra-ethernet-consortium-super-ethernet
title: "4. UEC (Ultra Ethernet Consortium) / Super Ethernet"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "China", "Cohere", "Google", "Intel", "Meta", "Microsoft", "Nvidia", "Oracle", "TSMC"]
dates: ["2023-07-19", "2025-06-11", "2025-09", "2026-03", "2026-06", "2026-06-09", "2026-07", "2026-07-16", "2026-08-25"]
keywords: ["ethernet", "amd", "asic", "benchmark", "chiplet", "coherent optics", "cost", "cpo", "datacenter", "dci", "dsp", "energy"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [807, 875]
section: "Data-Center Networking for AI (2026)"
sha256: 259dffd6a48daac90a6594a4187df89aa3cebd06df53791cc3c244d041f9aebd
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

- TrendForce: **800G+ shipments ~24M units (2025) → ~63M units (2026)**, 2.6× growth; 1.6T shipments **~2.5M (2025) → 20M+ by end-2026** (per marketminute coverage of OFC 2026 analyst confirmation; 1.6T "scaling faster than any previous generation"). [secondary]
- LightCounting-derived TAM (cited in a July 2026 investor research note): total transceivers $23.8B (2025); datacenter modules ~$22.8B (2026), of which **800G+1.6T ~$14.6B**; 800G+ units ~24M (2025) → ~63M (2026). Note: scope disagreements are wide (MarketsandMarkets pegs DC transceivers at only $9.2B for 2025) — treat single points cautiously. [secondary] (github/roachx92 citing LightCounting)
- GM Insights (market-research firm, 2026): optical transceiver market $13.4B (2025) → $15.4B (2026) → $48.1B (2035), 13.5% CAGR; Coherent led with >22.2% share in 2025; top-5 (Coherent, Cisco, Broadcom, Lumentum, Accelink) held 72.2%. (Estimate — [secondary].) [secondary]
- Vendor landscape: **Innolight (~$5.6B datacenter modules)** and Eoptolink together hold **~60% of NVIDIA's 800G volume**; Coherent and Lumentum are the Western integrated laser suppliers (Lumentum the only volume 200G-EML shipper at ~50–60% share, +40% EML capacity in 2025 and again 2026); Coherent expanding 6-inch InP across 4 sites (~4× die/wafer at ~½ cost, capacity doubling 2026); Cisco (Acacia) took >$1B of optics orders in a single quarter (Q4 FY2026). [secondary] [vendor-reported] (mix: github research note citing J.P. Morgan/LightCounting; Cisco via ainvest)
- **EML bottleneck:** TrendForce reports an upstream laser bottleneck; NVIDIA pre-allocated large EML capacity (investor note claims "$4B lock-up of Lumentum+Coherent EML capacity" with multiyear purchase commitments — [unverified]); NVIDIA's March 2026 $2B direct investment in Lumentum to guarantee photonics capacity (marketminute). Extended EML lead times beyond 2027 reported. [secondary]
- Adopters: "1.6T is shipping now," with **NVIDIA and Google already integrating**; Meta and Oracle slated to follow (per Kozlov via eetimes). [independent]

### 5.4 NVIDIA silicon photonics

- NVIDIA's roadmap includes **Spectrum-X Ethernet Photonics (H2 2026)** and Quantum-X InfiniBand (early 2026) — i.e., CPO/silicon-photonics switch variants (per the github silicon-photonics research note). [unverified] — no dated NVIDIA product announcement of these variants was located in vendor sources; treat H2 2026 as roadmap guidance, not confirmed shipping.
- NVIDIA announced **Quantum-X and Spectrum-X silicon photonics networking switches** (referenced in the Spectrum-XGS release as enabling "AI factories to connect millions of GPUs across sites while reducing energy consumption and operational costs"). [official]
- NVIDIA CTO Gilad Shainer: CPO "already shipping and ramping H2 2026" (response to June 2026 SemiAnalysis yield-risk report). [vendor-reported]
- NVIDIA's slower-than-anticipated silicon-photonics/CPO progress is one reason for continued dependence on pluggable modules and the EML pre-allocation strategy (TrendForce). [secondary]

### 5.5 Data center interconnect (DCI) notes

- 800G digital coherent optics (DCO): $963M market (2025) → $1,713M (2032), 8.7% CAGR (MarketPublishers, Jan 2026 — [secondary] estimate); key players II-VI/Coherent, Lumentum, Innolight, Hisense, Accelink. [secondary]
- Scale-across DCI is emerging as its own category: Cisco's Silicon One **P200** is explicitly positioned for scale-across deployments (three hyperscaler wins in Q4 FY2026 incl. a P200 scale-across win); NVIDIA Spectrum-XGS is the Ethernet answer; Arista's 1.6T platforms target scale-across. [vendor-reported] [official]
- Corning announced a multi-billion-dollar "rack-scale" fiber infrastructure deal with Meta for gigascale AI data centers (March 2026). [secondary] (marketminute)

---

