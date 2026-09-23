---
id: etape7-phased-proxmox-backup/00-proxmox-backup/16-gaps-conflicts-and-unverified-claims
title: "16. Gaps, conflicts and unverified claims"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["cost", "cyber", "datacenter", "distribution", "license", "pricing", "research"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [285, 369]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: d6f16531c6449b4be6994c1f8238beaa5feefc313d2f2f2a9b4c07587587c776
---

# 16. Gaps, conflicts and unverified claims

## 16. Gaps, conflicts and unverified claims

- **G1 — PVE Community €120/socket:** reported by one third-party 2026 overview; absent from the official pricing page. Possibly outdated or reseller-specific. `[conflict/secondary]`
- **G2 — Zmanda PBS pricing table:** lists Basic/Standard/Premium at ~€119/€299/€599 per host with "unlimited email tickets" — contradicts the official PBS agreement (€560/€1,120/€2,240/€4,480 per server). The Zmanda page appears to conflate VE and PBS tiers. Official figures used. `[conflict]`
- **G3 — Vates VMS 2026 pricing:** latest public figures are the 2024 bundle announcement (€1,000/€1,800 per host/yr, Essential €2,000/yr). Current 2026 pricing must be re-verified on vates.tech. `[unverified-current]`
- **G4 — Veeam 2026 list pricing:** no authoritative per-workload list price captured; Community Edition free tier (≤10 workloads) confirmed by community sources only. `[gap]`
- **G5 — Restic latest version:** downstream changelogs confirm 0.18.x current in 2026; upstream release page not checked for 0.19+. `[unverified-current]`
- **G6 — Kopia/Duplicacy/Barman/WAL-G/Bareos 2026 versions:** feature descriptions included; exact current versions not verified. `[unverified-current]`
- **G7 — oVirt 2026 status:** no verified release data; project in long-term decline post-RHV deprecation. `[gap]`
- **G8 — PVE 9.3:** no public confirmation found as of 2026-09-22; watch item. `[unverified]`
- **G9 — PBS S3 cost modeling:** S3-native datastores shift cost to request/API pricing; no 2026 TCO study captured comparing S3-backed PBS vs local NVMe for dedup-heavy workloads. `[gap]`
- **G10 — Borg 1.4.x default chunker:** official 1.4.5 docs describe Buzhash; upstream development trees reference FastCDC/Gear as default — version-dependent, verify per release. `[conflict-nuance]`
- **G11 — PDM HA:** active-standby for PDM itself is roadmap-evaluation only; two instances side-by-side is the current workaround. `[official-roadmap]`
- **G12 — Proxmox VE 9.2 adoption data:** no independent market-share or migration-volume figures captured (only vendor/community anecdote on VMware→Proxmox movement). `[gap]`

## 17. Glossary

- **AGPLv3:** copyleft license used by Proxmox VE/PBS; network use counts as distribution.
- **BlueStore:** Ceph's default OSD backend (see also Phase 7E).
- **CRS:** Cluster Resource Scheduler (Proxmox dynamic load balancing).
- **Dirty bitmap:** QEMU change-tracking bitmap enabling incremental VM backups.
- **Fencing:** forcibly isolating a misbehaving cluster node (power/reset) to protect shared resources.
- **Fleecing:** offloading backup read I/O to a snapshot replica.
- **PITR:** point-in-time recovery (database replay to an arbitrary timestamp).
- **pmxcfs:** Proxmox cluster filesystem (configuration database replicated over Corosync).
- **RPO/RTO:** recovery point objective / recovery time objective.
- **VDDK:** VMware Virtual Disk Development Kit (used by XO's VMware migration tool).
- **WAL:** Write-Ahead Log (PostgreSQL durability/recovery log).
- **XOSTOR:** Vates hyper-converged storage add-on for XCP-ng.

## 18. Source index (verbatim URLs)

- https://www.theregister.com/software/2025/08/11/virtualization-vet-pushes-out-proxmox-ve-9-backup-server-4/1221487
- https://linuxiac.com/proxmox-ve-9-0-launches-with-debian-13-under-the-hood/
- https://www.storagenewsletter.com/2025/08/21/availability-of-proxmox-virtual-environment-v-9-0-with-debian-13/
- https://www.webpronews.com/proxmox-ve-9-0-released-debian-13-base-sdn-enhanced-ha-features/
- https://github.com/community-scripts/ProxmoxVE/discussions/6600
- https://proxmox.com/en/products/proxmox-virtual-environment/pricing
- https://www.proxmox.com/en/products/proxmox-backup-server/pricing
- https://proxmox.com/images/download/pbs/agreements/Proxmox-Backup-Server-Subscription-Agreement-V2.5.pdf
- https://www.starwindsoftware.com/blog/proxmox-ve-and-proxmox-backup-server-an-in-depth-overview-of-open-source-virtualization-and-data-protection/
- https://edywerder.ch/proxmox-backup-server/
- https://www.zmanda.com/blog/proxmox-backup-server-subscription-cost/
- https://proxmox.com/en/about/company-details/press-releases/proxmox-datacenter-manager-1-1
- https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Datacenter-Manager-1-1-20260528-en.pdf
- https://pdm.proxmox.com/docs/roadmap.html
- https://www.starwindsoftware.com/blog/proxmox-datacenter-manager-1-1-vcenter-alternative/
- https://linuxiac.com/proxmox-virtual-environment-9-2-released-with-dynamic-load-balancer/
- https://www.notebookcheck.net/Proxmox-9-2-Virtual-Environment-launches-with-the-7-0-Linux-kernel-as-default.1303073.0.html
- https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Virtual-Environment-9-2-20260521-en.pdf
- https://github.com/fiberoptix/home-lab-setup/blob/HEAD/phases/phase1b_proxmox_kernel_upgrade_safe_try.md
- https://github.com/jnprautomate/fw-skills-share/blob/HEAD/skills/clearpass-proxmox-deploy/SKILL.md
- https://xcp-ng.org/blog/2024/10/07/xcp-ng-8-3/?ref=xen-orchestra.com
- https://www.itspyworld.com/2024/08/xen-orchestra-590-released-xcp-ng-830.html
- https://xcp-ng.org/forum/topic/3170/xen-orchestra-editions/4
- https://github.com/xcp-ng/xcp-ng-org/blob/HEAD/docs/intro.md
- https://www.vladan.fr/vmware-to-vates-migration-tool-for-xcp-ng-and-xen-orchestra/
- https://github.com/harvester/harvester/blob/HEAD/README.md
- https://github.com/harvester/release-notes/blob/HEAD/v1.7.2.md
- https://github.com/harvester/release-notes/blob/HEAD/v1.5.0.md
- https://github.com/harvester/docs/blob/HEAD/docs/index.md
- https://medium.com/@0.all_existence.0/kubernetes-meets-virtualization-installing-harvester-on-bare-metal-ed04c3ac10bc
- https://www.vladan.fr/veeam-vbr-13-and-scality-artesca-integration-a-game-changer-for-cyber-resilient-backups/
- https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/
- https://forums.veeam.com/object-storage-as-backup-target-f52/immutable-flag-duration-somewhere-visible-in-veeam-t69753.html
- https://github.com/davperez-tech/backup-disaster-recovery
- https://vmorecloud.com/download-veeam-backup-replication-v13/
- https://borgbackup.readthedocs.io/_/downloads/en/1.4.5/pdf/
- https://github.com/borgbackup/borg/blob/HEAD/docs/internals.rst
- https://www.freshports.org/archivers/py-borgbackup/
- https://github.com/fatgrizzly/community-content/blob/HEAD/tutorials/install-and-configure-borgbackup/01.en.md
- https://mangohost.net/blog/borgbackup-in-2025-efficient-compression-and-deduplication-for-linux-servers/
- https://github.com/backube/volsync/blob/HEAD/mover-restic/restic/CHANGELOG.md
- https://github.com/backube/volsync/blob/HEAD/CHANGELOG.md
- https://go.libhunt.com/restic-changelog
- https://github.com/pgstef/pgbackrest
- https://github.com/percona/postgresql-docs/blob/HEAD/docs/release-notes/release-notes-v18.6.1.md
- https://github.com/r314tive/pgdrill/blob/HEAD/compatibility/evidence/pgbackrest-v2.58.0-postgresql-18.3-linux-amd64-pgdrill-v0.3.0-alpha.10-emulated/README.md
- https://github.com/percona/docs-home/blob/HEAD/docs/new/posts/MySQL/percona-xtrabackup-8.4.0-6-release.md
- https://github.com/percona/docs-home/blob/HEAD/docs/new/posts/MySQL/percona-xtrabackup-9.7.1-rc1-release.md
- https://github.com/evgeniypatlan/percona-mysql-skills/blob/HEAD/percona-xtrabackup/SKILL.md
- https://documentation.commvault.com/2024e/expert/mysql_backup_using_percona_xtrabackup.html

*End of Step 7 Phase D. Single writer; no other workspace files modified. Research cutoff 2026-09-22. File written append-only in 5 waves.*

