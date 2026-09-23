---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-15-7-stackstorm-pack-examples-network-relevant
title: "E4.15.7 StackStorm pack examples (network-relevant)"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: ["2026-01-15", "2026-01-27", "2026-01-28", "2026-01-29", "2026-03-23", "2026-05-13", "2026-07-11", "2026-08-19", "2026-09-06", "2026-09-14", "2026-09-22"]
keywords: ["aws", "compute", "latency", "packaging", "research"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [647, 721]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 361efae04d5a9d14ee70db9e694048d279ddbc4d3b5166d543e807c7f1d1255b
---

# E4.15.7 StackStorm pack examples (network-relevant)

### E4.15.7 StackStorm pack examples (network-relevant)

- Packs exist for: network devices (via NAPALM/Netmiko-style actions in community packs), monitoring (Sensu, Nagios), ITSM (ServiceNow, Jira), chat (Slack), cloud (AWS, GCP) — 160+ packs / 6000+ actions total on the Exchange `[official]` — https://github.com/StackStorm/st2.
- Exact pack names/versions change; verify on the Exchange before depending on one `[unverified]`.

### E4.15.8 Timeline: 2026 milestones in this file's scope

- 2026-01-15: IP Fabric 7.9 unified cloud visibility `[vendor-reported]`.
- 2026-01-27/28: netascode nac-test 2.0.0 experimental pyATS integration (changelog window) `[secondary]`.
- 2026-01-29: Forward AI unveiled `[vendor-reported]`.
- 2026-03: Containerlab 0.73 review (git vars, netem events, cEOS certs) `[secondary]`; GNS3 2.2.57 `[official]`.
- 2026-04: Forward AI GA planned `[vendor-reported]`; nac-test 2.0.0 breaking async change `[secondary]`.
- 2026-05-13: CML 2.10 FCS `[official]`; Forward rebrand + Forward Predict `[independent]`.
- 2026-05: ipspace netlab 26.05 (min clab 0.75.0) `[secondary]`.
- 2026-06: EVE-NG v7 (24th), Community EOL; GNS3 3.1.0a3 packaging `[secondary]`.
- 2026-07-11: CML 2.10 release-notes update `[official]`.
- 2026-08-19: Robot Framework Browser 20.4.0 `[secondary]`.
- 2026-09-06: EVE-NG Pro 7.2.0-4 `[official]`.
- 2026-09-14: Robot Framework 7.5 stable `[secondary]`.
- 2026-09-22: this research cutoff.

### E4.15.9 When to use what (one-line rules)

- Proving a change is safe before it ships → Batfish (free) or Forward Predict (commercial twin) `[analysis]`.
- Watching a live network for drift → SuzieQ (OSS/Enterprise) or IP Fabric (commercial intent checks) `[analysis]`.
- Testing automation code itself → pyATS (Cisco depth) / Robot (readability) / pytest (custom) `[analysis]`.
- Reproducible labs in CI → Containerlab; biggest VM topologies → EVE-NG Pro; free VM labs → GNS3; Cisco exam/dev → CML `[analysis]`.
- Reacting to events automatically → EDA rulebooks (Ansible shops) or StackStorm (heterogeneous) `[analysis]`.
- Tying it together → Git as source of truth, plan-on-PR, Batfish pre-gate, SuzieQ post-gate `[analysis]`.

---

*End of Phase E4 — final. Single writer; append-only; no other workspace files modified.*

---

## E4.16 — Final additions

### E4.16.1 EVE-NG Pro feature detail

- RBAC multi-user, user roles, and lab sharing for teams `[secondary]` (cloudmylab comparison).
- Multi-node clustering for distributed labs `[secondary]` (vendor-adjacent comparison).
- Built-in Wireshark capture integration per interface `[secondary]` (cloudmylab).
- Cloud/NAT connectivity and custom topology link presets `[secondary]` (vendor-adjacent).
- 7.2.0-4 Pro build dated 2026-09-06 is the newest verified release `[official]` — https://eve-ng.net.
- Free/Community limitations remain: no multi-user RBAC, limited node types per the freemium model `[secondary]` (blog.cloudmylab.com/eve-ng-community-vs-eve-ng-professional).

### E4.16.2 Containerlab kinds and extensibility (detail)

- Kinds are typed node definitions: SR Linux (Nokia-first), Arista cEOS, Juniper cRPD/vMX/vQFX/vSRX, Cisco XRd/XRv/XRv9k/IOL, SONiC (multiple vendors), FRRouting, VyOS, Linux bridge/host, OVS, and generic `linux` containers `[secondary]` (containerlab SKILL).
- `vrnetlab`-wrapped VM images expose legacy platforms (IOSv, NX-OSv, ASA) as containerized nodes `[secondary]` (containerlab docs ecosystem).
- Startup-config binding, certificate injection (cEOS), and staged startup ordering support realistic multi-vendor labs `[secondary]` (Brian Linkletter review).
- `containerlab deploy --reconfigure` updates running labs without full teardown `[secondary]` (containerlab docs).

### E4.16.3 GNS3 server architecture (detail)

- Split design: gns3-gui (Qt client) + gns3-server (REST API + compute); the server can run local or remote, including headless on a VM `[official]` — https://github.com/GNS3/gns3-gui/blob/master/CHANGELOG.
- Appliances: marketplace of templates (Cisco IOSv, Juniper vMX trial, Arista vEOS, Docker containers); images supplied by the user `[secondary]` (gns3 community docs).
- REST API enables automation of lab lifecycle similar to Containerlab's CLI, though less CI-native `[analysis]`.
- 2.2.57 (2026-03-23) is the newest verified stable in the 2.2 line `[official]` (changelog).

### E4.16.4 nac-test changelog detail

- 2.0.0 (2026-01-28): BREAKING — `pyats` async mode change: synchronous parsing replaced by threaded execution via new `naclabs.test.helpers.parse`; `inventory.hosts` replaced by `naclabs.test.helpers.get_hosts`; pyats helpers fully async; docs restructured (tests, helpers, integrations) `[secondary]` — https://github.com/netascode/nac-test/blob/HEAD/CHANGELOG.md.
- 1.x (2025): added `naclabs.test.robot` namespace (Nov 2025); `test_connectivity` supports FQDN/IPv6 (Dec 2025) `[secondary]` (changelog).
- Status: experimental pyATS integration; Robot keywords stable-ish; production adoption `[unverified]` (§E4.11 O15).

### E4.16.5 Synthetic monitoring (scope note)

- Listed in the task scope; not deep-dived in this wave: candidates include Cisco ThousandEyes, Arista DANZ/DMF synthetic tests, Juniper Paragon Active Assurance — flagged as a future research item `[unverified]` (§E4.11 O13).
- Relationship to this phase: synthetic probes generate the traffic whose loss/latency the validation frameworks assert on; EDA/Kafka pipelines can consume synthetic results as events `[analysis]`.

### E4.16.6 Minimal CI job sketch (illustrative, not run)

```yaml
