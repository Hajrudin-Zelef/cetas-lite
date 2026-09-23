---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-25-juniper-jti-dial-out-configuration-anatomy
title: "Wave 25 — Juniper JTI dial-out configuration anatomy"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent", "ethernet", "revenue"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [532, 576]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 2d21ad23776838b6179965d47fadd08a6e0d99714b503adab747cce795de9753
---

# Wave 25 — Juniper JTI dial-out configuration anatomy

## Wave 25 — Juniper JTI dial-out configuration anatomy

### 25.1 The three-part config (streaming-server, export-profile, sensor)

JTI dial-out config under `[edit services analytics]` has three pieces [secondary](https://github.com/jawroper/jti_nus_fluentbit):

1. **Streaming server** — the collector:
   - `set services analytics streaming-server <name> remote-address <IP>`
   - `set services analytics streaming-server <name> remote-port 4729` (example port)
2. **Export profile** — transport and cadence:
   - `set services analytics export-profile <name> transport udp`
   - `set services analytics export-profile <name> reporting-rate 30`
   - Optional: `local-address <mgmt-IP>`, `routing-instance mgmt_junos`
3. **Sensor** — what to stream, one block per resource path:
   - `set services analytics sensor interfaces server-name <server>`
   - `set services analytics sensor interfaces export-name <profile>`
   - `set services analytics sensor interfaces resource /junos/system/linecard/interface/`
   - Chassis example: `resource /junos/chassis/`

### 25.2 Sensor refinements

- `resource-filter <regex>` narrows within a resource (e.g. `et-*` for Ethernet logical interfaces only) [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf).
- A sensor can name **multiple streaming servers** (bracketed list) — native fan-out to redundant collectors [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf).
- Verify: `show agent sensors`, `show services analytics` [secondary](https://github.com/jawroper/jti_nus_fluentbit).

### 25.3 Platform constraints (2026)

- Non-EVO Junos: native UDP streaming only via **revenue interface**; Junos EVO: revenue or management interface (in default routing instance) [secondary](https://github.com/jawroper/jti_nus_fluentbit).
- 25.2R1 enhancement: **multiple gRPC servers with distinct services, listening addresses, and ports** on ACX7024/ACX7024X/ACX7100-32C/ACX7100-48L/ACX7332/ACX7348/ACX7509, PTX10001-36MR/PTX10002-36QDD/PTX10003/PTX10004/PTX10008/PTX10016, QFX5130-32CD/48C/48CM/5130E-32CD, QFX5220, QFX5230-64CD, QFX5700/5700E — per-service ports plus TLS certificate configuration [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html).
- Native YANG state model + telemetry for network-stack protocol stats (TTP, ICMP, MPLS, TCP…) on ACX/QFX/PTX EVO — previously CLI-only counters, now streamable on-change or periodic [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html).
- MACsec statistics sensor path: `/macsec/` [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html).

## Wave 26 — Collector configuration patterns (Telegraf, Fluent Bit, JTI-native)

### 26.1 Telegraf dual-plugin pattern (Juniper AI/ML guide)

- `inputs.gnmi` for OpenConfig sensors (hostname, port 50051, credentials, redial interval; `inputs.gnmi.subscription` stanzas per sensor with name, path, interval) [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).
- `inputs.jti_openconfig_telemetry` for **native JTI sensors** (unique client ID per sensor, e.g. `telegraf3`) [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).
- `outputs.influxdb` (database `telegraf`); verify with `select * from <measurement> limit 1` in the InfluxDB CLI [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).

### 26.2 Fluent Bit for native UDP JTI

- Community `jti_nus_fluentbit` project: Juniper native-UDP sensor input plugin for Fluent Bit; device streams UDP to the Fluent Bit host [secondary](https://github.com/jawroper/jti_nus_fluentbit).
- Also observed: Synse JTI plugin ingesting UDP on port 5566 (`udp://0.0.0.0:5566`) — the native-UDP collector ecosystem is small but real [secondary](https://GitHub.Com/vapor-ware/synse-juniper-jti-plugin).

