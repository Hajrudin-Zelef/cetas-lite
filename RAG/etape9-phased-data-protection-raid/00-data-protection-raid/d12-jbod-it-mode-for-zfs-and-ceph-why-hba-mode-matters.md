---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d12-jbod-it-mode-for-zfs-and-ceph-why-hba-mode-matters
title: "D12 — JBOD / IT mode for ZFS and Ceph: why \"HBA mode\" matters"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: ["alignment", "capex", "cost", "gpu", "incident", "license", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [407, 487]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: dffbc807acd2f09b8cca8aa6c653fabe44e2c80874ce9e9cfabd831d0bb19e75
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

| Factor | Hardware RAID | ZFS | Ceph |
|---|---|---|---|
| License/capEx | Card + CacheVault modules | Free (HBA only) | Free (more nodes/HBA) |
| OpEx skill | Low (GUI, vendor support) | Medium (CLI, tuning ARC) | High (cluster ops) |
| Vendor support contract | Yes (Broadcom/Microchip) | Community or vendor (TrueNAS) | Community or vendor (IBM/Red Hat) |
| Portability on card failure | Needs identical/similar card | Any HBA, disks move freely | Any node, self-describing OSDs |
- **Rule of thumb:** single node + need simplicity → hardware RAID; single node + data integrity priority → ZFS on HBA; multi-node + scale → Ceph; 20+ NVMe at line rate → GPU/offload RAID class [secondary].
- **Ransomware/backup layering is orthogonal:** none of the above replaces offline backup — see D14/D15 and the 3-2-1-1-0 rule (3 copies, 2 media, 1 offsite, 1 offline/air-gapped, 0 unverified backups) [secondary].
- **How to use the matrices:** filter first by node count (single node eliminates Ceph; multi-node eliminates single-controller RAID), then by integrity requirement (regulated/silent-corruption-sensitive → ZFS or Ceph checksums over hardware RAID), then by skill/budget. Performance is the last filter, not the first — all three can be fast, only one ordering minimizes regret [secondary].
- **TCO framing:** hardware RAID = card + CacheVault + vendor support contract, low admin hours; ZFS = HBA + admin skill, near-zero license; Ceph = 3+ nodes + networking + significant ops skill. The cheapest capEx (Ceph on commodity nodes) can be the most expensive opEx without the team to run it [secondary].
- **Migration paths:** RAID → ZFS is a rebuild (no in-place conversion); ZFS → Ceph is a data migration; mdadm → hardware RAID is a rebuild. Choose with a 5-year horizon — re-platforming petabytes is the real cost [secondary].
- **Hybrid reality:** many 2026 estates run hardware RAID for boot/VM (simplicity), ZFS for bulk (integrity), and Ceph for scale-out — the matrices are per-workload, not per-company [secondary].
- **Erasure-coding overhead examples (Ceph):** k=8,m=3 → 37.5% overhead (vs 200% for 3× replication); k=8,m=2 → 25%. EC is the capacity answer at scale, paid for in CPU and rebuild network traffic [secondary].
- **ZFS RAIDZ overhead examples:** 8-wide RAIDZ2 → 25% overhead; 6-wide RAIDZ2 → 33%. Narrower vdevs cost more overhead but rebuild faster — the same tradeoff as RAID 50/60 subarrays [secondary].
- **Small-file workloads:** erasure coding and wide parity punish small random writes (read-modify-write across the stripe). Databases and VM images on EC pools need careful PG/stripe sizing or stay on replicated/mirrored pools [secondary].
- **The "which RAID" flowchart (single node):** need boot simplicity → HW RAID 1; VM datastore → HW RAID 10 or ZFS mirrors; bulk HDD → ZFS RAIDZ2 or HW RAID 6; max NVMe throughput → SupremeRAID/SmartRAID-4300 class or ZFS mirrors; integrity-first → ZFS [secondary].
- **Multi-node flowchart:** 3+ nodes → Ceph (replication for performance pools, EC for capacity); 2 nodes → ZFS replication (syncoid/zfs send) — Ceph needs 3+ for quorum [secondary].
- **RPO/RTO lens:** hardware RAID gives fast local rebuild (good RTO) but no snapshots; ZFS gives instant snapshots/clones (excellent RPO/RTO for logical errors); Ceph gives self-healing across nodes. Ransomware recovery is a snapshot/air-gap question, not a RAID-level question [secondary].
- **Skill inventory (honest):** hardware RAID needs vendor-GUI skill; ZFS needs CLI + tuning skill; Ceph needs distributed-systems ops. The technology you can operate at 3 a.m. beats the technology with better specs [secondary].
- **Compliance lens:** regulated data may require SED + sanitize logs (D4) regardless of redundancy choice — the decision matrices pick availability architecture; D4 picks the decommission story [secondary].
- **Snapshots change the backup math:** ZFS/Ceph snapshots give near-zero-RPO local recovery; they don't replace offsite/offline copies (a site loss takes snapshots with it) — snapshots are an RTO tool, 3-2-1-1-0 is the survival tool [secondary].
- **Vendor support as risk transfer:** a Broadcom/Microchip/IBM contract transfers 3 a.m. debugging to the vendor; community-supported ZFS/Ceph transfers it to your team. Price the transfer honestly [secondary].
- **Monitoring parity:** whatever the choice, the monitoring burden is similar (SMART, temperatures, scrub status) — the matrices differ in *failure handling*, not in whether you watch the drives [secondary].
- **Exit strategy:** document how to leave each choice (RAID→ZFS rebuild, ZFS→Ceph migration, Ceph→anything via radosgw/s3 copy). The absence of an exit plan is a decision too [secondary].
- **Proof of concept:** 30-day POC with production-like workload before committing petabytes — every vendor's numbers survive contact with your workload differently [secondary].
- **Workload fit examples:** video surveillance (sequential writes, long retention) → RAID 6/ZFS RAIDZ2 on HDD; VDI (random small I/O) → RAID 10/ZFS mirrors on SSD; analytics scratch → RAID 0/SupremeRAID with upstream backup [secondary].
- **Growth pattern:** hardware RAID grows by adding VDs; ZFS grows by adding vdevs; Ceph grows by adding nodes — match the growth axis to the business (capacity vs performance vs scale-out) [secondary].

