---
id: etape6-phasef3-platform-security/00-platform-security/4-firmware-cves-of-20252026
title: "4. Firmware CVEs of 2025–2026"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["EU", "Meta", "Microsoft"]
dates: ["2025-12", "2025-12-05", "2026-06", "2026-07-13", "2026-09"]
keywords: ["cybersecurity", "disclosure", "distribution", "hyperscaler", "memory"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [32, 59]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 7c7aea2255f08ebc9928a802d23c4b813c5e36d573636e54151f4e6e3d98b06f
---

# 4. Firmware CVEs of 2025–2026

- BlackLotus is a UEFI bootkit discovered by ESET exploiting CVE-2023-24932, a Secure Boot Security Feature Bypass vulnerability in the Windows boot manager; it loads before Windows, sits beneath OS-level controls, and survives clean OS reinstalls [secondary](https://csirt.ncc.gov.ng/index.php/resources/security-advisories/66-blacklotus-uefi-bootkit-malware-targeted-fully-patched-windows-11-systems).
- BlackLotus was sold commercially for ~$5,000, democratizing bootkit capability beyond nation-state actors [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- Microsoft's fix has been a multi-phase rollout since 2023, still in progress in 2026: updated bootloader, revocation of the vulnerable boot manager, and certificate replacement [independent](https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/).
- As of 2026, mitigations for CVE-2022-21894 (Baton Drop) and CVE-2023-24932 are still not enabled automatically; enabling requires registry edits, PowerShell commands, manual DBX additions, and applying the Secure Version Number (SVN) to firmware [secondary](https://www.makeuseof.com/why-windows-secure-boot-can-be-bypassed-so-easily/).
- Microsoft has been reluctant to force automatic DBX revocation because revoking compromised Secure Boot keys can brick PCs and break accessories; recovery media (CD/DVD, PXE, USB) also stops working after the patch without a separate recovery-media update [secondary](https://www.makeuseof.com/why-windows-secure-boot-can-be-bypassed-so-easily/).
- NVRAM variable abuse is a recurring attack primitive: MoonBounce hooked the `GetVariable` UEFI Runtime Service to inject kernel payloads; BlackLotus manipulated NVRAM variables to bypass Baton Drop restrictions and inject certificates into the db [independent](https://github.com/itsventie/whitepapers/blob/HEAD/CyberSecurity/dfir/firmware/2026-07-13-uefi-spi-flash-nvram-carving.md).
- Defensive triad cited for 2026: Secure Boot + current DBX (KB5025885) + HVCI (Hypervisor-Protected Code Integrity); TPM 2.0 + BitLocker with PCR 7 sealing detects boot-chain changes by refusing automatic decryption when measurements change; Secured-Core PCs with HVCI and DRTM are described as the best commercially available hardware protection against bootkits in 2026 [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).

## 4. Firmware CVEs of 2025–2026

- **UEFI Shell Secure Boot bypass (September 2026):** Eclypsium researcher Stas Lyakhov found that the firmware logic removing the diagnostic UEFI Shell boot option when Secure Boot is active can be overwhelmed by creating multiple redundant boot entries pointing to the shell. Once in the shell, an attacker with local privileges uses memory-modify commands to overwrite security-enforcement memory values and load unverified payloads [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
  - AMI Aptio V: CVE-2026-33197 [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
  - Cisco UCS servers/appliances: CVE-2026-20293 [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
  - Insyde Software: CVE-2026-6485 [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
  - GIGABYTE confirmed the issue in its AMI Aptio implementations [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
  - No confirmed active exploitation in the wild and no public PoC as of the September 2026 disclosure; true device count likely in the millions given supply-chain distribution [secondary](https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/).
- **Insyde SecureFlashCertData (CVE-2025-4275, CVSS 8.2):** Insyde H2O's unsafe implementation of the SecureFlashCertData NVRAM variable (used to store/exchange public keys in the trust validation chain) lacked protection and could be modified at runtime by an attacker with admin privileges, enabling digital certificate injection and pre-boot malware that survives firmware updates and OS reinstalls [secondary](https://cyberpress.org/uefi-vulnerability-in-insyde-allows-digital-certificate-injection/). Insyde released patches (INSYDE-SA-2025002) and notified OEM partners; remediation depends on device vendors shipping firmware updates [secondary](https://gbhackers.com/insyde-uefi-flaw/). CERT/CC recommended firmware updates, inspection of firmware images/NVRAM behavior with specialized tools, and removal of unprotected variables like SecureFlashCertData from trust management [secondary](https://securityonline.info/insyde-uefi-flaw-cve-2025-4275-secure-boot-bypass-allows-rootkits-undetectable-malware/).
- BMCs remain high-value targets: the MegaRAC CVE-2024-54085 vulnerability is cited in 2026 market analysis as evidence that BMC compromise can threaten entire data-center fleets [secondary](https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf).

## 5. Open firmware: Coreboot, OpenBMC, LinuxBoot

- **Coreboot:** open-source firmware project replacing proprietary UEFI/BIOS init code; 3mdeb is a licensed coreboot service provider (since 2016) and UEFI Adopter (2018), also official consultant for the Linux Foundation fwupd/LVFS project [independent](https://github.com/3mdeb/conferences/blob/HEAD/unveiling-openbmc.md). Practical porting experience (e.g., 3mdeb's Gigabyte MZ33-AR1 coreboot port, December 2025) shows vendor BMC/BIOS interop remains the hard part — Gigabyte's proprietary IPMI-over-VGA-MMIO BIOS-to-BMC communication was judged not worth reimplementing, pushing toward OpenBMC for standardized BIOS-to-BMC methods [independent](https://github.com/3mdeb/news-and-ideas/blob/HEAD/blog/content/post/2025-12-05-upstream-and-bmc-ipmi-gigabyte-mz33-ar1.md).
- **OpenBMC:** Linux-Foundation open-source BMC firmware stack (Yocto-based). HPE enables OpenBMC on select ProLiant Gen11 servers, providing silicon-Root-of-Trust firmware verification from lowest-level firmware through BIOS [secondary](https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/). The CanopyBMC community distribution released 2026.06 (June 2026) with Yocto 6.0 (Wrynose), per-platform SBOMs, HPE ProLiant Gen11 host-BIOS update support, and UBM/NVMe inventory features [independent](https://github.com/canopybmc/canopybmc/blob/HEAD/release-notes/2026.06.md). 2026 market analysis positions OCP/OpenBMC as a differentiation opportunity for modular software services on standardized hardware, while noting supply-chain fragmentation and EU CRA compliance as headwinds [secondary](https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf).
- **LinuxBoot:** replaces most of UEFI DXE with a Linux kernel as the boot loader (used by Meta/Facebook at scale); adoption remains hyperscaler-centric in 2026 [unverified — no 2026-specific source located; flagged as gap].
- Adoption reality check: outside hyperscalers and enthusiasts, OpenBMC/coreboot adoption is limited because most organizations lack the engineering staff to maintain custom firmware; OEM support as a primary platform is the gating factor [independent](https://forums.servethehome.com/index.php?threads/a-hardware-enthusiast-view-on-the-usefulness-of-open-source-firmwares-like-coreboot.27027/).

## 6. Root of Trust concepts and NIST SP 800-193

