---
id: etape6-phasef3-platform-security/00-platform-security/35-linux-secure-boot-shim-mok-distro-signing
title: "35. Linux Secure Boot: shim, MOK, distro signing"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "Apple", "EU", "Google", "Intel", "Microsoft", "Nvidia", "Qualcomm"]
dates: []
keywords: ["amd", "cost", "custom silicon", "cybersecurity", "embedding", "intel", "luna", "memory", "nvidia"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [483, 527]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 20c0df371957c27bf0cc2cde0538025bce235640e28d6d79ce52e65119f61676
---

# 35. Linux Secure Boot: shim, MOK, distro signing

## 35. Linux Secure Boot: shim, MOK, distro signing

- Most Linux distros boot on Secure Boot PCs via **shim**: a small Microsoft-signed first-stage bootloader embedding the distro's CA [secondary].
- **MOK (Machine Owner Key):** lets machine owners enroll their own keys (e.g., for DKMS kernel modules, NVIDIA drivers) without disabling Secure Boot [secondary].
- Distro signing: Fedora/RHEL (Red Hat CA), Ubuntu (Canonical CA), SUSE — each maintains its own signing infrastructure chaining to shim [secondary].
- The 2026 certificate transition required distros to re-issue shims signed against the 2023 Microsoft certs; stale shims risk revocation via SBAT [secondary].
- **SBAT (Secure Boot Advanced Targeting):** UEFI variable allowing revocation of vulnerable shims/GRUB versions without full dbx churn — used in the 2024 GRUB2 revocation wave [secondary].
- Kernel lockdown mode: when Secure Boot is on, the kernel restricts features that would allow unsigned code (e.g., /dev/mem, kexec of unsigned kernels) [secondary].

## 36. Windows Secured-Core and System Guard

- **Secured-core PCs** (Microsoft OEM program): require TPM 2.0, Secure Boot, HVCI (memory integrity), VBS, DRTM where available, and System Guard Secure Launch [vendor-reported].
- **HVCI (Hypervisor-Protected Code Integrity) / Memory Integrity:** uses VBS to isolate code-integrity decisions; BlackLotus deliberately disabled it [secondary].
- **Windows Defender System Guard:** runtime attestation of system integrity via the TPM; Secure Launch (DRTM) establishes a trusted launch even if firmware is untrusted [vendor-reported].
- **BitLocker:** seals disk-encryption keys to PCRs 0, 2, 4, 11 by default (PCR 7 with Secure Boot); firmware/boot changes trigger recovery-key prompts — the user-visible cost of measured boot [secondary].
- **Pluton:** Microsoft-designed security processor in AMD Ryzen 6000+, Intel 13th-gen+ (select), Qualcomm Snapdragon 8cx Gen 3+ devices; functions as TPM 2.0, protects credentials/keys/identities, receives firmware updates via Windows Update [vendor-reported].
- Pluton vs discrete TPM: integrated in the SoC die (no bus to sniff), firmware-updatable through the OS — but increases dependence on the CPU vendor + Microsoft update cadence [independent].

## 37. Apple and Google platform RoTs

- **Apple:** T2 chip (Intel Macs) then Secure Enclave in Apple Silicon; signed boot chain from Boot ROM; Startup Security Utility policies on Intel Macs [vendor-reported].
- Apple Silicon: Boot ROM → LLB → iBoot, each stage verified; Secure Enclave handles keys/biometrics; no user-disables without recoveryOS [vendor-reported].
- **Google Titan:** Titan M2 (Pixel) and Titan C (Chromebook) security chips; OpenSK open-source security-key firmware; Titan in Google Cloud for platform RoT [vendor-reported].
- Google's approach: custom silicon RoT + open firmware components + transparency logs (binary transparency) [independent].
- These closed-ecosystem RoTs contrast with the open PC-server model (UEFI + TPM + Caliptra) — relevant when comparing procurement options [independent].

## 38. FIPS 140-3 and Common Criteria programs

- **FIPS 140-3** (ISO/IEC 19790:2012) replaced FIPS 140-2; CMVP (Cryptographic Module Validation Program, NIST/CSE Canada) issues certificates; ESV (Entropy Source Validation) is now mandatory [official].
- Validation levels 1–4 as in §10; HSMs target Level 3; drive crypto modules (Kioxia CM7) validated at Level 2 [vendor-reported].
- Certificate lifecycle: validations are version-specific (firmware/hardware version bound); a firmware update can require re-validation or a "version update" submission [official].
- **Common Criteria:** ISO/IEC 15408; Protection Profiles define requirements; EAL ratings 1–7; Luna HSMs carry EAL4+; Thales MultiApp v5.2 smart-card platform reached ANSSI CC EAL6+ [secondary].
- **EUCC (EU Cybersecurity Certification Scheme on CC):** EU's successor framework under the Cybersecurity Act; relevant for EU public procurement of HSMs/firmware [secondary].
- Practical note: always verify the certificate number (e.g., Luna T7 cert 5450) on the NIST CMVP / CC portal rather than trusting marketing claims [official].

## 39. Confidential-containers and enclave frameworks

- **Confidential Containers (CoCo):** CNCF project running Kata Containers inside TDX/SEV-SNP CVMs; attestation via Trustee/verifier [independent].
- **Kata Containers:** VM-isolated containers — the substrate CoCo builds on [independent].
- **Gramine / Occlum:** SGX library OSes running unmodified Linux binaries in enclaves; relevant while SGX persists on Xeon [independent].
- **Enarx:** runtime for TEE-agnostic workloads (SGX + SEV-SNP backends) [independent].
- **Microsoft Azure confidential containers:** AKS confidential containers on SEV-SNP via Hyper-V nested virtualization (mshv + Cloud Hypervisor L1) [independent](https://github.com/enclavehost/enclave/blob/HEAD/windows/EDITIONS.md).
- Attestation plumbing: Trustee (CoCo), Intel Tiber Trust Authority, Azure Attestation — verifiers for CVM/enclave evidence [official][independent].
- Programming-model guidance: new greenfield AI/data workloads → CVMs (TDX/SEV-SNP); legacy enclave apps on Xeon → maintain SGX path; client SGX → migrate off (removed from client silicon) [independent].

