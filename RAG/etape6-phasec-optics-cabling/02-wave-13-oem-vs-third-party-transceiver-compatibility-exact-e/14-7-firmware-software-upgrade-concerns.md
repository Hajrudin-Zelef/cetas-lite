---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/14-7-firmware-software-upgrade-concerns
title: "14.7 Firmware / software upgrade concerns"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Broadcom", "Nvidia", "United States"]
dates: ["2024-09", "2025-06", "2025-10", "2026-07", "2026-09-22"]
keywords: ["asic", "cpo", "dsp", "ethernet", "inference", "license", "lpo", "nvidia", "optics", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2598, 2653]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 2b3e10ff591252b62ada1e40136d7ffd9a5705215203b0a2e373edebc63a80df
---

# 14.7 Firmware / software upgrade concerns

### 14.7 Firmware / software upgrade concerns

- **MikroTik — documented breakage case:** RouterOS **v7.12** refactored SFP/QSFP handling: "SFP/QSFP functionality has been refactored for consistent behavior and better scalability. Now, compliance with SFP/SFP+/QSFP MSA standard is mandatory. This may cause issues with SFP/QSFP modules that are not fully compliant." Users reported non-fully-MSA-compliant modules (e.g., some Ubiquiti modules, some DACs) stopped working after upgrading from ROS v6 or pre-7.12, requiring downgrade; the original v7.12 announcement text has since been removed from MikroTik's site [secondary] (https://forum.mikrotik.com/t/compatibility-error-ccr2216-1g-12xs-2xq-with-ubiquiti-modules/177364). Related thread: "RouterOS v7 bug on ROS6 working unrecognized 10Gb SFP DAC" [secondary] (forum.mikrotik.com). **This is the only documented OS-upgrade-breaks-third-party-optic case found in this wave.**
- **Aruba — explicit non-guarantee:** "No guarantees are implied that a thirdparty transceiver will continue to work from release to release" [official doc, mirrored] (AOS-S/AOS-CX Transceiver Guide Ed. 17).
- **Ubiquiti — stated variance:** third-party modules "may be some modules or cables that are not fully compatible", with "revision-to-revision variance; release-to-release variance" listed as risks [official-mirrored] (UniFi help article).
- **Cisco:** no official documented case found of an IOS/IOS XE upgrade breaking third-party optic recognition. Cisco bug IDs CSCut94443 ("GLC-SX-MMD is not recognized after OIR") and CSCuj31712 ("certain Vendor Sfp force ports to errdisable upon OIR") exist but concern Cisco/edge-case optics, not third-party lock changes [vendor-reported — Cisco TAC article]. Cisco-staff caveat: "a Cisco-branded optic from an OEM worked in Cisco product, but the 'identical' optic from the same OEM purchased through a 3rd-party did not work… because the OEM is supplying Cisco with a slightly different version number… with tweaks to address issues Cisco found during qualification" [vendor-reported — community.cisco.com post].
- **10Gtek warranty exclusion ties to firmware:** 10Gtek's warranty "excludes third-party software/upgraded IOS issues" [official] (http://cn.10gtek.com/support.html) — i.e., if a host OS upgrade breaks the coded module, 10Gtek's own warranty excludes it.
- **Juniper, Dell, NVIDIA:** no documented firmware-upgrade breakage cases found — not found.

### 14.8 Counterfeit optics — risks, Cisco guidance, and how legitimate third parties distinguish themselves

- **Cisco definition and alleged risks:** Cisco treats counterfeit transceivers as non-Cisco products bearing Cisco branding passed off as genuine; alleged risks include network disruption, privacy/security breaches, data loss, unsafe/unpredictable malfunction [independent/legal filing] (https://regmedia.co.uk/2020/07/22/cisco.pdf — old court filing, not 400G/800G-specific; do not overgeneralize).
- **Cisco Brand Protection guidance (official):** buy from authorized Cisco channels; check labels (carton/PCBA/module-security holograms — optical transceivers use holographic module-security labels: tilt to see padlock, "BP", check marks); use the "Report Counterfeit Products" form to escalate to Cisco Brand Protection (reviewed daily); ask whether equipment is new, genuine, and not second-hand/pre-owned; be suspicious of too-good-to-be-true prices [official] (https://sec.cloudapps.cisco.com/security/center/resources/anticounterfeit and https://blogs.cisco.com/networking/3-key-steps-to-protect-your-network-from-counterfeit-cisco-products). Cisco also pursues investigations with law enforcement globally [official].
- **Cisco "Buy Right" flyer (official):** warns that purchases from unauthorized resellers may carry **no Cisco support, no Cisco warranty, and no valid end-user software license**, citing risks including counterfeit, modified/upgraded, diverted, gray, stolen, and obsolete products [official] (http://cisco.com/c/dam/en_us/about/legal/brand-protection/buy-right-flyer.pdf).
- **Independent warning:** Computer Weekly quotes Cisco warning that counterfeit products pose risks to network quality, performance, safety, and reliability, recommending Cisco or authorized partners [independent] (https://www.computerweekly.com/news/252486166/Warning-over-security-holes-in-fake-Cisco-kit).
- **Official 2026 vendor statements specifically targeting counterfeit 400G/800G optics: not found.** Do not overgeneralize the above to 800G-specific enforcement.
- **How legitimate third parties distinguish themselves (synthesis — [secondary]/analysis, not a vendor quotation):** all four profiled vendors sell openly **as their own brand** (10Gtek/ProLabs/Approved/FLEXOPTIX), list *equivalent* OEM part numbers only as compatibility references, carry their own warranties and serial/asset tracking, and do not label modules as genuine Cisco. Counterfeits, by contrast, carry unauthorized Cisco trademarks/holograms to pass as genuine. Approved Networks explicitly frames itself as "OEM alternative optical networking connectivity" [official] (https://approvednetworks.com/about-us/).

### 14.9 Wave 14 conflicts, anomalies, and gaps register

1. **Identity:** task label "Approved Optics" is an unofficial variant; official name is Approved Networks. Reconciled in §14.2.1 — both names recorded.
2. **Approved Networks founding year:** conflicting secondary sources (1992 vs 2009); official page silent. Unresolved.
3. **ProLabs founding year:** 2004/carved out 2013 [independent] vs 1995 [secondary]. Unresolved.
4. **ProLabs negative regional stock** (−158, −557, −1): site display/data-feed artifact, not real inventory.
5. **10Gtek US$259 price page** appears stale (search metadata suggests several years old); "price displayed", not recently maintained. 10Gtek 400G/800G optical prices: not found. 10Gtek 400G/800G warranty length: not specified officially.
6. **10Gtek warranty conflict:** retailer "3-year warranty" claims vs official 12-month 40G/100G term — unverified.
7. **Cisco EEPROM check fields:** secondary only; no official source.
8. **Aruba authentication claim:** third-party-vendor mirror only; unverified.
9. **Ubiquiti no-restriction quote:** verbatim from UniFi help article but retrieved via distributor mirror, not help.ui.com directly — [official-mirrored].
10. **Cisco Nexus leniency vs Catalyst:** community consensus only; unverified.
11. **Dell SONiC acceptance:** community consensus only; official docs cover OS9 tiered policy only.
12. **Juniper no-lock:** inference from official JTAC wording + absence of lock documentation; flagged as such.
13. **NVIDIA Cumulus codes 1025/1027/1028:** mirrored NVIDIA docs [secondary]; not verified at docs.nvidia.com.
14. **NVIDIA official third-party policy statement:** not found. NVIDIA official ConnectX-8 announcement/GA page: not found (FS.com catalog shows retail availability only).
15. **All OEM price ratios:** indicative snapshots; the only Cisco figure is a reseller-stated MSRP (US$2,440.10 for QSFP-100G-SR4-S), not a verified Cisco GPL. Do not present ratios as verified OEM-vs-third-party differentials.
16. **Approved Networks & ProLabs public prices:** quote-only; not found. ProLabs 800G product: not found. FlexOptix detailed QA claims: not found. Official per-OEM lists for Ubiquiti/MikroTik (Approved), NVIDIA/Mellanox (10Gtek), FlexOptix target brands: not found.

---

## Wave 15 — 800G ecosystem readiness, silicon photonics, 1.6T signals, availability, alternatives & consolidated gaps (research date 2026-09-22)

**Method:** read-only web research (browser_search / browser_open), public pages only. Complements Waves 2–3 (which covered 800G module types, 1.6T early status, DSP/LPO/CPO) with: shipping 800G switch systems, 800G NICs, named AI deployments, retail price snapshots from FS.com alternatives, silicon photonics shipment status, 1.6T demo evidence, regional availability/lead times, and a consolidated open-items log.

### 15.1 800G switches — shipping systems (as of Sept 2026)

- **Arista 7800R4** — official Arista release dated **29 October 2025**: up to **576×800GbE** with 36-port 800G linecards; **7280R4** variants include **32×800GbE** and **10×800GbE + 64×100GbE**, all shipping [official] (https://investors.arista.com/Communications/Press-Releases-and-Events/Press-Release-Detail/2025/Arista-Networks-Unveils-Next-Generation-Data-and-AI-Centers/default.aspx).
- **Arista 7060X6** — secondary reporting identifies it as **64×800G on Broadcom Tomahawk 5**; the 7800R4 is described as using **Broadcom Jericho3-AI** [secondary] (https://www.demandtalk.com/news/it-infra-news/arista-launches-the-etherlink-ai-platforms-for-ai-workloads/).
- **Cisco 8122-64EH/EHF** — **64×800G**, 2RU, built on **Cisco Silicon One G200**, with QSFP-DD800 or OSFP optics options [independent] (https://www.networkworld.com/article/3564900/cisco-pumps-up-data-center-networking-with-ai-large-workloads-in-mind.html). **Cisco Silicon One G200 datasheet** [official] supports configurations from **64×800GE to 512×100GE** (https://www.cisco.com/c/en/us/solutions/collateral/silicon-one/silicon-one-g200-ds.pdf).
- **Cisco 8122X-64EF-O** — **64×800G SONiC** switch from the G200 family with Cisco 800G LPO support [independent] (https://www.networkworld.com/article/4130263/cisco-amps-up-silicon-one-line-delivers-new-systems-and-optics-for-ai-networking.html).
- **FS 800G TH5 switch** — FS says its Tomahawk-5 system offers **64×800G OSFP** (BCM78900), listed "available now" (announcement circa September 2024) [vendor-reported] (https://www.fs.com/sg/blog/fs-unveils-512t-400g-and-800g-ethernet-switches-powered-by-broadcom-tomahawk-5-10112.html).
- **Broadcom Tomahawk 5 (BCM78900)** — press-release copy: **64×800GbE**, 51.2 Tb/s aggregate, "ships" status [vendor-reported — hosted press release, not Broadcom's own page] (https://www.nasdaq.com/press-release/broadcom-ships-tomahawk-5-industrys-highest-bandwidth-switch-chip-to-accelerate-ai-ml).
- **Broadcom Tomahawk 6** — secondary report dated **4 June 2025**: Broadcom shipping the 102.4 Tb/s chip with 1.6T-port support [secondary] (https://www.storagenewsletter.com/2025/06/04/broadcom-ships-tomahawk-6-first-102-4tb-s-switch). **TH6-based shipping switch systems with model numbers and GA dates: not found.**
- **Marvell Teralynx 10 (ASIC)** — official product brief: **64×800G / 128×400G / 32×1.6T**, 51.2 Tb/s [official] (https://cn.marvell.com/content/dam/marvell/en/public-collateral/switching/marvell-teralynx-10-data-center-ethernet-switch-product-brief.pdf). A July 2026 secondary report says Teralynx 10 had begun production/customer deployment [secondary] (https://convergedigest.com/marvell-ships-51-2-tbps-ethernet-switch-for-ai-data-centers/) — **specific shipping switch-system models based on Teralynx 10: not found.**
- **NVIDIA Spectrum-4 / Spectrum-X800 800G systems** (model numbers, port counts, GA dates): **not found** in the evidence gathered — do not treat catalog presence as GA.
- White-box / ODM 800G systems (Edgecore, Celestica, etc.): not systematically covered in this wave — gap.

