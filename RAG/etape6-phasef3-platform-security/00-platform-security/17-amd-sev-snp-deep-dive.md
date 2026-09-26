---
id: etape6-phasef3-platform-security/00-platform-security/17-amd-sev-snp-deep-dive
title: "17. AMD SEV-SNP deep dive"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Google", "Intel", "Microsoft"]
dates: ["2026-07-13", "2026-08"]
keywords: ["amd", "aws", "consumer", "cybersecurity", "exploit", "intel", "memory", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [201, 238]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 01e1c7e074a3435babd4c9e51eadaab58c0f4dd8c9e3f28ea4f2ae3cec28e583
---

# 17. AMD SEV-SNP deep dive

## 17. AMD SEV-SNP deep dive

- SEV lineage: SEV (memory encryption) → SEV-ES (encrypted state) → SEV-SNP (Secure Nested Paging: integrity protection against malicious hypervisor remapping) [secondary].
- SNP attestation: the AMD Secure Processor signs attestation reports covering VM measurements, TCB version, and guest-provided data; remote verifiers check the report against AMD's root keys [secondary].
- VMPL (VM Privilege Levels): four privilege levels inside the guest allowing in-guest security managers (e.g., a paravisor or vTPM) to run more privileged than the guest OS [secondary].
- Page-relocation feature (RMPUPDATE/PSMASH) is the mechanism DDRop abused on SNP — see §16 [secondary](https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say).
- Nested SNP: Hyper-V can virtualize SNP into a child partition (virtualized RMPUPDATE/PSMASH MSRs, ACPI PSP) — how AKS confidential containers work on Azure hosts (mshv + Cloud Hypervisor as L1); KVM nested-SNP-host series was an unmerged 2023 RFC; the mshv L1VH + SNP series dates to August 2026 and is still in review; no nested-TDX-host mode exists [independent](https://github.com/enclavehost/enclave/blob/HEAD/windows/EDITIONS.md).
- AMD fTPM runs on the PSP; SNP guests commonly use vTPMs backed by the host for measured boot inside CVMs [independent](https://github.com/johnforfar/xnode-tpm-attest).
- Cloud availability 2026: Azure DCasv5/v6 + ECasv5/v6, Google Cloud, AWS select instances, OVHcloud (France) [secondary](https://ayinedjimi-consultants.fr/static/pdf/confidential-computing-tee-2026.pdf).

## 18. UEFI internals: boot phases, NVRAM, SPI flash, SMM

- UEFI PI boot phases: SEC → PEI → DXE → BDS → TSL → RT; drivers execute in DXE before the OS loads — the window bootkits exploit [secondary].
- Platform firmware lives in SPI flash; the flash descriptor, ME region, BIOS region, and NVRAM variable store share the chip; weak SPI protections allow persistent implants (CosmicStrand, LoJax/MosaicRegressor families) [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- NVRAM variable store holds Setup, BootOrder, and Secure Boot key databases (PK/KEK/db/dbx); variable parsers with integer-overflow flaws enable pre-OS code execution via malformed variables [independent](https://github.com/itsventie/whitepapers/blob/HEAD/CyberSecurity/dfir/firmware/2026-07-13-uefi-spi-flash-nvram-carving.md).
- System Management Mode (SMM): the most privileged x86 execution mode; SMM isolation (locking SMRAM, SMM_BWP/SMM write protection) is a core CHIPSEC check and an SP 800-193 protection control [independent](https://github.com/houdini91/uefi-supply-chain/blob/HEAD/COMPLIANCE-MATRIX.md).
- UEFI variable services attacked: `GetVariable` (UEFIcanhazbufferoverflow), `SetVariable` (SecureFlashCertData injection, CVE-2025-4275) [secondary].
- SPI flash carving/DFIR: analysts carve NVRAM regions from flash dumps to find injected PE32 drivers and shellcode [independent](https://github.com/itsventie/whitepapers/blob/HEAD/CyberSecurity/dfir/firmware/2026-07-13-uefi-spi-flash-nvram-carving.md).
- Write protection mechanisms: BIOS_CNTL/BIOS lock, protected range registers (PR0–PR4), Intel Boot Guard (measured/verified boot from ACM), AMD Hardware Validated Boot [secondary].
- On consumer laptops since ~2019, Intel ME usually locks flashrom read access; on Dell iDRAC/HPE iLO servers, flashrom access via the BMC is often possible with root privileges and BIOS SPI-unlock settings [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- Intel Boot Guard profiles: measured boot (TPM extends measurements) vs verified boot (ACM verifies IBB signature, fused keys); fused Boot Guard cannot be disabled — a supply-chain consideration for refurbished hardware [secondary].
- AMD PSP performs analogous platform secure-boot verification on AMD systems [secondary].
- SMM callouts and SMI handlers have been a recurring CVE source; SMM isolation testing is part of the uefi-supply-chain compliance matrix [independent](https://github.com/houdini91/uefi-supply-chain/blob/HEAD/COMPLIANCE-MATRIX.md).

## 19. Measured boot, DICE, SPDM, and attestation standards

- **Measured boot:** each boot component hashes the next into TPM Platform Configuration Registers (PCRs); PCRs 0–7 cover firmware/Boot Guard/Secure Boot state; PCR 7 specifically reflects Secure Boot policy — BitLocker seals to PCR 7 so boot-chain changes block automatic decryption [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- **TPM 2.0 PCR banks:** SHA-1 (legacy) and SHA-256 banks; modern attestation uses the SHA-256 bank; TPM2_Quote signs PCR values with an Attestation Key (AK) [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **DICE (Device Identifier Composition Engine, TCG):** layered attestation architecture where each layer measures the next and derives a Compound Device Identifier (CDI); Caliptra implements the DICE Protection Environment (DPE) with UDS/CDI protection validated by NCC Group [independent](https://github.com/chipsalliance/Caliptra/raw/a8d574c3d8824f137b481136eadb5ae84d11636d/doc/NCC_Group_Microsoft_MSFT283_Report_2023-10-13_v1.2.pdf).
- **SPDM (Security Protocol and Data Model, DMTF):** authentication and measurement protocol for devices (PCIe, MCTP); used for component attestation — e.g., wolfTPM added SPDM secured transport for Nuvoton NPCT75x; Caliptra reference stack implements MCTP/PLDM/SPDM [independent](https://github.com/wolfssl/microchip/blob/HEAD/wolfboot-2.9.0-commercial/lib/wolfTPM/ChangeLog.md) [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md).
- **IETF RATS / EAT:** Remote ATtestation procedureS architecture; Entity Attestation Token (EAT) is the standard token format — Arm CCA uses EAT as its attestation wire format, enabling uniform verification pipelines across CCA/TDX/SEV-SNP [independent](https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md).
- **OCP security:** OCP's Security WG defines platform/peripheral security architecture recommendations; Caliptra's OCP Recovery (streaming boot) and OCP LOCK (fuse ratcheting, zeroization) are part of the 2.1 subsystem [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/Caliptra.md).
- **DRTM (Dynamic Root of Trust for Measurement):** late-launch measured environment (Intel TXT, AMD SKINIT); Secured-Core PCs combine DRTM with HVCI as the strongest commercial bootkit defense [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- **Intel TXT (Trusted Execution Technology):** DRTM via measured launch (tboot); complements static RTM [secondary].
- Attestation services 2026: Intel Tiber Trust Authority (SaaS, successor to EPID IAS), Azure Attestation, AWS Nitro attestation, open-source TrusTEE / Confidential Containers Trustee [official](https://www.intel.com/content/www/us/en/developer/archive/tools/sgx-attestation-service-utilizing-epid.html).

## 20. Server OEM firmware management and supply-chain security

