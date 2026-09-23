---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/overview
title: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["license", "nvidia", "optics", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2280, 2352]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 8e1a203494304544bad3be661366b691066828d83402fec22457d92f6fe95315
---

# Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior

*Wave 13 research report, appended 2026-09-22 (UTC). Scope: OEM vs third-party transceiver compatibility — exact error messages verbatim quotes (Cisco NX-OS, IOS/IOS XE, Dell OS9/OS10, Aruba AOS-CX, MLNX-OS, Arista EOS, Juniper, MikroTik/Ubiquiti), EEPROM/identification mechanics, OEM support/compatibility policy (Cisco, Juniper, Dell, Arista, NVIDIA/Mellanox, HPE/Aruba), third-party vendors (FlexOptix, ProLabs, AddOn, FS, 10Gtek, Approved Optics), DOM/DDM failure diagnosis and failure modes, FEC, dated price snapshots, practical compatibility notes, explicit verification gaps. Research date 2026-09-22; all observations from the English-language public web. No files written by the research pass.*

## How to read this report

- Every factual claim carries a tag: `[official]` = OEM/spec/vendor official doc; `[vendor-reported]` = third-party optic vendor's own claim; `[independent]` = industry/independent source; `[secondary]` = reseller/blog/distributor or unverified copy of an official doc; `[unverified]` = community/one-off claim.
- All prices are dated snapshots (observed 2026-09-22), not timeless prices.
- Section 9 lists every error string that could NOT be verified verbatim.

## 1. Exact error messages (verbatim quotes)

### 1.1 Cisco NX-OS — `[official]` verified from Cisco NX-OS troubleshooting guide

- `PORT-3-IF_UNSUPPORTED_TRANSCEIVER: Transceiver for interface [chars] is not supported.`
- Cisco's explanation: `Transceiver (SFP) is not from an authorized vendor.`
- Example: `Nov_1 14:07:58 excal-113 %PORT-3-IF_UNSUPPORTED_TRANSCEIVER: Transceiver for interface fc1/13 is not supported`
- Sources: https://www.cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/101x/troubleshooting/cisco-nexus-9000-nx-os-troubleshooting-guide-101x/m-overview.html and the NX-OS system messages reference (sl_nxos_book).

### 1.2 Cisco IOS / IOS XE — `[secondary]`, NOT verified from an official source

- Secondary copy (Scribd copy of a Cisco SFP maintenance document) gives:
  `SYS-3-TRANSCEIVER_NOTAPPROVED:Transceiver on port [dec]/[dec] is not supported`
  Source: https://www.scribd.com/document/133653647/Cisco-Small-Form-Factor-Pluggable-SFP-Transceiver-Modules-Maintenance-and-Troubleshooting
- A separate secondary compilation quotes Cisco's GBIC security messaging:
  `%GBIC_SECURITY_CRYPT-4-VN_DATA_CRC_ERROR: GBIC in port [dec] has bad crc`
  `%GBIC_SECURITY-4-EEPROM_SECURITY_ERR: GBIC in port [dec] failed security check`
  `%GBIC_SECURITY-4-EEPROM_CRC_ERR: GBIC in port [dec] has bad EEPROM checksum`
  `%PMBUY-7-GBIC_INSERTED: GBIC inserted in port [dec]`
  `%PMBUY-7-NON_CISCO_GBIC_INSERTED: Non Cisco GBIC inserted in port [dec]`
  Tag: `[secondary/unverified as Cisco-authored]`. Do NOT treat these as officially verified until confirmed against Cisco's IOS/IOS XE system message references.

### 1.3 Dell EMC Networking OS9 — `[official]` verified from Dell OS9 config guide

- Permitted but nonqualified media:
  `Apr 29 05:09:16: %S4048-ON:1 %IFAGT-5-UNSUP_OPTICS: Non-qualified optics in slot 1 port 49`
- Disallowed nonqualified LR4/SR4/LM4/PSM4-family optics:
  `Apr 29 05:09:16: %S4048-ON:1 %IFAGT-2-TRANSCEIVER_UNSUPPORTED_ERROR: Transceiver in slot 1 port 50 unrecognized, putting interface in operational-down state`
- Interface state shown: `fortyGigE 1/50 is up, line protocol is down(error-disabled[Transceiver Unsupported])`
- Module line shown: `Non-qualified pluggable media present, QSFP type is 40GBASE-SR4`
- Source: https://www.dell.com/support/manuals/en-sg/dell-emc-os-9/c9010-9.14.2.8-config-pub/Non-Dell-Qualified-Transceivers?guid=guid-46826015-c0df-4b18-9534-96701fd44009&lang=en-us

### 1.4 Dell OS10 — `[unverified/community]` only

- A Dell Community user reported the string `serial number:MT1841VS05210 is third party and is disabled` and said `allow unsupported-transceiver` / `unlock third-party transceiver` did not work in OS10 10.5.5.3. No official OS10 rejection string was found.
- Source: https://www.dell.com/community/en/conversations/networking-general/unlock-third-party-transceiver-doesnt-work-on-os10-10553-license-question/653af5dc7af2d67a933c8281

### 1.5 Aruba/HPE (AOS-CX) — `[official]` verified from HPE/Aruba tech docs

- The enable command emits this warning text verbatim:
  `Warning: The use of unsupported transceivers, DACs, and AOCs is at your own risk and may void support and warranty. Please see HPE Warranty terms and conditions.`
  `Do you agree and do you want to continue (y/n)?`
- Disabling it warns: `Warning: Unsupported transceivers, DACs, and AOCs will be disabled, which could impact network connectivity. Use 'show allow-unsupported-transceiver' to identify unsupported transceivers, DACs, and AOCs.`
- AOS-CX transceiver event IDs (official event reference): 3801 `allow-unsupported-transceiver feature enabled`; 3802 `allow-unsupported-transceiver feature disabled`; 3803 `allow-unsupported-transceiver feature enabled: Unsupported transceivers found in: <list>`; 3806 `Interface <interface> transceiver attempted link recovery <count> times` (Warning); 3808 `Transceiver <xcvr_desc> inserted in <interface> is <status>. <reason>`.
- Sources: https://arubanetworking.hpe.com/techdocs/AOS-CX/10.14/HTML/fundamentals_6200/Content/Chp_IfaceCfg/Iface_cmds/all-uns-tra.htm and https://arubanetworking.hpe.com/techdocs/AOS-CX/10.14/HTML/elmrg/Content/events/XCVR.htm
- NOTE: these are consent/UT-mode messages, not an automatic "rejected" syslog. No automatic-rejection syslog string was verified. HPE/Comware (FlexNetwork) exact messages: NOT found — see gap list.

### 1.6 NVIDIA/Mellanox MLNX-OS — `[official]` verified from MLNX-OS user manual

- `Warning: High power transceiver is not supported` — when this condition occurs the link does not come up. This is a port power-capability error, NOT proof of a vendor whitelist.
- Source: https://docs.nvidia.com/networking/display/nvidia-mlnx-os-user-manual-v3-12-3000.3000.pdf
- No verified generic "third-party Mellanox cable rejected" string was found.

### 1.7 Arista EOS — NOT found. No trustworthy verbatim EOS unsupported-optic syslog was found in this wave. Explicitly unverified; do not quote one.

- A secondary Arista-authored FAQ copy (Scribd) says third-party copper cables are not restricted if IEEE-compliant and unrecognized interfaces are disabled, but the copy is not a strong primary source: `[secondary/unverified]`.
- An Arista-authored technical note for NetVisor states: `Arista NetVisor OS does not lock Transceiver support`, while warning third-party firmware/implementation variations can prevent off-the-shelf compatibility. This is NetVisor, not standard EOS; frame narrowly. `[official, product-specific]`: https://techdocassets.pluribusnetworks.com/F9000-HW/transceivers/DEON/TechnicalNote.html
- WARNING: the string `L1SOURCE-6-UNSUPPORTED: The configured Layer1 source for interface name is unsupported.` is about a Layer-1 source configuration, NOT an unsupported transceiver. Do not misuse it.

### 1.8 Juniper — no unsupported-vendor syslog was found, and official policy explains why: Juniper does NOT software-lock non-Juniper-certified optics (see §3.2). Speed/optic mismatch alarms exist but are not vendor-authentication messages.

### 1.9 MikroTik / Ubiquiti — no official generic-optic policy or error strings were found in this wave. Only marketplace evidence that generic/open-coded modules are commonly marketed for both. `[gap]`

