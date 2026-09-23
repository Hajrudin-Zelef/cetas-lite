---
id: etape7-phased-proxmox-backup/00-proxmox-backup/11-open-source-backup-tools
title: "11. Open-source backup tools"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-06-24", "2026-07", "2026-07-15", "2026-07-20", "2026-07-21", "2026-08-17", "2026-09-09"]
keywords: ["agent", "distribution", "license", "pricing", "robotics"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [173, 237]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 319f5b1e95a1494bc1ae17f1f6094fd560e3775e8c942b5fa7876f0094d1e646
---

# 11. Open-source backup tools

## 11. Open-source backup tools

### 11.1 BorgBackup (Borg)

- BorgBackup is a deduplicating, compressing, authenticated-encryption backup program; BSD-3-Clause license; main goal is efficient and secure daily backups `[official]` (https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/).
- Current stable line: **1.4.x** — official docs at release 1.4.5; FreeBSD port at 1.4.4_2 updated 2026-09-09 `[official/secondary]` (https://www.freshports.org/archivers/py-borgbackup/).
- Deduplication: content-defined chunking (Buzhash rolling hash in 1.4.x; newer development trees move to FastCDC/Gear as default) — global across all archives and all hosts sharing a repository, on data and metadata; unaffected by file renames, timestamp changes or shifted offsets; only changed chunks of large files (VMs, raw disks) are stored `[official]` (https://github.com/borgbackup/borg/blob/HEAD/docs/internals.rst).
- Security: 256-bit AES encryption with HMAC-SHA256 integrity/authentication, client-side; the key object is the sole trust boundary, so backups can target untrusted storage `[official]` (https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/).
- Performance: chunking/compression/encryption in C/Cython; local file/chunks index cache; fast detection of unmodified files `[official]` (https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/).
- Architecture: repository (key-value store) holding multiple archives; internals layering — commands → Archive/Item data-access layer → compression (RepoObj) → authenticated encryption (key) → repository (local or remote via SSH agent protocol) `[official]` (https://github.com/borgbackup/borg/blob/HEAD/docs/internals.rst).
- Operations: `borg init --encryption=repokey`, `borg create`, `borg prune` (retention policies), `borg check` (integrity), `borg mount` (FUSE browse), `borg extract`; SSH transport with a borg agent on the server side reduces data over the wire vs plain SSHFS `[secondary]` (https://github.com/fatgrizzly/community-content/blob/HEAD/tutorials/install-and-configure-borgbackup/01.en.md).
- Caveats: one repository should be modified from one place (exclusive lock during create/delete; simultaneous backups fail); multi-client single-repo use is technically possible but not recommended under Borg's security model (shared-repo attack-model limits); cache rebuilds after external repo modification can be slow `[official]` (https://lira.no-ip.org:8443/doc/borgbackup-doc/html/faq.html — mirror of official FAQ `[secondary]`).
- Typical targets: Hetzner Storage Box and similar SSH targets are popular in the homelab/SMB segment `[secondary]` (https://github.com/fatgrizzly/community-content/blob/HEAD/tutorials/install-and-configure-borgbackup/01.en.md).
- Decision guidance (community 2026): Borg for Linux servers wanting dedup+encryption+speed; Restic/Duplicacy for cross-platform or S3-native; tar/rsync for trivial cases `[secondary]` (https://mangohost.net/blog/borgbackup-in-2025-efficient-compression-and-deduplication-for-linux-servers/).

### 11.2 Restic

- Restic is a fast, cross-platform (Go single binary), secure backup program: content-defined chunking deduplication, snapshots, authenticated encryption, with backends including local disk, SFTP, S3-compatible object storage, rest-server, rclone and more `[secondary]`.
- Version observed in the wild 2026: **0.18.x** (0.18.0/0.18.1 referenced in downstream changelogs; 0.18.1 fixes: `--stdin-filename` with directory paths, `RESTIC_HOST` honored by `forget`, chmod-on-CIFS/WebDAV-FUSE errors ignored, xattr EOPNOTSUPP ignored, rare crash on directory removal fixed, download-retry handling in `check`) `[secondary]` (https://github.com/backube/volsync/blob/HEAD/mover-restic/restic/CHANGELOG.md).
- Feature set: `backup`, `snapshots`, `forget` (retention policies), `check` (repository integrity), `copy` (snapshot copy between repos; `--copy-chunker-params` preserves cross-repo dedup), `mount` (FUSE), `restore`, `stats` with tag/path filtering; snapshots are full point-in-time views with dedup underneath `[secondary]` (https://go.libhunt.com/restic-changelog).
- Performance work: optimized Minio SHA-256 for chunk hashing (10–30% faster; ~2–3x on CPUs with SHA extensions); improved large-file restore path `[secondary]` (https://go.libhunt.com/restic-changelog).
- Ecosystem: resticprofile (declarative profiles), backrest (web UI), autorestic; Kubernetes movers (VolSync) use restic for PVC replication `[secondary]` (https://github.com/backube/volsync/blob/HEAD/CHANGELOG.md).
- Exact latest upstream release number not verified in this wave — downstream references confirm 0.18.x is current; upstream release page should be checked for 0.19+ `[unverified-current]`.

### 11.3 Kopia, Duplicacy

- Kopia: open-source backup with content-addressable deduplication, compression, client-side encryption, policy-driven snapshots, S3-compatible and cloud backends, CLI + KopiaUI graphical client; popular for workstation/server backup with repository-server mode for multi-client setups `[secondary]`. Latest 2026 version not verified in this wave `[unverified-current]`.
- Duplicacy: lock-free deduplication backup; CLI (free for personal use) and web GUI (licensed); supports major cloud backends; erasure-coding option for added durability `[secondary]`. Pricing/version details not re-verified 2026 `[unverified-current]`.

### 11.4 Bacula / Bareos

- Bacula: enterprise network backup suite (Director, Storage daemon, File daemon, catalog in SQL database); mature, complex, suited to tape libraries and large fleets `[secondary]`.
- Bareos: the open-source Bacula fork, actively maintained with its own release line; common in European public-sector/enterprise Linux estates `[secondary]`. Current 2026 version numbers not verified in this wave `[unverified-current]`.
- Positioning: heavier operational lift than Borg/Restic/PBS but strong for heterogeneous OS fleets, tape robotics and compliance cataloging; evaluate against Veeam for Windows-heavy estates.

## 12. Database backup: PostgreSQL and MySQL ecosystems

### 12.1 PostgreSQL — WAL archiving fundamentals

- PostgreSQL point-in-time recovery (PITR) rests on two pillars: a base backup plus continuous WAL (Write-Ahead Log) archiving via `archive_command`; recovery replays WAL segments to any timestamp `[official-concept]`.
- `pg_basebackup` is the built-in physical base-backup tool; production systems layer pgBackRest/Barman/WAL-G on top for compression, encryption, retention, parallelism and object storage `[secondary]`.
- WAL volume planning: busy OLTP systems can generate tens of GB of WAL per hour; `archive_timeout`, WAL segment size (16 MB default) and replication slots interact with backup tooling; unarchived WAL accumulation is a classic disk-full outage vector `[secondary]`.

### 12.2 pgBackRest

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

