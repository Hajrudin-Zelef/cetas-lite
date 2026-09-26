---
id: etape7-phased-proxmox-backup/00-proxmox-backup/3-proxmox-ve-architecture-kvm-lxc-clustering-ha
title: "3. Proxmox VE architecture: KVM, LXC, clustering, HA"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2025-08", "2025-08-09", "2026-05-21", "2026-08", "2026-08-13", "2026-09-22"]
keywords: ["amd", "datacenter", "distribution", "intel", "memory", "nvidia", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [23, 53]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 5fa8e6416db98166b868846222890497941b24dc88d4f80ee12ecacb941c5d87
---

# 3. Proxmox VE architecture: KVM, LXC, clustering, HA

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

## 4. Proxmox VE storage: Ceph, ZFS, LVM, PBS integration

