---
id: etape6-phasec-optics-cabling/01-wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra/overview
title: "Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal"
domain: wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["optics", "attribution", "ethernet", "memory", "pricing", "research", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1938, 2005]
section: "Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal"
sha256: e0d719cae72f0d9964bf5bb9fd44e8b193853f6af8aebdeade26d2dc00d113f0
---

# Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal

*Wave 12 research report, appended 2026-09-22 (UTC). Scope: third-party coded optics — EEPROM/MSA mechanics, programmers, vendor ecosystem (12 named suppliers), warranty models, OEM lock-in behavior (Cisco, Juniper, HPE/Aruba, Dell, Arista, Extreme), price gaps, legal/gray areas, enterprise/hyperscale practice. Research date context 2026-09-22; all page observations made 2026-09-22. Each factual claim carries one evidentiary label: `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`. Pricing figures are point-in-time observations from search crawls (last crawl dates noted); e-commerce prices are dynamic and vary by region, currency, channel, and inventory. "Coded" means: EEPROM/management-map identity fields (vendor name, OUI, PN, revision, serial, checksums, compliance bytes, vendor-specific bytes) populated to satisfy a target OEM platform's transceiver validation routine — whether at factory or via a programmer.*

## 0. Evidence legend and method notes

- `[official]` — claim on the subject company's own published material (vendor/OEM sites, docs, PDFs) — official to the *claimant*, not necessarily independent verification.
- `[vendor-reported]` — claim by a third-party vendor about its own product/compat/pricing.
- `[independent]` — claim by a party with no sale interest in compatible optics (standards body, open-source project, trade press, court record, academic source).
- `[secondary]` — reseller listing, community forum, aggregator, blog re-posting, market-research site.
- `[unverified]` — claim encountered in research that could not be traced to a citable source in this pass, or attribution is ambiguous.

## 1. Coding mechanics — EEPROM fields, checksums, standards, OUI validation, recoding

### 1.1 SFF-8472 (SFP/SFP+ serial ID, 2-wire A0h/A2h)

- SFF-8472 defines the serial-ID (A0h) base table fields used by host software for identification, and separate CC checksum bytes that the host can recompute to validate integrity `[independent]` — mirrors: https://www.f-tone.com/wp-content/uploads/2026/01/SFF-8472.pdf and https://www.gigalight.com/downloads/standards/sff-8472.pdf
- A0h bytes 20–35: Vendor Name (16 ASCII chars, left-aligned, padded with 20h) `[independent]` — field spec mirror: https://git.wertzui.xyz/wertzui/sfp_util/raw/commit/f35280c1990e882901e0285e17ede5d8b0236ef2/docs/data_fields.pdf
- A0h bytes 37–39: Vendor IEEE company ID (OUI), 3 bytes; a value of all-zero marks the OUI as unspecified `[independent]` — same data-fields reference.
- A0h bytes 40–55: Vendor Part Number (16 ASCII chars) `[independent]` — same reference.
- A0h bytes 56–59: Revision (4 ASCII chars) `[independent]` — same reference.
- A0h bytes 60–61: wavelength/cable compliance fields (laser wavelength encoded as value/20 in nm, or cable attenuation) `[independent]` — same reference.
- A0h byte 63: CC_BASE — check code (low-order 8 bits of the sum) over bytes 0–62 `[independent]` — same reference.
- A0h bytes 68–83: Vendor Serial Number (16 ASCII chars) `[independent]` — same reference.
- A0h bytes 84–91: Date Code (8 ASCII chars) `[independent]` — same reference.
- A0h byte 94: SFF-8472 compliance revision claimed by the module `[independent]` — same reference.
- A0h byte 95: CC_EXT — check code over bytes 64–94 `[independent]` — same reference.
- A0h bytes 96–127: vendor-specific data area, definition left to the vendor `[independent]` — same reference.
- A2h holds alarm/warning thresholds, calibration constants, control/status bytes, and the live DOM/DDM monitor registers (temperature, Vcc, TX bias, TX power, RX power) `[independent]` — same reference.
- An open-source Python library implements reading/decoding SFF-8472 EEPROM data, confirming the byte-level layout is publicly parseable: https://github.com/better-internet-ltd/py-sfp-eeprom `[independent]`.
- Real-world decoded example of a third-party module: an FS-coded QSFP28 LR4 module reported Vendor name "FS", Vendor OUI 00:00:00 (all-zero = unspecified), Vendor PN "QSFP28-LR4-100G", Vendor rev "01", Vendor SN "C2003030166", Date code "20041619", claiming SFF-8636 Rev 2.5/2.6/2.7 compliance `[secondary]` — decoded-output document at https://manuals.plus/m/20de4fde8820eaa8360353df84213a6e721338dd14a5605beae3efa37ab793ac.pdf
- This example shows third-party modules can legally ship with their own vendor name and an unspecified (all-zero) OUI rather than an OEM's OUI, while remaining SFF-compliant `[secondary]` — same decoded-output source.
- Cisco's own maintenance documentation (reproduced across mirrors) states that Cisco-approved SFP EEPROMs contain serial number, vendor name and ID, a unique security code, and a CRC, and that Cisco software rereads and recomputes these fields; a port can be placed in err-disable when validation fails `[official]` — https://www.sanooksales.com/docs/cisco/SFP_maintenance.pdf and https://www.scribd.com/document/133653647/Cisco-Small-Form-Factor-Pluggable-SFP-Transceiver-Modules-Maintenance-and-Troubleshooting

### 1.2 SFF-8636 (QSFP+/QSFP28, paged 2-wire map)

- SFF-8636 uses a paged memory map: Lower Page 00h (bytes 0–127) for ID/status/interrupts/monitors/controls, and Upper Page 00h (bytes 128–255) for the serial-ID block; optional pages 01h–1Fh (user EEPROM, thresholds, vendor-specific) are selected via the page-select byte at Lower Page byte 127 `[independent]` — SFF-8636 Rev 2.11 mirror: http://iommu.com/datasheets/ethernet/phy/qsfp%2B.pdf
- Upper Page 00h bytes 129–191 are the Base ID Fields; byte 191 is CC_BASE, check code over bytes 128–190 `[independent]` — SFF-8636 Rev 2.9 mirror: https://www.gigalight.com/downloads/standards/sff-8636.pdf
- Upper Page 00h bytes 148–163: Vendor Name (16 ASCII, left-aligned/padded); bytes 165–167: Vendor OUI (3 bytes); bytes 168–183: Vendor PN (16 ASCII); bytes 184–185: Vendor rev (2 ASCII) `[independent]` — same Gigalight mirror.
- Upper Page 00h bytes 196–211: Vendor Serial Number (16 ASCII); bytes 212–219: Date Code (8 ASCII) `[independent]` — same Gigalight mirror.
- Upper Page 00h bytes 192–223: Extended ID; bytes 224–255: Vendor Specific ID `[independent]` — same iommu.com Rev 2.11 mirror.
- The spec states "at least one of the vendor name or the vendor OUI fields shall contain valid serial data," i.e., the fields are functional identifiers, not proof of OEM origin `[independent]` — SFF-8436/QSFP-form-factor mirror text: https://www.newnets.ru/upload/medialibrary/7e8/jpx3x3u3ko9ntd4b6jb1du8setxa9huj/17_SFF-8436-QSFP-Form-Factor.pdf
- Lower Page 00h: bytes 0–2 ID and status; bytes 3–21 interrupt flags (clear-on-read); bytes 22–33 device monitors; bytes 86–99 controls; bytes 119–122 optional password change (write-only); bytes 123–126 optional password entry (write-only); byte 127 page select `[independent]` — same iommu.com mirror.
- The password-change/entry bytes show the map includes a write-protection mechanism a host must authenticate past before recoding certain fields `[independent]` — same iommu.com mirror.
- Legacy QSFP28 modules inserted into QSFP-DD host cages fall back to the SFF-8636 management interface rather than CMIS `[independent]` — QSFP-DD hardware specification mirror: https://www.fluxlight.com/content/Tech-Docs/QSFP%20DD%20Hardware%20Specification.pdf

### 1.3 CMIS (QSFP-DD/OSFP, common management interface)

- CMIS defines a 256-byte host-addressable management area split into a 128-byte Lower Memory and a 128-byte Upper Memory; complex modules expose larger maps through dynamically paged 128-byte pages and, where needed, banks `[independent]` — CMIS Rev 4.0: http://www.qsfp-dd.com/wp-content/uploads/2019/05/QSFP-DD-CMIS-rev4p0.pdf
- Flat 256-byte maps suit passive copper/basic devices; complex modules expose serial ID, thresholds, controls, lane data, and application advertisements through pages `[independent]` — same CMIS Rev 4.0.
- CMIS Upper Page 00h layout (per Rev 3.0 mirror): byte 128 identifier; bytes 129–144 vendor name; bytes 145–147 vendor OUI; bytes 148–163 vendor PN; bytes 164–165 vendor rev; bytes 166–181 vendor SN; bytes 182–189 date code; bytes 190–199 CLEI code; bytes 200–201 module power characteristics; byte 202 cable assembly length; byte 203 media connector type; bytes 204–209 copper cable attenuation; bytes 210–211 cable assembly lane information; byte 212 media interface technology; byte 221 custom; byte 222 checksum over bytes 128–221; bytes 223–255 custom info (non-volatile) `[independent]` — CMIS Rev 3.0 mirror: https://www.fluxlight.com/content/Tech-Docs/QSFP%20DD%20Management%20Specification.pdf
- SONiC's open-source platform API documents CMIS and C-CMIS support for coherent ZR modules, confirming CMIS parsing is implemented outside OEM stacks `[independent]` — https://github.com/sonic-net/sonic/blob/HEAD/doc/platform_api/CMIS_and_C-CMIS_support_for_ZR.md

### 1.4 OUI validation, checksums, and "encryption" claims

- Host validation typically recomputes CC_BASE/CC_EXT checksums and matches identity fields against an OEM allow-list; the checksum mechanism itself is public arithmetic (low-order 8 bits of a byte sum), not a secret `[independent]` — same SFF-8472/SFF-8636 mirrors as above.
- Claims that vendor-specific bytes contain "encrypted" hashes or signed data were found only in third-party/vendor marketing and reseller material in this pass; no OEM published specification confirming cryptographic transceiver validation for any of the five named OEMs was found `[unverified]`.
- 10Gtek's compatibility matrix historically described Cisco, Extreme, and Arista coding as "special/encrypted," with older per-unit coding surcharges (Cisco-compatible QSFP+ USD 25/pc more; Extreme-compatible USD 15/pc more), while listing Force10, Juniper, IBM, HP Storage, EMC, Dell, Sun, Emulex, NetApp, Mellanox, and QLogic as accepting its "standard OEM coding" in that tested context `[vendor-reported]` — historical page: https://cn.10gtek.com/new-905 — treat as dated, vendor-reported, platform-specific.
- A Cisco field document confirms Cisco's validation includes a "unique security code" in EEPROM, which is the closest OEM-published reference to non-public validation data found `[official]` — same SFP maintenance mirror.
- Reporting rule: MSA compliance (physical/electrical/management behavior) is conceptually separate from OEM allow-list acceptance; a module can be fully MSA-compliant and still be rejected by a host that enforces vendor-ID matching `[independent]` — derived from SFF-8472/8636 scope plus Cisco/Aruba/Dell/Extreme policy docs below.

### 1.5 Recoding mechanics and limits

- Recoding = writing new identity fields (vendor name, OUI, PN, revision, serial, date, compliance bytes, checksums, vendor-specific bytes) into module EEPROM so a host's allow-list matches; on multi-code/universal modules it can be repeated `[vendor-reported]` — Flexoptix: https://www.flexoptix.net/en; Solid Optics Multi Fiber Tool: https://www.fibre-systems.com/product/solid-optics-releases-multi-function-fibre-tool
- Hardware constraints that can limit recoding: write-protected EEPROM regions, MCU-enforced password authentication (SFF-8636 password bytes), signed/security fields, availability of current vendor coding profiles, and model-level programmer support `[independent]` — SFF-8636 mirror plus FS BOX limitation below.
- Documented programmer-side limitation: FS BOX "does not currently support changing H3C and HPE compatibility for DAC cables" on a FS HPE-Aruba DAC product page `[vendor-reported]` — https://www.fs.com/products/166688.html
- Therefore the FS BOX is designed for supported FS-programmable products, not as an unrestricted programmer for arbitrary OEM optics `[vendor-reported]` — same page and FS BOX datasheet.

