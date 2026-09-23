---
id: etape7-phased-proxmox-backup/00-proxmox-backup/21-hypervisor-alternatives-deep-dive-supplement
title: "21. Hypervisor alternatives deep-dive (supplement)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "aws", "cost", "license", "packaging"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [423, 469]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: c177c96103f7a4e592cdc6af092d0ca5609050ea8f79448e889db70a408821e6
---

# 21. Hypervisor alternatives deep-dive (supplement)

## 21. Hypervisor alternatives deep-dive (supplement)

- XCP-ng storage: Storage Repositories (SR) — local EXT/LVM, NFS, iSCSI, FC, Ceph RBD via drivers; thin provisioning on file-based SRs; CBT (changed block tracking) for fast incremental backups `[secondary]`.
- XCP-ng networking: vSwitch (Open vSwitch based) with VLANs, bonds (LACP/active-backup), jumbo frames; SDN controller concepts for private networks across hosts `[secondary]`.
- Xen Orchestra backup: full/delta backups, continuous replication, disaster-recovery replication to remote XO; backup proxy (XO Proxy) for offloading; S3-compatible remote targets `[secondary]`.
- XO access model: per-XOA subscription covers unlimited pools/hosts; self-compiled XO from sources is free but lacks some premium features and the appliance conveniences `[secondary]` (https://forums.lawrencesystems.com/t/xcp-ng-orchestra/15640).
- XCP-ng on older hardware: community reports running on 4-year-old non-HCL hardware (e.g., Lenovo x240 M5 class) with best-effort support from Vates `[secondary]` (https://xcp-ng.org/forum/topic/3170/xen-orchestra-editions/4).
- Harvester networking: cluster network + VM networks via Kube-OVN; VLAN support; SR-IOV for direct VF assignment; multus for multi-interface pods/VMs `[secondary]`.
- Harvester storage: Longhorn replicated volumes (default 3 replicas); RWX/RWO access modes; snapshots, backups to S3/NFS; storage network separation recommended `[secondary]`.
- Harvester upgrades: single-click-ish upgrades via the UI with preflight checks; version skew policies between Rancher/KubeVirt components managed by the release train `[official]` (https://github.com/harvester/release-notes/blob/HEAD/v1.7.2.md).
- Harvester backup: VM backup/restore to S3-compatible targets; Longhorn volume backups; etcd/Rancher state protection needed for full cluster recovery `[secondary]`.
- Migration Proxmox→Harvester: no first-party importer; typical path is virt-v2v or manual disk import into Harvester images; plan network/storage mapping carefully `[secondary]`.
- VMware-exit decision factors 2026: license cost delta, staff skills (KVM vs Xen vs K8s), storage reuse (Ceph on Proxmox, XOSTOR on XCP-ng, Longhorn on Harvester), Windows guest support, backup tooling compatibility `[secondary]`.
- oVirt deprecation timeline: RHV 4.4 was the last major Red Hat line; migration targets commonly Proxmox or OpenShift Virtualization; treat oVirt as legacy-only in 2026 plans `[unverified]`.

## 22. Veeam deep-dive (supplement)

- Editions history: v12 line (v12.0 2023, v12.1, v12.2, v12.3) introduced immutable-by-design posture, IoC scanner, four-eyes authorization, YARA rule scanning; v13 (Nov 2025) is the current major `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- Licensing: Veeam Universal License (VUL) per-workload portable across hypervisor/cloud/physical; legacy per-socket (retired for new sales); subscription vs perpetual+S&S options historically — verify current packaging with Veeam/partner `[secondary]`.
- Veeam ONE: monitoring, reporting and analytics; Advanced license tier required for data-observability features in v13 `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- SureBackup/SureReplica: automated recovery verification — boots VMs from backup in an isolated virtual lab and runs test scripts; the commercial answer to "untested backups" `[secondary]`.
- Veeam Agent: Windows/Linux/macOS workload protection (physical servers, cloud VMs) managed from the same console; used in the v13 Community Edition lab pattern `[secondary]` (https://github.com/davperez-tech/backup-disaster-recovery).
- NAS/file backup: unstructured data backup with changed-file tracking; ransomware detection via the same inline-entropy engine `[secondary]`.
- Scale-out Backup Repository (SOBR): performance tier + capacity tier (object storage) with move/copy policies; immutability supported per extent `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- Hardened Repository: Linux XFS-based with immutability flag; single-use credentials; recommended as the primary on-prem immutable tier before cloud copy `[secondary]`.
- VeeamZip and ad-hoc backups: one-off full backups for migrations and pre-change checkpoints `[secondary]`.
- Kasten K10: Veeam's Kubernetes-native backup (acquired); relevant for Harvester/K8s estates alongside Velero `[secondary]`.
- Recovery orchestration: Veeam Recovery Orchestrator for documented, tested DR plans with runbooks and compliance reporting `[secondary]`.
- Threat-hunting integration: SIEM/SOAR webhooks (Splunk, ServiceNow) and EDR/XDR integrations (CrowdStrike, Palo Alto) for coordinated response `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- v13 Linux backup server: reduces Windows licensing for the backup infrastructure itself; PostgreSQL-backed configuration database option replaces MSSQL dependency in newer architectures `[secondary]` (https://vmorecloud.com/download-veeam-backup-replication-v13/).
- Upgrade note: v13 forum activity (Jan 2026) shows "Immutable Until" visibility quirks on hardened repositories after upgrade — verify display on non-object extents post-upgrade `[secondary]` (https://forums.veeam.com/object-storage-as-backup-target-f52/immutable-flag-duration-somewhere-visible-in-veeam-t69753.html).

## 23. Open-source backup tools deep-dive (supplement)

- Restic backends: local, SFTP, HTTP rest-server, S3 (AWS/MinIO/Ceph RGW/Garage), Azure Blob, GCS, B2, rclone (any rclone remote); `RESTIC_*` env vars for scripting `[secondary]`.
- Restic retention: `forget --keep-daily/weekly/monthly/yearly --prune`; policies per tag/path; `check --read-data` for full content verification (slow, schedule sparingly) `[secondary]`.
- Restic encryption: AES-256-CTR with Poly1305-AES message authentication; keys derived from the repository password; no key recovery — password loss = data loss `[secondary/unverified-detail]`.
- rest-server: lightweight Go HTTP server for restic with `--append-only` mode — clients can write but not delete, a cheap immutability primitive against compromised clients `[secondary]`.
- Automation: systemd timers or cron; lock handling for overlapping runs; `unlock` for stale locks; metrics via exit codes + `restic snapshots --json` for monitoring `[secondary]`.
- Borg patterns: `borg create --exclude 'sh:/home/*/.cache'` etc.; `--compression zstd,8`; `--stats`; prune policies `--keep-daily=7 --keep-weekly=4 --keep-monthly=6` `[secondary]`.
- Borg append-only: `borg serve --append-only` restricts clients from deleting archives — ransomware-resistant repos without full immutability `[secondary]`.
- Borg key management: `repokey` (stored in repo, needs passphrase), `keyfile` (local), `repokey-blake2`; export with `borg key export`; paper backup recommended `[official]` (https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/).
- Kopia policies: per-path compression (zstd/s2/pgzip), scheduling, retention, ignore rules; repository server mode with user accounts and ACLs; KopiaUI for desktop users `[secondary]`.
- Duplicacy details: chunk-level dedup with lock-free design (multiple clients, no coordinator); `-stats`, check/ prune commands; web GUI license per computer for commercial use `[secondary]`.
- Bareos specifics: Director/Storage/File daemons, PostgreSQL/MySQL/SQLite catalog; NDMP support; bareos-webui; packages for major distros; commercial subscriptions from Bareos GmbH `[secondary]`.
- Tool selection heuristics: single Linux server → Borg; mixed OS + S3 → Restic/Kopia; VM fleet on Proxmox → PBS; Postgres → pgBackRest; MySQL → XtraBackup; heterogeneous enterprise with tape → Bareos/Veeam `[secondary]`.

