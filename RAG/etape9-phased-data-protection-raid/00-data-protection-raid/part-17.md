---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-17
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 17)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: ["capex", "cost", "gpu", "license", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [460, 487]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: e0e2f80e0e440fca65164f42faef4426f78ce10a97aa1caa00f02ff7bcdf4c61
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 17)

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

