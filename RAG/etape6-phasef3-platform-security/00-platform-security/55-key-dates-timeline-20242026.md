---
id: etape6-phasef3-platform-security/00-platform-security/55-key-dates-timeline-20242026
title: "55. Key dates timeline 2024–2026"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "EU", "Intel", "Microsoft"]
dates: ["2025-04-02", "2025-09-18", "2025-10-01", "2026-01-06", "2026-06-30", "2026-08-20", "2026-09-22", "2026-12-31"]
keywords: ["acquisition", "amd", "cyber", "intel", "luna", "memory", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [692, 752]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: ff3dc4ad1d87a3f8b4bd8a17f4529629a11c9a032f99f4d999474b2635051db6
---

# 55. Key dates timeline 2024–2026

## 55. Key dates timeline 2024–2026

- 2024-04: UEFIcanhazbufferoverflow disclosed (Phoenix SecureCore GetVariable; Intel Skylake→Raptor Lake) [secondary].
- 2024-10: MegaRAC BMC CVE-2024-54085 disclosed [secondary].
- 2025-01: Sectigo acquisition of Entrust's public TLS business announced [secondary].
- 2025-04-02: Intel discontinued SGX Attestation Service (EPID) [official].
- 2025-09-18: Sectigo/Entrust TLS deal completed [secondary].
- 2025-10-01: Intel archived sgx-tdx-dcap-quoteverificationservice [official].
- 2026-01-06: Kioxia BG7 (BiCS8, Opal 2.01) at CES 2026 [vendor-reported].
- 2026-01: CanopyBMC 2026.06 (Yocto 6.0, SBOM, Gen11) [independent].
- 2026-04: Kioxia EG7 QLC (Opal 2.02) [vendor-reported].
- 2026-05: Azure Intel TDX confidential VMs GA (DCesv6/DCedsv6, ECesv6/ECedsv6) [vendor-reported].
- 2026-06: Microsoft Windows Production PCA 2011 expiry per one source (conflict C1) [secondary].
- 2026-06-30: Azure DCsv2 (SGX) retired [independent].
- 2026-07: Supermicro Intel TDX + HGX B200 white paper [vendor-reported].
- 2026-08: Thales Luna 8 launched; AMD Caliptra integration announced for 2026+ products [secondary].
- 2026-08-20: Caliptra 2.1 first-silicon power-on (reported) [secondary].
- 2026-09: UEFI Shell Secure Boot bypass claims (AMI/Insyde/Cisco CVEs); DDRop research disclosed [secondary].
- 2026-10: Microsoft Windows Production PCA 2011 expiry per second source (conflict C1) [secondary].
- 2026-12-31 (expected): PCCS rebranded to Collateral Caching Service (CCS) [official].
- 2027: ANSSI stops certifying products without quantum-safe encryption [secondary].
- 2030/2031: CNSA 2.0 PQC deadlines (key establishment / signatures) [secondary].

## 56. Cross-references to other Phase F tracks

- NIC/DPU/SmartNIC platform security (secure boot of DPU ARM cores, NIC firmware signing): see Phase F1 file [unverified].
- FPGA bitstream authentication and tamper protection: see Phase F2 file [unverified].
- virtio/SR-IOV device isolation, VT-x/VT-d IOMMU, TDX Connect device assignment: see Phase F4 file [unverified].
- This file (F3) owns: system firmware, platform RoT, TPM, HSM, SED, confidential computing.

*End of Phase F3. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*

## 57. Secure Boot: servers vs clients

- Servers: Secure Boot typically enforced with OEM/vendor keys; custom key enrollment via BIOS setup or Redfish; headless verification through BMC [secondary].
- Clients: Microsoft keys pre-enrolled; Linux via shim/MOK; user can disable in firmware setup (reduces to audit value) [secondary].
- Virtualization: vTPM + UEFI Secure Boot in VMs (Hyper-V Generation 2, KVM OVMF) extends the chain into guests [secondary].
- Containers: no firmware boundary — container images rely on host Secure Boot + image signing (Sigstore/cosign), not UEFI [independent].
- Embedded/IoT: U-Boot verified boot (FIT signatures) plays the Secure Boot role where UEFI is absent [secondary].
- Supply-chain note: refurbished/secondary-market servers may carry fused Boot Guard or unknown PK — verify PK ownership before deployment [secondary].

## 58. Glossary additions

- **CRTM:** Core Root of Trust for Measurement [secondary].
- **SRTM/DRTM:** Static/Dynamic RTM [secondary].
- **IBB:** Initial Boot Block [secondary].
- **FSP:** Firmware Support Package (Intel) [secondary].
- **AGESA:** AMD Generic Encapsulated Software Architecture [secondary].
- **BMC:** Baseboard Management Controller [secondary].
- **PBA:** Pre-Boot Authentication [secondary].
- **OPAL SSC:** TCG Storage Security Subsystem Class [secondary].
- **RTMR:** Runtime Measurement Register (TDX) [vendor-reported].
- **HKID:** Host Key ID (TDX memory encryption) [secondary].
- **SEAM:** Secure Arbitration Mode (TDX) [vendor-reported].
- **TD:** Trust Domain (TDX) [vendor-reported].
- **HES:** Hardware Enforced Security (Arm CCA attestation root) [independent].
- **GPT:** Granule Protection Table (Arm RME) [independent].
- **LVFS:** Linux Vendor Firmware Service [independent].
- **CRA:** Cyber Resilience Act (EU) [secondary].

*End of Phase F3. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*
