---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d12-jbod-it-mode-for-zfs-and-ceph-why-hba-mode-matters
title: "D12 — JBOD / IT mode for ZFS and Ceph: why \"HBA mode\" matters"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: ["alignment", "cost", "incident", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [407, 459]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 02c88c230773c40b5803d11650a68051a1763612cac0fd942ec882c75f3298c8
---

# D12 — JBOD / IT mode for ZFS and Ceph: why "HBA mode" matters

## D12 — JBOD / IT mode for ZFS and Ceph: why "HBA mode" matters

- **ZFS and Ceph manage redundancy themselves** (ZFS RAIDZ/mirroring + checksums + self-healing; Ceph replication/erasure coding). They require direct, unfiltered access to each physical disk: identity, SMART, error counters, flush/FUA semantics, and write ordering [secondary].
- **Hardware RAID (or RAID-card "JBOD mode") breaks this contract:** virtualized disks hide SMART, mask medium errors until scrub time, reorder or coalesce flushes, and a controller failure takes the whole array offline. TrueNAS community guidance is explicit: use a genuine HBA in IT mode, not RAID mode, not per-disk RAID 0 volumes [secondary].
- **The JBOD-mode trap:** many RAID cards offer a "JBOD" passthrough, but the firmware still abstracts the disks (identity mapping, error handling, cache behavior differ from true HBA firmware). Community testing and forum guidance treat this as insufficient for ZFS production [secondary].
- **Ceph:** OSDs similarly want raw block devices; BlueStore issues its own checksums and needs direct device control. Deploy Ceph on HBA/JBOD or NVMe-direct nodes, not behind hardware RAID [secondary].
- **When hardware RAID is still right:** boot volumes, small VM datastores needing simple mirroring, Windows-only shops without ZFS expertise, or workloads needing the controller's write-back cache — provided cache is protected (CacheVault) [secondary].
- **SED note for both worlds:** hardware RAID cards (Broadcom SafeStore, SmartRAID 4300 SED support) and OS tools (ZFS native encryption, LUKS) solve different problems — drive-level theft protection vs dataset-level encryption. They compose; pick per threat model [secondary].
- **Known-good HBA silicon for IT mode:** Broadcom/LSI SAS3008 (9300-8i/16i, 12G), SAS3408/SAS3416 (9400), SAS3808 (9500), SAS4124-based 9600 in eHBA mode; all present each disk as a direct `/dev/sdX` or NVMe namespace with full SMART/log-page passthrough [secondary].
- **Crossflashing reality:** many "HBA" cards on the secondary market are RAID cards flashed to IT firmware (or need flashing). A proper IT flash behaves like a factory HBA; a RAID card left in "JBOD mode" does not — verify with `sas3flash -list` / `storcli` and by checking SMART passthrough, not by the seller's label [secondary].
- **ZFS disk-access requirements (why IT mode is non-negotiable):** ZFS needs stable device identity (use /dev/disk/by-id, not sdX), direct SMART, honest FLUSH/FUA, and visibility into medium errors for self-healing. Hardware RAID virtualizes all four away [secondary].
- **ashift and alignment:** set `ashift=12` (4K sectors) or `13` (8K) at pool creation for modern drives; wrong ashift is a permanent performance tax [secondary].
- **ECC RAM:** recommended for ZFS (a bit flip in ARC can theoretically corrupt data before checksumming) but not required — the "ZFS needs ECC" claim is overstated; what ZFS needs is the HBA passthrough above [secondary].
- **Ceph BlueStore:** one OSD per physical drive (or per NVMe namespace), no RAID underneath; BlueStore keeps its own checksums and RocksDB metadata. RAID-under-Ceph double-parities capacity and blinds placement groups to real failure domains [secondary].
- **Proxmox VE note:** Proxmox's ZFS-on-root and ZFS storage backends assume direct disks; passing RAID virtual disks to Proxmox ZFS inherits all the D12 problems [secondary].
- **The "one disk per RAID 0" anti-pattern:** creating N single-disk RAID 0 volumes to fake JBOD is the worst of both worlds — you keep the controller as a single point of failure *and* lose SMART passthrough. TrueNAS forums document real data-loss incidents from exactly this setup [secondary].
- **HBA firmware updates:** update HBA firmware with the same discipline as RAID cards (config backup, maintenance window) — an HBA flash gone wrong takes every disk offline at once [secondary].
- **SAS address (WWN) stability:** record HBA SAS addresses; after replacement, zoning and multipath bindings may need updating — another line for the runbook [secondary].
- **NVMe-direct for ZFS/Ceph:** CPU-attached NVMe needs no HBA at all — the "HBA question" only applies to SAS/SATA shelves. Don't buy an HBA for an all-NVMe node [secondary].
- **TrueNAS hardware guidance (community):** genuine HBA in IT mode, no RAID abstraction, ECC recommended, and enough RAM for ARC (1 GB per TB of storage is the old rule of thumb; L2ARC needs even more). The forums are blunt because the failure mode (silent corruption) is the worst kind [secondary].
- **Proxmox Backup Server:** PBS datastores on ZFS with sync writes — put the datastore on PLP-SSD or a SLOG, or backup throughput collapses [secondary].
- **Ceph drive classes:** use CRUSH device classes (hdd/ssd/nvme) so pools land on the right media; mixing classes in one pool gives you the worst of both [secondary].
- **HBA queue depth and ZFS:** ZFS issues deep concurrent I/O; an HBA with shallow queues becomes the bottleneck on all-flash shelves — same check as D10, now for the software-defined stack [secondary].
- **Multipath with ZFS:** ZFS prefers its own multipath handling; disable OS dm-multipath on ZFS members to avoid double-path confusion — document the choice [secondary].
- **RAID 6 dual-failure math:** RAID 6 survives any 2 failures; the 3rd failure during rebuild is the data-loss event — MTTDL models this as AFR² × rebuild-window terms, which is why rebuild speed (D8's 60× claim) matters [secondary].
- **Rebuild vs re-silver terminology:** hardware RAID "rebuild," ZFS "resilver," Ceph "backfill/recovery" — same concept (restore redundancy), different scope (whole disk vs allocated blocks vs placement groups) [secondary].
- **Predictive spare activation:** some controllers pre-copy to a hot spare on SMART predictive-failure alerts, shrinking the degraded window to near zero — enable where supported [secondary].
- **Controller event log retention:** export and archive controller event logs with backups — post-incident forensics without the controller log is guesswork [secondary].
- **CacheVault lifecycle:** supercaps age; monitor CacheVault health in MSM/StorCLI and replace on warning — a failed supercap silently forces write-through (safe but slow) or blocks write-back [secondary].

## D13 — Decision matrices: hardware RAID vs ZFS vs Ceph erasure coding

- **Matrix 1 — redundancy technology by workload:**

| Workload | Hardware RAID | ZFS (mirrors/RAIDZ) | Ceph (repl/EC) |
|---|---|---|---|
| Boot / OS mirror | Best fit — simple, BIOS-visible | OK (mirrors) | Overkill |
| Single-node VM datastore | Good (RAID 10 + protected WB cache) | Best fit (mirrors, ARC, snapshots) | Not applicable |
| Single-node bulk storage (HDD) | OK (RAID 6) | Best fit (RAIDZ2, self-healing) | Not applicable |
| Multi-node scale-out | Not applicable | Via replication tools only | Best fit (EC for capacity) |
| 8–32 NVMe high-throughput | SupremeRAID / SmartRAID 4300 class | Good (mirrors; RAIDZ parity cost) | EC on NVMe (CPU-heavy) |
| Archival / backup target | RAID 6 acceptable | RAIDZ2/3 good | EC best (capacity efficiency) |
- **Matrix 2 — failure-domain and rebuild behavior:**

| Property | HW RAID 6 | ZFS RAIDZ2 | Ceph EC (k+m) |
|---|---|---|---|
| Tolerated disk failures | 2 per array | 2 per vdev | m per placement group |
| Silent corruption detection | No (unless DIF/DIX end-to-end) | Yes (checksums, self-heal) | Yes (BlueStore checksums) |
| Rebuild scope | Whole disks | Only allocated blocks | Only affected PGs |
| Controller as SPOF | Yes (mitigate: dual-path SAS) | No (HBA passthrough) | No (CRUSH, no single controller) |
| Write penalty (random) | 6 (RAID 6) | ~parity overhead, ARC-assisted | Network + EC encode |
- **Matrix 3 — cost/complexity/skill:**

