---
id: etape6-phasec-optics-cabling/00-front-matter/10-6-deployment-signals-named-operators
title: "10.6 Deployment signals (named operators)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "CoreWeave", "EU", "Google", "Huawei", "Intel", "Lambda", "Meta", "Microsoft", "Nvidia", "OpenAI", "Oracle", "xAI"]
dates: ["2024-10", "2024-10-16", "2024-10-28", "2025-06-04", "2025-10-08", "2025-10-16", "2026-03", "2026-09", "2026-09-16", "2026-09-22"]
keywords: ["asic", "cpo", "dsp", "ethernet", "full-duplex", "gpu", "gpus", "hbm", "intel", "latency", "liquid cooling", "lpo"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1559, 1621]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 83b4d971670402bc52263a6023e0c23297482e64d44dba7fd8622b9a82412b8f
---

# 10.6 Deployment signals (named operators)

### 10.6 Deployment signals (named operators)

- **xAI Colossus** (Memphis): revealed October 2024 to use NVIDIA Spectrum-X Ethernet rather than InfiniBand; each GPU paired with a 400GbE BlueField-3 SuperNIC and Spectrum SN5600 64-port 800G switches; NVIDIA claims 95% sustained data throughput with zero application latency degradation or packet loss from flow collisions across all three fabric tiers (vendor-claimed — treat the performance figure as such) [official] (globenewswire.com, 2024-10-28).
- Colossus was built with 100,000 Hopper GPUs (H100/H200); xAI/NVIDIA announced expansion toward 200,000 GPUs in a single building; ServeTheHome's facility tour corroborated the SN5600 + BlueField-3 pairing [independent] (sdxcentral.com; techradar.com).
- **Meta**: secondary technical writing states Meta built two parallel 24,576-GPU clusters — one InfiniBand, one RoCEv2 — and found the Ethernet cluster matched InfiniBand training throughput with topology-aware scheduling and tuning; strong but secondhand signal, not a Meta-published deployment spec [secondary] (foundation-model-engineering, GitHub).
- **Oracle**: at AI World (Las Vegas, October 16, 2025) Oracle unveiled **OCI Zettascale10** (up to 16 zettaFLOPS, targeting up to 800,000 NVIDIA GPUs per cluster) on the new **Acceleron RoCE** architecture with built-in GPU-NIC switching, multiple isolated network planes, and LPO/LRO optics sustaining 400G/800G throughput; flagship deployment is the **Stargate** supercluster with OpenAI in Abilene, Texas, with customer availability targeted H2 2026 [official] (oracle.com AI World roundup; corroborated by convergedigest.com).
- Oracle's Q1 FY2027 results (reported September 2026) show 850 MW of AI capacity with 300,000+ GPUs brought online and 97.9% GPU utilization, consistent with large-scale fabric buildouts, though the release does not break out switch models or port counts [secondary] (nationalcioreview.com).
- **CoreWeave**: NVIDIA stated in 2025 that CoreWeave would be among the first users of Spectrum-XGS to connect distributed data centers; on 2026-09-16 CoreWeave announced a multi-rack Vera Rubin NVL72 cluster connected with Spectrum-X Ethernet, with two ConnectX-9 SuperNICs per GPU giving 1.6 Tbps scale-out connectivity per GPU and a vendor-claimed modular nonblocking design supporting roughly 128,000 GPUs per rail [vendor-reported] (hpcwire.com). The underlying NVIDIA Spectrum-XGS announcement (2025) frames the distributed-data-center ("giga-scale AI super-factory") use case [official] (investor.nvidia.com, 2025).
- Arista's 2026 launch material references longstanding work with Meta, Microsoft, and Oracle, but no exact customer deployment model, port count, or date verified — vendor customer-logo signal only [vendor-reported].
- **No adequate evidence** found for named 800G-switch deployments at Google, Lambda, or exact Microsoft/Meta/Oracle installations (models, quantities, dates) — report as gaps, not negative findings [unverified].

### 10.7 Power, cooling, and optics power budgets

- System power (representative 51.2T air-cooled fixed): NVIDIA SN5600 ~940 W typical with passive copper; Cisco N9364E-SG2X 972 W typical / 2,779 W max; Cisco N9364E-SG2 official pages conflict between 2,270 W and 2,779 W maximum — flag the conflict; Edgecore AIS800-64O up to 2,775 W; Cisco N9364E-SP2R (HBM variant) 2,500 W typical / 5,100 W max [official] (cisco.com/static-cisco.com datasheets).
- 102.4T class: HPE Juniper QFX5250 3,375 W typical / 3,435 W max with direct liquid cooling and a fanless liquid-cooled variant at 10 lpm coolant flow; Alpha Networks' Tomahawk 6 liquid-cooled design claims removal of more than 1.8 kW of heat with 64 × 1.6T OSFP [official] (HPE QFX5250 document).
- 800G optic power classes (per-module, industry compilations): DSP-based pluggable 14–18 W; LPO 5–8 W (~40–50% lower, hot-swappable); LRO 9–12 W (~25% lower); CPO <3 W equivalent (board-level only, not field-serviceable) [secondary] (ascentoptics.com 800G power blog).
- FS's 800G DR8 OSFP LPO module specified at 8.5 W maximum, described as ~50% below DSP-based modules [vendor-reported] (thefastmode.com, FS 800G LPO launch).
- Adtran **LiteWave800** (March 2026 announcement, shown at OFC 2026): 800G DR8 LPO in OSFP claiming 1 pJ/bit ≈ 0.8 W using single-mode VCSELs and in-house low-power electronics — 12–18× below typical DSP-based optics per Adtran; vendor-claimed until independently tested [vendor-reported] (thefastmode.com).
- Broadcom CPO claim (via secondary compilation): an 800G CPO port consumes about 6.4 W versus ~16–18 W for regular pluggable modules/PCBs [secondary].

### 10.8 Pricing (public data is scarce)

- Nearly all enterprise/OEM and whitebox 51.2T/102.4T systems are quote-only; public pricing below excludes optics, cables, NICs, software/support, and volume discounts unless stated [independent].
- FS.com **N9600-64OD** (51.2T, 64 × 800G, Tomahawk 5): €33,894.00 VAT excl. (€40,333.86 incl.) EU / £30,136.00 VAT excl. (£36,163.20 incl.) UK — the only official street price found for a new 800G-port switch chassis [official] (FS.COM eu-en/uk product 250955).
- NVIDIA SN5600 (secondary compiled reseller snapshot, volatile single-unit pricing): Hardware Nation ~$67,213, an "Avendor" sale ~$70,391, claimed list $87,125, used/refurbished ~$18,750 [secondary] (gpusmith.com datasheet mirror).
- No trustworthy FS.com switch-chassis price beyond the N9600-64OD listing above was found, and no verified street prices were found for Edgecore, Micas, Dell Z9864F-ON, Arista, Cisco, or Juniper/HPE 800G systems — report the gap honestly [independent].

### 10.9 Ecosystem timeline (anchor dates)

- **2022-08**: Broadcom ships Tomahawk 5 (51.2T) [official].
- **2023-01**: Intel exits future network-switching investment (Tofino line wound down to support-only) [independent].
- **2023-03**: Marvell previews Teralynx 10 (51.2T), sampling guided for Q2 [independent].
- **2023-04**: Broadcom Jericho3-AI available to qualified customers [independent].
- **2023-10**: Edgecore announces AIS800-64O/64D 800G whiteboxes for OCP/Fyuz demos [independent].
- **2024-10-16**: Micas M2-W6940-64OC announced as globally shipping [official].
- **2024-10-28**: NVIDIA/xAI disclose Colossus Spectrum-X Ethernet fabric [official].
- **2025-06-04**: Broadcom Tomahawk 6 (102.4T) shipping reported [vendor-reported].
- **2025-10**: Arista 7800R4/7280R4 800G announcement (shipping now per PR); Broadcom Tomahawk 6 Davisson CPO announced 2025-10-08 [official].
- **2025-10-16**: Oracle unveils OCI Zettascale10 + Acceleron RoCE (400G/800G, H2 2026 customer target) [official].
- **2026-02**: Cisco unveils Silicon One G300 (102.4T) and N9364F-SG3 [vendor-reported].
- **2026-03**: Adtran announces LiteWave800 0.8W 800G LPO (OFC 2026) [vendor-reported].
- **2026-09-16**: CoreWeave brings up multi-rack Vera Rubin NVL72 on Spectrum-X Ethernet [vendor-reported].
- **2026-Q4 → 2027-Q1** (planned, future relative to research date): Arista 7060XE7-64PS (Q4 2026), 7060XE7-64PRS/128PE (early 2027), 7060XE7-64PRS-RV3-L (Q1 2027) — announced, not confirmed shipping [independent].

### 10.10 Gaps, rumors, and unverified claims (do not use without confirmation)

- **SN5800**: no evidence of a real announced NVIDIA system by 2026-09-22 [unverified].
- **Juniper Express5 / Trio 800G data-center ASIC**: no fresh primary source captured; Juniper's documented 800G systems use merchant silicon (Tomahawk 5/6, Trident 5) [unverified].
- **Arista "7388X5"** as an 800G fixed system: not found [unverified].
- **Celestica 800G model** (e.g., DS-series 800G SKU with ASIC mapping): not found [unverified].
- **ZTE native 800G data-center switch**: not found [unverified].
- **Ruijie RG-S6990-64OC2XS**: reseller-only, no official datasheet; claimed 102.4T figure is full-duplex-implied and non-comparable [unverified].
- **Huawei 16800-X 288 × 800GE**: secondary claim only; official datasheet dual capacity figures (179/387 etc.) need accounting clarification; exact 800G line card unverified [unverified].
- **Cisco N9364F-SG3** details (cages, buffer, power, form factor, ship date): announced Feb 2026, awaiting official datasheet [unverified].
- **Edgecore per-cage optic budget**: 24 W vs 30 W across sources — unresolved [unverified].
- **Cisco N9364E-SG2 max power**: 2,270 W vs 2,779 W across official pages — unresolved [unverified].
- **Q3200-RA aggregate bandwidth**: 57.6 Tbps vs 2 × 28.8 Tbps across mirrored datasheets — inconsistent aggregation [unverified].
- **Tomahawk 6 production ramp scale** through 2026 (volumes, lead customers beyond announced OEM systems): not quantified in captured sources [unverified].
- **Arista 7800R4 line-card ASIC** (Jericho3-AI vs merchant deep-buffer silicon): datasheet excerpt lists SKUs but not the ASIC — do not assert a mapping [unverified].
- **Dell SN6800-LD ASIC identity**: sheet states 102.4 Tb/s modules without naming silicon — do not label "Spectrum-X successor" as fact [unverified].
- **Adtran LiteWave800 0.8W** and **Micas 30% CPO power saving**: vendor claims, not independently measured [unverified].
- **Meta/Microsoft/Google/Lambda/xAI (beyond Colossus) named 800G deployments** with models and quantities: not verified [unverified].

