---
id: etape7-phasee-storage-software/00-storage-software/j-decision-guides-and-comparison
title: "J. Decision guides and comparison"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: ["2026-03"]
keywords: ["benchmark", "benchmarks", "consumer", "dram", "latency", "liability", "pricing", "throughput"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [330, 401]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 739f1f26ee885cf7db55a7e4a95c3fd59d00f3eab3b717d7f6c7a6740f58a75d
---

# J. Decision guides and comparison

## J. Decision guides and comparison

### J1. Distributed vs single-box storage

| Need | Best fit | Why |
|---|---|---|
| Multi-petabyte, block+file+object, strong consistency | **Ceph** (RBD/CephFS/RGW) | Unified, EC, self-healing; needs expertise and ≥~10 nodes [secondary] |
| Kubernetes persistent volumes | **Ceph via Rook** (RBD/CephFS CSI) | Native CSI, snapshots, RWX via CephFS [secondary] |
| S3 object on-prem, simple ops | **Ceph RGW** or MinIO/Garage | RGW = Ceph-integrated; MinIO = single-binary simplicity [secondary] |
| Single-box ZFS integrity | **TrueNAS CE** | Free, polished, ZFS-native [secondary] |
| Mixed-size disks, media server | **Unraid** | Parity array flexibility, best Docker UX; $49–249 [secondary] |
| Debian box + web UI, low power | **OpenMediaVault** | Free, light, plugin ZFS [secondary] |
| Set-and-forget appliance | **Synology/QNAP/Asustor** | DSM/QTS/ADM polish; watch drive-lock-in and exposure hardening [secondary] |

### J2. Ceph pool design cheat sheet

- Default for latency-sensitive/RBD/DB: **replicated size=3, min_size=2**, failure domain `host` [secondary].
- Cold/object/archive: **EC 4+2 or 8+3** (k/(k+m) usable); RGW data pool only; metadata/index pools stay replicated [official][secondary].
- Never put hardware RAID under OSDs; use separate 10–25 GbE public/cluster networks [official][secondary].
- PG target 100–200 per OSD; leave the autoscaler on [secondary].
- HDD OSDs: DB on NVMe at 4% (1% for pure block), WAL small; all-NVMe: colocate [official][secondary].
- Enterprise NVMe with PLP for OSD duty; consumer DRAM-less drives are a liability [secondary].

### J3. ZFS pool design cheat sheet

- Mirrors for performance/IOPS; RAIDZ1/2/3 for capacity; dRAID for wide greenfield arrays with fast resilver [secondary].
- RAM first (ARC); L2ARC only when ARC-hit is low and working set > RAM; SLOG only with PLP SSDs for sync-heavy NFS/iSCSI/VM [secondary].
- Special vdev (mirrored) for metadata/small blocks; Fast Dedup quota-bounded, only for dedup-friendly workloads [secondary].
- RAIDZ expansion exists (2.3+) but cannot change parity level; plan parity at creation [secondary].
- Scrub monthly-ish off-peak; snapshot before upgrades; replicate offsite (3-2-1) [secondary].

### J4. NAS hardening checklist (all vendors)

- Never expose the NAS management UI to the internet; disable port forwarding/UPnP and cloud-relay services you don't use (EZ Connect, QuickConnect relay if unused) [secondary].
- Change default ports (8000/8001, 80/443), disable SSH/Telnet/SFTP when not needed, enforce 2FA and strong admin passwords [secondary].
- Keep the OS and all apps (Photo Station, Plex, etc.) patched — DeadBolt/QSnatch both rode unpatched services [secondary].
- Snapshots are not backups: keep versioned offsite/offline copies (USB rotation, cloud, second NAS); ransomware encrypts mounted/cloud-synced copies too [secondary].
- TrueNAS/OMV/Unraid: same rules apply to any internet-facing service (Nextcloud, Jellyfin, etc.) [secondary].

---

## K. Conflicts, gaps and unverified claims

### K1. Conflicts observed

1. **Unraid pricing**: legacy docs show $59/89/129 lifetime tiers; current (post-2024) pricing is $49/$109/$249 + optional $36/yr. The legacy figures are superseded, not concurrent [secondary].
2. **OpenMediaVault version**: some 2026 sources discuss OMV 7.x; the current release is **OMV 8 "Synchrony"** (Dec 2025, Debian 13). OMV 7 references are stale [secondary].
3. **Ceph Reef EOL**: stated as March 2026 by a community runbook; official ceph.io EOL pages were not directly opened in this pass — treat the month as [secondary], not confirmed official.
4. **TrueNAS 26 REST API removal**: reported by community docs ("deprecated in 25.04, removed in 26"); no direct iXsystems release-notes page opened in this pass [secondary].

### K2. Gaps (not covered in this pass)

- G1: QNAP 2026 model lineup and street pricing (TVS/TS-x64 refreshes) — not captured.
- G2: Asustor 2026 lineup/pricing — not captured.
- G3: TrueNAS Enterprise appliance list pricing — quote-only, no public figures.
- G4: Independent 2026 Ceph benchmarks (Squid vs Tentacle, EC vs replication throughput) — no rigorous third-party benchmark set found.
- G5: Independent 2026 NAS throughput comparisons (DS1825+ vs QNAP/Asustor equivalents) — not captured.
- G6: rclone release version at cutoff (v1.x) — version number not pinned in this pass.
- G7: Ceph "U-release" (post-Tentacle) name and date — unannounced at cutoff.
- G8: OpenZFS 2.4 final feature list — TrueNAS 26 target, not released at cutoff.
- G9: Synology DSM 7.3 feature detail — not covered in depth.
- G10: rclone crypt vs restic encryption audit status — no audit claims made.

### K3. Explicitly unverified items

- Post-Tentacle Ceph release name/date [unverified].
- CERN deployment size figures [unverified] (only qualitative "one of the largest" captured).
- TrueNAS Enterprise pricing [unverified].
- QNAP/Asustor 2026 street pricing [unverified].

---

