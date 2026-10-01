---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-snuffy2-aiopnsense-async-python-client-library-for-opnsense
title: "github-snuffy2-aiopnsense-async-python-client-library-for-opnsense"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-snuffy2-aiopnsense-async-python-client-library-for-opnsense.md
source_anchor: ""
source_lines: [1, 32]
sha256: ca69e2bcc690ef6004124ce7f28263f63b956562df4a42ba263560eb5b0e03c9
---

# github-snuffy2-aiopnsense-async-python-client-library-for-opnsense

`aiopnsense` is an async Python client library for OPNsense.

`aiopnsense` wraps supported OPNsense REST endpoints behind a single async client, `OPNsenseClient`. It is designed for applications that need to query router state or trigger supported OPNsense actions without manually building HTTP requests.

The client currently includes helpers for:

- system information, notices, certificates, CARP, Wake-on-LAN, reboot, and interface reloads
- firmware version checks, update status, and upgrade actions
- interface, gateway, CPU, memory, filesystem, and temperature telemetry
- diagnostics traffic snapshots and live interface traffic stream samples
- DHCP lease and ARP table access
- firewall rules, NAT rules, alias toggling, and state killing
- service status lookup and service start/stop/restart operations
- SMART device listing and per-device SMART detail queries
- Unbound blocklist management
- OpenVPN and WireGuard status plus VPN instance toggling
- vnStat metrics, captive portal vouchers, and speed test data

- 
  - For firmware < 26.1.1, the `Firewall and NAT` methods will return empty data.
- For firmware < 26.1.1, the 

- Read the Docs: https://aiopnsense.readthedocs.io

The `scripts/` directory includes live diagnostic helpers for maintainers:

- `aiopnsense_dump.py` runs supported`aiopnsense` client endpoints and prints the raw data returned by the library.
- `opnsense_api_call.py` calls an arbitrary OPNsense API endpoint with a specified method and optional POST payload.

Both scripts read credentials from `scripts/aiopnsense.env` by default. Copy `scripts/aiopnsense.env.example` to create your local env file; it is ignored by git. Full CLI reference is in the Read the Docs script documentation.

`aiopnsense` was extracted from the hass-opnsense integration. It is primarily used as an external dependency by hass-opnsense and Home Assistant for its OPNsense Integration.
