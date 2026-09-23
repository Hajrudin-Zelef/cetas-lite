---
id: etape7-phasee-storage-software/00-storage-software/n-nas-vendors-and-zfs-further-detail-expansion
title: "N. NAS vendors and ZFS — further detail (expansion)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: ["Anthropic", "Google", "Microsoft"]
dates: ["2026-04", "2026-04-25"]
keywords: ["agent", "apache", "attention", "claude", "compute", "license", "licenses", "mcp", "pricing", "revenue"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [532, 580]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 93cdb44806769d43d94232033943dbaf27bb57cb156884e3acace1312841a9ec
---

# N. NAS vendors and ZFS — further detail (expansion)

## N. NAS vendors and ZFS — further detail (expansion)

### N1. Synology DSM applications in depth

- **Active Backup Suite**: centralized backup for Windows/Linux/macOS endpoints, VMs (VMware/Hyper-V), and SaaS (Microsoft 365/Google Workspace); supports 3-2-1 with offsite copies; free of license fees on Plus/XS models (device CALs apply on some tiers) [secondary].
- **Synology Drive**: private-cloud file sync/share with cross-platform clients, on-demand sync, and site-to-site sync (ShareSync) for distributed teams [secondary].
- **Surveillance Station**: on-prem NVR; camera licenses bundled (typically 2 free) with additional licenses purchasable; multi-site monitoring [secondary].
- **Container Manager**: Docker/Compose support on Plus and higher models; used for self-hosted apps (Jellyfin, Immich, etc.) though RAM ceilings on non-expandable models limit it [secondary].
- **Hyper Backup**: versioned backup to local/remote/cloud with client-side encryption and deduplication [secondary].

### N2. QNAP QuTS hero and 2026 positioning

- **QuTS hero** is QNAP's ZFS-based OS for higher-end models (TVS-h series etc.), offering WORM, inline dedup/compression, and snapshots — QNAP's answer to TrueNAS for data-integrity buyers [secondary].
- QTS remains the mainstream OS for the volume lineup (TS-x53D, TS-x64, TS-hx77XU families); specific 2026 refresh pricing was not captured in this pass — see gap G1 [unverified].
- QNAP's security track record (DeadBolt, QSnatch) makes the hardening checklist (section J4) especially relevant to QNAP owners; QNAP publishes its own security best-practice guidance [secondary].

### N3. Asustor ADM applications

- ADM includes App Central (Plex, Docker via Portainer, etc.), EZ Connect remote access, and Snapshot Center (Btrfs snapshots on supported models) [secondary].
- Asustor's post-DeadBolt guidance (disable EZ Connect, change default ports, disable SSH/SFTP, keep backups) remains the reference hardening set for ADM [secondary].

### N4. UGREEN and UniFi disruptors

- **UGREEN NASync** (from $374, no license): strong hardware value, Docker-capable, keeps OS options open (users commonly install TrueNAS) [secondary].
- **UniFi UNAS Pro** ($499, 7-bay rackmount, no license): storage-and-backups appliance with minimal management attention; no app/Docker story — pair with a mini PC for compute [secondary].
- **HexOS** ($199 early access, $299 after): TrueNAS-based simplified layer ("TrueNAS without learning it"), one-drive-at-a-time expansion [secondary].

### N5. ZFS replication, send/receive and tooling

- `zfs send | zfs receive` is the native replication primitive: full + incremental streams, `-w` raw (encrypted) sends, resume support (`-s`) for interrupted transfers [secondary].
- TrueNAS replication tasks, Proxmox PBS-adjacent ZFS replication, and homelab tools (**Sanoid/Syncoid**) build on send/receive for automated snapshot + replication policies [secondary].
- `zfs snapshot -r pool@name` recursive snapshots; snapshot retention policies in TrueNAS periodic-snapshot tasks [secondary].

### N6. Distributed-storage alternatives in brief

- **GlusterFS**: Red Hat deprecated GlusterFS in favor of Ceph for new deployments; community use is in long-term decline — included here only as a migration source, not a recommendation [secondary].
- **Linstor/DRBD**: replicated block for HA pairs (2–3 nodes); simpler than Ceph for small HA VM storage [secondary].
- **MinIO/AIStor, SeaweedFS, Garage**: lightweight S3 object stores; a 2026 community comparison notes MinIO community was archived 2026-04-25 (AIStor proprietary continues), Garage AGPLv3 replication-only, SeaweedFS Apache-2.0 with paid tier beyond 25 TB free [secondary].
- **Longhorn/OpenEBS**: Kubernetes-native block for small clusters; simpler than Rook/Ceph but single-cluster scope [secondary].

### N7. TrueNAS 26 and ecosystem notes

- TrueNAS 26 (beta April 2026): annual-major cadence begins; Docker Engine 29.0.4; REST API removed (WebSocket/`midclt` only); Linux 6.18 LTS; OpenZFS 2.4 [secondary].
- **TrueNAS Connect**: Foundation (free) / Plus ($60/yr/3 systems) for SSO and fleet monitoring — iXsystems' attach-revenue play for the CE base [secondary].
- Community MCP/API tooling (e.g. a 2026 Claude-oriented MCP server exposing 18 read-only TrueNAS tools) shows the WebSocket API is scriptable enough for agent-driven ops [secondary].
- Version-detection discipline matters: Apps/VM backends changed twice (k3s → Docker; libvirt API namespaces), so runbooks must read `system.version` before choosing method names [secondary].

---

