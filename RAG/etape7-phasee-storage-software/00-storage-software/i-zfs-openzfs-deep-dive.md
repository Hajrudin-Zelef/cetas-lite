---
id: etape7-phasee-storage-software/00-storage-software/i-zfs-openzfs-deep-dive
title: "I. ZFS / OpenZFS deep-dive"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: ["2024-10-04"]
keywords: ["latency", "memory", "pruning"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [269, 329]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 03e19849db6fc3df27c342c25846fcce7ce3e1bd8b4826b50b2e87dbfa7d7a65
---

# I. ZFS / OpenZFS deep-dive

## I. ZFS / OpenZFS deep-dive

### I1. OpenZFS release train

- **OpenZFS 2.1.x**: dRAID introduced; final maintenance releases reached 2.1.16+ by late 2024 [secondary].
- **OpenZFS 2.2** (released 2023): **block cloning** headline feature, **Blake3** checksums, Linux 6.5 kernel support, OverlayFS support, idmapped mounts, `zpool scrub -e` (targeted repair of known-damaged blocks), vdev properties for observability, and a new `zfs receive` mode that can heal corrupted data [secondary].
- **OpenZFS 2.3** (branched 2024-10-04, GA early 2025): **RAIDZ expansion** and **Fast Dedup** headline features; 2.2.x maintenance continued to 2.2.7+ [secondary][official].
- **TrueNAS integration**: Electric Eel (24.10) first integrated OpenZFS 2.3 features; Dragonfish (24.04) and CORE 13.3 carried 2.2 (dRAID, block cloning); TrueNAS 26 targets **OpenZFS 2.4** on Linux 6.18 LTS [official][secondary].
- Proxmox VE 9.x documents Fast Dedup operation for its ZFS users [secondary].

### I2. dRAID

- **dRAID (distributed RAID)** arrived with OpenZFS 2.1: distributes parity and spare capacity across all disks instead of dedicated RAIDZ + hot-spare disks, giving **much faster resilver** on wide arrays (typical for 10+ disk NAS/enterprise) [secondary].
- Notation: `draid1:2d:4c:0s` = 1 parity, 2 data disks per redundancy group, 4 children per group, 0 distributed spares [secondary].
- Trade-offs: less flexible expansion than classic RAIDZ, fixed stripe width, and `zpool remove`/device-removal limitations; best for greenfield large arrays [secondary].

### I3. Block cloning

- Block cloning (2.2) lets a file or part of a file be **cloned with shared data blocks** (shallow copy); later modifications copy-on-write only the changed blocks — the ZFS equivalent of reflinks/file-level CoW [secondary].
- Common Linux copy tools (newer `cp`) attempt to create clones automatically via `copy_file_range` [secondary].
- Caveat: not all OS/filesystem bridges wire it up (e.g. some niche ports lack the `vop_copy_file_range` path), so verify on non-mainstream platforms [secondary].

### I4. RAIDZ expansion

- RAIDZ expansion (2.3) **adds a drive to an existing RAIDZ vdev and rebalances data across all disks**; the array stays usable and fault-tolerant during the (long, possibly multi-day) rewrite; progress survives reboots and a drive failure during expansion is tolerated; repeatable [secondary].
- Limitation: **RAIDZ level cannot change** — a RAIDZ1 cannot become RAIDZ2; additional parity drives cannot be added by expansion [secondary].

### I5. Fast Dedup

- Legacy ZFS dedup (since ~2010) was memory-hungry and slow: the dedup table (DDT) lived in ARC with unbounded growth; guidance on 8 GB boxes was "do not enable" [secondary].
- **Fast Dedup (2.3)** reworks the backend with four mechanisms: **DDT log** (batch sorted DDT writes into ZAP leaves instead of random single-record writes, killing write inflation), **prefetch**, **pruning** (`zpool ddtprune`), and a **quota** (`dedup_table_quota`) bounding DDT size [official][secondary].
- Operational model (Proxmox VE 9.x guide): enable `feature@fast_dedup` on the pool, set `dedup_table_quota`, reserve ARC, optionally pin DDT on dedicated flash; monitor DDT size, ARC hit ratio, dedup efficiency [secondary].
- Workload fit: **good** for VM templates, VDI, backups; **unsuitable** for databases and encrypted data [secondary].
- Important: existing data is not retroactively deduplicated — only data written after enabling; to convert, copy/replicate data with the new setting active; per-dataset dedup is possible while the DDT spans the pool [official].

### I6. Caching architecture: ARC, L2ARC, ZIL, SLOG

- **ARC (Adaptive Replacement Cache)**: in-RAM read cache; primary ZFS performance lever — "add RAM first". Stats via `arcstat` and `/proc/spl/kstat/zfs/arcstats` [secondary].
- **L2ARC**: SSD-backed second-level read cache; useful only when ARC hit rate is low and workload is read-heavy and larger than RAM; consumes some ARC for its index (roughly 1 GB ARC per ~50–100 GB L2ARC, version-dependent) [secondary].
- **ZIL (ZFS Intent Log)** / **SLOG (separate log device)**: the ZIL handles synchronous writes; placing it on a dedicated fast, low-latency, **power-loss-protected** SSD (SLOG) dramatically improves sync-write workloads (NFS, iSCSI, databases, VMs). Without PLP the SLOG is pointless — sync writes must reach stable media [secondary].
- **Special vdevs (special allocation class)**: dedicated fast vdev for small blocks and metadata; accelerates metadata-heavy workloads and small-file performance; **a special vdev failure can take the pool down** — mirror it like data vdevs [secondary].
- Dedup vdevs (DDT on special class) are a related pattern for fast-dedup deployments [secondary].

### I7. Compression, checksums, encryption, snapshots

- Compression: lz4 default recommendation; **zstd** available (levels 1–19) with excellent ratio/speed trade-offs; gzip legacy [secondary].
- Checksums: fletcher4 default, sha256/sha512/skein/edonr/**blake3** options; checksums on every block enable self-healing against redundant copies [secondary].
- Native encryption: `aes-256-gcm` (also aes-128-ccm/gcm); per-dataset, keylocation prompt/file/URI; raw sends (`zfs send -w`) preserve encryption [secondary].
- Snapshots/clones: cheap, instant, CoW; replication via `zfs send | zfs receive` (TrueNAS replication tasks, Proxmox, Sanoid/Syncoid in homelabs) [secondary].
- Recordsize tuning: 128K default; 1M for large sequential/media; 16K/32K for databases/VMs; `recordsize` is settable per dataset [secondary].
- Scrub: periodic integrity verification; schedule off-peak; `zpool scrub -e` for targeted repair of known-bad blocks [secondary].

### I8. ZFS on the NAS OSes (cross-reference)

- **TrueNAS**: ZFS-only data path; full pool topology management in UI (data/cache/log/spare/special vdevs), scrubs, snapshots, replication, cloud sync [secondary].
- **Unraid**: ZFS pools as cache/array-adjacent devices; foreign pool import since 7.1; parity array remains the default [secondary].
- **OMV**: ZFS via plugin; mdraid+ext4/xfs/btrfs default [secondary].
- **Proxmox VE**: ZFS as first-class local storage (documented Fast Dedup ops in VE 9.x); Ceph as the clustered alternative [secondary].

---

