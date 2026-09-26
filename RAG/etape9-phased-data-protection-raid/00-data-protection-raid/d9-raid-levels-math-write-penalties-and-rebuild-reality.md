---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d9-raid-levels-math-write-penalties-and-rebuild-reality
title: "D9 — RAID levels: math, write penalties, and rebuild reality"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: ["consumer", "energy", "latency", "nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [275, 299]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 880162560ea8cf76e09bd784b9b4ef3f3a5fa7b36ed0016a6588277d24fce9ec
---

# D9 — RAID levels: math, write penalties, and rebuild reality

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

