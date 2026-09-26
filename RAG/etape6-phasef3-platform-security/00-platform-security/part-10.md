---
id: etape6-phasef3-platform-security/00-platform-security/part-10
title: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 10)"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["China", "EU", "Samsung"]
dates: []
keywords: ["luna"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [239, 247]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 4bfa7c59189675b8f2834739a187be7749c81202db0d108691801cd07ed01e49
---

# Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 10)

- **Dell iDRAC, HPE iLO, Lenovo XClarity Controller (XCC):** BMCs providing out-of-band firmware update, with signed update packages and (on recent generations) silicon RoT verification; HPE ProLiant Gen11 adds OpenBMC enablement on select models [secondary](https://programminginsider.com/top-openbmc-ecosystem-providers-2026-comparison/).
- MegaRAC (AMI) is the dominant commercial BMC firmware; CVE-2024-54085 demonstrated BMCs as fleet-wide attack vectors [secondary](https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf).
- **Firmware SBOM:** CanopyBMC 2026.06 provides per-platform SBOM information; SBOMs are becoming a procurement requirement under EU CRA [independent](https://github.com/canopybmc/canopybmc/blob/HEAD/release-notes/2026.06.md).
- **fwupd/LVFS:** Linux Vendor Firmware Service distributes signed firmware updates on Linux; 3mdeb is an official consultant [independent](https://github.com/3mdeb/conferences/blob/HEAD/unveiling-openbmc.md).
- NIST SP 800-193's authenticated-update requirement pushes vendors toward signed, versioned, rollback-protected updates with per-system unique keys [official](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-193.pdf).
- U.S. federal supply-chain pressure: Thales TCT's U.S.-designed/manufactured/supported Luna line exists specifically to meet government supply-chain requirements (all employees U.S. citizens, no outsourced support) [vendor-reported](https://www.thalestct.com/wp-content/uploads/2022/09/luna-network-hsm-pb-11-29-23.pdf).
- Geopolitical bifurcation: 2026 market analysis notes nations seeking domestic alternatives to foreign-designed silicon and firmware, complicating global product development [secondary](https://pdf.marketpublishers.com/profresearch/baseboard-management-controller-global-market-profresearch.pdf).
- China-specific TPM ecosystem: Nationz (China) listed among TPM key players alongside Infineon, Nuvoton, Samsung, SK hynix [secondary](https://www.verifiedmarketreports.com/product/trusted-platform-module-tpm-market/).

