---
id: etape7-phased-proxmox-backup/00-proxmox-backup/27-capacity-planning-formulas-and-rules-of-thumb
title: "27. Capacity planning formulas and rules of thumb"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-08"]
keywords: ["datacenter", "distribution", "gpus", "pricing", "throughput", "training"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [529, 591]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 7d2b887d0575f0d65e7f9a6a8dcc9ac626732c0c4b5e1f415b8d1c6be02cc3c4
---

# 27. Capacity planning formulas and rules of thumb

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

- VE 9.0 installer: based on Debian 13 netinst flow with ZFS (RAID0/1/10/RAIDZ) and Btrfs options at install time `[secondary]`.
- VE 9.0 LVM snapshot chains: child volumes track parent diffs; snapshot deletion merges chains; monitor chain depth on busy SAN LUNs `[secondary]`.
- VE 9.0 OpenFabric: link-state protocol for fabrics as an alternative to OSPF; useful for data-center underlays `[secondary]`.
- VE 9.0 affinity rules: `positive` (keep together) and `negative` (keep apart) rules evaluated by the HA manager `[secondary]`.
- VE 9.0 UI: dark mode in the reworked mobile interface; desktop UI refinements across guest, storage and SDN panels `[secondary]`.
- VE 9.1 (point release): kernel 6.17 series rollout; continued QEMU 11 and ZFS 2.4 stabilization before 9.2 `[secondary]` (https://github.com/fiberoptix/home-lab-setup/blob/HEAD/phases/phase1b_proxmox_kernel_upgrade_safe_try.md).
- VE 9.1→9.2 upgrade: community field notes show `apt full-upgrade` pulling kernel 7.0.6-2-pve, systemd 257.13 and new shim/systemd-boot; kernel metapackage handling required care (install `proxmox-kernel-6.17` metapackage, not just the image) `[secondary]` (https://github.com/fiberoptix/home-lab-setup/blob/HEAD/phases/phase1b_proxmox_kernel_upgrade_safe_try.md).
- VE 9.2 CRS: the classic static scheduler remains; dynamic mode is opt-in per cluster `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf).
- VE 9.2 WireGuard fabric: each node gets a keypair; public keys distributed via cluster config; encrypted VXLAN/overlay transport option for stretched clusters `[secondary]`.
- VE 9.2 BGP fabric: iBGP/eBGP sessions between nodes; route maps filter what is advertised; prefix lists bound what is accepted `[secondary]`.
- VE 9.2 EVPN: controller options expanded; IPv6 underlay support; SLAAC handling in EVPN contexts per video documentation `[secondary]`.
- VE 9.2 OSPF: route redistribution into fabrics for hybrid routing designs `[secondary]`.
- VE 9.2 CPU flags selector: compares host CPU flags cluster-wide in the UI; highlights migration blockers before they bite `[secondary]`.
- VE 9.2 custom CPU models: define once at datacenter level; assign per VM; version-controlled via cluster config `[secondary]`.
- VE 9.2 HA disarm: `Freeze` keeps guests running untouched; `Ignore` detaches HA tracking; both avoid fencing during maintenance `[secondary]`.
- VE 9.2 Ceph Tentacle: opt-in per cluster; Squid remains supported; upgrade path documented separately from VE upgrade `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf).
- VE 9.2 kernel 7.0: newer hardware enablement (NICs, GPUs); ZFS 2.4 feature compatibility; test with passthrough devices before fleet rollout `[secondary]`.
- VE 9.2 download: ISO includes the full feature set; alternate ISO download mirrors available `[secondary]`.
- Upgrade cadence: Proxmox ships point releases roughly every 4–6 months; plan maintenance windows accordingly `[secondary]`.
- EOL awareness: VE 8.4 security fixes end August 2026 — 8.x estates should be on 9.x by then `[secondary]` (https://github.com/community-scripts/ProxmoxVE/discussions/6600).

