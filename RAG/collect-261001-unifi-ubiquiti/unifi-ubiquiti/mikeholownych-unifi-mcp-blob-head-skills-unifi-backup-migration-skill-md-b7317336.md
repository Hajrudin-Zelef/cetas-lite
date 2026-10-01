---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/mikeholownych-unifi-mcp-blob-head-skills-unifi-backup-migration-skill-md-b7317336
title: "mikeholownych-unifi-mcp-blob-head-skills-unifi-backup-migration-skill-md-b7317336"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/mikeholownych-unifi-mcp-blob-head-skills-unifi-backup-migration-skill-md-b7317336.md
source_anchor: ""
source_lines: [1, 22]
sha256: bcd5dfa95c15507f8ec20cff5c5c076f837f3c76da67728785901f5a42f7c863
---

# mikeholownych-unifi-mcp-blob-head-skills-unifi-backup-migration-skill-md-b7317336

| name | unifi-backup-migration | 
|---|---|
| description | Use when backing up UniFi controller settings, exporting configuration, preparing a gateway replacement or migration, checking if backups exist, or recovering a dead console. | 
| argument-hint |  | 
| disable-model-invocation | true | 
| backup \| migrate \| check | 
Protect the configuration investment: verify backups exist, create them, understand what they contain, and move between consoles without surprises.
- ✅ Sites, networks/VLANs, WLANs, firewall policies/groups, port profiles, device provisioning data, admin list, hotspot config.
- ❌ Time-series statistics/history beyond short retention, UniFi Protect recordings (Protect has its OWN backup/restore path), client history depth varies. Say this explicitly when someone asks for "a full backup".
- UI path: Settings → System → Backups → Create Backup (choose to include statistics only if explicitly wanted — large files).
- Verify recency: the UI lists existing backups w/ timestamps. If none within ~30 days, make one now.
- Download the .unf file OFF the console to another disk — a backup stored
only on the dying box is not a backup.
- Target console should run the same or newer Network application major version as the export; cross-version restores fail or degrade silently.
- Same-model swap: restore → adopt devices → done (devices re-appear because provisioning info travels).
- Different site/controller: devices must be re-adopted; SSH adoption or the "Set Inform" step may be needed — expect it, don't panic when APs show pending.
- Protect/NVR: migrate separately via Protect's own backup or disk moves.
Save get_networks, get_wlans, get_firewall_policies, list_devices JSON
to ./unifi-backups/<date>-pre-migration.json — invaluable if zone/policy
mapping gets lost in translation.
- Never store credentials in exported files shared with others (.unf contains config, but treat it as sensitive).
- Confirm-gate any destructive step during recovery discussions.
