---
id: etape7-phased-proxmox-backup/00-proxmox-backup/24-database-backup-operations-deep-dive-supplement
title: "24. Database backup operations deep-dive (supplement)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-09-01"]
keywords: ["datacenter", "governance", "latency", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [470, 528]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 3e9ce91a26507d1cdaae26ef3191c293bec4eeb4f62de7bf7b9bf8314cff6582
---

# 24. Database backup operations deep-dive (supplement)

## 24. Database backup operations deep-dive (supplement)

- pgBackRest stanza setup: `pgbackrest --stanza=main stanza-create`; `archive_command = 'pgbackrest --stanza=main archive-push %p'`; verify with `pgbackrest --stanza=main check` (expects WAL archiving + DB connectivity) `[secondary]`.
- pgBackRest backup types: full / differential / incremental; `--type=incr` default daily with weekly full is a common pattern; block-level incremental on 2.59+ reduces transfer for large DBs `[official]` (https://github.com/pgstef/pgbackrest).
- pgBackRest repos: repo1 local (fast restore), repo2 S3 (long retention, off-site); per-repo retention (`--repo1-retention-full=2`, `--repo2-retention-full=52`) `[official]` (https://github.com/pgstef/pgbackrest).
- pgBackRest restore: `--delta` restores only changed files for speed; `--target-action=promote` for PITR to a timestamp (`--type=time --target="2026-09-01 12:00:00"`); timeline handling automatic `[official]` (https://github.com/pgstef/pgbackrest).
- Encryption at rest: `--repo-cipher-type=aes-256-cbc` with passphrase in a secured cipher file; separate from transport TLS `[official]` (https://github.com/pgstef/pgbackrest).
- Patroni + pgBackRest: Patroni can drive base backups via pgBackRest; replicas bootstrap from pgBackRest (`--standby`); keep `archive_command` consistent across failover — use a shared stanza or per-node stanzas with caution `[secondary]`.
- Barman setup: `barman receive-wal` (streaming) or `archive_command = 'rsync ... barman@host:...'`; `backup_method = rsync` or `postgres`; `retention_policy` and `wal_retention_policy`; `barman check` for preflight `[secondary]`.
- Barman Cloud: `barman-cloud-backup` / `barman-cloud-restore` / `barman-cloud-wal-archive` for S3-compatible object storage without a Barman server `[secondary]`.
- WAL-G usage: `wal-g backup-push`, `wal-g wal-push` (as `archive_command`), `wal-g backup-fetch LATEST`, `wal-g wal-fetch`; delta-from backup chains; LZ4/zstd/br compression options; S3/GCS/Azure backends `[secondary]`.
- WAL-G in Kubernetes: standard sidecar in Zalando/CloudNativePG-style operators; object storage as the durability layer; PITR via `wal-g backup-fetch --target-time` `[secondary]`.
- PostgreSQL PITR procedure: stop Postgres, restore base backup, configure `restore_command` / `recovery_target_time`, start in recovery, promote; always restore into a *new* data directory or host — never overwrite production in place `[secondary]`.
- RPO math for Postgres: WAL archiving every segment (16 MB) or `archive_timeout` (e.g., 5 min) bounds data loss; synchronous standby bounds it to ~zero for Tier 0 `[secondary]`.
- MySQL XtraBackup flow: `xtrabackup --backup --target-dir=...` (hot copy), `--prepare` (crash recovery), `--copy-back` or `--move-back`, fix permissions, start MySQL; incremental with `--incremental-lsn` `[secondary]` (https://github.com/evgeniypatlan/percona-mysql-skills/blob/HEAD/percona-xtrabackup/SKILL.md).
- MySQL PITR: XtraBackup full/incremental + mysqlbinlog replay from `--binlog-pos`; GTID-based replication simplifies replica rebuilds `[secondary]`.
- mysqldump for small DBs: `mysqldump --single-transaction --routines --triggers --events`; restore is single-threaded SQL replay — parallel restore via `mydumper/myloader` for speed `[secondary]`.
- Backup validation for DBs: automated restore to a staging host + `pg_amcheck`/`mysqlcheck` + application smoke queries; the pgdrill evidence format (JSON reports + SHA-256 manifests) is a reusable pattern `[secondary]` (https://github.com/r314tive/pgdrill/blob/HEAD/compatibility/evidence/pgbackrest-v2.58.0-postgresql-18.3-linux-amd64-pgdrill-v0.3.0-alpha.10-emulated/README.md).
- Common failure modes: WAL archiving silently broken (monitor `pg_stat_archiver` failed counts); XtraBackup version mismatch across major versions; replica lag masking as "backup OK"; untested `--prepare` on encrypted backups `[secondary]`.

## 25. Proxmox VE security hardening checklist (supplement)

- Change default `root@pam` password immediately after install; create named admin users instead of sharing root `[secondary]`.
- Enable two-factor authentication (TOTP) for all privileged accounts `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Restrict web UI (port 8006) to a management VLAN/VPN; never expose it to the internet `[secondary]`.
- Use the datacenter firewall to default-drop inbound to hosts; allow only management subnets `[secondary]`.
- Separate networks: management, Corosync/cluster (low latency, isolated), storage/Ceph (high bandwidth), VM traffic `[secondary]`.
- Corosync on its own NIC/VLAN; multicast or unicast per version guidance; latency spikes cause fencing storms `[secondary]`.
- Disable password auth for SSH; key-only + fail2ban or port knocking for any exposed SSH `[secondary]`.
- API tokens with expiration and least-privilege for automation; never embed root passwords in scripts `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Keep the enterprise repository enabled on production; no-subscription repo lags on security fixes `[secondary]`.
- Subscribe to the Proxmox security announcements mailing list / RSS `[secondary]`.
- Enable Secure Boot on hosts where hardware supports it; use vTPM + Secure Boot for Windows 11 guests `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Prefer unprivileged LXC containers; drop capabilities; use nesting only when required `[secondary]`.
- Apply AppArmor/SELinux profiles where applicable; Proxmox ships LXC AppArmor profiles `[secondary]`.
- Encrypt VM disks at rest where required (LUKS inside guest or storage-level encryption); PBS backups are encrypted client-side anyway `[secondary]`.
- Backup the Proxmox host configs (`/etc/pve`, `/etc/network/interfaces`, firewall rules) — cluster config lives in pmxcfs but node-local files do not `[secondary]`.
- Document the qdevice for 2-node clusters; losing quorum during maintenance is the classic outage `[secondary]`.
- Test fencing: a fence device that cannot actually fence is worse than none — verify IPMI credentials work `[secondary]`.
- Snapshot before upgrades; have a rollback plan for kernel/QEMU bumps `[secondary]`.
- Limit who can use PCIe passthrough (it bypasses IOMMU isolation boundaries if misconfigured) `[secondary]`.
- Audit logs: review task log and auth logs; forward syslog to a central collector `[secondary]`.
- Disable unused services on hosts (e.g., no extra daemons on hypervisors) to shrink attack surface `[secondary]`.

## 26. Backup scheduling and retention policy examples

- Bronze (file servers): daily incremental, weekly full, 30-day retention, off-site weekly copy `[secondary]`.
- Silver (app VMs): daily incremental with dirty-bitmap CBT, 14-day local retention, 90-day S3 immutable copy `[secondary]`.
- Gold (databases): pgBackRest/XtraBackup daily full or weekly full + daily diff, WAL/binlog continuous, 7-day local + 1-year S3/archive tier `[secondary]`.
- PBS prune example: keep-last=3, keep-daily=7, keep-weekly=4, keep-monthly=6, keep-yearly=1 — tune per datastore `[secondary]`.
- Restic forget example: `--keep-daily 7 --keep-weekly 5 --keep-monthly 12 --keep-yearly 3` then `--prune` `[secondary]`.
- Borg prune example: `--keep-daily=7 --keep-weekly=4 --keep-monthly=6` `[secondary]`.
- Grandfather-father-son (GFS): daily son, weekly father, monthly grandfather — classic tape rotation still valid for LTO pools `[secondary]`.
- Backup windows: stagger start times; PBS fleecing avoids production I/O contention; verify jobs run outside backup windows `[secondary]`.
- Full-backup frequency vs restore speed: more frequent fulls = faster restores but more storage; synthetic fulls (Veeam) merge incrementals server-side `[secondary]`.
- Retention vs immutability: Object Lock duration must cover the retention you promise; extending retention later does not extend existing locks `[secondary]`.
- Legal hold: compliance mode Object Lock for regulated data; governance mode allows privileged deletion with bypass permission `[secondary]`.
- Test restores on a cadence: monthly automated, quarterly full DR drill, annual tabletop with stakeholders `[secondary]`.

