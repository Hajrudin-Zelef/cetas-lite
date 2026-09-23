---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/overview
title: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: []
dates: ["2025-11-07", "2025-12-11", "2026-01-14", "2026-08-03", "2026-09", "2026-09-01", "2026-09-22"]
keywords: ["benchmark", "optics", "research"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [1, 46]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 665fea2252a20d2f98026dd74c7c7fa172d62f3c2914bd0df8d217466aa67ebb
---

# Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability

> **Scope:** OpenConfig models and 2026 working-group status, vendor adoption; gNMI/gNOI protocols, dial-in vs dial-out telemetry, subscription modes; vendor streaming-telemetry implementations; open-source collectors; streaming telemetry pipeline architectures; real-time topology tools and NetBox/Nautobot integration; NETCONF/RESTCONF vs gNMI in 2026; YANG Suite and pygnmi; telemetry for AI-fabric monitoring (RoCE counters, PFC/ECN, congestion visibility).
> **Coverage date:** 2026-09-22. All facts current through this date unless noted.
> **Method:** read-only web research (browser_search / browser_open). No live-browser visits, no sign-ins, nothing sent externally. Single writer, append-only file.
> **Provenance legend:** `[official]` = vendor/standards-body documentation or announcements · `[vendor-reported]` = vendor blog/whitepaper/benchmark claims not independently verified · `[independent]` = third-party lab, test, or analyst measurement · `[secondary]` = press, community summaries, forum/Reddit · `[unverified]` = single-source or unconfirmed claim.
> **Identifier rule:** no SKUs, model names, URLs, or numbers were guessed; all URLs are verbatim from search results.

---

## Wave 1 — OpenConfig 2026: models, working group, and vendor adoption

### 1.1 OpenConfig release cadence and versioning policy

- OpenConfig publishes the full compatible model set from `openconfig/public` roughly every quarter; tags follow semantic versioning (`vx.y.z`, e.g. `v1.2.0`); non-backward-compatible changes are deliberately infrequent and batched [official](https://github.com/openconfig/public/blob/HEAD/doc/releases.md).
- A YANG "release bundle" concept (machine-readable module catalog `openconfig-module-catalog.yang`) underpins the quarterly tags [official](https://github.com/openconfig/public/blob/HEAD/doc/releases.md).

### 1.2 Model revisions observed in 2026

- `openconfig-terminal-device.yang` (optical transport): version 1.12.0, revision 2026-01-14 — "Removal of references to unused/removed LLDP groupings" [official](https://github.com/openconfig/public/blob/master/release/models/optical-transport/openconfig-terminal-device.yang).
- Prior revisions in that file: 1.11.0 (2025-12-11, moved `fec-uncorrectable-blocks` leaf), 1.10.0 (2025-11-07, FEC counter semantics aligned to implementation intent and PM interval) — evidence of active 2025→2026 work on coherent-optics PM modeling [official](https://github.com/openconfig/public/blob/master/release/models/optical-transport/openconfig-terminal-device.yang).
- Platform models (2026-09-01): `openconfig-platform-types.yang` 1.13.0 adds `component-power-oper-type` typedef (POWER_ENABLED / POWER_DISABLED / operational-only POWER_TRANSITIONING); `openconfig-platform-common.yang` 0.33.0 adds read-only `power-oper-state` as the operational counterpart of `power-admin-state`; controller-card, fabric, linecard models bumped to track [official](https://github.com/openconfig/public/commit/af2b0f11def9fb0cad7f4a04d28d6d5bef2b7da8).
- Platform pipeline counters (2026-08-03): `openconfig-platform-pipeline-counters.yang` 0.6.1 (0.5.2 in an earlier revision of the same PR) clarifying trigger semantics of `threshold` and `active` leaves; review thread notes a conflicting change that bumped the file to 6.0/6.1 — evidence of parallel vendor/operator branches landing in 2026 [official](https://github.com/openconfig/public/pull/1494).
- AFT models: `openconfig-aft*.yang` version 3.1.0 after moving counters out of atomic sub-containers in the `afts/` tree (~late 2025) [official](https://github.com/openconfig/public/pull/1330).

### 1.3 Known working-group scale and community signals (2026)

- The `openconfig/public` repo carried ~993 stars as of 2026-09 [secondary — GitHub UI snapshot via search index].
- The `featureprofiles` repo (functional test coverage gating backward-incompatible changes) remains the mechanism that decides whether a breaking change ships [official](https://github.com/openconfig/public/blob/HEAD/doc/releases.md).
- **Gap:** the exact 2026 quarterly tag list (e.g. Q1/Q2/Q3 2026 tag names) was not enumerated in the sources found this wave; treat any "latest release" claim without a tag as [unverified] until checked against the releases page [unverified].

## Wave 2 — gNMI protocol: specification, RPCs, and subscription modes

### 2.1 Protocol fundamentals

- gNMI = gRPC Network Management Interface: protocol-buffer encoding over HTTP/2; RPCs are `Capabilities`, `Get`, `Set`, `Subscribe`; data is addressed by YANG-derived paths (e.g. `Set(system/config/hostname, "tor14")`, `Subscribe(interfaces/interface/state/counters)`) [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).
- Subscribe modes: `ON_CHANGE`, `SAMPLED` (periodic), `TARGET_DEFINED` (device chooses); responses carry timestamp, path, and value [secondary](https://files.botwerks.net/presentations/20220922-mnnug-openconfig.pdf).
- Dial-in (collector initiates to device) vs dial-out (device initiates to collector) are deployment patterns: dial-out is favored where devices sit behind NAT/firewalls or where the collector scales horizontally [secondary — widely documented in vendor telemetry guides; specific dial-out scale numbers per vendor in Wave 4].

### 2.2 gnmic — the reference open-source gNMI client/collector (2026)

- `gnmic` (openconfig/gnmic) changelog shows **v0.49.0 released 16 September 2026** — roughly monthly cadence through 2026 [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/changelog.md).
- Install surface: script installer (`https://get-gnmic.openconfig.net`), deb/rpm packages, Docker images (`gnmic/gnmic`, `ghcr.io/openconfig/gnmic`) tagged per release, `latest` pointer; upgrade via `gnmic version upgrade` [official](https://github.com/openconfig/gnmic/blob/HEAD/docs/install.md).
- Cisco IOS-XR (KNE lab example, 8000e): `gnmic capabilities` against a live device reported **gNMI version 0.8.0**, encodings `JSON_IETF`, `ASCII`, `PROTO`; supported models included `openconfig-bgp-types` / `openconfig-bgp-errors` at 5.3.1 [secondary — lab example README](https://github.com/openconfig/kne/blob/HEAD/examples/cisco/8000e/README.md).
- Same lab verified `gnoic` (gNOI System service ping) and `gribic` (gRIBI flush) against IOS-XR 8000 on ports 9337/9340 — evidence of the gNXI suite (gNMI/gNOI/gRIBI) working together on a production NOS [secondary](https://github.com/openconfig/kne/blob/HEAD/examples/cisco/8000e/README.md).

