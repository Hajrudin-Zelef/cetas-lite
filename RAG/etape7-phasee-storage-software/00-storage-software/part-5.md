---
id: etape7-phasee-storage-software/00-storage-software/part-5
title: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) (part 5)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["compute"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [184, 188]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: e9bdb4a9aed8d09afc08292f7ae8b5f81103196076b9d547ffa80251e085f444
---

# Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS) (part 5)

- TrueNAS CE: free, ZFS-native (data integrity/silent-corruption protection), Docker apps; weaknesses vs Unraid = no mixed-drive single pool (RAIDZ expands one drive at a time), higher RAM floor (~16 GB practical for ZFS), compute is secondary to storage [secondary].
- Community rule of thumb: **compute-first → Proxmox VE; storage-first → TrueNAS**; a mature pattern is running TrueNAS as a VM inside Proxmox to add ZFS later without a second box [secondary].

---

