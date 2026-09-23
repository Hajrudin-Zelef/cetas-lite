---
id: etape7-phasee-storage-software/00-storage-software/o-further-operational-detail-final-expansion
title: "O. Further operational detail (final expansion)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: ["2025-04"]
keywords: ["gpu", "latency", "lean", "pricing", "throughput"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [581, 651]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 1d8fe753f402328c8926d5320b5c53189a0ebbed594b1dc9d6789d3b301f496c
---

# O. Further operational detail (final expansion)

## O. Further operational detail (final expansion)

### O1. ZFS resilver, scrub, spares and ZED

- **Resilver** (rebuilding redundancy after disk replacement) is ZFS's equivalent of RAID rebuild; sequential resilver (2.x) dramatically speeds it versus legacy healing traversal on mostly-empty pools [secondary].
- **Hot spares**: `zpool add pool spare <disk>`; spares activate automatically on drive failure; dRAID distributes spare capacity instead [secondary].
- **ZED (ZFS Event Daemon)**: monitors pool events and can trigger scripts (email alerts, spare activation, LED blinking); TrueNAS wires its alerting into ZED-equivalent middleware [secondary].
- **Scrub vs resilver**: scrub verifies all checksums (integrity), resilver restores redundancy (rebuild); schedule scrubs monthly off-peak, more often on large/old arrays [secondary].
- `zpool history` and `zpool events` give the audit trail for pool operations [secondary].

### O2. ZFS data-protection extras

- **`copies=2|3`**: ditto blocks — extra copies of data/metadata within the same pool, protecting against bitrot even on single-disk pools (not a substitute for redundancy) [secondary].
- **Self-healing**: on every read, checksums are verified; on redundant vdevs, a bad block is transparently repaired from a good copy [secondary].
- **Snapshots are read-only and instant**; **clones** are writable snapshots (used by TrueNAS for VM/test workflows) [secondary].
- **Bookmarks**: lightweight send/receive resume points without keeping the snapshot [secondary].
- **Native encryption gotchas**: encrypted datasets can't use dedup effectively; key management (passphrase vs keyfile vs TPM) is the operational burden; `zfs change-key` rotates keys [secondary].

### O3. ZFS recordsize and workload tuning (detail)

- Default `recordsize=128K`; guidance: 1M for media/sequential, 64K–128K general, 16K/32K for databases and VM images (match guest filesystem/DB page where practical) [secondary].
- `volblocksize` is the zvol equivalent (set at creation, immutable); 16K–64K typical for VM zvols [secondary].
- `primarycache`/`secondarycache` per-dataset ARC/L2ARC control (e.g. `primarycache=metadata` for backup targets) [secondary].
- `logbias=throughput` for workloads that don't need low-latency sync (avoids SLOG contention) [secondary].
- `atime=off` universally recommended; `xattr=sa` and `dnodesize=auto` for modern workloads [secondary].

### O4. rclone advanced usage

- **Filters**: `--include`, `--exclude`, `--exclude-from` (file), `--filter-from` with `+`/`-` rules, `--min-size`/`--max-size`, `--min-age`/`--max-age`, `--delete-excluded` [secondary].
- **Backup-dir pattern**: `--backup-dir remote:archive/$(date +%F)` keeps overwritten/deleted files versioned — the poor-man's versioning for `sync` [secondary].
- **Union policies**: `epall`/`epmfs` (existing path), `newest`, `lfs` (least free space) for write placement across member remotes [secondary].
- **VFS cache modes** for `rclone mount`: `off` (read-only streaming), `minimal`, `writes`, `full` (full read/write caching; required by most apps) [secondary].
- **rc API**: `rclone rc` exposes every command over HTTP+JSON (`rclone rcd` daemon); used with `--rc-user/--rc-pass` and `--rc-addr` for dashboards and automation [secondary].
- **Config encryption**: `rclone config` → set password (`s` option) encrypts tokens in `rclone.conf`; without it, OAuth tokens sit obfuscated-but-recoverable [secondary].

### O5. TrueNAS apps ecosystem notes

- The Community Apps catalog (Docker Compose era) is curated by iXsystems; **TrueCharts** — the dominant third-party k3s catalog — ended its TrueNAS SCALE support when k3s was removed (25.04), a major ecosystem event for upgraders [secondary].
- Custom apps: any Docker Compose YAML can be installed as a custom app, preserving flexibility beyond the catalog [secondary].
- GPU passthrough to apps: supported for VMs straightforwardly; containers need manual device configuration [secondary].

### O6. Synology SHR and filesystem notes

- **SHR (Synology Hybrid RAID)**: Synology's LVM/mdraid-based flexible RAID allowing mixed drive sizes with 1- or 2-disk redundancy; the key reason buyers accept the Synology premium [secondary].
- DSM uses **Btrfs** on Plus/XS models (snapshots, checksums, quota) and ext4 on value/J models; Btrfs is not ZFS — no native encryption-at-rest on older DSM, RAID scrub exists but no block-pointer rewrite [secondary].
- Drive compatibility enforcement (April 2025 policy) interacts with SHR expansion: replacing with non-certified drives on 25-series Plus models loses health/firmware features [secondary].

### O7. QNAP model and pricing signals (partial)

- Popular 2026 QNAP families: TS-x53D (2.5GbE, Celeron), TS-x64 (N5095/5105), TS-hx73A (QuTS hero, Ryzen), TVS-hx74 (QuTS hero, Core i5/i7) [secondary].
- Street pricing observed in 2026 passes: TS-464 (4-bay) commonly ~$600–700 diskless; TS-673A (6-bay) ~$900–1,000 diskless — reseller-dependent, treat as signals not MSRP [secondary].
- QuTS hero models carry a premium over QTS equivalents for the ZFS feature set [secondary].

### O8. Ceph miscellany worth keeping

- **mClock scheduler** (`osd_mclock_profile`) shapes QoS between client I/O, recovery, and scrub; profiles like `balanced` are default [secondary].
- **Balancer modes**: `crush-compat` (legacy), `upmap` (PG-level remapping without data movement) — Squid/Tentacle lean on automatic read balancing on top [secondary].
- **ceph-volume** supports `lvm` and `raw` modes; Tentacle reworked OSD mapper lifecycle for activate paths [official].
- **BlueStore compression**: per-pool `compression_mode` (none/passive/aggressive/force), `compression_algorithm` (lz4/zstd/snappy); BlueStore compresses only blocks that benefit, tracked per-extent [secondary].
- **RBD features**: layering (clones), exclusive-lock, object-map, fast-diff, deep-flatten, journaling (legacy), mirroring (`rbd-mirror` for DR, journal or snapshot-based) [secondary].
- **RBD encryption**: LUKS2-based per-image encryption (`rbd encryption format`), kernel 5.4+ for krbd [secondary].
- **Stretch clusters**: 2-site + tiebreaker MON for metro HA (niche, documented) [secondary].

### O9. Small-box Ceph and alternatives recap

- 3-node converged Ceph (e.g. 3× mini-PCs with 2–4 NVMe each) is a popular 2026 homelab pattern for learning; production guidance still says ≥5 OSD hosts for real durability [secondary].
- For 1–3 nodes, ZFS replication (TrueNAS/Proxmox) or Linstor/DRBD usually beats Ceph on simplicity; Ceph's value starts at scale and multi-tenancy [secondary].
- **MicroCeph** (Canonical): snap-based single-to-few-node Ceph for edge/IoT — the low-end on-ramp [secondary].

---

