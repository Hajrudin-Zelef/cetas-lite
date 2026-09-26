---
id: etape6-phasef3-platform-security/00-platform-security/7-silicon-rot-implementations-intel-pfr-amd-psp-pluton
title: "7. Silicon RoT implementations: Intel PFR, AMD PSP, Pluton"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "Google", "Intel", "Microsoft", "Nvidia", "Qualcomm"]
dates: ["2018-05", "2025-04-29", "2025-10", "2026-02", "2026-08", "2026-10"]
keywords: ["amd", "intel", "compute", "gpus", "nvidia", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [60, 89]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 8a3a7279d141750dac87f192c295781cf68445cc7c491def3cf36b0786bb1a53
---

# 7. Silicon RoT implementations: Intel PFR, AMD PSP, Pluton

- A Root of Trust (RoT) is the inherently trusted component anchoring a chain of trust; it extends trust to other components in a secure manner [official](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-193.pdf).
- **NIST SP 800-193, Platform Firmware Resiliency Guidelines** (first published 4 May 2018) defines technical guidelines for protecting platform firmware against destructive attacks, organized around three functions [official](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-193.pdf):
  - **Protection:** firmware kept in a state of integrity (usable as RoT); cryptographic write protection so flash regions can only be programmed/erased via authenticated processes with per-system unique keys [official][secondary](https://staging.embeddedcomputing.com/technology/security/iec-iso-other-standards/how-nist-sp800-193-supports-resiliency).
  - **Detection:** unauthorized changes to firmware and critical data detected before execution or use [official].
  - **Recovery:** rapid, secure recovery from attacks (e.g., automatic reflash from a known-good recovery image) [official].
- Key normative rules (§4.1): all RoTs and Chains of Trust shall be immutable or integrity-protected; RoTs/CoTs for Update, Detection, and Recovery in nonvolatile storage shall be implemented in platform firmware; RoTs shall resist tampering by software running under or as part of the host OS; information from host software to platform firmware is untrusted [official](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-193.pdf).
- The standard defines a **Root of Trust for Update (RTU)** / Chain of Trust for Update (CTU): every platform device with mutable firmware relies on an RTU to authenticate firmware updates [official](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-193.pdf). Related RoTs: RTM (Measurement), RTD (Detection), RTRec (Recovery) [official].
- Compliance note: community compliance matrices (e.g., the uefi-supply-chain project) map SP 800-193 controls to evidence such as CHIPSEC module results — chipsec-posture, platform-protection-posture, component-byte-integrity, firmware-digest-anchor [independent](https://github.com/houdini91/uefi-supply-chain/blob/HEAD/COMPLIANCE-MATRIX.md).

## 7. Silicon RoT implementations: Intel PFR, AMD PSP, Pluton

- **Intel Platform Firmware Resilience (PFR):** introduced to Xeon Scalable with 3rd Gen Xeon (Ice Lake); uses an Intel FPGA as the platform Root of Trust to validate critical-to-boot firmware components before any firmware executes. Protected components can include BIOS flash, BMC flash, SPI descriptor, Intel Management Engine firmware, and power-supply firmware; PFR detects and can correct (recover) corrupted firmware [vendor-reported](https://www.techpowerup.com/273349/intel-introduces-new-security-technologies-for-3rd-generation-intel-xeon-scalable-platform-code-named-ice-lake).
- PFR has moved downmarket: ASUS Xeon E-2300/E-2400 entry servers and motherboards (RS300-E11-PS4/RS4, RS300-E12-PS4, P12R-M) ship with an "integrated PFR FPGA as the platform Root-of-Trust solution," targeting ROBO/branch-office deployments [vendor-reported](https://hssl.us/asus-rs300-e12-ps4-is-an-intel-xeon-e-2400-1u-server-with-support-for-four-ddr5-ecc-udimm-four-nvme-one-m-2-one-pcie-5-0-slot-and-one-pcie-4-0-slot-plus-dual-lan-80-plus-gold-power-supply-platform-firmware-resilience-pfr-rs300-e12-ps4-1g2/).
- **AMD PSP (Platform Security Processor):** on-die ARM Cortex-A5 security coprocessor present in AMD CPUs/APUs; hosts AMD's firmware TPM (fTPM) implementation and platform secure-boot functions [secondary]. AMD's fTPM root certificates ship out-of-band (not bundled with standard EK-cert tooling), which has complicated automated attestation deployments [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **Microsoft Pluton:** security processor architecture integrated into AMD, Intel, and Qualcomm CPUs (Xbox-derived design); provides TPM functionality, Secure Boot key protection, and firmware-update isolation. AMD publicly compared Caliptra favorably against its own Pluton integration in 2024, calling Caliptra more useful given its open-source/OCP character [secondary](https://www.phoronix.com/news/AMD-Project-Caliptra-2026).
- Competitive positioning 2026: Intel PFR (discrete FPGA RoT, NIST 800-193 aligned), AMD moving toward open Caliptra integration (see §8), Microsoft Pluton (OEM/Windows-integrated) — three different philosophies for the same problem [secondary].

## 8. Caliptra and OpenTitan: open silicon Roots of Trust

- **Caliptra** is an open-source silicon Root of Trust project born in the Open Compute Project (OCP), backed by Google, Microsoft, NVIDIA, and AMD, and developed under the CHIPS Alliance [secondary](https://www.phoronix.com/news/AMD-Project-Caliptra-2026). It targets data-center silicon (CPUs, GPUs, TPUs, DPUs) with both silicon-level functionality and firmware guarantees [secondary].
- **Caliptra 2.0** defines a design standard for a silicon internal RoT baseline satisfying Root of Trust for Measurement (RTM) plus cryptographic services for the SoC; Caliptra stores boot measurements and reports them with signed attestations rooted in unique per-asset entropy, serving as a Root of Trust for Identity (RTI); the 2.0 subsystem adds RTU (Update) and RTRec (Recovery) [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/Caliptra.md).
- PQC is the headline 2.0 feature: ML-DSA (Dilithium) for secure boot and attestation, ML-KEM (Kyber) for key wrapping, hardware-accelerated cryptography with side-channel countermeasures; the 2.1 subsystem adds Adams Bridge external Mu, ML-KEM spec/implementation, standalone SHAKE256 hardware, AES-GCM DMA, AXI streaming boot over I3C, and OCP LOCK (fuse ratcheting, zeroization) [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md).
- **Milestones 2026:** Caliptra 2.0 RTL released 2025-04-29; 2.0 Subsystem HW released with ROM/FMC/RT firmware finals in October 2025; 2.1 Subsystem HW released 10 October 2025 with firmware releases January–February 2026; **first 2.1 subsystem silicon power-on achieved ~20 August 2026** [official](https://lists.chipsalliance.org/g/caliptra-wg/topics?threadid=97432215). Caliptra Workshop 2026 scheduled 12 October 2026 at the Microsoft campus in Mountain View (integrations, PUF lessons, confidential computing with Caliptra, roadmap) [official](https://lists.chipsalliance.org/g/caliptra-wg/topics?threadid=97432215).
- **AMD to integrate Caliptra from 2026:** announced ahead of OCP Global Summit 2024 — "AMD has strategic plans to integrate Project Caliptra into its 2026+ product lineup" (Alex Tzonkov), expected at least in EPYC data-center products, with Zen 6 timing [vendor-reported](https://www.phoronix.com/news/AMD-Project-Caliptra-2026).
- Microsoft commissioned NCC Group to assess Caliptra's DICE Protection Environment (DPE): objectives included protecting the Unique Device Secret (UDS) and Composite Device Identifier (CDI), preventing firmware-loading bypass, preventing silent dropping of measurements from DPE derivations, and reviewing DPE signing for side-channel leakage [independent](https://github.com/chipsalliance/Caliptra/raw/a8d574c3d8824f137b481136eadb5ae84d11636d/doc/NCC_Group_Microsoft_MSFT283_Report_2023-10-13_v1.2.pdf).
- **OpenTitan:** open-source silicon RoT project (lowRISC stewardship, Google-originated); designed as a discrete/standalone RoT chip with OpenTitan Earl Grey tape-outs. 2026 commercial adoption status not confirmed in this research pass — flagged as gap [unverified].
- Subsystem FIPS certificate availability for Caliptra 2.0/2.1 is still TBD per the official roadmap [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md).

## 9. TPM 2.0: chips, firmware TPMs, and the PQC transition

