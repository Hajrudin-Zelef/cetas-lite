---
id: etape9-phased-data-protection-raid/00-data-protection-raid/overview
title: "Step 9 — Phase D: Data Protection & RAID Hardware"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: ["2026-09", "2026-09-22"]
keywords: ["research"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [1, 11]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: ce7097cf3de66fe9131c6aea77b9581e54ad4ec63c6eab6e2b24aba5032833b9
---

# Step 9 — Phase D: Data Protection & RAID Hardware

- **Scope:** Data Loss Protection (DLP) concepts, end-to-end data path integrity, firmware resilience, secure erase/sanitize, reliability metrics (MTBF/AFR/UBER/SMART), wear leveling and thermal management, hardware RAID controllers (Broadcom MegaRAID 96xx/95xx, Microchip Adaptec SmartRAID 4300), cache protection (CacheVault/supercap), HBAs and SAS expanders, NVMe hardware RAID (GRAID SupremeRAID) vs software RAID, JBOD vs RAID for ZFS/Ceph, tape (LTO-9/LTO-10) and optical archival as last-resort backup, decision matrices for hardware RAID vs ZFS vs Ceph erasure coding.
- **Research date / cutoff:** 2026-09-22. English throughout.
- **Method:** Read-only web research (search + page fetch), September 2026. Vendor datasheets, program announcements, and independent press. Every factual claim carries a provenance tag. Facts contradicted across sources are preserved as conflicts, not resolved by invention.
- **Provenance legend:** `[official]` = vendor/standards-body documentation or press release; `[vendor-reported]` = vendor performance/claim in marketing collateral; `[independent]` = third-party testing or analyst data; `[secondary]` = press, retail listings, community documentation; `[unverified]` = claim seen in only one weak source or not confirmed. `[conflict]` marks contradictory sourced claims and is always paired with a tag.

---

## D1 — Data Loss Protection (DLP): what "protection" means at drive level

