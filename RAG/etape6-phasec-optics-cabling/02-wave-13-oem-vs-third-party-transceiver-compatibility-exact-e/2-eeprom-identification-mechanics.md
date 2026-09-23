---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/2-eeprom-identification-mechanics
title: "2. EEPROM / identification mechanics"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Nvidia"]
dates: []
keywords: ["ethernet", "memory", "nvidia", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2353, 2426]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: fd8d0cc9f9233eb46616a573017a6ff0eb1a2f8d1fc41474c2e10a8633f59487
---

# 2. EEPROM / identification mechanics

## 2. EEPROM / identification mechanics

- SFF-8636 Rev. 2.9: Vendor OUI on upper page 00h bytes 165-167, a three-byte IEEE Company Identifier; all-zero means unspecified. Vendor PN bytes 168-183. `[official/specification]`: https://www.gigalight.com/downloads/standards/sff-8636.pdf
- SFF-8636 Rev. 2.11 common memory map: lower page 00h holds status, monitors, controls, interrupt masks, page select; upper page 00h bytes 128-191 identifier/base ID, 192-223 extended ID, 224-255 vendor-specific ID; page 03h holds module/channel thresholds and controls. `[official/specification]`: http://iommu.com/datasheets/ethernet/phy/qsfp%2B.pdf
- CMIS/QSFP-DD Rev. 3.0 upper page 00h: Vendor name bytes 129-144, Vendor OUI 145-147, Vendor PN 148-163, revision 164-165, serial number 166-181, date code 182-189. `[official/specification copy]`: https://www.fluxlight.com/content/Tech-Docs/QSFP%20DD%20Management%20Specification.pdf
- Do NOT conflate with the SFP/SFP+ (SFF-8472-style) layout, where vendor name is bytes 20-35 and OUI is bytes 37-39 of A0h. `[secondary — from search material, not direct PDF verification in this wave]`
- Cisco-oriented maintenance material (secondary) states switch software reads serial EEPROM fields including serial number, vendor name/ID, a security code, and CRC, then may reject the module if validation fails. `[secondary]`: https://www.scribd.com/document/133653647/Cisco-Small-Form-Factor-Pluggable-SFP-Transceiver-Modules-Maintenance-and-Troubleshooting
- MLNX-OS reads readonly inventory attributes burned in EEPROM: identifier, cable/module type, speeds, vendor, part number, revision, serial number. `[official]`

## 3. OEM support / compatibility policy

### 3.1 Cisco — `[official]` Non-Entitlement Policy v2.0, section 5.1 "Third-Party Components"

- "(a) When You report a product fault or defect and Cisco believes the fault or defect can be traced to the use of third-party memory products, cables, interface components, filters, or other non-Cisco authorized components by You or a reseller, then, at Cisco's discretion, Cisco may withhold support under warranty or a Cisco support program. If Cisco concludes that the fault or defect is not attributable to the use of third-party memory, cables, interfaces, filters, or other non-Cisco authorized components installed by You or a reseller, Cisco will continue to provide support for the affected product under warranty or covered by a Cisco support program."
- "(b) Cisco also reserves the right to charge You for services provided to You when Cisco determines, after having provided such services, that the root cause of the defective product was caused by a non-supported third-party vendor supplied product."
- Source: https://www.cisco.com/c/en/us/products/warranties/warranty-doc-c99-740959.html
- Cisco ISR1000 third-party SFP config guide (IOS XE Fuji 16.9.1): non-Cisco-certified SFPs are called "third-party SFPs"; Cisco-approved modules have undergone rigorous testing; verbatim: `Cisco does not provide any kind of support for the third-party SFPs because they are not validated by Cisco.` From IOS XE Fuji 16.9.1, ISR1000 routers recognize third-party SFPs, though platform-specific qualification rules remain. `[official]`: https://www.cisco.com/c/en/us/td/docs/routers/access/1100/software/configuration/xe-16-10/cisco_1100_series_swcfg_xe_16_10_x/config-third-party-sfps.html
- Widely cited hidden commands `service unsupported-transceiver` and `no errdisable detect cause gbic-invalid`: `[secondary]` — NOT verified from an official Cisco source in this wave.
- TAC behavior summary from a Cisco Community response (links Cisco's official warranty guideline): Cisco may withhold troubleshooting if it suspects third-party optics, then resume after replacement with a supported Cisco optic. `[secondary]`: https://community.cisco.com/t5/optics-and-optical-networking/sfp-compatibility/td-p/4832838

### 3.2 Juniper — `[official]` strong evidence, from JTAC User Guide and optics guides

- "Juniper Networks does not prevent the use of non-Junipercertified (thirdparty) optical transceivers (e.g., by means of a softwareenforced lock) but cannot guarantee the proper operation of thirdparty optical transceivers and/or cables that are not qualified and supplied by Juniper Networks."
- JTAC fully supports Juniper-supplied optics; JTAC does NOT support unqualified third-party optical modules; JTAC may diagnose host-related issues if it believes the issue is unrelated to the third-party module; engineers will likely request testing or replacement with an equivalent Juniper-qualified component.
- High-power third-party coherent ZR/ZR+ optics can potentially cause thermal damage or reduce host lifespan; damage caused by such use remains the user's responsibility.
- Sources: https://supportportal.juniper.net/sfc/servlet.shepherd/document/download/069Dp00000IlqtZIAR?operationContext=S1 ; https://www.juniper.net/documentation/us/en/hardware/coherent-optics-cables-guide/coherent-optics-cables-guide.pdf ; http://juniper.net/documentation/us/en/hardware/100g-optics-cables-guide/100g-optics-cables-guide.pdf
- Practical consequence: Juniper is the major OEM least likely to hard-reject a third-party optic; risk is thermal/operational, not a lockout.

### 3.3 Dell — `[official]` OS9 behavior is media-dependent

- Some non-Dell DACs, AOCs, fan-outs, 1G optics, and adapters can become operational with a warning; certain LR4, SR4, LM4, PSM4, and PSM4-LR modules must be Dell-qualified or the port enters operational-down/error-disabled state; performance of nonqualified media is not guaranteed.
- Source: Dell OS9 config guide link in §1.3.
- OS10 release notes document real third-party interoperability regressions: a third-party QSFP28 100GBASE-SR4 could be misidentified as QSFP56 200GBASE-SR4 after an upgrade leaving links down; enabling DOM for DACs that do not implement DOM could cause link flaps. `[official]`: https://dl.dell.com/content/manual37615048-dell-smartfabric-os10-10-5-5-7-release-notes.pdf?language=en-us
- Dell TAC policy toward third-party optics: only a secondary FS.com summary was found (Dell TAC may ask you to replace with Dell-qualified optics before troubleshooting). `[secondary]` — no official Dell TAC policy page verified in this wave.

### 3.4 Arista — reputation for permissiveness, but thin official evidence in this wave. See §1.7. Do not overstate: no verified EOS policy document was found.

### 3.5 NVIDIA/Mellanox — documentation emphasizes electrical/thermal capability and FEC rather than a vendor-authentication lock. `[official]` See §1.6 and section 6.

### 3.6 HPE/Aruba — AOS-CX provides an official consent-gated `allow-unsupported-transceiver` mode (UT-mode), enabled per the warning text in §1.5; support is on a best-effort basis once enabled. One reseller (EdgeOptic) claims AOS-CX 10.06.0100 enabled allow-unsupported-transceiver by default so third-party optics with a valid EEPROM are accepted without CLI override — `[secondary]`, treat as version-specific and unverified from HPE docs in this wave. Older AOS (e.g., 2930F) documents the same command interactively. `[secondary]`: https://www.winsysadminblog.com/2024/10/aruba-switches-and-trancievers/
- HPE/Comware (FlexNetwork) official policy/messages: NOT found — gap.

## 4. Third-party vendors: coding, testing, warranty claims

### 4.1 FlexOptix (official homepage claims) — `[vendor-reported]`: https://www.flexoptix.net/en

- FLEXBOX configures universal transceivers "in just seconds"; supports nearly 200 vendors; form factors up to 800G; DIY recoding through the FLEXOPTIX app.
- A distributor describes separately configuring each end of DACs for different host systems and lists "MSA Standard" coding. `[secondary]`: https://ausoptic.com.au/flexoptix-100gigabit-qsfp28-to-2x-qsfp28-breakout-dac-cable-passive-copper-cable-1-5-m.html
- FlexOptix's own general transceiver warranty terms: NOT found — gap. Do not infer lifetime coverage (distributor pages show a one-year distributor warranty, which is not FlexOptix's own policy).

### 4.2 ProLabs — `[vendor-reported]`

- Optics coded for the specific target system/class; programmed and tested in-house under operating conditions; backed by a non-transferable lifetime warranty when purchased through an authorized reseller.
- Sources: https://www.prolabs.com/supporting-altnets-and-service-providers-with-compatible-optical-transceivers ; https://www.prolabs.com/dependable-fiber-optics-for-uninterrupted-operations

### 4.3 AddOn Networks — `[vendor-reported]`

- 800G OSFP multi-code transceivers support Arista, Cisco, Juniper, and Mellanox; 100% tested; lifetime warranty: https://www.addonnetworks.com/why-addon/multicoded-optics
- 2026 overview claims form-and-function compatibility across 100 OEMs and 20,000+ systems/platforms, lifetime warranty, 24/7/365 field engineering: https://fe5e0932bbdbee188a67-ade54de1bba9a4fe61c120942a09245b.ssl.cf1.rackcdn.com/AddOn_Product_Overview_2026(2).pdf
- A reseller's AddOn Cisco-compatible 100G SR4 listing: programmed, uniquely serialized, data-traffic/application tested, DOM-capable, limited lifetime warranty. `[secondary]`: https://buy.bluum.com/addon-cisco-qsfp-100g-sr4-s-compatible-taa-compliant-100gbase-sr4-qsfp28-transceiver-mmf-850nm-100m-mpo-dom/

### 4.4 FS — `[vendor-reported]`

- Product pages claim target-switch testing, QSFP28/QSFP-DD MSA compliance, FS BOX real-time configuration; 100G SR4 Cisco-, Arista-, Juniper-, and NVIDIA/Mellanox-coded variants listed at the same price.
- Sources: https://www.fs.com/c/100g-qsfp28-sfp-dd-1159 ; https://www.fs.com/uk/products/199401.html
- FS's own transceiver warranty terms: NOT verified in this wave — gap. Do not claim a lifetime replacement warranty for FS.

### 4.5 10Gtek — `[vendor-reported]`

- Maintains original-brand switches in its test lab; tests Juniper compatibility: https://cn.10gtek.com/new-1424
- Official warranty page: 3-year (36-month) warranty for 1G/10G transceivers and DAC cables; 1-year (12-month) warranty for 40G/100G transceivers and for NICs and media converters; limited warranty extends only to original purchasers, non-transferable: http://cn.10gtek.com/support.html

### 4.6 Approved Optics — NOT adequately researched in this wave — gap (no verified coding/warranty/MSA claims collected).

