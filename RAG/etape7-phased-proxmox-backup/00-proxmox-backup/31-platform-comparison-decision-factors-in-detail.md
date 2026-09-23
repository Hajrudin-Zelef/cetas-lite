---
id: etape7-phased-proxmox-backup/00-proxmox-backup/31-platform-comparison-decision-factors-in-detail
title: "31. Platform comparison: decision factors in detail"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["Broadcom"]
dates: []
keywords: ["agent", "cost", "cyber", "datacenter", "license"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [592, 641]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 8aeebc9cea4f8fda70fb42d3854179a0548f128b928cf304782591e0dc31173f
---

# 31. Platform comparison: decision factors in detail

## 31. Platform comparison: decision factors in detail

- Team skills: KVM/Linux teams ramp fastest on Proxmox; Xen concepts (domains, SRs) need retraining; Harvester needs Kubernetes fluency `[secondary]`.
- Windows guests: all three run Windows well with VirtIO/Xen/KubeVirt drivers; vGPU/SR-IOV passthrough supported on Proxmox and XCP-ng `[secondary]`.
- Storage reuse: existing SAN → Proxmox LVM snapshots or XCP-ng SRs; no SAN → Ceph (Proxmox), XOSTOR (XCP-ng), Longhorn (Harvester) `[secondary]`.
- Backup fit: Proxmox+PBS is the cheapest integrated story; XCP-ng needs XO backup or third-party; Harvester leans on Longhorn/S3 backups or Kasten `[secondary]`.
- Multi-site: PDM 1.1 centralizes Proxmox sites; XO manages many pools from one XOA; Harvester uses Rancher multi-cluster `[secondary]`.
- Cost modeling: Proxmox per-socket vs XCP-ng per-host vs Harvester free+support; a 3-node dual-socket Proxmox Standard cluster ≈ €3,300/yr subscriptions before hardware `[official/secondary]`.
- VMware exit: ESXi import (Proxmox) and VDDK warm migration (XO) both avoid OVA export round-trips `[official/secondary]`.
- Compliance: Secure Boot + vTPM on Proxmox; Shielded VMs concept on XCP-ng; confidential-computing features vary — verify per workload `[secondary]`.
- Ecosystem: Proxmox has the largest homelab/community content volume in 2026; XCP-ng has the most VMware-like operational model; Harvester is the Kubernetes-native bet `[secondary]`.
- Risk: Proxmox = single vendor (GmbH, Vienna); XCP-ng = Vates (France); Harvester = SUSE — all smaller than Broadcom; weigh vendor viability explicitly `[secondary]`.

## 32. Veeam v13 feature inventory (itemized)

- Immutable-by-default backup chains on hardened and object repositories `[secondary]` (https://vinfrastructure.it/2025/11/security-and-anti-ransomware-features-in-veeam-backup-replication-13/).
- "Immutable Until" column in the console for object storage and hardened repositories `[secondary]` (https://forums.veeam.com/object-storage-as-backup-target-f52/immutable-flag-duration-somewhere-visible-in-veeam-t69753.html).
- AI Malware Analysis agent with classification and guided remediation `[secondary]`.
- Proactive investigation: automatic signature scans of suspicious objects; auto-mark clean `[secondary]`.
- Inline entropy analysis during backup streams `[secondary]`.
- Guest filesystem activity analysis `[secondary]`.
- IoC Tools Scanner (MITRE ATT&CK indicators, v12.3+, Advanced license) `[secondary]`.
- Linux-based backup server role `[secondary]` (https://vmorecloud.com/download-veeam-backup-replication-v13/).
- Direct-to-object backup (S3, Azure Blob) without a performance tier `[secondary]`.
- Storage snapshot integration for backup-from-snapshot `[secondary]`.
- Secure HTTPS web console with modern UI `[secondary]`.
- RBAC improvements for delegated administration `[secondary]`.
- Kubernetes and cloud-native workload protection `[secondary]`.
- CrowdStrike / Palo Alto / Splunk / ServiceNow integrations `[secondary]`.
- Scality ARTESCA unified appliance (VBR + S3 object storage, CORE5) `[secondary]` (https://www.vladan.fr/veeam-vbr-13-and-scality-artesca-integration-a-game-changer-for-cyber-resilient-backups/).
- Community Edition: free protection for small estates (10-workload band per community reports) `[secondary]`.
- Four-eyes authorization for destructive operations (carried from v12) `[secondary]`.
- YARA rule support for custom malware scanning (v12+) `[secondary]`.
- Encrypted backups at rest and in transit; lost-password recovery via enterprise key management `[secondary]`.

## 33. Monitoring and alerting stack for Proxmox estates

- Built-in: task log, cluster log, notification targets (SMTP/Gotify/webhook) with severity matchers `[secondary]`.
- Metrics: InfluxDB/Graphite metric servers; `pve-metrics` collection; PDM 1.1 central dashboards with gauges and maps `[official]` (https://pdm.proxmox.com/docs/roadmap.html).
- Prometheus: `pve-exporter` exposes node/VM/CT/storage metrics; pair with Grafana dashboards (community templates exist) `[secondary]`.
- Ceph: `ceph -s`, `ceph health detail`, MGR dashboards; PDM 1.1 surfaces Ceph health centrally `[official]` (https://proxmox.com/en/about/company-details/press-releases/proxmox-datacenter-manager-1-1).
- ZFS: `zpool status`, `zpool iostat -v`, scrub schedules; alert on degraded/failed vdevs `[secondary]`.
- Backup: PBS task history + datastore usage; alert on failed verify jobs and missed backup windows `[secondary]`.
- Database: `pg_stat_archiver` (failed WAL archives), `pgbackrest info --output=json` for stanza status; XtraBackup exit codes + `xtrabackup_info` `[secondary]`.
- Restic/Borg: wrapper scripts emitting Prometheus textfiles or Gotify alerts on non-zero exit `[secondary]`.
- Log aggregation: rsyslog/syslog-ng forwarding to Loki/ELK/Graylog; Proxmox firewall logs per rule when enabled `[secondary]`.
- Uptime: Blackbox-style checks against the PVE API (`/api2/json/version`) and PBS API for synthetic monitoring `[secondary]`.
- Capacity alerts: datastore/ZFS/Ceph at 80% warn, 90% critical; backup repos need GC headroom `[secondary]`.
- Certificate expiry: PVE/PBS web certificates and S3 credential rotation tracked in the runbook `[secondary]`.

