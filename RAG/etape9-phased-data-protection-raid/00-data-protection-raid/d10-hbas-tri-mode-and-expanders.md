---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d10-hbas-tri-mode-and-expanders
title: "D10 — HBAs, tri-mode, and expanders"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom", "Nvidia"]
dates: []
keywords: ["accelerator", "alignment", "benchmark", "cost", "gpu", "latency", "license", "licenses", "nvidia", "serdes", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [334, 406]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 2008bc812df40d33cb54fe5adccb525f01f187dfd50c5314b537b4089e79f6d1
---

# D10 — HBAs, tri-mode, and expanders

## D10 — HBAs, tri-mode, and expanders

- **HBA vs RAID card:** an HBA (Host Bus Adapter) presents disks directly to the OS with no RAID logic; an eHBA (enhanced HBA, Broadcom 9600 terminology) adds enterprise management features (secure boot, SPDM attestation, SED key management, UBM) while still passing disks through [secondary].
- **Broadcom HBA 9600 series:** PCIe Gen4 x8, 24G SAS tri-mode; 9600-16i (05-50111-00, internal, 16 ports, 15 W, MTBF 5M hours per reseller spec sheet); 9600-24i (05-50111-01, 24 ports); 9600W-16e (external, SFF-8674); up to 240 SAS/SATA or 32 NVMe devices per controller [secondary].
- **Broadcom HBA 9500 series:** PCIe Gen4 tri-mode, 12G SAS generation (e.g. 9500-16e 05-50075-00 external, 9500-8i 05-50077-03); connects up to 1024 SAS/SATA or 32 NVMe [secondary].
- **Tri-mode SerDes:** one adapter concurrently serves NVMe, SAS, and SATA devices in the same backplane, auto-negotiating protocol and speed (SAS 22.5/12/6 Gb/s lanes, SATA 6 Gb/s, PCIe up to 16 GT/s per lane); enables non-disruptive migration from SAS/SATA to NVMe [secondary].
- **SAS expanders:** fan out one HBA port to dozens of drives; Broadcom SAS4x24 expander (24G SAS-4) doubles per-lane bandwidth vs 12G; AIC launched a SAS4 JBOD family (4U 108-bay, 78-bay, 2U 24-bay) on Broadcom SAS4x expanders at SC24 (Nov 2024) with dual-hub expander modules for dual-path redundancy [secondary].
- **Enclosure services (SES/SES-2):** expanders and backplanes report drive presence, temperature, voltage, and locator/fault LEDs; UBM (Universal Bay Management, SFF-TA-1005) standardizes tri-mode bay management [secondary].
- **Cabling notes:** SFF-8654 (SlimSAS) internal and SFF-8674 external on 24G parts; existing SFF-8639 (U.2) backplanes are compatible; DataBolt2 lets 24G controllers use 12/6G drives and backplanes at aggregated bandwidth [secondary].
- **Crossflashing caution:** many HBAs ship as RAID cards flashed to IT/HBA firmware (or vice versa). For ZFS/Ceph, use genuine HBA firmware or properly crossflashed IT-mode firmware; a RAID card in "JBOD mode" is not equivalent (see D12) [secondary].
- **9500 vs 9600 at a glance:** 9500 = PCIe Gen4, 12G SAS tri-mode (mature, cheaper, plenty for HDD/SATA-SSD estates); 9600 = PCIe Gen4, 24G SAS tri-mode, eHBA security features (secure boot, SPDM attestation), higher NVMe device counts. Buy 9600 for new NVMe-heavy builds; 9500 remains rational for bulk HDD [secondary].
- **Port/cable map:** internal SFF-8654 (SlimSAS x4/x8), external SFF-8674; U.2 (SFF-8639) backplanes compatible via tri-mode. One x8 SlimSAS cable carries 8× 24G SAS lanes ≈ 192 Gb/s raw to an expander/backplane [secondary].
- **Expander topology rules:** SAS4x24-class expanders fan one HBA port to dozens of bays; cascade expanders enclosure-to-enclosure (AIC's SAS4 JBODs support this); keep cascade depth shallow (1–2) to bound latency and failure blast radius [secondary].
- **Dual-domain multipath:** dual expander modules + dual-port SAS drives give two independent paths; OS multipath (Linux dm-multipath, Windows MPIO) survives a cable/expander/HBA-port failure. SATA drives are single-port — no true dual-path for SATA [secondary].
- **Zoning:** expander zoning carves one JBOD into isolated groups for multiple hosts — useful for shared-JBOD clusters without a full SAN [secondary].
- **SES/SES-2 and UBM:** enclosure services report per-slot presence, fault/locate LEDs, temperature, and voltage; UBM (SFF-TA-1005) unifies this across NVMe/SAS/SATA tri-mode bays so one backplane design serves all three [secondary].
- **Power/thermal:** 9600-16i eHBA is rated ~15 W (reseller spec) — plan PCIe slot power and airflow; 24-port and RAID cards with cache run hotter [secondary].
- **Bandwidth math for planning:** 24G SAS = 24 Gb/s ≈ 2,400 MB/s per lane after encoding; a x4 SFF-8654 link to an expander ≈ 9,600 MB/s raw. 12 HDDs at 250 MB/s = 3,000 MB/s — one x4 24G link feeds them with headroom; 24 SATA SSDs at 550 MB/s = 13,200 MB/s — needs dual x4 links or fewer drives per expander uplink [secondary].
- **Oversubscription rule:** expander uplinks are routinely oversubscribed for HDD (disks rarely stream simultaneously); for all-flash NVMe/SAS-SSD shelves, provision uplinks near 1:1 with aggregate drive bandwidth or accept contention [secondary].
- **SATA tunneling caveat:** SAS expanders tunnel SATA via STP; SATA drives don't support dual-porting, so a "dual-domain" shelf with SATA drives still has single-path drives — dual-path requires SAS drives [secondary].
- **Interop checklist before buying:** HBA firmware ↔ expander firmware ↔ backplane (SES/UBM) ↔ drive firmware. Tri-mode multiplies the matrix — get the vendor's tested-configuration list in writing for NVMe+SATA mixed shelves [secondary].
- **HBA queue depth:** enterprise HBAs expose deep queues (1K+ per device); a shallow-queue HBA starves NVMe SSDs capable of 64K+ queue depth — check queue-depth specs when the workload is high-IOPS NVMe [secondary].
- **Boot support:** HBAs (and eHBAs) support boot from attached devices via UEFI driver; RAID cards additionally offer bootable virtual disks. For ZFS-root or mdadm-root, HBA boot is sufficient [secondary].
- **Driver maturity:** `mpt3sas` (12G) and `mpi3mr`/`mpt3sas`-successor drivers for 24G parts are in-tree on modern Linux; verify the exact kernel's driver supports the card's PCI ID before deploying a new distro release [secondary].
- **LED/locator integration:** SES-managed fault/locate LEDs are the physical-safety layer for drive pulls — confirm the HBA↔backplane SES path works (test a locate blink) before you need it at 2 a.m. [secondary].
- **Cable length and signal integrity:** 24G SAS passive cables are length-limited (typically ≤1 m internal, check vendor spec); longer runs need active cables or retimers. CRC errors (SATA SMART 199 / SAS PHY error counters) that rise with cable length are a cabling problem, not a drive problem [secondary].
- **PHY error counters:** `smartctl -l sasphy` exposes per-PHY invalid DWORDs, disparity errors, and link resets — the diagnostic layer below SMART for flaky links [secondary].
- **Expander firmware:** expanders run firmware too (SES processor) — include it in the firmware inventory (D3); a hung expander looks like many simultaneous drive failures [secondary].
- **External SAS for JBOD shelves:** 9600-16e/9500-16e external ports (SFF-8674) connect rack JBOD shelves; keep external cable runs short and latching — a bumped external SAS cable is a classic multi-disk "failure" [secondary].
- **Enclosure numbering:** record enclosure:slot topology (`sas3ircu`/`sas3flash` topology pages); after expander replacement, re-verify the map — SES addresses can shift [secondary].
- **RAID 0 stripe width:** wider stripes help sequential throughput; RAID 0's risk scales with member count — 8-disk RAID 0 has ~8× the failure exposure of one disk [secondary].
- **Nested RAID naming:** RAID 10 = stripe of mirrors (1+0); RAID 01 = mirror of stripes (obsolete, weaker) — the order matters [secondary].

## D11 — NVMe hardware RAID: GRAID SupremeRAID and the software alternative

- **GRAID SupremeRAID SR-1010:** PCIe Gen4 x16 card built on an NVIDIA GPU (Ampere GA106 / RTX A2000 6 GB in current bundles; predecessor SR-1000 used the T1000); RAID 0/1/5/6/10 over up to 32 native NVMe SSDs; out-of-path design — data flows CPU→SSD directly, the GPU handles parity/RAID math without consuming CPU cycles or throttling throughput [vendor-reported].
- **Vendor performance claims — conflict flagged:** GRAID's newer brochure claims up to 28M 4K random-read IOPS and 260 GB/s sequential read; launch-era coverage (StorageReview, TweakTown) reported 19M IOPS and 110 GB/s reads, 1.5M IOPS / 22 GB/s writes. The discrepancy likely reflects different test configs (Gen5 platforms, drive counts) but neither figure is independently verified at cutoff [vendor-reported] [conflict].
- **Also claimed:** NVMe-oF initiator/target support, compression, encryption, thin provisioning, boot-drive protection; no battery/supercap needed because there is no write-back cache on the card (no cache = nothing to lose on power failure) [vendor-reported].
- **Linux mdadm RAID:** kernel software RAID for NVMe (and SATA/SAS); RAID 0/1/5/6/10; mature, free, no vendor lock-in; CPU cost for parity on modern CPUs is modest, but rebuilds are single-threaded-ish and slow relative to GPU-offloaded designs; widely used for boot/OS mirrors [secondary].
- **When SupremeRAID-class products make sense:** 8–32 NVMe SSDs needing RAID 5/6 at near-native throughput for AI checkpointing, video, HPC scratch — workloads where mdadm parity would cap throughput or burn excessive CPU [secondary].
- **When they don't:** small arrays (mdadm or ZFS mirrors are simpler and cheaper), capacity-oriented HDD pools (HBA + ZFS/Ceph), or any environment where GPU driver/firmware adds operational risk [secondary].
- **Vendor lock-in note:** SupremeRAID arrays depend on GRAID's software stack and license; mdadm/ZFS arrays are portable across any Linux host [secondary].
---
- **SupremeRAID generations:** SR-1000 (launch, NVIDIA T1000 GPU, PCIe Gen4, up to 32 drives); SR-1010 (current, Ampere GA106 / RTX A2000 6 GB, same 32-drive, RAID 0/1/5/6/10 envelope, NVMe-oF initiator/target, compression, encryption, thin provisioning). The performance-claim conflict in D17 spans these generations and test configs [vendor-reported].
- **Why out-of-path wins at NVMe scale:** 32× Gen4 x4 SSDs ≈ 224 GB/s of raw read bandwidth; no x16 RAID-on-Chip can inline that. SupremeRAID keeps data on CPU-attached PCIe lanes and uses the GPU only for parity math — the card never becomes the bottleneck the way an inline ROC does [secondary].
- **No cache battery by design:** SupremeRAID has no write-back cache on the card, so there is nothing to protect across power loss — one BOM item and failure mode fewer than classic RAID cards [vendor-reported].
- **mdadm (Linux MD) profile:** RAID 0/1/4/5/6/10 in-kernel, free, portable across any Linux host; write-intent bitmaps bound resync work after unclean shutdown; reshape (grow/shrink) supported. Parity RAID 5/6 throughput is CPU-bound — fine at 4–8 SSDs, increasingly wasteful at 16+ [secondary].
- **ZFS on NVMe:** mirrors recommended over RAIDZ for performance; special allocation vdevs (metadata/small blocks on NVMe) accelerate HDD pools; ZIL/SLOG on PLP-NVMe for sync-write latency. RAIDZ expansion (adding disks to RAIDZ) matured in OpenZFS 2.x — check version before planning [secondary].
- **btrfs warning:** btrfs RAID 5/6 still carries the write-hole risk and is not recommended for production parity RAID — use mirrors or RAID 1C3/1C4 profiles [secondary].
- **Windows:** Storage Spaces (mirror/parity) is the native option; for performance-critical NVMe RAID on Windows, a hardware/offload card is usually the pragmatic choice [secondary].
- **Lock-in ledger:** SupremeRAID arrays need GRAID's driver/license stack; mdadm/ZFS arrays move between hosts freely. Price the exit, not just the entry [secondary].
- **NVMe RAID option comparison:**

| Option | Max drives (typical) | RAID levels | Parity bottleneck | License/lock-in | Best for |
|---|---|---|---|---|---|
| GRAID SupremeRAID SR-1010 | 32 | 0/1/5/6/10 | GPU offload (vendor claim) | Commercial, GRAID stack | 8–32 NVMe at line rate |
| Microchip SmartRAID 4300 | 32 | 0/1/10/5/50 | Accelerator offload (vendor claim) | Commercial, Microchip stack | NVMe RAID 5/50 on Gen4/5 |
| Linux mdadm | Platform-limited | 0/1/4/5/6/10 | Host CPU | None (in-kernel) | ≤8 NVMe, boot mirrors |
| ZFS mirrors | Platform-limited | mirror/RAIDZ | Host CPU | None (OpenZFS) | Integrity-first single node |
- **Throughput sizing sanity check:** one Gen4 x4 NVMe ≈ 7 GB/s read; 32 of them ≈ 224 GB/s — verify the *platform* (PCIe lane budget, NUMA) can actually feed that before buying any RAID product; most real deployments are lane-limited, not RAID-limited [secondary].
- **GPU driver as operational surface:** SupremeRAID depends on the NVIDIA driver stack in the storage host — kernel upgrades that break the NVIDIA driver break the array. Pin and test driver+kernel combos; this is the hidden opEx of GPU-offloaded RAID [secondary].
- **Power and thermals:** an A2000-class GPU adds ~70 W and its own thermal/fan requirements to the storage node — include it in power and cooling budgets [secondary].
- **mdadm reshape caution:** growing/reshaping mdadm arrays rewrites the whole array; keep backups and never reshape the only copy. Hardware RAID expansion has the same warning [secondary].
- **ZFS RAIDZ width guidance:** 2^n + parity disks (e.g. 6+2, 10+2) for alignment efficiency; very wide RAIDZ (12+3) slows rebuilds and random I/O — prefer multiple narrower vdevs striped [secondary].
- **SupremeRAID data services (vendor-reported):** thin provisioning, compression, encryption, snapshots/clones, and NVMe-oF target/initiator — positioning as a storage-services card, not just parity math [vendor-reported].
- **Boot support:** SupremeRAID supports protecting the OS boot volume — relevant because mdadm/ZFS-boot has its own complexity on some platforms [vendor-reported].
- **When mdadm wins:** existing Linux expertise, small arrays, no budget for licenses, need for reshape/grow flexibility, and environments where a GPU in the storage node is a non-starter [secondary].
- **Benchmark before buying:** any NVMe RAID decision above 8 drives deserves a POC with your workload (`fio` profiles matching production) — vendor IOPS are measured with queue depths your application will never produce [secondary].
- **Multi-card scaling:** SupremeRAID supports multiple cards per host for >32 drives — verify PCIe slot/bifurcation and NUMA placement; crossing NUMA nodes costs latency [vendor-reported].
- **ZFS dRAID:** distributed RAIDZ variant that rebuilds faster via distributed spares — the ZFS answer to wide-parity rebuild windows, worth evaluating against hardware RAID 60 for large HDD pools [secondary].
- **Expander fan-out limits:** a single SAS4x24 expander fans to 24+ devices; cascade for more, but each hop adds latency and a failure domain — measure before cascading deep [secondary].
- **SES heartbeat:** poll enclosure status periodically; a silent expander (stale SES data) masks drive failures — the monitor must watch the watcher [secondary].

