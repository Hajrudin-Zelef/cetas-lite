---
id: etape6-trackb-arista-sonic/00-arista-sonic/1-arista-networks
title: "1. ARISTA NETWORKS"
domain: step-6-track-b-arista-sonic-cumulus-data-center-fabric-open-
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Broadcom", "Meta", "Microsoft", "Oracle"]
dates: ["2025-12", "2026-05", "2026-06", "2026-06-04", "2026-06-09", "2026-08-04"]
keywords: ["accelerator", "agentic", "amd", "aws", "capex", "compute", "ethernet", "hyperscaler", "latency", "lpo", "maia", "optics"]
source: docs/RAG/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [19, 95]
section: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
sha256: 2689a9b8f41afaade560cd4dd12ef9d61d30d132ecf4e130526fc553c35a7ab5
---

# 1. ARISTA NETWORKS

## 1. ARISTA NETWORKS

### 1.1 Financial results 2026

**Q2 2026 (reported August 4, 2026) — first $3B quarter**
- Revenue **$3.036 billion**, +12.1% QoQ, **+37.7% YoY** vs Q2 2025 **[official — Arista Q2 2026 earnings release; confirmed by The Next Platform, Sep 20, 2026]**.
- GAAP diluted EPS **$0.95**, non-GAAP diluted EPS **$1.02** (+39.7% YoY from $0.73) **[official]**.
- GAAP operating margin 45.4%, non-GAAP operating margin **49.9%** (vs 44.7%/48.8% in Q2 2025) **[official]**.
- CEO Jayshree Ullal: "As we deliver our first $3 billion quarter in Q2 2026, it is clear that our Arista 2.0 platform strategy is compelling. Customers see networking as the central nervous system for infrastructure from the client to campus to data and AI centers." **[official]**.
- CFO Chantelle Breithaupt (appointed CFO 2025) cited "strong, broad-based growth" **[official]**.
- Full-year 2026 revenue guidance raised to approximately **$12.6 billion**, implying roughly 40% growth **[secondary — Tickeron earnings recap, ~Sep 14, 2026]**.
- Beat Wall Street consensus (revenue consensus ~$2.83B, EPS consensus ~$0.88) **[secondary — Tickeron]**.
- Non-GAAP gross margin slipped to 63.4% on customer mix and higher component costs **[secondary — Tickeron]**.

**Q1 2026 (reported May 2026)**
- Revenue **$2.71 billion**, +35.1% YoY (vs $2.62B expected); non-GAAP EPS **$0.87** vs $0.81 forecast **[secondary — CoinCentral, ~June 9, 2026; ainvest]**. GAAP operating margin 42.7%, non-GAAP 47.8% **[secondary]**.
- Q2 guidance at the time: ~$2.8B **[secondary]**.
- Balance sheet (Mar 31, 2026): $2.79B cash + $9.56B marketable securities; $1.69B operating cash flow in Q1; total liabilities $8.17B, equity $13.49B; low-debt profile **[secondary]**.

**Context (secondary investment-press framing, treat figures as approximate)**
- Zacks "Bull of the Day" (Sep 16, 2026): revenue more than doubled 2022→2025 ($4.38B → $9.00B); GAAP EPS $1.07 → $2.75 over the same period; $13.3B cash and equivalents, $23.7B total assets vs near-zero debt and $8.9B liabilities; five consecutive years of EPS beats; Zacks Rank #1 (Strong Buy) **[secondary — Zacks]**.
- Bank of America raised price target to **$200** (Buy) after the 1.6T launch; average analyst target ~$186.47 **[secondary — CoinCentral]**.
- Major shareholder Andreas Bechtolsheim (co-founder) sold 240,000 shares on June 4, 2026 under a pre-arranged 10b5-1 plan **[secondary — CoinCentral]**.
- AI hyperscalers projected to spend ~$800B in AI-related capex in 2026 **[secondary — Zacks]**.

### 1.2 AI networking revenue and Etherlink portfolio

**AI revenue trajectory (vendor-reported/management statements — handle with care)**
- Management expects AI revenue to reach **at least $3.6B in 2026**, supported by scale-up, scale-out and scale-across deployments **[secondary — Zacks, Sep 16, 2026]**.
- Earlier framing (late 2025 / Feb 2026 reports): Arista **doubled its AI networking revenue goal to $3.25B for FY2026**, up from ~$1.5B in FY2025 **[secondary — Financial Freedom is a Journey, Feb 2026]**.
- ⚠️ The $3.6B figure supersedes the $3.25B figure; use the most recent (Sep 2026) as the current target. Both are management guidance, not independently audited.
- Etherlink AI fabric portfolio: **more than 100 cumulative customers** as of Q2 2026, up from "only a handful of early adopters in 2024" **[secondary — Tickeron Q2 2026 recap; Zacks]**.

**Etherlink AI — platform positioning**
- Arista's AI Ethernet fabric portfolio spans **400G, 800G and (since June 2026) 1.6T** switching **[secondary — The CODEW, Aug 2026]**.
- Technical foundation: RoCEv2 + PFC + DCQCN congestion control, Arista dynamic load balancing (DLB), plus per-packet spraying, congestion signaling, and trimming features aligned with the Ultra Ethernet transport trajectory — described as a "bridge today (RoCEv2 + PFC + DLB) with UEC plumbing being added" **[secondary — LLM-Systems-Wiki Arista/Etherlink page, updated Aug 2026; vendor-reported]**.
- **7700R4 Distributed Etherlink Switch**: single-hop distributed architecture designed to scale to **more than 30,000 connected accelerators** while preserving lossless, deterministic forwarding — an alternative to conventional multi-tier spine-leaf for very large clusters **[secondary — The CODEW]**.
- EOS implements RoCEv2, DCQCN, and priority flow control alongside Arista's own load-balancing for AI workloads **[secondary — The CODEW]**.

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

