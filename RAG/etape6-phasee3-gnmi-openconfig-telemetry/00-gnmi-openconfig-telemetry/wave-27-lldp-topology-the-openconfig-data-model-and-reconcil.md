---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-27-lldp-topology-the-openconfig-data-model-and-reconcil
title: "Wave 27 — LLDP topology: the OpenConfig data model and reconciliation loop"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["agent", "optics"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [577, 629]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 77640a1e9b60a7ada03fc17653fccab7ee9a530cde70e9b892f774eb4377b708
---

# Wave 27 — LLDP topology: the OpenConfig data model and reconciliation loop

## Wave 27 — LLDP topology: the OpenConfig data model and reconciliation loop

### 27.1 Model surface

- `openconfig-lldp` exposes `lldp/interfaces/interface[name]/neighbors/neighbor[id]/state` (chassis-id, port-id, system-name, capabilities) — the vendor-neutral adjacency feed [secondary].
- Nokia `srl_nokia-interfaces` and Arista/Cisco equivalents carry LLDP TLV state in native models; the orb-agent gNMI discovery layer normalizes discovery results into profiles with OpenConfig origins (Wave 7.2) [official](https://github.com/netboxlabs/orb-agent/blob/HEAD/orb-discovery/gnmi-discovery/README.md).

### 27.2 Reconciliation loop (practice pattern)

1. Subscribe on-change to LLDP neighbor state per device → adjacency events.
2. Reconcile against NetBox/Nautobot cable objects (Wave 7.3): create missing, flag stale, quarantine mismatches.
3. Feed confirmed topology into NetBox for the E1 file's source-of-truth role and into Batfish for analysis (E4) [analysis].
- Event-driven LLDP reconciliation replaces periodic SNMP-walk discovery with sub-second topology truth — the mechanism behind "real-time topology" in the scope [analysis].

## Wave 28 — NETCONF vs RESTCONF vs gNMI: operation-by-operation

| Concern | NETCONF (RFC 6241) | RESTCONF (RFC 8040) | gNMI |
|---|---|---|---|
| Transport | SSH (830) | HTTPS | gRPC/HTTP2 (TLS) |
| Data modeling | YANG | YANG | YANG (OpenConfig/native) |
| Config ops | edit-config, candidate/commit, confirmed-commit, rollback | PATCH/PUT/POST/DELETE on datastores | Set (replace/update/delete) |
| State read | get/get-config | GET | Get |
| Streaming | Notifications (RFC 8639/8641) | SSE / subscribed notifications | Subscribe (SAMPLE/ON_CHANGE/TARGET_DEFINED/ONCE/POLL) |
| Device ops | RPCs | Actions | gNOI services |
| 2026 network role | Transactional config, especially multi-vendor push (Wave 17.2) | Controller NB APIs, lightweight tooling | Streaming telemetry + emerging config |
| Strength | Transactions, confirmed commit | HTTP familiarity, tooling | Sub-second streams, single schema w/ OC |

- Cisco NX-OS/IOS-XE ship gNOI + gNMI + NETCONF side by side; the operator chooses per plane (Wave 6) [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf).
- YANG Suite (Cisco) remains the 2026 model-exploration GUI; 2026 release version unverified (open item, Wave 15.2 #9).

## Wave 29 — Prometheus metric patterns for network telemetry

### 29.1 Naming and labels

- gnmic Prometheus output maps gNMI path elements to metric names and keys to labels; device identity (`target`) is the primary label — cardinality scales with targets × paths × label values (Wave 5.2) [official](https://github.com/openconfig/gnmic/blob/HEAD/README.md).
- High-cardinality traps: per-lane optics, per-queue counters, per-neighbor BGP RIB entries — downsample or route to ClickHouse (Wave 5.2) [analysis].

### 29.2 Query patterns (operator's sketch)

- Interface error rate: `rate(if_in_errors[5m])` per target/interface — the canonical first dashboard panel in the SR Linux reference stack (Wave 3) [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md).
- Absence-of-data alerting on heartbeat metrics distinguishes dead streams from quiet ones (Wave 22.3) [analysis].
- BGP session state transitions are better consumed as on-change events (Wave 11.1) than scraped gauges [analysis].

## Wave 30 — Closing supplement: what changed since Wave 15's audit

- Waves 16–18: Arista CloudVision/NetDL/AVA/CV UNO/AI-agent observability; gNMI Set semantics; deployment playbook with hardening checklist and alerting patterns.
- Waves 19–21: gRPC tunnel dial-out (Arista/Cisco/gnmic); OpenConfig model-structure map (bgp/interfaces/platform/terminal-device/qos); Nokia EDA positioning.
- Waves 22–24: troubleshooting playbook; OpenConfig 2026 process and adoption-lag analysis; trajectory outlook.
- Waves 25–26: Juniper JTI dial-out config anatomy (25.2R1 multi-gRPC-server enhancement, protocol-stack sensors, MACsec path); Telegraf dual-plugin and Fluent Bit/Synse native-UDP collectors.
- Waves 27–29: LLDP reconciliation loop; NETCONF/RESTCONF/gNMI operation table; Prometheus patterns.

*File complete per assignment: English, current through 2026-09-22, ≥750-line target to be verified at close, sole writer, append-only, provenance-tagged, no invented identifiers.*

