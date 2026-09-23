---
id: etape7-phasee-storage-software/00-storage-software/overview
title: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: ["2025-11", "2026-03", "2026-04-06", "2026-08", "2026-08-05", "2026-09", "2026-09-22"]
keywords: ["benchmarks", "distribution", "research"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [1, 45]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 8bff2a15c94584953c88e7f60944680c965ded5832cd652d059908a5229d4fa3
---

# Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)

## File header

- **Scope:** Software-defined storage and turnkey NAS systems relevant to homelab, SMB and enterprise data infrastructure: Ceph (releases, architecture, tuning), rclone, TrueNAS SCALE/CORE, OpenMediaVault, Unraid, QNAP/Synology/Asustor appliance vendors, and a ZFS/OpenZFS deep-dive.
- **Research cutoff date:** 2026-09-22. All prices, versions and claims are stated as of this date unless otherwise marked.
- **Method:** Read-only public web research (vendor docs, release notes, community runbooks, independent comparisons). No live browsing sessions, no accounts, no purchases.
- **Provenance legend:** `[official]` = vendor/project official documentation or announcement. `[vendor-reported]` = vendor statements in interviews/PR not independently verified. `[independent]` = independent testing/benchmarks. `[secondary]` = third-party writeups, blogs, forums, AI-community notes (claims should be cross-checked). `[unverified]` = single-source or speculative claim; do not rely on it.
- **Deliberate overlaps:** backup tooling detail lives in Phase 7D; this file covers only rclone-relevant backup adjacency briefly. Step 8 of the roadmap covers Redis/SQL from the developer angle; Phase 7F covers the ops angle. Phase F1 of Step 6 covered Ceph NVMe-oF gateway overlap from the DPU/NIC angle; this file covers it from the storage-ops angle.

---

## A. Ceph — release train 2024–2026

### A1. Release timeline and support status

- Ceph **Reef (v18.2)** was released in 2023 and reached **EOL in March 2026**; it is no longer a supported upgrade target [secondary].
- Ceph **Squid (v19.2)** was released in late 2024; community support was stated to **end in September 2026**, i.e. effectively EOL at the research cutoff [secondary].
- Ceph **Tentacle (v20.2)** was released in **November 2025** as the then-current stable line [official].
- Tentacle point releases in 2026: **v20.2.1 released 2026-04-06** with EC recovery and BlueFS WAL envelope fixes, plus an overhauled NVMe-oF gateway fast-failover logic [official]; a later Tentacle release dated **2026-08-05** added dashboard hardware-monitoring (node-proxy metrics), CSV host import, TLS 1.3 `ssl_ciphersuites` for the Beast RGW frontend, mTLS Kafka notification auth, and reworked `ceph-volume` OSD-mapper lifecycle handling [official].
- **v20.2.4 (August 2026)** reworked CephX with the **`aes256k` key type**, described as the first new CephX key type in Ceph's history [secondary].
- The next major after Tentacle was expected under the project's alphabetical naming (post-"T" would be "U", e.g. *Undine*), but no official 2026 GA announcement for it had surfaced by the cutoff [unverified].

### A2. Squid (19.2) highlights

- **Automatic read balancing** (`read_balance_score` visible in `ceph osd pool ls detail`): Squid evens out read-primary distribution online, whereas Reef only had the manual offline `pgremapper`/`upmap-read` style mechanism [secondary].
- **BlueStore/RocksDB tuning**: faster OSD startup, better compaction defaults, improved space accounting [secondary].
- **NVMe-oF gateway**: export RBD images as NVMe/TCP targets to non-Ceph clients (VMware, bare metal), positioned as Ceph's enterprise-block push [secondary].
- **RGW multisite**: live resharding and smoother replication [secondary].
- Scrub scheduling made less impactful and more controllable [secondary].
- The **EC default changed to ISA-L** for new pools in the Squid→Tentacle upgrade path (existing pools unchanged) [secondary].
- Final Squid point release observed in research: **v19.2.4** (2026), including BlueFS default WAL volume size raised to 1GB to prevent ENOSPC failures and RGW Kafka integration at internal version 3.9.2 [official].

### A3. Tentacle (20.2) highlights and breaking changes

- **CephFS**: directories can be configured with case-insensitive or normalized entry names; changing `max_mds` on an unhealthy cluster now requires the `--yes-i-really-mean-it` confirmation flag; the FUSE client returns `EOPNOTSUPP` for default-mode `fallocate` [official].
- **Crimson/SeaStore**: SeaStore is deployable alongside Crimson-OSD as a **tech preview** for early testing only, not a production default; community feedback explicitly solicited [official]. Crimson/SeaStore remained non-default through 2026 [secondary].
- **Dashboard**: NVMe/TCP gateway groups with multiple namespaces, multi-cluster management, OAuth 2.0 integration, enhanced RGW/SMB features (multi-site automation, tiering, policies, lifecycles, notifications, granular replication) [official].
- **Integrated SMB**: a new `smb` MGR module (cephadm-based) auto-creates Samba-backed SMB shares on CephFS, supports AD-domain or standalone-user auth, hosts virtual SMB clusters via Samba CTDB, deploys `samba-container` images, and adds a `cephfs-proxy` daemon for Samba↔CephFS scalability; Dashboard-configurable [official].
- **MGR cleanup**: the long-deprecated `restful` and `zabbix` mgr modules (deprecated since 2020) were **removed**; users can now force-disable always-on modules [official]. Community upgrade guides note `zabbix` removal requires switching to external monitoring [secondary].
- **Fast EC optimizations** (opt-in): had PG-inconsistency bugs through 20.2.2; attempts to enable fast EC on non-4K-aligned chunk sizes are now rejected because unaligned chunks perform poorly and trigger bugs [secondary].
- **RGW tenant-level IAM deprecated** in the Tentacle upgrade path; migration to User Account–scoped IAM required if S3 IAM policies are used [secondary].

---

