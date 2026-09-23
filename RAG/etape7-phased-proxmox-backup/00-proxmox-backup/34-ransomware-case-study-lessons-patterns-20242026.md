---
id: etape7-phased-proxmox-backup/00-proxmox-backup/34-ransomware-case-study-lessons-patterns-20242026
title: "34. Ransomware case-study lessons (patterns, 2024–2026)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["cyber", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [642, 697]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 01e4e2c8f2bfb5f6f317bf693ce6e7f465c2f76c7af237b3462c74e491d8cbef
---

# 34. Ransomware case-study lessons (patterns, 2024–2026)

## 34. Ransomware case-study lessons (patterns, 2024–2026)

- Attackers target backups first: backup servers joined to the same AD forest are compromised with production — separation of identity is lesson #1 `[secondary]`.
- Dwell time matters: median dwell before ransomware detonation is measured in days to weeks — immutability windows shorter than dwell time fail `[secondary]`.
- Exfiltrated backups: unencrypted off-site copies become double-extortion leverage — client-side encryption neutralizes this `[secondary]`.
- Deleted cloud backups: attackers with cloud console access delete unprotected snapshots — Object Lock (compliance mode) or separate-account backup vaults prevent it `[secondary]`.
- Untested restores: organizations discovering corrupt/unrestorable backups during incidents is a recurring post-mortem finding — SureBackup-style automated verification exists for this reason `[secondary]`.
- Configuration backup gaps: hypervisor and backup-server configs unprotected — rebuilds take days without them `[secondary]`.
- Tape's comeback: air-gapped LTO copies remain the cheapest last-resort tier; several 2025–2026 incidents were recovered from tape after online backups were destroyed `[secondary]`.
- Tabletop exercises: teams that rehearsed ransomware playbooks restored 2–5× faster than those improvising `[secondary]`.
- Communication plans: pre-written stakeholder/regulatory notification templates save critical hours `[secondary]`.
- Insurance: cyber-insurance questionnaires increasingly ask for immutable backups and tested restores — document both `[secondary]`.

## 35. Backup encryption key management

- PBS: AES-256-GCM with passphrase-derived keys; export keys via the UI/CLI; store offline in two locations (sealed envelope + password manager) `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Borg: `borg key export` / `borg key import`; paperkey printouts for `repokey` modes; test import on a scratch host annually `[official]` (https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/).
- Restic: password in file or command; no recovery path — treat the password as the key; consider Shamir-split storage for teams `[secondary]`.
- pgBackRest: cipher passphrase in a restricted file (`0600`, owned by postgres/barman user); rotate per policy; old backups need old keys — archive retired keys `[official]` (https://github.com/pgstef/pgbackrest).
- Veeam: password loss on encrypted backups is unrecoverable without the enterprise key-management integration; document escrow before enabling encryption `[secondary]`.
- Rotation: rotating keys does not re-encrypt existing backups; plan retention overlap so old-key backups age out `[secondary]`.
- Separation: backup encryption keys must not live only on the backup server — ransomware encrypts that too `[secondary]`.

## 36. Command cheat-sheet (quick reference)

- `pvecm status` — cluster membership and quorum state `[secondary]`.
- `pvecm nodes` — list cluster nodes `[secondary]`.
- `qm list` / `pct list` — VMs and containers inventory `[secondary]`.
- `qm migrate <vmid> <target> --online` — live migration `[secondary]`.
- `qm snapshot <vmid> <name>` — VM snapshot `[secondary]`.
- `pvesh get /cluster/resources --type vm` — API query for all guests `[secondary]`.
- `pveperf` — quick host performance sanity check `[secondary]`.
- `proxmox-backup-client backup root.pxar:/ --repository <repo>` — file backup to PBS `[secondary]`.
- `proxmox-backup-client list --repository <repo>` — list snapshots `[secondary]`.
- `borg create --stats --compression zstd,8 <repo>::{hostname}-{now} /etc /home` — borg backup `[secondary]`.
- `borg prune --keep-daily=7 --keep-weekly=4 --keep-monthly=6 <repo>` — borg retention `[secondary]`.
- `borg check <repo>` — repository integrity `[secondary]`.
- `restic -r <repo> backup /data --tag daily` — restic backup `[secondary]`.
- `restic -r <repo> forget --keep-daily 7 --keep-weekly 5 --prune` — restic retention `[secondary]`.
- `restic -r <repo> check` — repository check `[secondary]`.
- `restic -r <repo> snapshots` — list snapshots `[secondary]`.
- `pgbackrest --stanza=main --type=full backup` — full Postgres backup `[secondary]`.
- `pgbackrest --stanza=main check` — stanza health `[secondary]`.
- `pgbackrest --stanza=main info` — backup inventory `[secondary]`.
- `pgbackrest --stanza=main --delta restore` — delta restore `[secondary]`.
- `xtrabackup --backup --target-dir=/backups/base` — MySQL hot backup `[secondary]`.
- `xtrabackup --prepare --target-dir=/backups/base` — crash-recovery prepare `[secondary]`.
- `xtrabackup --copy-back --target-dir=/backups/base` — restore data files `[secondary]`.
- `mysqldump --single-transaction --routines --triggers --all-databases > full.sql` — logical backup `[secondary]`.
- `zpool status -v` — ZFS pool health `[secondary]`.
- `zfs list -t snapshot` — ZFS snapshots `[secondary]`.
- `ceph -s` / `ceph health detail` — Ceph cluster state `[secondary]`.
- `ceph osd df tree` — OSD utilization `[secondary]`.
- `ha-manager status` — Proxmox HA manager state `[secondary]`.
- `pvesr status` — storage replication job status `[secondary]`.

