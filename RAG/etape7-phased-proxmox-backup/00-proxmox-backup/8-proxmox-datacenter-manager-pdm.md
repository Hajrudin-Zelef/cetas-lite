---
id: etape7-phased-proxmox-backup/00-proxmox-backup/8-proxmox-datacenter-manager-pdm
title: "8. Proxmox Datacenter Manager (PDM)"
domain: step-7-phase-d-proxmox-ve-virtualization-platforms-backup-re
role: deep-dive
task: reference
actors: ["AWS", "Broadcom", "Microsoft"]
dates: ["2024-10-07", "2025-07-31", "2025-11", "2026-05-28"]
keywords: ["datacenter", "agent", "agents", "aws", "cost", "cyber", "gpu", "license", "open source", "parameters", "pricing", "training"]
source: docs/RAG/etape7_phaseD_proxmox_backup.md
source_anchor: ""
source_lines: [112, 172]
section: "Step 7 — Phase D: Proxmox VE, Virtualization Platforms & Backup — Research"
sha256: 45ed1a8dd1570224ef3ca91e2eceb12795ec360ff5fc45fa6c038ee60c771c6c
---

# 8. Proxmox Datacenter Manager (PDM)

## 8. Proxmox Datacenter Manager (PDM)

- Proxmox Datacenter Manager is the centralized management plane for distributed Proxmox infrastructures (multiple VE clusters, PBS instances, sites) — positioned as the "single pane of glass" / vCenter alternative for Proxmox estates `[secondary]` (https://www.starwindsoftware.com/blog/proxmox-datacenter-manager-1-1-vcenter-alternative/).
- PDM 1.0 was the initial release line; **PDM 1.1 released May 28, 2026** `[official]` (https://proxmox.com/en/about/company-details/press-releases/proxmox-datacenter-manager-1-1).
- PDM 1.1 highlights `[official]` (https://www.proxmox.com/images/proxmox/press/Release-Proxmox-Datacenter-Manager-1-1-20260528-en.pdf):
  - Automated installation workflows: PDM acts as central configuration server for provisioning; answer-file configurations with predefined installation parameters served to remotes for unattended installs; install progress tracked in the PDM web UI; token-based security so prepared configs are accessed only by authorized installs.
  - Centralized subscription management: subscription registry with a central pool of keys assignable to specific remotes; answer files can embed a key so a new node registers automatically.
  - Unified Ceph cluster monitoring: health, capacity, performance, MON/MGR/OSD and flag status across multiple Ceph clusters on connected hyper-converged VE remotes.
  - Expanded central guest and snapshot management; world-map view of remote locations; local metric collection of the PDM host itself.
  - Base: Debian 13.5, Linux kernel 7.0 stable default, OpenZFS 2.4.2.
- Roadmap note: active-standby architecture for PDM itself is under evaluation to remove the single point of failure; two instances side-by-side already work in practice at the cost of doubled metric collection `[official]` (https://pdm.proxmox.com/docs/roadmap.html).
- PDM access: included in VE/PBS Basic and higher subscriptions (Enterprise Repository + manager) per the 2026 pricing pages; Community tier lists "No support" for the manager `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).
- Third-party operations commentary: PDM 1.1 solves visibility/standardization across sites, but operating model (ownership, build standards, subscription assignment) still needs to be defined by the team — automation makes a standard repeatable, it does not decide the standard `[secondary]` (https://medium.com/@computerportitsolutions/proxmox-datacenter-manager-1-1-what-changes-beyond-one-cluster-computer-port-it-solutions-5b212fba6a0c).

## 9. Hypervisor alternatives

### 9.1 XCP-ng + Xen Orchestra (Vates)

- XCP-ng is a community fork of Citrix XenServer (Xen hypervisor) maintained by Vates (France); free to use with optional paid support `[secondary]` (https://www.vladan.fr/vmware-to-vates-migration-tool-for-xcp-ng-and-xen-orchestra/).
- **XCP-ng 8.3** released 2024-10-07; XCP-ng 8.2 LTS "remains the gold standard" for stability with monthly security updates, while 8.3 evolves faster with more features `[vendor-reported]` (https://xcp-ng.org/blog/2024/10/07/xcp-ng-8-3/?ref=xen-orchestra.com).
- XCP-ng 8.3 was the first release under Vates' consolidated business model (no per-socket/per-core pricing), credited by the vendor for growth following the Broadcom/VMware licensing changes `[vendor-reported]` (https://xcp-ng.org/blog/2024/10/07/xcp-ng-8-3/?ref=xen-orchestra.com).
- Xen Orchestra (XO) is the web management plane (VM on the pool), with XO Lite as a lightweight embedded UI; backup, ACLs, load balancing and REST API live in XO `[secondary]` (https://github.com/xcp-ng/xcp-ng-org/blob/HEAD/docs/intro.md).
- Vates VMS bundles (announced 2024, latest public figures found): Vates VMS Pro €1,000/host/year (XCP-ng Standard Pro Support + XO Enterprise); Vates VMS Enterprise €1,800/host/year (Enterprise support + XO Premium); Vates Essential €2,000/year flat for small estates (Standard support + XO Starter); EUR/USD parity pricing; bundles target 3+ hosts `[secondary]` (https://www.itspyworld.com/2024/08/xen-orchestra-590-released-xcp-ng-830.html). Current 2026 pricing should be re-verified on vates.tech — flagged as possibly superseded `[unverified-current]`.
- Xen Orchestra pricing is per XOA instance: one XOA subscription covers the entire infrastructure regardless of pool/host count (forum clarification by Vates staff) `[secondary]` (https://xcp-ng.org/forum/topic/3170/xen-orchestra-editions/4).
- Migration story: "VMware to Vates" tool in XO uses VMware VDDK for conversion with warm migration support, reading only allocated blocks (faster transfers, no temporary NFS/VSAN staging), handling disks beyond 2 TB `[secondary]` (https://www.vladan.fr/vmware-to-vates-migration-tool-for-xcp-ng-and-xen-orchestra/).
- Feature parity claims vs vSphere: live migration, HA, GPU passthrough, XOSTOR (hyper-converged storage add-on) and XO Proxy (backup proxy) as options `[secondary]`.
- Vates reports Gartner guide inclusion three years running and a training/certification program `[vendor-reported]` (https://xcp-ng.org/blog/2024/10/07/xcp-ng-8-3/?ref=xen-orchestra.com).

### 9.2 Harvester HCI (SUSE/Rancher)

- Harvester is an open-source hyper-converged infrastructure platform built on Kubernetes: Elemental/SUSE Linux Micro immutable OS, KubeVirt for VM management on KVM, Longhorn for distributed block storage, Kube-OVN networking, Grafana/Prometheus observability `[secondary]` (https://github.com/harvester/docs/blob/HEAD/docs/index.md).
- Positioned as the cloud-native HCI alternative for operators wanting VMs and containers on one Kubernetes-native platform, with Rancher integration for multi-cluster management `[secondary]` (https://github.com/harvester/harvester/blob/HEAD/README.md).
- **Harvester v1.7.2** is the latest patch line observed; v1.7.0 component versions: KubeVirt v1.6.3→v1.6.6, Longhorn v1.10.1→v1.10.2, CDI v1.62→v1.64, Kube-OVN v1.14.10, embedded Rancher v2.13.0→v2.13.3, RKE2 v1.34.2→v1.34.9, SUSE Linux Micro 6.1 `[official]` (https://github.com/harvester/release-notes/blob/HEAD/v1.7.2.md).
- RKE reached end of life 2025-07-31; Harvester v1.6.0+ supports RKE2 only (replatforming required, no in-place upgrade from RKE) `[official]` (https://github.com/harvester/release-notes/blob/HEAD/v1.5.0.md).
- Lab/practical notes: single-node installs work on mini-PCs (v1.7 demonstrated air-gapped on x86_64 mini-PC with 16–32 GB RAM, UEFI boot, static networking via cloud-init config); production HA wants 3+ nodes; resource footprint is heavier than Proxmox/XCP-ng per node (8 GB+ RAM minimum commonly cited) `[secondary]` (https://medium.com/@0.all_existence.0/kubernetes-meets-virtualization-installing-harvester-on-bare-metal-ed04c3ac10bc).
- Trade-offs: complete HCI (VM+container) with web UI and no CLI required, but steeper learning curve, larger attack surface and overkill for container-only workloads `[secondary]`.

### 9.3 oVirt and others

- oVirt (upstream of Red Hat Virtualization) has been in long-term decline since Red Hat deprecated RHV (announced 2021, end of maintenance approaching mid-2020s); community activity continues but it is no longer a recommended greenfield choice — no verified 2026 release data found, flagged as a gap `[unverified]`.
- Nutanix AHV and Microsoft Hyper-V/Azure Stack HCI remain in the enterprise consideration set, and Broadcom's VMware Cloud Foundation per-core subscription model is the cost driver behind 2026 Proxmox/XCP-ng migrations; no new primary data collected in this wave `[secondary/context]`.
- KVM-native lightweight alternatives (plain libvirt/QEMU, OpenNebula) exist but were outside this wave's scope; Proxmox's own unattended-install and Terraform/OpenTofu providers cover the IaC story for greenfield KVM shops `[official]` (https://proxmox.com/en/products/proxmox-virtual-environment/pricing).

## 10. Veeam — Backup & Replication v12/v13

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

