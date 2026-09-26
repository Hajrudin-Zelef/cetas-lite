---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d4-secure-erase-and-sanitize-retiring-drives-without-leaking
title: "D4 — Secure erase and sanitize: retiring drives without leaking data"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: ["containment", "research"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [81, 108]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 45628df1786dae00f982e7f973d49ca16b239d4f40a6589ef4a86723b4f92cc7
---

# D4 — Secure erase and sanitize: retiring drives without leaking data

- **NIST SP 800-193 "Platform Firmware Resiliency Guidelines"** defines three principles for platform firmware: Protection (integrity of firmware code and critical data, authentic updates), Detection (detect corruption of firmware/critical data), Recovery (restore to a state of integrity after detection) [official].
- **The guideline explicitly covers storage:** the draft lists the mass-storage Host Controller and the HDD/SSD itself (microcontroller + firmware) as devices whose firmware compromise "can be used as a launch pad for other exploits in the system, or could be used to compromise user and/or platform data" [official].
- **Three platform properties:** Protected (meets protection requirements, may not fully recover), Recoverable (detect + recover), Resilient (all requirements; a compromised device must not impact platform security) [official].
- **Storage-vendor implementations observed in research:** Broadcom 9600 adapters ship hardware Secure Boot plus SPDM (Security Protocol and Data Model) attestation [secondary]; Microchip Adaptec SmartRAID 4300 lists hardware root of trust, secure boot, secure update, attestation, and SED support [vendor-reported]; HPE SSDs advertise Digitally Signed Firmware to block firmware-based attacks [secondary].
- **Practical firmware hygiene for storage admins:** keep drive and controller firmware on vendor-validated releases (not latest-by-default); verify signed-update chains where available; quarantine drives with unknown firmware provenance; log firmware versions in the asset inventory so a vulnerability (e.g. a 2026 SSD firmware CVE) can be triaged fleet-wide [secondary].
- **Firmware bricking during power loss** is a real failure mode Kingston calls out: if a drive bricks during unsafe-power-loss qualification, engineering restarts the whole qualification — a reminder that PLP and firmware robustness are co-dependent [official].
- **Gap:** no public consolidated database was found of SSD/HBA firmware CVEs for 2025–2026 with severity and affected families; this is a genuine research gap for fleet risk assessment [unverified].
- **The three Roots of Trust (800-193 §4):** RTU — Root of Trust for Update (authenticates firmware updates and critical-data changes, enforces anti-rollback); RTD — Root of Trust for Detection (measures/detects corruption of code and critical data); RTRec — Root of Trust for Recovery (restores code and critical data to integrity). A storage device claiming resilience should be able to name which RoT covers its firmware update path [official].
- **External PFR implementations:** FPGAs (Lattice, Microchip) are commonly used as board-level PFR roots of trust that monitor the BMC and UEFI SPI flash independently of the host CPU — the same pattern applies to storage controller firmware: an external RoT can verify the RAID card's firmware before the host boots [secondary].
- **SPDM (DMTF Security Protocol and Data Model):** the attestation protocol Broadcom cites for 9600 adapters; lets the host challenge a device to prove its firmware measurements before trusting it — relevant to supply-chain attacks where a counterfeit or tampered HBA/SSD is inserted [secondary].
- **NVMe firmware update mechanics:** up to 7 firmware slots, `fw-download` + `fw-commit` with commit actions (downloaded, downloaded+activate on reset, activate immediately); dual-bank layouts allow safe rollback if the new image fails to boot [secondary].
- **SED firmware as attack surface:** self-encrypting drive firmware handles the media encryption key; a compromised SED firmware can exfiltrate keys — which is why signed firmware (HPE's digitally-signed SSD firmware) and TCG's firmware-update authentication matter beyond "just" boot integrity [secondary].
- **SBOM for firmware:** NIST and OMB guidance increasingly expect software bills of materials for firmware blobs; for storage fleets this means knowing which drives run which firmware build when a CVE drops — the inventory discipline from D3's hygiene bullet [secondary].
- **Gap:** no SSD or RAID vendor found at cutoff publishes an explicit NIST SP 800-193 conformance claim for drive firmware; "secure boot" marketing does not equal 800-193 Recoverable/Resilient properties [unverified].
- **800-193 properties mapped to storage:** Protected ≈ signed firmware updates with anti-rollback; Recoverable ≈ corruption detection + golden-image restore; Resilient ≈ a compromised SSD/HBA cannot undermine host boot or peer devices. Ask vendors which property their storage firmware claims — most can only honestly claim Protected [official].
- **Recovery image storage:** the golden/recovery image must live in storage the attacker can't rewrite (write-protected flash region, separate SPI device) — recovery from the same rewritable flash the attacker corrupted is not recovery [official].
- **Detection at runtime, not just boot:** 800-193's Detection includes runtime re-measurement; a drive that only verifies firmware at power-on is blind to runtime compromise [official].
- **TPM integration:** platform TPM can seal SED authentication secrets and record firmware measurements (PCRs) for remote attestation of the storage stack [secondary].
- **Supply-chain controls:** verify drive/HBA firmware hashes against vendor-published values on receipt; counterfeit storage devices with backdoored firmware are a documented attack vector [secondary].
- **Measured boot for storage controllers:** a platform that measures HBA/RAID firmware into TPM PCRs can detect a swapped or tampered card at boot — the server-side complement to the card's own secure boot [secondary].
- **Anti-rollback (version binding):** RTU's rollback protection prevents downgrading to a vulnerable firmware after patching — verify the mechanism exists before assuming "we patched" is permanent [official].
- **Update delivery:** vendor tools (Broadcom MSM/StorCLI, Microchip maxView/ARCCONF, `nvme fw-download`) plus OS package channels; air-gapped estates need an offline firmware bundle process — don't discover this during a CVE fire drill [secondary].
- **Firmware inventory format:** record vendor, model, serial, current firmware, and last-update date per device; this is the table you join against a CVE notice [secondary].
- **Pre-boot DMA protection:** a malicious storage device could abuse bus-master DMA; IOMMU/VT-d containment is the platform control complementing device attestation [secondary].
- **Event logging:** enable and forward controller event logs (thermal events, PD errors, firmware crashes) to the SIEM — storage firmware anomalies are security telemetry, not just ops telemetry [secondary].

## D4 — Secure erase and sanitize: retiring drives without leaking data

