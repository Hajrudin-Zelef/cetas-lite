---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/3-2-cisco-ios-xr-model-driven-telemetry
title: "3.2 Cisco (IOS-XR model-driven telemetry)"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent", "ethernet"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [51, 77]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: f66dce2c4d4b286e245480ef96088d1f0de822fe246f63ebab412086d3fc80ca
---

# 3.2 Cisco (IOS-XR model-driven telemetry)

- Junos Telemetry Interface (JTI): dial-in by default — collector initiates gRPC/gNMI to the device, default port **50051**; sensors addressed by native (`/junos/system/linecard/interface/`) or OpenConfig paths; verify with `show agent sensors` [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf).
- Dial-out exists as "Telemetry Interface Sensor" streaming-server config: `streaming-server` with remote-address/remote-port, `export-profile` (format `gpb`, transport `udp`, reporting-rate, DSCP/forwarding-class) — classic GPB-over-UDP path alongside gNMI [official — Juniper support KB](https://supportportal.juniper.net/s/article/Junos-Telemetry-Interface-Sensor-not-supported-in-routing-instance?language=en_US).
- Limitation: Telemetry Interface Sensor only works in the default/primary routing instance; interfaces in a VRF cannot source the telemetry stream [official](https://supportportal.juniper.net/s/article/Junos-Telemetry-Interface-Sensor-not-supported-in-routing-instance?language=en_US).
- `TARGET_DEFINED` subscription mode supported since Junos OS 20.2R1 on MX5/10/40/80/104/150/204/240/480/960/2008/2010/2020/10003/10008/10016 — the device picks STREAM vs ON_CHANGE per leaf; default periodic cadence 30 s unless the collector specifies `sample_interval` [official — Junos JTI user guide, mirrored](https://content.etilize.com/Additional-pdf4/1069319663.pdf).
- Junos OS Evolved 25.2R1 additions: physical Ethernet interface sensor on **ACX7509** via gRPC/gNMI with ON_CHANGE + periodic streaming, model `openconfig-if-ethernet.yang` v2.6.2 (sensor `/interfaces/interface/ethernet/state/`); MACsec statistics on ACX7100-32C/48L/ACX733/ACX7348/ACX7509 via `openconfig-macsec.yang` (sensor `/macsec/`) with additional out-of-model leaves augmented [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html).
- gNMI on Junos EVO: Get/Set/Subscribe(SAMPLE, ON_CHANGE, ONCE)/Capabilities; cited platforms QFX5220/5240 and PTX10000; OpenConfig paths used for multi-vendor consistency [secondary](https://github.com/chrishuffman5/domain-expert/blob/HEAD/./plugins/networking/skills/juniper-junos/SKILL.md).
- Juniper's own AI/ML-workloads telemetry guide pipelines JTI into **Telegraf** using two plugins side by side: `inputs.gnmi` for OpenConfig sensors and `inputs.jti_openconfig_telemetry` for native sensors, then `outputs.influxdb` into InfluxDB for Grafana — the documented reference pipeline for AI-fabric monitoring [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).

### 3.2 Cisco (IOS-XR model-driven telemetry)

- Dial-out MDT model: `telemetry model-driven` → `destination-group` (address-family ipv4/IPv6 + port, encoding `self-describing-gpb`, protocol `tcp`) → `sensor-group` (YANG `sensor-path`, e.g. `Cisco-IOS-XR-wdsysmon-fd-oper:system-monitoring/cpu-utilization`) → `subscription` binding sensor-group to destination with `sample-interval` in ms [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/710x/configuration/guide/b-telemetry-cg-ncs5500-710x/dial-out-telemetry-session-from-router-to-destination.html).
- Dial-out destinations support IPv4, IPv6, and FQDN (DNS, 128-char limit); on DNS failure the router retries every 30 s [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/710x/configuration/guide/b-telemetry-cg-ncs5500-710x/dial-out-telemetry-session-from-router-to-destination.html).
- gNMI encoding dial-out: `encoding json`, `protocol grpc` to a collector port (example 56000); **sample-interval 0 = event-driven incremental updates** (stats cache changes streamed on change, not polled); per-subscription counters (bytes/packets sent, send errors/drops) visible via `show telemetry model-driven subscription <name> internal` [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/telemetry/77x/b-telemetry-cg-ncs5500-77x.pdf).
- "gNMI Dial-Out via Tunnel Service" listed as an enhancement in the 7.7.x-era guide; a telemetry configuration guide exists for Cisco 8000 Series at release **26.x** (2026 train) [official](https://www.cisco.com/c/en/us/td/docs/iosxr/cisco8000/telemetry/26xx/configuration/guide/b-telemetry-cg-8000-26xx/dial-out-telemetry-session-from-router-to-destination.pdf).
- Scale guidance (community/Cisco): dial-out is the most widely used method; for collector control use dial-in instead; `max-containers-per-path 1024` needed for large sensor paths; dial-out → InfluxDB scaling documented in Cisco Community [secondary](https://community.Cisco.com/t5/service-providers-knowledge-base/how-to-scale-ios-xr-telemetry-with-influxdb/ta-p/4442024).

### 3.3 Nokia SR Linux

- SR Linux serves gNMI on TCP **57400** with its native YANG models (`srl_nokia-interfaces`, etc.); OpenConfig models must be explicitly enabled, and OpenConfig config can be pushed over NETCONF [secondary](https://github.com/netpilot-labs/example-prompts/blob/HEAD/advanced/netconf-gnmi-openconfig.md).
- Example lab flow (2026): `gnmic -a clab-...-leaf1,clab-...-spine1 subscribe --stream-mode SAMPLE --path "/interface[name=ethernet-*]/oper-state"`; native paths like `/interface[name=ethernet-1/*]/statistics`, `/interface[name=ethernet-1/*]/traffic-rate`, `/interface[name=ethernet-1/*]/oper-state` with `mode: stream`, `stream-mode: sample`, `sample-interval: 5s` [secondary](https://github.com/vista-/srlinux-hackathon-elisa/blob/HEAD/50_telemetry_gnmic/README.md).
- Nokia-authored docs (srexperts, updated Sept 2026) standardize the reference stack: **gnmic → Prometheus (scrape :9273/metrics every 5 s) → Grafana**, with gnmic event processors (`trim-sros-prefixes`, `add-labels`, `group-by-interface`) reshaping labels; a Docker label-based `loader` auto-discovers `nokia_srlinux` containerlab nodes [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md).
- gnmic itself is described as "an OpenConfig project developed by Nokia" — Nokia is both NOS vendor and collector author [secondary](https://github.com/alejo-guevara/srexperts/blob/HEAD/hackathon/activities/srlinux-b-streaming-telemetry/README.md).
- Contrast demonstrated in labs: SR Linux native path `/interface[name=...]/statistics` vs Arista cEOS OpenConfig path `/interfaces/interface[name=Ethernet1]/state/counters` — same counters, different module/hierarchy/field names [secondary](https://github.com/martimy/model-driven-configuration-tutorial/blob/HEAD/docs/tasks/12-stream-telemetry.md).
- SR Linux YANG browser available at `https://yang.srlinux.dev/v25.10.1` [secondary](https://github.com/vista-/srlinux-hackathon-elisa/blob/HEAD/50_telemetry_gnmic/README.md).

### 3.4 Dell SmartFabric OS10

