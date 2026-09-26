---
id: etape6-phasef3-platform-security/00-platform-security/part-4
title: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 4)"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "China", "Intel", "Microsoft"]
dates: ["2026-05", "2026-08"]
keywords: ["amd", "intel", "license"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [90, 101]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 428551c2438e01827ee8756d60e84b600412887b775afa9d0b50c79af8335254
---

# Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 4)

- The TPM market was valued at USD 3.63 billion in 2026, projected to reach USD 6.00 billion by 2031 at 10.60% CAGR; virtual TPM is the fastest-growing type (12.8% CAGR) driven by cloud, virtualization, and edge AI; automotive electronics is the fastest-growing end-use segment (12.3% CAGR); North America held 38.2% share in 2025 [secondary](https://www.mordorintelligence.com/industry-reports/trusted-platform-module-market). A second analyst report gives USD 3.8B (2024) → USD 12.5B (2032) at 16.5% CAGR — different scope/period; treat as non-comparable [secondary](https://www.verifiedmarketreports.com/product/trusted-platform-module-tpm-market/).
- Windows 11 (and Server 2025) keeps TPM 2.0 as a core security requirement, turning OS migration into a hardware review/refresh cycle for enterprises [secondary](https://www.mordorintelligence.com/industry-reports/trusted-platform-module-market).
- **Discrete TPM vendors:** Infineon (OPTIGA TPM, SLB 9672/SLB 9673), Nuvoton (NPCT75x), STMicroelectronics (ST33K/ST33TP*), plus Nationz (China) [secondary](https://www.verifiedmarketreports.com/product/trusted-platform-module-tpm-market).
- **Firmware TPMs:** Intel PTT (Platform Trust Technology, in-chipset/ME-based), AMD fTPM (PSP-based), Microsoft Pluton-based implementations. Community attestation tooling (xnode-tpm-attest) fully verified the canonical 7-step TPM2 quote/seal/unseal/credential-activation flow against Intel PTT (Meteor Lake, May 2026) and treats Infineon/AMD-fTPM/Nuvoton/STMicro paths as experimental due to vendor-specific EK cert chains and firmware quirks [independent](https://github.com/johnforfar/xnode-tpm-attest).
- **PQC status (June–August 2026):** the TCG v1.85 library spec defining PQC TPM capabilities is final, but **no discrete hardware TPM ships with PQC capabilities yet** [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
  - Infineon SLB 9672/9673: PQC protection over the TPM *update channel* (XMSS signatures on firmware pushes); the PQC boundary stops at the chip — attestation and key exchange still classical [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
  - STMicro ST33K: similar situation [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
  - SEALSQ: engineering samples of QS7001 secure element / QVault TPM targeted for Q3 2026 [vendor-reported, cited in independent roundup](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
  - wolfSSL wolfTPM firmware TPM: available for v1.85 API development/testing (GPLv2 for development; commercial license to ship); wolfboot 2.9.0 changelog shows 105/113 v1.38 commands (93%), socket/TIS transports, SPDM secured transport for Nuvoton NPCT75x, and STMicro ST33KTPM2X firmware update with LMS support [independent](https://github.com/wolfssl/microchip/blob/HEAD/wolfboot-2.9.0-commercial/lib/wolfTPM/ChangeLog.md).
  - Microchip MEC175xB embedded controller offers immutable hardware PQC (ML-DSA, ML-KEM, LMS) but is *not* a TCG-compliant discrete TPM [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
- Attestation flows: EK certificate chains (Infineon publishes RSA+ECC roots via pki.infineon.com), PCR extend/quote, sealed storage; TPM 2.0 attestation underpins BitLocker PCR-7 sealing, measured boot, and confidential-computing guest attestation [independent][secondary].

