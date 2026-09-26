---
id: etape7-phased-proxmox-backup/00-proxmox-backup/27-capacity-planning-formulas-and-rules-of-thumb
title: "27. Capacity planning formulas and rules of thumb"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["datacenter", "distribution", "governance", "pricing", "throughput", "training"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [516, 570]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: af4f9b0b1a3567a4f2bc0ee31ccac9afa8513d619429555fd0c46b1d1000fad2
---

# 27. Capacity planning formulas and rules of thumb

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

## 27. Capacity planning formulas and rules of thumb

- Proxmox host RAM: sum of VM allocations + 2–4 GB host overhead + ZFS ARC (cap at ~50% RAM if ZFS is primary storage) + Ceph OSD ~4 GB per OSD daemon `[secondary]`.
- Ceph raw-to-usable: usable ≈ raw / replication size (e.g., /3); minus ~10–15% overhead for PG distribution and balancer slack `[secondary]`.
- Ceph OSD count: minimum 3 nodes for replicated pools; production guidance often 4+ nodes to survive maintenance + failure simultaneously `[secondary]`.
- PBS datastore: first full ≈ unique data; daily delta ≈ change rate; with dedup across similar VMs, 20 VMs × 40 GB may store as ~200–400 GB after dedup `[secondary]`.
- PBS CPU: 1 modern core per ~100–200 MB/s of backup throughput is a rough planning figure; verify with your data mix `[secondary]`.
- Network: 10 GbE minimum for Ceph storage networks in production; 25 GbE+ for all-flash; management can stay 1 GbE `[secondary]`.
- Backup window math: window hours × throughput MB/s × 3.6 = GB per window; size jobs to fit with 30% headroom `[secondary]`.
- Tape: LTO-9 native 18 TB, LTO-10 in market (verify current generation pricing); 2.5:1 compression ratio is marketing — plan on native capacity `[secondary/unverified-current]`.
- Cloud egress: restores from S3 incur egress fees; keep a local fast copy for RTO-sensitive restores `[secondary]`.
- Database WAL disk: provision archive destination at ≥ 2× expected daily WAL volume; alert at 70% `[secondary]`.

## 28. Common errors and troubleshooting

- `TASK ERROR: cluster not quorate` — lost quorum; check Corosync network, votes, qdevice `[secondary]`.
- Fencing loops after network blip — increase fence delay, fix the underlying network flapping `[secondary]`.
- VM fails to migrate: CPU model mismatch — align CPU types or use a common baseline/custom model (VE 9.2 UI) `[secondary]`.
- `storage 'X' is not available on node 'Y'` — storage not marked shared or node missing the mount; check storage config scope `[secondary]`.
- PBS slow backups: check CPU saturation first (not network); enable dirty bitmaps; verify datastore is local SSD `[secondary]`.
- PBS verify failures: investigate disk health (SMART/ZFS scrub) — do not ignore checksum errors `[secondary]`.
- `borg check` errors after unclean shutdown — run `borg check --repair` only after reading docs; keep key backups `[secondary]`.
- Restic `lock` errors — `restic unlock` after confirming no active run; stale NFS locks are common `[secondary]`.
- pgBackRest `check` failing on archive-push — verify `archive_command` permissions and stanza config on the DB host `[secondary]`.
- XtraBackup version mismatch error — install the PXB major matching the server major; do not bypass with `--no-server-version-check` `[secondary]` (https://github.com/evgeniypatlan/percona-mysql-skills/blob/HEAD/percona-xtrabackup/SKILL.md).
- Ceph HEALTH_WARN: nearfull OSDs — add capacity or reduce pool size; never let OSDs hit full `[secondary]`.
- ZFS pool degraded — replace disk, `zpool replace`, resilver; keep hot spares for large pools `[secondary]`.
- Time drift breaking Corosync/TLS — run chrony/NTP on all nodes; >~1s skew causes cluster instability `[secondary]`.

## 29. Training, certification and community resources

- Proxmox official training courses and the Proxmox Academy video library (including PDM 1.1 "what's new" videos) `[official]` (https://proxmox.com/en/services/training-courses/videos/proxmox-datacenter-manager/whats-new-in-proxmox-datacenter-manager-1-1).
- Proxmox community forum (forum.proxmox.com) — primary community support channel, especially for Community-tier users `[official]`.
- Proxmox bug tracker and public git repositories for source-level debugging `[secondary]`.
- Community-scripts/ProxmoxVE helper scripts for fast LXC/VM deployment in labs `[secondary]` (https://github.com/community-scripts/ProxmoxVE/discussions/6600).
- Vates training and certification program for XCP-ng/Xen Orchestra `[vendor-reported]` (https://xcp-ng.org/blog/2024/10/07/xcp-ng-8-3/?ref=xen-orchestra.com).
- Veeam Vanguard community and R&D forums — deep technical threads (e.g., v13 immutability UI behavior) `[secondary]` (https://forums.veeam.com/object-storage-as-backup-target-f52/immutable-flag-duration-somewhere-visible-in-veeam-t69753.html).
- Vendor docs as ground truth: pgBackRest (pgbackrest.org), Borg (borgbackup.readthedocs.io), Proxmox wiki (pve.proxmox.com/wiki), Harvester docs (docs.harvesterhci.io) `[official]`.
- Homelab learning path: 1 node PVE + PBS VM → 3-node cluster + Ceph → PDM + off-site sync; each stage teaches clustering, storage and backup independently `[secondary]`.

## 30. Proxmox VE 9.x detailed changelog notes (supplement)

