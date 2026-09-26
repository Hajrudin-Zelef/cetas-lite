---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d17-worked-numerical-examples-reproducible-arithmetic
title: "D17 — Worked numerical examples (reproducible arithmetic)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [552, 580]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: e9bac1c59b0773186d2826f2ee9acd8914951aacff8a3cae23c1e9cd10a4b4c2
---

# D17 — Worked numerical examples (reproducible arithmetic)

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

