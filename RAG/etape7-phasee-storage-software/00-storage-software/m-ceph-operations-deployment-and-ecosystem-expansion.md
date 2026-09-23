---
id: etape7-phasee-storage-software/00-storage-software/m-ceph-operations-deployment-and-ecosystem-expansion
title: "M. Ceph — operations, deployment and ecosystem (expansion)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: ["2026-08"]
keywords: ["consumer", "cost", "distribution", "memory"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [474, 531]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 212dc2a9bbf5ef1c9f58ed948024fea8c24affc5205f742227a5fd12ee29fdd5
---

# M. Ceph — operations, deployment and ecosystem (expansion)

## M. Ceph — operations, deployment and ecosystem (expansion)

### M1. cephadm and deployment models

- **cephadm** is the current official deployment/orchestration tool: it manages Ceph daemons as containers on hosts, replacing the older ceph-deploy/ansible paths [official].
- Typical production layout: 3 or 5 MONs (odd count, on hosts with SSD-backed RocksDB), 2 MGRs (active/standby), N OSDs (one per device), MDS only if CephFS is used, RGW daemons on edge/load-balanced hosts [secondary].
- Rook is the Kubernetes-native alternative: operator + CephCluster CR + CSI drivers; external-cluster mode lets Kubernetes consume a bare-metal Ceph (e.g. on Proxmox hosts) without running Ceph daemons in-cluster [secondary].
- Proxmox VE integrates Ceph management in its UI/API (pveceph), and community runbooks document the Reef→Squid→Tentacle upgrade path for Proxmox-managed clusters including the MON-address/CephFS-client edge cases [secondary].

### M2. Pool and CRUSH operations

- Creating pools: `ceph osd pool create <name> <pg_num>` then `ceph osd pool set <name> size 3` / `min_size 2`; EC pools need a profile: `ceph osd erasure-code-profile set <profile> k=4 m=2` (ISA-L is the default plugin since Tentacle) [secondary].
- CRUSH rule example for rack-level failure domain and EC pools; `ceph osd crush rule create-erasure <name> <profile>` [secondary].
- **PG autoscaler** (`pg_autoscale_mode`) keeps PG counts near the 100–200-per-OSD target automatically; Reef hardened it and Squid/Tentacle continued default-tuning [secondary].
- Upmap/read-balancer: Squid's automatic read balancing is surfaced via `read_balance_score` in `ceph osd pool ls detail`; Reef only had the manual offline path [secondary].

### M3. RGW multisite and S3 surface

- RGW multisite: realms → zonegroups → zones; active-active or active-passive; Tentacle added live resharding and smoother replication flows [official][secondary].
- S3 API coverage includes versioning, Object Lock, lifecycle, bucket policies, IAM users/subusers, STS; Swift API also served [secondary].
- 20.2.x RGW fixes observed: `ssl_ciphersuites` for Beast TLS 1.3, canned-ACL/`BlockPublicAcls` fix, CopyObject for encrypted objects, mTLS Kafka notifications, delete-marker lifecycle fix [official].
- RGW + EC: RGW is the canonical EC-pool consumer (full-object writes); bucket-index and service pools stay replicated [official].

### M4. CephFS operations

- CephFS needs at least two pools (metadata = replicated, data = replicated or EC with `allow_ec_overwrites`); multiple active MDS ranks with subtree pinning for scale-out metadata [secondary].
- Tentacle CephFS changes: case-insensitive/normalized directory entry names, `max_mds` confirmation guard, FUSE `fallocate` behavior [official].
- Known operational sharp edge: wedged CephFS clients after reboots/MON-address changes (documented in Proxmox runbooks) [secondary].
- CephFS snapshots are cheap and instant; the new integrated **SMB module** (Tentacle) exposes CephFS as clustered SMB via Samba CTDB, managed from the Dashboard [official].

### M5. Monitoring, alerting and tooling

- MGR hosts the **Prometheus module** (metrics), **Dashboard**, **balancer**, **pg_autoscaler**, **crash** module; removed in Tentacle: `restful` and `zabbix` modules (use Dashboard API / external monitoring instead) [official][secondary].
- `ceph status`, `ceph osd df tree`, `ceph pg stat`, `ceph health detail` are the daily-driver commands; pre-upgrade backup runbooks snapshot cluster state (status, osd tree, pools, CRUSH dump, config dump, versions) [secondary].
- Node-proxy–backed hardware monitoring was added to the Dashboard in the August 2026 Tentacle release [official].

### M6. Upgrade paths and compatibility

- Supported path: Reef → Squid → Tentacle; skipping a major is not supported. Rook v1.19+ required for Tentacle targets [secondary].
- Breaking changes at Tentacle: tenant-level RGW IAM deprecated (migrate to User Accounts), `restful`/`zabbix` removed, ISA-L default for new EC pools, `osd_repair_during_recovery` option removed [secondary].
- CephX `aes256k` key type (20.2.4, Aug 2026): first new CephX key type in project history — relevant for key-rotation and compliance posture [secondary].

### M7. Commercial ecosystem and support

- **Red Hat Ceph Storage** (IBM/Red Hat) is the dominant enterprise distribution; version 8 hardware guide current in 2026 [official]. IBM's Red Hat ownership continues to shape Ceph's enterprise roadmap [secondary].
- Alternatives: Canonical (Charmed Ceph / MicroCeph for edge), SUSE Enterprise Storage, and specialists **Croit**, **42on**, **SoftIron** (HyperDrive appliances) [secondary].
- SoftIron, Croit and 42on are repeatedly cited as the go-to Ceph support/appliance vendors for teams without in-house Ceph expertise [secondary].

### M8. Ceph TCO and sizing rules of thumb

- Minimum production: 3 MON + 5+ OSD hosts + 2+ RGW ≈ 10 nodes for a serious object cluster; smaller converged/Rook clusters exist for homelab [secondary].
- Memory: ~4–6 GB per OSD baseline, plus MON/MGR/RGW headroom; NVMe-oF gateways and RGW are CPU-hungry [secondary].
- Network: 10 GbE absolute minimum, 25 GbE+ recommended; separate public (client) and cluster (replication/recovery) networks; EC recovery is network-intensive [secondary].
- Capacity math: usable = raw × k/(k+m) for EC; usable = raw/3 for 3x replication; always leave 15–20% headroom before `nearfull` ratios throttle [secondary].
- The dominant TCO driver is operational expertise — mis-tuned PG counts, scrub storms, or recovery throttling are the classic cost multipliers [secondary].

---

