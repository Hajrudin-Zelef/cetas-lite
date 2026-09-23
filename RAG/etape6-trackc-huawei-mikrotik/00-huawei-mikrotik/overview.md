---
id: etape6-trackc-huawei-mikrotik/00-huawei-mikrotik/overview
title: "Step 6 — Track C: Huawei + MikroTik (Networking Hardware)"
domain: step-6-track-c-huawei-mikrotik-networking-hardware
role: deep-dive
task: hardware
actors: ["Broadcom", "China", "DeepSeek", "EU", "Huawei", "Nvidia", "United States"]
dates: ["2025-04", "2026-01-01", "2026-03-17", "2026-09-22"]
keywords: ["accelerator", "agent", "agentic", "ascend", "compute", "cost", "datacenter", "deepseek", "dsp", "energy", "ethernet", "gpu"]
source: docs/RAG/etape6_trackC_huawei_mikrotik.md
source_anchor: ""
source_lines: [1, 71]
section: "Step 6 — Track C: Huawei + MikroTik (Networking Hardware)"
sha256: f0763d3e7641659f07a7e60c50663fbea246241e60f9ba4de510735f70e247b8
---

# Step 6 — Track C: Huawei + MikroTik (Networking Hardware)
## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 6 (Réseau & sécurité), Track C
**Coverage window:** January 1, 2026 → September 22, 2026
**Scope:** Huawei networking portfolio (data-center switches, AI fabric, international status), MikroTik 2026 launches (routers/switches, RouterOS v7), switch pricing trends (25G/100G/400G), open-source NOS news (brief)

### Provenance legend
- **[official]** — vendor's own site, press release, or documentation.
- **[vendor-reported]** — figure claimed by the vendor without independent audit.
- **[independent]** — reputable third-party press (IDC, Dell'Oro, The Register, SemiAnalysis, Telegeography) or independent measurement.
- **[secondary]** — lower-tier press, marketplace listings, blogs, community coverage; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

---

## 1. HUAWEI — 2026 networking portfolio

### 1.1 HUAWEI CONNECT 2026: Stellar AI Fabric upgrade (Sep 21, 2026)

At HUAWEI CONNECT 2026 (Shanghai), Huawei hosted the **AI DC Innovation Summit** ("Leading AI DC Innovation, Shaping the Agentic World", 600+ industry leaders and researchers) and unveiled an upgraded **Stellar AI Fabric Solution** for production AI clusters **[official — https://www.prnewswire.com/in/news-releases/huawei-connect-2026--huawei-upgrades-stellar-ai-fabric-solution-to-build-efficient-ai-computing-production-networks-302884451.html]**.

Software/management upgrades:
- **Rock-Solid Architecture 2.0**, **StarryWing Digital Map 2.0** (network digital twin), **Hyper-Converged Fabric (HCF)**, **Network Packet Load Balancing (NPLB)**.
- **NetMaster network AI agent**: automates root-cause analysis for **95% of network faults within minutes** (vendor-reported) to maintain production uptime. Keynote by Erick Zhang, President of Data Center Network Domain at Huawei Data Communication Product Line: "AI is turning traditional data centers into compute factories… the underlying network fabric directly dictates token throughput and overall compute yield" **[official]**.

Two new switch families launched for high-density computing:
- **CloudEngine SF9300 (UBG switch)** — two-layer multi-plane architecture that cuts buildout costs by **30%**; simplified UB forwarding lowers end-to-end latency by **40%**; **Link-Layer Retransmission (LLR)** executes microsecond-level recovery, eliminating packet loss during link flapping **[official, vendor-reported]**.
- **CloudEngine XH9300 (NPO switch)** — "Near-Packaged Optics" switch delivering proprietary **100T and 51.2T** (switching capacity figures); centralized light source with proprietary **3.2T Optical Engine (OE)** lowers interconnect power draw by **40%**; eliminating the optical DSP (oDSP) stage cuts forwarding latency by **26%**; snap-in pluggable optical engines speed up maintenance **10×** **[official, vendor-reported]**.

Standards and research deployments:
- The **Beijing FinTech Industry Alliance, Huawei, and 17 partners** released the *Technical Specification for Network High Availability of Financial Data Center*, defining technical baselines and testing methods across two categories and five network types for intra-DC and cross-DC environments (architectural redundancy, fault recovery, proactive prevention) **[official]**.
- **IHEP** (Institute of High Energy Physics, Chinese Academy of Sciences): deploying the Stellar AI Fabric lifted computing efficiency by **over 10%**; network-wide load balancing resolved training traffic imbalances for deep-space signal analysis workloads (per Computing Center Director Qi Fazhi) **[official, vendor-reported]**.

### 1.2 AI data-center networking: CloudMatrix 384

Huawei's AI cluster networking centers on the **CloudMatrix 384** supernode: 384× Ascend 910C + 192× Kunpeng CPUs, connected via Huawei's self-developed **Unified Bus (UB)** in a fully non-blocking all-to-all interconnect — a deliberate alternative to copper/NVLink-style scale-up **[independent — https://semianalysis.com/2025/04/16/huawei-ai-cloudmatrix-384-chinas-answer-to-nvidia-gb200-nvl72/; secondary — https://chinaresearchcollective.substack.com/p/huawei-ascend-cloudmatrix-384-supernode]**.

- **~289–300 PFLOPS (BF16)**, **~48 TB HBM** total, 16 racks; UB bandwidth **196 GB/s per NPU** (vs PCIe 5.0's 64 GB/s); 7-layer sub-plane design, each layer an independent full-mesh **[secondary]**.
- Optical interconnect: **6,912× 400G SiPh LPO (Linear Pluggable Optics) modules** per pod — 5,376 for the vertical (scale-up) expansion network, 1,536 for the horizontal (RDMA scale-out) network; 3,168 optical fibers; each NPU gets 2.8 Tbps interconnect bandwidth (7× 400G transceivers per chip); flat single-tier topology with ~16,800 modular switch elements using cell-spraying (mechanism similar to Broadcom Jericho3 + Ramon3) **[independent/secondary — SemiAnalysis via fibermall.com; https://www.prnewswire.com — Huawei Connect 2025 materials]**.
- Huawei positions LPO as a power-saving measure vs DSP-based optics; tradeoff noted by analysts: higher transceiver count raises cost, power, airflow, and reliability concerns, mitigated by fault-tolerant training software **[independent — SemiAnalysis]**.
- First deployment context: Inner Mongolia Ulanqab Data Center (April 2025 launch); 2026 role: domestic compute foundation for models such as DeepSeek V4 Pro **[secondary]**.

### 1.3 Enterprise Wi-Fi 7 leadership (Mar 2026)

- **March 17, 2026**: Huawei announced it ranked **No. 1 both globally and in China** in enterprise-class **Wi-Fi 7 market share and shipments** per IDC's *Worldwide Quarterly Wireless LAN Tracker (2025 Q4)* — the **second consecutive year** of leadership since IDC began publishing Wi-Fi 7 statistics in 2024 **[official — https://e.huawei.com/en/news/2026/solutions/enterprise-network/wi-fi7-market-share-shipments; independent — IDC]**.
- **2M+ enterprise-class Wi-Fi 7 products shipped**; Huawei is also No. 1 globally in Wi-Fi 7 standards contributions **[official, vendor-reported]**.
- Background: world's first enterprise-class Wi-Fi 7 AP (2022); first commercial Wi-Fi 7 network deployment (2023, Tolly-tested); industry's first complete Wi-Fi 7 portfolio for all scenarios (2024); world's first enterprise-class Wi-Fi CERTIFIED 7 certificate **[official]**.

### 1.4 Market position: Q2 2026 results (IDC, Dell'Oro)

**IDC — global Ethernet switch market, Q2 2026** (reported ~Sep 21, 2026) **[independent — https://www.idc.com/resource-center/blog/ethernet-switch-market-surges-43-4-to-18-9b-in-2q26-as-ai-infrastructure-demand-drives-record-datacenter-spending/; https://www.kad8.com/network/global-ethernet-switch-market-hits-18.9-billion-usd-in-q2-2026/]**:
- Worldwide Ethernet switch revenue hit **$18.9B, +43.4% YoY** — AI infrastructure demand driving record datacenter spending.
- **Huawei: ~$1.5B Ethernet switch revenue, +28.6% YoY, 8.2% global market share.**
- **Huawei router revenue: $1.3B, +21.6% YoY, 29.9% global router market share** — "underscoring continued strength in service provider networking, particularly in China and select emerging markets."
- IDC's three signals: (1) AI cluster construction is a top driver of datacenter network spend; (2) east-west GPU/server traffic makes bandwidth/latency directly affect accelerator utilization; (3) the fabric is becoming part of the AI compute system.

**Dell'Oro — High End Routing and Aggregation, Q2 2026** (reported ~Sep 10, 2026) **[independent — https://www.prnewswire.com/news-releases/high-end-routing-and-aggregation-market-grew-25-percent-in-2q-2026-according-to-delloro-group-302872936.html]**:
- Segment revenue **+25% YoY**; vendor **direct sales to cloud providers +94% YoY** (hyperscaler AI infrastructure push).
- Trailing-4-quarter vendor rank: **High End Routing & Aggregation: 1 Cisco, 2 Huawei, 3 Nokia**; **Cloud Provider segment: 1 Cisco, 2 HPE Juniper, 3 Nokia**; **Communication Service Provider segment: 1 Huawei, 2 Cisco, 3 Nokia**.

**Fortune Business Insights (enterprise networking)**: Huawei listed among top vendors with ~18% enterprise networking market share (vs Cisco 22%) **[secondary — https://www.fortunebusinessinsights.com/enterprise-networking-market-105887]**.

### 1.5 International market status

- Huawei's networking growth remains anchored in **China and select emerging markets** (Middle East, Africa, Latin America, Southeast Asia) — per IDC's characterization of its router business **[independent]**.
- Export-control environment: Huawei continues to operate its networking business under US entity-list restrictions (imposed 2019); no lifting or change in 2026 found in sources fetched — treat current status as **[unverified]** for any specific 2026 regulatory change.
- No 2026 Huawei telecom/datacom wins in US, EU-5, Japan, Australia, or UK markets were found in fetched sources — consistent with the ongoing restrictions **[secondary]**.

---

