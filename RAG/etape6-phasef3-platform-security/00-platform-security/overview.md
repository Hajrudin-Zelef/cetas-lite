---
id: etape6-phasef3-platform-security/00-platform-security/overview
title: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "EU", "Intel", "Microsoft", "Nvidia"]
dates: ["2026-06", "2026-07-13", "2026-09", "2026-09-22", "2026-10"]
keywords: ["amd", "cyber", "cybersecurity", "disclosure", "distribution", "intel", "luna", "memory", "nvidia", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [1, 50]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: c998d357f6dffcbde79b949a6895baac7fb8a1f8773845e02ac7fa608775f469
---

# Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust

**Scope:** BIOS/UEFI supply chain and CVEs, Secure Boot (2026 certificate transition), open firmware (Coreboot/OpenBMC/LinuxBoot), Root of Trust (NIST SP 800-193, Intel PFR, AMD PSP, Pluton), Caliptra and OpenTitan, TPM 2.0 (chips, PQC status, market), HSMs (Thales Luna, Entrust nShield, YubiHSM, CloudHSM, FIPS 140-3), self-encrypting drives (TCG Opal, SED vendors), confidential computing (Intel TDX, AMD SEV-SNP, Arm CCA, NVIDIA CC).
**Research date / cutoff:** 2026-09-22. **Method:** read-only web research (browser_search / browser_open); no live-browser visits, nothing sent externally. **Provenance legend:** [official] = vendor/standards-body publication; [vendor-reported] = vendor marketing/press without independent confirmation; [independent] = third-party technical analysis; [secondary] = press/aggregator; [unverified] = claim found but not corroborated. No identifiers guessed; all URLs verbatim from search results.

---

## 1. BIOS/UEFI fundamentals and the firmware supply chain

- UEFI (Unified Extensible Firmware Interface) is the firmware interface replacing legacy BIOS; it connects the OS to platform hardware and implements Secure Boot, measured boot, and the pre-OS driver environment (DXE phase) [secondary].
- The x86 server firmware supply chain is highly concentrated: AMI (American Megatrends, Aptio V), Insyde Software (H2O), and Phoenix Technologies (SecureCore) supply the base UEFI codebase — including the Initial Boot Block (IBB) — to nearly all OEMs (Dell, HPE, Lenovo, Supermicro, GIGABYTE, ASUS, Cisco UCS) [secondary]. A single codebase flaw therefore propagates across hundreds of models and brands simultaneously [secondary].
- Intel's open-source EDK2 (TianoCore) is the common base that AMI, Phoenix, and Insyde build on; flaws in EDK2-derived network/USB stacks (e.g., the PixieFail campaign, CVE-2023-45229 through CVE-2023-45237, nine vulnerabilities in the UEFI IPv6 network stack) simultaneously affected servers from Lenovo, Dell, HP, Intel, Microsoft (Hyper-V), and AMD [secondary].
- Eclypsium's "UEFIcanhazbufferoverflow" research demonstrated buffer-overflow code execution in Phoenix SecureCore UEFI firmware via an unsafe `GetVariable` UEFI service call, affecting Intel CPUs from 14th Gen Raptor Lake back to 6th Gen Skylake, plus Lenovo, Intel, Insyde, and AMI codebases; a TPM alone does not save an affected system since the compromise occurs below the measurement point [secondary](https://www.tomshardware.com/pc-components/cpus/firmware-flaw-affects-numerous-generations-of-intel-cpus-uefi-code-execution-vulnerability-found-for-intel-cpus-from-14th-gen-raptor-lake-to-6th-gen-skylake-cpus?rand=12330).
- France's ANSSI recommends hardened firmware posture for the most sensitive information systems (OIV/OSE operators) under its security framework [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- The EU Cyber Resilience Act (CRA), entering into force in 2027, will impose legal security-update obligations for the product lifetime on all manufacturers of connected equipment sold in Europe — a structural change expected to transform firmware patching practices [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).
- French critical organizations are increasingly writing contractual firmware-SLA clauses (e.g., maximum 30 days between a critical firmware CVE publication and a signed update, priority notification channels, UEFI source access under NDA for independent audit); Dell, Lenovo, and HPE reportedly accept such clauses for large accounts [secondary](https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf).

## 2. Secure Boot architecture and the 2026 certificate transition

- Secure Boot verifies each boot component against keys enrolled in firmware: the Platform Key (PK), Key Enrollment Key (KEK), signature database (db), and forbidden signature database (dbx) [secondary].
- Three Microsoft 2011-era certificates baked into virtually every Windows device's firmware expire in 2026 [independent](https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/):
  - **Microsoft Corporation KEK CA 2011** — expires June 2026 [independent].
  - **Microsoft Corporation UEFI CA 2011** — signs third-party UEFI components and option ROMs; expires June 2026 [independent].
  - **Microsoft Windows Production PCA 2011** — signs the Windows boot manager; expires June 2026 per one source, October 2026 per another (conflict registered; see §12) [independent].
- Replacements are the 2023 certificate set: Microsoft Corporation KEK CA 2023, Microsoft UEFI CA 2023, and Windows UEFI CA 2023 [independent](https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/).
- Microsoft's coordinated migration adds Windows UEFI CA 2023 to firmware trust, replaces the boot manager, and revokes old vulnerable boot managers via DBX in one pass; all three steps must complete or the host fails to boot or stays vulnerable [secondary](https://github.com/fleetdm/fleet/issues/45514).
- Patching alone does not apply the migration: Microsoft gates it behind a registry opt-in or Group Policy; Windows Server gets no automatic rollout at all [secondary](https://github.com/fleetdm/fleet/issues/45514).
- Detection is complicated because the Windows boot manager exists in two places: `C:\Windows\Boot\EFI\bootmgfw.efi` (staging copy) and the ESP copy that firmware actually loads. The staging copy refreshes on a later Windows Update cycle, sometimes weeks later, so a host can be mitigated while the staging copy still shows the old PCA 2011 signature — signature checks on the staging copy alone produce false negatives [secondary](https://github.com/fleetdm/fleet/issues/45514).

## 3. BlackLotus and the Secure Boot bypass problem

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

