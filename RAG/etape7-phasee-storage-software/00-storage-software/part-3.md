---
id: etape7-phasee-storage-software/00-storage-software/part-3
title: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) (part 3)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [103, 110]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 52f01bd0c577834f43b2675758e675c03c6a65d0186f26ed2462ff46f0b3c69f
---

# Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) (part 3)

- The **NVMe-oF gateway** (introduced in Squid, hardened in Tentacle 20.2.1 with fast-failover rework and `nvme-gw listeners` command) exports RBD as NVMe/TCP targets so VMware/bare-metal clients consume Ceph block without a Ceph client; the Dashboard manages gateway groups, namespaces, and IO statistics [official][secondary].
- **Ceph Dashboard** (MGR module): multi-cluster management, OAuth 2.0, NVMe-oF group management, RGW/SMB lifecycle features; 20.2.x backports added hardware monitoring via node-proxy metrics and CSV host import [official].
- **Rook** (Kubernetes operator): v1.19.x current in mid-2026 homelab/enterprise guides; typical pattern = operator Helm chart + `CephCluster` CR + StorageClasses via CSI (`rook-ceph.rbd.csi.ceph.com`); Ceph Squid→Tentacle upgrades under Rook documented with pre-upgrade backup runbooks (ceph status, osd tree/df, pool detail, CRUSH dump, config dump, versions) [secondary]. Rook v1.19 supports Ceph Tentacle targets (v19.2.x → v20.2.x) [secondary].
- **Adoption signals**: CERN operates one of the world's largest Ceph deployments (multi-petabyte, referenced as a canonical large-scale adopter) [secondary]; Proxmox VE ships integrated Ceph management and documents Reef→Squid→Tentacle upgrade runbooks [secondary]; OpenStack Cinder/Glance/Manilla integrate RBD/CephFS/RGW [secondary]; commercial support from Red Hat/IBM (Red Hat Ceph Storage 8 hardware guide current), Canonical, SUSE, plus specialists Croit, 42on, SoftIron [secondary].
- **TCO notes**: 10–25 GbE minimum with separate public/cluster networks; EC pools cut raw cost ~40% vs 3x replication at the price of CPU and recovery I/O; operational expertise is the dominant cost driver [secondary].

---

