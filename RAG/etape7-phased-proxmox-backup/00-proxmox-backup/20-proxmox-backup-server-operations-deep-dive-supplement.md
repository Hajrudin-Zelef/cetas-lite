---
id: etape7-phased-proxmox-backup/00-proxmox-backup/20-proxmox-backup-server-operations-deep-dive-supplement
title: "20. Proxmox Backup Server operations deep-dive (supplement)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "latency", "memory", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [372, 403]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: f4657cda01445acce8e45664f028585b5bf41d250f6c6a3f70b4bc6d7b9d3cc0
---

# 20. Proxmox Backup Server operations deep-dive (supplement)

- Repositories: `pve-enterprise` (stable, subscription), `pve-no-subscription` (testing-ish, free), `pvetest`; production guidance is enterprise repo + subscription `[secondary]`.
- Update model: Debian APT-based; `apt update && apt full-upgrade` on a schedule; major upgrades (8→9) via official upgrade guide with pre-upgrade checker `pve8to9` `[secondary]`.
- Kernel policy: Proxmox ships its own newer kernel (6.14 in VE 9.0, 7.0 in VE 9.2) for hardware enablement and ZFS; pinning kernels via APT holds is a documented practice during soak periods `[secondary]` (https://github.com/fiberoptix/home-lab-setup/blob/HEAD/phases/phase1b_proxmox_kernel_upgrade_safe_try.md).
- QEMU machine versions: pin VM machine type (e.g., `q35`) per VM for stable live-migration across QEMU upgrades; changing defaults on upgrade can break migration compatibility `[secondary]`.
- CPU models: `host` for max performance on homogeneous clusters; `kvm64`/`x86-64-v2-AES` baselines for heterogeneous clusters; VE 9.2 custom CPU profiles managed in the web UI `[secondary]`.
- VirtIO drivers: Windows guests need the VirtIO ISO (viostor/vionet/balloon/QEMU guest agent); guest agent enables clean shutdown, freeze/thaw for backups, IP reporting `[secondary]`.
- Ballooning: dynamic memory management for overcommit; KSM (kernel samepage merging) dedups identical host memory pages across VMs — useful for VDI and homogeneous fleets `[secondary]`.
- HugePages: 2 MB/1 GB pages reduce TLB pressure for databases and DPDK-style workloads; configure via kernel cmdline or sysctl `[secondary]`.
- NUMA: pin vCPUs and memory to NUMA nodes for latency-sensitive workloads; Proxmox exposes NUMA topology per VM `[secondary]`.
- Storage plugins: Directory, LVM, LVM-thin, NFS, CIFS/SMB, iSCSI, Ceph RBD, CephFS, ZFS (local and over iSCSI), Btrfs (experimental-ish), ESXi import source `[secondary]`.
- ZFS specifics: ARC cache sizing (`zfs_arc_max`), `ashift=12` for 4K drives, compression=lz4 default recommended, snapshots + `pvesr` replication jobs for async DR between nodes `[secondary]`.
- Ceph deployment from UI: install Ceph packages per node, create MON/MGR, create OSDs on dedicated disks (BlueStore), define pools (replicated size=3/min_size=2 typical); separate cluster/public networks recommended (ideally 25 GbE+) `[secondary]`.
- Ceph PG sizing: use the PG calculator; too few PGs hurts parallelism, too many hurts MON/OSD memory; autoscaler helps but manual review is advised `[secondary]`.
- Backup jobs: datacenter-level backup schedules (day-of-week, start time, selection pools); modes snapshot/suspend/stop; fleecing target for I/O offload; notification via email/Gotify/webhook targets `[secondary]`.
- Replication jobs: storage replication for ZFS/local volumes on schedules (e.g., every 15 min) for warm-standby DR `[secondary]`.
- Firewall: datacenter/node/VM levels; default policies (in/out accept/drop); security groups reusable across VMs; ipsets for IP lists; logging options per rule `[secondary]`.
- SDN IPAM: phpIPAM/NetBox-style integrations coordinate subnet allocation with SDN zones `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- HA details: requires shared storage for VM disks (or replicated ZFS); quorum via Corosync votequorum; fencing via IPMI/iLO/DRAC or watchdog; 2-node clusters need a qdevice (third vote) — commonly a Raspberry Pi or small VM `[secondary]`.
- Maintenance: VE 9.2 arm/disarm avoids fencing during rolling upgrades; prior practice was manual HA disable + node migrate `[secondary]`.
- Monitoring: built-in metric server targets (InfluxDB, Graphite); external Prometheus via `pve-exporter`; PDM 1.1 centralizes metrics and adds host utilization gauges `[official]` (https://pdm.proxmox.com/docs/roadmap.html).
- Notifications: notification targets (SMTP, sendmail, Gotify, webhook) with matchers routing by severity; backup/HA/fencing events covered `[secondary]`.
- User management: PAM, PVE, LDAP, AD, OpenID Connect realms; TOTP 2FA; pools for delegated administration; fine-grained privileges (e.g., VM.PowerMgmt without VM.Config) `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Templates: cloud-init enabled templates for fast VM provisioning; linked clones for instant copies sharing base image (storage must support it) `[secondary]`.
- Cloud-init: user-data/meta-data/network-config via snippets or the UI; standard path for immutable-ish VM provisioning alongside Terraform/OpenTofu `[secondary]`.
- ESXi import: wizard imports VMs/disks from ESXi/vCenter into Proxmox storage with driver adjustments; part of the 2026 VMware-exit toolkit alongside Vates' VDDK-based tool `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- ARM64: Proxmox advertises ARM64 support via sales for bleeding-edge server hardware (pricing page banner) — x86_64 remains the mainstream platform `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Common homelab topology: 3-node mini-PC cluster, Ceph on 2.5" SATA SSDs or NVMe, PBS on a 4th node or NAS-backed VM (local SSD preferred), 2.5 GbE management + 10 GbE storage network `[secondary]`.
- Split-brain avoidance: never run two clusters on the same Corosync network without isolation; use separate cluster networks or VLANs `[secondary]`.
- Backup encryption keys: PBS client-side encryption keys must be escrowed (paperkey/KDF); losing the key = losing all backups — a top operational footgun `[secondary]`.

## 20. Proxmox Backup Server operations deep-dive (supplement)

