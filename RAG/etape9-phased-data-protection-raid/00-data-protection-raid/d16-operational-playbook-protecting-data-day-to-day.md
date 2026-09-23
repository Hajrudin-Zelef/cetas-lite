---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d16-operational-playbook-protecting-data-day-to-day
title: "D16 — Operational playbook: protecting data day to day"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["AMD"]
dates: ["2026-09-22"]
keywords: ["consumer", "cost", "energy", "incident", "nand", "pricing", "research"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [550, 613]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 7561a9525bec2aa6427344e9a8fd06643d7a9d253ef38c8887cc9f9773d5a4da
---

# D16 — Operational playbook: protecting data day to day

## D16 — Operational playbook: protecting data day to day

- **Power:** enterprise SSDs with PLP + RAID cards with CacheVault/supercap + UPS for the rack = three independent layers. Test the failure path (pull power during write load) in staging, not in production [secondary].
- **Monitoring baseline:** ingest NVMe SMART (`critical_warning`, `percentage_used`, `media_errors`, temperature) and SAS/SATA SMART (5/197/198/199) into fleet telemetry; alert on any non-zero critical_warning or media_errors immediately — these are leading indicators, not lagging ones [secondary].
- **Scrubbing:** schedule ZFS scrubs / RAID consistency checks / Ceph deep-scrubs; patrol reads on hardware RAID catch latent sector errors before a rebuild needs those sectors [secondary].
- **Firmware lifecycle:** inventory drive + HBA/RAID firmware versions; apply vendor-validated updates on a maintenance cadence; keep a rollback plan (dual-bank where supported) [secondary].
- **Sanitize on decommission:** NVMe sanitize (block erase for non-SED, crypto erase for SED), record per-serial logs; never trust single-pass overwrite on SSDs [secondary].
- **Backup rule:** 3-2-1-1-0 — three copies, two different media, one offsite, one offline/air-gapped (tape/LTO-10 or disconnected disk), zero unverified backups (test restores) [secondary].
- **Spares on the shelf:** keep cold-spare drives of each deployed model (firmware-matched where possible); a rebuild waiting on procurement is how degraded arrays become dead arrays [secondary].
- **Label everything:** serial, firmware, purchase date, and array membership on physical labels — during a 2 a.m. drive pull, `by-id` strings are not readable in a dark rack [secondary].
- **Test the power-fail path:** pull power (or simulate via managed PDU) during write load in staging with PLP SSDs + CacheVault; confirm no filesystem or array inconsistency. Untested protection is a hope, not a control [secondary].
- **Firmware rollback plan:** stage previous-known-good firmware for HBAs, RAID cards, and SSDs before fleet updates; dual-bank devices make this trivial, single-bank devices make it a maintenance window [secondary].
- **Restore drills:** the "0 unverified backups" in 3-2-1-1-0 means scheduled test restores with checksums, including from the tape/optical copy — a backup never restored is Schrödinger's backup [secondary].
- **Document the array:** RAID level, stripe size, cache policy, spare policy, and controller model per array in the runbook; the person rebuilding at 2 a.m. is rarely the person who built it [secondary].- **Runbook template (per array):** controller/HBA model + firmware; RAID level, stripe size, cache policy; member serials + slot map; spare policy; patrol/scrub schedule; firmware versions of member drives; last consistency-check date; restore-test date. One page per array, reviewed quarterly [secondary].
- **Change control:** firmware updates, drive replacements, and cache-policy changes go through change tickets with rollback plans — storage changes are the highest-blast-radius changes in the fleet [secondary].
- **Monitoring dashboard minimum:** per-drive SMART status, array optimal/degraded state, temperature top-talkers, scrub/patrol completion, backup completion + restore-test status. If it's not on a dashboard, it's not monitored [secondary].
- **Incident rehearsal:** practice the degraded-array drill (identify failed drive via SES locate, pull, replace, monitor rebuild) before it happens — the first rebuild should not be during a real failure [secondary].
- **Decommission checklist:** sanitize (D4) → verify → log serial → remove from monitoring/inventory → physical destruction or resale per classification [secondary].
- **Firmware lifecycle calendar:** quarterly review of vendor firmware bulletins; semi-annual maintenance windows for HBA/RAID/SSD updates; emergency process for CVEs. Unpatched storage firmware is technical debt with data-loss interest [secondary].
- **Capacity planning:** track usable-vs-raw, endurance burn rate (percentage_used/month), and temperature trends together — the three curves that tell you when to buy [secondary].
- **Documentation location:** runbooks in version control, not in someone's head or a wiki nobody reads; the 2 a.m. test — can a competent stranger rebuild the array from the runbook alone? [secondary].
- **Power-fail testing cadence:** repeat the PDU pull test after every firmware update and annually — protection is a property of the current firmware, not a permanent attribute [secondary].
- **Post-incident review:** every degraded array, thermal event, or sanitize failure gets a blameless review with runbook updates — the playbook compounds [secondary].
- **Onboarding checklist (new array):** burn-in drives → record firmware/serials → configure RAID/ZFS → set cache policy → enable patrol/scrub → configure SMART alerting → power-fail test → first backup + restore test → runbook entry. Skipping steps is borrowing trouble [secondary].
- **Offboarding checklist (retired array):** final backup → sanitize per D4 → verify → logs → inventory removal → destruction/resale per classification [secondary].
- **Quarterly storage review:** AFR vs expected, endurance burn, temperature trends, scrub completion, backup/restore-test status, firmware currency, spare inventory — one meeting, all arrays [secondary].
- **Capacity runway:** alert at 80% usable on every tier; procurement lead times exceed fill rates on fast-growing datasets [secondary].
- **Single-pane-of-glass:** array state, SMART, backup status, and tape library status in one view — incidents correlate across layers (a thermal event causing both throttling and backup misses) [secondary].

## D17 — Worked numerical examples (reproducible arithmetic)

- **Example 1 — RAID 6 capacity:** 12 × 20 TB HDD in RAID 6: usable = (12 − 2) × 20 TB = 200 TB raw (≈ 181.9 TiB). Parity overhead = 2/12 ≈ 16.7% [secondary].
- **Example 2 — URE probability during RAID 5 rebuild (the famous calculation):** expected uncorrectable errors = bits_read ÷ UBER_bits. 8 × 20 TB RAID 5 rebuild reads 7 × 20 TB = 140 TB = 1.12×10^15 bits. At consumer UBER 10^15: expected errors ≈ 1.12; P(no error) ≈ e^−1.12 ≈ 33% → ~67% chance the rebuild hits an uncorrectable error. This single calculation is why RAID 5 is retired for large HDDs [secondary].
- **Example 3 — same array at enterprise UBER 10^16:** expected errors ≈ 0.112; P(at least one) ≈ 1 − e^−0.112 ≈ 10.6%. Better, but still a one-in-ten rebuild failure — RAID 6, not RAID 5 [secondary].
- **Example 4 — SSD RAID 5 at UBER 10^17:** 8 × 3.84 TB RAID 5 rebuild reads 7 × 3.84 TB = 26.88 TB = 2.15×10^14 bits; expected errors ≈ 0.0022 → P(failure) ≈ 0.2%. Negligible — SSD RAID 5 is UBER-safe (though write-penalty and wear still apply) [secondary].
- **Example 5 — rebuild time exposure:** 20 TB HDD at 250 MB/s sustained ≈ 22.2 hours minimum per-disk rebuild; during this window the array is degraded and all surviving disks stream continuously — the period Examples 2–3 quantify [secondary].
- **Example 6 — fleet AFR sizing:** 1,000 drives at 0.44% AFR → λ = 4.4 expected failures/year. Poisson P(≥8 failures) ≈ 7% — size cold spares for the tail (8–10), not the mean (4.4) [secondary].
- **Example 7 — DWPD headroom:** workload writes 5 TB/day to a 3.84 TB drive rated 1 DWPD: actual = 5 ÷ 3.84 ≈ 1.30 DWPD → exceeds rating by 30%; endurance life ≈ 5 yr ÷ 1.3 ≈ 3.8 years. Size up capacity or pick a 3-DWPD SKU [secondary].
- **Example 8 — write-penalty IOPS:** 8 × 10K-IOPS SSDs: RAID 10 write IOPS = (4 × 10K)/2 = 20K; RAID 6 = (6 × 10K)/6 = 10K; RAID 5 = (7 × 10K)/4 = 17.5K. Read IOPS ≈ sum of data disks in all cases [secondary].
- **Example 9 — LTO-10 archive media count:** 1 PB (≈ 1,000 TB) on 40 TB native cartridges: 1,000 ÷ 40 = 25 cartridges + 10% spare/handling ≈ 28 cartridges. At 2.5:1 compressible data the same 25 cartridges hold up to 2.5 PB — never budget on this [official].
- **Example 10 — 3-2-1-1-0 capacity overhead:** 100 TB primary → 100 TB (onsite backup) + 100 TB (offsite) + 100 TB (offline tape) = 400 TB total raw, before RAID/erasure overhead. Tape's $/TB is what makes the fourth copy affordable [secondary].
- **Example 11 — expander uplink sizing:** 24 × 550 MB/s SATA SSDs = 13.2 GB/s aggregate; one x4 24G SAS uplink ≈ 9.6 GB/s raw → oversubscribed 1.4:1; provision dual x4 uplinks (19.2 GB/s) for full-bore flash shelves [secondary].
- **Example 12 — NVMe lane budget:** 32 × Gen4 x4 SSDs need 128 PCIe 4.0 lanes for full bandwidth — a 2-socket EPYC platform can do this; a single-socket workstation cannot. The RAID product is rarely the bottleneck; the platform is [secondary].
- **Example 13 — ZFS RAIDZ2 capacity:** 8 × 20 TB RAIDZ2: usable ≈ 6 × 20 TB = 120 TB (before metadata/slop). Two vdevs of 8 (16 disks) striped: 240 TB, and a rebuild touches one 8-disk vdev — narrower failure domains, same math as RAID 60 [secondary].
- **Example 14 — Ceph EC 8+3 on 12 nodes:** usable = 8/11 ≈ 72.7% of raw; a 12 × 20 TB cluster ≈ 174 TB usable. Rebuild traffic spreads across all nodes — fast, but mind the network [secondary].
- **Example 15 — SSD fleet endurance:** 100 × 3.84 TB 1-DWPD drives = 384 TB raw; daily fleet write budget at rating = 384 TB/day. A 200 TB/day logging workload fits at ~52% of endurance budget [secondary].
- **Example 16 — RAID 10 vs RAID 6 usable on 12 × 20 TB:** RAID 10: 120 TB usable, best random write; RAID 6: 200 TB usable, 6× write penalty. The 80 TB difference is the price of write performance [secondary].
- **Example 17 — tape restore time:** 25 × 40 TB cartridges at 400 MB/s-class drives ≈ 28 hours per cartridge worst case — the offline copy is for disasters, not for quick restores. Keep a disk-based copy for operational recovery (3-2-1-1-0 has *two* onsite tiers for this reason) [secondary].
- **Example 18 — $/TB sanity (illustrative, not quotes):** at street prices, enterprise NVMe ≈ 10–20× the $/TB of HDD, LTO media ≈ 1/5 the $/TB of HDD. Tiering (NVMe → HDD → tape) is an economic necessity, not a preference [secondary].
- **Example 19 — RAID 10 rebuild scope:** RAID 10 rebuild copies one mirror disk (≈ 20 TB at 250 MB/s ≈ 22 h) but reads only *one* surviving disk — the URE exposure of Example 2 shrinks ~7× vs RAID 5. This is the quantitative case for RAID 10 on HDDs [secondary].
- **Example 20 — endurance in years:** drive rated 1 DWPD/5 yr, workload 0.5 DWPD → endurance life ≈ 10 years (warranty still 5). Workload 2 DWPD → ≈ 2.5 years. The rating is a budget; spend it deliberately [secondary].
- **Example 21 — CacheVault energy (order of magnitude):** destaging 8 GB of write-back cache to NAND at ~1 GB/s takes ~8 s of hold-up; supercap modules are sized with margin for this plus controller shutdown sequencing — which is why a degraded/failed supercap forces write-through mode (the controller refuses unsafe write-back) [secondary].
- **Example 22 — patrol read coverage:** 20 TB at 250 MB/s ≈ 22 h per full patrol pass; weekly patrol on a 12-disk array means each disk's surface is fully scanned roughly weekly — latent errors surface within days, not months [secondary].
- **Example 23 — RAID 5 vs RAID 6 usable (12 × 20 TB):** RAID 5: 220 TB usable, 1-disk fault tolerance, ~67% rebuild-failure risk at consumer UBER (Example 2). RAID 6: 200 TB usable, 2-disk tolerance. The 20 TB "savings" of RAID 5 buys a two-in-three chance of rebuild failure — the worst trade in storage [secondary].
- **Example 24 — Ceph 3× replication vs EC 8+3:** 100 TB raw → 33 TB usable (3×) vs 73 TB usable (EC 8+3). EC more than doubles usable capacity at the cost of encode CPU and small-write penalties [secondary].
- **Example 25 — supercap hold-up energy (illustrative):** destaging 8 GB at 1 GB/s needs ~8 s at ~25 W ≈ 200 J plus sequencing margin — the reason CacheVault modules are supercapacitor banks, not coin cells [secondary].
- **Example 26 — fleet temperature:** 24 × 10 W SSDs = 240 W in a 2U shelf; at 20 CFM airflow that's roughly a 20 °C rise over inlet — inlet 25 °C → drives near 45–50 °C before workload heat. Cooling math belongs in the BOM [secondary].
- **Example 27 — RAID 1 vs RAID 10 (4 disks):** RAID 1 with 4 members is possible on some controllers (n-way mirror, 25% usable); RAID 10 gives 50% usable with better read scaling — prefer RAID 10 at 4+ disks [secondary].
- **Review this file's numbers annually:** vendor lineups (9600/4300/SupremeRAID), LTO media pricing, and firmware CVEs move fast — the 2026-09-22 cutoff is the freshness boundary [secondary].
- **Cross-reference:** this phase pairs with Step 9 Phase C (enterprise SSD models/endurance) for drive selection and Phase B (filesystems) for ZFS/btrfs internals [secondary].
- **Feedback loop:** every incident in production should add a bullet here — a research file that doesn't grow with operations goes stale [secondary].
- **Done beats perfect:** ship the protection you can operate today (PLP SSDs + monitored RAID + tested backups) over the ideal architecture you'll deploy "next quarter" [secondary].
- **Example 28 — 3-2-1-1-0 on 50 TB:** primary 50 TB → onsite backup 50 TB → offsite 50 TB → offline tape 50 TB = 200 TB raw; with RAID 6 (16.7% overhead) on disk tiers ≈ 220 TB disk + 2 LTO-10 cartridges. Small enough to be cheap, complete enough to survive ransomware [secondary].

