---
id: etape6-phasec-optics-cabling/01-wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra/1-4-oui-validation-checksums-and-encryption-claims
title: "1.4 OUI validation, checksums, and \"encryption\" claims"
domain: wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["memory", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1986, 2005]
section: "Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal"
sha256: b9d037c5e2972b2ef4cc76fc9521259a8161f3c4554c5ef44ddcbbc3f6a29c7f
---

# 1.4 OUI validation, checksums, and "encryption" claims

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

