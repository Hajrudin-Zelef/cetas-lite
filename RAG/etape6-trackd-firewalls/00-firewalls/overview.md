---
id: etape6-trackd-firewalls/00-firewalls/overview
title: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["Apple", "Intel"]
dates: ["2026-02-01", "2026-03-10", "2026-05-06", "2026-07-28", "2026-07-29", "2026-09-22"]
keywords: ["advisory", "agent", "agents", "asic", "benchmarks", "chiplet", "cybersecurity", "disaggregated", "foundry", "governance", "intel", "latency"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [1, 58]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: dc39f9a46ba90d7a54c84bdbf970f7cfa87f730e0339c08f0b15e593411837b8
---

# Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)

**Project:** RAG data-collection, Step 6 (Réseau & sécurité)
**Coverage window:** February 1, 2026 → September 22, 2026
**Status:** Research snapshot. Prices and version figures are dated snapshots as of September 22, 2026.

### Provenance legend
- **[official]** — vendor's own blog, docs, release notes, earnings statement, or advisory.
- **[vendor-reported]** — figure claimed by the vendor (benchmarks, market share) without independent audit.
- **[independent]** — reputable third-party press (BleepingComputer, The Register, Dark Reading, Reuters, IDC, Gartner) or independent measurement.
- **[secondary]** — lower-tier press, blogs, analyst summaries, community trackers; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

---

## 1. FORTINET / FortiGate

### 1.1 FortiOS 8.0 — the headline OS release (announced March 10, 2026)

- Fortinet launched **FortiOS 8.0** at its **Accelerate 2026 conference in Las Vegas** on **March 10, 2026** — the new OS release powering the Fortinet Security Fabric **[secondary — https://www.stocktitan.net/news/FTNT/fortinet-introduces-forti-os-8-0-to-expand-secure-networking-with-iv0y66jrnxym.html; https://www.coffeefranchisehub.com/archives/42978]**.
- FortiOS 8.0 groups its updates into three pillars: **AI-driven security, next-generation SASE, and quantum-safe protection** **[secondary — https://securitybrief.co.uk/story/fortinet-s-fortios-8-0-adds-ai-quantum-safe-tools]**.
- **AI usage controls / governance** (new in 8.0):
  - **FortiView for AI**: attack-surface and shadow-AI visibility — shows how AI applications/services are used across the organization, distinguishing sanctioned vs unsanctioned tools **[secondary]**.
  - **AI-aware application control** for generative-AI usage; visibility into **Model Context Protocol (MCP) and agent-to-agent (A2A) interactions** to reduce blind spots **[secondary]**.
  - Enhanced data loss prevention (DLP) powered by **optical character recognition (OCR)** **[secondary]**.
  - AI agents across the Fortinet Security Fabric enabling conversational troubleshooting and configuration in firewall and SD-WAN environments **[secondary]**.
- **SASE in 8.0**: **FortiSASE Outpost** and sovereign deployment options for data residency; unified SD-WAN bundles; multipath IPsec tunnels **[secondary]**.
- **Quantum-safe**: quantum-resilient cryptographic controls, hybrid post-quantum cryptography (PQC) SSL deep inspection **[secondary]**.
- **ZTNA in 8.0** (from the official FortiOS 8.0 release notes): a FortiGate can now act as a **ZTNA service connector**, reverse-proxying service connections upstream to the ZTNA Edge (e.g., FortiPAM or FortiProxy acting as ZTNA access proxy); configuration simplification/modularization (`firewall.access-proxy` split into `ztna.traffic-forward-proxy`, `ztna.web-proxy`, `ztna.web-portal`; new `ztna.destination` object); IPv6 security posture tags/groups for dual-stack ZTNA policies; tag sharing with FortiClient EMS **[official — https://docs.fortinet.com/document/fortigate/8.0.0/fortios-release-notes/743723/new-features-or-enhancements]**.
- Ken Xie quote (CEO): *"FortiOS 8.0 reflects more than 25 years of continued innovation at the intersection of networking and security. As organizations embrace AI, cloud, and increasingly encrypted environments, a unified operating system is essential to reduce complexity, improve visibility, and ensure security can scale without slowing the business."* **[secondary]**

### 1.2 New FortiGate hardware in 2026 (G series)

**FortiGate 3500G + FortiGate 400G — announced May 6, 2026**
- New additions to the **FortiGate G series**, positioned to secure AI workloads and encrypted traffic in data centers and enterprise edges **[official/via GlobeNewswire — https://www.globenewswire.com/news-release/2026/05/06/3288892/0/en/Fortinet-Expands-FortiGate-G-Series-to-Secure-AI-from-the-Data-Center-to-Modern-Enterprise-Edges.html]**.
- Powered by Fortinet's **NP7 and SP5** custom ASICs and FortiOS **[official]**.

**FortiGate 1200G — announced July 28, 2026**
- Midrange G-series appliance: **10G, 25G and 100G connectivity**; **397 Gbps firewall throughput**; aimed at campus, data center and hybrid environments inspecting large volumes of encrypted traffic **[secondary — https://www.csoonline.com/article/4202566/fortinets-new-fortigate-platform-converges-firewall-sase-technologies-2.html; https://www.itp.net/cybersecurity/fortinet-launches-fortigate-1200g-to-advance-firewall-and-sase-convergence]**.
- Fills the gap between campus/edge/branch appliances (e.g., 400G) and flagship data-center platforms (e.g., 3500G) **[secondary]**.
- Ships with **FortiOS 8.0** and can operate as a **FortiSASE Outpost**: the 1200G becomes a local SASE point of presence (POP) inside customer-controlled environments (on-prem sites, private DCs, colocation), enforcing security locally while staying centrally managed from the cloud — targeting data-sovereignty and latency-sensitive workloads **[secondary]**.
- **Availability expected in Q3 2026** **[secondary]**.
- Integrates **FortiGuard AI-Powered Security Services** and **FortiAI** for real-time threat intelligence and AI-assisted SecOps **[secondary]**.

### 1.3 Intel SP6 collaboration (July 29, 2026)

- Fortinet announced a **strategic collaboration with Intel to develop Fortinet Security Processor 6 (SP6)**, combining Fortinet's proprietary security-processor expertise with Intel's design, development, packaging and manufacturing capabilities, strengthening supply-chain resilience and diversity **[official — via Q2 2026 earnings, https://www.globenewswire.com/news-release/2026/07/29/3335555/0/en/fortinet-reports-strong-second-quarter-2026-financial-results.html]**.
- Secondary reporting (TradingView/UDIS): Intel Foundry will co-develop and manufacture the SP6 ASIC; Fortinet is reportedly **the first external cybersecurity customer for Intel Foundry and the first external customer on the Intel 4 node**, with production at **Fab 34 in Ireland**; the SP6 moves from the monolithic 7nm design of the SP5 to a **disaggregated chiplet architecture** for better yields and lower unit costs **[secondary — https://www.tradingview.com/chart/FTNT/vj3CqrYb-Why-Is-Intel-Building-Fortinet-s-Next-Firewall-Chip/; unverified in part — Intel Foundry customer specifics not independently confirmed in this research]**.

### 1.4 Financials — Q2 2026 results (July 29, 2026)

- Revenue **$2.05B (+26% YoY)**; product revenue **$773M (+52% YoY)**; billings **$2.37B (+33% YoY)**; GAAP operating margin 34%; non-GAAP operating margin **38%**; GAAP EPS $0.82 (+44%); non-GAAP EPS $0.90 (+41%); operating cash flow $1.04B; free cash flow $966M **[official — https://www.globenewswire.com/news-release/2026/07/29/3335555/0/en/fortinet-reports-strong-second-quarter-2026-financial-results.html]**.
- FY2026 revenue guidance raised to **19% YoY growth**; Q3 2026 EPS guidance $0.83–0.87 (non-GAAP); FY2026 EPS guidance $3.41–3.47 **[official/secondary — https://www.americanbankingnews.com/2026/09/17/newedge-advisors-llc-increases-holdings-in-fortinet-inc-ftnt.html]**.
- FY2025 context: total revenue **$6.80B**, net income $1.85B, free cash flow $2.21B **[secondary — StockTitan 10-K reference]**.
- CEO Ken Xie framing: customers value Fortinet's **"SASE Firewall"** — firewall, SD-WAN and SASE integrated on a single FortiOS, powered by FortiASIC, with sovereign deployment options **[official]**.

### 1.5 Market position 2026

