---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-6-netconf-restconf-vs-gnmi-in-2026-yang-suite
title: "Wave 6 — NETCONF/RESTCONF vs gNMI in 2026; YANG Suite"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [146, 193]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 5f5d988263f60ef22ce1aec17b8d900a78d125fbe699699074eba8cc4c552668
---

# Wave 6 — NETCONF/RESTCONF vs gNMI in 2026; YANG Suite

## Wave 6 — NETCONF/RESTCONF vs gNMI in 2026; YANG Suite

### 6.1 gNOI — the operational companion to gNMI (2026)

- gNOI = gRPC Network Operations Interface: gRPC microservices for operational commands (not config). Cisco's formulation: "the gNMI service defines operations for configuration management, operational state retrieval, and bulk data collection through streaming telemetry. gNOI only allows the adoption of services that a device supports" [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf).
- Standard services/RPCs:
  - **System**: Ping, Traceroute, Time, SwitchControlProcessor, Reboot, RebootStatus, CancelReboot, SetPackage [official](https://www.cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/102x/programmability/cisco-nexus-9000-series-nx-os-programmability-guide-release-102x/m-gnoi-grpc-network-operations-interface.html).
  - **OS**: Activate, Verify (image lifecycle) [official](https://www.cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/102x/programmability/cisco-nexus-9000-series-nx-os-programmability-guide-release-102x/m-gnoi-grpc-network-operations-interface.html).
  - **Cert**: Install, Rotate, LoadCertificate [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf).
  - **File**: Get, Put, Stat, Remove, TransferToRemote [official](http://www.static-cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/106x/programmability/cisco-nexus-9000-series-nx-os-programmability-guide-106x/gnoi---operation-interface.pdf).
  - **FactoryReset**: Start; **Bootstrapping** service [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf).
- Cisco NX-OS specifics: gNOI configuration identical to gNMI config; max **16 active gNOI RPCs**; N9K runs one endpoint with one gNMI service + two gNOI microservices; Nexus container services (ListImage/RemoveImage/ListContainer/StartContainer/StopContainer/UpdateContainer/Log/CreateVolume/RemoveVolume/ListVolume/StartPlugin/StopPlugin/ListPlugin/RemovePlugin) also exposed [official](http://www.static-cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/106x/programmability/cisco-nexus-9000-series-nx-os-programmability-guide-106x/gnoi---operation-interface.pdf).
- Cisco IOS-XE: gNOI covered in the programmability configuration guide at release **17.18** (2026 train); user authentication via `gnxi securepasswordauth`, or OpenConfig `openconfig-system-management.yang` [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf).
- IOS-XR KNE lab: `gnoic system ping` verified against 8000e — System service live on production code [secondary](https://github.com/openconfig/kne/blob/HEAD/examples/cisco/8000e/README.md).
- Extreme Networks **ONE OS Switching v22.2.0.0** ships a dedicated gNOI Reference Guide (2026): OpenConfig-compliant `file.proto` and `os.proto` plus Extreme extensions — `auth.proto` (JWT auth, cert generation), `cfgmgmt.proto` (default/running/file config states), `firmware.proto` (commit/rollback/uninstall), network-protocol services (`acl/bfd/bgp` incl. L2VPN-EVPN/MAC-VRF route clearing, `lacp/mlag/qos/vlan`), `iah.proto` (VM control), `programmable.proto` (component firmware), `ssh.proto`, `systemUtils.proto` (reboot), `techsupport.proto` [official](https://documentation.extremenetworks.com/Extreme%20ONE%20OS%20Switching%20v22.2.0.0%20GNOI%20Reference%20Guide/Extreme_ONE_OS_Switching_22_2_0_0_GNOI_Reference_Guide.pdf).
- **Note:** the Extreme gNOI guide URL above was returned verbatim by search; treat as cited [official].

### 6.2 NETCONF/RESTCONF status vs gNMI in 2026

- Division of labor as of 2026: NETCONF (SSH/830, transactional candidate/commit/rollback) remains the config-transaction plane; gNMI (gRPC, efficient binary) dominates streaming telemetry; RESTCONF/JSON sits between for ad-hoc automation [secondary — consensus across 2026 programmability docs].
- Evidence both planes coexist on current NOS: Nokia SR Linux and Arista cEOS both serve NETCONF (port 830) and gNMI (57400 / 6030) natively; OpenConfig interface description pushed via NETCONF, counters streamed via gNMI in the same lab [secondary](https://github.com/netpilot-labs/example-prompts/blob/HEAD/advanced/netconf-gnmi-openconfig.md).
- **Gap:** a current quantitative survey of NETCONF vs gNMI deployment share (2026) was not found; record as open item.

### 6.3 YANG Suite

- Cisco YANG Suite: GUI/API toolkit for YANG model exploration, NETCONF/RESTCONF/gNMI client operations against devices [vendor-reported — Cisco developer tooling; version currency not verified this wave].
- **Gap:** YANG Suite 2026 release/version and feature status not verified this wave.

## Wave 7 — Real-time topology tools and NetBox/Nautobot integration

### 7.1 gnmic operator ↔ NetBox (2026)

- The gNMIc Kubernetes operator ships documented NetBox target-discovery flows: an HTTP `TargetSource` pulls `/api/dcim/devices/`, transforms records into gNMIc `Target` resources via **CEL expressions**, and creates/updates them in-cluster (sync example: every 30 min) [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/examples/NetBox/REST%20API/_index.md).
- Alternate path: NetBox **export templates** (`?export=gNMIc%20Device%20Export`) feeding the same `TargetSource`; labels like `inventory: netbox` mark discovered targets [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/examples/NetBox/Export%20Template/_index.md).
- A **webhook** flow lets NetBox push device events to the operator API (bearer-token + signature secrets) [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/examples/NetBox/webhook/_index.md).
- Credentials handled via Kubernetes Secrets (`TargetProfile` + `credentialsRef`); NetBox API token stored as secret, never in manifests [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/examples/NetBox/REST%20API/_index.md).

### 7.2 NetBox Labs orb-agent gNMI discovery

- `netboxlabs/orb-agent` includes a `gnmi_discovery` backend: targets as host / CIDR / range (`10.0.0.0/24`, `10.1.0.0-50`); CIDRs/ranges are probed and only answering addresses subscribed (max **1024 addresses per policy**); per-target `mode` (`auto`/`on_change`/`sample`/`get`), `profile` (e.g. `arista_eos`, auto-detected when omitted), `origin` (default `openconfig`), TLS settings, and `netbox_id` pinning to an existing NetBox device ID [secondary — NetBox Labs repo docs](https://github.com/netboxlabs/orb-agent/blob/HEAD/orb-discovery/gnmi-discovery/README.md).
- Interface name regexes classify interface types and exclude patterns (e.g. skip `^Management`) [secondary](https://github.com/netboxlabs/orb-agent/blob/HEAD/docs/backends/gnmi_discovery.md).

### 7.3 Real-time topology derivation

- LLDP remains the base source: device LLDP neighbor tables (OpenConfig `openconfig-lldp`) streamed via gNMI ON_CHANGE give link-level topology in near real time; BGP-LS (for link-state IGP domains) is the routed-fabric complement [secondary — standard practice].
- Combining LLDP streams + inventory (NetBox/Nautobot) + config (intent) yields live topology graphs used for impact analysis and event correlation [secondary].
- **Gap:** a named 2026-vintage open-source "real-time topology from gNMI LLDP" project comparison was not completed this wave; SuzieQ covers related ground (its own Phase E4 file) — record as cross-reference, not duplication.

