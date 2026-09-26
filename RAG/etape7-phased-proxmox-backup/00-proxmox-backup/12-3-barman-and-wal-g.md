---
id: etape7-phased-proxmox-backup/00-proxmox-backup/12-3-barman-and-wal-g
title: "12.3 Barman and WAL-G"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-06-24", "2026-07", "2026-07-15", "2026-07-20", "2026-07-21", "2026-08-17"]
keywords: ["distribution"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [218, 237]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 3637e1fad20a0bfb1060639368048cc03b077d9cd70dd36f022a33844762c7e5
---

# 12.3 Barman and WAL-G

- pgBackRest is the reference open-source PostgreSQL backup/restore tool: parallel backup/restore, lz4/zstd compression, AES-256 encryption, multiple repositories (e.g., local fast + remote long-retention), S3/Azure/GCS support, delta restore, page-level checksums, backup info/verify commands `[official]` (https://github.com/pgstef/pgbackrest).
- **Current stable: 2.59.1 (released 2026-08-17)**; 2.59.0 released 2026-07-20; new distribution tarball July 2026 `[official]` (https://github.com/pgstef/pgbackrest).
- Validated against PostgreSQL 18: Percona Distribution for PostgreSQL 18.6.1 bundles pgBackRest 2.59.0 (18.3.1→2.58.0, 18.1.1→2.57.0 progression through 2026) `[secondary]` (https://github.com/percona/postgresql-docs/blob/HEAD/docs/release-notes/release-notes-v18.6.1.md).
- Independent restore-drill evidence: pgdrill ran disposable-container PITR drills with pgBackRest 2.58.0 + PostgreSQL 18.3 (2026-07-21/28), proving timestamp-PITR replay correctness (101 vs 102 rows boundary assertions), while noting the drill does not constitute RTO evidence or object-storage compatibility proof `[secondary]` (https://github.com/r314tive/pgdrill/blob/HEAD/compatibility/evidence/pgbackrest-v2.58.0-postgresql-18.3-linux-amd64-pgdrill-v0.3.0-alpha.10-emulated/README.md).
- Typical topology: dedicated backup host running `pgbackrest` stanza per cluster; `archive-push` via `archive_command`; retention (full/diff/incr) per repository; `pgbackrest --stanza=X check` as a monitoring hook `[secondary]`.

### 12.3 Barman and WAL-G

- Barman (2ndQuadrant/EDB): PostgreSQL backup manager supporting rsync/SSH and streaming (pg_receivewal) methods, WAL archiving, PITR, retention policies, hook scripts, and Barman Cloud for S3-compatible object storage `[secondary]`. 3.x series current; exact 2026 point version not verified `[unverified-current]`.
- WAL-G: Go-based WAL archiver with delta backups, multi-threaded upload, S3 (and compatible) backends, used heavily in Kubernetes Postgres operators; lighter than full backup suites for WAL-shipping pipelines `[secondary]`. Exact 2026 release not verified `[unverified-current]`.
- Selection note: pgBackRest for full-featured on-prem/object backup with parallelism; Barman for EDB-aligned shops and simple rsync flows; WAL-G for cloud-native/K8s WAL pipelines.

### 12.4 MySQL — Percona XtraBackup and mysqldump

- Percona XtraBackup (PXB) is the 100% open-source hot physical backup tool for InnoDB/MyRocks/MyISAM: copies data files while the server runs, then replays redo/undo logs in `--prepare` to produce a consistent snapshot; supports incrementals, compression, encryption and streaming; used under the hood by Percona XtraDB Cluster and Percona Operators `[secondary]` (https://github.com/evgeniypatlan/percona-mysql-skills/blob/HEAD/percona-xtrabackup/SKILL.md).
- **Hard version rule: PXB major version must equal the server major version** — PXB 8.4 backs up MySQL/Percona Server 8.4 only (not 8.0, not 9.x); minor versions need not match within a major; `--no-server-version-check` can override but risks corrupt backups `[secondary]` (https://github.com/evgeniypatlan/percona-mysql-skills/blob/HEAD/percona-xtrabackup/SKILL.md).
- Releases 2026: **PXB 8.4.0-6 released 2026-06-24** (for Percona Server/MySQL 8.4.x); **PXB 9.7.1-rc1 released 2026-07-15** (release candidate for MySQL 9.7 LTS; new `--check-tables` InnoDB B-tree validation during `--prepare`, backup-size reporting in `xtrabackup_info`, sparse-file handling improvements) `[secondary]` (https://github.com/percona/docs-home/blob/HEAD/docs/new/posts/MySQL/percona-xtrabackup-8.4.0-6-release.md).
- `mysqldump`: logical backup (SQL text); `--single-transaction` for consistent InnoDB dumps without long locks; slow to restore at scale — suited to small DBs, migrations and logical exports, not large-scale DR `[secondary]` (https://documentation.commvault.com/2024e/expert/mysql_backup_using_percona_xtrabackup.html).
- MariaDB note: `mariabackup` (fork lineage) is the native physical tool; PXB supports MariaDB 10.x/11.x via symlinks per Commvault's matrix — verify against MariaDB's own docs before relying on it `[secondary]` (https://documentation.commvault.com/2024e/expert/mysql_backup_using_percona_xtrabackup.html).

