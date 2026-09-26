---
id: etape7-phased-proxmox-backup/00-proxmox-backup/part-16
title: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 16)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-08"]
keywords: ["datacenter", "gpus"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [571, 591]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 779381f487b1eb957ed112465b274c171f9b73e10283a719bc4528cdaee33d5e
---

# Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 16)

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

