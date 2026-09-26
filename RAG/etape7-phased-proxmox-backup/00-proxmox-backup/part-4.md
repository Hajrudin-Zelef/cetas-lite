---
id: etape7-phased-proxmox-backup/00-proxmox-backup/part-4
title: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 4)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2020-07", "2025-08", "2026-01-12", "2026-04", "2027-01-10"]
keywords: ["agent", "cost", "license", "pricing", "throughput"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [89, 111]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 7f84245b7de0fbca623c0b5cb8b63f2e4f9ffa28232114b54be76493298025ac
---

# Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 4)

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

