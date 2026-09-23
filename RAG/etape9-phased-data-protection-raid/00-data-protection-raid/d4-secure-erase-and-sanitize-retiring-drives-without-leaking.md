---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d4-secure-erase-and-sanitize-retiring-drives-without-leaking
title: "D4 — Secure erase and sanitize: retiring drives without leaking data"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["consumer", "datacenter", "memory", "nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [107, 172]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 7571487d53042f6918152d9b65c9be1fd50d6c2addb8a36c0a2acaae35cbb953
---

# D4 — Secure erase and sanitize: retiring drives without leaking data

## D4 — Secure erase and sanitize: retiring drives without leaking data

- **NVMe defines two distinct erasure mechanisms** and they are not interchangeable: `sanitize` (whole controller — all namespaces, caches, over-provisioned areas; survives reboot; maps to NIST SP 800-88 "Purge") vs `format --ses` (single namespace; may not survive interruption; "Clear" level) [secondary].
- **NVMe Sanitize actions:** sanact=2 Block Erase (resets all NAND cells, 1–5 minutes typical), sanact=3 Overwrite (writes a pattern across all blocks — slow, adds P/E wear, not recommended for NAND), sanact=4 Crypto Erase (destroys the media encryption key on self-encrypting drives, completes in ~1 second) [secondary].
- **NVMe Format Secure Erase Settings:** ses=0 no erase, ses=1 user-data erase, ses=2 cryptographic erase (namespace-scoped) [secondary].
- **Check capability first:** `nvme id-ctrl /dev/nvme0 | grep -i sanicap`; poll progress with `nvme sanitize-log` — SSTAT 0x101 with SPROG 65535 indicates successful completion [secondary].
- **SATA/SAS SSD path:** ATA Enhanced Secure Erase via hdparm (set password, issue erase); `blkdiscard --secure` only on drives reporting deterministic TRIM with read-zero-after-trim; single-pass `shred` is NOT sufficient on SSDs because wear-leveling reserve blocks hide stale data [secondary].
- **SED shortcut:** on drives with documented always-on encryption, Crypto Erase is the preferred retirement method — destroying the media encryption key renders raw flash unreadable even under chip-off forensics [secondary].
- **Operational rules:** never sanitize the boot drive from the running OS — use a live USB or pull the drive to a sanitization workstation; record sanitize completion logs per serial number for compliance; for regulated data, Purge (sanitize) or physical destruction, not Clear [secondary].
- **HDD note:** for spinning media, cryptographic erase on SED HDDs or multi-pass overwrite remains standard; degaussing is ineffective on SSDs and only partially effective on modern high-coercivity HDDs — physical destruction is the fallback [secondary].
- **NIST SP 800-88 Rev. 1 levels:** Clear (logical removal, e.g. single-pass overwrite — data not recoverable by standard utilities), Purge (firmware-level, e.g. NVMe sanitize block erase or crypto erase — not recoverable even in a lab), Destroy (disassembly, shredding, incineration). Match the level to data classification; regulated data generally requires Purge or Destroy [secondary].
- **Sanitize progress and failure:** the sanitize-log's SSTAT field reports status; SPROG 65535 = complete. If sanitize fails partway, the controller may leave the device in a degraded state — verify, don't assume; retry or escalate to physical destruction [secondary].
- **Overwrite details:** one full-device pass of a fixed pattern is sufficient for modern media per NIST 800-88 (multi-pass DoD 5220.22-M is legacy for modern densities); the `no-deallocate-after-sanitize` behavior varies — check whether deallocated blocks read back as zeros [secondary].
- **SATA Secure Erase via hdparm:** `--security-set-pass` then `--security-erase` (normal) or `--security-erase-enhanced` (also erases reallocated/reserved areas on drives that implement it). BIOS "frozen" security state blocks this — suspend/resume (S3) usually unfreezes without data loss on modern systems [secondary].
- **TCG revert:** PSID revert (physical label code) cryptographically resets an Opal/Enterprise drive to factory state, destroying all keys and data — the fastest compliant decommission for SED fleets, but it requires physical access and the printed PSID [secondary].
- **Verification is part of the procedure:** sample read-back after sanitize (first/middle/last LBAs plus random offsets) and log the results per serial number; auditors accept logs, not assertions [secondary].
- **Time budgeting:** crypto erase ≈ seconds; block erase ≈ minutes; full overwrite ≈ hours on multi-TB drives and adds a full P/E cycle of wear — schedule accordingly and prefer crypto erase wherever SED is confirmed [secondary].
- **Regulatory mapping:** GDPR right-to-erasure, HIPAA, PCI-DSS, and government classifications each define acceptable methods — Purge-level sanitize with per-device logs satisfies most; "we deleted the files" satisfies none [secondary].
- **NVMe sanitize procedure (typical flow):** 1) `nvme id-ctrl /dev/nvme0 | grep -i sanicap` — confirm sanitize supported; 2) back up anything needed — sanitize is whole-controller and irreversible; 3) `nvme sanitize /dev/nvme0 --sanact=4` (crypto erase on SED) or `--sanact=2` (block erase); 4) poll `nvme sanitize-log /dev/nvme0` until SSTAT shows complete (SPROG 65535); 5) record serial + SSTAT in the decommission log [secondary].
- **SATA secure erase procedure:** 1) check not frozen: `hdparm -I /dev/sdX | grep frozen`; 2) if frozen, suspend/resume (S3) to unfreeze; 3) `hdparm --user-master u --security-set-pass PASS /dev/sdX`; 4) `hdparm --user-master u --security-erase-enhanced PASS /dev/sdX` (or `--security-erase`); 5) confirm with `hdparm -I` (not locked, not frozen) and sample read-back [secondary].
- **PSID revert procedure (TCG SED):** 1) record the PSID from the drive label photographically; 2) boot a trusted environment; 3) run the vendor/TCG tool's PSID-revert (e.g. `sedutil-cli --PSIDrevert <PSID> /dev/nvme0`); 4) verify the drive is factory-fresh (no locking ranges); 5) log per serial [secondary].
- **Post-sanitize verification sampling:** read first, middle, and last 1 MB plus ≥10 random offsets across the LBA range; on crypto erase, confirm reads return zeros or crypto-erase patterns per the sanitize spec; any residual structured data = failed sanitize → escalate to destruction [secondary].
- **IEEE 2883-2022:** the current standard for sanitizing storage media, superseding older DoD 5220.22-M guidance for modern devices — reference it in policy instead of legacy multi-pass standards [secondary].
- **Sanitize vs format decision tree:** decommissioning the whole drive → sanitize; repurposing a namespace within a live drive → format with ses; SED confirmed → crypto erase (fastest Purge); non-SED SSD → block erase; non-SED with unknown encryption → block erase + verification [secondary].
- **RAID member sanitization:** break the array first and sanitize member disks individually — controller-level "erase VD" may not reach over-provisioned areas on all members [secondary].
- **Failed drives:** a drive that won't respond to sanitize commands can't be trusted to have erased itself — physical destruction is the only compliant path for dead drives holding sensitive data [secondary].
- **Cloud/leased hardware:** you cannot sanitize what you can't touch — contractual data-destruction attestation from the provider plus encryption-before-write (so the provider never holds plaintext) is the control [secondary].
- **Enhanced vs normal erase (SATA):** Enhanced Secure Erase additionally wipes reallocated and reserved areas on drives that implement it; normal erase may leave remapped sectors' stale data — prefer enhanced when available [secondary].
- **`blkdiscard` semantics:** plain `blkdiscard` issues TRIM/deallocate (fast, not a sanitize); `--secure` requests secure-trim only where the drive reports deterministic TRIM + read-zero-after-trim. Neither equals NVMe sanitize [secondary].
- **Documentation to keep:** sanitize method, tool + version, operator, timestamp, serial, result code, and verification sample results — the audit packet for ISO 27001 A.8.10 / NIST 800-88 [secondary].
- **Resale of sanitized drives:** only resell after Purge-level sanitize with logs; Clear-level is insufficient for drives that held sensitive data [secondary].
- **Sanitize of failed SED:** if the drive is dead, the key is (presumably) unrecoverable — but "presumably" isn't Purge. Dead SED with sensitive data → physical destruction, same as non-SED [secondary].
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

