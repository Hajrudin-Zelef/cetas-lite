---
id: etape9-phased-data-protection-raid/00-data-protection-raid/part-6
title: "Step 9 — Phase D: Data Protection & RAID Hardware (part 6)"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["nand"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [109, 138]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 6dbbe997cf5e2344c2ec0917c601de7b9ee7048b73f69fb03078658a8d48b985
---

# Step 9 — Phase D: Data Protection & RAID Hardware (part 6)

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
