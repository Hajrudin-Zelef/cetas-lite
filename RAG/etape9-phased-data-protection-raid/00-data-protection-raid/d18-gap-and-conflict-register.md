---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d18-gap-and-conflict-register
title: "D18 — Gap and conflict register"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom", "Samsung"]
dates: []
keywords: ["asic", "benchmarks", "dram", "memory", "nand", "pricing", "serdes", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [614, 683]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 187b38b38417a9e223532cbb09b617290cbcfa73b2e69da6df2f70c100665a50
---

# D18 — Gap and conflict register

## D18 — Gap and conflict register

- **Conflict 1 — SupremeRAID performance:** GRAID's current brochure claims 28M IOPS / 260 GB/s; launch-era press reported 19M IOPS / 110 GB/s read, 1.5M IOPS / 22 GB/s write. Likely different test configs, but neither is independently verified — do not compare generations or competitors on these numbers [vendor-reported] [conflict].
- **Conflict 2 — PM9A3 TCG profile:** Samsung U.2 datasheet says TCG Opal; Samsung product brief says TCG Enterprise. Reconcile before compliance claims [secondary] [conflict].
- **Conflict 3 — 9660-16i "up to" performance:** Newegg lists 28 GB/s and 2.4M IOPS (x8 interface!) while other listings say 7 GB/s and 2.4M 4K IOPS. The 28 GB/s figure likely aggregates across a topology, not the x8 link — treat as listing error until the datasheet confirms [secondary] [conflict].
- **Gap 1 — SmartRAID 4300 independent benchmarks:** none found at cutoff; all throughput figures are Microchip/vendor-reported [unverified].
- **Gap 2 — ODA 2026 roadmap:** no current Sony statement on next-gen Optical Disc Archive found; availability uncertain [unverified].
- **Gap 3 — LTO-10 drive/media street pricing:** not captured; launch pricing volatile [unverified].
- **Gap 4 — SSD/HBA firmware CVE database:** no consolidated 2025–2026 source found [unverified].
- **Gap 5 — DIF/DIX deployment guides:** IBM T10 DIF docs and Linux DIX references are aging; current kernel/HBA support matrix for 2026 not verified [unverified].
- **Non-comparable benchmarks:** vendor "up to" IOPS/GB/s across MegaRAID 9600, SmartRAID 4300, and SupremeRAID use different drive counts, queue depths, and block sizes — cross-vendor comparison is invalid [secondary].
- **Gap 6 — LTO-10 native transfer rate:** Fujifilm/HPE spec tables captured through LTO-9 rates; LTO-10 drive data rate not found at cutoff [unverified].
- **Gap 7 — Backblaze SSD coverage:** Backblaze's public drive-stats are HDD-centric; SSD AFR/SMART fleet data at Backblaze scale is thinner — vendor MTBF claims for SSDs have weaker independent cross-checks [independent].
- **Gap 8 — Current Linux DIX/DIF support matrix:** most DIX documentation dates to the 2010s; which 2026 kernels, filesystems, and HBA drivers support which PI types is not verified [unverified].
- **Gap 9 — CacheVault hold-up validation data:** vendor describes the mechanism; independent measurements of cache-destage success across supercap aging were not found [unverified].
- **ZFS spare behavior:** ZFS hot spares activate on device fault automatically; unlike RAID copyback, ZFS keeps the spare in place (detach/replace when convenient) — simpler operations [secondary].
- **Ceph OSD memory:** BlueStore wants GBs of RAM per OSD for RocksDB; under-provisioned OSD hosts flap — size RAM with the OSD count, not just the CPU [secondary].

## D19 — Glossary

- **AFR:** Annualized Failure Rate — % of a drive population failing per year; the practical form of MTBF [secondary].
- **ARC:** Adaptive Replacement Cache — ZFS's RAM cache [secondary].
- **Backfill:** Ceph's term for redistributing data after topology change [secondary].
- **CRUSH:** Controlled Replication Under Scalable Hashing — Ceph's data-placement algorithm [secondary].
- **dRAID:** distributed RAIDZ — ZFS distributed-parity variant with faster rebuilds [secondary].
- **HCTM:** Host-Controlled Thermal Management (NVMe) [secondary].
- **Resilver:** ZFS term for rebuilding redundancy onto a replacement device [secondary].
- **BBU:** Battery Backup Unit — legacy cache-protection battery, superseded by CacheVault/supercap [secondary].
- **CacheVault:** Broadcom's supercapacitor + NAND flash module protecting RAID write-back cache across power loss [secondary].
- **DIF/DIX:** T10 Data Integrity Field / Data Integrity eXtensions — SCSI/NVMe end-to-end protection information [secondary].
- **DWPD:** Drive Writes Per Day — endurance rating over the warranty period [secondary].
- **eHBA:** enhanced HBA — Broadcom 9600 passthrough adapters with secure boot, attestation, SED management [secondary].
- **FTL:** Flash Translation Layer — SSD firmware mapping logical to physical NAND addresses [secondary].
- **LTFS:** Linear Tape File System — open tape-as-filesystem format [secondary].
- **OP:** Over-provisioning — spare NAND reserved for GC/wear leveling [secondary].
- **PLP:** Power Loss Protection — capacitor-backed flush of volatile SSD state to NAND [official].
- **SED:** Self-Encrypting Drive — hardware encryption with media encryption key, TCG Opal/Enterprise [secondary].
- **SES:** SCSI Enclosure Services — expander/backplane management protocol [secondary].
- **SPDM:** Security Protocol and Data Model — DMTF device attestation/authentication standard [secondary].
- **SPOR:** Sudden Power-Off Recovery — firmware restoration after unsafe shutdown (does not save in-flight data) [secondary].
- **TBW:** Terabytes Written — absolute endurance rating [secondary].
- **Tri-mode:** single SerDes serving NVMe + SAS + SATA on one backplane [secondary].
- **UBER:** Uncorrectable Bit Error Rate — e.g. 1 sector per 10^17 bits read [official].
- **UBM:** Universal Bay Management (SFF-TA-1005) — tri-mode backplane management [secondary].
- **WAF:** Write Amplification Factor — NAND writes ÷ host writes [secondary].
- **3-2-1-1-0:** 3 copies, 2 media types, 1 offsite, 1 offline/air-gapped, 0 unverified backups [secondary].
- **BER:** Bit Error Rate — raw media error rate before ECC [secondary].
- **CCT/WCT:** Critical/Warning Composite Temperature thresholds (NVMe SMART) [secondary].
- **Copyback:** RAID controller operation returning data from hot spare to the replaced disk's slot [secondary].
- **DIX modes:** READ_STRIP/WRITE_INSERT, READ_INSERT/WRITE_STRIP, READ_PASS/WRITE_PASS — controller behaviors for PI handling [secondary].
- **ECC RAM:** error-correcting memory; recommended (not mandatory) with ZFS [secondary].
- **FUA:** Force Unit Access — command flag requiring media commit before completion [secondary].
- **HBA:** Host Bus Adapter — passthrough disk controller, no RAID logic [secondary].
- **IT mode:** initiator-target (passthrough) firmware on a SAS HBA [secondary].
- **JBOD:** Just a Bunch Of Disks — disks presented individually [secondary].
- **MTTDL:** Mean Time To Data Loss — array-level reliability metric incorporating rebuild windows [secondary].
- **NVDIMM-N:** DRAM DIMM with NAND + supercap for persistent memory [secondary].
- **Patrol read:** background media scan for latent sector errors on hardware RAID [secondary].
- **P/E cycles:** program/erase cycles — the finite endurance unit of NAND blocks [secondary].
- **ROC:** RAID-on-Chip — the ASIC at the heart of a hardware RAID card [secondary].
- **RTU/RTD/RTRec:** 800-193 Roots of Trust for Update, Detection, Recovery [official].
- **SES:** SCSI Enclosure Services — expander/backplane management protocol [secondary].
- **SPDM:** Security Protocol and Data Model — DMTF device attestation/authentication standard [secondary].
- **SPOR:** Sudden Power-Off Recovery — firmware restoration after unsafe shutdown (does not save in-flight data) [secondary].
- **STP:** SATA Tunneling Protocol — how SAS expanders carry SATA traffic [secondary].
- **TCG:** Trusted Computing Group — Opal/Enterprise SED standards [secondary].
- **UBM:** Universal Bay Management (SFF-TA-1005) — tri-mode backplane management [secondary].
- **WAF:** Write Amplification Factor — NAND writes ÷ host writes [secondary].
- **WORM:** Write Once Read Many — non-rewritable compliance media [official].

