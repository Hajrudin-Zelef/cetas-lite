---
id: collect-261001-cisco/cisco/etape6-tracka-cisco-juniper-7
title: "Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia", "Oracle"]
dates: ["2026-03", "2026-04", "2026-06", "2026-09"]
keywords: ["accelerator", "acquisition", "amd", "backlog", "blackwell", "coherent optics", "dci", "ethernet", "gpus", "helios", "neocloud", "nvidia"]
source: docs/RAG/collect-261001-cisco/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [333, 367]
sha256: cc169d43ab305ed5bf55b6e2fd1ca52a779dc40ab5faa48c0b4ae9fbaa94c2a7
---

# Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026

### 3.1 Q2 FY2026 (quarter ended 30 April 2026; reported ~June 2026)
- Total HPE revenue **$10.7B, +40% y/y** (above $10B high-end guidance); GAAP net income $624M (vs $1B loss a year ago); non-GAAP EPS $0.79, +108%. [official]
- **Networking revenue: $2.7B, +148.2% reported / +10% normalized** (normalized to include pre-acquisition Juniper y/y); network operating margin **21.6%** (down from 25.0% on integration costs); Cloud & AI revenue $7.7B, +22.9%. [official]
- Sub-numbers: campus & branch networking revenue **$1.3B**; data-center networking revenue **$320M**; routing revenue **$775M**; campus & branch orders record +20% normalized; enterprise DC switching orders +~20% normalized; routing orders +~30% normalized; security orders +15–19% normalized; Wi-Fi 7 AP sales 7×. [official — earnings call transcript via Fool/Constellation]
- **Networks for AI order target raised to ≥$2B cumulative by fiscal year-end** (June 2026 guidance). [official]
- AI systems: $1.8B new orders in the quarter; **$16.4B cumulative AI bookings**; AI systems backlog record **$5.9B**. [official]

### 3.2 Q3 FY2026 (reported early September 2026)
- Total revenue **$12.2B, +34% y/y** (record); GAAP EPS $1.06 (vs $0.21); non-GAAP EPS $1.11 (above $0.88–0.93 outlook); non-GAAP gross margin 40.4%. FY26 revenue-growth outlook raised to **34–37%**; FCF floor raised to $3.75B; FY27 revenue growth guided **13–17%**. [official — HPE Q3 FY26 release via multiple outlets]
- **Networking revenue: $2.9B, +74.9% reported / +10% normalized**; operating margin **22%**; normalized networking **orders +36%** (well ahead of revenue — demand exceeds shipments; supply constraints are limiting conversion, with purchase commitments more than doubled sequentially). [official — via Zacks, Constellation, ConvergeDigest]
- Sub-numbers: campus & branch **$1.44B** (+31% reported); data-center networking **$382M** (+112%); routing **$788M** (+270% — includes the Juniper routing line); DC switching & routing orders up at high double-digit rates. [official]
- **Networks for AI: $700M orders in Q3; $2.2B cumulative; FY26 target raised to $2.5–3.0B.** Combined AI backlog (AI Systems + Networks for AI): **$7.6B**. [official]
- **Customer win: Oracle** — gigawatt-scale, multi-gigawatt AI cloud buildout deploying HPE Juniper Networking switches and routers: **QFX switching for "Scale Out" + PTX routing for "Scale Across."** Announced on the Q3 earnings call. [official — CEO Antonio Neri via ConvergeDigest]
- HPE said the upcoming AMD Helios platform will incorporate HPE Juniper Scale-Up switches — but the Helios opportunity is **excluded** from the current FY27 networking growth framework. [official — ConvergeDigest]

### 3.3 Customer deployments (non-OEM)
- **Oracle AI cloud** (gigawatt scale, Sept 2026) — see above. [official]
- **2026 Winter Olympics** infrastructure — HPE networking (>1M clients). [vendor-reported]
- Neocloud/CSP AI-fabric deployments are cited qualitatively ("AI factories and inferencing… hyperscalers, neoclouds") but few are named publicly; PTX orders up ~30% normalized on cloud-service-provider deployments in Q2. [official/secondary]
- **Gap:** no named per-quarter Juniper-only AI revenue split post-acquisition; HPE reports "Networks for AI" orders as the AI proxy.

## 4. AI data-center networking strategy

- **Ethernet for AI / open standards:** HPE Juniper's strategy is explicitly standards-based Ethernet: Junos AI load balancing (GLB/DLB), multivendor fabric management (Apstra), and Mist AIOps layered on merchant silicon — Juniper's message is "secure, high-performance, multivendor, lower TCO" vs. closed stacks. [vendor-reported — Praveen Jain quote in Broadcom Tomahawk 6 PR]
- **Ultra Ethernet / UEC participation:** not confirmed. Juniper praised Broadcom's Tomahawk 6 as an "open, **UEC-compliant** architecture," but **no source in this pass confirms Juniper/UEC Consortium membership**. [flagged — unverified]
- **Congestion management (DCQCN / RoCEv2):** Junos EVO docs reference RoCE/rdma-opcode support and AI-ML load-balancing features, but **no 2026-dated Juniper announcement on DCQCN or RoCEv2 congestion features** was found. [gap flagged]
- **400G/800G portfolio:** QFX5240/5241 (64×800G / 32×800G, Tomahawk 5), PTX10002 (800G fixed), PTX12000/PTX10000 (800GbE, 1.6T-ready platforms); 1.6T readiness is a 2026 theme but no 1.6T ports ship yet. [official/secondary]
- **Partnerships:**
  - **Broadcom:** deepening. Dec 18, 2025 — HPE announced the **first AMD "Helios" AI rack-scale architecture** with integrated **scale-up Ethernet built with Broadcom**: purpose-built HPE Juniper Networking scale-up switch + software, using Broadcom's **Tomahawk 6** (102.4 Tbps), based on the open **Ultra Accelerator Link over Ethernet (UALoE)** standard, supporting trillion-parameter training. HPE claims **first OEM to productize a 100% liquid-cooled Tomahawk 6 switch** (from the Q2 FY26 earnings call). The Tomahawk 6 launch PR quotes Juniper SVP Praveen Jain (AI Clusters & Cloud Ready Data Center). Broadcom's industry-first **800G Thor Ultra NIC** is positioned alongside the HPE Juniper AI-optimized fabric. [official/secondary]
  - **AMD:** co-launch of the Helios rack-scale architecture (scale-up Ethernet over UALoE). [official]
  - **NVIDIA:** HPE "AI Factory" and **AI Grid with NVIDIA** (announced at NVIDIA GTC, March 2026): HPE Juniper PTX/MX routers + **800G ZR+ coherent optics**, MACsec, IP/MPLS for scale-across/DCI; HPE AI Factory with **Spectrum-X Ethernet**, BlueField-3, Blackwell GPUs, Vera Rubin NVL72 on the roadmap; AI Factory Lab in Grenoble (announced Dec 1, 2025, opened Q2 2026); a "3-2-1 integration strategy" (three silicon platforms from two companies into one fabric). HPE says the NVIDIA relationship is a two-way dependency: HPE adds integration, sovereignty, financing, GSI co-sell. [official/secondary — businesswire; Futurum]
- **Scale-out / scale-up / scale-across taxonomy:** HPE now frames AI networking as scale-out (QFX, e.g. Oracle), scale-up (Tomahawk 6-based switch for Helios), scale-across/DCI (PTX12000). [official]

## 5. Positioning vs Cisco in AI data-center networking (2026)

