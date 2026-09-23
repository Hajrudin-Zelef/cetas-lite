---
id: etape6-phasef3-platform-security/00-platform-security/46-caliptra-subsystem-integration-notes
title: "46. Caliptra subsystem integration notes"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "CISA", "EU", "Google", "Intel", "Meta", "Microsoft", "Nvidia"]
dates: ["2025-04", "2026-06", "2026-08", "2026-09-22", "2026-10"]
keywords: ["amd", "cybersecurity", "intel", "nvidia", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [581, 635]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 2a4c715b0fa34988ab12b120a1fe2c9d7cdb7f25f1dd075cbd66254841fb894c
---

# 46. Caliptra subsystem integration notes

## 46. Caliptra subsystem integration notes

- Caliptra 2.0 core: RISC-V microcontroller + cryptographic accelerators (SHA-256/512, ECC-384, HMAC, AES, ML-DSA, ML-KEM), DICE/DPE identity, RTM/RTI services, MCTP/SPDM/PLDM transport [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/Caliptra.md).
- 2.1 adds: RTU (update), recovery (OCP Recovery), OCP LOCK (fuse ratcheting/zeroization), enhanced DPE commands — first-silicon power-on reported ~20 August 2026 [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md) [secondary](https://lists.chipsalliance.org/g/caliptra-wg/topics?threadid=97432215).
- Microsoft's DPE/iRoT audit by NCC Group (Oct 2023, report MSFT283): 5 consultants, 60 person-days; 5 findings (1 high — CDI/UDS protection); Caliptra claims audit-driven hardening [independent](https://github.com/chipsalliance/Caliptra/raw/a8d574c3d8824f137b481136eadb5ae84d11636d/doc/NCC_Group_Microsoft_MSFT283_Report_2023-10-13_v1.2.pdf).
- AMD announced Caliptra integration in 2026+ products — first major x86 vendor commitment [secondary](https://www.phoronix.com/news/AMD-Project-Caliptra-2026).
- NVIDIA, Microsoft, Google, Meta, Arm are Caliptra consortium participants — cross-industry RoT standardization [official].
- Subsystem FIPS certification: still TBD on the roadmap — procurement blocker for regulated buyers until resolved [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md).
- Open-source RTL + firmware enables independent audit — contrasts with closed PSP/ME/Pluton firmware [official].

## 47. BMC security hardening

- Change default BMC credentials (ADMIN/admin) on first boot; disable IPMI anonymous/NULL cipher suites [secondary].
- Segment the management LAN; never expose BMC web/IPMI to the internet or general user VLANs [secondary].
- Enable signed BMC firmware updates; verify update signatures out-of-band where supported [secondary].
- Monitor BMC audit logs; MegaRAC CVE-2024-54085 showed unauthenticated BMC access enables persistent, OS-invisible control [secondary](https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf).
- Prefer OpenBMC-based BMCs (auditable source, reproducible builds) where the OEM offers them [independent](https://github.com/canopybmc/canopybmc/blob/HEAD/release-notes/2026.06.md).
- BMC supply chain: AMI MegaRAC dominates commercial BMC firmware; CanopyBMC, 3mdeb, and ODMs provide OpenBMC alternatives [secondary](https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/).

## 48. Secure Boot certificate transition operations (2026)

- Trust store: enroll the 2023 Microsoft certs (Windows UEFI CA 2023, Microsoft UEFI CA 2023) alongside 2011 certs before any revocation [secondary](https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/).
- Boot manager: deploy the 2023-signed bootmgr/bootmgfw before dbx revokes 2011-signed binaries [secondary].
- Revocation: apply the updated dbx (via Windows Update KBs on Windows; manual on Linux fleets) only after steps 1–2 are confirmed fleet-wide [secondary].
- Windows Server: no automatic rollout — script and verify manually [secondary](https://github.com/fleetdm/fleet/issues/45514).
- Linux fleets: re-issue shims against 2023 certs; manage SBAT levels to avoid booting revoked GRUB [secondary].
- Rollback plan: keep 2011-signed recovery media until the fleet is fully migrated; test bare-metal recovery paths [secondary].
- Expiry conflict C1 (June vs October 2026) means: treat June 2026 as the planning deadline regardless [unverified].

## 49. Glossary extension

- **SBAT:** Secure Boot Advanced Targeting (shim/GRUB revocation) [secondary].
- **DBX:** forbidden signature database [secondary].
- **MOK:** Machine Owner Key [secondary].
- **RTU:** RoT for Update [official].
- **DPE:** DICE Protection Environment [independent].
- **RATS:** Remote ATtestation procedureS (IETF) [independent].
- **EAT:** Entity Attestation Token [independent].
- **SPDM/MCTP/PLDM:** device security/management protocols (DMTF) [official].
- **CC:** Common Criteria (ISO/IEC 15408) [secondary].
- **EUCC:** EU Cybersecurity Certification scheme [secondary].
- **CMVP:** Cryptographic Module Validation Program [official].
- **ESV:** Entropy Source Validation [official].
- **CNSA 2.0:** NSA Commercial National Security Algorithm suite [secondary].
- **MSID/PSID:** manufacturer/physical secure IDs (TCG Storage) [secondary].
- **KMIP:** Key Management Interoperability Protocol (OASIS) [secondary].
- **PKCS#11:** cryptographic token API standard [secondary].
- **CoCo:** Confidential Containers [independent].
- **VMPL:** VM Privilege Levels (SEV-SNP) [secondary].
- **RMP:** Reverse Map Table (SEV-SNP) [secondary].
- **TDX Module:** Intel's trusted firmware managing trust domains [vendor-reported].
- **EPID/DCAP:** Intel attestation schemes (EPID EOL April 2025) [official].

*End of Phase F3. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*

