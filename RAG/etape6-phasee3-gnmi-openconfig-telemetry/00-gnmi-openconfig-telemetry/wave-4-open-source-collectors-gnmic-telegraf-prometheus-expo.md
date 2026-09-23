---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-4-open-source-collectors-gnmic-telegraf-prometheus-expo
title: "Wave 4 — Open-source collectors: gnmic, Telegraf, Prometheus exporters, pmacct"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [97, 145]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: cb4973382783ae3246a2b13bf1a365e6ed828cebb83d04012f2bee46abd0a8cd
---

# Wave 4 — Open-source collectors: gnmic, Telegraf, Prometheus exporters, pmacct

## Wave 4 — Open-source collectors: gnmic, Telegraf, Prometheus exporters, pmacct

### 4.1 gnmic as collector (OpenConfig project, 2026)

- Outputs (2026): `nats`, `jetstream`, `kafka`, `prometheus` (scrape endpoint) / `prometheus_write` (remote write), `influxdb`, `clickhouse`, `file`, `tcp` — plus multi-output fan-out from one instance [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/deployments/deployments_intro.md) [official](https://github.com/openconfig/gnmic/blob/HEAD/README.md).
- Deployments: single instance, **clustered/HA mode** (target connections load-shared across instances, data replicated), and forked/serial **data pipelines** (NATS→Prometheus, NATS→InfluxDB) with containerlab and docker-compose reference stacks [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/deployments/deployments_intro.md).
- Embedded **tunnel server** for gRPC-tunnel-based dial-out telemetry [official](https://github.com/openconfig/gnmic/blob/HEAD/README.md).
- **Event processors** transform data in flight (templates, label rewrites, grouping); **dynamic target discovery** loads targets at runtime from external systems; prompt mode offers YANG-based path auto-suggestions [official](https://github.com/openconfig/gnmic/blob/HEAD/README.md).
- Kubernetes operator (`gnmic/operator`): `Output` CRDs with `serviceRef`/`serviceSelector` for dynamic NATS/Kafka/Prometheus/InfluxDB endpoint discovery, cross-namespace references [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/user-guide/output.md).
- 2026 changelog highlights: Kafka `add-headers` per-message Go templates; InfluxDB drops non-finite floats per field instead of poisoning whole writes (with `gnmic_influxdb_output_non_finite_values_dropped_total` counter); bounded input queue + cancellable backpressure on Prometheus remote write; NATS input/output reconnect fixes [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/changelog.md).

### 4.2 Telegraf gNMI input plugin

- Telegraf's `inputs.gnmi` plugin is the documented OpenConfig sensor path in Juniper's AI/ML telemetry guide: per-sensor stanzas with addresses, credentials, `redial` on failure, subscription name + path + interval [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).
- Juniper's native JTI sensors use the separate `inputs.jti_openconfig_telemetry` plugin (unique client ID per sensor) in the same `telegraf.conf`, feeding one `outputs.influxdb` [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).
- **Gap:** Telegraf `inputs.gnmi` version matrix and ON_CHANGE vs SAMPLE support specifics per Telegraf release not verified this wave.

### 4.3 pmacct

- pmacct is a long-standing open-source flow/BGP/BMP collector; gNMI-native ingestion is not its core path — BMP/BGP-LS remain its network-visibility strength [unverified — flagged for a dedicated pass].

### 4.4 pygnmi and custom collectors

- **Gap:** pygnmi (Python gNMI client) current version and 2026 maintenance status not verified this wave; custom collectors typically built on the `openconfig/gnmi` Go protobufs or `pygnmi`.

## Wave 5 — Streaming telemetry pipeline architectures

### 5.1 Canonical network streaming-telemetry pipeline (2026)

- Reference layers, from vendor and open-source practice:
  1. **Devices** — dial-out (Cisco IOS-XR, Dell OS10, Juniper JTI streaming-server) or dial-in (Juniper default, gnmic/Telegraf poll-subscribe) gNMI/gRPC streams, sample or on-change cadences [official — Cisco, Dell, Juniper docs cited in Wave 3].
  2. **Collector tier** — gnmic (single, clustered/HA, or tunnel-server for dial-out), Telegraf `inputs.gnmi` [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/deployments/deployments_intro.md) [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).
  3. **Message bus** — NATS/JetStream or Kafka for buffering, fan-out, and replay; gnmic publishes natively to both [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/deployments/deployments_intro.md).
  4. **Storage** — Prometheus (short-term high-resolution), InfluxDB, ClickHouse (columnar, high-cardinality), TimescaleDB (SQL-native) [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/deployments/deployments_intro.md) [secondary](https://github.com/DhanrajGangnaik/Self-Hosted-Data-Platform).
  5. **Presentation/alerting** — Grafana dashboards, Alertmanager/webhook alerts, with telemetry-backed automation triggers [secondary](https://github.com/DhanrajGangnaik/Self-Hosted-Data-Platform).
- Nokia's published reference stack is the minimal viable version of this: **gnmic (collect + label processing) → Prometheus scrape → Grafana**, 5 s cadence [secondary](https://github.com/nokia/srexperts/blob/HEAD/docs/nos/srlinux/beginner/53-SR_Linux_Streaming_Telemetry.md).
- Juniper's AI-workloads guide is the InfluxDB variant: **JTI/gNMI → Telegraf → InfluxDB → Grafana** [vendor-reported](https://manuals.plus/juniper-networks/telemetry-in-junos-for-ai-ml-workloads-software-manual).

### 5.2 Scaling lessons (2026 practice)

- ClickHouse ingestion of high-velocity telemetry needs micro-batching (Kafka as "shock absorber") to avoid the `MergeTree` "too many parts" problem; tiered memtable→disk→S3 designs (e.g. CtrlB) address the "Kafka tax" [secondary](https://medium.com/google-cloud/ctrl-a-for-clickhouse-ctrl-f-for-ctrlb-redefining-the-telemetry-stack-f4e72c67b1bc).
- Cardinality: Prometheus ~10k labels practical limit vs ClickHouse ~1M per event (practical ~50k) — relevant when streaming per-queue/per-lane counters from AI fabrics [secondary](https://github.com/kubaik/kubaik.github.io/blob/HEAD/docs/observability-2026-ditch-the-three-pillars/index.md).
- gnmic's 2026 hardening directly targets this scale: bounded input queues, cancellable backpressure, per-message goroutine elimination, output-write lifecycle decoupling [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/changelog.md).

### 5.3 Event-driven automation off telemetry

- Patterns in use: threshold/correlation rules over the TSDB or bus triggering remediation (port bounce, BGP session flap correlation, ECN-storm → QoS policy), webhook/Kafka triggers into Ansible/event-driven automation [secondary — practice widely described in 2026 network-automation content; exact productized implementations vary by shop].
- ON_CHANGE subscriptions (gNMI) and sample-interval-0/event-driven streams (Cisco, Dell OS10) are the device-side enablers that make sub-second detection possible without polling storms [official — Cisco, Dell docs in Wave 3].

