---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-17-multicast-in-practice-market-data-trm-checklist-trou
title: "Wave 17 — Multicast in practice: market data, TRM checklist, troubleshooting"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["asic", "cost", "latency", "nvidia", "voice"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [486, 534]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 904cb151b9267b1f722402ab199ab1917b03b321f49b116f71862079d2b001bb
---

# Wave 17 — Multicast in practice: market data, TRM checklist, troubleshooting

## Wave 17 — Multicast in practice: market data, TRM checklist, troubleshooting

### 17.1 Financial market-data multicast (the canonical DC multicast workload)
- Exchanges distribute price feeds as UDP multicast; consumers use PIM-SSM (source known) with IGMPv3 joins; latency sensitivity makes RP-free SSM and static joins preferable to dynamic discovery `[secondary]` — industry practice; PIM-SSM mechanics in Wave 7.
- Requirements: deterministic low-jitter paths, no RPF failures during reconvergence, redundant feeds on diverse paths with fast failover — anycast-RP + MSDP mesh or dual independent SSM trees `[secondary]`.
- PTP often coexists on the same fabric for MiFID II timestamping — Wave 9/15.

### 17.2 TRM deployment checklist (Nexus 9000-based, per Cisco)
1. Underlay: IPv4 PIM-SM everywhere VTEPs peer; RP reachable; do NOT use SSM/BiDir/IR in the underlay `[official]` — Cisco TRM whitepaper (Wave 8).
2. One BUM mode fabric-wide; if TRM is in the future, deploy underlay multicast from day one `[official]` — same.
3. Overlay: tenant VRF with PIM-SM or SSM; overlay RP reachable INSIDE the tenant VRF; source TTL high enough to survive FHR/LHR `[official]` — same.
4. 224.0.0.0/24 stays bridged, never routed `[official]` — same.
5. Verify: `show ip pim rp`, `show ip mroute vrf <tenant>`, EVPN Type-6/7 routes for selective forwarding `[secondary]` — operator CLI practice (command names per NX-OS; verify per version).

### 17.3 Multicast troubleshooting primitives
- `show ip mroute` / `show ip mroute vrf`: (*,G) vs (S,G) state, incoming interface, OIL (outgoing interface list), RPF neighbor `[secondary]` — standard CLI.
- RPF failures: check unicast table toward source/RP in the CORRECT VRF; most "multicast broken after VRF change" incidents are RPF in the wrong table `[secondary]` — operator practice.
- No-receiver flooding: check IGMP snooping querier election and EVPN Type-6/7 propagation `[secondary]` — Wave 8.
- Duplicate packets: overlapping RP domains, two DFWs forwarding, or IR+multicast-underlay mix (see Wave 8 "don't mix") `[secondary]`.

### Wave 17 verification
- Sources: 5. Checklist transcribed from Cisco's TRM whitepaper; market-data section is industry practice, flagged as such.

---

## Wave 18 — QoS in practice: traffic-class plan, buffer math, validation

### 18.1 Reference traffic-class plan (8-class DC model)
- A widely used 8-class plan: TC7 network control (CS6/CS7, strict priority, tiny bandwidth), TC6 management/PTP, TC5 voice/signaling (EF), TC4 real-time video, TC3 lossless RDMA/storage (PFC+ECN), TC2 transactional data (AF31), TC1 bulk data, TC0 best effort / scavenger `[secondary]` — synthesized from Cisco/Dell/Arista QoS guides and the Wave 6 RoCE convention; treat as a template, not a standard.
- Rules: exactly ONE lossless class per direction unless the platform is explicitly validated for more; never put PFC on the control class; scavenger (less-than-best-effort) protects everything else from bulk floods `[secondary]` — operator practice.
- Marking map example: RoCE → DSCP 26 or 48 (operator choice — document it), CNP → DSCP 46, storage (iSCSI/NVMe-TCP) → AF31 or dedicated class, control → CS6 `[secondary]` — conventions observed across vendor guides; no universal standard exists, which is itself a finding.

### 18.2 Buffer math (conceptual)
- Bandwidth-delay product sets the minimum buffer to absorb an RTT of line-rate burst; microburst absorption needs headroom beyond the average — deep-buffer ASICs (Jericho family per Arista) trade buffer for burst tolerance at higher cost/power `[vendor-reported]/[secondary]` — Arista positioning (Wave 6).
- PFC threshold sizing: XOFF (pause) threshold must leave enough buffer for in-flight frames during the pause round-trip; XON (resume) needs hysteresis to avoid pause flapping — vendor calculators/tools exist per ASIC; hand-tuning without them is guesswork `[secondary]` — operator practice.
- Shared-buffer architectures (NVIDIA Spectrum) dynamically allocate from a common pool with per-class minimum guarantees; fixed-carve designs partition statically — the operational difference is how bursts in one class affect others `[vendor-reported]/[secondary]`.

### 18.3 Validation playbook
1. Baseline: `show` PFC pause counters per priority per interface — zero under normal load is the target; rising counters locate the congestion point `[secondary]`.
2. ECN efficacy: CNP counters on NICs + ECN-mark counters on switches; if PFC fires but ECN never marks, thresholds are inverted (fix per Wave 12 §12.4) `[secondary]` — Kar lab method (Wave 6).
3. Loss test: RoCE traffic at line rate with incast (many-to-one); measure drops, PFC pause duration, and tail latency before/after tuning `[secondary]` — Kar lab.
4. Storm-control/CoPP audit (Wave 9): DHCP/ARP/IGMP punted traffic must be policed so control-plane floods cannot starve BGP or pause the fabric `[secondary]`.
5. Change discipline: QoS changes are fabric-wide fate-sharing — stage on one leaf pair, measure, then roll; keep the previous config archived `[secondary]` — operator practice.

### Wave 18 verification
- Sources: 5. Template presented as template; buffer math conceptual (no per-ASIC numbers claimed).

---

