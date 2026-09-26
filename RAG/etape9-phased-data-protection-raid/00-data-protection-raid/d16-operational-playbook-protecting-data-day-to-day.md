---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d16-operational-playbook-protecting-data-day-to-day
title: "D16 — Operational playbook: protecting data day to day"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["consumer", "pricing", "throughput"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [528, 551]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 63457999e59893154d1d46a6bb8da4515c0a0a3513436dafbf760e9736998531
---

# D16 — Operational playbook: protecting data day to day

- **Sony Optical Disc Archive (ODA):** Gen 3 cartridges hold 5.5 TB (11 × 500 GB discs), rated for 100-year media life in controlled conditions; Gen 2 was 3.3 TB. Positioned for cold archive, broadcast, and compliance [secondary].
- **Performance envelope (secondary reporting):** ~187.5 MB/s write, ~375 MB/s read per drive; library pricing historically ~$9,000/drive with ~$184 per 5.5 TB cartridge — an order of magnitude pricier per TB than tape [secondary].
- **Archival Disc standard:** Sony/Panasonic 300 GB (later 500 GB/1 TB roadmap) write-once media with 50+ year claims; consumer Blu-ray M-DISC (100 GB BDXL) marketed at 1,000-year claims by vendors — treat longevity claims as accelerated-aging extrapolations, not observed data [secondary].
- **2026 status:** optical archive remains a niche vs tape; no 2025–2026 ODA generation refresh was found at cutoff. Availability of drives/media and vendor roadmap commitment in 2026 is a flagged gap — verify with Sony Professional before designing around it [unverified].
- **Decision note:** for petabyte-scale cold archive in 2026, LTO-10 dominates on $/TB; optical fits small, truly write-once compliance sets where cartridge-level immutability and random access matter [secondary].
- **ODA generations (sourced):** Gen 2 cartridges: 3.3 TB; Gen 3 (ODS-D380U class drives): 5.5 TB per cartridge (11 × 500 GB discs), 100-year media-life rating. Gen 3 per-drive throughput ~3 TB/hour-class (verify against current Sony spec — figures vary by source) [secondary].
- **M-DISC (consumer write-once):** 100 GB BDXL discs with an inorganic recording layer marketed on millennium-scale longevity claims — accelerated-aging extrapolations, not observed data; suitable for small personal/legal archives, not data-center scale [secondary].
- **Historical hyperscale dabbling:** Facebook's 2014 cold-storage Blu-ray experiments showed optical could work technically but lost economically to HDD/tape at scale — the lesson still holds in 2026 [secondary].
- **When optical still fits:** small write-once compliance sets, legal hold, and air-gapped vault copies where cartridge-level WORM semantics and random access beat tape's sequential restore [secondary].
- **ODA vs tape economics (secondary reporting):** ~$184 per 5.5 TB cartridge vs LTO media at a fraction of that per TB; ~$9,000 per ODA drive. Optical only pencils out where its random-access WORM semantics are specifically required [secondary].
- **Disc rot and media QC:** cheap consumer optical media degrades; archival-grade (M-DISC, ODA) uses inorganic layers. Whatever the medium, verify-on-write and periodic re-reads are the actual guarantee — the "100-year" number is a lab extrapolation [secondary].
- **Format obsolescence risk:** a 100-year disc needs a working drive in 100 years. Tape has the same problem (mitigated by the LTO compatibility rule and active migration); optical's smaller market makes drive availability the bigger risk — the D18 gap [secondary].
- **Optical jukeboxes/libraries:** Sony's PetaSite libraries scale ODA to hundreds of cartridges with robotic handling — the optical answer to tape libraries, at optical economics [secondary].
- **Write verification:** ODA's verify-after-write and defect management are the actual integrity mechanism; the century rating is meaningless without them enabled and logged [secondary].
- **Blu-ray cold storage economics:** consumer BDXL media is cheap per disc but the $/TB and robotic-handling story never reached data-center viability — the niche stays personal/legal archiving [secondary].
- **Tape roadmaps as planning input:** the LTO Program's LTO-14 913 TB target signals decades of continued investment — tape is not a dying medium in 2026, it's a growing one by exabytes shipped [official].
- **Air-gap procedure:** write → verify (read-back compare) → eject → barcode-log → offsite vault. A tape verified and shelved is the strongest ransomware control in this file [secondary].
- **Tape drive sharing:** libraries share drives across hosts via partitioning or ISLs — size drive counts for concurrent backup windows, not for total slots [secondary].
- **Media serialization:** barcode-label every cartridge and track it in the library ISV — an unlabeled tape is an unfindable tape during a restore [secondary].
- **Offsite rotation:** grand-father-son or Tower of Hanoi rotation for the offline set; the offline copy only works if it's actually offsite and rotated [secondary].
- **Optical as ransomware vault:** written-once optical cartridges shelved offline share tape's air-gap property at smaller scale — viable for the "1 offline" copy of modest datasets [secondary].

## D16 — Operational playbook: protecting data day to day

