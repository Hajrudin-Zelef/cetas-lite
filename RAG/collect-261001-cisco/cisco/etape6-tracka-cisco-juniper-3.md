---
id: collect-261001-cisco/cisco/etape6-tracka-cisco-juniper-3
title: "Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026"
domain: cisco
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "Nvidia", "OpenAI", "United States"]
dates: ["2026-04", "2026-04-14", "2026-05-04", "2026-06", "2026-06-02", "2026-07"]
keywords: ["agent", "agentic", "agents", "amd", "aws", "cost", "cpo", "cyber", "data centre", "dci", "dsp", "ethernet"]
source: docs/RAG/collect-261001-cisco/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [132, 185]
sha256: fa9cc3a4f7acaf57541d17379ef2dbe50fbce3ce817b382bea9ee816dbb68db8
---

# Step 6 Track A — Cisco + Juniper (Enterprise Data-Center Networking), 2026

### Cisco Cloud Control + AgenticOps (Cisco Live, Las Vegas, June 2, 2026)
- **Cisco Cloud Control** (controlled availability in the US from June 2, 2026; global availability planned July 2026): unified platform across networking, security, observability, collaboration; foundation of the "AgenticOps" model [independent via CRN/Presidio].
- **AI Canvas**: multiplayer workspace where operators + AI agents work from shared live telemetry; context persists across shifts [independent].
- **Agent Builder / App Builder / Cloud Control Marketplace** (integrations: ServiceNow, Atlassian, BMC, Okta, Ping, Entra ID, Jamf, LiveAction, Panduit, NetBox, Device42, Vertiv, Anthropic, OpenAI, NVIDIA, Collibra) [independent via GIXtools].
- **Agentic Actions for networking** (beta June 2026 via Meraki): sense → diagnose → remediate → validate → deploy loop; Experience Metrics; Deep Reasoning; **Digital Twin** (alpha July 2026) — emulated network replica using actual software images [independent via GIXtools].
- **Cisco Multicloud Fabric**: cloud-delivered service connecting branches/DCs/clouds (AWS, Azure, GCP, neoclouds), managed overlay, no customer hardware, zero-trust routing, ThousandEyes + Splunk observability [independent].

### Congestion-control / AI-cluster networking features (Nexus 9000)
Cisco's AI/ML fabric feature set (from Cisco blog + white papers, [official]):
- **Dynamic Load Balancing (DLB)**: distribute traffic across equal-cost paths/links.
- **Priority Flow Control (PFC)**: pause frames to prevent Ethernet frame drops → lossless Ethernet for RoCEv2.
- **Explicit Congestion Notification (ECN)**: end-to-end per-IP-flow congestion marking without drops.
- **WRED** (Weighted Random Early Detection) for traffic-class drop probabilities.
- **Cisco Intelligent Packet Flow**: proactive congestion management for non-uniform inter-switch link utilization, "tailored to customer-specific requirements" [official via Cisco AI-networking white paper].
- Lossless transport for **RoCEv2** with Data Center Bridging (DCB); sub-microsecond latency on 400G/800G interfaces [official/independent].
- DCN licensing note: fine-grain load balancing on 800G switches sold as perpetual license **DCN-FGLB-XF3** (seen in AI POD BOM) [official via CVD].
- PFC/ECN are also the technologies the **Ultra Ethernet Consortium** is developing for AI infrastructure — Cisco positioning aligns [independent via Network World].
- **Nexus Dashboard** as ops/automation platform: Insights (visibility), Fabric Controller (automation) [official].

### Nexus Hyperfabric AI
- Cloud-managed fabric with intent-based networking: cloud design tools, zero-touch deployment, declarative policy, automatic GPU-cluster detection; workload-aware observability (network behavior correlated with AI job performance), ML-powered root-cause [secondary via Medium author's article — flagged secondary; Cisco's official Hyperfabric branding exists but details above are from a secondary source].

---

## 7. 2026 acquisitions and partnerships (networking/AI infrastructure relevance)

### Acquisitions
1. **Galileo Technologies Inc.** — AI observability/"AI trust" platform; announced ~April 2026 (Cisco's "second planned deal of 2026"); terms **undisclosed**; expected close in **Q4 FY2026**; strengthens Splunk Observability Cloud AI-agent monitoring [secondary via CRN/TipRanks]. Network World (April 2026) described Galileo as "purposeful observability for AI agents across the full AI development lifecycle" [independent].
2. **Astrix Security** — Israeli startup (founded 2021; non-human identity / AI-agent security: discovers AI agents, maps non-human identities — API keys, service accounts, OAuth tokens, MCP servers; lifecycle management, automated remediation). Timeline: "advanced talks" reported April 14, 2026 at **$250–350M** [secondary via The Information/CRN]; Cisco announced **intent to acquire May 4, 2026** [independent via SecurityWeek mirror (obstracts)]; **~$400M** per Calcalist/CTech [secondary, flagged as press-reported not official]. Terms officially undisclosed. Astrix had raised $85M total ($45M Series B Dec 2024, Menlo Ventures-led) [secondary]. Planned integration into identity intelligence, secure access, Duo IAM [independent].
3. **NeuralFabric Corp.** — enterprise AI platform (generative AI, data sovereignty); announced end of 2025 (context, just outside 2026) [secondary via CRN].
4. **EzDubs** — real-time speech translation; end of 2025 [secondary via CRN].

### Partnerships
1. **AMD + Cisco + HUMAIN joint venture** (announced Nov 19, 2025; operational ramp 2026) — up to **1 GW of AI infrastructure in Saudi Arabia by 2030**; initial **100 MW** phase with AMD Instinct MI450-series GPUs + Cisco networking/critical infrastructure; HUMAIN (PIF-backed) provides data centers. AMD and Cisco are **exclusive technology partners and founding investors**. First customer: **Luma AI** contracted the full 100 MW. By **Aug 31, 2026**, the first segment became operational running **AMD MI355X GPUs**; up to 250 MW more planned from 2027 [independent via Data Centre Magazine/StorageReview/CryptoBriefing]. Note: announced Nov 2025, so 2026 status = first phase operational.
2. **Cisco + Teleport** (announced Aug 25, 2026) — strategic partnership on "Infrastructure Identity" (short-lived, task-scoped access for humans and AI agents); Cisco became Teleport's **largest strategic investor**; first use case: privileged access for network/infrastructure admins [secondary via Cyber Magazine]. Security-flavored, not DC networking hardware.
3. **Cisco joins Verizon's 6G Innovation Forum** (Sept 17, 2026) as core partner — AI-native networking research for enterprise/carrier infrastructure [secondary via Simply Wall St].
4. **Splunk + AWS** multi-year joint-development pact for agentic SOC products (announced ~Sept 2026 alongside Splunk.conf) [secondary].
5. AI POD ecosystem partners (ongoing 2026): NVIDIA, VAST Data, NetApp, Pure Storage, Nutanix, WWT, CDW, Accenture, bitsIO, Wipro, Computacenter [official/vendor-reported].

---

## 8. 400G/800G portfolio, optics, CPO

### Pluggable optics launched Feb 2026 (with G300)
- **1.6T OSFP** pluggable optics: ultra-high-bandwidth for AI scale-out; targets 1.6T switch-to-NIC links and 1.6T/800G/400G/200G switch-to-server links; high performance + reliability [official]. Both ship **this calendar year** per Cisco (CRN) [independent].
- **800G Linear Pluggable Optics (LPO)**: **50% lower optical-module power vs. retimed modules**; with LPO-supporting N9000/8000 systems, customers can cut **overall switch power by 30%** [official]. Built with Cisco silicon photonics technology [secondary via Zacks].
- AI POD 2-8-9-400 BOM shows: QSFP-400G-DR4 (400G QSFP112), QDD-8X100G-FR dual-port 800G, QSFP-200G-SR4-S in use [official via CVD].

### Optical transport (DCI/backbone, context)
- **Cisco Open Transport 3000 Series** multi-rail open line system (2026): Cisco claims 75% power reduction and 80% rack-space reduction per rail vs. prior single-rail systems (C-band and C&L-band) [independent via Light Reading]. Primarily hyperscalers/SPs; "14× capacity" framing vs. legacy DCI [independent].
- **NCS 1014** upgraded with 800G pluggable transponder line card: 12.8T capacity (16× 800G in one card), 50% less rack space, 38% less power vs. prior 400G generation [independent via Light Reading].
- **New 100ZR coherent pluggable** using Cisco's own silicon photonics + DSP; 100–150 km reach; extends Routed Optical Networking into access/edge [independent via Light Reading].
- Acacia (Cisco's coherent unit): record Q3 FY2026 — **>750,000 400G + 40,000 800G coherent pluggables shipped**; five new hyperscaler design wins (incl. Silicon One P200/G200 systems); CEO: Acacia to **grow 200% in FY2026** [independent via Lightwave]. "Commanding share" of the ZR/ZR+ market [independent].

