---
id: etape7-phased-proxmox-backup/00-proxmox-backup/overview
title: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2025-08", "2025-08-09", "2026-05-21", "2026-08", "2026-08-13", "2026-09-22"]
keywords: ["research", "amd", "benchmarks", "datacenter", "distribution", "intel", "memory", "nvidia", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [1, 51]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 94f4134b8e65442d66e3288641ffe188d45b144191dcfa371c89265a805e3f8e
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

- **Proxmox VE 9.0** released August 2025 on Debian 13 "Trixie" (Debian 13 itself released 2025-08-09) `[official/secondary]` (https://www.storagenewsletter.com/2025/08/21/availability-of-proxmox-virtual-environment-v-9-0-with-debian-13/).
- VE 9.0 core stack: Linux kernel 6.14.8-2 (stable default; newer than Debian 13's 6.12), QEMU 10.0.2, LXC 6.0.4, OpenZFS 2.3.3, Ceph Squid 19.2.3 `[official/secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- VE 9.0 headline storage feature: snapshot support for VMs on thick-provisioned LVM shared storage (including iSCSI and Fibre Channel SANs), implemented via volume chains where a child volume tracks differences from its parent snapshot `[secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- The same snapshot-as-volume-chain model extends to Directory, NFS and CIFS backends, giving storage-vendor-independent snapshots without clustered filesystems or proprietary hooks `[secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- ZFS 2.3 allows adding new devices to existing RAIDZ pools with minimal downtime (RAID-Z expansion, long awaited) `[secondary]` (https://www.theregister.com/software/2025/08/11/virtualization-vet-pushes-out-proxmox-ve-9-backup-server-4/1221487).
- VE 9.0 networking: "SDN Fabrics" for building routed, resilient topologies (two-layer spine-leaf with multipath between nodes, automatic NIC failover); dynamic routing via OpenFabric and OSPF; EVPN underlays or Ceph full-mesh networks easier to stand up `[secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- VE 9.0 HA: HA Resource Affinity Rules for fine-grained workload placement (e.g., pin a database and its app server to the same node; spread redundant services across nodes) `[secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- VE 9.0 mobile interface completely reworked on the Rust-based Yew framework with a new widget toolkit: quick service overviews, guest start/stop, basic configuration from a mobile browser `[secondary]` (https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/).
- Upgrade guidance from community docs: 8.4 clusters with Ceph Reef must first upgrade Ceph Reef→Squid, then Proxmox VE 8.4→9.0; VE 8.4 receives bug/security fixes until August 2026 (about one year of overlap) `[secondary]` (https://github.com/community-scripts/ProxmoxVE/discussions/6600).
- **Proxmox VE 9.1** shipped as a point release in the 9.x line; community field notes document pve-manager 9.1.4 running kernel 6.17.x (6.17.2-1-pve through 6.17.13-13-pve), qemu-kvm 11.0, ZFS 2.4.2, systemd 257.13 during a 9.1.4→9.2.3 upgrade path `[secondary]` (https://github.com/fiberoptix/home-lab-setup/blob/HEAD/phases/phase1b_proxmox_kernel_upgrade_safe_try.md).
- **Proxmox VE 9.2** released May 21, 2026 on Debian 13.5 "Trixie" with Linux kernel 7.0 as the new stable default `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf).
- VE 9.2 stack: QEMU 11.0, LXC 7.0, OpenZFS 2.4, Ceph Tentacle 20.2.1 as a stable option alongside Ceph Squid 19.2 `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf).
- VE 9.2 headline feature: Dynamic Load Balancer — a dynamic mode for the cluster resource scheduler using real-time CPU/memory utilization data from nodes and guests; it can automatically live-migrate HA-managed guests to reduce node imbalance while adhering to user-defined HA rules, with configurable sensitivity `[official/secondary]` (https://linuxiac.com/proxmox-virtual-environment-9-2-released-with-dynamic-load-balancer/).
- VE 9.2 HA maintenance: cluster-wide "disarm" and "arm" controls to temporarily suspend the HA Manager during planned maintenance and prevent false-positive fencing; two modes — Freeze (halts changes to current guest operational state) and Ignore (HA stops tracking guests so they appear HA-unmanaged); HA resource states preserved across cycles `[secondary]` (https://www.notebookcheck.net/Proxmox-9-2-Virtual-Environment-launches-with-the-7-0-Linux-kernel-as-default.1303073.0.html).
- VE 9.2 SDN: native WireGuard and BGP fabric support with automatic crypto key distribution across joined cluster nodes; BGP/EVPN filtering via route maps and prefix lists; OSPF route redistribution for fabrics; IPv6 underlay support for EVPN; additional EVPN controller options `[secondary]` (https://linuxiac.com/proxmox-virtual-environment-9-2-released-with-dynamic-load-balancer/).
- VE 9.2 virtualization: dedicated web-UI interface (Datacenter section) to create/edit/remove custom CPU models, plus an integrated CPU flags selector showing supported flags across cluster nodes to avoid live-migration incompatibilities `[secondary]` (https://www.notebookcheck.net/Proxmox-9-2-Virtual-Environment-launches-with-the-7-0-Linux-kernel-as-default.1303073.0.html).
- Confirmed current line as of 2026-09-22: Proxmox VE 9.2.x is the latest stable point release documented; VE 8.4 support window ends August 2026 `[secondary]`; no public confirmation of a VE 9.3 release found — flagged as a watch item `[unverified]`.

## 3. Proxmox VE architecture: KVM, LXC, clustering, HA

- QEMU/KVM full virtualization: hardware-assisted virtualization (Intel VT-x / AMD-V), OVMF/UEFI with Secure Boot, vTPM 2.0, VirtIO paravirtual drivers (net, block, balloon, rng), PCIe passthrough and SR-IOV, NVIDIA vGPU profiles, live migration, snapshots `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- LXC system containers: OS-level virtualization sharing the host kernel, near-native performance, suitable for dense microservice/DNS/DHCP/reverse-proxy workloads; LXC 7.0 in VE 9.2 `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Unprivileged containers are the recommended default for hardening LXC workloads; privileged containers retained for nested-docker-style use cases `[secondary]`.
- Clustering: up to 32 nodes per cluster historically (pmxcfs cluster filesystem over Coroutines/Corosync); multi-master design with no single management point of failure `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- HA stack: watchdog-based fencing; quorum required for reliable fencing; HA groups with configurable fencing delays in VE 9.0+ reduce split-brain risk in distributed setups `[secondary]` (https://www.webpronews.com/proxmox-ve-9-0-released-debian-13-base-sdn-enhanced-ha-features/).
- Dynamic HA balancing (VE 9.2): real-time CRS scheduling migrates HA-managed guests automatically on imbalance `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf).
- Migration tooling: ESXi import wizard eases VMware→Proxmox migration of VMs, listed as a first-class product feature `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Community/lab note: ClearPass Policy Manager 6.14 validated by third parties on Proxmox VE 9.2 (pve-manager, QEMU 11.0.3) as a KVM guest, 2026-08-13 — evidence of broad third-party appliance compatibility `[secondary]` (https://github.com/jnprautomate/fw-skills-share/blob/HEAD/skills/clearpass-proxmox-deploy/SKILL.md).

