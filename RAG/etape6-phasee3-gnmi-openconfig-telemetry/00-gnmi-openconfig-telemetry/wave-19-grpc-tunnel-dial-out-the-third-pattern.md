---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-19-grpc-tunnel-dial-out-the-third-pattern
title: "Wave 19 — gRPC tunnel dial-out: the third pattern"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: ["Huawei"]
dates: ["2026-01-14", "2026-09-01"]
keywords: ["agent", "coherent optics", "dci", "ethernet", "optics"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [433, 478]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 3688560aaa65c54ccc682ac9bc0ca3dcde635b245d1dec494dfb60dee50bc907
---

# Wave 19 — gRPC tunnel dial-out: the third pattern

## Wave 19 — gRPC tunnel dial-out: the third pattern

### 19.1 Why tunnels exist

- Classic dial-in fails when the collector cannot reach the device (NAT, cloud-hosted collector, untrusted network); plain dial-out puts subscription control on the device. The **gRPC tunnel** (openconfig/grpctunnel) splits the difference: the device dials out to establish a tunnel, then the collector issues gNMI requests *through* it — collector-driven subscriptions over device-initiated transport [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md) [official](https://docs.defenseorchestrator.com/cdfmc/c_openconfig_streaming_telemetry.html).
- Cisco's framing: DIAL-IN is ideal on trusted networks where collectors are trusted; DIAL-OUT (tunnel) is ideal when collectors are cloud-hosted or outside the trusted network. Both modes use TLS; dial-out adds tunnel-infrastructure keys [official](https://docs.defenseorchestrator.com/cdfmc/c_openconfig_streaming_telemetry.html).

### 19.2 Arista: gNMI dial-out via gRPC tunnel

- Under `management api gnmi`, configure `transport grpc-tunnel <name>`:
  - `vrf MGMT` (optional source VRF), `destination 192.185.128.100 port 30000` (required — the collector's tunnel server), `local interface Management1 port 50000` (optional source), `target spine1` (**required** user-defined ID used in tunnel establishment), `provider eos-native` [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md).
- Sequence: tunnel client on switch dials out → collector's tunnel server accepts → secure gRPC tunnel established → collector sends gNMI requests via tunnel → device responds via the same tunnel [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md).
- Debug: `show agent gnmireverse logs` [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md).

### 19.3 Cisco: grpctunnel dial-out

- `gnxi grpctunnel target GNMI_GNOI` (only GNMI_GNOI supported; GNMI_GNOI_INSECURE is test-only) and `gnxi grpctunnel destination <name>` with `address`, `port`, `identity trustpoint`, `source-address`, `source-vrf` [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1717/b_1717_programmability_cg/gnmi-dial-out-using-the-grpc-tunnel-service.pdf).

### 19.4 gnmic tunnel-server

- gnmic runs a tunnel server the devices dial into; combined with `gnmi-server` mode, Get/Subscribe RPCs against gnmic are routed to the relevant registered targets (`--target` flag or all) [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/user_guide/tunnel_server.md).
- Knobs: `target-wait-time: 2s`, `client-auth`, `enable-metrics` (Prometheus gRPC server metrics), `debug` [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/user_guide/tunnel_server.md).
- Kubernetes operator: `TunnelTargetPolicy` resources select tunnel-mode targets by label; the cluster must have `grpcTunnel` configured or the pipeline errors [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/user-guide/pipeline.md).

## Wave 20 — OpenConfig model-structure walkthroughs (operator's map)

### 20.1 `openconfig-bgp`

- Structure: `bgp/neighbors/neighbor[neighbor-address]/afi-safis/afi-safi[afi-safi-name]/` for per-AF state, `bgp/peer-groups`, `bgp/global` [secondary].
- 2026 field reality: `bgp-types` 5.3.1 in current IOS-XR (4.x line on Cisco 8000E lab), 3.2.0-era on Junos 25.2, **4.1.0 on Huawei** (gnmic issue #451) — the same model family at three different generations across vendors [secondary](https://github.com/openconfig/gnmic/issues/451).
- EVPN rides the BGP model via `l2vpn-evpn` AFI-SAFI — the primary OpenConfig telemetry surface for EVPN fabrics (E1/E4 cross-ref) [secondary].

### 20.2 `openconfig-interfaces`

- `interfaces/interface[name]/` with `config`, `state`, `subinterfaces`, `hold-time`; operational counters under `state/counters` [secondary].
- Junos 25.2R1 ACX7509 ships `openconfig-if-ethernet` 2.6.2 as a sensor path — interfaces are the most universally supported OC telemetry surface [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html).

### 20.3 `openconfig-platform` and `openconfig-terminal-device`

- Platform model carries component inventory (chassis/linecard/port/optics), power states (updated 2026-09-01), and is the LLDP-adjacency reconciliation anchor for source-of-truth sync (Wave 7.3) [secondary](https://github.com/openconfig/public/commit/af2b0f11def9fb0cad7f4a04d28d6d5bef2b7da8).
- Terminal-device 1.12.0 (2026-01-14) covers coherent optics — the DWDM/ZR+ telemetry surface for AI-fabric interconnects and DCI (cross-ref E3 optics-adjacent scope) [official](https://github.com/openconfig/public/blob/master/release/models/optical-transport/openconfig-terminal-device.yang).

### 20.4 `openconfig-qos`

- Queue/counter paths used in featureprofiles DP-1.4; OpenConfig QoS → Junos mapping published (Wave 10.2); PFC/ECN-relevant queue stats are the AI-fabric monitoring surface (Wave 8) [secondary](https://github-wiki-see.page/m/openconfig/featureprofiles/wiki/qos_output_queue_counters_test) [official](https://www.juniper.net/documentation/us/en/software/junos/open-config/topics/concept/open-config-qos-mapping.html).

