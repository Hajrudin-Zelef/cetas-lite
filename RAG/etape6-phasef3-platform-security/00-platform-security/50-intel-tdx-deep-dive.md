---
id: etape6-phasef3-platform-security/00-platform-security/50-intel-tdx-deep-dive
title: "50. Intel TDX deep dive"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["EU", "Google", "Intel", "Nvidia", "United States"]
dates: ["2026-05", "2026-07", "2026-07-13", "2026-09-22"]
keywords: ["intel", "amd", "blackwell", "cyber", "cybersecurity", "disclosure", "distribution", "gpu", "gpus", "hbm", "memory", "nvidia"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [636, 691]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: d073db1bdb419c9c6c7bf6c7d7e3f417a3cfdf9265f804e6eed9c0cb75fd9738
---

# 50. Intel TDX deep dive

## 50. Intel TDX deep dive

- **Trust Domains (TDs):** hardware-isolated VMs; the TDX Module (Intel-signed SEAM firmware) manages TD entry/exit, memory assignment, and measurement [vendor-reported].
- **SEAM (Secure Arbitration Mode):** new CPU mode hosting the TDX Module, isolated from the hypervisor [vendor-reported].
- **Memory protection:** per-TD private HKID (Host Key ID) encrypts memory with AES-XTS; default logical integrity; optional cryptographic integrity (blocks DDRop-class write-dropping except attestation forgery — see §16) [secondary](https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say).
- **TD quote:** TDX Module generates attestation quotes (RTMRs — runtime measurement registers — plus TD attributes) signed via SGX/DCAP collateral chain [official](https://github.com/intel/confidential-computing.tee.dcap.pccs/releases/tag/DCAP_1.27).
- **PCCS → CCS:** Provisioning Certificate Caching Service 1.27 rebrands to Collateral Caching Service by end-2026; Windows PCCS support ends next release [official].
- **TDX Connect:** upcoming PCIe device assignment to TDs (GPU direct assignment) — relevant for confidential AI [vendor-reported].
- Azure TDX GA SKUs (May 2026): DCesv6/DCedsv6 (general purpose), ECesv6/ECedsv6 (memory-optimized); up to 128 vCPUs / 512 GiB; West US, West US 3, West Europe [vendor-reported](https://techcommunity.microsoft.com/blog/azureconfidentialcomputingblog/announcing-general-availability-of-azure-intel®-tdx-confidential-vms/4495693).
- Guest support: Windows Server 2025, Ubuntu 22.04/24.04, Red Hat [vendor-reported].
- Supermicro July 2026 white paper: TDX + up to 8× HGX B200 Blackwell GPUs for confidential AI at scale [vendor-reported](https://www.supermicro.com/white_paper/white_paper_Intel_TDX.pdf).
- Maintenance note: Google engineers took over open-source TDX upkeep in Cloud Hypervisor as Intel contributions declined — watch upstream velocity [secondary](http://portallinuxferramentas.blogspot.com/2025/05/cloud-hypervisor-46-drops-intel-sgx.html).

## 51. Arm CCA deeper

- **Four worlds:** Normal, Secure, Realm, Root — extending TrustZone's two-world model; Root world hosts the EL3 monitor firmware [independent](https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md).
- **RME (Realm Management Extension):** Armv9-A ISA extension; per-page world assignment via the Granule Protection Table (GPT) programmed by EL3 [independent].
- **TF-RMM:** open-source Realm Management Monitor (Trusted Firmware project, `tf-rmm`); 1.0 released; runs in Realm world EL2 managing realm VMs [independent].
- **Attestation:** CCA realms produce EAT tokens per IETF RATS; platform token (signed by HES/root) + realm token (signed by RMM) [independent].
- **Silicon 2026:** Cortex-X4/A720 client (Armv9.2-A, RME); Neoverse V3/N3 server cores shipping 2026 [independent].
- **Kernel support:** Linux 6.7+ mainline RME host support; KVM realm guests in development [independent].
- **Gap (G8):** no production cloud CCA offering comparable to TDX/SEV-SNP GA observed as of early 2026; academic work (CAGE confidential accelerators, IEEE TDSC 2026) active [secondary](https://export.arxiv.org/pdf/2605.26018).

## 52. NVIDIA confidential computing deeper

- H100/H200 GPU CC mode: GPU TEE with encrypted HBM, secure boot of GPU firmware, attestation of GPU measurements [secondary](https://ayinedjimi-consultants.fr/static/pdf/confidential-computing-tee-2026.pdf).
- **CPU↔GPU chain:** CVM (TDX/SEV-SNP) + GPU CC establishes end-to-end encrypted path: data encrypted in CPU TEE, transferred encrypted, decrypted only inside GPU package [secondary].
- Azure: NCCadsH100v5 confidential GPU VMs (SEV-SNP + H100); Azure Confidential AI pairs H100 CC with SEV-SNP hosts [independent](https://blog.interian.be/2026/09/19/amd-sev-snp-intel-tdx-intel-sgx/).
- Supermicro: TDX + HGX B200 (Blackwell) confidential-AI reference (July 2026) [vendor-reported](https://www.supermicro.com/white_paper/white_paper_Intel_TDX.pdf).
- DDRop researchers note NVIDIA CC GPUs are out of reach for interposer attacks because memory sits inside the chip package [secondary](https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say).
- Indicative pricing: Azure confidential H100 ~$14/hr (May 2026, third-party comparison — treat as indicative) [secondary](https://voltagegpu.com/blog/azure-confidential-computing-alternative-in-2026-intel-tdx-on-eu-hardware-at-1-4).

## 53. Firmware forensics and analysis tools

- **UEFITool:** parses UEFI firmware images (volumes, files, sections); standard first step for image triage [independent].
- **EDK2 / UDK debugger:** source-level debugging of DXE drivers where symbols exist [independent].
- **efixplorer (Binarly):** IDA plugin for UEFI reverse engineering; commercial efiXplorer monitors runtime integrity [secondary](https://gbhackers.com/insyde-uefi-flaw/).
- **CHIPSEC:** runtime platform-security auditing (see §41) [secondary].
- **flashrom:** SPI read/write/erase; hardware programmers (CH341A, Bus Pirate) for bricked/externally-dumped chips [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- **NVRAM carving:** extract variable store from flash dumps to find injected drivers/shellcode [independent](https://github.com/itsventie/whitepapers/blob/HEAD/CyberSecurity/dfir/firmware/2026-07-13-uefi-spi-flash-nvram-carving.md).
- **Eclypsium Automata:** automated binary analysis that found the Phoenix SecureCore flaw hands-off [secondary](https://www.tomshardware.com/pc-components/cpus/firmware-flaw-affects-numerous-generations-of-intel-cpus-uefi-code-execution-vulnerability-found-for-intel-cpus-from-14th-gen-raptor-lake-to-6th-gen-skylake-cpus?rand=12330).
- **fwupd/LVFS:** update distribution; version metadata useful for fleet patch-gap analysis [independent].

## 54. UEFI supply chain: IBV and OEM model

- **IBVs (Independent BIOS Vendors):** AMI, Insyde, Phoenix write the base UEFI; OEMs (Dell, HPE, Lenovo, Supermicro, GIGABYTE, ASUS, Cisco) customize and ship [secondary].
- **AMI Aptio V:** dominant server/client UEFI; MegaRAC SP-X for BMC [secondary].
- **Insyde H2O:** strong in client/notebook; SecureFlashCertData flaw (CVE-2025-4275) hit its certificate handling [secondary](https://cyberpress.org/uefi-vulnerability-in-insyde-allows-digital-certificate-injection/).
- **Phoenix SecureCore:** UEFIcanhazbufferoverflow (2024) affected its GetVariable [secondary](https://www.tomshardware.com/pc-components/cpus/firmware-flaw-affects-numerous-generations-of-intel-cpus-uefi-code-execution-vulnerability-found-for-intel-cpus-from-14th-gen-raptor-lake-to-6th-gen-skylake-cpus?rand=12330).
- **EDK2 (TianoCore):** Intel's open reference; PixieFail showed EDK2 network-stack flaws propagate to all IBV derivatives [secondary].
- Economic structure: OEMs pay IBVs per-unit royalties; customization happens in the OEM layer — patches flow IBV → OEM → customer, adding delay [secondary].
- Open alternatives disrupt this chain: coreboot/Dasharo (3mdeb) and OpenBMC/CanopyBMC let ODMs own the firmware [independent](https://github.com/canopybmc/canopybmc/blob/HEAD/release-notes/2026.06.md).
- EU CRA (Cyber Resilience Act) pushes SBOM and vulnerability-disclosure duties onto firmware vendors — reshaping IBV/OEM responsibilities through 2026–2027 [secondary].

*End of Phase F3. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*

