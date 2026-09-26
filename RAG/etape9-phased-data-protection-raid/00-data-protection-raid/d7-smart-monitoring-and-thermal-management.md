---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d7-smart-monitoring-and-thermal-management
title: "D7 — SMART monitoring and thermal management"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["alignment", "consumer", "cost", "dram", "latency", "memory", "nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [175, 204]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 2c5c8e9b5a105a6121a157521b6e41ce02ab230e972f4c67c9377fc91f733a39
---

# D7 — SMART monitoring and thermal management

- **Wear leveling:** the FTL spreads program/erase cycles across all NAND blocks because each block tolerates a finite number of P/E cycles (fewer per cell as bits-per-cell rise: SLC > MLC > TLC > QLC). Static wear leveling also relocates cold data so no block sits under-cycled [secondary].
- **Garbage collection:** NAND cannot overwrite in place — pages are written once, then whole blocks must be erased. GC reclaims blocks containing invalidated pages by copying still-valid pages elsewhere, which creates **write amplification** (more NAND writes than host writes) [secondary].
- **Write amplification factor (WAF):** the ratio of NAND writes to host writes; high WAF burns endurance and depresses sustained write performance. Random-write workloads and nearly-full drives raise WAF; over-provisioning and TRIM lower it [secondary].
- **Over-provisioning (OP):** spare NAND invisible to the host, used for GC, wear leveling, and bad-block replacement. Enterprise drives ship with higher factory OP (often ~28% on write-intensive SKUs); the NVMe `available_spare` SMART field reports remaining spare pool [secondary].
- **TRIM / Deallocate:** the OS tells the SSD which LBAs are free so GC doesn't waste cycles preserving deleted data. Without TRIM (or with filesystems that don't issue it), write performance and endurance degrade over time. `blkdiscard --secure` semantics depend on deterministic-TRIM + read-zero-after-trim support [secondary].
- **Thermal coupling:** sustained GC under write load is a major heat source; see D7 — thermal throttling can be triggered by GC storms on nearly-full drives, not just host I/O [secondary].
- **Monitoring:** track `percentage_used` (NVMe) / wear-leveling attributes (ATA 177/230/231/233, vendor-specific), `available_spare`, and host-vs-NAND write counters where exposed; a climbing WAF with flat host writes signals GC distress [secondary].
- **Dynamic vs static wear leveling:** dynamic WL only spreads writes across free blocks (cold data never moves — a weakness); static WL periodically relocates cold data so all blocks cycle evenly. Enterprise firmware implements static WL; cheap controllers may not [secondary].
- **Read disturb and read reclaim:** repeated reads of one page can flip bits in neighboring pages; firmware counters trigger read-reclaim (rewrite the block) before errors accumulate — another background NAND write source alongside GC [secondary].
- **Data retention (powered off):** JEDEC specifies client SSDs retain data ~1 year and enterprise ~3 months at end-of-life *at rated temperature* — powered-off archival on SSDs is not a strategy; retention is worse at high temperature. Tape/optical own the cold-archive role (D14/D15) [secondary].
- **TRIM flavors:** queued TRIM (SATA) / Deallocate (NVMe); deterministic vs non-deterministic read-after-trim. Filesystems must be configured to issue discards (e.g. `discard` mount option or periodic `fstrim`) — it is not always on by default [secondary].
- **OP sizing:** ~7% factory OP is typical for client drives; write-intensive enterprise SKUs ship ~28%+. Some enterprise drives allow flexible OP (sacrificing user capacity for endurance) — a tuning lever for write-heavy workloads [secondary].
- **SLC cache write cliff (consumer QLC):** drives advertise burst speeds from a pseudo-SLC cache; sustained writes past cache exhaustion fall to native QLC speed (often <200 MB/s). Enterprise TLC with PLP is rated at steady-state, not burst [secondary].
- **HMB (Host Memory Buffer):** DRAM-less NVMe SSDs borrow host RAM for FTL tables — fine for clients, absent from enterprise drives (which carry their own DRAM + PLP) [secondary].
- **WAF estimation in practice:** WAF ≈ NAND writes ÷ host writes; some enterprise drives expose both counters. Sustained WAF > 3 on a random-write workload signals the drive is nearly full or TRIM-starved — add OP or free space before endurance burns [secondary].
- **Filesystem choices affect WAF:** CoW filesystems (ZFS, btrfs) amplify writes vs ext4/XFS for some workloads; databases with small random writes are the classic high-WAF case — match drive DWPD to measured WAF, not to host write volume alone [secondary].
- **Partition alignment:** misaligned partitions (legacy 63-sector offset) turn every 4K write into a read-modify-write, doubling WAF. Modern installers align to 1 MB; verify on migrated images [secondary].
- **Discard granularity:** the NVMe deallocate granularity (reported in `nvme id-ns`) determines how efficiently TRIM works; coarse granularity wastes GC on partially-trimmed blocks [secondary].
- **P/E cycle budgets by NAND type (typical, varies by vendor/process):** SLC ~100K, MLC ~10K, TLC ~3K, QLC ~1K. This 100:1 SLC-to-QLC gap is why QLC is read/cold-storage NAND and TLC remains the enterprise workhorse [secondary].
- **DWPD classes:** read-intensive (~1 DWPD, e.g. PM9A3), mixed-use (~3 DWPD), write-intensive (10+ DWPD, often higher OP + SLC-cache designs). Match the class to measured workload, not to price [secondary].
- **Zoned Namespaces (ZNS):** NVMe ZNS lets the host control placement, eliminating GC/WAF inside the drive for ZNS-aware software (e.g. some Ceph/RocksDB work) — the architectural answer to write amplification, at the cost of software complexity [secondary].
- **Flexible data placement (FDP):** the incremental step before ZNS — host hints guide the FTL's GC without full zone management [secondary].
- **GC under RAID:** the RAID layer can't see GC pressure; a "healthy" array of nearly-full SSDs can still hit synchronized GC stalls (all members pausing together). Keep SSD arrays under ~80% full for performance-critical roles [secondary].
- **Deterministic latency and GC:** predictable-latency-mode drives bound GC pauses — pair with RAID levels whose penalties you can model, unlike best-effort consumer drives [secondary].
- **PLP testing across temperature:** capacitor hold-up degrades at temperature extremes — run power-fail acceptance at the chassis's hottest operating point, not on a lab bench [secondary].
- **Drive power states:** NVMe APST power states trade latency for power; aggressive power management on enterprise drives can interact with latency SLAs — validate, don't default [secondary].
- **Firmware-signed RMA:** replacement drives may ship older firmware — re-flash to the fleet standard before adding to an array [secondary].

## D7 — SMART monitoring and thermal management

