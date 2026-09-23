---
id: etape7-phased-proxmox-backup/00-proxmox-backup/4-proxmox-ve-storage-ceph-zfs-lvm-pbs-integration
title: "4. Proxmox VE storage: Ceph, ZFS, LVM, PBS integration"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["Broadcom"]
dates: ["2020-07", "2025-08", "2026-01-12", "2026-04", "2026-10-19", "2027-01-10"]
keywords: ["agent", "cost", "datacenter", "distribution", "latency", "license", "pricing", "throughput"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [52, 111]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 667d4dadd38d2eae51081f7ec8f24fc09b2bf2c1989e2b02b3c89f3baf9b1936
---

# 4. Proxmox VE storage: Ceph, ZFS, LVM, PBS integration

## 4. Proxmox VE storage: Ceph, ZFS, LVM, PBS integration

- Integrated Ceph SDS: hyper-converged Ceph deployable from the Proxmox UI/API (MON/MGR/OSD roles); Ceph Squid 19.2.3 in VE 9.0, Ceph Tentacle 20.2(.1) stable option in VE 9.2 `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf).
- Ceph provides unified block (RBD for VM disks), object (RGW) and file (CephFS) services on the same cluster `[secondary]`.
- Ceph upgrade path: Reef→Squid documented for VE 8.4→9.0 transitions; two-step process (Ceph first, then VE) `[secondary]` (https://github.com/community-scripts/ProxmoxVE/discussions/6600).
- OpenZFS: local ZFS pools with snapshots, replication (ZFS replication built in for DR), compression; ZFS 2.3 adds devices to existing RAIDZ pools with minimal downtime; ZFS 2.4 in VE 9.2 `[official/secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- LVM thick-provisioned shared storage: VE 9.0 added VM snapshot support via volume chains, covering iSCSI/FC SAN environments that previously lacked snapshots `[secondary]` (https://www.storagenewsletter.com/2025/08/21/availability-of-proxmox-virtual-environment-v-9-0-with-debian-13/).
- Backup fleecing: backup jobs can offload read I/O to a snapshot replica so production storage sees less backup load `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- PBS integration: native backup target for VE guests with incremental, deduplicated, encrypted backups; dirty-bitmap incremental backups transfer only changed blocks `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- PBS datastore placement guidance: local SSD (RAID-10 with ZFS) or all-NVMe backing; avoid NFS/SMB for datastores because dedup creates many small-file operations that are brutal on network-share latency `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- PBS is CPU-bound (SHA-256, zstd, verification), not network-bound; single-thread performance matters more than core count; community reports of 10 GbE links idling while throughput sits at 80–90 MB/s on weak CPUs `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).

## 5. Proxmox VE networking & security: SDN, firewall, identity

- SDN stack: zones (VLAN, QinQ, VXLAN, EVPN), VNets and subnets managed centrally; VE 9.0 introduced "SDN Fabrics" for routed resilient topologies (spine-leaf with multipath, automatic NIC failover) `[secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- VE 9.2 SDN protocols: native WireGuard fabric (encrypted inter-node traffic across untrusted networks, automatic key distribution) and BGP fabric; BGP/EVPN route filtering via route maps and prefix lists; OSPF route redistribution; IPv6 EVPN underlay `[secondary]` (https://linuxiac.com/proxmox-virtual-environment-9-2-released-with-dynamic-load-balancer/).
- Distributed stateful firewall: cluster → node → VM/CT level rules from the web UI/CLI/API; typical deployments allow VM self-service firewall management delegated via permissions `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- IPAM integration: coordinated IP address management for SDN-deployed networks `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Authentication: Proxmox realm plus AD, LDAP and OpenID Connect realms; two-factor authentication (TOTP); granular RBAC with roles/privileges, ACL propagation; API tokens and service accounts for automation without user credentials `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Guest compliance features: UEFI Secure Boot and vTPM 2.0 for Windows 11 and compliance-sensitive Linux guests; measured-boot capable via host TPM `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Automation: full REST API; unattended bare-metal installation; official Ansible playbooks; Terraform and OpenTofu providers for declarative VM/CT provisioning `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Note on overlap: Ansible/Terraform networking automation was covered in Step 6 Phase E2; this file covers the platform/virtualization angle only.

## 6. Proxmox VE licensing & subscription pricing (2026)

- Licensing model: per occupied physical CPU socket, per year; core count does not affect price; every server in a cluster needs coverage and all nodes in a cluster must use the same subscription level `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Official paid tiers (prices net, VAT extra; one-year term; page live 2026-09) `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing):
  - Premium: €1,100/socket/year — unlimited tickets, 2-hour first response (Severity 1), 24/7 support from 2026-10-19, remote SSH support, offline updates/key activation, Enterprise Repository, Datacenter Manager included.
  - Standard: €550/socket/year — 10 tickets/year, 4-hour response, 24/7 support "by the end of 2026", remote SSH support, offline updates, Enterprise Repository, Datacenter Manager included.
  - Basic: €370/socket/year — 3 tickets/year, 1 business day response, stable/secure updates, all features, Enterprise Repository, Datacenter Manager included.
- Community tier: access to the Enterprise Repository but no ticket support from the enterprise team; assistance via public forum `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing). Community pricing reported at €120/socket/year by a third-party 2026 overview — not shown on the official pricing page snapshot; treat as `[secondary]`, possibly outdated or region-specific (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- Free tier: fully functional for production; subscription adds the tested/stable Enterprise Repository stream and vendor support `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- Support coverage notes (2026): Premium 24/7 available from October 19, 2026 (including existing Premium customers); Standard 24/7 "by the end of 2026"; business-hours definitions vary by region (North America → NA hours, APAC → Eastern Australian hours, others CET/CEST) `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- TCO angle repeatedly cited in 2026 coverage: Broadcom/VMware licensing changes pushed cost-conscious shops toward Proxmox; per-socket model with no per-core or per-VM metering is a differentiator vs. per-core subscription models `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-datacenter-manager-1-1-vcenter-alternative/).

## 7. Proxmox Backup Server (PBS)

- Proxmox Backup Server is the companion backup product: open-source (AGPLv3), free to download/install/run with no license key; first public beta July 2020 `[secondary]` (https://edywerder.ch/proxmox-backup-server/).
- PBS 4.x released alongside VE 9.0 (August 2025) on Debian 13; PBS 4.2 released April 2026 on Debian 13 "Trixie" with ZFS 2.4 `[secondary]` (https://edywerder.ch/proxmox-backup-server/).
- Core engine features `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing):
  - Incremental backups — transfer only changed data (QEMU dirty-bitmap support for VM backups).
  - Global deduplication — duplicate chunks eliminated across backups, clients and VMs.
  - Zstandard compression; fixed and variable chunking to tune dedup for different data types.
  - Client-side AES-256-GCM encryption — data encrypted before transfer, suitable for untrusted targets.
  - SHA-256 checksums and scheduled verification — detects bit rot and damaged backups; verification workers/readers tunable (community reports ~800 MB/s default, 5+ GB/s tuned) `[official/secondary]`.
  - Roles and permissions — granular access to backup resources.
  - Live restore for VMs — start workloads while restore completes; granular file/directory recovery; snapshot catalogs; interactive recovery shell.
  - Remote synchronization — incremental sync jobs between datastores for off-site copies.
  - Native S3 datastore backends — public and private object storage as backup targets, with S3 usage monitoring (requests/traffic thresholds and notifications).
  - Tape: LTO-5 and later, tape autoloaders for long-term offline media.
- PBS subscription pricing (official; price list valid 2026-01-12 to 2027-01-10; per server/year; unlimited storage and unlimited backup clients included) `[official]` (https://proxmox.com/images/download/pbs/agreements/Proxmox-Backup-Server-Subscription-Agreement-V2.5.pdf):
  - Premium: €4,480/year — unlimited tickets, 2-hour response, 24/7, SSH remote support, offline activation, Enterprise Repository.
  - Standard: €2,240/year — 15 tickets/year, 4-hour response (24/7 "by end of 2026").
  - Basic: €1,120/year — 5 tickets/year, 1 business day response.
  - Community: €560/year — Enterprise Repository access, forum support only.
- Third-party pricing writeups conflict with official numbers (e.g., a Zmanda blog lists Basic/Standard/Premium at ~€119/€299/€599 per host with "unlimited email tickets" — that page appears to conflate VE and PBS tiers and is outdated); use the official agreement figures above `[official]`; conflicting third-party table flagged `[secondary/conflict]` (https://www.zmanda.com/blog/proxmox-backup-server-subscription-cost/).
- Sizing guidance (community 2026): datastore on local SSD/NVMe; verification throughput scales with `default-verification-readers/workers`; enable QEMU guest agent + dirty bitmaps to minimize backup windows `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- Homelab economics: PBS routinely deployed on sub-$200 mini-PCs for 20+ VMs, with dedup making small nodes viable; enterprise repo subscription optional `[secondary]` (https://edywerder.ch/proxmox-backup-server/).
- PBS backup client (`proxmox-backup-client`) also runs on any Debian-based Linux host, so physical servers can be protected without VE `[secondary]` (https://edywerder.ch/proxmox-backup-server/).

