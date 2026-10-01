---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-12
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Meta", "Microsoft", "Nvidia", "xAI"]
dates: ["2023-07-19", "2025-06-11", "2025-09", "2026-06", "2026-07-16", "2026-08-25"]
keywords: ["amd", "asic", "benchmark", "blackwell", "cpo", "energy", "ethernet", "gpus", "hyperscaler", "inference", "intel", "latency"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [778, 836]
sha256: 1c962d2d12a17227758d5b8d599e81e24f3bcfa2c3b8b00bc303fac0ec1513d5
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

- Spectrum-X800 SN5600 switches are based on the Spectrum-4 switch ASIC (800G ports, 51.2T). [secondary] (tomshardware citing NVIDIA)

---

## 3. InfiniBand vs Ethernet in 2026 AI fabrics

### 3.1 2026 state of play

- Both NVIDIA Quantum-X800 (InfiniBand) and Spectrum-X800 (Ethernet) now ship at 800 Gb/s per port; Quantum-X800 at 115.2 Tb/s total switching capacity vs Spectrum-X800 51.2 Tb/s per the ascentoptics comparison (note: different switch configurations; treat as [secondary], vendor/product-page sourced). [secondary]
- 8-node reference numbers (secondary, 2026): IB NDR 400G (Quantum-2) ~350 GB/s effective all-reduce at sub-1 µs latency with SHARP in-switch reductions; well-tuned 400GbE RoCEv2 ~270–290 GB/s at ~2.4 µs; Spectrum-X within 5–10% of IB. [secondary] (rdp.in, temperature2.com)
- Decision guidance in 2026 (secondary consensus): IB for large training where all-reduce latency compounds every step; Ethernet (Spectrum-X/RoCE) for inference, multi-tenant clouds, and open-ecosystem deployments. [secondary]
- **Meta example:** two 24,576-H100 clusters run side by side — one RoCE-over-Ethernet (Arista 7800), one NVIDIA Quantum-2 InfiniBand — both reaching 90%+ network utilization training Llama 3. [secondary] (temperature2.com)
- Dell'Oro Q1 2026: Ethernet ~2/3 of AI-backend DC switch sales; IB sales tripled on the Blackwell Ultra 800G ramp (brownfield-upgrade caveat noted). [independent]

### 3.2 NVIDIA InfiniBand updates: Quantum-X800 / Quantum-3 / Quantum-X

- **Quantum-X800 (XDR InfiniBand, 800G/port):** shipping with NVIDIA's Blackwell Ultra platform ramp (Dell'Oro Q1 2026: IB sales "more than tripled ... supported by the ramp of 800 Gbps switches shipping with NVIDIA's Blackwell Ultra platform"). [independent]
- Quantum-X800 features SHARP v4 in-network computing for collective operations (9× improvement claim in the ascentoptics comparison table — [secondary], treat improvement factor as vendor-reported via that source). [secondary]
- NVIDIA newsroom framed the IB vs Ethernet choice for xAI: Colossus chose Spectrum-X Ethernet over InfiniBand — a landmark design decision for the world's largest AI supercomputer at the time. [official]
- **Quantum-X silicon photonics switches:** NVIDIA's announced Quantum-X and Spectrum-X silicon photonics networking switches "enable AI factories to connect millions of GPUs across sites while reducing energy consumption and operational costs" — referenced in the Spectrum-XGS release. [official]
- TrendForce roadmap: Quantum-2 (400G/port, 51.2 T) → Quantum-X800 (800G, 115.2 T) → Quantum-X1600 (1.6T, 230.4 T, paired with Rubin) → Quantum-X3200 (3.2T, paired with Feynman). [secondary]
- **Uncertainty flags:** (a) No Quantum-3 shipping announcement was found in the sources retrieved; the platform table lists Quantum-2 → Quantum-X800 → Quantum-X1600, so "Quantum-3" may be a skipped/renamed generation — do not assert its existence. (b) Quantum-X800 CPO/silicon-photonics variant shipping dates were described as "early 2026" roadmap (per the github silicon-photonics research note) but no dated production shipment was confirmed. (c) The `<100 ns` latency figure in the ascentoptics table vs "sub-1 µs" elsewhere is inconsistent — do not quote either as settled fact.

### 3.3 Broadcom vs NVIDIA "scale-out tech war"

- TrendForce (2025, still cited 2026): frames the AI fabric contest as Broadcom (Tomahawk/Thor/UEC) vs NVIDIA (Spectrum-X/Quantum-X) camps; Marvell Teralynx 10 and Cisco Silicon One G200 (both 51.2T, 2023) also compete; Cisco has CPO prototypes. [secondary]

---

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

