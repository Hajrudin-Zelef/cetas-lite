---
id: etape7-phased-proxmox-backup/00-proxmox-backup/part-6
title: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 6)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2025-11"]
keywords: ["agent", "agents", "aws", "cost", "cyber", "license", "open source", "pricing"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [157, 172]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 63508c6c22b57d5cf0168823f8c9ca7605231df3e6ad747a35c0f3bf5cb112c6
---

# Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research (part 6)

- Veeam Backup & Replication **v13 released November 2025** with major security, performance and object-storage advances `[secondary]` (https://www.vladan.fr/veeam-vbr-13-and-scality-artesca-integration-a-game-changer-for-cyber-resilient-backups/).
- v13 headline: **Linux-based backup server support** (no longer Windows-only for the backup server role) `[secondary]` (https://vmorecloud.com/download-veeam-backup-replication-v13/).
- v13 security: hardened Linux repositories with **immutability by default**; backups immutable by default to align with ransomware best practices; S3 Object Lock support; "Immutable Until" visibility added in the v13 UI for object storage (Sept 2025) and confirmed working in v13.0.1.1071 for hardened repositories (Jan 2026 forum thread) `[secondary]` (https://forums.veeam.com/object-storage-as-backup-target-f52/immutable-flag-duration-somewhere-visible-in-veeam-t69753.html).
- v13 malware/cyber features `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/):
  - AI-driven Malware Analysis agent: detects, classifies and reports malware/suspicious activity with guided remediation.
  - Proactive investigation (new in v13): automatically starts signature-based scans of suspicious objects and auto-marks backups clean when nothing is found — automates the malware-detection flow (some capabilities require Advanced licensing).
  - Retained v12 methods: inline entropy analysis of data streams, guest file-system activity analysis, IoC Tools Scanner (v12.3, MITRE ATT&CK-based, Advanced license).
  - Backup data encryption at rest; zero-trust architecture with granular access; MFA; event-driven alerting.
  - Integrations: CrowdStrike, Palo Alto Networks, Splunk, ServiceNow for unified detection/investigation/response.
- v13 platform: agentless backups for vSphere/Hyper-V; backup from snapshots and storage integrations; direct backup to object storage (S3, Azure Blob); Kubernetes/cloud-native backup; secure HTTPS web console; RBAC improvements; centralized job management `[secondary]` (https://vmorecloud.com/download-veeam-backup-replication-v13/).
- Ecosystem: **Scality ARTESCA + Veeam Unified Software Appliance** — bundles VBR with S3-compatible ARTESCA object storage on one hardened platform (CORE5: architectural immutability, zero-trust, no exposed endpoints), qualified as both repository and object storage with S3 Object Lock `[secondary]` (https://www.vladan.fr/veeam-vbr-13-and-scality-artesca-integration-a-game-changer-for-cyber-resilient-backups/).
- Community/lab pattern (2026): Veeam B&R v13 **Community Edition** (free) protecting lab estates with agent-based image-level backups, forever-forward-incremental chains, local NTFS repositories plus AWS S3 with Object Lock as immutable cloud tier (14-day retention example); ~1.6x dedup+compression ratios reported on mixed lab data `[secondary]` (https://github.com/davperez-tech/backup-disaster-recovery).
- 3-2-1 reference architecture used in the field: Copy 1 production VMs; Copy 2 local repository for sub-10-minute RTO file-level recovery; Copy 3 S3 with Object Lock (ransomware-resistant off-site), bucket versioning + Object Lock enabled at creation, least-privilege IAM, all public access blocked `[secondary]` (https://github.com/davperez-tech/backup-disaster-recovery).
- Pricing: Veeam sells per-workload (Veeam Universal License, per-VM/server) subscription and perpetual+S&S; no authoritative 2026 list price captured in this wave — flagged as a gap; Community Edition remains free for up to 10 workloads `[secondary/unverified-pricing]`.
- Positioning vs open source: Veeam covers VMware/Hyper-V/Nutanix/AHV/cloud-native with application-aware processing; Proxmox shops more commonly pair PBS (free) or Veeam agents; cost comparison should include storage, cloud egress and support — see Section 15 matrix.

