---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-16-arista-deep-dive-eos-streaming-cloudvision-netdl-ava
title: "Wave 16 — Arista deep-dive: EOS streaming, CloudVision NetDL, AVA, AI observability"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["agent", "agents", "gpu", "nvidia"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [377, 432]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 48a838cc9f26c7d261c5c3c513be95b5861f309c55a82ace4ea857870354515d
---

# Wave 16 — Arista deep-dive: EOS streaming, CloudVision NetDL, AVA, AI observability

## Wave 16 — Arista deep-dive: EOS streaming, CloudVision NetDL, AVA, AI observability

### 16.1 EOS streaming architecture

- EOS devices stream granular network state (sub-second granularity in some cases) to CloudVision over **gRPC and OpenConfig** — Arista's positioning against legacy SNMP polling [vendor-reported](https://intelligentvisibility.com/unified-observability-logicmonitor).
- The on-device streaming agents ship packaged with EOS; the state-streaming concept grew out of NetDB (EOS's centralized network-state database) [vendor-reported](https://www.sdxcentral.com/news/arista-rolls-out-real-time-streaming-of-state-data/).

### 16.2 NetDL — the network data lake

- NetDL maintains real-time state streamed from all Arista EOS devices plus third-party services, kept as time series so operators can snapshot or go back in time for root-cause review [vendor-reported](https://hello.doclang.workers.dev/best_codes/how-https-www.arista.com/assets/data/pdf/Datasheets/EOSCloudVision_DataSheet.pdf).
- Multi-domain: data center, campus/Wi-Fi, branch, WAN, private cloud under one UI with tailored dashboards [vendor-reported](https://www.businesswire.com/news/home/20240924881511/en/).

### 16.3 AVA and CV UNO

- **AVA (Autonomous Virtual Assist)**: CloudVision's AI/ML engine over NetDL — anomaly detection, predictive analytics (outage forecasting), intelligent root-cause analysis [vendor-reported](https://intelligentvisibility.com/unified-observability-logicmonitor).
- **CV UNO (Universal Network Observability)**: extends beyond network-centric views with an application-to-network graph, flow data ingestion, and ML correlation across topology/time/function dimensions [vendor-reported](https://intelligentvisibility.com/unified-observability-logicmonitor).

### 16.4 AI-fabric observability

- Arista EOS-based **AI Agent** installs on NVIDIA BlueField-3 SuperNICs: remote control + visibility into the NIC, QoS coordination with the fabric, and NIC telemetry streamed upstream into NetDL for analytics/reporting [secondary](https://www.sdxcentral.com/news/arista-weaves-in-optimized-ai-fabric-for-networking/).
- **AI Job-Centric Observability**: AI job health metrics, deep-dive analytics into network/server NIC performance for AI workloads, flow visualization for AI job traffic [vendor-reported](https://intelligentvisibility.com/unified-observability-logicmonitor).
- Arista × VAST Data partnership: VAST's DASE storage architecture integrated into the AI fabric story for GPU data pipelines [secondary](https://www.sdxcentral.com/news/arista-weaves-in-optimized-ai-fabric-for-networking/).

## Wave 17 — gNMI Set: config management over gRPC

### 17.1 Set semantics

- gNMI `Set` supports replace/update/delete of config subtrees addressed by YANG paths — the transactional alternative to NETCONF edit-config for programmatic config, without NETCONF's candidate/commit ceremony [secondary].
- Dell's SmartFabric director uses gNMI Set against the OS10 gNMI agent to configure the telemetry agent itself (OpenConfig-format streaming) — a concrete production use of Set for telemetry bootstrapping [official](https://www.dell.com/support/manuals/en-ie/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-2-6/gnmi-agent?guid=guid-9d4c66d3-178d-4122-baed-6914977fe93e&lang=en-us).
- Junos EVO supports gNMI Set for config modification alongside Get/Subscribe/Capabilities [secondary](https://github.com/chrishuffman5/domain-expert/blob/HEAD/./plugins/networking/skills/juniper-junos/SKILL.md).

### 17.2 Config-vs-telemetry division (2026 practice)

- Push config via NETCONF (transactional) or gNMI Set / Ansible / Terraform; stream state via gNMI Subscribe; operate the device via gNOI. The three planes are complementary, not competing — the Nokia/Arista 2-node lab demonstrates NETCONF config-push + gNMI telemetry on the same devices [secondary](https://github.com/netpilot-labs/example-prompts/blob/HEAD/advanced/netconf-gnmi-openconfig.md).
- OpenConfig as the shared schema makes one config path work across vendors — the explicit goal of the multi-vendor lab pattern [secondary](https://github.com/netpilot-labs/example-prompts/blob/HEAD/advanced/netconf-gnmi-openconfig.md).

## Wave 18 — Deployment playbook: building the stack in practice

### 18.1 Minimal viable stack (lab)

1. Devices: SR Linux (gNMI :57400), cEOS (gNMI :6030), or IOS-XR — enable gNMI, note per-NOS port and model differences (Wave 3/12).
2. Collector: `gnmic subscribe --config gnmic.yml` with named subscriptions (sample 5 s for counters, on-change for oper-state) → Prometheus output on :9273 [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md).
3. TSDB: Prometheus scrape every 5 s; dashboards in Grafana.
4. Target discovery: Docker-label loader for containerlab, or NetBox `TargetSource` for production (Wave 7).

### 18.2 Production hardening checklist

- TLS everywhere; cert rotation via gNOI Cert service (Wave 13).
- Clustered gnmic for connection load-sharing and data replication; NATS/Kafka buffer between collectors and TSDB (Wave 5).
- Heartbeats on on-change subscriptions so silent streams are distinguishable from dead ones [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).
- Per-subscription cardinality budget: per-queue/per-lane AI counters explode label sets — ClickHouse or downsampling for high-cardinality paths (Wave 5.2).

### 18.3 Alerting patterns

- Threshold alerts (interface errors, PFC storm rate, buffer occupancy), rate-of-change (BGP session flaps, route churn), and absence-of-data (heartbeat loss) — computed in the TSDB or stream processor, routed via Alertmanager/webhooks into event-driven automation (Wave 5.3) [secondary].

