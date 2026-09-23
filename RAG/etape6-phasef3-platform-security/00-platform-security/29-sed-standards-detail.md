---
id: etape6-phasef3-platform-security/00-platform-security/29-sed-standards-detail
title: "29. SED standards detail"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "EU", "Intel", "Microsoft", "Samsung"]
dates: ["2026-07-13"]
keywords: ["amd", "cybersecurity", "datacenter", "incident", "intel", "nand"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [356, 401]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: b17cbd41d1f611ffb6621cf78f7941ed4bae8465e2c6bfb7cc778c105419a8af
---

# 29. SED standards detail

## 29. SED standards detail

- **TCG Storage SSCs:** Opal 2.0 (client), Opalite (lightweight client), Enterprise SSC (datacenter), Ruby (NVMe-specific), Key Per I/O (emerging) [secondary].
- **Locking ranges:** the drive's LBA space is divided into ranges, each with its own access credentials; global range + up to N user ranges [secondary].
- **MBR shadow / PBA:** pre-boot authentication environment that unlocks ranges before the OS loads; sedutil (Linux) implements PBA for TCG Opal drives [secondary].
- **PSID revert:** physical printed PSID allows cryptographic erase/revert-to-factory when credentials are lost — operational recovery path [secondary].
- **Crypto erase vs sanitize:** crypto erase destroys the media-encryption key (instant, O(1)); NVMe Sanitize (block erase/overwrite) physically sanitizes NAND; FIPS drives distinguish the two in certification [secondary].
- **eDrive (IEEE 1667 + TCG):** Microsoft's BitLocker hardware-encryption path using Opal drives; requires PBA or OS-managed unlock [secondary].
- **FIPS-validated drives:** Kioxia CM7 SED options (FIPS 140-3 L2); historically ~25% premium for FIPS 140-2 drives (2026 premium unconfirmed — G5) [vendor-reported][secondary](https://www.redeweb.com/en/present/Kioxia%27s-NVME-SSD-cryptographic-module-obtains-FIPS-140-3-Level-2-validation/).
- **KMIP integration:** enterprise SED fleets pair Opal drives with KMIP servers for centralized key escrow/lifecycle; 2026 KMIP server market not researched (G4) [unverified].
- **Risk — unmanaged SED keys:** drives shipped with default credentials (e.g., well-known MSID) provide zero protection; procurement must verify credential provisioning and PSID handling procedures [secondary].
- **Vendor coverage 2026:** Kioxia BG7 (Opal 2.01), EG7 (Opal 2.02), CM7 (Opal + FIPS 140-3 L2); Samsung/Micron/WD/Seagate ship Opal SED lines — exact 2026 SKUs not verified per model (see gap in §11) [vendor-reported].

## 30. Procurement checklist: platform security

- Require signed, versioned, rollback-protected firmware updates from the OEM (NIST SP 800-193 aligned) [official].
- Verify Secure Boot is enabled and the 2023 certificate set is installed; confirm dbx is current [secondary].
- Confirm TPM 2.0 present and enabled (dTPM preferred for servers; fTPM acceptable with documented risk acceptance) [secondary].
- For confidential-computing workloads: specify TDX or SEV-SNP capable SKUs and attestation-service compatibility [vendor-reported].
- For SED fleets: specify TCG Opal version, FIPS 140-3 requirement where needed, and credential-provisioning/PSID procedures in the contract [vendor-reported][secondary].
- For HSMs: require FIPS 140-3 Level 3 (current validation, not "designed to meet"), PQC roadmap in writing, and per-firmware-version validation scope [vendor-reported].
- Request firmware SBOMs from the OEM (EU CRA direction) [independent].
- Verify BMC firmware is signed and updatable independently of host firmware; change default BMC credentials [secondary].
- For PQC-sensitive lifetimes (>2030): require ML-DSA/ML-KEM support statements with target firmware versions [secondary].
- Confirm supply-chain provenance: U.S. federal may require domestic manufacture (e.g., Thales TCT) [vendor-reported].

## 31. Firmware update and recovery playbook

- Maintain a firmware inventory: BIOS/UEFI, BMC, NIC/DPU, drive, PSU versions per asset (CHIPSEC/firmware scanners, LVFS reports) [secondary].
- Stage updates: lab validation → canary group → fleet rollout; verify Secure Boot still passes and PCR 7 values update as expected [secondary].
- Apply dbx updates only after the 2023 trust store and updated boot managers are in place (brick risk otherwise) [secondary].
- Incident — suspected firmware compromise: dump SPI flash (flashrom via BMC or hardware programmer), carve NVRAM for injected drivers, compare against vendor golden image [independent](https://github.com/itsventie/whitepapers/blob/HEAD/CyberSecurity/dfir/firmware/2026-07-13-uefi-spi-flash-nvram-carving.md).
- Recovery: SP 800-193 recovery requires a protected recovery image + manual/automated reflash; Intel PFR automates this via FPGA RoT; Caliptra RTU/OOB recovery extends it to device RoTs [official].
- SED incident: crypto-erase via PSID revert or KMIP-driven key destruction for instant sanitization [secondary].
- HSM incident: M-of-N quorum procedures for key zeroization/rotation; verify backup integrity [secondary].

## 32. Measured-boot and attestation flow (reference)

- SRTM: CRTM → BIOS → Option ROMs → bootloader → OS, each measuring into PCRs 0–7 [secondary].
- DRTM: late launch (Intel TXT/tboot, AMD SKINIT) establishes a measured environment regardless of prior boot state [secondary].
- TPM2_Quote: AK signs PCR set + nonce; verifier checks EK/AK cert chain, PCR golden values, event-log replay [independent](https://github.com/johnforfar/xnode-tpm-attest).
- DICE layering: UDS → CDI per layer; Caliptra DPE protects CDI from firmware [independent].
- Device attestation: SPDM GET_MEASUREMENTS over MCTP/PCIe; Caliptra streams evidence [official].
- Confidential VM attestation: TDX quote (TDX Module) / SEV-SNP attestation report (AMD SP) / CCA realm token (EAT/RATS); verified by Intel Tiber Trust Authority, Azure Attestation, or open TrusTEE [official][independent].
- Seal/unseal: keys bound to PCR policy (TPM2_PolicyPCR) — BitLocker, LUKS, or CVM disk encryption keys release only in known-good boot state [secondary].

