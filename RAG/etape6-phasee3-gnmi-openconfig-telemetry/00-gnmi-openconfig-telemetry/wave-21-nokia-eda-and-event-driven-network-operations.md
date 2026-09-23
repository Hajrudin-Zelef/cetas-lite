---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-21-nokia-eda-and-event-driven-network-operations
title: "Wave 21 — Nokia EDA and event-driven network operations"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: ["Huawei", "Nvidia"]
dates: ["2026-01-14", "2026-09-01", "2026-09-22"]
keywords: ["agent", "decode", "nvidia"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [479, 531]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: ba3dbb657182ced8d2ef496bf1462629c72ad29b07519356dd8c2e66ab13c829
---

# Wave 21 — Nokia EDA and event-driven network operations

## Wave 21 — Nokia EDA and event-driven network operations

### 21.1 Where EDA fits

- Nokia's **EDA (Event Driven Automation)** is the intent-based platform direction for SR Linux fabrics; srexperts material positions gNMI streaming as the telemetry feed that event-driven automation consumes [secondary — Wave 7 source family; EDA specifics beyond this remain a 2026 tracking item] [unverified — detailed EDA telemetry integration].
- The generic pattern (vendor-neutral): on-change gNMI subscriptions → event bus (NATS/Kafka) → automation rules (Wave 5.3) — EDA productizes this loop for SR Linux fabrics [analysis].

## Wave 22 — Troubleshooting playbook: telemetry failures

### 22.1 Stream down / no data

- Check TLS first: cert expiry, CA trust, hostname verification — the most common dial-in failure class across vendors [analysis].
- Juniper dial-in: verify client cert+key presented and protobuf definitions present on collector for decode [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf).
- Arista tunnel dial-out: `show agent gnmireverse logs` [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md).
- Dell dial-out: confirm only one destination group/profile exists (config rejected otherwise) and check `show telemetry` stream state [official](https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/configure-telemetry?guid=guid-a56d32ee-b7b9-4667-9156-7b62845701f6&lang=en-us).

### 22.2 Stream alive but path returns nothing

- Run gNMI `Capabilities` — the path or model version may not exist on this NOS release (Wave 3.4 symptom class; Arista 4.33.1F `no shutdown` rejection is the config-side analogue) [secondary](https://github.com/gesh75/multivendor-ai-network-lab/blob/HEAD/docs/STREAMING_TELEMETRY_GAPS.md).
- Check model version skew (Wave 12 table): the path may exist under a different revision's tree.

### 22.3 Silent stream vs dead stream

- On-change subscriptions without heartbeats are ambiguous — always configure `--heartbeat-interval` so absence of data is distinguishable from absence of change [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).
- gnmic cluster lock contention: check `--lock-retry` behavior and cluster membership when targets flap between collectors [official](https://github.com/nokia/gnmic/blob/HEAD/docs/cmd/subscribe.md).

## Wave 23 — OpenConfig 2026 process: how the models evolve

### 23.1 Release mechanics

- Quarterly semantic-versioned bundles with breaking-change policy (Wave 1); model-level revisions land continuously (terminal-device 1.12.0 on 2026-01-14, platform power states 2026-09-01) [secondary](https://github.com/openconfig/public/blob/HEAD/doc/releases.md).
- Featureprofiles provides the compliance test suite that turns model text into verifiable device behavior (qos_output_queue_counters_test DP-1.4 as the QoS example) [secondary](https://github-wiki-see.page/m/openconfig/featureprofiles/wiki/qos_output_queue_counters_test).

### 23.2 The adoption lag

- Vendor implementation lags HEAD by years (bgp 3.2.0 on Junos 25.2 vs 5.3.1 in IOS-XR; interfaces 2.3.0 on Huawei) — OpenConfig is a *direction*, not a *uniform API*, in 2026 [secondary](https://github.com/openconfig/gnmic/issues/451).
- Operational consequence: per-device capability discovery and path inventories are mandatory; multi-vendor subscriptions must be parameterized by NOS/release [analysis].

## Wave 24 — Outlook and final notes

### 24.1 Trajectories visible in 2026

- gNMI/gNOI convergence with AI-fabric operations: Arista's NIC-level telemetry into NetDL, NVIDIA-adjacent ecosystem pressure, and 800G test methodologies (Wave 8.4) point to congestion-signal standardization as the next modeling frontier [secondary](https://www.sdxcentral.com/news/arista-weaves-in-optimized-ai-fabric-for-networking/) [secondary](https://xenanetworks.com/wp-content/uploads/documentation/whitepaper/AI_Switch_testing_WP_TeledyneLecroyXena_Sierra.pdf).
- gRPC tunnel dial-out normalizes the NAT/cloud collector topology problem across Arista and Cisco — expect broader vendor uptake [official](https://github.com/aristanetworks/openmgmt/blob/HEAD/docs/telemetry/adapters/gnmi-dial-out/index.md).
- gnmic operator maturity (K8s-native pipelines, TunnelTargetPolicy, NetBox TargetSource) makes the Kubernetes-orchestrated telemetry stack the 2026 reference deployment [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/user-guide/pipeline.md).

### 24.2 What this file does not cover (delegated)

- Step-by-step lab build: see Wave 18.1 pointer to srexperts; full reproducible lab is E2/E4-file territory.
- NetBox/Nautobot data models: E1 file. SuzieQ/Batfish depth: E4 file.

*Supplementary audit (2026-09-22): Waves 16–24 appended. File now holds 24 waves, sole writer, append-only. Wave 15's open-item list (12 items) stands; Wave 19.4 and 20.x partially address items 7 and 6 only at the documentation level. Provenance tags verified throughout; no identifiers invented.*

