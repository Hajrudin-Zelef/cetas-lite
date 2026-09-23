---
id: etape6-phasef3-platform-security/00-platform-security/40-key-management-and-code-signing-infrastructure
title: "40. Key management and code-signing infrastructure"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: []
dates: ["2026-08"]
keywords: ["cybersecurity", "luna"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [528, 580]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 3cadad5bf23301d6a0302f24c177e8afc7bdbfef10ab05761fc9fe0457d570c9
---

# 40. Key management and code-signing infrastructure

## 40. Key management and code-signing infrastructure

- **KMIP (OASIS):** protocol for HSM ↔ key-manager communication; enterprise SED/PKI deployments centralize keys in KMIP servers [secondary].
- **HashiCorp Vault:** software key management with HSM auto-unseal (Luna/nShield/CloudHSM as seal); transit engine offloads crypto to HSM [vendor-reported].
- **Thales CipherTrust Manager:** centralized key/lifecycle management integrating Luna HSMs (Luna T7 validated inside CipherTrust Manager deployments) [vendor-reported](https://www.executivebiz.com/articles/thales-tct-luna-t-series-hsm-fips-140-3-validation).
- **DigiCert Trust Lifecycle Manager + Luna:** PKI roots in HSM, ACME/CMPv2/EST enrollment, PQC-ready [vendor-reported](https://www.digicert.com/content/dam/digicert/pdfs/solution-brief/secure-trust-with-digicert-and-thales.pdf).
- **Code signing:** EV code-signing certs require FIPS HSM-backed private keys (CA/Browser Forum rules); signing HSMs typically network-attached with timestamping [secondary].
- **Sigstore:** open-source code-signing (Fulcio/Rekor) with short-lived certs; HSM-backed CA roots in production deployments [independent].
- **5G SIM/eSIM and payment:** HSM use cases include SIM provisioning and payment HSMs (Thales payShield lineage) [secondary](https://www.saitechincorporated.com/saitech-inc-strengthens-port-of-oaklands-cybersecurity-with-entrust-nshield-hsm-solutions/).

## 41. CHIPSEC module catalog (audit reference)

- CHIPSEC modules audit: `smm` (SMM isolation), `smm_dma` (SMM DMA protection), `spi_lock` (SPI flash locking), `spi_desc` (flash descriptor), `bios_wp` (BIOS write protection) [independent].
- `secureboot` modules: `variables` (db/dbx/PK/KEK inspection), `sb` (Secure Boot config) [independent].
- `cpu` modules: `cpuid`, `msr`, `smm_ptr` (SMM pointer protections) [independent].
- `uefi` modules: `uefi_vars`, `uefi_tables`, image parsing checks [independent].
- Usage: boot a Linux live USB, install CHIPSEC, run `chipsec_main.py -m common.<module>`; ~15-minute baseline audit per machine [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- Enterprise use: baseline PCR/SPI measurements per fleet SKU, re-run after firmware updates, alert on drift [secondary].

## 42. fwupd/LVFS operations

- **fwupd:** Linux daemon for firmware updates; LVFS (Linux Vendor Firmware Service) distributes vendor-signed capsules; `fwupdmgr get-updates` workflow [independent].
- Vendors publishing on LVFS include Dell, Lenovo, HP; 3mdeb is an official fwupd/LVFS consultant for ODM enablement [independent](https://github.com/3mdeb/conferences/blob/HEAD/unveiling-openbmc.md).
- Capsule updates: UEFI UpdateCapsule mechanism; signed capsules verified against db/KEK or vendor keys [secondary].
- Limitation: LVFS coverage is strongest on client devices; server BMC/BIOS updates still flow through iDRAC/iLO/XCC [secondary].

## 43. TPM provisioning and TCG guidance

- TCG TPM 2.0 Provisioning Guidance covers EK cert provisioning, platform attestation CA setup, and supply-chain verification of TPM authenticity [secondary].
- OEMs provision EK certificates at manufacture; verifiers fetch EK certs from vendor repositories or the TPM's NV storage [independent](https://github.com/johnforfar/xnode-tpm-attest).
- Privacy: EK is unique per TPM — direct EK use would track devices; AKs (with privacy CA or DAA) decouple identity from attestation [secondary].
- TPM 2.0 NV indices store EK certs, platform certificates, and small secrets with access policies [secondary].
- Dictionary-attack protection: TPMs throttle authorization failures (lockout) — brute-force PIN protection [secondary].

## 44. SED operations detail

- **MSID (Manufacturer Secure ID):** default credential printed on the drive label; must be changed at provisioning — drives left on MSID are effectively unencrypted to anyone with physical access [secondary].
- **SID/Admin SP:** TCG Storage security providers; Admin SP manages locking ranges and credentials; SID is the initial authority [secondary].
- **RevertSP / PSID revert:** returns the drive to factory state with cryptographic erase; the operational "break glass" for lost credentials [secondary].
- **TCG Storage Opal SSC feature sets:** Single User Mode, additional DataStore table, PSID, global-range locking — check drive datasheets for which are implemented [secondary].
- **NVMe sanitize commands:** Crypto Erase, Block Erase, Overwrite; distinct from TCG crypto-erase but complementary; FIPS drives certify specific sanitize semantics [secondary].
- Fleet practice: pre-provisioning scripts (sedutil-cli) set SID passwords, enable locking ranges, record PSIDs in the asset database before deployment [secondary].

## 45. PQC migration playbook (platform security)

- Inventory: list every place classical crypto anchors platform trust — Secure Boot signers, TPM EK/AK algorithms, HSM keys, SED auth, firmware-update signers, attestation CAs [secondary].
- Prioritize by data lifetime: anything protecting data/keys that must survive past 2030 migrates first (harvest-now-decrypt-later) [secondary](https://highways.today/2026/08/05/quantum-deadlines/).
- HSMs: confirm per-firmware PQC validation (Thales TCT v7.15.1 ML-DSA/ML-KEM/LMS as reference point); plan Luna 8 / nShield PQC firmware upgrades [vendor-reported](https://www.executivebiz.com/articles/thales-tct-luna-t-series-hsm-fips-140-3-validation).
- TPMs: no PQC discrete TPM silicon as of August 2026; track SEALSQ QS7001 and TCG v1.85 implementations; use LMS/XMSS for firmware signing in the interim [independent](https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md).
- Caliptra 2.x provides ML-DSA/ML-KEM in the RoT — prefer Caliptra-based platforms for PQC-ready device identity [official](https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md).
- Hybrid deployments: run classical + PQC in parallel during transition (X.509 hybrid certs, dual signatures) to avoid flag days [secondary].
- CNSA 2.0 deadlines: key establishment 2030, signatures 2031 (U.S. NSS); ANSSI 2027 certification cutoff (France) [secondary](https://highways.today/2026/08/05/quantum-deadlines/).

