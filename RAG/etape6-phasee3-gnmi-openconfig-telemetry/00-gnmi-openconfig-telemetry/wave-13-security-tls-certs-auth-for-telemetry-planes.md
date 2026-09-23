---
id: etape6-phasee3-gnmi-openconfig-telemetry/00-gnmi-openconfig-telemetry/wave-13-security-tls-certs-auth-for-telemetry-planes
title: "Wave 13 — Security: TLS, certs, auth for telemetry planes"
domain: step-6-phase-e3-gnmi-openconfig-model-driven-telemetry-progr
role: deep-dive
task: reference
actors: ["Broadcom", "Huawei", "Nvidia"]
dates: []
keywords: ["agent", "asic", "benchmarks", "distribution", "latency", "nvidia"]
source: docs/RAG/etape6_phaseE3_gnmi_openconfig_telemetry.md
source_anchor: ""
source_lines: [304, 376]
section: "Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability"
sha256: 7614590f17c2c6859137a201751298df55ddf0052314403a463d54b0df681088
---

# Wave 13 — Security: TLS, certs, auth for telemetry planes

## Wave 13 — Security: TLS, certs, auth for telemetry planes

### 13.1 Transport security

- gNMI/gRPC runs over TLS by default on production NOS; lab shortcuts (`--skip-verify`, `insecure`, `skip_verify`) are pervasive in docs but must not be cargo-culted into production [secondary — lab configs across Waves 3–7].
- Dell OS10 gNMI agent: TLS verification + username/password, and **only Super Admin roles** are accepted — authentication failure means total rejection, not degraded access [official](https://www.dell.com/support/manuals/en-ie/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-2-6/gnmi-agent?guid=guid-9d4c66d3-178d-4122-baed-6914977fe93e&lang=en-us).
- Juniper dial-in: client certificate + key required when TLS enabled [official](https://www.juniper.net/documentation/us/en/software/junos/interfaces-telemetry/interfaces-telemetry.pdf).

### 13.2 Certificate lifecycle via gNOI

- gNOI Cert service (`Install`, `Rotate`, `LoadCertificate`) is the programmatic path for device certificate lifecycle — directly relevant to keeping telemetry TLS trust fresh at scale [official](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1718/b-1718-programmability-cg/gnoi.pdf).
- Extreme ONE OS extends this with `auth.proto` (JWT auth + certificate generation) [official](https://documentation.extremenetworks.com/Extreme%20ONE%20OS%20Switching%20v22.2.0.0%20GNOI%20Reference%20Guide/Extreme_ONE_OS_Switching_22_2_0_0_GNOI_Reference_Guide.pdf).

### 13.3 Exposure hygiene

- Telemetry ports (50051 Junos, 6030 Arista, 57400 SR Linux) and dial-out destinations are management-plane attack surface; recommended posture: dedicated management VRF/ACLs, mTLS, credential rotation via gNOI — **operational guidance synthesized from vendor docs, not a single cited source** [analysis].
- NetBox API tokens for target discovery stored as Kubernetes Secrets (gnmic operator pattern) — the same secret-hygiene applies [official](https://github.com/gnmic/operator/blob/HEAD/docs/content/docs/examples/NetBox/REST%20API/_index.md).

## Wave 14 — Adjacent tooling and cross-references

### 14.1 Cross-references to sibling Phase E files

- **SuzieQ** (E4 file): complementary poller-based observability — SuzieQ's value is normalized multi-vendor state tables, while gNMI streaming gives sub-second event latency; shops run both [secondary].
- **Batfish** (E4 file): config *analysis* (what will the network do) vs telemetry *observation* (what is it doing) — the CI/CD loop closes when Batfish pre-change validation and gNMI post-change telemetry share the same source of truth [secondary].
- **NetBox/Nautobot** (E1 file): inventory → target discovery (Wave 7) and topology overlay (Wave 7.3).

### 14.2 Nautobot apps touching telemetry

- **Gap:** Nautobot app ecosystem specifics for streaming telemetry (Golden Config telemetry hooks, etc.) are covered in the E1 file; not duplicated here.

### 14.3 In-band telemetry / INT (out-of-scope note)

- In-band Network Telemetry (INT), IOAM, and Broadcom/ASIC-level congestion signals are the *data-plane* complement to gNMI's *management-plane* streaming; 2026 AI-fabric congestion visibility increasingly combines both. Dedicated coverage deferred — open item [unverified].

## Wave 15 — Final audit: scope coverage, gaps, and close

### 15.1 Scope-to-wave map

| Original scope item | Covered in |
|---|---|
| OpenConfig 2026 models, WG status, vendor adoption | Waves 1, 10 |
| gNMI/gNOI protocol status, dial-in vs dial-out, subscription modes | Waves 2, 6.1, 11 |
| Vendor implementations | Waves 3, 12 |
| Open-source collectors (Telegraf, Prometheus exporters, pmacct, custom) | Waves 4, 7.1–7.2 |
| Streaming pipeline architectures (Kafka, TimescaleDB/ClickHouse, Grafana) | Wave 5 |
| Real-time topology, LLDP mapping, NetBox/Nautobot integration | Wave 7 |
| NETCONF/RESTCONF vs gNMI 2026, YANG Suite, pygnmi | Wave 6 |
| AI-era telemetry (RoCE counters, PFC/ECN, congestion) | Wave 8 |
| Security of telemetry planes | Wave 13 |

### 15.2 Consolidated open items (all flagged inline as [unverified] or gaps)

1. 2026 quarterly OpenConfig tag list (Q1/Q2/Q3 2026 tag names) not enumerated.
2. Arista TerminAttr/CloudVision dial-out specifics and 2026 evolution.
3. SONiC-distribution gNMI/OpenConfig maturity (Dell Enterprise SONiC, NVIDIA Cumulus, community).
4. Telegraf `inputs.gnmi` version matrix; pygnmi 2026 status; pmacct gNMI role.
5. NVIDIA Spectrum-X gNMI/Redfish AI-fabric telemetry specifics.
6. Per-vendor gNMI path inventories for PFC/ECN/buffer counters.
7. gNMI history/depth extension device support.
8. NETCONF vs gNMI deployment-share survey (2026).
9. YANG Suite 2026 release status.
10. In-band telemetry / INT adoption 2026.
11. Huawei gNMI integration resolution status (gnmic issue #451).
12. Quantitative dial-in vs dial-out scale benchmarks per vendor.

### 15.3 Provenance audit

- All facts carry one of `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`, or `[analysis]` (synthesis clearly labeled).
- No SKUs, URLs, or numbers guessed; all URLs verbatim from search results.
- Conflicts registered: (C1) Arista 4.33.1F lab image gNMI quirk vs general EOS gNMI support; (C2) SONiC lossless priorities (3,4) vs per-platform defaults; (C3) OpenConfig model versions in the wild lagging HEAD by years.

*End of Phase E3 — 24 waves, append-only, single writer. All scope items covered; 12 open items logged in Wave 15.2 (see supplementary audit in Wave 24.2).* Note: the Wave 15 marker above was written when the file held 15 waves; Waves 16–24 extend it.

