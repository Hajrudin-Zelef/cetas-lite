---
id: collect-261001-general-networking/general-networking/etape6-trackb-arista-sonic-2
title: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Broadcom", "Meta", "Microsoft", "Nvidia", "Oracle"]
dates: ["2025-05", "2025-12", "2026-04-15", "2026-05-27", "2026-06", "2026-06-09", "2026-10"]
keywords: ["accelerator", "agentic", "amd", "aws", "compute", "ethernet", "hyperscaler", "inference", "latency", "lpo", "maia", "nvidia"]
source: docs/RAG/collect-261001-general-networking/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [58, 116]
sha256: 4492f94689651182af90f7b94e6e58e70d34e70c6572bf1dd44e3ef0950e6a48
---

# Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)

**7060XE7 Series — 1.6T launch (June 9, 2026)**
- Announcement: "Arista Introduces Next-Generation 1.6Terabit Portfolio for AI Fabrics" — new portfolio of **1.6T networking platforms for rack-scale AI infrastructure** **[official — Arista press release, June 9, 2026]**.
- Built on **Broadcom Tomahawk 6** silicon; up to **100 Tbps** system switching capacity; 1.6 Tbps per port **[official]**.
- Models (per press release and press coverage):
  - **7060XE7-64PS** and **7060XE7-64PRS**: air-cooled 4U rack switches supporting pluggable IHS (integrated heat sink) and RHS (riding heat sink) optics; IHS for current air-cooled data centers, RHS for future liquid-cooled fabrics **[official]**.
  - **7060XE7-64PRS-RV3-L**: specialized 2OU liquid-cooled platform for high-density clusters, **224G SerDes**, DC power from ORv3 rack, no internal fans, integrates with liquid-cooled XPU servers; availability **Q1 2027** **[official / secondary — Network World]**.
  - **7060XE7-128PE**: **128 ports of 800G** in an air-cooled 4RU design (100G SerDes), for environments needing deployment flexibility and backward compatibility **[official]**.
  - Air-cooled rack switches available **Q4 2026** **[secondary — Network World]**.
- **Linear Pluggable Optics (LPO)**: headline claim of ~**60% reduction in interconnect power consumption** vs traditional optics **[official]**.
- **XPO high-density liquid-cooled pluggable optics**: claimed to reduce networking racks by up to 75% and save up to 44% of floor space vs traditional optics **[secondary — ainvest, citing launch materials]**.
- Runs **Arista EOS** (low-latency, intelligent packet buffering for AI microbursts and collective patterns) and also supports **open network OS options** (SONIC, OpenSwitch) **[official / secondary — Network World, Zacks]**.
- **Hyperscaler endorsements in the launch**:
  - Meta, Microsoft and Oracle contributed statements to the product announcement **[official — Arista press release]**.
  - Microsoft's **Rani Borkar** noted collaboration on the 1.6T Ethernet interface for **Azure Maia**, Microsoft's AI accelerator chip **[secondary — CoinCentral]**.
  - Meta endorsement quote (from press release): "Arista Networks' 1.6T platforms provide the throughput, determinism, and stability needed for our RDMA-based AI fabrics, while Arista EOS delivers operational consistency and performance at scale across our global AI infrastructure." **[official — Arista press release]**.
  - Platform designed to work with **AMD's compute silicon and network interface cards** **[official]**.
- Arista's strategic framing: transition "from providing high-performance switches to delivering comprehensive rack-scale systems" — the network as "an elastic and integrated backplane" / "tightly integrated AI supersystem" rather than a standalone layer **[official — Tyson Lamoreaux, SVP Cloud and AI Networking, Arista]**.

### 1.3 EOS, CloudVision, NetDL, AVA — 2026 software updates

- **EOS (Extensible Operating System)**: single-OS architecture running across Arista's entire hardware line from campus edge to AI spine — longstanding differentiator vs multi-OS vendors **[secondary — The CODEW]**.
- **CloudVision + NetDL**: network-wide management and telemetry layer feeding a **Network Data Lake (NetDL)** that ingests telemetry from Arista hardware, third-party systems, server NICs, and AI job schedulers — explicitly built to correlate network behavior with AI training-job performance **[secondary — The CODEW]**.
- **AVA (AI-driven network automation)**: built on EOS with NetDL as unified data repository; 2026 enhancements announced **December 2025** (GA anticipated Q1 2026):
  - Agentic AI framework: multi-domain event correlation, agentic conversational/troubleshooting with context-aware dialogue, continuous monitoring and automated root cause analysis across wired, wireless, data center and security domains **[secondary — ITDigest, Dec 11, 2025]**.
  - Announced as part of "AI-Powered Campus Mobility at Scale" launch alongside new campus products **[secondary]**.
- **Campus / enterprise (Dec 2025 announcements, GA expected Q1 2026)**: two ruggedized switches (**710HXP-28TXH** and **710HXP-20TNH**) for industrial/outdoor environments (extreme temperatures, vibration, shock; IP50 Din-Rail 20-port and IP30 1RU 24-port with 90W PoE); Wi-Fi 7 access points; all run EOS and integrate with CloudVision **[secondary — ITDigest]**.
- One secondary report targets **$1.25B in campus revenue in 2026**, taking share from Cisco in corporate/branch **[secondary — Financial Freedom is a Journey, Feb 2026 — unverified]**.

### 1.4 Customer concentration and risks (analyst framing)

- Top three hyperscalers account for roughly **55% of net revenue** (alternative report: ~50% from Meta/Microsoft/Amazon/AWS/Alphabet) — concentration is the top cited risk **[secondary — research-ANET GitHub brief, Jul 2026; FFJ]**.
- Margin pressure: as AI clusters grow, buyer leverage forces volume discounts; management guided 2026 operating margins down slightly to ~46% from historic 48%+ highs **[secondary — FFJ]**.
- Supply-chain risk: 1.6T requires cutting-edge optics and cooling; optics hiccups could delay 2026 deployments **[secondary — Tamar Securities]**.
- Cisco competitive resurgence: Cisco data-center segment +43% YoY in Q1 2026 (IDC); campus/enterprise refresh is a grind **[secondary — ainvest]**.
- Customer-concentration example of volatility: a spending pause by a single Cloud Titan (e.g., Microsoft) can create revenue volatility **[secondary — Tamar Securities]**.

---

## 2. SONiC (SOFTWARE FOR OPEN NETWORKING IN THE CLOUD)

### 2.1 Community status and Microsoft's role

- SONiC is a **Linux Foundation project** (moved there in 2022) developing an open-source, Debian-based network operating system for switches, hardware-agnostic and modular **[independent — Network World, Jan 2026]**.
- Started life at **Microsoft**, adapted from its Debian-based Azure Cloud Switch; Microsoft remains a core contributor/user **[independent — The Register, May 27, 2026]**.
- SONiC 4.5 released May 2025, community-supported until at least October 2026 **[independent — Network World]**.
- Commercial vendors (Dell, Cisco, Juniper, Arista, NVIDIA, etc.) support SONiC on their hardware; Arista's 7060XE7 explicitly supports open network OS options (SONIC/OpenSwitch) **[independent/secondary]**.

### 2.2 2026 releases and distributions

**Enterprise SONiC 4.6 (Stordis) — GA June 2026**
- Operations-focused release: route-table **overflow detection with automatic recovery and retry**; improved route consistency checker (per-prefix inspection, transient-route ignore, CLI-triggered recovery); better telemetry; tighter access control; broader hardware list **[secondary — Stordis release breakdown]**.
- Base system moved to **Debian 12 Bookworm**, Linux kernel 6.1; Broadcom SAI Adapter 15.3.0; Broadcom SDK 6.5.35; FRR 8.2.2; Config DB version_4_6_1 **[secondary]**.
- Covers leaf-spine fabrics, Ethernet networks for AI clusters, and campus edge **[secondary]**.
- ⚠️ Release-notes documentation footnote labels Debian 12 as "Bullseye" in one component table; the feature section confirms Bookworm — minor doc inconsistency flagged by Stordis **[secondary]**.

**Netberg SONiC 202511.n0 — released April 15, 2026**
- Successor to the 202411 release; positioned for data centers, AI-driven infrastructure, and campus networks **[secondary — PRLog/Netberg]**.
- Key features: **AI workload optimization** (enhanced **SRv6** support for high-performance AI training/inference backend networks with efficient load balancing); advanced observability (high-frequency telemetry, advanced diagnostics); **Debian 13 (Trixie)** base with updated kernel, security patches, modern toolchain **[secondary]**.

