---
id: etape6-phasef3-platform-security/00-platform-security/33-threat-model-matrix-platform-security
title: "33. Threat-model matrix (platform security)"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "Intel", "Microsoft"]
dates: ["2025-12-05", "2026-07-13", "2026-09-22"]
keywords: ["amd", "aws", "cybersecurity", "intel", "luna", "memory", "research"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [402, 482]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: ebfc94e31670c5d8cdcc42af13ea5340a71f49277564c08b963d5eb24b479f44
---

# 33. Threat-model matrix (platform security)

## 33. Threat-model matrix (platform security)

- Threat: bootkit persistence in SPI flash (CosmicStrand/BlackLotus-class). Controls: Secure Boot + current dbx, SPI write protection, measured boot, SP 800-193 detection/recovery [secondary].
- Threat: SMM exploitation. Controls: SMM isolation, SMRAM locking, CHIPSEC SMM modules, signed SMM drivers [independent].
- Threat: BMC compromise as fleet pivot (CVE-2024-54085-class). Controls: signed BMC updates, network segmentation of management LAN, default-credential rotation, OpenBMC auditability [secondary].
- Threat: supply-chain firmware tampering. Controls: vendor SBOMs, signed updates, Intel Boot Guard/AMD HVB, Caliptra/Ocp LOCK [official][secondary].
- Threat: memory-snooping/interposer on CVM hosts (DDRop-class). Controls: cryptographic-integrity TDX mode where available; physical security; note: no full software fix exists [secondary](https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say).
- Threat: harvest-now-decrypt-later on long-lived data. Controls: PQC migration per CNSA 2.0 timelines; HSM PQC firmware; crypto-agility [secondary].
- Threat: lost/stolen drives. Controls: TCG Opal SED with managed credentials, PBA, crypto-erase procedures [secondary].
- Threat: key exfiltration from software keystores. Controls: HSM (FIPS 140-3 L3), M-of-N quorum, no exportable keys [vendor-reported].
- Threat: rollback to vulnerable firmware. Controls: anti-rollback version fuses/counters, SP 800-193 authenticated update [official].
- Threat: malicious Option ROMs/peripherals. Controls: Secure Boot option-ROM policy, SPDM device authentication, IOMMU/VT-d [secondary].

## 34. Source index (verbatim URLs used)

- https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-193.pdf
- https://www.techpowerup.com/273349/intel-introduces-new-security-technologies-for-3rd-generation-intel-xeon-scalable-platform-code-named-ice-lake
- https://github.com/chipsalliance/caliptra/blob/HEAD/doc/Caliptra.md
- https://github.com/chipsalliance/caliptra/blob/HEAD/doc/caliptra_20/Roadmap.md
- https://lists.chipsalliance.org/g/caliptra-wg/topics?threadid=97432215
- https://www.phoronix.com/news/AMD-Project-Caliptra-2026
- https://github.com/tpm2dev/tpm.dev.tutorials/blob/HEAD/PQC/pqc-ready-tpm.md
- https://github.com/wolfssl/microchip/blob/HEAD/wolfboot-2.9.0-commercial/lib/wolfTPM/ChangeLog.md
- https://github.com/johnforfar/xnode-tpm-attest
- https://www.mordorintelligence.com/industry-reports/trusted-platform-module-market
- https://blog.thomasmarcussen.com/secure-boot-certificate-expiry-2026-blacklotus/
- https://github.com/fleetdm/fleet/issues/45514
- https://www.makeuseof.com/why-windows-secure-boot-can-be-bypassed-so-easily/
- https://csirt.ncc.gov.ng/index.php/resources/security-advisories/66-blacklotus-uefi-bootkit-malware-targeted-fully-patched-windows-11-systems
- https://securityonline.info/secure-boot-bypass-vulnerability-uefi-shell-flaw/
- https://cyberpress.org/uefi-vulnerability-in-insyde-allows-digital-certificate-injection/
- https://gbhackers.com/insyde-uefi-flaw/
- https://securityonline.info/insyde-uefi-flaw-cve-2025-4275-secure-boot-bypass-allows-rootkits-undetectable-malware/
- https://github.com/canopybmc/canopybmc/blob/HEAD/release-notes/2026.06.md
- https://github.com/3mdeb/news-and-ideas/blob/HEAD/blog/content/post/2025-12-05-upstream-and-bmc-ipmi-gigabyte-mz33-ar1.md
- https://forums.servethehome.com/index.php?threads/a-hardware-enthusiast-view-on-the-usefulness-of-open-source-firmwares-like-coreboot.27027/
- https://www.gartner.com/reviews/product/thales-luna-hsm?marketSeoName=operational-technology-security&vendorSeoName=thales&productSeoName=thales-luna-network-hsm&industry=265
- https://www.digicert.com/content/dam/digicert/pdfs/solution-brief/secure-trust-with-digicert-and-thales.pdf
- https://www.executivebiz.com/articles/thales-tct-luna-t-series-hsm-fips-140-3
- https://highways.today/2026/08/05/quantum-deadlines/
- https://www.techjockey.com/detail/entrust-nshield-connect
- https://www.saitechincorporated.com/saitech-inc-strengthens-port-of-oaklands-cybersecurity-with-entrust-nshield-hsm-solutions/
- https://www.thalestct.com/wp-content/uploads/2022/09/luna-network-hsm-pb-11-29-23.pdf
- https://www.peerspot.com/products/comparisons/entrust-nshield-hsm_vs_thales-luna-hsm
- https://www.thetechbag.com/entrust
- https://technology.toolsinfo.com/compare/entrust-nshield-hsm-vs-aws-cloudhsm
- https://www.businesswire.com/news/home/20260421384424/en/KIOXIA-Unveils-Value-Oriented-QLC-based-EG7-Series-SSDs-for-PC-OEMs
- https://www.storagenewsletter.com/2026/01/08/ces-2026-kioxia-unveils-bg7-series-next-gen-up-to-2tb-ssds-for-pc-oems/
- https://www.redeweb.com/en/present/Kioxia%27s-NVME-SSD-cryptographic-module-obtains-FIPS-140-3-Level-2-validation/
- https://www.crn.com/news/storage/227400315/seagate-adds-opal-fips-140-2-standards-to-self-encrypting-hard-drives
- https://HEXUS.net/tech/items/storage/24352-new-internal-self-encrypted-drive-upgrade-laptop-security/
- https://techcommunity.microsoft.com/blog/azureconfidentialcomputingblog/announcing-general-availability-of-azure-intel®-tdx-confidential-vms/4495693
- https://github.com/intel/confidential-computing.tee.dcap.pccs/releases/tag/DCAP_1.27
- https://www.supermicro.com/white_paper/white_paper_Intel_TDX.pdf
- https://github.com/kurt-r2c/llm-security-research-expertise/blob/HEAD/software/crypto-trust/confidential_computing_expertise.md
- https://export.arxiv.org/pdf/2605.26018
- https://ira.lib.polyu.edu.hk/handle/10397/117541?mode=full
- https://voltagegpu.com/blog/azure-confidential-computing-alternative-in-2026-intel-tdx-on-eu-hardware-at-1-4
- https://medium.com/@tfdigitalapp/revolutionary-confidential-computing-adoption-2026-securing-uk-data-in-untrusted-environments-4e12c34a60dc
- https://ayinedjimi-consultants.fr/static/pdf/confidential-computing-tee-2026.pdf
- https://ayinedjimi-consultants.fr/static/pdf/uefi-firmware-bootkits-persistance-2026.pdf
- https://github.com/itsventie/whitepapers/blob/HEAD/CyberSecurity/dfir/firmware/2026-07-13-uefi-spi-flash-nvram-carving.md
- https://github.com/houdini91/uefi-supply-chain/blob/HEAD/COMPLIANCE-MATRIX.md
- https://github.com/pqctoday-org/pqctoday-hub/blob/HEAD/src/data/product-extractions/csc_043_extractions_03312026.md
- https://www.tomshardware.com/pc-components/cpus/firmware-flaw-affects-numerous-generations-of-intel-cpus-uefi-code-execution-vulnerability-found-for-intel-cpus-from-14th-gen-raptor-lake-to-6th-gen-skylake-cpus?rand=12330
- https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf
- https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/
- https://github.com/3mdeb/conferences/blob/HEAD/unveiling-openbmc.md
- https://www.verifiedmarketreports.com/product/trusted-platform-module-tpm-market/
- https://malware.news/t/nist-compliance/74695
- https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say
- https://support.avax.network/en/articles/6449817-how-does-intel-support-sgx-technology
- https://www.intel.com/content/www/us/en/developer/archive/tools/sgx-attestation-service-utilizing-epid.html
- https://github.com/intel/sgx-tdx-dcap-quoteverificationservice/blob/HEAD/README.md
- https://blog.interian.be/2026/09/19/amd-sev-snp-intel-tdx-intel-sgx/
- https://github.com/enclavehost/enclave/blob/HEAD/windows/EDITIONS.md
- http://portallinuxferramentas.blogspot.com/2025/05/cloud-hypervisor-46-drops-intel-sgx.html
- https://github.com/chipsalliance/Caliptra/raw/a8d574c3d8824f137b481136eadb5ae84d11636d/doc/NCC_Group_Microsoft_MSFT283_Report_2023-10-13_v1.2.pdf

*End of Phase F3. Single writer; no other workspace files modified. Research cutoff 2026-09-22.*

