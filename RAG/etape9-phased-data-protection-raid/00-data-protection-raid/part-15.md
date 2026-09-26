---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-15
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 15)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Nvidia"]
dates: []
keywords: ["accelerator", "alignment", "benchmark", "cost", "gpu", "latency", "license", "licenses", "nvidia", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [369, 406]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 0c0313d84560b3a45697766502098e39c38a90df84d136424da8c1ae74a32cd3
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 15)

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

