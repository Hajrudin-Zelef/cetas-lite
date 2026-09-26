---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-2
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 2)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["benchmark", "consumer", "cost", "dram", "latency", "memory", "nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [12, 38]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: bd297a59dfe76f68f758949e8020c2685b50c03f169eaa6d7cfe5cc6a669e22c
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 2)

- **DLP (Data Loss Protection)** in this file refers to drive- and controller-level mechanisms that prevent in-flight or cached data from being lost or corrupted during unexpected power events, not DLP software for exfiltration prevention (out of scope) [secondary].
- **The core problem:** SSDs buffer writes in volatile DRAM (mapping tables/FTL plus write cache). A sudden power cut erases DRAM contents instantly, which can destroy both user data in flight and the flash translation layer metadata needed to locate any data on the drive [secondary].
- **Enterprise SSDs solve this with PLP (Power Loss Protection):** onboard capacitors (supercapacitors or tantalum polymer capacitors) provide a few milliseconds of hold-up power after voltage drop is detected, giving the controller time to flush DRAM contents and updated FTL mapping tables to NAND before shutdown [official].
- **Kingston's documented implementation:** hardware (supercaps or tantalum polymer capacitors for short-term hold-up) plus firmware routines; the controller flushes in-flight data from cache and updates the Flash Translation Layer; if updates were interrupted, firmware rebuilds structures on reboot; qualification requires surviving numerous unsafe power-loss events and booting fully functional, or the whole qualification restarts [official].
- **ATP Electronics' four-layer industrial model:** (1) an MCU watches the input rail and de-glitches it so ripples/dips don't trigger false flushes and over-voltage cuts input; (2) polymer tantalum capacitors supply reserve power to flush DRAM cache to NAND and complete the last program operation — the only layer that protects in-flight data; (3) Sudden Power-Off Recovery (SPOR) firmware restores firmware and mapping tables on next boot (cannot save in-flight data); (4) capacitor health monitoring, because hold-up capacity degrades and usually fails silently [secondary].
- **Consumer vs enterprise:** PLP capacitors are essentially standard on enterprise/data-center SSDs and rare on consumer SSDs; the Samsung PM9A3 lists "Enhanced Power Loss Data Protection" (capacitor-backed) as a defining enterprise feature while consumer drives omit it [secondary].
- **Controller-level analog:** the same in-flight-write problem exists one layer up at the RAID controller's write-back cache — solved by CacheVault-style supercap + NAND flash offload (see D8), not by UPS alone [secondary].
- **A UPS is not a substitute:** site UPS covers AC mains failure but not PSU failure, cable dislodging, firmware lockups, or host crashes mid-write; PLP is the last-line guarantee inside the drive itself [secondary].
- **Volatile write cache size:** enterprise NVMe SSDs typically carry ~1 GB DRAM per 1 TB NAND (e.g. ~4 GB on a 3.84 TB drive) holding both user write cache and the FTL mapping table; losing it means losing not just seconds of writes but the map needed to find *any* data [secondary].
- **FTL mapping table criticality:** the logical-to-physical map is the drive's most valuable metadata — Kingston's firmware explicitly prioritizes flushing updated FTL structures during the PLP window, and rebuilds them on reboot if interrupted [official].
- **Hold-up time budget:** capacitor hold-up is measured in tens of milliseconds — enough to program cached pages and metadata sequentially to NAND, not enough for random host I/O to continue; firmware therefore switches to an emergency flush routine the instant the rail monitor trips [secondary].
- **Supercapacitors vs tantalum polymer:** supercaps offer high capacitance in small volumes but narrower temperature ratings; tantalum polymer caps are smaller per joule but wider-temp and longer-lived. Kingston documents both being used across its enterprise line depending on form factor and thermal envelope [official].
- **Capacitor aging is the silent killer:** hold-up capacitance degrades with temperature and age, and failure is usually invisible to the host — which is exactly why ATP's fourth PLP layer is continuous capacitor health monitoring rather than assuming the caps still work [secondary].
- **Unsafe shutdown audit trail:** NVMe SMART exposes `unsafe_shutdowns` (count of shutdowns without STANDBY IMMEDIATE / controller shutdown notification). A climbing counter with no site power events points at PSU, cable, or backplane issues — investigate before data loss, not after [secondary].
- **Flush semantics the host relies on:** ATA FLUSH CACHE, SCSI SYNCHRONIZE CACHE, NVMe flush command, and FUA (Force Unit Access) bits are the contracts by which filesystems (ext4 journal, XFS) and databases (WAL) order writes. If a drive lies about flush completion (volatile cache enabled, no PLP), an OS crash can corrupt the filesystem above it [secondary].
- **Consumer-drive tradeoff:** many consumer SSDs ship with volatile write cache enabled for benchmark numbers and no PLP; disabling it (`hdparm -W0`, or NVMe volatile-write-cache feature) restores ordering safety at a steep write-performance cost. Enterprise drives keep the cache fast *and* safe via PLP [secondary].
- **The RAID-card mirror of the same problem:** a RAID controller's DRAM write-back cache has identical volatility exposure, solved by the same pattern one layer up — supercapacitor + NAND offload (CacheVault), covered in D8 [secondary].
- **NVDIMM/persistent memory alternative:** NVDIMM-N (DRAM + NAND + supercap on the DIMM) gives software a byte-addressable persistent write buffer with DRAM latency — used for ZFS SLOG/ZIL and database journals where even PLP-SSD latency is too high [secondary].
- **UPS ≠ PLP:** a rack UPS covers AC mains failure only. PSU failure, power-cable dislodge, backplane faults, firmware hangs, and host crashes all bypass the UPS — PLP inside the drive is the last line, and it must be validated, not assumed [secondary].
- **NVMe shutdown notification types:** normal shutdown notification lets the controller flush gracefully (fast, no PLP needed); sudden power loss is the PLP path. A host that hard-resets without notification converts every reboot into an unsafe shutdown — fix the host, not the drive [secondary].
- **STANDBY IMMEDIATE and SYNCHRONIZE CACHE:** the commands hosts issue before power-down; drives without PLP depend on the OS actually sending them. Crash/power-cut scenarios never send them — the entire PLP justification [secondary].
- **PLP and the FTL commit point:** enterprise firmware checkpoints FTL metadata transactionally so a torn write during the PLP window is detected and rolled back on next boot (SPOR) rather than silently accepted [secondary].
- **Capacitor redundancy:** some enterprise drives split hold-up across multiple capacitor banks so a single shorted cap doesn't kill PLP — ask vendors about single-point-of-failure in the hold-up circuit for critical deployments [unverified].
- **PLP validation in qualification:** Kingston's documented practice — repeated unsafe power-loss cycles with full functional boot checks, restarting qualification on any brick — is the industry pattern; request the vendor's power-loss test report for drives going into write-heavy roles [official].
- **Interaction with RAID write-back cache:** drive-level PLP protects the SSD's internal cache; it does not protect data sitting in the *controller's* DRAM — both layers (PLP + CacheVault) are needed for an end-to-end power-safe write path [secondary].
- **Virtualized environments:** a hypervisor host crash is a "power loss" from the guest's perspective — guest filesystems depend on the virtual disk honoring flush, which depends on the physical drive's PLP. The chain is only as strong as the physical layer [secondary].
- **Testing PLP yourself:** managed-PDU power cuts during `fio` write load, followed by filesystem check and data-integrity verification (checksummed dataset), is the acceptance test. Do it on a sample of every new drive model before fleet rollout [secondary].
