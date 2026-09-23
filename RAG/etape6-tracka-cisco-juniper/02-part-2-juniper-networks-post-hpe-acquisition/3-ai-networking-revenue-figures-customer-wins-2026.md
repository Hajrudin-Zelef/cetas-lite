---
id: etape6-tracka-cisco-juniper/02-part-2-juniper-networks-post-hpe-acquisition/3-ai-networking-revenue-figures-customer-wins-2026
title: "3. AI networking revenue figures & customer wins (2026)"
domain: part-2-juniper-networks-post-hpe-acquisition
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Huawei", "Nvidia", "Oracle", "TSMC", "United States"]
dates: ["2026-03", "2026-04", "2026-05", "2026-06", "2026-09"]
keywords: ["revenue", "accelerator", "acquisition", "amd", "backlog", "blackwell", "coherent optics", "compute", "cpo", "dci", "ethernet", "gpus"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [329, 410]
section: "PART 2 — JUNIPER NETWORKS (post-HPE acquisition)"
sha256: a160aa42e66aab696fb33e4e6282dc0fe0758a92b80f16e224bf9c2dfee57f6a
---

# 3. AI networking revenue figures & customer wins (2026)

## 3. AI networking revenue figures & customer wins (2026)

HPE fiscal year ends 31 October. Juniper figures are reported inside HPE Networking.

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

- **Market data (IDC, Q2 2026)** via Next Platform (Sept 20, 2026) and kad8:
  - Global Ethernet switch revenue **$18.9B** in Q2 2026 (+64.5% y/y for the data-center segment, $12.3B). [independent — IDC via Next Platform/kad8]
  - **Cisco: $5.43B, +36.2%** (28.7% global share); ~$2.2B from data-center switches; router business $1.6B (+31.6%). Largest vendor overall. [independent]
  - **NVIDIA: ~$2.5B (per kad8, +181% y/y; ~20% of DC Ethernet) — but Next Platform cites $3.86B growing 2.8×.** The two figures appear to use different scopes; **flagged as unresolved.** [independent — discrepancy noted]
  - **Arista: $2.53B, +37.6%** (13.4% global; 18.7% DC). Arista launched a modular router at >115 Tbps in early 2026 aimed at AI. [independent — IDC via Next Platform; Arista launch via Think Insights]
  - **HPE (incl. Juniper): $1.15B Ethernet, +6.1%** — "a lot of incremental business for HPE, whose Ethernet business was largely wireless under Aruba." Huawei $1.55B (+28.6%); ODMs ~$3.56B (+2.43×). [independent — Next Platform's IDC augmentation]
  - **Reading:** HPE-Juniper is #4 by Ethernet revenue (behind Cisco, NVIDIA/Arista, Huawei) but is the only vendor with full campus-to-DC-to-routing plus compute/storage; its growth lags the AI-cluster leaders (NVIDIA, Arista) in DC switching share.
- **Silicon positioning:** Juniper claims Express 5 (28.8 Tb/s) beats Cisco P100 (19.2 Tb/s) by ~33% throughput — a vendor claim Cisco declined to confirm or deny. [vendor-reported]
- **Analyst commentary:**
  - Futurum (Discover 2026 recap): "the networking story is the strongest, most defensible part of HPE's hand"; honest tension is the two coexisting management planes (HPE Mist vs Aruba Central); key tests = self-driving delivery in brownfield, security story above the network, NVIDIA dependency. [independent]
  - ACG's Ray Mota: "HPE delivers the architectural and operational foundation service providers need to fully participate in the AI value chain." [independent quote]
  - EMA's Shamus McGillicuddy: Juniper was showing more value than anyone else; HPE's acquisition centers on that. [independent]
  - Think Insights industry analysis: HPE-Juniper is a new full-stack competitor "explicitly positioned to challenge Cisco and Nvidia in AI-native networking"; Arista retains hyperscale/DC advantage; white-box/SONiC competes on price. [independent]
  - ETR AI survey cited by Futurum: ~21% of orgs building AI apps use HPE for hosting/deployment (and plan to stay), 14% considering — "most of the field is still open." [independent]
  - HPE's own claim (Feb 2026, MWC): "I don't think you could find one company who has a leadership in compute, storage, security and networking across the board" — the full-stack pitch vs Cisco/Arista. [vendor-reported]
- **Customer wins vs Cisco (2026):** Oracle's gigawatt-scale HPE Juniper win is the marquee 2026 reference; Wi-Fi 7 AP sales 7× and campus record orders suggest campus displacement momentum; no 2026-dated named hyperscale DC wins against Cisco surfaced. [official on Oracle; gaps flagged]

## 6. Optics and co-packaged optics (CPO), 2026

- **No Juniper-specific CPO product announcement found in 2026.** This is a notable absence given Broadcom's 2026 CPO push. [explicit gap]
- **Adjacent CPO facts (context):**
  - Broadcom's **Tomahawk 6** (102.4 Tbps) ships with a **native CPO variant** (built on TSMC photonics tech); Juniper builds on Tomahawk 6 for its AI switches and endorsed the chip's "UEC-compliant architecture" and CPO direction. [official — Broadcom PR; vendor-reported Juniper quote]
  - Industry context: TSMC's first true CPO implementation (COUPE photonics) entering production in 2026; NVIDIA's Spectrum-X Ethernet Photonics in production alongside Vera Rubin; 1.6T optical modules ramping in 2026. [secondary]
- **Juniper optics that did move in 2026:** **800G ZR+ coherent optics** on PTX/MX for distributed AI-factory DCI (HPE AI Grid / GTC 2026 announcements; also the AI Factory Grid architecture with NVIDIA). [official/secondary]
- Broadcom's 800G **Thor Ultra NIC** is positioned with the "HPE Juniper Networking AI-optimized fabric" — end-to-end fabric messaging rather than a Juniper transceiver product. [official — Broadcom PR]

## 7. Key uncertainties & gaps

1. **UEC membership:** no 2026 source confirms Juniper in the Ultra Ethernet Consortium; only that Juniper endorses UEC-compliant merchant silicon. [unverified]
2. **CPO product:** no Juniper CPO product announcement; strategy appears to ride Broadcom's CPO roadmap. [gap]
3. **Apstra 2026 version:** no discrete 2026 Apstra release found. [gap]
4. **Junos OS 2026 release:** no new named Junos version for 2026 found; latest AI-ML docs are 24.4-era. [gap]
5. **ACX 2026:** no announcements found. [gap]
6. **DCQCN/RoCEv2 2026 news:** none found (only pre-existing Junos EVO RoCE/AI-load-balancing docs). [gap]
7. **NVIDIA IDC revenue discrepancy** ($2.5B vs $3.86B for Q2 2026) unresolved. [flagged]
8. **Deal value** cited as both $14B and $13.4B (likely gross vs net of cash); used "$14B (some outlets $13.4B)". [flagged]
9. **Instant On divestiture** details (buyer, timing) rest on a single buyer-guide source; DOJ-approved buyer unknown. [unverified]

## 8. Sources consulted (all via web search/fetch, Sept 2026)
- CRN (deal close); The Stack (close/analyst quotes); ainvest (Aug 2026 ruling, Judge Pitts); appliedantitrust.com (HPE court filing, Dec 2025); networkdevicesinc.com (buyer guide); NAND Research (Discover 2026 recap); Futurum Group (Discover 2026 analysis, 24 June 2026); SiliconANGLE (Discover Barcelona, Dec 2025); The Register (723H AP, 8 May 2026); Tech Field Day (Wi-Fi 7 updates); Network World (PTX12000, Express 5); Data Center Knowledge (PTX, MWC 2026); Fierce Network (MWC AI); HostingJournalist (PTX); SDxCentral (Express 5 vs Cisco P100/Nokia FP5); stocktitan (HPE MWC PR); HPE Store UK (QFX5240 SKUs); Juniper.net docs (Junos EVO AI-ML, GLB/DLB); blocksandfiles.com (Q2 FY26); Fool (Q2 FY26 earnings transcript); control.vg (Q2 FY26); Constellation Research (Q2 & Q3 FY26); ConvergeDigest (Q3 FY26, Oracle win); Zacks (Q3 FY26); marketbusinessnews (Q3 FY26); skn-finance (Q3 FY26); dividendinformer (Q3 FY26); Next Platform (IDC Q2 2026 Ethernet data, 20 Sept 2026); kad8 (IDC Q2 2026); Think Insights (industry analysis); ad-hoc-news.de (analyst coverage); Broadcom PR via GlobeNewswire (Tomahawk 6); cxotoday.com (AMD Helios, Dec 2025); businesswire (HPE AI Grid with NVIDIA, 17 Mar 2026); business-news-today/expresscomputer (HPE-NVIDIA sovereign AI); CRN/stocktitan/chartmill/barchart/nomios/saudishopper (SRX400, RSA March 2026); markets.financialcontent.com/urdupure.com/github silicon-photonics research (CPO context).

---

