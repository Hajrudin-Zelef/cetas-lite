---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-13
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom", "China", "Cohere", "Google", "Intel", "Meta", "Microsoft", "Nvidia", "Oracle", "TSMC"]
dates: ["2026-03", "2026-05", "2026-06", "2026-06-09", "2026-07", "2026-08-04"]
keywords: ["asic", "chiplet", "coherent optics", "cost", "cpo", "datacenter", "dci", "dsp", "energy", "ethernet", "foundry", "gpus"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [837, 887]
sha256: 22bc1973d91046a4bf3e96f67ee981caf73d6d65220b0ea71c88ff3d4500601f
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

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

## 6. Arista: 2026 AI networking revenue, wins, earnings

### 6.1 Q2 2026 results (reported August 4, 2026)

- Revenue **$3.036B** (+12.1% QoQ, **+37.7% YoY**) — first $3B quarter. GAAP operating margin 45.4%, non-GAAP 49.9%. Non-GAAP diluted EPS $1.02 (+39.7% YoY); GAAP EPS $0.95. [official] (Arista Q2 2026 release; also hosted at s21.q4cdn.com)
- CEO Jayshree Ullal: "As we deliver our first $3 billion quarter in Q2 2026, it is clear that our Arista 2.0 platform strategy is compelling. Customers see networking as the central nervous system for infrastructure from the client to campus to data and AI centers." [official]
- CFO Chantelle Breithaupt: "Our second quarter 2026 reflects strong, broad-based growth, with revenue up 37.7% and EPS up 39.7%." [official]
- Recognized on the 2026 Fortune 500 list; named a Leader in the 2026 Gartner Magic Quadrant for Enterprise Wired and Wireless LAN. [official]
- Q1 2026 (reported ~May 2026): $2.71B revenue (+35.1% YoY), diluted EPS $0.87 (+31.8%), raised 2026 guidance to $11.5B with $3.5B from AI. [secondary] (fxempire)

### 6.2 Guidance and AI-fabrics revenue

