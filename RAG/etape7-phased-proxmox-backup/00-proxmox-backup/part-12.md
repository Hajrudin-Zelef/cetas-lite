---
id: etape7-phased-proxmox-backup/00-proxmox-backup/part-12
title: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 12)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["datacenter", "pricing", "throughput"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [404, 422]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: c2d4590fcdde6080998d92d01466565f331ee45f8c9c2a0c0b4d282f92f442c3
---

# Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 12)

- Datastore: chunk store on a filesystem (XFS/ext4/ZFS typical); namespaces organize clients; retention/prune policies per namespace `[secondary]`.
- Chunking: fixed chunker (e.g., 4 MB) vs variable (content-defined, better for shifted data); variable chunking improves dedup on VM images with insertions `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Deduplication scope: across all backups in a datastore — 20 similar VMs share the OS chunks, so incremental storage growth is small after the first full `[secondary]`.
- Verification: scheduled verify jobs re-read chunks and check SHA-256; catches bit rot before a restore does; tune readers/workers for throughput `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- Garbage collection: prune removes snapshots per retention; GC reclaims unreferenced chunks; schedule GC after prune windows `[secondary]`.
- Sync jobs: pull/push between PBS instances (datastore → remote namespace); incremental; used for off-site copies and 3-2-1 `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- S3 backend: datastore on S3-compatible storage (public cloud or MinIO/Garage/Ceph RGW on-prem); usage monitoring with request/traffic thresholds; mind API request costs on high-churn datastores `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Tape: media pools, retention, barcode label management; LTO-5 through current LTO generations; autoloader/changer support; offline tapes = air gap `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Encryption: AES-256-GCM client-side; key derivation from passphrase; key backup/escrow mandatory; encryption happens before network transfer `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Access control: API tokens with granular permissions; separate tokens per client host with least privilege; audit log of operations `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- `proxmox-backup-client`: backup arbitrary paths (`backup root.pxar:/ --repository ...`), dump databases to files then back up; cron/systemd timers for scheduling; works on any Debian-based host `[secondary]` (https://edywerder.ch/proxmox-backup-server/).
- pxar format: Proxmox's archive format with metadata, ACLs, xattrs; used for container and file backups `[secondary]`.
- Live restore: boot a VM directly from backup while blocks stream in the background — RTO measured in minutes even for large VMs `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- File-level restore: mount backup snapshots via the UI or recovery shell; download single files without full VM restore `[official]` (https://www.proxmox.com/en/products/proxmox-backup-server/pricing).
- Performance tuning: `default-verification-readers 16`, `default-verification-workers 32` community recipe; NVMe datastore; 10 GbE+ for multi-client; CPU with strong single-thread (SHA-256/zstd) `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/).
- Sizing rule of thumb: datastore usable ≈ (sum of unique data × retention depth × change rate) / dedup ratio; plan 20–30% headroom for GC and verification temp space `[secondary]`.
- Monitoring: built-in task log, datastore usage graphs, S3 request metrics; alerting via notification targets `[secondary]`.
- PDM integration: PBS remotes visible in Datacenter Manager with location tags and centralized subscription handling `[official]` (https://proxmox.com/en/about/company-details/press-releases/proxmox-datacenter-manager-1-1).

