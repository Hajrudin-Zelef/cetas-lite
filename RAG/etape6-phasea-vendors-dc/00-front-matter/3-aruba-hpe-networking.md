---
id: etape6-phasea-vendors-dc/00-front-matter/3-aruba-hpe-networking
title: "§3 Aruba (HPE Networking)"
domain: front-matter
role: reference
task: reference
actors: ["AMD"]
dates: ["2026-05"]
keywords: ["acquisition", "agentic", "amd", "backlog", "cost", "inference", "license", "nand", "optics", "pricing", "research", "revenue"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [145, 221]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: a3ced48485652d20f6a8b790d4cd3cae15d7a6e804978e072b7a0743fdc3b3cb
---

# §3 Aruba (HPE Networking)

## §3 Aruba (HPE Networking)

### 3.1 Portfolio positioning after the Juniper acquisition

- HPE's stated 2026 portfolio split: **Aruba owns the campus; Juniper owns the data center and edge routing** — both lines under active R&D; "no product lines are being deprecated" and consolidation "is not imminent, or even a priority" (TechTarget, Jun 2026) [independent; vendor-reported].
- QFX5250 (Tomahawk-6, 102.4T, liquid-cooled) is HPE's AI-scale DC flagship (Juniper line); Aruba CX covers distributed-services ToR, aggregation, and campus [secondary — networkdevicesinc.com buyer guide].

### 3.2 CX 10000 — distributed services switch (AMD Pensando)

- **HPE Aruba Networking CX 10000-48Y6C**: combines AOS-CX with **AMD Pensando Elba DPU (P4-programmable)**; **3.6 Tbps** line-rate switching; 48× 1/10/25G SFP28 + 6× 40/100G QSFP28, 1U; delivers **800G of distributed stateful services** — east-west firewalling, NAT, encryption, telemetry, micro-segmentation — inline on every port, eliminating backhaul to centralized firewalls; managed via **Aruba Fabric Composer** with unified network+security policy [official — HPE datasheets; secondary — reseller specs].
- 2025 follow-on: **CX 10040** platform pairing AMD programmable DPU with HPE DC networking for 400G-class distributed services; HPE claims 10× scale/performance at 1/3 the cost of traditional enterprise solutions [vendor-reported — HPE press, 2025-05].
- Pricing snapshot (2026, reseller street): CX 10000-48Y6C listed **~$44,452–$46,570** (Avendor/spaceboundsolutions), vs ~$55,639 list [secondary — unverified street, list inferred].
- No 2026 hardware refresh of CX 10000 was announced; the line continues as HPE's distributed-services story [gap noted].

### 3.3 CX data-center/aggregation models

- **CX 9300**: data-center/aggregation switch family (100G/400G class) — current in 2026 lineup; detailed new 2026 SKUs not surfaced [gap].
- **CX 8360**: data-center switch (25/100G ToR) — current; 2026 updates not located [gap].
- **CX 6300M** (2025): campus access/aggregation, VSF stacking to 10 members, BGP/EVPN/VXLAN, Smart Rate multi-gig, 90W PoE, MACsec-256 [vendor-reported].
- **CX 6000**: 8-port compact models added at NRF 2026 (Jan) for retail/checkout-lane deployments, PoE/non-PoE [official].

### 3.4 Management: Aruba Central + Mist convergence (2026)

- **HPE Discover 2026** (Las Vegas, Jun 2026): "build once, deploy twice" — **Marvis AI actions extended into Aruba Central** (wired port remediation, self-driving capabilities from May 2026); **Aruba CX switches now manageable from HPE Mist** with predictive DC analytics (hardware/optics failure prediction) [vendor-reported — Nand Research, Network World, CRN].
- The combined Aruba+Juniper AI teams formed a single data-sciences group producing an **agentic mesh framework** (agentic orchestration layer across Aruba and Juniper domains); GreenLake Intelligence as the strategic control plane [vendor-reported — TechTarget].
- **HPE Networking Data Center Director**: single console for Juniper-derived switching (replacing separate Apstra/Junos Space workflows) [secondary — Nand Research].
- Unified AI-native SASE platform on EdgeConnect (SD-WAN + SSE, embedded ZTNA connector, sovereign SASE) announced at Discover 2026 [vendor-reported].

### 3.5 AI networking features

- CX 10000's DPU-based inline telemetry/encryption and Fabric Composer automation are HPE's "networks for AI" edge story; Mist DC assurance adds predictive analytics across power, temperature, optics, and system health [vendor-reported].
- AI-fabric scale (Tomahawk-6, 1.6T, UEC) sits on the Juniper side (QFX5250/PTX12000); Aruba CX is positioned for enterprise DC, inference edge, and distributed services [secondary].

### 3.6 Financials and market

- HPE Networking: **$2.7B Q1 FY26** (+10% normalized); **$2.9B Q3 FY26** (+74.9% reported, 22% op margin, orders +36%); **Networks-for-AI: $700M Q3 orders, $2.2B cumulative, FY26 target $2.5–3.0B**; combined AI backlog $7.6B [vendor-reported — HPE earnings, via step-6 track A]. Aruba-specific revenue is not broken out [gap].
- Aruba CX switch street pricing beyond CX 10000 was not systematically located [gap].

---

## §4 Cisco Meraki

### 4.1 Portfolio structure (2026)

Meraki MS is Cisco's cloud-managed switching line; the 2026 lineup spans compact access to 100G aggregation [official — Meraki documentation]:

| Family | Positioning | Key specs |
|---|---|---|
| **MS130 / MS130R** | cloud-first access, compact/rugged | 8/24/48× 1G, mGig options, 10G SFP+ uplinks; MS130R for warehouses/outdoor |
| **MS150** | stackable multigigabit access | 12 models, 24/48-port, up to 60W PoE++ per port |
| **MS210** | mid/large access (5 models) | 24/48× 1G, PoE variants (370W) |
| **MS390** | high-performance L3 access — first Catalyst-hardware-based Meraki switch | 24/48-port, multigigabit, **480G stacking**, modular 10/40G uplinks, StackPower, Adaptive Policy (SGT over-the-wire segmentation) |
| **MS355** | aggregation | up to **688 Gbps** switching, 400G stacking, mGig + 10G SFP+, 2× 40G QSFP+, UPoE 740W |
| **MS450** | 100G aggregation | 12× 40G QSFP+ + 2× 100G QSFP28 uplinks, L3 w/ OSPF, 400G stacking, hot-swap PSU/fans |

Datasheets for MS390 and MS450 were refreshed **Sep 4, 2026** and **Jul 21, 2026** respectively — confirming both families current [official].

### 4.2 2026 dashboard and platform updates

- **AI-powered support cases** (Aug 31, 2026): Client VPN and switching cases can be submitted through an AI Assistant inside the Meraki Dashboard, with environment analysis and troubleshooting suggestions before case creation [official — Cisco Community announcement].
- **Catalyst cloud convergence continues**: IOS XE 17.15.3+ cloud operating mode and hybrid operating mode for Catalyst 9000 (public beta → stable RC track, 2025); **Catalyst 9500 and C9610 smart-switch device-configuration support** in dashboard (2025), extending dashboard management from access to core; Cloud Monitoring EoS path ended Mar 31, 2026 [official — Meraki Community].
- **Cisco Live 2026** (Jun): new "agentic network" devices branded **Cisco** (not Meraki/Catalyst) — one hardware, one OS, choice of Meraki dashboard / Catalyst Center / hybrid; includes a **new 9550 fixed core campus smart switch with 400G uplinks**, 8600 secure router, 8100/8200/8300 variants, Wi-Fi 7 outdoor AP [secondary — RCR Wireless analyst angle].
- Meraki for Government (FedRAMP ATO) available [secondary].
- No brand-new MS hardware family was announced in 2026; the MS130/MS150/MS210/MS355/MS450/MS390 stack stands [gap noted].

### 4.3 Cloud-managed DC positioning

- Meraki's DC story is **aggregation/campus-core for enterprise and mid-market**, not hyperscale AI fabric: dashboard-managed MS355/MS450 aggregation, Catalyst 9500/C9610 for core, with single-pane visibility, virtual stacking across sites, L7 application visibility, and Adaptive Policy zero-trust segmentation [official; secondary].
- Cisco's broader DC/AI-fabric play (Nexus, Silicon One, AI PODs) sits outside Meraki; the two converge via cloud-managed Catalyst [secondary].

### 4.4 Pricing

- Meraki uses hardware + **per-device cloud license** (1/3/5/7/10-yr terms); switch hardware pricing is partner/quote-based. UK reseller listings confirm MS355-48X2-HW as a current orderable SKU (specs as above); numeric street prices were not consistently published [secondary — cloudappliances.co.uk].
- License pricing specifics for 2026 were not located [gap].

---

