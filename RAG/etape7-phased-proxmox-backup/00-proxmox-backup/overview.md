---
id: etape7-phased-proxmox-backup/00-proxmox-backup/overview
title: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["research", "benchmarks", "datacenter", "distribution", "nvidia", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [1, 22]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 500224bd77fc29b1847eaa827547498d0d28b9d72a74d62413fef7d9d969bc3b
---

# Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research

- **Scope:** Proxmox Virtual Environment (VE) 8.x/9.x, Proxmox Backup Server (PBS), Proxmox Datacenter Manager (PDM); hypervisor alternatives (XCP-ng/Xen Orchestra, Harvester HCI, oVirt); backup software (Veeam, BorgBackup, Restic, Kopia, Duplicacy, Bacula/Bareos); database backup (PostgreSQL WAL ecosystem: pgBackRest, Barman, WAL-G; MySQL: Percona XtraBackup, mysqldump); ransomware-proof backup design, 3-2-1 strategies, RTO/RPO, restore testing.
- **Observation date:** 2026-09-22. Research cutoff: 2026-09-22.
- **Method:** read-only web research via search and page fetch. No live login, no lab testing. Facts are tagged per provenance legend.
- **Provenance legend:** `[official]` = vendor/official docs or press releases; `[vendor-reported]` = vendor blogs/claims without independent confirmation; `[independent]` = independent press, labs, benchmarks; `[secondary]` = secondary coverage (community docs, forums, aggregators); `[unverified]` = single-source or unconfirmed claim. Gaps, conflicts and unverified claims are listed at the end, not silently dropped.
- **Language:** English per project convention. One writer on this file; no other workspace files modified.

## 1. Proxmox VE — product overview

- Proxmox Virtual Environment is a Debian-based open-source virtualization platform from Proxmox Server Solutions GmbH (Vienna, Austria); the company celebrated its 20th anniversary in 2025 `[official]` (https://www.theregister.com/software/2025/08/11/virtualization-vet-pushes-out-proxmox-ve-9-backup-server-4/1221487).
- The platform combines KVM/QEMU virtual machines and LXC system containers under one web UI (plus CLI and REST API), targeting hyper-converged infrastructure (HCI) deployments `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Proxmox modifies stock Debian to add enterprise virtualization features rather than shipping a vanilla distribution `[secondary]` (https://www.theregister.com/software/2025/08/11/virtualization-vet-pushes-out-proxmox-ve-9-backup-server-4/1221487).
- The source code is open (AGPLv3 for VE and PBS), and all features are available at every subscription level; plans differ only in support entitlement, ticket allowance and response time `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Feature surface per the official product page: Dynamic HA balancing with real-time CRS scheduling; multi-master clustering with no management single point of failure; zero-downtime live migration; HA maintenance workflows preserving resource state; unified QEMU/KVM and LXC; NVIDIA vGPU for AI/ML, rendering and VDI; PCIe passthrough; ESXi import and CPU controls; integrated Ceph SDS; Fibre Channel and iSCSI SAN support; multipath and shared LVM; templates and linked clones `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Backup/disaster-recovery surface: Proxmox Backup Server integration (incremental, encrypted, deduplicated); live restore and file recovery; backup fleecing (backup from snapshot replica to reduce production I/O); ZFS replication `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Networking surface: EVPN, VXLAN and BGP for multi-tenancy; SDN fabrics and routing policies; IPAM integration; distributed stateful firewall from cluster down to guest `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Identity/automation surface: AD, LDAP and OpenID Connect; granular RBAC and ACLs; API tokens and service accounts; UEFI Secure Boot and vTPM for compliance-ready guests; REST API; unattended installations; Ansible playbooks; Terraform and OpenTofu providers; solution-provider hardware ecosystem `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Community-scripts/ProxmoxVE is a large community project maintaining helper scripts for LXC and VM deployment on Proxmox VE `[secondary]` (https://github.com/community-scripts/ProxmoxVE/discussions/6600).

## 2. Proxmox VE 9.x release line (2025–2026)

