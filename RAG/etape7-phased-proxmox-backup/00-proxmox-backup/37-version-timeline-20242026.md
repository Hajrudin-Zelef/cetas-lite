---
id: etape7-phased-proxmox-backup/00-proxmox-backup/37-version-timeline-20242026
title: "37. Version timeline (2024–2026)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2024-10-07", "2025-07-20", "2025-08-09", "2025-08-11", "2026-05-21", "2026-05-28", "2026-06-24", "2026-07-15", "2026-08", "2026-08-17", "2026-09-09", "2026-09-22", "2026-10-19"]
keywords: ["aws", "benchmark", "datacenter", "pricing", "research", "throughput"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [698, 755]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 8ff057fa09feebb001b2681d4addbf6a4276ad2d4561a8fd8dd42e4b647ba42a
---

# 37. Version timeline (2024–2026)

## 37. Version timeline (2024–2026)

- 2024-10-07: XCP-ng 8.3 released (first under consolidated Vates model) `[vendor-reported]`.
- 2025-07-20: pgBackRest 2.59.0 released `[official]`.
- 2025-08-09: Debian 13 "Trixie" released `[secondary]`.
- 2025-08-11: Proxmox VE 9.0 + PBS 4.0 announced in press `[secondary]`.
- 2025-11: Veeam Backup & Replication v13 released `[secondary]`.
- 2026-04: Proxmox Backup Server 4.2 released (Debian 13, ZFS 2.4) `[secondary]`.
- 2026-05-21: Proxmox VE 9.2 released (Debian 13.5, kernel 7.0, QEMU 11, Ceph Tentacle) `[official]`.
- 2026-05-28: Proxmox Datacenter Manager 1.1 released `[official]`.
- 2026-06-24: Percona XtraBackup 8.4.0-6 released `[secondary]`.
- 2026-07-15: Percona XtraBackup 9.7.1-rc1 released `[secondary]`.
- 2026-08-17: pgBackRest 2.59.1 released (current stable) `[official]`.
- 2026-08: VE 8.4 support window ends (security fixes until August 2026) `[secondary]`.
- 2026-09-09: FreeBSD borgbackup port at 1.4.4_2 (Borg 1.4.x line current) `[secondary]`.
- 2026-10-19: Proxmox Premium 24/7 support takes effect (including existing Premium customers) `[official]`.
- End of 2026: Proxmox Standard 24/7 support planned `[official]`.
- 2026-09-22: observation date for this file.

## 38. Decision trees (quick selectors)

- New virtualization cluster, Linux-skilled team, tight budget → Proxmox VE 9.2 + PBS `[secondary]`.
- VMware-like operations, Xen acceptable, per-host billing preferred → XCP-ng + Xen Orchestra `[secondary]`.
- Kubernetes-first org, VMs as KubeVirt workloads → Harvester HCI `[secondary]`.
- Windows-heavy enterprise needing app-aware backup + compliance reports → Veeam v13 `[secondary]`.
- Single Linux server, dedup + encryption, SSH target → BorgBackup `[secondary]`.
- Mixed OS fleet, S3-native, single binary → Restic `[secondary]`.
- Desktop/workstation with GUI needs → Kopia/KopiaUI `[secondary]`.
- Heterogeneous estate with tape library → Bareos or Veeam `[secondary]`.
- PostgreSQL production → pgBackRest (repos: local + S3) + streaming replica `[secondary]`.
- MySQL 8.4 production → Percona XtraBackup 8.4.x + binlog archiving `[secondary]`.
- Ransomware resilience on any stack → immutable copy (S3 Object Lock / hardened repo / tape) + tested restores `[secondary]`.
- RTO < 1 h for VMs → local fast copy + PBS live restore or Veeam Instant Recovery `[secondary]`.

## 39. Open questions for future waves

- Re-verify Vates VMS bundle pricing for 2026 on vates.tech before publishing any TCO comparison `[unverified]`.
- Capture authoritative Veeam VUL per-workload list pricing from a Veeam price list or partner quote `[unverified]`.
- Check restic upstream releases for 0.19+ and Kopia/Duplicacy/Barman/WAL-G/Bareos current versions `[unverified]`.
- Confirm whether Proxmox VE 9.3 exists or is planned; watch the roadmap wiki `[unverified]`.
- Collect an S3-costed PBS TCO study (API request pricing vs local NVMe) for dedup-heavy datastores `[unverified]`.
- Validate the €120/socket Proxmox Community figure against the shop or reseller price lists `[unverified]`.
- Track oVirt community activity to decide whether it stays in scope for future phases `[unverified]`.
- Benchmark PBS 4.2 S3-backend throughput vs local NVMe on identical hardware `[unverified]`.
- Document Harvester + Kasten K10 backup integration status for Kubernetes-native DR `[unverified]`.
- Survey Proxmox VE 9.2 WireGuard fabric performance overhead vs plain VXLAN `[unverified]`.
- Record Ceph Tentacle 20.2 vs Squid 19.2 real-world upgrade experiences in Proxmox clusters `[unverified]`.
- Map Proxmox Datacenter Manager 1.1 API coverage gaps vs VE/PBS native APIs `[unverified]`.
- Investigate Azure/AWS-native backup for Proxmox-hosted workloads (hybrid estates) `[unverified]`.
- Compare ZFS replication vs PBS sync jobs for cross-site DR RTO in the 15-minute tier `[unverified]`.
- Evaluate Proxmox ARM64 builds maturity for edge deployments `[unverified]`.
- Review LTO-10 drive/library pricing and PBS tape-pool ergonomics in 2026 `[unverified]`.
- Assess Valkey/Redis persistence backup interplay with VM-level backups (covered in Phase 7F) `[unverified]`.
- Track Veeam v13.x patch releases and any immutability-display fixes for hardened repos `[unverified]`.
- Collect field data on Proxmox→Proxmox cross-site stretched clusters with WireGuard fabrics `[unverified]`.
- Decide whether OpenNebula / plain libvirt deserve a dedicated alternatives section in a later phase `[unverified]`.

*End of Step 7 Phase D file. Single writer; no other workspace files modified. Research cutoff 2026-09-22. Written append-only in 8 waves; Section 39 added as final wave.*
