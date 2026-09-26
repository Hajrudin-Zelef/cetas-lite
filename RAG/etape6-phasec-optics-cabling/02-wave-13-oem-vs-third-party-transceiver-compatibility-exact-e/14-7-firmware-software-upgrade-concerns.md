---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/14-7-firmware-software-upgrade-concerns
title: "14.7 Firmware / software upgrade concerns"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Nvidia", "United States"]
dates: ["2026-09-22"]
keywords: ["nvidia", "optics", "pricing"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2585, 2608]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: dd041927a1cabca1efd8a9c3a76af16ab3430e52b879e5cce04b133c3d490bb6
---

# 14.7 Firmware / software upgrade concerns

- **Standing warning (repeated from Waves 1/7/13):** the only Cisco figure available is a reseller-stated MSRP, not a verified Cisco GPL/list price. All ratios below are indicative snapshots as of 2026-09-22; street/refurbished prices can be far below nominal MSRP.
- **100G QSFP28 SR4:**
  - Genuine Cisco QSFP-100G-SR4-S: reseller-stated MSRP **US$2,440.10**, street price US$1,936.61 [vendor-reported] (https://techstore.friendsoffice.com/cisco-qsfp-100g-sr4-s-network-transceiver-module-fiber-optic-100000-mbit-s). Another reseller listed factory-sealed units at US$311.99 [secondary] (https://www.cablesandkits.com/mc/transceivers/qsfp-100g-sr4-s/fam-246/fp-17700/) — illustrating how far street/refurbished pricing can fall below nominal MSRP.
  - Third-party Cisco-coded equivalents: AddOn via SHI MSRP $1,200 / $787 [vendor-reported] (https://www.shi.com/product/31696350/); Eaton/Tripp Lite MSRP $1,722.14 [vendor-reported] (https://www.shi.com/product/46750415/); Axiom MSRP $1,096 / $364 [vendor-reported] (https://www.directdial.com/us/item/axiom-100gbase-sr4-qsfp28-transceiver-for-cisco/axg96475); Fleet Network "MSRP" $406.63 / $80 [vendor-reported] (https://fleetnetwork.ca/cisco-compatible-qsfp28-...).
  - Retail third-party: FS.com compatible at US$99.00 [vendor-reported] (https://www.fs.com/c/qsfp28-100g-transceivers-1159); QSFPTEK Cisco/Arista/Juniper/Ciena-compatible at **US$39.90** [vendor-reported] (https://store.qsfptek.com/products/qsfp-100g-sr4-ai); FlexOptix Q.851HG.02 as low as **EUR 73.18** [official] (https://www.flexoptix.net/en/transceiver).
  - Indicative ratio (genuine reseller-stated MSRP vs. lowest third-party retail): ~30:1 (2,440.10 ÷ ~80); vs. mid-range branded compatibles (AddOn/Eaton MSRP): ~1.4–2x. **Do not quote "30:1" without the MSRP caveat.**
- **400G QSFP-DD DR4:**
  - Genuine Cisco QDD-400G-DR4-S: reseller "list" US$5,195.85 vs street US$1,558.87 [vendor-reported] (https://www.hummingbirdnetworks.com/cisco-qsfp-dd-transceiver-module-...qdd-400g-dr4-s); Zones US$1,784.99 [vendor-reported] (https://www.zones.com/site/product/index.html?id=109098466&page_name=product_alert); ESPHERE €1,370 [vendor-reported]. None is a verified Cisco list price.
  - Third-party Cisco-coded: FS.com QDD-DR4-400G-Si US$749 [vendor-reported] (https://www.fs.com/products/128242.html); LINK-PP US$368.94–$429 [vendor-reported] (https://www.l-p.com/store-26044-100-200-400-800g-transceiver-modules.htm); Legrand via SHI MSRP $1,192.66/$1,156 [vendor-reported] (https://www.shi.com/product/47010742/).
  - Retail third-party: QSFPTEK Arista-compatible US$449.90 [vendor-reported] (https://www.qsfptek.com/product/102618.html); PRO-OPTICS generic SiPh US$699.00 [vendor-reported] (https://www.pro-optics.com/store/product/400g-qsfp-dd-dr4); LightOptics (UK) US$519.50 [vendor-reported] (https://www.lightoptics.co.uk/collections/400g-transceiver); FlexOptix D.134HG.05 as low as EUR 490.60 [official] (https://www.flexoptix.net/en/transceiver).
  - Indicative ratio vs. reseller-stated list: ~7x (5,195.85 ÷ 749); vs. street ~US$1.6k: ~2x. **Indicative only.**
- **Reference anchors:** Ubiquiti's own post-cut pricing — SFP+ ~US$9, QSFP28 100G ~US$39 [independent] (https://www.notebookcheck.net/Unifi-Ubiquiti-unlocks-SFP-and-QSFP-modules-with-new-SFP-Wizard.1149839.0.html) — shows how an OEM can undercut third-party retail when it chooses. FS.com generic (non-coded) 400G QSFP-DD DR4 also US$749 [vendor-reported] (https://www.fs.com/products/128245.html).

### 14.7 Firmware / software upgrade concerns

- **MikroTik — documented breakage case:** RouterOS **v7.12** refactored SFP/QSFP handling: "SFP/QSFP functionality has been refactored for consistent behavior and better scalability. Now, compliance with SFP/SFP+/QSFP MSA standard is mandatory. This may cause issues with SFP/QSFP modules that are not fully compliant." Users reported non-fully-MSA-compliant modules (e.g., some Ubiquiti modules, some DACs) stopped working after upgrading from ROS v6 or pre-7.12, requiring downgrade; the original v7.12 announcement text has since been removed from MikroTik's site [secondary] (https://forum.mikrotik.com/t/compatibility-error-ccr2216-1g-12xs-2xq-with-ubiquiti-modules/177364). Related thread: "RouterOS v7 bug on ROS6 working unrecognized 10Gb SFP DAC" [secondary] (forum.mikrotik.com). **This is the only documented OS-upgrade-breaks-third-party-optic case found in this wave.**
- **Aruba — explicit non-guarantee:** "No guarantees are implied that a thirdparty transceiver will continue to work from release to release" [official doc, mirrored] (AOS-S/AOS-CX Transceiver Guide Ed. 17).
- **Ubiquiti — stated variance:** third-party modules "may be some modules or cables that are not fully compatible", with "revision-to-revision variance; release-to-release variance" listed as risks [official-mirrored] (UniFi help article).
- **Cisco:** no official documented case found of an IOS/IOS XE upgrade breaking third-party optic recognition. Cisco bug IDs CSCut94443 ("GLC-SX-MMD is not recognized after OIR") and CSCuj31712 ("certain Vendor Sfp force ports to errdisable upon OIR") exist but concern Cisco/edge-case optics, not third-party lock changes [vendor-reported — Cisco TAC article]. Cisco-staff caveat: "a Cisco-branded optic from an OEM worked in Cisco product, but the 'identical' optic from the same OEM purchased through a 3rd-party did not work… because the OEM is supplying Cisco with a slightly different version number… with tweaks to address issues Cisco found during qualification" [vendor-reported — community.cisco.com post].
- **10Gtek warranty exclusion ties to firmware:** 10Gtek's warranty "excludes third-party software/upgraded IOS issues" [official] (http://cn.10gtek.com/support.html) — i.e., if a host OS upgrade breaks the coded module, 10Gtek's own warranty excludes it.
- **Juniper, Dell, NVIDIA:** no documented firmware-upgrade breakage cases found — not found.

### 14.8 Counterfeit optics — risks, Cisco guidance, and how legitimate third parties distinguish themselves

