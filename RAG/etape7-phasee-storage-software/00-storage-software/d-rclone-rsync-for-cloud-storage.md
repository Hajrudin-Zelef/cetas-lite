---
id: etape7-phasee-storage-software/00-storage-software/d-rclone-rsync-for-cloud-storage
title: "D. rclone — \"rsync for cloud storage\""
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: ["Google", "Intel", "United States"]
dates: ["2025-03-06", "2026-04", "2026-06"]
keywords: ["agent", "compute", "intel", "memory", "pricing", "research"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [111, 188]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 06f6b5753ed3c454b53dc6c6ab99cc22f68a64d62ce367f299beae6b53ad30ba
---

# D. rclone — "rsync for cloud storage"

## D. rclone — "rsync for cloud storage"

### D1. Project basics

- rclone is a Go-based CLI (single static binary, `CGO_ENABLED=0` in release builds) for syncing/moving data between local storage and cloud/remote backends [secondary].
- The project documents **70+ storage providers** on its overview page: Google Drive, Dropbox, OneDrive, S3, GCS, Azure Blob/Files, Backblaze B2, SFTP, WebDAV, SMB, Swift, iCloud Drive and others [secondary].
- Core commands: `rclone copy` (new/changed files), `rclone sync` (one-way, make destination identical), `rclone check` (hash-equality verification), `rclone move`, `rclone mount` (FUSE), `rclone serve` (HTTP/WebDAV/FTP/SFTP/DLNA), `rclone bisync`, `rclone dedupe`, `rclone lsjson` (machine-readable), `rclone rc` (remote-control HTTP API) [secondary].
- Integrity: MD5/SHA-1 hashes checked at all times; timestamps preserved; multi-threaded downloads to local disk; partial syncs on whole-file basis [secondary].
- The `rclone.conf` holds remotes and tokens; passwords are obscured (`rclone obscure`), and a config password encrypts tokens at rest [secondary].

### D2. Backends and virtual (wrapping) backends

- Storage backends include: S3, GCS, Dropbox, Azure Files/Blob, iCloud Drive, OneDrive, Google Drive, WebDAV, SMB, SFTP, plus `local` and `memory` [secondary].
- **Virtual backends** adapt or modify other remotes: `alias` (rename), `cache` (deprecated), `chunker` (split large files), `combine` (merge remotes into one tree), `compress` (transparent compression), `crypt` (client-side encryption), `hasher` (extra hash tracking), `union` (join remotes, with policies) [secondary].
- **`crypt` remote**: wraps any backend, encrypting filenames and contents before upload; AES-based; filename encodings (standard/base64/obfuscate); **losing the crypt password means unrecoverable data** — the single most repeated operational warning [secondary].

### D3. bisync, mount and serve

- **`rclone bisync`**: bidirectional sync keeping two locations in sync both ways, with `--resync` baseline establishment and conflict handling (strategies: newer/older/larger/smaller/local/remote/both) [secondary].
- **`rclone mount`**: FUSE-mount a remote as a local filesystem; `--vfs-cache-mode full` is the common production flag for write support and performance; `--allow-other` for multi-user access [secondary].
- **`rclone serve`**: expose local or remote files over HTTP, WebDAV, FTP, SFTP or DLNA — used for ad-hoc sharing and media serving [secondary].
- Automation fit: JSON output + `rclone rc` HTTP API make rclone scriptable for backup jobs, cross-cloud migrations and scheduled syncs; frequently used for agent-style automation of storage management [secondary].

### D4. Performance tuning

- Commonly tuned flags in production/field guides: `--fast-list` (fewer API listing calls), `--transfers N` and `--checkers N` (parallelism), `--check-first`, `--tpslimit` (transactions per second, to respect provider rate limits), `--timeout`/`--contimeout`, `--retries`/`--retries-sleep`/`--low-level-retries`, `--exclude-from`, `--delete-excluded`, `--ignore-existing`, `--log-file` [secondary].
- A representative 2026 backup recipe: `rclone sync /Photos/ remote:Photos --fast-list --transfers 4 --checkers 4 --check-first --tpslimit 10 --retries 5 --log-file=...` [secondary].

### D5. Alternatives and overlap notes

- Overlap with backup tooling (restic, Borg, Kopia) is acknowledged: rclone is a sync/transfer engine, not a deduplicating snapshot backup tool; detail on those lives in Phase 7D [secondary].
- Third-party GUIs/wrappers exist (e.g. Obsidian sync apps built on rclone with in-app `crypt` setup and `bisync`), showing rclone's role as an embeddable sync engine [secondary].

---

## E. TrueNAS — SCALE vs CORE in 2026

### E1. Product lineup and naming

- TrueNAS is developed by **iXsystems**; the free community product was **SCALE** (Debian Linux based) and **CORE** (FreeBSD based) [secondary].
- In **2025 iXsystems rebranded SCALE → "TrueNAS Community Edition (CE)"**; version 25.04 "Fangtooth" carried the CE branding [secondary]. Research in Sept 2026 references TrueNAS CE 25.10 as the current stable line and **TrueNAS 26 (beta April 2026)** as the next annual major [secondary].
- **CORE (FreeBSD) is "no longer under active development"** — no container story (no Docker/catalog/Compose), jails only; community guidance is unambiguous: do not build new infrastructure on CORE [secondary].

### E2. Release timeline 24.10 → 26

- **24.10 "Electric Eel"**: introduced Docker Compose–based Apps (replacing Kubernetes) and was the first TrueNAS to integrate **OpenZFS 2.3** features (Fast Dedup, RAIDZ expansion) [official][secondary].
- **25.04 "Fangtooth"**: Docker-based Apps fully in place; the old built-in k3s catalog (`truenas/charts`) was **removed** [secondary].
- **25.10 "Goldeye"**: 25.10.4 observed as stable tracker state in June 2026 [secondary].
- **TrueNAS 26 (beta April 2026)**: Linux 6.18 LTS base, OpenZFS 2.4, Docker Engine 29.0.4; **breaking change — the TrueNAS REST API was deprecated in 25.04 and is removed in 26** (automation moves to WebSocket API / `midclt`) [secondary].
- Annual major cadence announced from TrueNAS 26 onward [secondary].

### E3. Apps system (post-2025)

- Since 24.10/25.04, Apps are **Docker Compose** stacks managed through the TrueNAS UI, with a **Community Apps catalog** (single-node, orchestration-lite — explicitly not Kubernetes) [secondary].
- Pre-24.10 k3s-era catalog entries (e.g. `truenas/charts` registry chart 2.8.3, last touched 2025-03-06) are stale; upgrade is required before use [secondary].
- `midclt` (local middleware CLI) remains the stable automation path across versions, unaffected by the REST API removal: `midclt call app.query`, `pool.query`, `system.version`, etc. [secondary].

### E4. Storage and services surface

- TrueNAS exposes SMB, NFS, iSCSI (targets/extents/portals/initiators/CHAP), S3-compatible object sharing, NVMe-oF (subsystems/namespaces/ports/hosts), rsync tasks, cloud-sync tasks (S3/Drive/B2), ZFS replication tasks, periodic snapshots, scrubs, and VMs (KVM) — all manageable via WebSocket API/`midclt` [secondary].
- ZFS pool/dataset operations: `zpool list`, `zpool status -x`, `zfs get all`, ARC stats via `arcstat`/`/proc/spl/kstat/zfs/arcstats` [secondary].
- TrueCommand provides single-pane-of-glass multi-system management (monitoring, reports, predictive capacity/health analytics, audits) [vendor-reported].

### E5. Hardware: Mini line and Enterprise

- The **TrueNAS Mini** family is iXsystems' SOHO line; launch pricing **started at US$699** (Mini family), with a 70 TB Mini X+ configuration retailing **under $3,500** and all-flash configs up to 50 TB [vendor-reported].
- Mini X+ specs (community listings): 8-core 2.2 GHz Intel Atom CPU, 32 GB ECC RAM, 5× 3.5" + 2× 2.5" hot-swap bays, dual 1/10 GbE, IPMI, diskless options; ships with WD Red Plus drives 1–14 TB [secondary].
- Third-party 2026 listings show Mini X+ around **R 61,578** (South Africa, duties/VAT incl.) and Mini XL+ around **AED 9,842** (UAE) — import/reseller pricing, not iXsystems MSRP [secondary].
- **TrueNAS Enterprise** appliances are iXsystems' supported line (HA options, support contracts); Enterprise pricing is quote-based and not publicly listed [unverified].
- **TrueNAS Connect** SSO/monitoring add-on: Foundation tier free; Plus at **$60/year for 3 systems** [secondary].

### E6. Competitive positioning (2026 independent takes)

- TrueNAS CE: free, ZFS-native (data integrity/silent-corruption protection), Docker apps; weaknesses vs Unraid = no mixed-drive single pool (RAIDZ expands one drive at a time), higher RAM floor (~16 GB practical for ZFS), compute is secondary to storage [secondary].
- Community rule of thumb: **compute-first → Proxmox VE; storage-first → TrueNAS**; a mature pattern is running TrueNAS as a VM inside Proxmox to add ZFS later without a second box [secondary].

---

