---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/wave-14-vendor-compatibility-matrix-third-party-coded-optics
title: "Wave 14 — Vendor compatibility matrix: third-party coded optics vs OEM platforms (research date 2026-09-22)"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["optics", "research", "inference", "nvidia", "parameters", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2484, 2503]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: b9f93bb9edf41b066dbfb0fa4d0ff3f18f80511a46abe0051ac60d138b9c27bc
---

# Wave 14 — Vendor compatibility matrix: third-party coded optics vs OEM platforms (research date 2026-09-22)

## Wave 14 — Vendor compatibility matrix: third-party coded optics vs OEM platforms (research date 2026-09-22)

**Method:** read-only web research (browser_search / browser_open), public pages only, no logins or forms. This wave fills the explicitly promised "vendor compatibility" gap: per-vendor profiles (Approved Networks, 10Gtek, ProLabs, FlexOptix), per-OEM policy matrix (Cisco, Aruba/HPE, Dell, Ubiquiti, MikroTik, Juniper, NVIDIA), official compatibility-check tools, TAC/support implications, indicative price ratios, firmware-upgrade concerns, and counterfeit risks.
**Naming correction (important):** the task label "Approved Optics" is imprecise — the official brand is **Approved Networks**, a Legrand brand [official] (https://approvednetworks.com/about-us/). The name "Approved Optics" does not appear to be an official name. Both names are recorded here so the record stays searchable, but the official name is used throughout.

### 14.1 How transceiver coding / EEPROM compatibility works

- An optical transceiver carries an EEPROM (two-wire serial interface per SFF-8472 / SFF-8636 / CMIS) that identifies the module to the host: vendor name, part number, serial number, revision, wavelength, link length, and checksum/CRC fields. Host OSes read these fields at insertion and decide whether to bring the port up, warn, or error-disable it [independent/standard references; see Wave 13 §2 for SFF-8636 DOM field mechanics].
- "Coding" (also called programming or burning) a third-party optic means writing the EEPROM identification fields so the host recognizes the module — in many cases writing the same vendor strings/part numbers the OEM's own modules carry, plus module-specific serial numbers and correct checksums [vendor-reported] (https://www.10gtek.com — "To meet system compatibility, we need to modify the EEPROM information"; modifiable fields incl. "P/N, Length, Vendor, etc." [official] (https://cn.10gtek.com/new-904)).
- "Compatible with Cisco" therefore means: the module's EEPROM is programmed (coded) to Cisco's expected identification and the optical/electrical parameters conform to the relevant MSA/IEEE spec — it does **not** mean the module is manufactured by Cisco or carries Cisco branding [secondary/analysis].
- **Cisco EEPROM check fields — caveat:** the specific fields Cisco checks (vendor name, vendor ID, serial number, "security code", CRC) were found only in a secondary CCIE blog citing cisco.com; no official Cisco document listing the checked fields was found. Treat as [secondary/unverified].
- **Aruba authentication fields — caveat:** a claim that "Aruba switches perform authentication on connected optical modules by default" was found only in a third-party-vendor mirror of an Aruba support document; no HPE.com page confirming which EEPROM fields are checked was found. Treat as [unverified].
- **MikroTik's stated requirement** is simpler and official: no vendor restrictions at all; modules must comply with the transceiver MSA and "should be compatible" if they do [official-mirrored] (https://help.mikrotik.com/docs/display/ROS/MikroTik+wired+interface+compatibility — retrieved via mirror https://my.mbuzztech.com/portal/en/kb/articles/https-help-mikrotik-com-docs-display-ros-mikrotik-wired-interface-compatibility).
- **Ubiquiti's stated position** is official and unusually explicit: "We do not impose any artificial restrictions to limit third-party SFP modules or DAC cables. However, we also do not formally test these. As a result, there may be some modules or cables that are not fully compatible with UniFi devices. Ubiquiti SFP modules and DAC cables are guaranteed to work with UniFi devices." [official — UniFi help article "SFP Modules and DAC Cables", viewed via distributor mirror https://www.ubiquitidubai.com/unifi-sfp-modules; tag the verbatim quote as [official-mirrored]].
- **Juniper's stance:** no vendor-ID lock is enforced; Junos does not block third-party optics from linking. This is inferred from the official JTAC policy wording (recommends Juniper optics, describes support behavior) plus the absence of any documented lock/err-disable behavior — the no-lock characterization is an inference and is flagged as such [official policy text; characterization [secondary]].

### 14.2 Third-party vendor profiles (coding, warranty, compatibility, prices)

#### 14.2.1 Approved Networks (official name; "Approved Optics" is an unofficial variant)

