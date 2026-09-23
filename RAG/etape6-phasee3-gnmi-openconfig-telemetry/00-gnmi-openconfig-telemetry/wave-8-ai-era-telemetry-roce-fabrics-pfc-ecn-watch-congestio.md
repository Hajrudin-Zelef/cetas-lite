---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-8-ai-era-telemetry-roce-fabrics-pfc-ecn-watch-congestio
title: "Wave 8 — AI-era telemetry: RoCE fabrics, PFC/ECN watch, congestion visibility"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: ["Broadcom", "Huawei", "Nvidia"]
dates: ["2026-01-14", "2026-08-03", "2026-09-01"]
keywords: ["decode", "ethernet", "nvidia", "optics", "training"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [194, 254]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 7e405ba623b93ffa035048d1ba918e113aaa6a08ba210e937773a5dbfd4064da
---

# Wave 8 — AI-era telemetry: RoCE fabrics, PFC/ECN watch, congestion visibility

## Wave 8 — AI-era telemetry: RoCE fabrics, PFC/ECN watch, congestion visibility

### 8.1 What AI fabrics demand from telemetry

- RoCEv2 fabrics (the Ethernet answer to InfiniBand for AI training) require **lossless** behavior: Priority Flow Control (802.1Qbb) creates per-priority no-drop classes; ECN marks instead of dropping; DCQCN (end-to-end, QCN-based) throttles senders. The stack is commonly summarized PFC → ECN/QCN → DCQCN/IBCC [secondary](https://www.perfecxion.ai/pdfs/congestion-control-telemetry-security-ai-fabrics-PDF.pdf).
- FS.com's 2026 positioning: RoCE's advantage is combining RDMA efficiency with Ethernet flexibility, enabled by PFC, ECN, DLB/GLB on standard leaf-spine [vendor-reported](https://www.fs.com/blog/infiniband-vs-roce-how-to-choose-a-network-for-data-center-2521.html).
- Congestion observability = watching the control loops work: PFC PAUSE frame rates per priority, ECN mark rates, queue depths, buffer-pool occupancy, and NIC-side retransmit/CNP counters [secondary — operational practice].

### 8.2 PFC watchdog and storm detection

- PFC watchdog (Arista EOS pattern, mirrored in guides): monitors interfaces for PFC pause storms; actions include ignoring received pause frames, dropping on the interface, or error-disabling the port — because a pause storm signals a misbehaving downstream node [secondary](https://sup1rppylvxprc.vcaahomes.top/assets/data/pdf/AI-Network-Fabric_Deployment_Guide.pdf).
- SONiC: default lossless priorities 3 and 4 (DSCP 3/4 mapped); optional PFC watchdog (`pfcwd_sw_enable`); every PFC-enabled queue gets a WRED/ECN profile (`AZURE_LOSSLESS` style: `ecn_all`, green/yellow/red thresholds) for proactive congestion management [secondary](https://netbergtw.com/wp-content/uploads/Files/netberg_sonic_ai_fabric_rdma.pdf) [official — test plan](https://github.com/sonic-net/sonic-mgmt/blob/HEAD/docs/testplan/PFC-test-plan.md).
- Watchdog tuning example from a 2026 AI-fabric guide: `priorityflowcontrol pause watchdog default timeout 0.20`, `recoverytime 0.20`, `pollinginterval 0.100`, `action drop` [secondary](https://sup1rppylvxprc.vcaahomes.top/assets/data/pdf/AI-Network-Fabric_Deployment_Guide.pdf).
- Telemetry mapping: PFC counters (pause frames sent/received per priority, watchdog events), ECN-marked packet counts, WRED drop counters, and queue/buffer-pool stats are the per-vendor sensor set to subscribe on-change or at 5–15 s cadence; Dell OS10 exposes these via `base-qos/queue-stat`, `priority-group-stat`, `buffer-pool-stat` [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/example-configure-streaming-telemetry?guid=guid-90cd942e-bd77-4c20-a2e6-73aa0d4f8d7e&lang=en-us).

### 8.3 Test and validation of ECN/PFC behavior

- Xena/Teledyne LeCroy 2026-era whitepaper: verify ECN marking by ramping utilization 0→100% and counting ECN-11-marked packets; if ECN unsupported or egress queue overfills, PFC activates and the generator receives pause requests; SierraNet M1288 analyzers decode ECN/PFC at 25G–800G [vendor-reported](https://xenanetworks.com/wp-content/uploads/documentation/whitepaper/AI_Switch_testing_WP_TeledyneLecroyXena_Sierra.pdf).
- SONiC PFC test plan (updated Sept 2026): PAUSE frame format (ethertype, opcode 01-01, priority-enable vector, per-priority timers); lossless priorities only react/generate PFC; SONiC ignores 802.3x link-level PAUSE [official](https://github.com/sonic-net/sonic-mgmt/blob/HEAD/docs/testplan/PFC-test-plan.md).

### 8.4 Open items specific to AI telemetry

- Per-vendor gNMI path inventories for PFC/ECN/buffer counters (Cisco `Cisco-IOS-XR-*` vs OpenConfig `openconfig-qos`) are not consolidated this wave [unverified — open item].
- Congestion-visibility products (Broadcom/others' in-band telemetry/INT offerings) and their 2026 adoption are out of scope for this wave; record as follow-up.
- NVIDIA Spectrum-X telemetry specifics (what counters Spectrum-X switches expose over gNMI/Redfish for AI fabrics) not verified this wave [unverified — open item].

## Wave 9 — Coverage audit, open items, and close

## Wave 10 — OpenConfig model catalog: key models, versions, and coverage gaps

### 10.1 Model versions observed in the wild (2026)

| Model | Version seen | Where | Provenance |
|---|---|---|---|
| `openconfig-bgp` | 4.1.0 | Huawei-target capabilities dump (dated) | [secondary](https://github.com/openconfig/gnmic/issues/451) |
| `openconfig-bgp-types` / `openconfig-bgp-errors` | 5.3.1 | Cisco 8000 IOS-XR live (KNE lab) | [secondary](https://github.com/openconfig/kne/blob/HEAD/examples/cisco/8000e/README.md) |
| `openconfig-bgp-policy` | 5.0.0 | capabilities dump | [secondary](https://github.com/openconfig/gnmic/issues/451) |
| `openconfig-interfaces` | 2.3.0 | capabilities dump | [secondary](https://github.com/openconfig/gnmic/issues/451) |
| `openconfig-if-ethernet` | 2.6.2 | Junos Evolved 25.2R1 (ACX7509 sensor) | [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html) |
| `openconfig-acl` | 1.0.0 | capabilities dump | [secondary](https://github.com/openconfig/gnmic/issues/451) |
| `openconfig-aft` | 3.1.0 | post counters-rework | [official](https://github.com/openconfig/public/pull/1330) |
| `openconfig-terminal-device` | 1.12.0 (2026-01-14) | optical transport | [official](https://github.com/openconfig/public/blob/master/release/models/optical-transport/openconfig-terminal-device.yang) |
| `openconfig-platform-types` | 1.13.0 | power-oper-state addition (2026-09-01) | [official](https://github.com/openconfig/public/commit/af2b0f11def9fb0cad7f4a04d28d6d5bef2b7da8) |
| `openconfig-platform-pipeline-counters` | 0.6.1 | trigger-semantics clarification (2026-08-03) | [official](https://github.com/openconfig/public/pull/1494) |
| `openconfig-macsec` | (sensor `/macsec/`) | Junos Evolved 25.2R1 (ACX series) | [official](https://www.juniper.net/documentation/us/en/software/junos/release-notes/25.2/junos-evo-release-notes-25.2r1/topics/new-features/feature-descriptions/junos-telemetry-interface.html) |
| `openconfig-lldp-types` | 0.1.1 | Arista capabilities dump | [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf) |
| `openconfig-isis-lsdb-types` | 0.4.2 | Arista capabilities dump | [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf) |

### 10.2 QoS model and featureprofiles coverage

- Telemetry paths `/qos/interfaces/interface/output/queues/queue/state/{transmit-pkts, transmit-octets, dropped-pkts, dropped-octets}` are validated by featureprofiles **DP-1.4** (`qos_output_queue_counters_test`): send NC1/AF4/AF3/AF2/AF1/BE1 traffic, subscribe every 30 s, collect ≥10 samples, verify counter increments [official](https://github-wiki-see.page/m/openconfig/featureprofiles/wiki/qos_output_queue_counters_test).
- QoS config via OpenConfig maps to vendor native: Juniper documents `openconfig-qos` scheduler-policy → Junos `class-of-service` mapping (scheduler-maps, traffic-control-profiles; STRICT vs WRR handling) [official](https://www.juniper.net/documentation/us/en/software/junos/open-config/topics/concept/open-config-qos-mapping.html).
- Arista maintains a fork `aristanetworks/openconfig-featureprofiles` with QoS OTG tests (egress classification/rewrite, queue counters) — evidence vendors run featureprofiles against their own code [secondary](https://github.com/openconfig/gnmic/issues/451).

### 10.3 Gaps vs vendor YANG (persistent 2026 themes)

- **Version skew:** devices in the field expose OpenConfig models years behind HEAD (e.g. Junos support-portal table lists BGP 2.0.1/2.1.1-era models for 16.1R/17.1R; the gnmic issue dump shows interfaces 2.3.0 while 2026 revisions exist) — multi-version fleets are the norm [official](https://supportportal.juniper.net/sfc/servlet.shepherd/document/download/0693c00000LXblkAAD/?operationContext=S1) [secondary](https://github.com/openconfig/gnmic/issues/451).
- **Deviations:** vendors publish deviation modules (e.g. Arista `arista-bfd-deviations`, `arista-exp-eos-*`) where OpenConfig doesn't fit the implementation [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).
- **Coverage holes:** platform pipeline counters, MACsec, terminal-device optics, and power-state are 2025–2026 additions — recent enough that not all vendors implement them [official — dated revisions in Waves 1–2].
- **gNMI `Capabilities` as the discovery mechanism:** clients must query supported models/encodings per device rather than assuming a uniform OpenConfig surface [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).

