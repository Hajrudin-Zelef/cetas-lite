---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-11-gnmi-deep-dive-encodings-auth-cli-patterns-target-di
title: "Wave 11 — gNMI deep-dive: encodings, auth, CLI patterns, target discovery"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: model-release
actors: ["Huawei"]
dates: []
keywords: ["agent", "decode", "ethernet"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [255, 303]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: b311ef123c6966392e9864d8e4948440004c7dd1acf1eaffbc2714bfe01a0afb
---

# Wave 11 — gNMI deep-dive: encodings, auth, CLI patterns, target discovery

## Wave 11 — gNMI deep-dive: encodings, auth, CLI patterns, target discovery

### 11.1 gNMI subscription primitives (gnmic reference)

- Modes: `STREAM` (with stream-mode `target-defined` [default, 10 s], `sample`, `on-change`), `ONCE` (single snapshot), `POLL` (client-triggered) [official](https://gnmic.openconfig.net/user_guide/subscriptions/).
- CLI patterns (nokia/gnmic docs, 2026):
  - target-defined streaming, 10 s default: `gnmic -a <ip:port> sub --path /state/port[port-id=*]/statistics`
  - sampled, 30 s: `... --sample-interval 30s`
  - on-change with 1 min heartbeat: `... --stream-mode on-change --heartbeat-interval 1m`
  - one-shot: `... --mode once` [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).
- Heartbeats matter: on-change subscriptions with `--heartbeat-interval` keep the stream provably alive without change traffic — the pattern for event-driven monitoring [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).
- Multiple subscriptions can share one gRPC stream with mixed modes (e.g. `/configure/system/name` on-change + `/state/port/statistics` sampled 10 s) [official](https://gnmic.openconfig.net/user_guide/subscriptions/).

### 11.2 gNMI extensions in the wild

- **History extension** (`gnmi-history.md`): `--history-snapshot`, `--history-start`, `--history-end` — time-range queries over recorded telemetry (RFC3339 or ns-since-epoch) [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).
- **Depth extension** (`gnmi-depth.md`): `--depth` limits subtree depth in Subscribe responses [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).
- **Vendor support for extensions is uneven** — the history extension in particular requires device-side recording; treat availability as per-device [unverified — flagged].

### 11.3 Encodings, auth, and transport

- Encodings seen in 2026: `JSON_IETF`, `ASCII`, `PROTO` (Cisco IOS-XR); `bytes`/`proto` (Nokia); `json` (Cisco dial-out) [secondary](https://github.com/openconfig/kne/blob/HEAD/examples/cisco/8000e/README.md) [official](https://www.Cisco.com/c/en/us/td/docs/iosxr/ncs5500/telemetry/77x/b-telemetry-cg-ncs5500-77x.pdf).
- TLS: default on for gNMI in most NOS; `--skip-verify` / `insecure` / `skip_verify` in labs; Dell OS10 gNMI agent uses TLS + username/password with Super-Admin-only access [official](https://www.dell.com/support/manuals/en-ie/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-2-6/gnmi-agent?guid=guid-9d4c66d3-178d-4122-baed-6914977fe93e&lang=en-us) [secondary — lab configs].
- Juniper dial-in requires client cert+key when TLS enabled; protobuf definitions must be present on the collector to decode [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf).

### 11.4 Target discovery at scale

- gnmic dynamic loaders: containerlab/Docker label discovery (Nokia srexperts pattern), plus the operator's NetBox REST/export-template/webhook `TargetSource` (Wave 7) and K8s service discovery for outputs [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md) [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/examples/NetBox/REST%20API/_index.md).
- Clustering: multiple gnmic instances share targets via lock keys (`--cluster-name`, `--lock-retry 5s` default) for connection load-sharing and resiliency [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).

## Wave 12 — Vendor matrix: gNMI/OpenConfig support at a glance (2026)

| Vendor / NOS | gNMI port (default) | Dial-in / Dial-out | OpenConfig posture | Notable 2026 facts | Provenance |
|---|---|---|---|---|---|
| Cisco IOS-XR | gRPC agent port (configurable; KNE lab used 9339 ext) | Both; dial-out MDT primary | OC models served; native `Cisco-IOS-XR-*` sensor paths common | gNMI 0.8.0; OC bgp-types 5.3.1; telemetry guide at 26.x train; gNMI dial-out via tunnel service | [official](https://www.cisco.com/c/en/us/td/docs/iosxr/cisco8000/telemetry/26xx/configuration/guide/b-telemetry-cg-8000-26xx/dial-out-telemetry-session-from-router-to-destination.pdf) [secondary](https://github.com/openconfig/kne/blob/HEAD/examples/cisco/8000e/README.md) |
| Cisco NX-OS | gRPC agent | gNMI dial-in; MDT dial-out on some platforms | OC via NX-API/gNMI; gNOI full (System/OS/Cert/File/FactoryReset) | Max 16 active gNOI RPCs; container services exposed | [official](http://www.static-cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/106x/programmability/cisco-nexus-9000-series-nx-os-programmability-guide-106x/gnoi---operation-interface.pdf) |
| Cisco IOS-XE | gRPC agent | Both | gNOI in 17.18 prog. guide; OC `openconfig-system-management` for gNOI auth | gNMI/gNOI parity with NX-OS on Cat9k | [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf) |
| Juniper Junos | 50051 | Dial-in default; dial-out via JTI streaming-server | Native + OC paths; OC version table published per release | TARGET_DEFINED since 20.2R1; ACX7509 OC if-ethernet 2.6.2 sensor (25.2R1); JTI sensor VRF limitation | [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf) [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html) |
| Arista EOS | 6030 (Octa) | Dial-in; TerminAttr/CVP dial-out | OC served; deviations published | `management api gnmi` CLI; 4.33.1F lab quirk rejecting `no shutdown` | [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md) |
| Nokia SR Linux | 57400 | Dial-in (gnmic); EDA dial-out patterns | Native first (`srl_nokia-*`); OC explicitly enabled | YANG browser `yang.srlinux.dev` v25.10.1; gnmic authored by Nokia | [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md) |
| Dell OS10 | (telemetry agent) | Dial-out only (one dest group, one profile) | OC streamed when SmartFabric director sets it via gNMI; native `base-*` sensors | 15 s min cadences; 0 = event; `.proto` from support site; gNMI agent needs SFD mode | [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/configure-telemetry?guid=guid-a56d32ee-b7b9-4667-9156-7b62845701f6&lang=en-us) |
| Extreme ONE OS | gRPC | Dial-in | Dedicated gNOI reference guide v22.2.0.0 | Extensive Extreme gNOI extensions (auth/cfgmgmt/firmware/proto-per-service) | [official](https://documentation.extremenetworks.com/Extreme%20ONE%20OS%20Switching%20v22.2.0.0%20GNOI%20Reference%20Guide/Extreme_ONE_OS_Switching_22_2_0_0_GNOI_Reference_Guide.pdf) |
| SONiC (whitebox) | varies | Dial-in | OC via management framework | FRR `openconfigd` partial in lab contexts; PFC/ECN via native models | [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md) |
| Huawei | (gNMIc issue) | Dial-in attempted | OC versions listed (bgp 4.1.0, interfaces 2.3.0, acl 1.0.0) | Integration issue reported in gnmic repo — multi-vendor friction evidence | [secondary](https://github.com/openconfig/gnmic/issues/451) |

**Reading notes:**
- "Dial-out only" (Dell) vs "dial-in default" (Juniper) is a real architectural fork: dial-out scales collectors horizontally and traverses NAT; dial-in gives the collector subscription control. gnmic's embedded tunnel server bridges the two [official — gnmic docs].
- Version skew is the norm, not the exception: always `Capabilities` first.

