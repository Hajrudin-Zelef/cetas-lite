---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d8-hardware-raid-controllers-2026-landscape
title: "D8 — Hardware RAID controllers: 2026 landscape"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom", "Stripe"]
dates: ["2025-08-05"]
keywords: ["accelerator", "benchmark", "compute", "consumer", "cost", "disaggregated", "energy", "latency", "nand", "pricing"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [245, 333]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 91ce82258f8422345b961e26abea29d9a3456822a5e5d50568edae72aac29f24
---

# D8 — Hardware RAID controllers: 2026 landscape

## D8 — Hardware RAID controllers: 2026 landscape

- **Market shape:** two vendors dominate discrete RAID — Broadcom (MegaRAID, ex-LSI) and Microchip (Adaptec SmartRAID/SmartHBA). Motherboard-integrated "fakeraid" and chipset RAID are unsuitable for production [secondary].
- **Broadcom MegaRAID 9600 series (current tri-mode generation):** 3rd-generation tri-mode x8/x16 NVMe/SAS/SATA adapters; PCIe Gen4 host; 24G SAS; models include 9660-16i (05-50107-00, LP-MD2, 4 GB cache), 9670-24i (05-50123-00, based on SAS4124 RAID-on-Chip with 2×72-bit DDR at 3200 MT/s), plus eHBA variants like 9600-16i (05-50111-00); RAID 0/1/5/6/10/50/60 + JBOD [secondary].
- **Broadcom's claimed generational gains (9600 vs 95xx):** 2× bandwidth, up to 2× 4K random-read IOPS, >3× RAID 5 IOPS (4K RW); the 9670-24i listing claims 4×+ RAID 5 random-write IOPS, 25× lower write latency, 60× faster recovery vs previous generation [vendor-reported].
- **Broadcom MegaRAID 95xx (previous gen, still widely deployed):** 12G SAS tri-mode; e.g. 9560-16i (05-50077-00, 8 GB cache per reseller spec) — common in refurbished/secondary market builds [secondary].
- **Indicative street pricing (Newegg listings, 2026 crawl — treat as secondary, volatile):** 9660-16i ~$1,493; 9560-16i ~$1,561; 9600-16i eHBA ~$841; HBA 9500-16e ~$776 [secondary].
- **CacheVault flash cache protection:** supercapacitor + NAND flash module (CVPM02/CVPM05 kits, FBU02 packs). On power loss the supercap powers the card just long enough to dump write-back cache to NAND; on restore, data returns to cache and pending writes resume. Replaces lithium-ion BBUs — no battery replacement cycles, lower TCO, greener disposal [secondary].
- **Microchip Adaptec SmartRAID 4300 (launched 2025-08-05):** NVMe RAID storage *accelerator* with disaggregated architecture — storage software runs on the host CPU, writes go directly CPU→NVMe at native PCIe speed, while parity/XOR offload goes to the accelerator; up to 32 CPU-attached x4 NVMe devices (Gen4/Gen5), 64 logical drives/arrays; RAID 0/1/10/5/50; up to 291 GB/s sequential read claimed at full 32-drive saturation; up to 7× I/O vs previous generation per Microchip internal testing [vendor-reported].
- **SmartRAID 4300 security:** hardware root of trust, secure boot, secure update, attestation, SED support; low-profile MD2, PCIe x16 slot [vendor-reported].
- **Why accelerators exist:** traditional in-line RAID cards bottleneck Gen4/Gen5 NVMe (a single Gen5 x4 SSD ≈ 14 GB/s; 32 of them ≈ 450 GB/s — no x16 ROC can inline that). Out-of-path/disaggregated designs keep data on CPU-attached PCIe lanes and only offload parity compute [secondary].
- **Gap:** no independent (non-vendor) benchmark of the SmartRAID 4300 was found at cutoff; treat all 7×/291 GB/s figures as vendor-reported [unverified].
- **9600-series model table (reseller spec sheets, volatile — verify before PO):** 9660-16i (05-50107-00, LP-MD2, 4 GB cache, x8 PCIe 4.0, 4× SFF-8654, ~$1,493 street); 9670-24i (05-50123-00, SAS4124 RoC, 2×72-bit DDR 3200 MT/s, MD2, ~16 internal ports, RAID 0/1/5/6/10/50/60); 9600-16i eHBA (05-50111-00, 15 W, 5M-hour MTBF claimed, ~$841); 9600-24i eHBA (05-50111-01, 24 ports); 9600W-16e (external SFF-8674). All tri-mode 24G SAS / SATA / NVMe [secondary].
- **95xx previous generation:** 9560-16i (05-50077-00, 8 GB cache per reseller spec, ~$1,561 street — note: previous-gen street price can exceed current-gen; always quote both). 12G SAS tri-mode; still common in production and the refurbished market [secondary].
- **SAS4124 RAID-on-Chip:** the silicon behind the 9670-class cards; Broadcom's claimed gains over 95xx (2× bandwidth, 2× 4K RR IOPS, >3× RAID-5 IOPS, 25× lower write latency, 60× faster recovery) are vendor-reported deltas on Broadcom's test config — useful directionally, not as procurement guarantees [vendor-reported].
- **CacheVault parts landscape:** CVPM05 kits (current, for 95xx/96xx), CVPM02 (older generation), FBU (Flash Backup Unit) packs; Scan UK lists CVPM kits for 9361/9380-class cards, aeonfly lists 05-50038-00 modules. Supercap + NAND replaces Li-ion BBU: no 1–2-year battery replacement cycle, no battery disposal, and cache destage is measured in seconds [secondary].
- **Management tooling:** MegaRAID Storage Manager (GUI) and StorCLI/StorCLI2 (CLI) for VD/PD management, patrol read scheduling, firmware updates, and event log export. Script StorCLI in provisioning — GUI-only shops can't audit at scale [secondary].
- **SED/SafeStore:** Broadcom's SafeStore branding covers SED key management on MegaRAID; combined with CacheVault this gives encrypted-data-at-rest plus power-safe write-back cache on one card [secondary].
- **Microchip Adaptec line (for the decision matrix):** SmartRAID 4300 (2025, NVMe accelerator, disaggregated) sits above SmartRAID 3200/3100 (12G SAS/SATA RAID) and SmartHBA 2200 (HBA line). Microchip's pitch is the same triad as Broadcom's: RAID for SAS estates, HBA for software-defined, accelerator for NVMe [secondary].
- **Lifecycle note:** 94xx/93xx (12G, pre-tri-mode) are legacy — fine on existing estates, but new 2026 purchases should be 95xx-minimum and preferably 96xx for 24G SAS and NVMe density [secondary].
- **StorCLI2 reference (script these, don't click through):** `storcli2 /c0 show` (controller summary); `storcli2 /c0 /eall /sall show` (all physical drives); `storcli2 /c0 /vall show` (virtual drives); `storcli2 /c0 add vd type=raid6 drives=... wb ra` (create VD, write-back + read-ahead); `storcli2 /c0 set patrolread on` ; `storcli2 /c0 show events` (event log export for audit) [secondary].
- **Patrol read and consistency check cadence:** weekly patrol read, monthly consistency check on parity RAID; schedule off-peak with bounded rate so application latency doesn't crater — and never disable them to "save performance" on large HDD arrays [secondary].
- **Firmware update order:** back up config (`storcli2 /c0 show file=...` / MSM backup), update controller firmware, then drive firmware per vendor matrix, then verify VD optimal state. A RAID card firmware flash that resets cache policy to write-through has caused real-world performance incidents — verify policy after every update [secondary].
- **Foreign config discipline:** when moving disks between cards, import foreign config deliberately; auto-import on boot can assemble the wrong array if disks were shuffled. Label slots (D16) so this never happens [secondary].
- **Controller failure planning:** keep a cold-spare identical (or vendor-confirmed compatible) RAID card; array metadata lives on the disks, so a like-for-like card imports the config — a *different-generation* card may not. This is the portability tax hardware RAID pays vs ZFS (D13) [secondary].
- **Warranty and support:** Broadcom MegaRAID cards typically carry 3-year warranties with advance-replacement options; factor the support contract into TCO — a dead RAID card without a contract is a multi-day outage [secondary].
- **96xx vs SmartRAID 4300 positioning:** both target NVMe RAID but differ architecturally — MegaRAID 9600 is an inline ROC (data through the card), SmartRAID 4300 is a disaggregated accelerator (data direct CPU→SSD). At 8–16 drives both work; at 32 Gen5 drives the out-of-path designs have the headroom argument [secondary].
- **Refurbished market caution:** 93xx/94xx cards flood the secondary market; fine for lab/HBA-flash use, but verify capacitor age on any card with a BBU-era design and confirm CacheVault module compatibility before trusting write-back cache [secondary].
- **Tri-mode backplane requirement:** the card is only as tri-mode as the backplane — a SAS-only backplane won't pass NVMe even with a 9600 card. Budget backplane + cables with the card [secondary].
- **9670W-16i (wide MD2):** the 16-port wide-form variant (distrinode leaflet: 05-50123-00 family, 240 SAS/SATA or 24 NVMe direct) for dense 2U NVMe shelves where LP-MD2 doesn't fit the airflow plan [secondary].
- **Legacy 9460/9361 (12G):** still underpinning many production estates; CacheVault CVPM02/FBU-era. Replacement strategy in 2026: 95xx/96xx on controller refresh cycles, not emergency swaps [secondary].
- **ROC cache sizing:** 4 GB (9660-16i) vs 8 GB (9560-16i per reseller spec) — cache size matters for write-back burst absorption; size to the workload's burst, not to the drive count [secondary].
- **PCIe lane negotiation:** an x8 card in an x4 slot (or bifurcated wrong) silently halves bandwidth — verify `lspci` link width/speed after installation; 24G SAS needs the full x8 Gen4 to breathe [secondary].
- **Secure boot chain:** 9600's hardware secure boot verifies the card's own firmware; the host UEFI secure boot verifies the card's Option ROM — both must be enabled for the chain to mean anything [secondary].
- **JBOD mode on MegaRAID:** supported per-VD/per-PD, but remember D12 — JBOD mode on a RAID card is not IT-mode HBA for ZFS purposes. Use eHBA SKUs for passthrough [secondary].
- **Mixed drive types:** tri-mode cards allow SAS HDD + SATA SSD + NVMe in one chassis on separate VDs — but don't mix media in one VD (performance collapses to the slowest member) [secondary].
- **Energy:** 24G parts run hotter than 12G; verify chassis airflow for MD2 cards with cache + supercap — the CacheVault module adds its own thermal load [secondary].
- **Driver in-box vs out-of-box:** 9600-series needs a recent `mpi3mr` driver; LTS distro kernels may lag — check HCL/driver availability for the exact OS version in the deployment plan [secondary].
- **NVMe boot from RAID:** 9600 supports bootable NVMe RAID volumes via UEFI driver — useful for mirrored OS on NVMe without mdadm [secondary].
- **SafeStore key management:** local key management vs remote (KMIP); losing the SED authentication key is equivalent to losing the data — back up keys per D4's escrow warning [secondary].
- **Composite vs sensor temperatures:** NVMe reports a composite plus per-sensor readings where supported; the hottest sensor (often NAND or controller) drives throttling, not the composite — monitor both [secondary].
- **Thermal pad maintenance:** repasting/repadding heatsinks on U.2 carriers during drive replacement restores thermal performance — dried pads are a silent throttling cause [secondary].

## D9 — RAID levels: math, write penalties, and rebuild reality

- **Level summary (n = data disks):** RAID 0 (striping, n-disk capacity, no redundancy); RAID 1 (mirror, 50% capacity); RAID 5 (single parity, n−1 capacity, survives 1 disk failure); RAID 6 (dual parity, n−2 capacity, survives 2); RAID 10 (striped mirrors, 50% capacity, survives ≥1 per mirror pair); RAID 50 (striped RAID 5 sets); RAID 60 (striped RAID 6 sets) [secondary].
- **Write penalties (backend I/Os per host write):** RAID 0 = 1, RAID 1 = 2, RAID 5 = 4 (read data + read parity + write data + write parity), RAID 6 = 6, RAID 10 = 2. Effective random-write IOPS ≈ (single-disk IOPS × data disks) / penalty [secondary].
- **Example:** 8× 10K HDDs in RAID 6: read ≈ 8-disk streaming; random write ≈ (8−2 disks × ~150 IOPS)/6 ≈ 150 IOPS — the classic "RAID 6 is write-slow on HDDs" result. The same array on SSDs is penalty-bound, not seek-bound [secondary].
- **RAID 5/6 on large HDDs — the rebuild problem:** with 20 TB HDDs and consumer-grade UBER 10^15 (1 error per 12.5 TB read), a full-stripe rebuild read of the surviving disks has a high probability of hitting an uncorrectable read error, which on RAID 5 (single parity) fails the rebuild. This is why RAID 5 is considered unsafe for large-capacity HDDs and RAID 6 / RAID 10 / erasure coding are recommended instead [secondary].
- **Controller-assisted mitigation:** patrol read / background consistency check, configurable rebuild priority, and (on 9600-class cards) much faster rebuild engines (Broadcom claims 60× recovery improvement on 9670-24i) reduce the exposure window but don't change the UBER math [vendor-reported].
- **RAID 50/60 use cases:** large spindle counts where a single RAID 6 set would be too wide (rebuild time) or too narrow (parity overhead); failure domains stay per-subarray [secondary].
- **Hot spares and spare policy:** dedicated vs global hot spares; with SSDs, consider *distributed* spare capacity over dedicated spare drives to keep all NAND wear-leveling evenly [secondary].
- **Write-back vs write-through cache:** write-back (with CacheVault/supercap protection) dramatically improves random-write latency; without cache protection, write-back is unsafe — a power event loses acknowledged writes. Default enterprise practice: write-back only when the cache is protected [secondary].
- **Level comparison table:**

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

