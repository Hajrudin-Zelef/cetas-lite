---
id: etape6-tracka-cisco-juniper/02-part-2-juniper-networks-post-hpe-acquisition/2-3-mist-ai-campus-branch-hpe-mist
title: "2.3 Mist AI / campus & branch (HPE Mist)"
domain: part-2-juniper-networks-post-hpe-acquisition
role: deep-dive
task: reference
actors: ["AMD", "Microsoft", "Oracle"]
dates: ["2026-04", "2026-05", "2026-06", "2026-09"]
keywords: ["acquisition", "agentic", "amd", "backlog", "copilot", "helios", "latency", "licenses", "nand", "neocloud", "research", "revenue"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [309, 355]
section: "PART 2 — JUNIPER NETWORKS (post-HPE acquisition)"
sha256: c5646240c738fa044e5cc91f71de3c470e2d89bad569aa8034c4b65f21edc658
---

# 2.3 Mist AI / campus & branch (HPE Mist)

### 2.3 Mist AI / campus & branch (HPE Mist)
- **HPE Mist** (renamed from Juniper Mist) **now manages HPE Networking CX switches** — the Juniper AIOps engine moved into the Aruba access layer. **Marvis** — the self-driving framework — is **embedded into Aruba Central**. Strategy: "build once, deploy twice," rolling out from **Q1 2026**. [secondary — NAND Research Discover 2026 recap]
- **Discover Barcelona (Dec 2025):** Juniper's **Large Experience Model** (analyzes data from Zoom/Teams etc.) added to Aruba Central; Aruba's **Agentic Mesh** coming to Mist, improving issue detection, root-cause and remediation. [secondary — SiliconANGLE]
- **HPE Networking 723H** — first unified Juniper×Aruba product, shipped **6 May 2026**: a three-radio 2×2 **Wi-Fi 7 (802.11be)** wall-plate/hospitality AP that boots into **either Mist or Aruba Central** management (auto-detects at claim; existing licenses honored across platforms). HPE says **all future wireless hardware will follow this dual-platform strategy**. [secondary — The Register; Tech Field Day]
- **"Self-Driving Networks"** launch (May 2026): autonomous workflows executed without operator approval — dynamic wireless capacity optimization (e.g. re-provisioning capacity for an all-hands gathering), VLAN mismatch/missing-VLAN correction, rogue DHCP server neutralization, RF channel retuning for interference, client roaming refinement, end-to-end latency measurement from client association to cloud, and **automatic avoidance of frequencies reserved for military/priority users** (policy-driven channel allocation beyond standard DFS). Humans can remain in the loop. [secondary — The Register, 8 May 2026]
- **Market signals:** HPE Q2 FY26 earnings — **Wi-Fi 7 AP sales up 7×**; campus & branch orders hit a record, +20% normalized, with multimillion-dollar deals across verticals. [official — HPE Q2 earnings call]
- HPE showcased its wireless infrastructure role at the **2026 Winter Olympics** (served over 1 million clients across 5/6 GHz bands); 6 GHz adoption at ~60% of HPE AP shipments; Wi-Fi 7 at 10–15% adoption per HPE. [vendor-reported — Tech Field Day]
- Older Marvis/Mist capabilities (Minis digital twins, Wired/WAN Assurance) remain current; no discrete 2026 Mist platform re-launch found beyond the HPE Mist rebrand and cross-pollination work.

### 2.4 Security
- **SRX400 series** firewalls launched at **RSA Conference 2026** (24–26 March, Moscone), available Q2 2026 — carrier-grade security for smaller/space-constrained edge sites, compact form factor, hardware tamper protection, part of the hybrid mesh firewall. HPE credited the launch with security orders +15–19% (normalized) in Q2 FY26. [official — HPE press release; earnings call]
- At Discover 2026: **quantum-safe SRX4700** (AI Predictive Threat Prevention), **SASE Orchestrator** (unifies SD-WAN and SSE in one console — due later 2026), **SASE Copilot**, **AI Firewall**, Universal ZTNA extended to non-human/agentic identities, Security Director Copilot. [secondary — Futurum]

### 2.5 Apstra / Junos OS
- **Apstra:** no discrete 2026 Apstra release found in this research. It remains the intent-based networking layer (Apstra Data Center Director) underpinning HPE's AI fabric management alongside Ops4AI; the QFX5240 store page cites Apstra Data Center Director for daily-operation assurance of AI training fabrics. [official — HPE Store; no 2026-dated version found — gap flagged]
- **Junos OS:** latest documented AI-ML feature set is in Junos EVO 24.4-era release notes (GLB, DLB FlowSet, reactive path rebalancing, rdma-opcode support); no new named Junos release found for 2026. Marketing emphasis is on "AI-optimized Junos software (e.g. AI load balancing)" rather than a versioned release. [official — Juniper docs; gap flagged]

### 2.6 ACX
- No 2026 ACX hardware announcements found. **Gap flagged** — if ACX launches occurred, they were not surfaced by this search pass.

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

