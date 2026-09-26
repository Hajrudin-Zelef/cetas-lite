---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/3-5-arista-eos-ceos-and-hardware
title: "3.5 Arista EOS (cEOS and hardware)"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: hardware
actors: ["Nvidia"]
dates: []
keywords: ["agent", "distribution", "nvidia"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [78, 96]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 6b5b5b9b9e18ebe12dcc9869fda51ee54cbb44fcf9fbb12bbb99a5ee6a7299cb
---

# 3.5 Arista EOS (cEOS and hardware)

- Streaming telemetry (release 10.5.6, 2026 train): `telemetry` → `enable` → `destination-group` (one IPv4/IPv6 dest + port) → `subscription-profile` binding pre-configured sensor groups with per-group cadence in ms (`bgp 300000`, `bgp-peer 0`, `buffer 15000`, `device 300000`, `environment 300000`, `interface 180000`, `lag 0`, `system 300000`) → `encoding gpb`, `transport grpc no-tls` [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/example-configure-streaming-telemetry?guid=guid-90cd942e-bd77-4c20-a2e6-73aa0d4f8d7e&lang=en-us).
- Hard limits documented: **only one destination group, only one destination address in the group, only one subscription profile** — a scale ceiling for dial-out fan-out [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/configure-telemetry?guid=guid-a56d32ee-b7b9-4667-9156-7b62845701f6&lang=en-us).
- Sampling rate **0 = near real-time event-driven** collection; minimum recommended intervals 15000 ms for most groups, 0 (event) for BGP-peer and LAG [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/configure-telemetry?guid=guid-a56d32ee-b7b9-4667-9156-7b62845701f6&lang=en-us).
- Sensor paths are Dell-native (e.g. `base-qos/queue-stat`, `base-pas/chassis`, `infra-bgp/peer-state/peer-status`); collector authors must download Dell's telemetry `.proto` files from the support site [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/example-configure-streaming-telemetry?guid=guid-90cd942e-bd77-4c20-a2e6-73aa0d4f8d7e&lang=en-us).
- gNMI agent: activated only in **SmartFabric director mode** (`switch-operating-mode sfd`, requires reload); TLS + username/password; only Super Admin roles accepted (all others rejected); SmartFabric director issues gNMI **Set** to configure the telemetry agent to stream in **OpenConfig format** [official](https://www.dell.com/support/manuals/en-ie/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-2-6/gnmi-agent?guid=guid-9d4c66d3-178d-4122-baed-6914977fe93e&lang=en-us).

### 3.5 Arista EOS (cEOS and hardware)

- gNMI served by the **Octa** agent on port **6030**; CLI: `management api gnmi` → `transport grpc default` → `no shutdown` [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md).
- Interop data point: cEOS-Lab image **4.33.1F** accepted the `management api gnmi` block but rejected `no shutdown` (`% Incomplete command`), so the Octa process never bound `:6030` and gnmic dial-in failed — a lab-image quirk, not a platform verdict [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md).
- OpenConfig counter path on cEOS: `/interfaces/interface[name=Ethernet1]/state/counters` [secondary](https://github.com/martimy/model-driven-configuration-tutorial/blob/HEAD/docs/tasks/12-stream-telemetry.md).
- gnmic `capabilities` against an Arista target shows extensive Arista deviation models (`arista-bfd-deviations`, `arista-exp-eos-*`) alongside OpenConfig modules (`openconfig-bgp-evpn`, `openconfig-lldp-types` 0.1.1, `openconfig-isis-lsdb-types` 0.4.2, `openconfig-segment-routing` 0.0.4) [secondary — capabilities dump](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).
- **Gap:** TerminAttr (Arista's CloudVision streaming agent) dial-out specifics and CloudVision 2026 telemetry evolution were not covered in this wave; record as open item.

### 3.6 SONiC / whitebox and FRR (status note)

- SONiC exposes gNMI via its management framework; a 2026 AI-lab note flags **FRR-based nodes (vtysh JSON)** as still polling-based (15 s docker-exec) in that lab's migration roadmap — i.e., FRR `openconfigd` gNMI readiness is partial in lab contexts [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md).
- **Gap:** per-vendor SONiC distribution (Dell Enterprise SONiC, NVIDIA Cumulus, community) gNMI/OpenConfig maturity in 2026 needs a dedicated pass.

