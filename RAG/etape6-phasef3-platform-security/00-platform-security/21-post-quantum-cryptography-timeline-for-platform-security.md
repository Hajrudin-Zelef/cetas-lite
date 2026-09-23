---
id: etape6-phasef3-platform-security/00-platform-security/21-post-quantum-cryptography-timeline-for-platform-security
title: "21. Post-quantum cryptography timeline for platform security"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "CISA", "Falcon", "Intel", "Microsoft"]
dates: ["2023-05", "2026-06", "2026-08", "2026-10"]
keywords: ["amd", "aws", "intel", "luna", "memory", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [248, 300]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 091d04e363e2f352d08b9337f1208f0679b398c4026a7b2b561e87c549abcd80
---

# 21. Post-quantum cryptography timeline for platform security

## 21. Post-quantum cryptography timeline for platform security

- **NIST PQC standards:** FIPS 203 (ML-KEM), FIPS 204 (ML-DSA), FIPS 205 (SLH-DSA) finalized 2024; FIPS 206 (FN-DSA/Falcon) in progress [secondary](https://github.com/pqctoday-org/pqctoday-hub/blob/HEAD/src/data/product-extractions/csc_043_extractions_03312026.md).
- **NSA CNSA 2.0:** requires PQC for National Security Systems — key establishment by 2030, digital signatures by 2031 (per June 2026 U.S. executive order); Thales TCT Luna T-Series v7.15.1 is the first HSM with all CNSA 2.0 PQC algorithms (ML-DSA, ML-KEM, LMS) in a FIPS 140-3 validation [vendor-reported](https://www.executivebiz.com/articles/thales-tct-luna-t-series-hsm-fips-140-3-validation).
- **France/ANSSI:** stops certifying security products without quantum-safe encryption from 2027; government/critical-infrastructure procurement of quantum-safe-only products expected by 2030 [secondary](https://highways.today/2026/08/05/quantum-deadlines/).
- **TPM PQC:** TCG TPM 2.0 library v1.85 (PQC commands) final; no discrete PQC TPM silicon shipping as of August 2026; SEALSQ QS7001 engineering samples targeted Q3 2026 [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
- **Caliptra PQC:** 2.0/2.1 subsystems implement ML-DSA (attestation/secure boot) and ML-KEM (key wrapping) in hardware with side-channel countermeasures [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md).
- **HSM PQC:** Luna 8 launched on a new PQC-oriented platform (Aug 2026); nShield advertises PQC support; migration priority rated "Critical" (plan and execute by 2030/2035) for harvest-now-decrypt-later exposure [secondary](https://highways.today/2026/08/05/quantum-deadlines/) [secondary](https://technology.toolsinfo.com/compare/entrust-nshield-hsm-vs-aws-cloudhsm).
- **LMS/XMSS:** stateful hash-based signatures used for firmware-update authentication (Infineon SLB 9672/9673 update channel; ST ST33K Gen2 LMS-format manifests) — the pragmatic PQC step for code signing before full ML-DSA deployment [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md) [independent](https://github.com/wolfssl/microchip/blob/HEAD/wolfboot-2.9.0-commercial/lib/wolfTPM/ChangeLog.md).

## 22. Standards and specification index

- NIST SP 800-147 (BIOS Protection Guidelines), SP 800-155 (BIOS Integrity Measurement), SP 800-193 (Platform Firmware Resiliency), SP 800-53 (controls), SP 800-37 (Risk Management Framework) — the firmware compliance stack [secondary](https://malware.news/t/nist-compliance/74695).
- TCG: TPM 2.0 Library (v1.85 PQC), DICE, DPE, TCG Opal 2.0/2.01/2.02 + Opalite + Enterprise SSC + Ruby, TPM 2.0 Provisioning Guidance [independent][secondary].
- UEFI Forum: UEFI Specification, PI Specification; dbx updates distributed via Windows Update (KB5025885 for the 2023-2024 revocation wave) [secondary].
- DMTF: SPDM, MCTP, PLDM, Redfish (BMC management; OpenBMC implements Redfish + IPMI) [official][secondary](https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/).
- OASIS: KMIP (key management interoperability) [secondary].
- IEEE: 1667 (transient storage authentication), 1619 (storage encryption) [secondary].
- IETF: RATS architecture, EAT tokens [independent](https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md).
- FIPS 140-3 (ISO/IEC 19790:2012 basis) levels 1–4; Common Criteria EAL ratings (Luna: EAL4+; Thales MultiApp v5.2: ANSSI CC EAL6+) [secondary](https://github.com/pqctoday-org/pqctoday-hub/blob/HEAD/src/data/product-extractions/csc_043_extractions_03312026.md).

## 23. Glossary

- **RoT / RTU / RTM / RTD / RTRec / RTI:** Root of Trust; for Update, Measurement, Detection, Recovery, Identity [official].
- **CoT:** Chain of Trust anchored by a RoT [official].
- **PK / KEK / db / dbx:** Secure Boot key hierarchy (Platform Key, Key Exchange Key, signature database, forbidden database) [secondary].
- **fTPM / dTPM / vTPM:** firmware, discrete, virtual TPM [secondary].
- **DRTM / SRTM:** Dynamic / Static Root of Trust for Measurement [secondary].
- **CVM:** Confidential VM (TDX/SEV-SNP/CCA Realm) [secondary].
- **RME / RMM:** Realm Management Extension / Realm Management Monitor (Arm CCA) [independent].
- **EPC:** Enclave Page Cache (SGX) [secondary].
- **SMM / SMRAM:** System Management Mode and its protected memory [independent].
- **PBA:** Pre-Boot Authentication (SED unlock) [secondary].
- **SIE:** Sanitize Instant Erase (NVMe) [vendor-reported].
- **PFR:** Platform Firmware Resilience [vendor-reported].
- **PSP:** Platform Security Processor (AMD) [secondary].
- **ME:** Management Engine (Intel) [secondary].
- **ACM:** Authenticated Code Module (Intel TXT/Boot Guard) [secondary].
- **UDS / CDI:** Unique Device Secret / Compound Device Identifier (DICE) [independent].
- **HSM partitions:** cryptographically isolated key containers (Luna: up to 100 per appliance) [secondary].

## 24. BlackLotus anatomy and Secure Boot bypass mechanics

- BlackLotus (CVE-2023-24932) is a UEFI bootkit that bypasses Secure Boot on fully patched Windows 11 by exploiting a 2012-era vulnerability in the Windows boot manager [secondary](https://csirt.ncc.gov.ng/index.php/resources/security-advisories/66-blacklotus-uefi-bootkit-malware-targeted-fully-patched-windows-11-systems).
- Persistence mechanism: installs itself in the EFI System Partition (ESP) and modifies the Boot Configuration Data (BCD); executes before the OS, defeating OS-level security tools [secondary].
- Anti-forensics: disables HVCI (Hypervisor-Protected Code Integrity), Windows Defender, and BitLocker; deletes its own on-disk artifacts after establishing persistence [secondary].
- Microsoft's response required a three-stage fix: (1) trust-store update (add the 2023 certificates), (2) boot-manager replacement (revoke vulnerable 2011-signed binaries), (3) DBX revocation (block the vulnerable bootmgr) [secondary](https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/).
- The dbx revocation wave was staged through Windows Update (notably KB5025885, May 2023 onward) because applying revocations to systems with unupdated boot managers bricks boot [secondary].
- BlackLotus bootkits were sold on underground forums for ~$5,000 (2023) — industrialization of firmware attacks [secondary](https://www.makeuseof.com/why-windows-secure-boot-can-be-bypassed-so-easily/).
- Secure Boot bypass classes: stale dbx entries (revocation not applied), vulnerable-but-signed bootloaders, NVRAM/db manipulation via SMM or variable-service flaws, and Option ROM abuse [secondary](https://www.makeuseof.com/why-windows-secure-boot-can-be-bypassed-so-easily/).
- Microsoft's 2011 certificate expiry (June/October 2026 — conflict C1) forced the industry-wide 2023-certificate migration; systems that never received the 2023 trust store cannot boot post-revocation binaries [secondary](https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/).
- Windows Server lacked an automatic rollout path for the migration — manual admin action required, creating enterprise exposure [secondary](https://github.com/fleetdm/fleet/issues/45514).

