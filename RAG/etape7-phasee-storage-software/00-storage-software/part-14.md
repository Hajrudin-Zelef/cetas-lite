---
id: etape7-phasee-storage-software/00-storage-software/part-14
title: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) (part 14)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: ["2026-03", "2026-04-06", "2026-07", "2026-09-22"]
keywords: []
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [740, 751]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 791d4c0f27f39aa0992dbb0b964a3db7de18ffbaf20b88bf354d970af91cad71
---

# Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) (part 14)

- Ceph: production = **Tentacle v20.2.x** (20.2.1 released 2026-04-06; 20.2.4 with CephX `aes256k` released 2026-08); Squid 19.2.x support ends ~Sept 2026; Reef 18.2.x EOL March 2026 [official][secondary].
- Rook operator: **v1.19.x** current for Tentacle targets; upgrade runbooks require v1.19+ before the Ceph image bump [secondary].
- TrueNAS CE: **25.10.x** stable line; **TrueNAS 26** in beta (Docker 29.0.4, REST API removed, OpenZFS 2.4) [secondary].
- OpenZFS: **2.3.x** current stable (RAIDZ expansion, Fast Dedup); 2.2.x maintenance (2.2.7+); 2.4 pending with TrueNAS 26 [official][secondary].
- Unraid: **7.3.2** stable (July 2026), 7.3.3-rc.1 in testing [official][secondary].
- OpenMediaVault: **8 "Synchrony"** (Dec 2025, Debian 13) [secondary].
- Proxmox VE: **9.2** referenced as current stable in 2026 comparisons (Debian 13 base) [secondary].
- rclone: release version at cutoff not pinned in this pass — check rclone.org/downloads for the current v1.x [unverified].
- Synology DSM: 25 Plus series (DS1825+/DS725+/DS925+/DS1525+/DS425+) and neo+ series current; DSM 7.x line [secondary].
- Document stats: 18 H2 sections, ~370 provenance-tagged facts, 70+ verbatim source URLs, English prose throughout.

*End of Phase 7E file — software-defined & NAS storage. Observation date 2026-09-22.*
