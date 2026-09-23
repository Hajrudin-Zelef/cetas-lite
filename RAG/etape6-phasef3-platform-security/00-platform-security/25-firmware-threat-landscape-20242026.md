---
id: etape6-phasef3-platform-security/00-platform-security/25-firmware-threat-landscape-20242026
title: "25. Firmware threat landscape 2024–2026"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Google", "Intel", "Meta", "Qualcomm"]
dates: ["2026-09"]
keywords: ["agent", "amd", "aws", "consumer", "cost", "cybersecurity", "distribution", "intel", "licenses", "luna", "pricing", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [301, 355]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 9a5ae60612beb81883fed4837307b5c5b5b9a33750cd05ff0135137f9411403f
---

# 25. Firmware threat landscape 2024–2026

## 25. Firmware threat landscape 2024–2026

- **CosmicStrand** (2022): UEFI firmware rootkit found on ASUS/GIGABYTE boards; persists in SPI flash, injects into kernel at boot [secondary].
- **LogoFAIL** (2023, Binarly): vulnerabilities in UEFI image parsers (BMP/PNG logos) across AMI/Insyde/Phoenix-derived firmware — LogoFAIL affected virtually all consumer and enterprise endpoints [secondary].
- **PixieFail** (2023, Quarkslab): nine CVEs (CVE-2023-45229–CVE-2023-45237) in EDK2's IPv6 network stack — remote code execution via PXE before OS boot [secondary].
- **UEFIcanhazbufferoverflow** (2024, Binarly/Eclypsium): Phoenix SecureCore `GetVariable` buffer overflow; affected Intel CPUs from 6th-gen Skylake through 14th-gen Raptor Lake [secondary](https://www.tomshardware.com/pc-components/cpus/firmware-flaw-affects-numerous-generations-of-intel-cpus-uefi-code-execution-vulnerability-found-for-intel-cpus-from-14th-gen-raptor-lake-to-6th-gen-skylake-cpus?rand=12330).
- **CVE-2024-54085** (MegaRAC BMC): unauthenticated remote access to BMCs — fleet-wide server compromise vector [secondary](https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf).
- **September 2026 UEFI Shell bypasses:** AMI CVE-2026-33197, Cisco UCS CVE-2026-20293, Insyde CVE-2026-6485 — no confirmed exploitation or public PoC per the reporting source [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
- **Insyde CVE-2025-4275 (SecureFlashCertData):** digital-certificate injection enabling Secure Boot bypass/rootkits; patched via INSYDE-SA-2025002 [secondary](https://cyberpress.org/uefi-vulnerability-in-insyde-allows-digital-certificate-injection/).
- Eclypsium's 2024 report documented outdated, unpatched UEFI firmware persisting in enterprise fleets years after vendor patches — the patch-gap problem [secondary].
- Trend 2026: firmware attacks move from proof-of-concept to commercialized tooling; supply-chain concentration (AMI/Insyde/Phoenix) means one flaw = hundreds of models [secondary].

## 26. Open firmware ecosystem 2026

- **coreboot:** open-source firmware platform replacing proprietary UEFI on supported boards; payloads include SeaBIOS, Tianocore (EDK2), LinuxBoot [secondary](https://forums.servethehome.com/index.php?threads/a-hardware-enthusiast-view-on-the-usefulness-of-open-source-firmwares-like-coreboot.27027/).
- Commercial coreboot vendors: System76, Purism, Star Labs (laptops/desktops); 3mdeb's Dasharo firmware (coreboot + EDK2 distribution) for business/ODM devices [secondary].
- **LinuxBoot:** replaces most of DXE with a Linux kernel as the boot payload; used by hyperscalers (Meta, Google) for fast, auditable boot; 2026 broader adoption unconfirmed (G2) [unverified].
- **OpenBMC:** Linux-Foundation BMC firmware stack (Yocto-based); CanopyBMC 2026.06 (Yocto 6.0, per-platform SBOM, HPE ProLiant Gen11 support, UBM/NVMe inventory) is a leading distribution [independent](https://github.com/canopybmc/canopybmc/blob/HEAD/release-notes/2026.06.md).
- AMI MegaRAC SP-X is AMI's OpenBMC-based BMC offering [secondary](https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/).
- OpenBMC security features: signed firmware updates, Redfish/IPMI/SSH interfaces, hardware RoT integration points [secondary](https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/).
- **EDK2/TianoCore:** Intel's open UEFI reference implementation; the base from which AMI/Insyde/Phoenix derive — open for audit but rarely the shipped binary [secondary].
- **oreboot:** Rust rewrite of coreboot concepts; research-grade in 2026 [secondary].
- Open firmware advantages for security: reproducible builds, auditable source, faster patching, no opaque binary blobs in the boot chain (except required silicon FSP/ME/PSP blobs) [secondary].
- Limitation: Intel FSP (Firmware Support Package), ME, and AMD AGESA/PSP remain closed blobs even in coreboot deployments — the "binary blob boundary" [secondary].

## 27. TPM 2.0 operational detail

- **Endorsement Key (EK):** unique per-TPM RSA/ECC key pair burned at manufacture; EK certificate signed by the TPM vendor vouches for the TPM's authenticity — the anchor for attestation [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **Attestation Key (AK):** derived from the EK; signs PCR quotes (TPM2_Quote) presented to remote verifiers [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **Storage Root Key (SRK):** protects keys stored outside the TPM (key blobs sealed to the TPM) [secondary].
- **Key hierarchy:** EK → AK for attestation; SRK → user/storage keys; owner hierarchy with authorization policies (TPM2_PolicyAuthorize etc.) [secondary].
- **PCR banks:** 24 PCRs × SHA-1 and SHA-256 banks; PCRs 0–7 firmware/boot, 8–15 OS, 16–23 flexible/debug [secondary].
- **fTPM vs dTPM:** firmware TPM runs in PSP/TrustZone (AMD fTPM, Intel PTT, Qualcomm QTEE); discrete TPM is a separate chip (Infineon, Nuvoton, ST); vTPM is hypervisor-emulated for VMs [secondary].
- **Windows:** TPM 2.0 is a hard requirement for Windows 11; BitLocker sealing, Windows Hello, Credential Guard, and Secured-Core depend on it [secondary].
- **Linux:** tpm2-tools, tpm2-abrmd resource manager, kernel TPM driver; IMA (Integrity Measurement Architecture) extends file measurements into PCRs [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **Remote attestation flow:** challenger sends nonce → agent collects PCR quote + event log → AK-signed quote returned → verifier checks signature, EK cert, PCR values against golden measurements, event-log replay [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **Known TPM issues:** Infineon RSA key-generation flaw (ROCA, CVE-2017-15361, historical); fTPM stuttering bugs on AMD platforms (2022–2023, fixed via AGESA) [secondary].
- **TPM 2.0 PQC commands (v1.85):** ML-DSA/ML-KEM key generation, signing, KEM operations; firmware-update use case is PQC-protected on SLB 9672/9673 and ST33K Gen2 before attestation use [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
- wolfTPM (wolfSSL) tracks PQC: ChangeLog documents LMS/XMSS signature verification support for TPM firmware manifests [independent](https://github.com/wolfssl/microchip/blob/HEAD/wolfboot-2.9.0-commercial/lib/wolfTPM/ChangeLog.md).
- **xnode-tpm-attest:** Nix-based 7-step TPM2 remote attestation self-test (open tooling) [independent](https://github.com/johnforfar/xnode-tpm-attest).

## 28. HSM deployment patterns and selection guidance

- **Form factors:** network appliance (Luna Network HSM, nShield Connect), PCIe card (Luna PCIe, nShield Solo), USB (YubiHSM 2), cloud dedicated (AWS CloudHSM), cloud multi-tenant service (Luna as a Service, Azure Dedicated HSM) [secondary].
- **Tenancy models:** single-tenant dedicated (CloudHSM, Dedicated HSM) vs multi-tenant with partitions (Luna: up to 100 partitions per appliance, each with own SO/users/keys) [secondary].
- **APIs:** PKCS#11 (universal), JCE (Java), CNG/KSP (Windows), OpenSSL ENGINE/provider, REST (CloudHSM, Luna SA), vendor SDKs [secondary].
- **M-of-N quorum:** key-management operations (backup, partition init) require M of N smart-card/PED tokens — split-knowledge control [secondary].
- **Key backup:** encrypted cloning domains (Luna), smart-card sets (nShield Security World), CloudHSM backup to S3 [secondary].
- **Use-case mapping:** public CA roots → network HSM (FIPS 140-3 L3, CC EAL4+); database TDE → PCIe or network; code signing → network + timestamping; blockchain/digital assets → FIPS L3 with custom CodeSafe apps; dev/test → YubiHSM 2 or cloud [secondary].
- **Certification as procurement gate:** FIPS 140-3 Level 3 is the standard HSM bar; Common Criteria EAL4+ adds assurance; U.S. federal may require U.S.-manufactured (Thales TCT line) [vendor-reported](https://www.thalestct.com/wp-content/uploads/2022/09/luna-network-hsm-pb-11-29-23.pdf).
- **PQC posture (2026):** Thales TCT Luna T-Series v7.15.1 (cert 5450) — ML-DSA/ML-KEM/LMS validated; Luna 8 — new PQC-oriented platform; nShield — PQC support advertised; verify per-firmware-version validation status before procurement [vendor-reported][secondary].
- **Price visibility (2026, indicative):** nShield Connect aggregator listing "starting at $10,000" (callback pricing); Port of Oakland nShield delivery reported >$19K per unit (secondary snapshot, not OEM list); exact list prices unpublished — G7 [secondary](https://www.techjockey.com/detail/entrust-nshield-connect) [secondary](https://www.saitechincorporated.com/saitech-inc-strengthens-port-of-oaklands-cybersecurity-with-entrust-nshield-hsm-solutions/).
- **Hidden costs:** per-partition licensing, client licenses, support contracts, PED/smart-card kits, HA clustering (second appliance), training [secondary].
- **Cloud vs on-prem trade-off:** CloudHSM = no hardware to manage, single-tenant, but cloud-region jurisdiction and hourly cost; on-prem = capital cost + operational burden, full physical control [secondary](https://technology.toolsinfo.com/compare/entrust-nshield-hsm-vs-aws-cloudhsm).

