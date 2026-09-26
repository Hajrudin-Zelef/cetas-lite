---
id: etape7-phasee-storage-software/00-storage-software/b-ceph-architecture-and-core-concepts
title: "B. Ceph — architecture and core concepts"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["backlog", "compute", "consumer", "cost", "datacenter", "dram", "latency", "memory", "nand"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [46, 102]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 15d5739d342607c7e395bd41669cc855dc0fdf844264bb1f9431ae3fffb98a71
---

# B. Ceph — architecture and core concepts

## B. Ceph — architecture and core concepts

### B1. Daemon topology

- **MON (monitors)**: maintain the cluster map via Paxos consensus, authenticate via CephX; odd counts (3 or 5) recommended [secondary].
- **MGR (managers)**: host Prometheus metrics, Dashboard, PG autoscaler, balancer; deployed as active/standby pair [secondary].
- **OSD (object storage daemons)**: one per device, store data on BlueStore, handle replication, recovery and scrub; each OSD has ~4–6 GB baseline memory footprint [secondary].
- **MDS (metadata servers)**: CephFS namespace (directory tree, inodes, client capabilities); can run multiple active ranks with subtree partitioning [secondary].
- **RGW (RADOS gateway)**: stateless S3/Swift HTTP gateway over RADOS; multi-site replication and bucket management [secondary].

### B2. CRUSH, pools and placement groups

- **CRUSH** is a deterministic pseudo-random placement algorithm: clients hash object IDs to placement groups (PGs), then map PGs to OSDs via the CRUSH hierarchy (root → datacenter → rack → host → OSD). No central lookup table; clients talk directly to OSDs [secondary].
- The **failure domain** in a CRUSH rule controls replica separation (typical: `host`; stricter: `rack`) [secondary].
- Pool data-protection types: **replicated** (size=3/min_size=2 typical; 33% usable capacity; robust, good read latency, easy scrubbing) vs **erasure-coded** (e.g. k=4,m=2 → 67% usable; k=8,m=3 ≈ same redundancy as 3x replication at ~1.5x raw overhead; compute-intensive on write/scrub; considered less battle-proven than 3x replication) [secondary].
- Red Hat guidance: with 1 PB usable, triple replication needs ≥3 PB raw; EC 8+3 needs ~1.5x raw. **Only the data pool can be erasure-coded**; service-data and bucket-index pools must be replicated [official].
- **Do not use hardware RAID under Ceph**: Ceph already provides durability; a degraded RAID hurts performance and rebuilds slower than Ceph's own deep-copy/EC-chunk recovery [official].
- **PG target**: 100–200 PGs per OSD across all pools; the pg-autoscaler (enabled by default) manages this automatically [secondary].
- Key PG states: `active+clean` (healthy), `active+degraded` (replicas missing, I/O continues), `peering` (I/O blocked), `inactive` (all OSDs down) [secondary].
- Reads go to the primary OSD (or any shard on EC pools); checksums are validated on every read [secondary].

### B3. The three interfaces: RBD, CephFS, RGW

- **RBD (RADOS Block Device)**: exposes Ceph as block devices ("images") via kernel `rbd`, `librbd`, or NBD; blocks are stored as ~4 MB RADOS objects; integrates with libvirt and OpenStack Cinder; Kubernetes PVs and VM disks are the canonical use cases [secondary].
- **CephFS**: POSIX file via kernel client or FUSE; metadata in MDS; small-file workloads are a known weakness; shared filesystems and NFS/SMB gateways are the use cases [secondary].
- **RGW**: object storage (S3 + Swift REST); full-object writes/appends fit EC pools natively; partial writes on EC pools require `allow_ec_overwrites` on BlueStore OSDs (FileStore unsafe/slow), and EC pools lack omap — so RBD/CephFS on EC need a replicated metadata pool plus `--data-pool` pointing at the EC pool [secondary].
- **Cache tiering** can front an EC pool with a fast writeback pool to recover omap/performance semantics: `ceph osd tier add ecpool hot-storage` + `cache-mode writeback` + `set-overlay` [secondary].
- A comparison table from a 2026 community source positions Ceph RGW (Tentacle 20.2.1) against MinIO/AIStor, SeaweedFS and Garage: Ceph is LGPL-2.1, multi-protocol (S3+Swift+STS), per-pool EC, bitrot detection via BlueStore per-block CRC32C, strong consistency within cluster, minimum production cluster ≈ 3 MON + 5+ OSD + 2+ RGW (~10 nodes) [secondary].
- Distributed-filesystem comparison (2026, community): Ceph's strengths = unified block/file/object, multi-petabyte scale, strong consistency, EC, mature snapshots/clones, OpenStack integration, excellent self-healing; weaknesses = famously high operational complexity (OSD/PG/BlueStore/scrub/recovery tuning is a full-time job), CephFS small-file penalty, ~4–6 GB RAM per OSD, ugly tail latency during recovery; TCO drivers = real expertise or paid support (Red Hat/IBM, Croit, 42on, SoftIron), EC cuts raw cost ~40% vs 3x replication but raises CPU/recovery I/O, 10–25 GbE minimum with separate cluster and public networks [secondary].

---

## C. Ceph — BlueStore internals and OSD planning

### C1. BlueStore vs FileStore vs SeaStore

- **BlueStore** is the production OSD backend: raw block device + embedded RocksDB for metadata + BlueFS (a tiny filesystem) for the RocksDB WAL/DB files. Per-block CRC32C checksums detect bitrot on every read [secondary].
- **FileStore** (XFS-backed legacy) is unsafe/slow for EC overwrites and is not a recommended path on modern releases [secondary].
- **SeaStore** (on Crimson-OSD) is the long-term rewrite, tech-preview only in Tentacle, still not the production default as of 2026-09 [official][secondary].

### C2. DB/WAL device sizing (separate vs colocated)

- Classic split layout for HDD OSDs: `--block.db` (RocksDB DB) and `--block.wal` (RocksDB WAL) on a fast device via `ceph-volume lvm create --data /dev/sda --block.db /dev/nvme0n1p1 --block.wal /dev/nvme0n1p2` [secondary].
- A community partitioning recipe for 4 HDD OSDs sharing one 400 GB NVMe: 2 GB WAL per OSD, 50 GB DB per OSD [secondary]. A homelab variant uses 5 GB DB and 1 GB WAL via `bluestore_block_db_size`/`bluestore_block_wal_size` [secondary].
- **Red Hat's sizing rule**: `block.db` at **4% of block.data** for object/file/mixed workloads, **1%** for pure block/Cinder workloads; `block.wal` small (~10 GB) and only on a device faster than the DB device [official].
- **For all-NVMe clusters, colocating DB/WAL on the same NVMe as data is the correct, recommended design** — a dedicated DB/WAL device only helps when the data device is much slower (HDD/SATA). One detailed 2026 home-ops analysis confirmed `bluefs_dedicated_db=0`/`bluefs_dedicated_wal=0` on colocated NVMe is not a bug; the real risk on consumer NVMe is the **lack of power-loss protection (PLP)**: BlueStore's RocksDB WAL commit issues synchronous flushed writes, and without PLP a sync write must reach NAND before ack, making RocksDB-commit latency the choke point (observed as a `_txc_apply_kv` OSD crash) [secondary].
- Practical NVMe cautions from that same 2026 field report: prefer enterprise/datacenter NVMe with PLP for Ceph; DRAM-less consumer drives (e.g. HMB-based) are a poor fit for OSD duty [secondary].

### C3. RocksDB/BlueStore tuning knobs

- Ceph's 2022 RocksDB deep-dive (still referenced in 2026 tuning guides) tested option sets on NVMe: highest-performing configurations combined `max_write_buffer_number=128`, `min_write_buffer_number_to_merge=16`, `write_buffer_size=8 MB`, level compaction, `max_background_jobs=4`, `max_total_wal_size=1 GB`; LZ4 compression lowered RGW write amplification with moderate performance impact (notably on bucket listing) [official].
- BlueStore cache hierarchy tuning includes `bluestore_cache_trim_interval`, cache size/meta ratio; a Rook homelab example sets `bluestoreCacheSize: 3G`, `bluestoreCacheMetaRatio: 0.8`, `bluestoreMinAllocSize: 4K`, and enables `bluestoreCacheAutotune` [secondary].
- Scrub/recovery operational knobs seen in production homelabs: restrict scrubs to off-peak (`osd_scrub_begin_hour`/`end_hour`), `osd_scrub_during_recovery=false`, deep-scrub interval 2 weeks, `osd_recovery_max_active=8`, `osd_max_backfills=4`, `osd_recovery_op_priority=1` (low priority to protect client I/O) [secondary].
- OS-level tuning bundled with Ceph NVMe guides: 256 MB socket buffers (`net.core.rmem_max/wmem_max`), `netdev_max_backlog=30000`, `vm.swappiness=1`, NVMe queue `nr_requests=1024`, I/O scheduler `none` for NVMe [secondary].

### C4. NVMe-oF gateway, dashboard, Rook, adoption

