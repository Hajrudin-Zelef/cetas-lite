---
id: etape7-phasee-storage-software/00-storage-software/p-final-additions
title: "P. Final additions"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: ["Google"]
dates: ["2026-03", "2026-08", "2026-09", "2026-09-22"]
keywords: ["distribution", "latency", "license", "research"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [652, 739]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 6c9a0655696aa922b635e65a9b861ed602188bc9c7d3029f8aa5b7c6f5b3b123
---

# P. Final additions

## P. Final additions

### P1. Ceph RGW S3 compatibility and IAM detail

- RGW exposes an S3-compatible API plus a Swift API; compatibility covers the commonly used subset (multipart, versioning, lifecycle, SSE-S3/SSE-C, Object Lock) with documented gaps on exotic headers — always test application-specific calls [secondary].
- IAM model: users, subusers, key pairs, bucket policies, and (current) User-Account–scoped policies; tenant-level IAM was deprecated in the Tentacle cycle [secondary].
- STS (temporary credentials) is supported for federated/OIDC-adjacent flows [secondary].
- **Resharding**: dynamic bucket-index resharding keeps large buckets performant; Tentacle enabled live resharding to avoid downtime windows [secondary].
- **Notifications**: bucket notifications to Kafka/AMQP/HTTP endpoints (used for event-driven pipelines); the August 2026 Tentacle release added mTLS client-certificate auth for Kafka targets [official].
- **DataSync/Cloud transition**: RGW tiering to cloud (S3) for cold data — part of the Dashboard's enhanced tiering/lifecycle features in Tentacle [official].

### P2. Ceph security and compliance posture

- **CephX** authentication between all daemons/clients; the 20.2.4 `aes256k` key type is the first new key type in project history (relevant to FIPS-adjacent and key-rotation discussions) [secondary].
- **dm-crypt/LUKS at the OSD device level** is supported (`ceph-volume` handles encrypted OSDs); Tentacle improved rotational-media detection under dm-crypt to bypass the workqueue correctly [official].
- **RBD LUKS2 per-image encryption** gives tenant-level encryption without device-level LUKS [secondary].
- RGW SSE-S3/SSE-C server-side encryption options; KMIP integration exists for key management in enterprise distributions [secondary].
- Compliance note: regulated deployments typically pair Ceph with a supported distribution (Red Hat/IBM, Canonical, SUSE) rather than upstream builds [secondary].

### P3. TrueNAS replication and backup scheduling detail

- Replication tasks: push/pull ZFS snapshots between TrueNAS systems (or compatible ZFS targets) on schedules, with retention policies and encryption of the stream [secondary].
- Cloud sync tasks: S3, Backblaze B2, Google Drive, Dropbox and others via the built-in cloud credentials; uses rclone under the hood [secondary].
- Periodic snapshot tasks + scrub tasks are the baseline data-protection schedule; alerts fire on failures via the alert system [secondary].
- Boot environments: TrueNAS keeps boot-pool snapshots so a bad update rolls back from the boot menu — the safety net behind the annual-major cadence [secondary].

### P4. Vendor warranty and support comparison (NAS appliances)

- Synology: 3-year warranty standard on Plus models (extendable to 5 with Extended Warranty Plus); support via tickets/KB; drive-compatibility policy ties support to certified media [secondary].
- QNAP: typically 2–3 year warranty depending on series; extended warranty purchasable; support quality varies by region [secondary].
- Asustor: 3-year warranty on most models; post-DeadBolt, firmware/support responsiveness improved but trust recovery is ongoing [secondary].
- TrueNAS: CE is community-supported (forums/Discord); iXsystems appliances carry hardware warranty + optional support contracts; TrueNAS Connect adds fleet monitoring tiers [secondary].
- Unraid: license includes forum/community support; paid support options limited — the community is the support channel [secondary].

### P5. Power, noise and homelab practicalities

- TrueNAS Mini X+: sub-80W typical draw cited at launch; Nanoxia quiet fans; a frequent 2026 homelab pick for a quiet office NAS [vendor-reported][secondary].
- HDD spin-down/idle: TrueNAS/OMV/Unraid all support disk standby timers; ZFS scrubs and SMART tests wake disks — schedule accordingly [secondary].
- UPS integration: TrueNAS, OMV, Unraid, Synology, QNAP and Asustor all support USB UPS (NUT/apcupsd) with graceful shutdown — the cheapest data-loss insurance [secondary].
- 2.5GbE is the 2026 baseline on new prosumer NAS (DS1825+, TrueNAS Mini X+, many QNAP/Asustor); 10GbE via PCIe/SFP+ on expandable models; plan switch ports accordingly [secondary].

### P6. Ceph version-pinning and lifecycle advice

- Pin production to the current stable major (Tentacle at cutoff) and track point releases; Reef EOL March 2026 and Squid support ending September 2026 force upgrades off both [secondary].
- Read the release notes' breaking-changes section before every major upgrade (Tentacle's RGW IAM and module removals are the template example) [official].
- Keep one test/cluster or Rook staging environment to rehearse major upgrades — the community runbooks all assume a rehearsal [secondary].
- For Kubernetes users: align Rook operator versions with the Ceph target (v1.19+ for Tentacle) before upgrading the Ceph image [secondary].

### P7. ZFS future watch (post-cutoff)

- OpenZFS 2.4 (targeted by TrueNAS 26) was not released at cutoff; watch for continued Fast Dedup maturation, RAIDZ expansion polish, and Linux 6.18 support [unverified].
- Persistent L2ARC, allocation-class improvements, and `zstd` early-abort were recurring upstream discussion topics in 2026 [secondary].
- dRAID adoption in TrueNAS UI workflows was still limited at cutoff; CLI-first for dRAID creation [secondary].

---

## Q. Quick-reference checklists

### Q1. Greenfield Ceph checklist

- Size: 3 MON (odd), 2 MGR, ≥5 OSD hosts for production; 10/25 GbE+ with separate public/cluster networks [secondary].
- Media: enterprise NVMe with PLP for OSDs; HDD OSDs get NVMe DB at 4% (1% pure block), WAL small; all-NVMe = colocate DB/WAL [official][secondary].
- Pools: replicated 3/2 for RBD/DB/latency; EC 4+2 or 8+3 for RGW/cold; metadata/index pools always replicated [official][secondary].
- PGs: 100–200 per OSD, autoscaler on; read-balance score monitored post-Squid [secondary].
- Ops: scrubs off-peak, recovery throttled, pre-upgrade state backups, rehearsal environment, Rook version aligned [secondary].

### Q2. Greenfield ZFS checklist

- Topology: mirrors for IOPS, RAIDZ2 default for capacity, dRAID for wide greenfield arrays; parity level fixed at creation [secondary].
- RAM first; L2ARC only when ARC-hit low and working set > RAM; SLOG only PLP-backed for sync-heavy NFS/iSCSI/VM [secondary].
- recordsize/volblocksize matched to workload; compression lz4/zstd; atime off; encryption planned before data lands [secondary].
- Snapshots scheduled, scrub monthly, replication offsite, UPS with graceful shutdown [secondary].
- Fast Dedup only for dedup-friendly data, quota-bounded, DDT monitored [official][secondary].

### Q3. NAS purchase checklist (2026)

- Decide: appliance polish (Synology/QNAP/Asustor) vs DIY flexibility (TrueNAS/Unraid/OMV) vs prebuilt-DIY (UGREEN/UniFi) [secondary].
- Check the drive-compatibility policy before buying Synology 25-series (certified media required for full features) [secondary].
- Budget: diskless appliance + drives + 2.5/10GbE switch ports + UPS + offsite backup [secondary].
- Harden on day one: no WAN exposure, default ports changed, 2FA, updates on, cloud-relay off unless needed [secondary].
- Plan the exit: SHR/ZFS/RAID choices affect migration; keep an offsite copy independent of the NAS vendor [secondary].

*End of file. Phase 7E complete: Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS deep-dive. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*

---

## R. Version-pinning reference (cutoff 2026-09-22)

