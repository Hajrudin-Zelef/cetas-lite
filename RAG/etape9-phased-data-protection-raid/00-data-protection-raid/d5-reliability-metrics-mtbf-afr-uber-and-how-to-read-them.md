---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d5-reliability-metrics-mtbf-afr-uber-and-how-to-read-them
title: "D5 — Reliability metrics: MTBF, AFR, UBER, and how to read them"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["consumer", "datacenter", "memory", "nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [139, 174]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 1c71dbac4db28a7256a0a4be8f3203086bb1cb6d5458894cd9e5709df435cd40
---

# D5 — Reliability metrics: MTBF, AFR, UBER, and how to read them

- **Key escrow vs sanitize:** enterprises escrowing SED passwords/keys must destroy escrowed copies as part of decommission — a sanitized drive with an escrowed key is not sanitized [secondary].
- **Host-visible PLP status:** NVMe critical_warning bit 4 (volatile memory backup device failed) is the host-visible PLP health signal — alert on it like any other critical bit [secondary].
- **PLP and namespace formats:** secure erase/sanitize operations also depend on capacitors for metadata consistency — a drive that can't complete sanitize after power events has a PLP problem, not just a firmware quirk [secondary].
- **Qualification sampling:** validate PLP on a sample of *every* firmware revision, not just every model — firmware changes the flush routine [secondary].

## D5 — Reliability metrics: MTBF, AFR, UBER, and how to read them

- **MTBF (Mean Time Between Failures):** statistical, not a lifespan promise. A 2,000,000-hour MTBF means that in a large population, one failure is expected per 2M device-hours — roughly a 0.44% annualized failure rate — not that any single drive lasts 228 years [secondary].
- **AFR (Annualized Failure Rate):** the practical translation; enterprise SSDs typically target ~0.35–0.44% AFR at 2M-hour MTBF. Backblaze-style fleet data is the independent check against vendor MTBF claims — vendor and observed AFR can diverge under real thermals and workloads [independent].
- **UBER (Uncorrectable Bit Error Rate):** the media-read reliability floor. Enterprise SSDs commonly specify 1 sector per 10^17 bits read (e.g. Samsung PM9A3 product brief, U.2/E1.S/M.2 family spec) [official].
- **UBER math that matters:** 10^17 bits = 12.5 petabytes read per expected uncorrectable sector. Consumer drives often specify 10^15 (1 sector per 125 TB read) — a 100x gap that drives RAID rebuild-risk calculations (see D9) [secondary].
- **DWPD/TBW endurance vs UBER are orthogonal:** a drive can be within endurance (PM9A3: 1.0 DWPD for 5 years) and still hit UBER on reads; monitor both `percentage_used` and `media_errors` [secondary].
- **Write endurance example:** 3.84 TB PM9A3 at 1 DWPD over 5 years ≈ 7,008 TBW (retailer-computed; Samsung publishes DWPD, not per-capacity TBW — do not present the computed TBW as a Samsung spec) [secondary].
- **Conflict flagged:** Samsung's own documents disagree on the PM9A3 encryption standard — the U.2 datasheet says TCG Opal while the product brief says TCG Enterprise. Treat any single-source claim about its TCG profile as suspect until reconciled [secondary].
- **AFR conversion math:** AFR ≈ 1 − e^(−8760/MTBF). For 2,000,000-hour MTBF: 8760/2,000,000 = 0.00438 → AFR ≈ 0.44%. For 1,000,000-hour MTBF → ≈ 0.87%. Quote AFR in procurement, not MTBF — it is the number capacity planners can actually use [secondary].
- **Spec conditions matter:** MTBF figures are typically quoted at a rated temperature and workload (often 40 °C, enterprise workload per JEDEC JESD219). A drive running at 60 °C ambient in a dense chassis is not operating under the spec's assumptions [secondary].
- **UBER scale table:** 1 sector / 10^17 bits ≈ one error per 12.5 PB read (enterprise SSD); 10^16 ≈ 1.25 PB (enterprise HDD); 10^15 ≈ 125 TB (consumer/desktop drives). RAID-5 rebuild viability on large HDDs is decided by this row (see D9) [secondary].
- **JEDEC JESD219:** the standard enterprise endurance workload (random + sequential mix) behind DWPD ratings — a drive rated 1 DWPD under JESD219 may wear faster under a pure 4K-random-write workload with high WAF [secondary].
- **Endurance formula:** TBW ≈ DWPD × capacity_TB × 365 × warranty_years. PM9A3 3.84 TB at 1 DWPD × 5 yr ≈ 7,008 TBW — a retailer-computed figure; Samsung publishes DWPD, so present computed TBW as derived, not as spec [secondary].
- **Bathtub curve:** infant mortality (first months — burn-in and early SMART screening catch these), useful life (flat low failure rate — where AFR applies), wear-out (NAND exhaustion, bearing wear on HDD). Fleet replacement policy should be driven by observed AFR + SMART, not by MTBF alone [secondary].
- **Warranty vs reliability:** enterprise SSD warranties are typically 5 years or the endurance limit, whichever comes first — warranty length is a commercial term, not a reliability prediction [secondary].
- **Backblaze as independent check:** Backblaze publishes quarterly drive-stats with observed AFR by model across hundreds of thousands of drives — the standard independent cross-check of vendor MTBF/AFR claims for HDDs (SSD stats coverage is thinner) [independent].
- **MTBF→AFR worked conversions:** 2,000,000 h → 0.44%/yr; 1,500,000 h → 0.58%/yr; 1,000,000 h → 0.87%/yr; 750,000 h → 1.16%/yr. Fleet of 1,000 drives at 0.44% AFR expects ~4.4 failures/year — size cold spares and support SLAs from this, not from "2M hours sounds like forever" [secondary].
- **Warranty-end planning:** at 5-year warranty with 0.44% AFR, cumulative failure ≈ 2.2% of the fleet — but the bathtub curve means year-1 and post-warranty rates differ; budget refresh at warranty end, not at MTBF-implied "lifetime" [secondary].
- **Enterprise HDD MTBF examples:** 2.5M-hour MTBF (common on datacenter HDDs) → ≈ 0.35% AFR. Compare like-for-like: a 2.5M-hr HDD and a 2M-hr SSD have similar *rated* AFRs; observed AFRs diverge by workload and thermals [secondary].
- **AFR vs observed (Backblaze):** published Backblaze AFRs for popular HDD models range ~0.5–2%+ depending on model and age — the spread between best and worst models exceeds the difference between most MTBF specs. Model selection beats spec-sheet shopping [independent].
- **Workload derating:** MTBF/AFR assume rated workload (often 550 TB/year for enterprise). A drive running 3× the rated workload ages proportionally faster — endurance (D6) and AFR interact [secondary].
- **MTTDL (array level):** Mean Time To Data Loss folds AFR, rebuild time, and UBER into one number; RAID 6 MTTDL on large HDDs is dominated by the rebuild-window double-fault term — the quantitative version of D9's rebuild discussion [secondary].
- **Non-operational specs:** shock/vibration ratings matter for shipping and mobile use; a drive that survives the datasheet's operating shock may not survive a drop onto concrete — handle RMAs accordingly [secondary].
- **Altitude and helium:** datacenter HDDs are sealed He (lower drag, more platters); altitude derating in specs refers to cooling efficiency, not the sealed drive [secondary].
- **Power-loss count vs AFR:** Backblaze-style analyses show unsafe-shutdown counts correlate weakly with failure on their own — but combined with SMART 5/197 they strengthen the signal [independent].
- **Misdirected-write anatomy:** the classic case DIF catches — firmware writes block N's data to LBA M; the reference tag (expected LBA N) mismatches and the drive rejects the write instead of silently corrupting M [secondary].
- **PI and RAID rebuilds:** a rebuild that copies a corrupted-but-valid sector propagates corruption; DIF/DIX lets the controller detect the guard mismatch during the rebuild read and reconstruct from parity instead [secondary].

## D6 — Wear leveling, garbage collection, over-provisioning, TRIM

