---
id: etape7-phased-proxmox-backup/00-proxmox-backup/11-open-source-backup-tools
title: "11. Open-source backup tools"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-09-09"]
keywords: ["agent", "license", "pricing", "robotics"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [173, 217]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: f0c0ff4a510cf720ba1807ec4e47e4854379d5d882fc07bce2d9392e54411f6c
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

