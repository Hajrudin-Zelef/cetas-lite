---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-11
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 11)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: ["2025-08-05"]
keywords: ["accelerator", "benchmark", "compute", "disaggregated", "latency", "nand", "pricing"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [247, 274]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 383d133366c75444fbe402b378e36def8b367cda52b73c841448be5e5473f3be
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 11)

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
