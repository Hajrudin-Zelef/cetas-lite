---
id: etape6-trackb-arista-sonic/00-arista-sonic/4-open-networking-landscape-2026
title: "4. OPEN NETWORKING LANDSCAPE 2026"
domain: step-6-track-b-arista-sonic-cumulus-data-center-fabric-open-
role: deep-dive
task: reference
actors: ["AMD", "Anthropic", "Broadcom", "Huawei", "Intel", "Meta", "Microsoft", "Nvidia", "OpenAI", "TSMC"]
dates: ["2023-07-19", "2025-06-11", "2026-03-12", "2026-07-16"]
keywords: ["3nm", "alignment", "amd", "asic", "cpo", "ethernet", "gpu", "intel", "latency", "lpo", "nvidia", "optics"]
source: docs/RAG/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [163, 235]
section: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
sha256: 4c58fae4ce82956385f1f776f45d70d34ccc389fc03f7d1d35142362a7e7a456
---

# 4. OPEN NETWORKING LANDSCAPE 2026

## 4. OPEN NETWORKING LANDSCAPE 2026

### 4.1 Ultra Ethernet Consortium (UEC) / Super Ethernet

**Specification status**
- UEC is a **Linux Foundation Joint Development Foundation** project founded **July 19, 2023** by AMD, Arista, Broadcom, Cisco, Eviden/Atos, HPE, Intel, Meta, Microsoft **[secondary — LLM-Systems-Wiki UEC page, updated Aug 26, 2026]**.
- **Specification v1.0**: 562-page document released **June 11, 2025**, defining a full Ethernet-based communication stack (NICs, switches, optics, cables) with the new **Ultra Ethernet Transport (UET)** **[secondary]**.
- Current release: **1.0.3 (July 16, 2026)**; 1.0.2 also released in 2026 **[secondary]**.
- Consortium scale: **100+ member companies, 1,500+ participants** **[secondary — Medium/Teradata Labs, Jun 2026]**.
- UET design goals: per-packet multipathing, lossy-or-lossless operation, fast loss recovery, libfabric compatibility; removes in-order-delivery requirement and lossless-fabric dependency of RoCEv2 **[secondary]**.

**2026 technical priorities** (per Chad Hintz, UEC marketing co-chair / AMD, via Network World):
1. **Programmable Congestion Management (PCM)** — implement congestion-control algorithms in a standard language, portable across UE NICs.
2. **Congestion Signaling (CSIG)** — packets carry high-fidelity congestion information for faster, more accurate transport reactions.
3. **Small-message performance** — UET 1.0's 104-byte headers are fine for 4KB packets (2.5% overhead) but costly for 256-byte transactions; UEC pursuing a reduced-size forwarding header to cut overhead roughly in half for optimized deployments (matters for HPC and local scale-up networks).
4. **In-Network Collectives (INC)** — move AI/HPC reduction operations from hosts into the network **[independent — Network World, ~Jan 2026]**.

**Vendor alignment**
- **Broadcom**: "first UE-compatible Tomahawk" switch with ~250ns latency (claimed); UEC-aligned features (Cognitive Routing 2.0, Global Load Balancing) in Tomahawk 6 **[secondary — Medium; financialcontent]**.
- **Arista**: founding/steering UEC member; Etherlink roadmap tracks UET (per-packet spraying, NSCC/RCCC, trimming, in-network collectives) **[secondary — LLM-Systems-Wiki]**.
- **NVIDIA**: UEC member (~2024) but pushes its proprietary **Spectrum-X** stack; Spectrum-X is the bundled alternative to UEC Ethernet **[secondary]**.
- Ethernet Alliance 2026 focus: higher bandwidth, AI demands, interoperability events, another 200G/lane plugfest **[independent — Network World]**.

### 4.2 Tomahawk 6 — shipments and deployments

- **Broadcom began volume shipments** of the Tomahawk 6 switch series on **March 12, 2026** — first 102.4 Tbps single-chip Ethernet switch; **TSMC 3nm**; 512×200G SerDes channels; ~250ns latency (UE-compatible variant); Cognitive Routing 2.0, Global Load Balancing to attack incast congestion **[secondary — financialcontent/marketminute, Mar 26, 2026; TechPowerUp]**.
- Positioning: unified scale-up and scale-out AI networks; support for 100G/200G SerDes and co-packaged optics (CPO); designed for clusters of **more than one million XPUs**; open specifications for easy integration **[secondary — TechPowerUp]**.
- **Arista and Cisco integrated TH6** into their latest chassis — reporting "flatter" topologies: large AI clusters built with only **two tiers** instead of three or four, cutting optical transceiver count and GPU-to-GPU latency **[secondary — financialcontent]**.
- **Initial deployments scheduled for clusters of 100,000+ XPUs** **[secondary — StorageReview]**.
- **DriveNets** built AI Fabric switches around the Tomahawk 6 ASIC, targeting AI infrastructures with hundreds of thousands of XPUs; systems expected to ship **Q3 2026** — a visible system-level demand signal for TH6 **[secondary — ainvest, Jul 2026]**.
- Broadcom Q1 FY2026 (Mar 2026): total revenue $19.31B (+29.5%); AI infrastructure revenue **$8.4B (+106%)**; networking segment revenue +60% YoY, with Tomahawk 6 "effectively secured Broadcom's technological lead in data center fabrics"; confirmed **OpenAI as sixth custom-silicon (XPU) customer**; Anthropic $21B "Ironwood Racks" order **[secondary — financialcontent, Mar 6, 2026]**.
- ⚠️ "Beat competitors to the 100T era by at least two full quarters" is a secondary outlet's characterization, not independently verified.

### 4.3 White-box / ODM switching

- IDC no longer breaks out ODM Ethernet revenue publicly; **The Next Platform estimates ODMs at $3.56B in Q2 2026, up 2.43× YoY** — its own estimate, **[unverified as to exact figure; directionally consistent with AI buildout]** **[secondary — The Next Platform, Sep 20, 2026]**.
- Open-hardware ecosystem for SONiC deployments: Accton/Edgecore, Celestica, Wistron, Cisco, NVIDIA (Aviz-certified images) **[secondary — Aviz]**.
- SONiC enables NOS choice on white-box fleets; Dell has promoted this model for enterprise **[independent — The Register]**.

---

## 5. ARISTA vs NVIDIA SPECTRUM — COMPETITIVE POSITIONING

### 5.1 Market-share data (IDC, via press)

**Q1 2026 — NVIDIA takes the data-center Ethernet lead**
- **NVIDIA became the #1 data-center Ethernet switch vendor**: ~$2.1B revenue, **+193% YoY**, **21.5% share** — ahead of Arista (20.7%) and Cisco (17.8%) **[secondary — ainvest, citing IDC; remio.ai]**.
- NVIDIA's share climbed from <4% two years earlier to 21.5% in one quarter; the vehicle is **Spectrum-X** **[secondary]**.
- Overall Ethernet switching market: **$15.4B in Q1 2026, +39.8% YoY** **[secondary — remio.ai]**.
- NVIDIA's data-center Ethernet share was ~11% in late 2025 / 11.6% in Q3 2025 **[secondary]**.

**Q2 2026 — the gap widens (IDC via The Next Platform, Sep 20, 2026)**
- Total Ethernet switch revenue grew strongly; **NVIDIA's Ethernet revenue hit $3.86B (2.8×)**, driven by the GenAI systems business and **Spectrum-X attach rates** on DGX nodes and NVL72 rack-scale systems **[secondary — The Next Platform]**.
- Arista: **$2.53B, +37.6%**; Cisco: $5.43B (+36.2% overall, data-center +43% in Q1); HPE (incl. Juniper): $1.15B (+6.1%); Huawei: $1.55B (+28.6%); ODMs est. $3.56B (2.43×) **[secondary — The Next Platform; Q1 HPE/Huawei figures from IDC via ainvest]**.
- Data-center Ethernet in Q2: **$12.3B, +64.5% YoY**, +23% sequentially; ~129.5M ports sold into the data center (analyst estimate); speeds above 200G are overwhelmingly data-center/AI **[secondary — The Next Platform]**.

### 5.2 Architecture comparison: Etherlink vs Spectrum-X

| Dimension | Arista Etherlink | NVIDIA Spectrum-X |
|---|---|---|
| Switch ASIC | Merchant (Tomahawk 5 / Tomahawk 6 / Jericho 3-AI) | In-house Spectrum-4 (+ Spectrum-6 on roadmap) |
| Endpoint / NIC | Third-party NICs (no own NIC) | BlueField-3 SuperNIC / ConnectX; DPUs bundled |
| Fabric stack | RoCEv2 + PFC + DCQCN today; UEC/NSCC/RCCC roadmap | Integrated platform: Spectrum switches + BlueField DPUs + LinkX cables + CUDA stack |
| Congestion control | DCQCN now; UEC NSCC/RCCC later | Closed-loop TCC via SuperNIC |
| Multipath | DLB, packet spraying, MRC/CSIG (EOS) | Per-packet spraying + TCC |
| UEC standing | **Founding/steering member** | Member (~2024), pushes Spectrum-X |
| Optics | LPO/XPO, multi-vendor | Integrated; Spectrum Photonic (TSMC COUPE silicon photonics) roadmap |
| Best-fit | Open multi-vendor AI fabrics; merchant-silicon + UEC trajectory; no NVIDIA networking lock-in | NVIDIA-GPU AI factories; pre-optimized full stack |

**Comparison sources:** LLM-Systems-Wiki Arista/Etherlink page (secondary, vendor-adjacent technical detail) and StorageReview TH6/Spectrum-Photonic comparison (secondary). Treat table cells as secondary-sourced where marked.

### 5.3 Analyst framing of the competitive dynamic

