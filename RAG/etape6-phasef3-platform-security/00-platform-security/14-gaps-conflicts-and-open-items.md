---
id: etape6-phasef3-platform-security/00-platform-security/14-gaps-conflicts-and-open-items
title: "14. Gaps, conflicts, and open items"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Google", "Intel", "Microsoft", "Nvidia"]
dates: ["2025-04", "2025-10", "2026-06", "2026-08", "2026-09", "2026-09-22", "2026-10"]
keywords: ["amd", "aws", "benchmarks", "distribution", "gpus", "intel", "memory", "nvidia", "pricing", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [152, 189]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 1f2a19e19e592efc9f6e218039e570a693718f5537af6ab1539d0bd96174ce4d
---

# 14. Gaps, conflicts, and open items

- **CHIPSEC:** open-source platform-security assessment framework (Intel); the reference tool for firmware integrity auditing — runs from a Linux live USB in ~15 minutes; modules cover SMM isolation, SPI flash protections, Secure Boot configuration [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- **flashrom / fwupd:** flashrom reads/writes/erases SPI flash chips (`flashrom -p internal -r firmware_backup.bin`); fwupd/LVFS is the Linux firmware-update distribution service (3mdeb is an official fwupd/LVFS consultant) [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- **Binarly efiXplorer:** commercial firmware-integrity/runtime-monitoring tooling cited for detecting NVRAM tampering [secondary](https://gbhackers.com/insyde-uefi-flaw/).
- **Eclypsium Automata:** automated binary-analysis system used to discover the Phoenix SecureCore "UEFIcanhazbufferoverflow" flaw hands-off [secondary](https://www.tomshardware.com/pc-components/cpus/firmware-flaw-affects-numerous-generations-of-intel-cpus-uefi-code-execution-vulnerability-found-for-intel-cpus-from-14th-gen-raptor-lake-to-6th-gen-skylake-cpus?rand=12330).
- Community attestation tooling: tpm2-tools / tpm2-abrmd (Linux), xnode-tpm-attest (Nix-based 7-step TPM2 remote attestation self-test) [independent](https://github.com/johnforfar/xnode-tpm-attest).

## 14. Gaps, conflicts, and open items

- **C1 (conflict):** Microsoft Windows Production PCA 2011 expiry reported as June 2026 by one source and October 2026 by another — unresolved [unverified].
- **C2 (conflict):** TPM market sizing differs across analysts (Mordor: $3.63B 2026→$6.0B 2031, 10.6% CAGR; Verified Market Reports: $3.8B 2024→$12.5B 2032, 16.5% CAGR) — different scopes/periods, non-comparable [secondary].
- **G1:** OpenTitan 2026 commercial adoption status not confirmed in this pass [unverified].
- **G2:** LinuxBoot 2026 adoption beyond hyperscalers not sourced [unverified].
- **G3:** YubiHSM 2 2026 pricing/refresh and AWS CloudHSM 2026 hourly pricing not captured [unverified].
- **G4:** KMIP server market 2026 and IEEE 1667 status not researched [unverified].
- **G5:** FIPS 140-3 premium on SEDs in 2026 unconfirmed (only historical FIPS 140-2 ~25% signal) [unverified].
- **G6:** Caliptra subsystem FIPS certificate availability still TBD; no confirmed shipping products with Caliptra silicon as of 2026-09-22 (AMD integration announced for 2026+ products) [official].
- **G7:** Entrust nShield exact list pricing not published (only "$10,000 starting" aggregator signal) [secondary].
- **G8:** Production cloud GA of Arm CCA Realms not yet observed as of early 2026 [independent].
- **G9:** Phoenix Technologies 2026 corporate/product status not researched in this pass [unverified].
- **G10:** No independent head-to-head 2026 benchmarks for TDX vs SEV-SNP vs CCA overhead found (only vendor/analyst-cited ranges) [unverified].

*End of Phase F3. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*

## 15. Intel SGX lifecycle in 2026: enclaves vs confidential VMs

- Intel SGX (Software Guard Extensions) is the application-enclave TEE model: small trusted code regions (enclaves) inside a process, with memory encryption plus a hardware integrity tree [secondary](https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say).
- Intel's stated position: "There are currently no plans to deprecate Intel SGX on the supported Intel Xeon Scalable processors" — Xeon platforms are "full steam ahead" with SGX [vendor-reported](https://support.avax.network/en/articles/6449817-how-does-intel-support-sgx-technology).
- Client SGX is effectively dead: Intel removed SGX from 11th-gen and later client cores [independent](https://github.com/enclavehost/enclave/blob/HEAD/windows/EDITIONS.md).
- **EPID attestation end-of-life:** on 2 April 2025 Intel discontinued the SGX Attestation Service utilizing Intel EPID (all API versions); migration paths: Intel Tiber Trust Authority (SaaS) or SGX DCAP with ECDSA-based attestation [official](https://www.intel.com/content/www/us/en/developer/archive/tools/sgx-attestation-service-utilizing-epid.html).
- Intel archived the `sgx-tdx-dcap-quoteverificationservice` GitHub repo (read-only from 1 October 2025, no further maintenance) [official](https://github.com/intel/sgx-tdx-dcap-quoteverificationservice/blob/HEAD/README.md).
- Cloud Hypervisor 46 deprecated Intel SGX support (removal planned for v48); Google engineers stepped in to maintain TDX support as Intel contributions declined [secondary](http://portallinuxferramentas.blogspot.com/2025/05/cloud-hypervisor-46-drops-intel-sgx.html).
- Azure SGX mapping (technical review 18 August 2026): DCsv3/DCdsv3 provide SGX-capable processors and EPC capacity for enclave-aware apps; the earlier DCsv2 series retired 30 June 2026; Microsoft will not deploy DCsv2/DCsv3/DCdsv3 in new regions [independent](https://blog.interian.be/2026/09/19/amd-sev-snp-intel-tdx-intel-sgx/).
- Azure confidential VM mapping (August 2026): AMD SEV-SNP families DCasv5/DCadsv5/DCasv6/DCadsv6 (general purpose) and ECasv5/ECadsv5/ECasv6/ECadsv6 (memory-optimized); Intel TDX families DCesv6/DCedsv6 and ECesv6/ECedsv6; NCCadsH100v5 combines SEV-SNP with NVIDIA H100 GPUs [independent](https://blog.interian.be/2026/09/19/amd-sev-snp-intel-tdx-intel-sgx/).
- SGX vs TDX vs SEV-SNP programming models: SGX = application enclaves (code changes required); TDX/SEV-SNP = whole-VM confidentiality (lift-and-shift) [independent](https://blog.interian.be/2026/09/19/amd-sev-snp-intel-tdx-intel-sgx/).
- A tiny enclave with a broad, weakly validated interface can be less trustworthy than a well-hardened confidential VM with simpler data flows — TCB size vs interface complexity is the key trade-off [independent](https://blog.interian.be/2026/09/19/amd-sev-snp-intel-tdx-intel-sgx/).

## 16. DDRop: September 2026 research undermining TEE memory protection

