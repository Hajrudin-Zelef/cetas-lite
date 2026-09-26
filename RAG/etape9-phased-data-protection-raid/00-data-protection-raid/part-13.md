---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-13
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 13)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Stripe"]
dates: []
keywords: ["compute", "cost", "energy"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [300, 333]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 5872b3355a8f7c333310b3ab0675e5273507853a1f406f471d737f37f7496377
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 13)

| Level | Min disks | Usable capacity | Fault tolerance | Write penalty | Typical use 2026 |
|---|---|---|---|---|---|
| 0 | 2 | 100% | 0 | 1 | Scratch, caches (backed up) |
| 1 | 2 | 50% | 1 | 2 | Boot, small critical volumes |
| 5 | 3 | (n−1)/n | 1 | 4 | Legacy; unsafe on large HDDs |
| 6 | 4 | (n−2)/n | 2 | 6 | HDD bulk, SSD general |
| 10 | 4 | 50% | ≥1 per mirror | 2 | VM datastores, databases |
| 50 | 6 | (n−k)/n, k=#subarrays | 1 per subarray | 4 | Large HDD pools (legacy) |
| 60 | 8 | (n−2k)/n | 2 per subarray | 6 | Large HDD pools |
- **RAID 5 write hole:** a crash during a partial-stripe write leaves data and parity inconsistent with no record of which is correct. Mitigations: battery/supercap-backed controller journal, ZFS copy-on-write (no overwrite → no hole), mdadm write-intent bitmap. Never run RAID 5/6 write-back without cache protection [secondary].
- **Rebuild time math:** 20 TB HDD at ~250 MB/s sustained ≈ 22 hours minimum per disk, during which the array runs degraded and every surviving disk is read in full — the window where UBER math (D5) decides whether RAID 5 survives. SSD rebuilds are faster but still saturate the controller [secondary].
- **Stripe size guidance:** 64–256 KB typical; align with workload (large stripes for streaming/media, smaller for OLTP). Misaligned stripes amplify the write penalty further [secondary].
- **Patrol read / consistency check:** schedule weekly patrol reads and monthly consistency checks on hardware RAID; they convert latent sector errors into reallocated sectors *before* a rebuild needs those sectors [secondary].
- **SSD-specific RAID notes:** URE-during-rebuild risk is negligible at 10^17 UBER; TRIM/discard is often *not* passed through hardware RAID (check per controller) — compensate with higher OP or periodic offline TRIM; RAID 10 vs RAID 6 on SSDs is a cost/performance trade, not a safety one [secondary].
- **Hot spare policy:** dedicated vs global spares; enable auto-rebuild and copyback (return to the original disk after replacement) to keep slot topology stable; on SSD arrays prefer distributed spare capacity so wear leveling stays even [secondary].
- **Controller cache policy per virtual disk:** write-back with read-ahead for sequential workloads; write-back with no read-ahead for random; write-through (or direct) only when cache is unprotected or data is already replicated above [secondary].
- **RAID 10 vs RAID 6 IOPS example (8× SSD, 10K random-write IOPS each):** RAID 10: (8/2 mirrors × 10K)/2 penalty = 20K write IOPS. RAID 6: (6 data × 10K)/6 penalty = 10K write IOPS. Same disks, 2× write-IOPS difference — the penalty table made concrete [secondary].
- **RAID 60 capacity example:** 16 disks as 2× 8-disk RAID 6: usable = (8−2)×2 = 12 disks; survives any 2 failures per subarray (up to 4 total if spread). Rebuild reads 7 disks instead of 15 — halved exposure vs one wide RAID 6 [secondary].
- **Rebuild priority tradeoff:** high rebuild priority shortens the degraded window (good for UBER math) but starves application I/O; low priority does the reverse. On large HDD RAID 6, bias toward rebuild — the array is one more failure from data loss [secondary].
- **Why RAID 5 persists anyway:** small arrays (3–4 disks) of SSDs where UBER math is safe, boot volumes, and legacy estates. The rule is not "never RAID 5" — it is "never RAID 5 on large HDDs" [secondary].
- **RAID 50 capacity example:** 12 disks as 3 × 4-disk RAID 5: usable = (4−1)×3 = 9 disks; survives 1 failure per subarray. Largely superseded by RAID 60/10 for HDDs, still seen in legacy video estates [secondary].
- **Nested-level failure domains:** RAID 60's advantage is not total parity disks but *rebuild scope* — a rebuild reads only its subarray. Size subarrays so a rebuild finishes inside your risk window (see D17 Examples 2–5) [secondary].
- **Minimum-disk guidance:** RAID 5: 3 (practical minimum 4–5 for sane overhead); RAID 6: 4 (practical 6–12); RAID 10: 4; RAID 60: 8. Below these, the overhead or fault-domain math doesn't justify the level [secondary].
- **SSD RAID 5/6 parity CPU:** on hardware ROC the XOR is offloaded; on mdadm it's host CPU — at 8+ SSDs measure `md0_raid6` thread CPU before assuming headroom [secondary].
- **mdadm chunk size:** `--chunk=512K` typical for HDD arrays; match to filesystem record/block where possible. Wrong chunk size is a permanent layout tax like stripe size on HW RAID [secondary].
- **RAID calculator discipline:** compute usable capacity, fault tolerance, and write penalty *before* buying disks — the number of surprises at 3 a.m. is inversely proportional to spreadsheet time [secondary].
- **Rebuild verification:** after any rebuild, run a consistency check / scrub — a rebuild that "completed" with a latent bad block elsewhere has just copied corruption into the new member [secondary].
- **Dual parity beyond RAID 6:** RAIDZ3 (ZFS, 3 parity) and Ceph EC k+3 exist for very large HDD vdevs where even RAID 6's rebuild window feels thin [secondary].
- **Uneven disk sizes:** RAID usable capacity is bounded by the smallest member — a 20 TB disk in a 12 TB array wastes 8 TB. Standardize disk sizes per array [secondary].
- **RAID 1E / RAID 1n (vendor variants):** some controllers offer interleaved-mirror variants for odd disk counts — non-standard, complicates migration; prefer standard levels [secondary].
- **Write journaling on Linux MD:** `mdadm --write-journal` on a fast device mitigates the RAID 5/6 write hole for md arrays — the software analog of CacheVault [secondary].
- **PLP in RAID rebuilds:** a power event during rebuild with PLP drives + CacheVault resumes cleanly; without either, expect to restart the rebuild and re-verify the array [secondary].
- **Capacitor tech choice:** tantalum-polymer for wide-temp industrial (ATP's choice), supercaps for high-energy enterprise flush — the application picks the chemistry [secondary].

