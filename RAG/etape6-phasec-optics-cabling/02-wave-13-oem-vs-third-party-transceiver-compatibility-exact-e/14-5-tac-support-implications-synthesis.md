---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/14-5-tac-support-implications-synthesis
title: "14.5 TAC / support implications — synthesis"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["coherent optics", "nvidia", "optics", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2564, 2584]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 8df7ac52e776ef65e40b11d259ffe5cef461e28de7a47ef1f82b62749daa0853
---

# 14.5 TAC / support implications — synthesis

- **Cisco Optics-to-Device Compatibility Matrix (TMG matrix)** [official] (https://tmgmatrix.cisco.com/home) — documented in https://blogs.cisco.com/sp/its-quick-and-easy-with-cisco-optics-online-tools. Inputs: Network Device Product Family/ID, Transceiver Product Family/ID, Data Rate, Reach, Cable Type, Form Factor. Outputs: table with data rate, form factor, reach, cable type, media type, connector, case temp, DOM capable, minimum software support, device/transceiver notes. Export to Excel/PDF/CSV [official — TMG Compatibility Matrix User Manual, https://www.cisco.com/c/en/us/td/docs/interfaces_modules/transceiver_modules/compatibility/matrix/TMG_CM_Tool_User_Manual.html]. **Scope: Cisco optics only — does not validate third-party modules.**
- **Cisco Optics-to-Optics Interoperability Matrix** [vendor-reported] (https://tmgmatrix.cisco.com/iop, also mirrored at https://optics.cisco.com/iop/home) — manual: https://www.cisco.com/c/en/us/td/docs/interfaces_modules/transceiver_modules/compatibility/matrix/TMG_IM_Tool_User_Manual.html. Inputs: transceiver product family/ID, data rate, reach, cable type, form factor. Outputs: which other Cisco optics/standards are optically interoperable over the fiber [vendor-reported; Cisco Blogs, official].
- **Juniper Pathfinder HCT** [official] (https://apps.juniper.net/hct/) — the only fully interactive web matrix found among the seven OEMs; inputs: hardware model / component P/N; outputs: Juniper P/N (740-xxxxx), speed, breakout, form factor, connector, DOM. Third-party optics not covered.
- **Aruba/HPE:** AOS-S and AOS-CX Transceiver Guide (Ed. 17) with per-transceiver compatibility tables and minimum software versions; separate Instant On Transceiver Guide; per-switch datasheets/QuickSpecs carry supported-transceiver lists [official doc, mirrored]. No interactive web matrix found — not found.
- **Dell:** SmartFabric OS10 Hardware Compatibility List (platform-vs-OS matrix per release) [official] (dell.com/support/manuals link in §14.3); Dell Networking Transceiver and Cable Solutions spec sheet (optics table with model, connector, wavelength, distance, Tx/Rx power, notes) [official] (delltechnologies.com PDF in §14.3). No single interactive transceiver↔switch web tool found — not found.
- **Ubiquiti:** no interactive matrix — not found. Closest official artifact: community-tested third-party module list in the EdgeSwitch help article [official] (https://help.ui.com/hc/en-us/articles/205220850).
- **MikroTik:** per-device compatibility tables in the wired-interface compatibility wiki (1G SFP, 10G SFP+/25G SFP28, 40G QSFP+, 100G QSFP28, breakout) [official-mirrored]. No interactive web tool — not found.
- **NVIDIA:** per-product firmware release notes "Validated and Supported Cables and Modules" [official] (docs.nvidia.com). No single interactive optics-compatibility tool found — not found.

### 14.5 TAC / support implications — synthesis

- **Cisco:** strictest documented posture — official Non-Entitlement Policy allows withholding support or charging for services when the fault traces to third-party components; TAC practice is to resume troubleshooting only after the suspect optic is swapped for a Cisco optic [official + vendor-reported] (§14.3).
- **Aruba/HPE:** consent-gated — enabling third-party mode requires acknowledging that use is "at your own risk and may void support and warranty"; Instant On support explicitly does not cover third-party transceivers [official doc, mirrored].
- **Juniper:** no lock, but JTAC "does not provide support for third-party optical modules and cables that are not qualified or supplied by Juniper Networks"; engineer will likely request a swap to a Juniper-qualified component; thermal damage from high-power third-party coherent optics is the user's responsibility [official].
- **Dell:** no official third-party-optics support policy document found — not found; only a secondary claim of a right to refuse repair if the fault is caused by third-party hardware [secondary/unverified].
- **Ubiquiti:** only Ubiquiti-branded modules/DACs get direct support; no lock, no formal entitlement gating found [official].
- **MikroTik:** no restrictions and no documented support gating — nothing found beyond the no-restriction statement.
- **NVIDIA/Mellanox:** no official third-party-optic support statement found — not found; community consensus is that support asks to reproduce with genuine parts [unverified/secondary].

### 14.6 Indicative price ratios (single-snapshot, 2026-09-22 — NOT verified OEM list prices)

