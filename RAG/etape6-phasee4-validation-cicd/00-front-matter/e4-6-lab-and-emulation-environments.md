---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-6-lab-and-emulation-environments
title: "E4.6 — Lab and emulation environments"
domain: front-matter
role: reference
task: reference
actors: ["Apple"]
dates: ["2026-03", "2026-05", "2026-05-13", "2026-06", "2026-06-23", "2026-07-11", "2026-09", "2026-09-06"]
keywords: ["claude", "compute", "license", "licenses", "mcp", "mit license", "model context protocol", "open source", "research"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [189, 244]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: eec8115f17f72398ce74a14365f64b090fc095795a50b56a013dc7273c78097c
---

# E4.6 — Lab and emulation environments

## E4.6 — Lab and emulation environments

### E4.6.1 Containerlab (SRL Labs / Nokia)

- Containerlab: open-source, container-based network lab orchestration (MIT license `[secondary]`); topologies defined in `.clab.yml` YAML; lifecycle commands deploy/inspect/destroy/graph/save `[secondary]` — https://github.com/chrishuffman5/domain-expert/blob/HEAD/./plugins/networking/skills/containerlab/SKILL.md.
- Supported NOS "kinds": Nokia SR Linux, Arista cEOS, Cisco XRd, SONiC VS/VM, Juniper cRPD/vMX, FRR, Linux; vrnetlab wraps VM-based images (vMX, XRv9K, CSR1000v) inside containers `[secondary]` (same source).
- Release 0.73 (reviewed March 2026, brianlinkletter.com): git variables in topology names (`__gitBranch__`, `__gitHash__` for CI/CD), netem events streaming via `events` command, Arista cEOS TLS certificate auto-loading for gRPC/RESTCONF `[secondary]` — https://github.com/srl-labs/containerlab/blob/HEAD/docs/rn/0.73.md and https://brianlinkletter.com/2026/03/containerlab-network-emulator-v0-73-review/.
- ipspace netlab 26.05 (May 2026) requires minimum Containerlab 0.75.0 `[secondary]` — https://github.com/ipspace/netlab/blob/HEAD/docs/release/26.05.md.
- Link impairment built in: `clab tool netem` sets speed/corruption/delay/jitter/loss `[secondary]` (brianlinkletter.com review).
- Graph outputs: HTML, draw.io, Mermaid, Graphviz DOT; VS Code extension; WSL2/macOS support `[secondary]` (skill source).
- CI/CD integration: GitHub Actions, GitLab CI for automated network testing `[secondary]` (skill source).
- Image licensing caveat: cEOS requires Arista download entitlement; XRd requires Cisco entitlement `[secondary]` (skill source).

### E4.6.2 EVE-NG (Emulated Virtual Environment)

- EVE-NG Professional 7.2.0-4 released 6 September 2026 — latest as of research date `[official]` — https://eve-ng.net.
- EVE-NG v7 released 24 June 2026 on Ubuntu 24.04 base with new installer and upgrade assistant from v6 `[secondary]` — https://blog.cloudmylab.com/eve-ng-community-vs-eve-ng-professional.
- Community Edition reached End of Life / End of Support June 2026 (ISO still downloadable, no further updates) `[secondary]` (same source).
- Freemium = Professional ISO running without license: 7 nodes, unlimited VPCS, admin-only access `[secondary]` (same source). The free path shrank from 63 nodes (old Community) to 7 `[secondary]`.
- Professional features: up to 1,024 nodes per lab; user roles (Admin/Editor/User); AD/RADIUS integration; MFA; hot-add links; link-quality simulation; dynamic TCP ports; EVE Cluster for multi-host labs; full APIs for Ansible/Python automation; Docker support (Wireshark etc.) `[secondary]` — https://blog.cloudmylab.com/elevate-networking-skills-explore-eve-ng-pro?hs_amp=true.
- 7.2.0-4 highlights: high-performance dataplane ("New Silicon"), large-lab scalability to 1024 nodes, faster Lab API/node startup, FRR Docker template, traffic filters and packet capture on links/hubs/clouds, QEMU startup-config inject/export on satellites, improved Docker provisioning, MAC prefix options `[official]` — https://www.eve-ng.net/index.php/documentation/release-notes/.
- Hardware sizing: 16GB+ RAM minimum recommended, NVMe, server-class CPU for large labs; images must be legally obtained by the user `[secondary]` (cloudmylab).
- Apple Silicon: not officially supported (images are x86-64) `[secondary]` — https://it-proacademy.com/eve-ng-download-install-guide-2024/.

### E4.6.3 GNS3

- GNS3 2.2.57 released 23 March 2026 (stable 2.2.x line): error reporting fixes, IOU defaults, SuperPuTTY VNC support, dependency upgrades (pytest, jsonschema, sentry-sdk, psutil) `[official]` — https://github.com/GNS3/gns3-gui/blob/master/CHANGELOG.
- 2.2.56 (Jan 2026): PyQt6 migration, XDG Config Home support; 2.2.55 (Nov 2025): Python 3.14 support `[official]` (same changelog).
- 3.x in development: 3.1.0a5 all-in-one seen on onworks.net; AUR package gns3-server 3.1.0a3-2 updated 2026-06-23 `[secondary]` — https://aur.archlinux.org/packages/gns3-server.
- Architecture: split GUI + server model; appliance library broad including niche devices; community MCP/automation tooling emerging (gns3-lab-automation skill v0.34.0 with 27 action tools, 21 MCP resources) `[secondary]` — https://github.com/majiayu000/claude-skill-registry/blob/HEAD/skills/data/gns3-lab-automation/SKILL.md.
- Positioning vs EVE-NG: EVE-NG single-VM browser-based vs GNS3 split GUI/engine; comparisons note EVE-NG easier to maintain, GNS3 broader appliance library for niche devices `[secondary]` — https://blog.cloudmylab.com/eve-ng-vs-gns3.

### E4.6.4 Cisco Modeling Labs (CML)

- CML 2.10 is the latest feature release (FCS May 13, 2026; release-notes doc updated July 11, 2026) `[official]` — http://developer.cisco.com/docs/modeling-labs/cml-release-notes/.
- New in 2.10: Model Context Protocol (MCP) Server integration for AI tools, packet capture streaming, lab autostart, node staging improvements; reference platform images now include XRd, Meraki vMX, Snort3, wireless images; experimental wireless networking and RADIUS user authentication `[official]` (same source).
- CML 2.8 notes (community coverage): FTD/FMC (Firepower) native support with day-zero config; enhanced node resource management (CPU/RAM/disk per node); "CML 2.8 is now free" per pinglabz (free tier positioning) `[secondary]` — https://www.pinglabz.com/cisco-modeling-labs-cml-2-8-free/.
- CML 2.7: IOL VM images, SD-WAN VM images (no additional licensing needed for included images), Azure install support, node pinning to compute (multi-user licenses) `[secondary]` — https://ciscolearningservices.my.site.com/cln/s/question/0D56e0000Dp1bCkCQI/cisco-modeling-labs-cml-270-release-is-available.
- Platform: KVM hypervisor, HTML5 UI, API-first REST design, labs persistent by default, custom simulation fabric, Breakout Tool for console multiplexing `[official]` — https://developer.cisco.com/docs/modeling-labs/2-7/overview-of-cml-2-x/.
- Upgrade constraint: cannot upgrade directly to 2.10 from 2.7.1 or older; must be on 2.7.2+ first; in-place upgrades from 2.2.3 or earlier unsupported (migration required) `[official]` (2.10 release notes).

### E4.6.5 Lab platform comparison (synthesis)

| Dimension | Containerlab | EVE-NG Pro | GNS3 | CML 2.10 |
|---|---|---|---|---|
| Model | Containers | VMs (+Docker) | VMs + Docker | VMs (KVM) |
| License | Open source (MIT) | Commercial (freemium 7 nodes) | Open source (GPL-3.0) | Commercial (free tier per 2.8 notes) |
| Latest (Sept 2026) | ≥0.75.0 | 7.2.0-4 (2026-09-06) | 2.2.57 stable / 3.1 alpha | 2.10 (2026-05-13) |
| Multi-vendor | Yes (kinds) | Yes (images) | Yes (appliances) | Cisco-focused + custom import |
| CI/CD fit | Native (YAML, git vars) | API-driven | API/MCP tooling emerging | REST API-first |
| AI angle | — | — | MCP skill (community) | MCP Server (official) |

- All version claims above carry the provenance of their rows; the table itself is `[analysis]` synthesis.

---

