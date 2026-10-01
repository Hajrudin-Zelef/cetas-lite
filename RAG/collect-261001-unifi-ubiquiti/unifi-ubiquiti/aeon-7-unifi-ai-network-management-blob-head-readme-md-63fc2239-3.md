---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/aeon-7-unifi-ai-network-management-blob-head-readme-md-63fc2239-3
title: "OpenClaw target"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "containment", "distribution", "incident"]
source: docs/RAG/collect-261001-unifi-ubiquiti/aeon-7-unifi-ai-network-management-blob-head-readme-md-63fc2239.md
source_anchor: ""
source_lines: [243, 308]
sha256: 6f16dd3d6250a2dd25d432d1bba716b4eb27c707a7b5c46fc55b4ef6653020f0
---

# OpenClaw target

| legacy-post /api/s/<site>/cmd/stamgr | Station manager commands | Reconnect/block/unblock a client. | Targeted response to a confirmed rogue or misbehaving client. | 
| legacy-post /api/s/<site>/cmd/devmgr | Device manager commands | Device actions such as provision/restart where supported. | Maintenance only; high-impact confirmation required. | 
See docs/API-ENDPOINTS.md and skill/unifi-api/references/endpoints.md for more detail.
- Find clients with weak signal, low data rates, excessive roaming, or bad AP association.
- Identify overloaded APs or imbalanced client distribution.
- Audit channel width, band use, SSID-to-VLAN mapping, and firmware update state.
- Detect offline or repeatedly reconnecting devices.
- Explain which switch port or AP a client is connected through.
- Build a maintenance plan that upgrades devices in safe batches instead of all at once.
- Identify unknown clients by MAC/IP/hostname/vendor and recent activity.
- Compare connected clients against an expected inventory.
- Detect risky firewall or guest network drift.
- Summarize security events and correlate them with clients, ports, APs, and VLANs.
- Recommend containment steps before action.
- With explicit permission and exact target verification, block or reconnect a suspicious client.
- Produce an incident record: what was observed, what changed, how it was verified, and rollback.
- Generate a human-readable network map from controller data.
- Export a site inventory report.
- Maintain a change log with before/after state.
- Verify that a requested change actually took effect.
- Prepare rollback plans from backups.
The skill instructs the agent to follow this pattern:
- Read current state.
- Identify exact site and object.
- Capture backup for major changes.
- Explain intended endpoint/action.
- Execute the smallest possible write.
- Verify state after the write.
- Report before/after and rollback.
High-impact actions require explicit user intent:
- gateway/AP/switch reboot
- firmware upgrades
- VLAN/DHCP/DNS/firewall changes
- SSID/security changes
- PoE cycling
- delete/adopt/forget/factory reset
- block/unblock clients when target identity is uncertain
.
|-- README.md
|-- setup.sh
|-- env.example
|-- AGENTS.md                 # installation/behavior notes for agent frameworks
|-- scripts/
|   |-- unifi-api-status
|   |-- unifi-api-configure
|   |-- unifi-api-disable
|   |-- unifi-api-enable
|   |-- unifi-config-backup
|   `-- unifi-config-restore
|-- skill/unifi-api/
|   |-- SKILL.md
|   |-- scripts/unifi_api.py
|   |-- references/
|   `-- agents/openai.yaml
`-- docs/
    |-- API-ENDPOINTS.md
    |-- OPSEC.md
    |-- USE-CASES.md
    `-- LOCAL_OWNER_RUNBOOK.md
- Linux or WSL-style shell environment
- Python 3.10+
- Bash
- OpenClaw or Hermes if you want automatic skill installation into their default skill directories
- UniFi Network version with API integrations for official local API access
The API helper intentionally uses Python standard library only.
MIT. Use at your own risk. Network automation can cause outages; read the plan before applying changes.
