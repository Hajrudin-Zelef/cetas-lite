---
id: etape7-phased-proxmox-backup/00-proxmox-backup/13-ransomware-proof-backup-design-2026
title: "13. Ransomware-proof backup design (2026)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["Broadcom"]
dates: []
keywords: ["apache", "cost", "cyber", "governance", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [238, 284]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 78788b7c18c8c6a1319b61e409a0d45774961d76c113b089f385c7d702231b3d
---

# 13. Ransomware-proof backup design (2026)

## 13. Ransomware-proof backup design (2026)

- **3-2-1 rule** (baseline): 3 copies of data, on 2 different media/types, 1 off-site; field-evolved to **3-2-1-1-0** (one immutable/offline copy, zero errors after verification) — Veeam's current best-practice framing `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- **Immutability options 2026:**
  - S3 Object Lock (compliance/governance mode) as backup target — Veeam, PBS (native S3 datastores), restic/kopia via S3 backends; enable at bucket creation (cannot be retrofitted) `[secondary]` (https://github.com/davperez-tech/backup-disaster-recovery).
  - Hardened Linux repositories with immutability (Veeam hardened repo; XFS immutability flag patterns) `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
  - Tape (LTO-5+ with autoloaders in PBS) as the true air gap — offline media cannot be encrypted by ransomware `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
  - Client-side encryption (PBS AES-256-GCM, Borg AES-256/HMAC-SHA256, restic) so stolen backup data is unreadable — encryption complements, not replaces, immutability `[official]` (https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/).
- **Separation of control planes:** backup infrastructure on separate credentials/IdP from production; dedicated backup admin accounts; MFA on backup consoles (Veeam v13 enforces MFA-capable flows); backup network segmentation so a domain compromise does not reach repositories `[secondary]`.
- **Malware-aware recovery:** scan backups before restore — Veeam v13 proactive investigation auto-scans suspicious objects and marks clean backups; inline entropy + filesystem-activity analysis catch crypto-ransomware patterns in backup streams `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- **Secure restore environments:** isolated "clean room" VLAN for restoring suspect VMs; verify before reconnecting to production; documented in Veeam's cyber-resilience guidance `[secondary]`.
- **Backup of the backup catalog:** protect the backup server/catalog/config itself (PBS datastore + config backups; Veeam configuration backup); a ransomware event that destroys the catalog destroys recoverability `[secondary]`.
- **Insider/misconfiguration risk:** immutability retention windows must exceed the maximum dwell time you plan for (common 14–30 days minimum; regulated estates longer); short Object Lock windows give false confidence `[secondary]`.

## 14. RTO/RPO guidance and restore testing

- Definitions: RPO = maximum acceptable data loss (backup frequency); RTO = maximum acceptable downtime (restore speed). Every backup design should state both per workload tier `[secondary]`.
- Tiering example: Tier 0 (RPO 15 min, RTO 1 h — databases with WAL archiving + standby); Tier 1 (RPO 1 h, RTO 4 h — critical VMs with hourly incrementals); Tier 2 (RPO 24 h, RTO 24 h — file servers, daily); Tier 3 (RPO 7 d — archives) `[secondary]`.
- Restore-speed levers: local fast copy for RTO (Tier 1/0), off-site immutable copy for survivability; PBS live restore (boot VM during restore); Veeam Instant Recovery; ZFS replication for near-instant failover `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Database RTO/RPO: PostgreSQL streaming replica (RPO ~seconds) + pgBackRest PITR (RPO minutes, RTO = restore+replay time); MySQL replica + XtraBackup `--prepare` restore; test replay time on production-sized data — logical dumps that restore in 6 hours fail a 1-hour RTO `[secondary]`.
- **Restore testing is the backup:** scheduled verification (PBS scheduled verification jobs; `borg check`; `restic check`; `pgbackrest check`), plus full DR drills (bare-metal/VM restore into isolated network, application smoke tests, timed runbooks). The pgdrill project publishes machine-checkable restore evidence (pass/fail JSON, SHA-256 manifests) as a model for auditable drills `[secondary]` (https://github.com/r314tive/pgdrill/blob/HEAD/compatibility/evidence/pgbackrest-v2.58.0-postgresql-18.3-linux-amd64-pgdrill-v0.3.0-alpha.10-emulated/README.md).
- Metrics to track: backup success rate, backup window duration, verify/check pass rate, last successful restore drill per tier, time-to-first-byte on restore, deduplication ratio trends `[secondary]`.

## 15. Comparison matrices

### 15.1 Virtualization platforms (2026 snapshot)

| Platform | Model | HA/Clustering | Storage | Containers | Management plane | Approx. cost signal |
|---|---|---|---|---|---|---|
| Proxmox VE 9.2 | Debian+KVM/LXC, AGPLv3 | Multi-master, HA w/ fencing, dynamic load balancer, arm/disarm | Ceph (Squid/Tentacle), ZFS 2.4, LVM snapshots, PBS | LXC 7.0 native | Web UI/CLI/API + PDM 1.1 | €370–1,100/socket/yr support (features free) `[official]` |
| XCP-ng 8.3 + XO | Xen, open core | Pool HA, live migration | XOSTOR HCI option, shared SR | Via VMs | Xen Orchestra / XO Lite | VMS Pro €1,000/host/yr (2024 figures) `[secondary]` |
| Harvester v1.7.x | K8s+KubeVirt+Longhorn | K8s-native HA | Longhorn distributed block | Native (K8s) + KubeVirt VMs | Rancher-embedded UI | Free; SUSE support via Rancher Prime (price n/a) `[secondary]` |
| VMware VCF (Broadcom) | Proprietary | vSphere HA/vMotion | vSAN | Tanzu | vCenter | Per-core subscription — the cost driver for 2026 migrations `[secondary]` |

### 15.2 Backup software (2026 snapshot)

| Tool | Type | Dedup | Encryption | Immutability | Targets | Cost |
|---|---|---|---|---|---|---|
| Proxmox Backup Server 4.2 | VM/CT/host image+file | Global, fixed/variable chunks | AES-256-GCM client-side | S3 Object Lock, tape air gap | Local, S3, tape, remote sync | Free; €560–4,480/server/yr support `[official]` |
| Veeam B&R v13 | Enterprise suite | Yes | At rest/in transit | Default-on, hardened repo, S3 Lock | Disk, object, tape, cloud | Commercial per-workload; CE free ≤10 workloads `[secondary]` |
| BorgBackup 1.4.x | File-level dedup | Content-defined, global | AES-256/HMAC-SHA256 | Via target (append-only mode) | SSH, local | Free (BSD) `[official]` |
| Restic 0.18.x | File-level dedup | Content-defined | Authenticated encryption | Via S3 Object Lock | S3/SFTP/rest-server/rclone/local | Free (BSD) `[secondary]` |
| Kopia | File-level dedup | Content-addressed | Client-side | Via S3 Object Lock | S3/cloud/local | Free (Apache-2.0) `[secondary]` |
| Bacula/Bareos | Enterprise suite | Varies | Yes | Tape air gap, catalog ACLs | Disk, tape, cloud | Free; commercial support available `[secondary]` |
| pgBackRest 2.59.1 | Postgres physical | Block-level incremental | AES-256 | Via S3 Object Lock | Local, S3/Azure/GCS | Free (MIT) `[official]` |
| Percona XtraBackup 8.4 | MySQL physical hot | Incremental pages | Yes | Via target | Local, streaming, cloud | Free (GPL) `[secondary]` |

