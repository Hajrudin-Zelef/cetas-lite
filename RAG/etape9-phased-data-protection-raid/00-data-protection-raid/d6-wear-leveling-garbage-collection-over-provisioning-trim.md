---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d6-wear-leveling-garbage-collection-over-provisioning-trim
title: "D6 — Wear leveling, garbage collection, over-provisioning, TRIM"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["alignment", "consumer", "cost", "dram", "latency", "memory", "nand", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [173, 244]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: a2a051bcc5b41430c065e5c0c9ec7e14996d4f9fa8eb71168804bab4e51356bb
---

# D6 — Wear leveling, garbage collection, over-provisioning, TRIM

## D6 — Wear leveling, garbage collection, over-provisioning, TRIM

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

- **NVMe SMART (Log 0x02) key fields:** `critical_warning` (bitmask — any non-zero bit is actionable: spare below threshold, temperature out of range, NVM subsystem reliability degraded, read-only mode, volatile backup device failed), `temperature` (composite), `available_spare` + `available_spare_threshold`, `percentage_used` (0 = new, 100 = endurance exhausted), `media_errors` (uncorrectable; should be 0), `data_units_read/written` (×512,000 bytes per unit), `power_cycles`, `unsafe_shutdowns`, error-log entries [secondary].
- **Alert thresholds used in practice:** critical_warning ≠ 0 → critical; media_errors > 0 → investigate; available_spare at/below threshold or < 20% → plan replacement; percentage_used ≥ 90% → endurance nearly exhausted [secondary].
- **ATA/SATA SSD attributes:** ID 5 Reallocated_Sector_Ct, 197 Current_Pending_Sector, 198 Offline_Uncorrectable, 199 CRC_Error_Count (cable/backplane), 177/202/230/231/233 wear-leveling variants (vendor-specific), 194 temperature [secondary].
- **SATA HDD attributes:** the same 5/197/198/199 quartet plus spin-retry and seek-error rates; Backblaze correlation studies link 5/187/188/197/198 to elevated failure probability [independent].
- **NVMe temperature thresholds:** drives expose Warning Composite Temperature Threshold (commonly ~70°C) and Critical (~85°C); exceeding warning sets critical_warning bit 1 and typically engages throttling [secondary].
- **Thermal throttling behavior:** controllers reduce clock/speed or suspend writes to hold temperature; simple implementations use 2–3 stages with visible performance cliffs, while multi-stage designs (e.g. ATP's AceTT: up to 18 stages starting at 85°C) taper gradually to maintain steadier throughput [vendor-reported].
- **Operational impact:** in dense chassis without directed airflow, Gen4/Gen5 NVMe SSDs under sustained load can sit at 70–85°C and throttle; enterprise ratings (e.g. PM9A3: 0–70°C operating) assume data-center airflow, not passive desktop conditions [secondary].
- **Monitoring stack:** `nvme smart-log` / `smartctl -a` for one-shots; node_exporter textfile collectors or smartctl-exporter for fleet telemetry; alert on any critical_warning bit, media_errors > 0, and temperature excursions — not just on FAILED self-assessment [secondary].

---
- **Backblaze's five failure-correlated HDD attributes (independent fleet data):** SMART 5 (Reallocated_Sector_Count), 187 (Reported_Uncorrectable_Errors), 188 (Command_Timeout), 197 (Current_Pending_Sector), 198 (Offline_Uncorrectable). In over 76% of observed failures, at least one of the five had a non-zero raw value before the drive died [independent].
- **Magnitude of correlation:** drives with non-zero SMART 5 show roughly 14× the failure rate of clean drives; non-zero 187/198 roughly 7.5×. Attribute 188 (Command_Timeout) is noisier — low counts are common from power management, but sustained high counts indicate controller/interconnect degradation [independent].
- **Rate matters more than count:** a drive jumping 0→20 on SMART 187 in one day is a far worse omen than a drive accumulating 60 over five years. Alert on *deltas*, not just thresholds [independent].
- **NVMe critical_warning bitmask (byte 0 of SMART log):** bit 0 = available spare below threshold; bit 1 = temperature above threshold or below lower threshold; bit 2 = NVM subsystem reliability degraded; bit 3 = media in read-only mode; bit 4 = volatile memory backup device failed (PLP/capacitor path — direct D1 relevance); bit 5 = persistent memory region read-only. Any set bit is a replace/mitigate-now signal [secondary].
- **SAS HDD deep telemetry:** `smartctl -l background` / SAS log pages expose grown-defect lists, IOEDC (I/O error detection code) counters, and background media scan results — richer than ATA SMART for pinpointing whether errors are media, head, or interconnect [secondary].
- **Monitoring stack (practical):** `smartctl -a` / `nvme smart-log` for spot checks; smartmontools + Prometheus `smartctl_exporter` or node_exporter textfile for fleets; Zabbix templates exist for both. Retention of SMART history enables the delta-based alerting Backblaze's data justifies [secondary].
- **Temperature policy:** NVMe exposes Warning and Critical Composite Temperature Thresholds (commonly ~70 °C / ~85 °C); exceeding warning sets critical_warning bit 1 and engages the controller's thermal policy. Enterprise ratings (PM9A3: 0–70 °C operating) assume data-center airflow — a Gen4/Gen5 drive in a passive desktop slot will throttle [secondary].
- **Throttling design spectrum:** simple controllers use 2–3 coarse stages (visible performance cliffs); ATP's AceTT uses up to 18 stages from 85 °C for gradual taper. Either way, sustained thermal throttling is a capacity-planning signal — add airflow or spread load, don't just accept it [vendor-reported].
- **Spot-check commands:** `smartctl -a /dev/sdX` (SATA/SAS), `smartctl -a /dev/nvme0` (NVMe via smartmontools), `nvme smart-log /dev/nvme0` (native NVMe log page 0x02), `nvme error-log /dev/nvme0` (error entries — any growth is significant) [secondary].
- **What "healthy" looks like:** `critical_warning: 0`, `media_errors: 0`, `num_err_log_entries` stable, `percentage_used` well under 100, `available_spare` at 100%, temperature under WCT, `unsafe_shutdowns` not climbing. Anything else is a ticket, not a shrug [secondary].
- **SSD-specific early warnings beyond SMART:** sudden sustained write-performance drop (GC distress / full drive), rising `percentage_used` faster than host writes imply (high WAF), and thermal throttling under previously-fine loads (heatsink/airflow degradation) [secondary].
- **HDD acoustic/mechanical tells:** SMART won't catch everything — new clicking, spin-up retries (SMART 10), and rising seek-error rates warrant proactive replacement even with SMART 5/197/198 at zero [secondary].
- **Retention policy for SMART history:** keep ≥12 months of per-drive SMART telemetry; Backblaze's rate-of-change finding (0→20 in a day vs gradual) is only actionable with history [independent].
- **Thermal design power reality:** a Gen4 NVMe SSD can draw 8–12 W under sustained load; 24 of them in a 2U shelf = ~250 W of heat in a small volume. Thermal throttling is a *cooling* problem first, a drive problem second [secondary].
- **Heatsinks and directed airflow:** motherboard M.2 slots with no heatsink routinely push client NVMe past 70 °C; enterprise U.2/E1.S carriers assume front-to-back chassis airflow. Match the carrier to the thermal spec, not just the connector [secondary].
- **Temperature vs retention:** high operating temperature accelerates NAND charge leakage (retention) and capacitor aging (PLP) simultaneously — the hot drive is attacked on two axes [secondary].
- **Alerting thresholds (sane defaults):** warning at WCT − 5 °C, critical at WCT; page on critical_warning bit 1 or any excursion past CCT. Log temperature histograms — a slow upward drift predicts cooling failure before throttling starts [secondary].
- **`nvme list-subsys` and topology:** map namespaces → controllers → PCIe addresses before alerting; a "failed drive" alert is useless without the slot/server mapping (combine with SES locate, D10) [secondary].
- **Error log triage:** `nvme error-log` entries with status codes like namespace-not-ready or LBA-out-of-range are often software bugs, not media; media/data-integrity errors are the ones that page [secondary].
- **Dashboarding:** Grafana + Prometheus smartctl_exporter gives per-drive temperature, percentage_used, and media_errors history — the three panels that predict most SSD incidents [secondary].
- **HDD-specific monitors:** load/unload cycle count (SMART 193) on laptop-class drives in 24/7 service; spin-retry (10); and helium-level attributes on He-filled drives where exposed [secondary].
- **Burn-in procedure:** before trusting a new drive, run full-device write + read-verify + SMART review; infant-mortality failures are cheapest when the drive holds no data [secondary].
- **Drive self-tests:** `smartctl -t long` (SATA) / NVMe device self-test log — schedule quarterly on HDDs; a self-test failure is a replace order even if all attributes look fine [secondary].
- **Log persistence across power loss:** SMART counters persist; temperature *history* may not on all drives — export telemetry continuously rather than relying on the drive's memory [secondary].
- **False positives:** a single CRC error (SMART 199) after reseating a cable is noise; a climbing 199 is a cable/backplane fault. Context distinguishes [secondary].
- **Over-provisioning via namespaces:** NVMe allows creating a smaller namespace than the physical capacity, donating the remainder as dynamic OP — a host-side lever when the drive's factory OP is insufficient [secondary].
- **TRIM and RAID passthrough matrix:** mdadm passes discards; ZFS issues them per vdev; hardware RAID usually doesn't — know your layer's behavior before assuming TRIM works [secondary].
- **Reporting hygiene:** when comparing drives, normalize to AFR at stated temperature/workload; "2M hr MTBF" without conditions is marketing, not engineering [secondary].
- **Infant mortality screening:** the first 90 days dominate early failures — burn-in plus close SMART watch in quarter one pays for itself [independent].

