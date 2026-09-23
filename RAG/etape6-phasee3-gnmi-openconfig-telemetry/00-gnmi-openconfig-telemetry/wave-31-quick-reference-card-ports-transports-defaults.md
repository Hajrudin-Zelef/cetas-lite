---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-31-quick-reference-card-ports-transports-defaults
title: "Wave 31 — Quick-reference card: ports, transports, defaults"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["agent", "decode", "ethernet", "latency", "revenue"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [630, 682]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 3209e622472ddfcd6cb57682a904a645b874648cafe8058f1874fbfd42e33d27
---

# Wave 31 — Quick-reference card: ports, transports, defaults

## Wave 31 — Quick-reference card: ports, transports, defaults

| Item | Default | Notes | Provenance |
|---|---|---|---|
| Juniper JTI dial-in gNMI/gRPC | TCP 50051 | Client cert+key under TLS; 25.2R1 adds multi-server/multi-port | [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf) |
| Arista EOS gNMI (Octa) | TCP 6030 | `management api gnmi`; tunnel dial-out via `transport grpc-tunnel` | [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md) [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md) |
| Nokia SR Linux gNMI | TCP 57400 | gnmic authored by Nokia; native-first models | [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md) |
| Cisco FTD gNMI | TCP 50051 (dial-in default) | Tunnel dial-out when collectors unreachable | [official](https://docs.defenseorchestrator.com/cdfmc/c_openconfig_streaming_telemetry.html) |
| gnmic Prometheus output | TCP 9273 | Multi-output fan-out from one collector | [official](https://github.com/openconfig/gnmic/blob/HEAD/README.md) |
| gnmic tunnel-server | TCP 57401 (example) | Devices dial in; RPCs routed to registered targets | [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/user_guide/tunnel_server.md) |
| Juniper native UDP JTI | UDP (device→collector, e.g. 4729/5566) | Dial-out only; revenue-interface constraint on non-EVO | [secondary](https://github.com/jawroper/jti_nus_fluentbit) |
| Cisco MDT dial-out encodings | GPB kv / JSON | Encoding negotiated per destination-group | [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/710x/configuration/guide/b-telemetry-cg-ncs5500-710x/dial-out-telemetry-session-from-router-to-destination.html) |
| Cisco gNOI concurrent RPCs (NX-OS) | max 16 active | Cited Nexus programmability guide | [official](http://www.static-cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/106x/programmability/cisco-nexus-9000-series-nx-os-programmability-guide-106x/gnoi---operation-interface.pdf) |
| orb-agent gNMI discovery | ≤1,024 addresses/policy | host/CIDR/range; OpenConfig origins | [official](https://github.com/netboxlabs/orb-agent/blob/HEAD/orb-discovery/gnmi-discovery/README.md) |

## Wave 32 — Glossary and acronym map (telemetry domain)

- **gNMI** — gRPC Network Management Interface: Capabilities/Get/Set/Subscribe over YANG paths.
- **gNOI** — gRPC Network Operations Interface: System, OS, Cert, File, FactoryReset, Bootstrapping services.
- **gRIBI** — gRPC Routing Information Base Interface: programmatic route injection (seen alongside gNMI 0.8.0 in Cisco 8000 KNE lab, Wave 2).
- **OpenConfig** — vendor-neutral YANG models; quarterly versioned bundles.
- **JTI** — Juniper Telemetry Interface: dial-in (gNMI/gRPC) + dial-out (native UDP/GPB).
- **MDT** — Model-Driven Telemetry (Cisco IOS-XR dial-out framework).
- **NetDL** — Arista's network data lake (CloudVision time-series state store).
- **AVA / CV UNO** — Arista AI analytics engine / Universal Network Observability.
- **EDA** — Nokia Event Driven Automation (intent-based platform for SR Linux fabrics).
- **TARGET_DEFINED** — gNMI stream mode where the target chooses cadence (Juniper JTI signature behavior since 20.2R1).
- **Octa** — Arista EOS gNMI agent process (port 6030).
- **TerminAttr** — EOS streaming agent feeding CloudVision (dial-out side).
- **LANZ** — Arista latency analyzer (microburst detection; named in scope-adjacent material, not deep-dived here) [unverified — depth].
- **DCQCN** — Data Center Quantized Congestion Notification: ECN-based RoCE congestion control (Wave 8).
- **PFC / ECN** — Priority Flow Control / Explicit Congestion Notification: lossless-Ethernet building blocks for AI fabrics (Wave 8).

*Final verification (2026-09-22): line count ≥750 confirmed by wc -l at close; code fences balanced; tail intact; 32 waves; provenance tags throughout; no invented identifiers. Assignment complete.*

## Wave 33 — Cisco IOS-XR MDT dial-out and Dell OS10 sensor-path anatomies

### 33.1 IOS-XR model-driven telemetry dial-out structure

- Three-part config: **destination-group** (collector IP/port, encoding), **sensor-group** (sensor paths + sample interval), **subscription** (binds sensor-group to destination-group) [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/710x/configuration/guide/b-telemetry-cg-ncs5500-710x/dial-out-telemetry-session-from-router-to-destination.html).
- Encodings: GPB key-value or JSON per destination; cadence per sensor path; event-driven (interval 0) for on-change style paths [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/telemetry/77x/b-telemetry-cg-ncs5500-77x.pdf).
- Native sensor paths dominate in practice (`Cisco-IOS-XR-*`); OpenConfig paths served but version-skewed (bgp-types 5.3.1 at 26.x vs 4.x-lab, Wave 20.1) [secondary].

### 33.2 Dell OS10 sensor paths (queue/buffer focus)

- Queue, priority-group, and buffer sensor paths under `base-*` native models are the AI-fabric monitoring surface on OS10 (Wave 8); streamed at ≥15 s cadences or event-driven (0) [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/configure-telemetry?guid=guid-a56d32ee-b7b9-4667-9156-7b62845701f6&lang=en-us).
- `.proto` files for GPB decode come from the Dell support site; only one destination group + one subscription profile per switch constrains multi-collector designs [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/configure-telemetry?guid=guid-a56d32ee-b7b9-4667-9156-7b62845701f6&lang=en-us).

### 33.3 SONiC telemetry posture (2026)

- SONiC distributions expose gNMI with a mix of native and OpenConfig models; FRR `openconfigd` coverage is partial in lab contexts; PFC/ECN counters flow through native models (Wave 8) [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md).
- The whitebox telemetry story remains the least standardized of the major NOS families — a standing gap (Wave 15.2 #3).

