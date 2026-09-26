---
id: etape6-phasef3-platform-security/00-platform-security/12-confidential-computing-tdx-sev-snp-arm-cca-gpu-tees
title: "12. Confidential computing: TDX, SEV-SNP, Arm CCA, GPU TEEs"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Google", "Intel", "Nvidia", "Samsung", "United States"]
dates: ["2026-01", "2026-04", "2026-05", "2026-07"]
keywords: ["gpu", "amd", "aws", "blackwell", "compute", "gpus", "intel", "memory", "nvidia", "pricing", "research", "sovereignty"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [125, 151]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 0fac499d114b00df35728bec47d32e6c0d95d6fd1b9a2df7b75baa8f5f01ab38
---

# 12. Confidential computing: TDX, SEV-SNP, Arm CCA, GPU TEEs

- SEDs perform always-on AES-256 hardware encryption of all data on the drive; the data-encryption key never leaves the drive; crypto-erase (key destruction) enables instant secure disposal/repurposing [secondary](https://HEXUS.net/tech/items/storage/24352-new-internal-self-encrypted-drive-upgrade-laptop-security/).
- **TCG Opal** is the Trusted Computing Group storage security specification family: Opal 2.0/2.01/2.02 for client drives, Opalite (lightweight), Enterprise SSC for data-center drives, plus Ruby (NVMe-focused) [secondary]. Pre-boot authentication (PBA) unlocks the drive before OS load [secondary].
- **Kioxia 2026 launches:**
  - **BG7 Series** (CES 2026, Jan 6): first client SSDs with BiCS FLASH Gen 8 + CBA (CMOS directly Bonded to Array); up to 1M random IOPS, 7,000 MB/s sequential read; PCIe 4.0, NVMe 2.0d, M.2 2230/2242/2280; **TCG Opal 2.01** SED support; 256GB–2TB [vendor-reported](https://www.storagenewsletter.com/2026/01/08/ces-2026-kioxia-unveils-bg7-series-next-gen-up-to-2tb-ssds-for-pc-oems/).
  - **EG7 Series** (April 2026): first client QLC (BiCS Gen 8, 4-bit/cell) drives; 1,000 KIOPS, 7,000 MB/s read / 6,200 MB/s write; **TCG Opal 2.02** SED support [vendor-reported](https://www.businesswire.com/news/home/20260421384424/en/KIOXIA-Unveils-Value-Oriented-QLC-based-EG7-Series-SSDs-for-PC-OEMs).
  - **CM7 Series (enterprise NVMe):** PCIe 5.0, 2.5" and E3.S, 1.6–30.72TB; security options include Sanitize Instant Erase (SIE), TCG Opal SED, and **FIPS 140-3 Level 2 validated** SED variants [vendor-reported](https://www.redeweb.com/en/present/Kioxia%27s-NVME-SSD-cryptographic-module-obtains-FIPS-140-3-Level-2-validation/).
- **Samsung, Micron, WD, Seagate:** all ship TCG Opal SED SSDs/HDDs across client and enterprise lines; FIPS 140-2/140-3 validated drive variants carry roughly a 25% price premium over non-validated SEDs (historical Seagate Momentus FIPS pricing signal; 2026 premium unconfirmed — flagged) [secondary](https://www.crn.com/news/storage/227400315/seagate-adds-opal-fips-140-2-standards-to-self-encrypting-hard-drives).
- **Key management:** KMIP (Key Management Interoperability Protocol, OASIS) for centralized key lifecycle; enterprise SED deployments typically pair Opal drives with KMIP servers or OS-native management (BitLocker eDrive, Linux sedutil). 2026 KMIP server market status not researched in this pass — flagged as gap [unverified].
- **IEEE 1667** (standard for authentication in host attachments of transient storage devices) underpins pre-boot auth on removable media; 2026 status not researched — flagged as gap [unverified].

## 12. Confidential computing: TDX, SEV-SNP, Arm CCA, GPU TEEs

- **Concept:** hardware-enforced memory confidentiality/integrity for VMs (CVMs) and enclaves, protecting data in use from the hypervisor, host OS, and cloud operator; remote attestation is the fundamental trust mechanism [secondary](https://ayinedjimi-consultants.fr/static/pdf/confidential-computing-tee-2026.pdf).
- **Intel TDX (Trust Domain Extensions):**
  - Azure GA (announced ~May 2026): next-gen confidential VMs on 5th Gen Intel Xeon with TDX — DCesv6/DCedsv6 (general purpose) and ECesv6/ECedsv6 (memory-optimized) series, up to 128 vCPUs and 512 GiB RAM, available in West US, West US 3, and West Europe; Windows Server 2025, Ubuntu 22.04/24.04, RedHat guests [vendor-reported](https://techcommunity.microsoft.com/blog/azureconfidentialcomputingblog/announcing-general-availability-of-azure-intel®-tdx-confidential-vms/4495693).
  - Intel SGX/TDX PCCS (Provisioning Certificate Caching Service) 1.27: PCCS will be **rebranded to Collateral Caching Service (CCS)** by end of 2026; Windows support for PCCS ends with the next release [official](https://github.com/intel/confidential-computing.tee.dcap.pccs/releases/tag/DCAP_1.27).
  - Supermicro white paper (July 2026): Intel TDX + NVIDIA HGX B200 (Blackwell) systems for "confidential AI at scale" — up to 8 HGX Blackwell GPUs, end-to-end data protection from CPU TEE through GPU [vendor-reported](https://www.supermicro.com/white_paper/white_paper_Intel_TDX.pdf).
  - Pricing signal: Azure Confidential H100 instances listed ~$14/hr (May 2026, VoltageGPU comparison; vendor marketing, treat as indicative) [secondary](https://voltagegpu.com/blog/azure-confidential-computing-alternative-in-2026-intel-tdx-on-eu-hardware-at-1-4).
- **AMD SEV-SNP (Secure Encrypted Virtualization – Secure Nested Paging):** memory encryption + integrity + VM isolation; available on Azure, Google Cloud, AWS (select instances), and OVHcloud (France, sovereignty positioning) [secondary](https://ayinedjimi-consultants.fr/static/pdf/confidential-computing-tee-2026.pdf). UK market comparison (NCSC Hardware Security Lab, Dec 2025, via Medium): TDX 4–7% overhead, SEV-SNP 3–5% overhead [secondary](https://medium.com/@tfdigitalapp/revolutionary-confidential-computing-adoption-2026-securing-uk-data-in-untrusted-environments-4e12c34a60dc).
- **Arm CCA (Confidential Compute Architecture):** Armv9-A extension of TrustZone's two-world model into four worlds (Normal, Secure, **Realm**, Root); Realm Management Extension (RME) ISA support with per-page Granule Protection Table; TF-RMM (Realm Management Monitor) trusted firmware, open-sourced as `tf-rmm` under Trusted Firmware; attestation uses IETF RATS EAT tokens (TEE-agnostic wire format) [independent](https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md).
  - Silicon: Cortex-X4/A720 (Armv9.2-A, RME support, production 2024–2025 flagships) client-side; **Neoverse V3/N3** server cores shipping in 2026 [independent](https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md). Linux kernel support mainline since 6.7; TF-RMM 1.0 released [independent].
  - As of early 2026, production cloud deployment of CCA Realms remains rare — architecture shipping in silicon, but no cloud-vendor offerings comparable to TDX/SEV-SNP GA'd yet [independent](https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md). Academic work (Fasco, 2026; CAGE for confidential accelerators on CCA, IEEE TDSC Jan–Feb 2026) is active [secondary](https://export.arxiv.org/pdf/2605.26018).
- **NVIDIA Confidential Computing (H100/H200):** GPU TEE with encrypted CPU↔GPU transfers (Confidential CUDA) and VRAM encryption; combined with TDX/SEV-SNP CVMs for full chain of trust; managed offerings emerging at hyperscalers in 2026 (Azure Confidential AI: H100 CC + AMD SEV-SNP; Google Cloud confidential VMs with accelerators) [secondary](https://ayinedjimi-consultants.fr/static/pdf/confidential-computing-tee-2026.pdf).
- **Adoption signals 2026:** techUK January 2026 survey — 63% of FTSE 100 companies mandate confidential-computing capabilities in cloud procurement (+210% since 2024); UK drivers: Data (Use and Access) Act 2025 processor accountability, AI Safety Act 2026 model protection, NCSC 2026 Cloud Security Principles for CNI [secondary](https://medium.com/@tfdigitalapp/revolutionary-confidential-computing-adoption-2026-securing-uk-data-in-untrusted-environments-4e12c34a60dc).

## 13. Verification and audit tooling

