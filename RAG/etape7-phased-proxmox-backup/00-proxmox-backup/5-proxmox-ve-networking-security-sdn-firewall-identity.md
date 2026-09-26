---
id: etape7-phased-proxmox-backup/00-proxmox-backup/5-proxmox-ve-networking-security-sdn-firewall-identity
title: "5. Proxmox VE networking & security: SDN, firewall, identity"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["Broadcom"]
dates: ["2026-10-19"]
keywords: ["cost", "datacenter", "distribution", "latency", "pricing", "throughput"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [54, 88]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 0f2da763ae8dab00256e202362e92e42e822fefde259e0f7a28a0f653bb40f61
---

# 5. Proxmox VE networking & security: SDN, firewall, identity

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

