---
id: etape6-tracka-cisco-juniper/02-part-2-juniper-networks-post-hpe-acquisition/overview
title: "PART 2 — JUNIPER NETWORKS (post-HPE acquisition)"
domain: part-2-juniper-networks-post-hpe-acquisition
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Microsoft", "United States"]
dates: ["2024-01", "2024-04", "2025-06", "2025-07", "2026-05", "2026-06", "2026-08", "2026-09"]
keywords: ["acquisition", "agentic", "amd", "asic", "benchmark", "copilot", "cpo", "dci", "gpu", "helios", "inference", "latency"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [273, 328]
section: "PART 2 — JUNIPER NETWORKS (post-HPE acquisition)"
sha256: 9c827241c8deb642c0b4db2ddabba1b455d4f087761f6d0ed4e6117fa3fc42aa
---

# PART 2 — JUNIPER NETWORKS (post-HPE acquisition)

# Juniper Networks, 2026 — Post-HPE Acquisition Research Report
**Report date:** 22 September 2026 (all figures "as of" this date)
**Provenance tags:** [official] = vendor press release/filing/earnings; [vendor-reported] = vendor claim quoted in media; [independent] = analyst firm/research; [secondary] = media reporting; [unverified] = single weak source / not corroborated. Claims without a tag prefix are the reporter's own synthesis and should be treated cautiously.

---

## 1. HPE–Juniper acquisition status

- **The deal closed on 2 July 2025** — ~18 months after its announcement on 9 January 2024 (Juniper shareholder approval: 2 April 2024). Consideration was ~$14 billion (some outlets cite $13.4B net). Juniper is a wholly owned subsidiary of HPE. [secondary — CRN; The Stack]
- **DOJ settlement (28 June 2025)** cleared the merger just before close. Remedies: HPE must divest its global Instant On campus/branch WLAN business (within 180 days to a DOJ-approved buyer) and must auction a license to Juniper's Mist AIOps source code to competitors. [secondary — NetworkDevicesInc buyer's guide citing the DOJ settlement; corroborated in broad strokes by ainvest reporting]
- **August 2026 court ruling (the item the task notes):** a coalition of state attorneys general (led by Colorado) had challenged the DOJ settlement as politically motivated and insufficient; U.S. District Judge **Casey Pitts** (N.D. Cal.) issued a **41-page ruling in mid-August 2026** dismissing the states' challenge, finding they failed to show the amended final judgment was not in the public interest. This is the final legal clearance: HPE may operate Juniper as a fully integrated entity. [secondary — ainvest, 14 Aug 2026]
- **Integration status:** According to HPE court filings (Dec 2025), HPE moved immediately after closing to integrate employees, products, systems, IP, finance, sales, channel and corporate structures; the business operates as a single **HPE Networking** operating unit. ~4,000 U.S. and ~1,700 non-U.S. former Juniper employees moved onto HPE IT/payroll/benefits; the former Juniper CPO, CSO, CFO, CMO, General Counsel and Chief People Officer are gone. **Rami Rahim** (ex-Juniper CEO) leads the unified HPE Networking business as EVP/President/GM. [official — HPE court filing via appliedantitrust.com]
- **Instant On divestiture status:** Struggling. Reporting says Extreme Networks and Tech Mahindra declined to bid; Fortinet offered only $5M–$15M; one investor group offered $1 for a standalone spinout. Guidance for buyers: do not build new deployments on Instant On. [unverified — single buyer-guide source; DOJ buyer approval status unknown]
- **Earnings commentary:** HPE says the Juniper integration and the "Catalyst initiative" are both running **ahead of schedule**, and raised its Networks-for-AI order target in Q3 FY26 (see §3). [official — HPE earnings calls]

## 2. 2026 product launches

### 2.1 Routing (PTX/MX) — Express 5 / Trio 6
- **Pre-MWC 2026 (announced ~25 Feb 2026):** two new Juniper PTX core router families for AI fabrics, distributed cloud and DCI, all built on the **Express 5** custom ASIC [official/secondary — HPE press release via Network World, Data Center Knowledge, HostingJournalist]:
  - **PTX12000** modular family: **PTX12008** (8 slots, 22RU, up to 345.6 Tbps) and **PTX12012** (12 slots, 32RU, up to 518.4 Tbps), 800GbE port density, **1.6T-ready** platforms, deep buffering, high-radix architecture.
  - **PTX10002** compact fixed-form 2RU routers for metro aggregation, peering, edge and AI data-center networking: up to **28.8 Tbps** throughput, ports up to 800G.
  - Express 5: each chip up to **28.8 Tb/s** (36× 800GbE equivalent); HPE claims **49% better power efficiency** than the previous generation. [vendor-reported]
- **Benchmark claim vs Cisco/Nokia:** Juniper VP Brendan Gibbs claimed Express 5 outperforms Cisco's P100 (19.2 Tb/s, "33% less throughput") and Nokia's FP5. Dell'Oro analyst Shin Umeda: Express 5 = core routing; Trio 6 = edge routing. Cisco declined to comment. [vendor-reported, unchallenged — treat as Juniper claim, not independently verified]
- **MX + Trio 6:** three MX-series routers powered by the **Trio 6** ASIC were announced alongside Express 5 (Oct 2025); Rahim reported "healthy demand" across MX and PTX. [secondary — SDxCentral]
- **MX301 inference-edge router** and **AI-Native Routing** (routers enforcing policy on access to public AI platforms): refreshed/launched at HPE Discover 2026 (June 2026). [secondary — Futurum recap]

### 2.2 Data-center switching (QFX)
- **HPE Discover 2026 (15–18 June, Las Vegas):** refreshed the AI switch lineup [secondary — Futurum]:
  - **QFX5252** — liquid-cooled **scale-up** switch designed for AMD Helios AI racks (see §4, partnerships);
  - **QFX5250** — scale-out switch;
  - **QFX5140** — inference switch.
- **QFX5240** (launched 2025, still the flagship 800G AI fabric platform): Broadcom **Tomahawk 5**-based, 64×800GbE OSFP in 2RU, 51.2 Tbps; positioned as the foundation of HPE's AI data-center fabric with Apstra assurance for AI/ML training, "fast job completion time" for GPU utilization. HPE Store lists multiple SKUs (e.g. QFX5240-64OD, QFX5241-32OD/64OD, AC/DC, ORv3 trays). [official — HPE Store]
- Junos EVO AI-ML fabric software features now documented for the QFX5xxx line: **BGP Global Load Balancing (GLB)** for AI/ML elephant-flow congestion, **Configurable FlowSet table in DLB flowlet mode** (QFX5130/5220/5230/5240/5700), reactive path rebalancing, selective DLB via firewall filters — collectively the **Ops4AI** (Operations for AI) extensions to Junos + Apstra. [official — Juniper docs (24.4 release notes)]

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

