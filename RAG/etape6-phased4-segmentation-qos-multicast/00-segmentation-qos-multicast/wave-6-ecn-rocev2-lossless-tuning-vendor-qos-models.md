---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-6-ecn-rocev2-lossless-tuning-vendor-qos-models
title: "Wave 6 — ECN, RoCEv2 lossless tuning, vendor QoS models"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Huawei", "Meta", "Nvidia"]
dates: []
keywords: ["asic", "distribution", "ethernet", "gpu", "latency", "mlperf", "nvidia", "throughput", "training", "voice"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [149, 206]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: e525f724e92262fb90072f1100602c0c25a58dae1c3c88c197e28dcfb35734cd
---

# Wave 6 — ECN, RoCEv2 lossless tuning, vendor QoS models

## Wave 6 — ECN, RoCEv2 lossless tuning, vendor QoS models

### 6.1 The three-layer lossless model
- Layer 1 — PFC: keeps a single hop lossless; does not scale across hops without a watchdog; not the answer to congestion `[secondary]` — Kar lab notebook.
- Layer 2 — ECN (RFC 3168): switches mark (not drop) packets when queues build; the receiver returns Congestion Notification Packets (CNPs) to throttle the sender before buffers overflow `[secondary]` — Kar; AICPLight.
- Layer 3 — DC-QCN: sender-side rate controller reacting to ECN/CNP; when tuned so ECN fires before PFC thresholds, PFC almost never fires and tail latency drops by an order of magnitude (lab observation: DC-QCN min 150 KB / max 1500 KB / 10% max-mark-prob on ConnectX-7 + Tomahawk 5) `[secondary]` — Kar notebook. Treat thresholds as testbed-specific, not defaults.
- End-to-end consistency requirement: DCB/PFC/ECN/DSCP must be configured identically on NICs, switches, and every link — otherwise PFC cannot prevent loss `[secondary]` — intelligentvisibility guide; FS.com whitepaper. https://intelligentvisibility.com/ai-networking-solutions/lossless-networking-ai

### 6.2 Concrete config snapshots (verbatim, dated)
- Dell OS10 (SmartFabric OS10 10.5.0) RoCEv2 with ECN `[official]` — Dell India manual: queuing class-maps Q0 (queue 0) and Q3 (queue 3); policy-map `policy_2Q` gives Q0 30% / Q3 70% bandwidth; WRED profile `wred_ecn` with `random-detect ecn`, green min 1000/max 2000, yellow 500/1000, red 100/500, drop-probability 100; qos-map maps qos-group 3 → queue 3 (lossless) and groups 0–2,4–7 → queue 0. https://www.dell.com/support/manuals/en-in/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-0/configure-roce-on-the-switch?guid=guid-ddfa4455-8f64-4014-8543-ceb6719c904d&lang=en-us
- NVIDIA/Mellanox Zero-Touch RoCE (ZTR) `[vendor-reported]` — FS.com whitepaper (2025-06): auto-configures PFC, ECN on all priorities, DCQCN based on ECN, RoCEv2 CNP marked DSCP 46, MTU 1500 default; `mlnx_qos` dump shows DCBX OS-controlled, dscp2prio mapping table, PFC per-priority enable bitmap, per-TC buffer sizes. https://resource.fs.com/mall/resource/cn_lossless-network-for-rdma-white-paper-20250626115109.pdf
- NetPilot AI-cluster pattern `[secondary]` (2026-09): eBGP underlay leaf-spine, PFC on traffic class 3 + ECN marking on the same class for RoCEv2, single flat VLAN 100 for GPU east-west; ECMP hashing tuned for elephant flows. https://github.com/netpilot-labs/example-prompts/blob/HEAD/data-center/ai-cluster-rocev2-fabric.md
- GitHub rdma-ai-cluster lab `[secondary]` (2026-01): Nexus + ConnectX bring-up scripts; claims 96% packet-drop reduction and 2× bandwidth improvement — self-reported by the repo author, NOT independently verified; treat as `[unverified]` for the performance numbers. https://github.com/enizaksoy/rdma-ai-cluster

### 6.3 Why AI fabrics need it
- AI/ML training (especially LLM all-to-all collectives) generates synchronized microbursts that overflow shallow buffers; RDMA has no TCP-style retransmission safety net at line rate, so drops collapse throughput `[secondary]` — intelligentvisibility; Kar.
- Strategic note (vendor-adjacent, 2026): MLPerf testing found the Ethernet-vs-InfiniBand gap statistically insignificant and Meta runs 24,576-GPU Ethernet fabrics; InfiniBand retained only for narrow extreme-synchronization niches `[vendor-reported]` — intelligentvisibility (pro-Ethernet vendor voice; InfiniBand camp would dispute the framing). Flagged as perspective, not consensus.
- Arista positioning `[vendor-reported]`: RoCEv2 across AI-ready EOS portfolio + deep-buffered ASICs (Jericho family) absorbing microbursts + Cluster Load Balancing for even load `[vendor-reported]` — intelligentvisibility citing Arista.

### 6.4 Vendor QoS model comparison (structural, not numeric)
- Cisco Nexus (NX-OS): MQC model — `class-map type qos` (classification) + `class-map type queuing` + `class-map type network-qos` (PFC/MTU per class) + `policy-map type ...`; system-defined classes; PFC enable per network-qos class `[official]` — Cisco NX-OS QoS guides (not deep-linked here; CLI nouns verified against Dell/Cisco docs above).
- Arista EOS: traffic-class mapping, per-TC ECN/WRED thresholds, PFC per 802.1p priority; AI fabrics commonly put RoCE on TC 3 `[vendor-reported]/[secondary]` — NetPilot pattern; Arista TOI docs.
- Dell OS10: `class-map type queuing`, `policy-map type queuing`, `qos-map traffic-class`, WRED+ECN profiles — full verbatim example in §6.2 `[official]`.
- NVIDIA Spectrum (Cumulus/ONYX or SONiC): shared-buffer architecture, lossless buffer pools, `mlnx_qos`/`ztc` RoCE profiles, DCBX application TLVs `[vendor-reported]` — FS.com whitepaper dump.
- Gap: side-by-side default queue counts, buffer depths, and PFC-watchdog defaults per 2026 ASIC (Spectrum-4/5, Tomahawk 5/6, Jericho3-AI) not compiled here — open item `[unverified]`.

### Wave 6 verification
- Sources: 8. Self-reported performance claims downgraded to `[unverified]`. Vendor-strategy claims labeled as perspective. Config snippets transcribed verbatim from dated docs.

---

## Wave 7 — Multicast: PIM-SM/SSM, RP design, MSDP, Anycast-RP

### 7.1 PIM-SM vs PIM-SSM
- PIM-SM (Sparse Mode): receivers join a shared tree (*,G) rooted at a Rendezvous Point; sources register to the RP; receivers then switch to source trees (S,G) for optimal paths. The single-RP-per-group constraint is PIM-SM's central scaling weakness `[official]` — RFC 3446 (Kim et al., Jan 2003, informational). https://www.rfc-editor.org/rfc/rfc3446
- PIM-SSM (Source-Specific Multicast): receivers subscribe to (S,G) directly via IGMPv3/MLDv2 — no RP, no shared tree, no MSDP; used where the source is known (content distribution, financial feeds) `[secondary]` — standard references; Cisco TRM doc notes TRM supports PIM-SSM and PIM-SM in the overlay `[official]` (Wave 8).
- PIM-BiDir: shared tree only, no (S,G) state — scales group state but with RP-through-traffic trade-offs; Cisco TRM explicitly does NOT support PIM-BiDir in underlay or overlay `[official]` — Cisco TRM whitepaper (Wave 8).

### 7.2 RP placement
- Single-RP problems at scale: traffic concentration/register-decap load on one box, slow convergence on RP failure, suboptimal forwarding, distant-RP dependencies — demonstrated in continental/inter-continental deployments `[official]` — RFC 3446.
- Design choices: on-path RP (RP on a device already in the multicast path — cheaper, but an RP node failure forces unicast RIB recomputation and RPF re-evaluation for every (S,G)) vs off-path RP (dedicated RPs out of the forwarding path — costs hardware, isolates failure domain) `[secondary]` — Cisco Community accepted solution (2026-05). https://community.cisco.com/t5/routing-and-sd-wan/multicast-rp-placement-msdp-vs-anycast/td-p/5551629
- RP discovery: static RP mapping (simplest, recommended with anycast-RP per Juniper), Auto-RP (Cisco), BSR (PIMv2 bootstrap) — Juniper recommends static over BSR/auto-RP when combined with anycast-RP for simplicity `[official]` — Juniper anycast-RP example. https://www.juniper.net/documentation/us/en/software/junos/multicast/topics/topic-map/mcast-pim-anycast-rp.html

### 7.3 Anycast-RP (RFC 3446)
- Mechanism: all RPs in the set share one anycast IP (typically a loopback), advertise the same group→RP mapping; sources/receivers join/register toward the topologically closest RP via unicast routing; RPs MSDP-peer with each other (using unique addresses — never the anycast address as router-ID) and exchange Source-Active (SA) messages so every RP knows all sources `[official]` — RFC 3446; Huawei anycast-RP docs; Juniper example. https://info.support.huawei.com/hedex/api/pages/EDOC1100149308/AEJ0713J/18/resources/admin/sec_admin_multicast_0167.html
- Benefits: RP load sharing, fast failover (unicast routing reconverges to the next-closest RP), optimal RP-proximity paths `[official]` — Huawei; RFC 3446.
- Platform nuance (conflict recorded): Cisco Community (2026): on IOS-XE, IPv4 anycast-RP requires MSDP (or MP-BGP IPv4-multicast AF); Nexus supports an MSDP-less PIM anycast-RP variant; Junos supports anycast-RP with or without MSDP (PIM-only for IPv4 and IPv6) `[secondary]/[official]` — Cisco Community thread; Juniper docs. Do NOT assume MSDP-less anycast-RP works on all platforms — verify per OS. https://community.cisco.com/t5/routing-and-sd-wan/multicast-rp-placement-msdp-vs-anycast/td-p/5551629
- Huawei implementation notes: with >3 RPs, full-mesh MSDP with all peers in one mesh-group; must configure a logical RP address for SA messages (peers discard SAs whose RP address equals their own) `[official]` — Huawei docs.

### 7.4 MSDP
- Multicast Source Discovery Protocol: TCP-based SA exchange between RPs in different PIM-SM domains (or anycast-RP sets); includes RPF checks on SA messages to prevent loops `[official]` — RFC 3446; Huawei.
- Mesh-groups suppress SA flooding between fully-meshed peers — required config hygiene at scale `[official]` — Huawei.

### Wave 7 verification
- Sources: 7 (RFC 3446, Huawei ×2, Juniper, Cisco Community, Huawei config guide). One platform conflict documented explicitly (MSDP-less anycast-RP support).

---

