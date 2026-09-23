---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/17-6-programming-coding-workflow-how-a-module-gets-its-oem-c
title: "17.6 Programming / coding workflow (how a module gets its OEM code)"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Broadcom", "Nvidia", "United States"]
dates: ["2026-09-22"]
keywords: ["cpo", "dsp", "ethernet", "lpo", "npo", "nvidia", "optics", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2942, 3012]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 2cc5821e72e30dee0c452eeabf211bbc44bc4a82e2e81dbf7758bba37e782b76
---

# 17.6 Programming / coding workflow (how a module gets its OEM code)

### 17.6 Programming / coding workflow (how a module gets its OEM code)

1. **Order-time coding:** select the OEM brand/model on the product page (FS.com per-brand SKUs; Approved "OP-" SKUs; ProLabs per-OEM product pages) — the module ships pre-coded [official] (Waves 1, 14).
2. **In-house programming (FS BOX):** FS BOX V4.0 (SKU 96657) performs real-time configuration/re-coding of transceivers, DAC and AOC [official] (Wave 1 §1.1).
3. **In-house programming (FLEXBOX):** FlexOptix universal transceivers are reconfigured in seconds via FLEXBOX/app; 400G QSFP-DD/OSFP use FLEXBOX with FLEXSOX; each cable end can be configured separately [official] (https://flexoptix.net/, https://www.flexoptix.net/en/dq-a854hg-z.html).
4. **In-house programming (OnePort):** Approved Networks OnePort codes/re-codes transceivers, DACs, AOCs for 90+ OEM systems at 1G–100G, with batch-coding and favorite lists by OEM/part number [official] (https://approvednetworks.com/blog/the-new-oneport-is-here/).
5. **EEPROM fields written:** P/N, vendor string, length/reach, serial, checksums ("P/N, Length, Vendor, etc.") [official] (https://cn.10gtek.com/new-904). 10Gtek 800G DAC: "EEPROM content can be customized" [official].
6. **Post-coding verification:** confirm on the target host OS that the port comes up without err-disable/unsupported warnings; log DOM baselines (see Wave 16 §16.2 playbook) [analysis].
7. **Recoding caveat:** 10Gtek's warranty "does not apply to products with modified internal electronic information, EEPROM information, etc. without authorized" — unauthorized recoding voids the 10Gtek warranty [official] (http://cn.10gtek.com/support.html).

### 17.7 Standards & SDO quick reference

- **IEEE 802.3bm** — 100GBASE-SR4/40GBASE-SR4 (cited in Approved Networks QSFP28 SR4 datasheet) [official].
- **IEEE 802.3bs** — 400GbE; 400GAUI-8 electrical (cited in 10Gtek Q400-DR4-M8-X2 datasheet) [official].
- **IEEE 802.3ck** — 100G/lane electrical; 800GAUI-8 (cited for 10Gtek 800G DAC/AEC) [official].
- **IEEE P802.3dj** — 200G/400G/800G/1.6T at 200G/lane; late-2026 target vs schedule-risk reports; not completed on collected evidence (Wave 15 §15.6) [independent/secondary].
- **SFF-8472 / SFF-8636** — DOM/DDM two-wire management interfaces (Wave 13 §5) [official/specification].
- **CMIS 5.0** — Common Management Interface Specification for 400G+ (10Gtek 800G AEC) [official].
- **QSFP-DD MSA / OSFP MSA / XFP MSA / QSFP28 MSA** — form-factor multi-source agreements cited across vendor datasheets [official].
- **OpenZR+ MSA** — coherent pluggable interoperability (FlexOptix coherent modules claim) [secondary].
- **UEC (Ultra Ethernet Consortium)** — open spec adopted by Broadcom Thor Ultra 800G NIC [independent].

### 17.8 Remaining micro-gaps (not yet logged elsewhere)

- DAC/AOC per-retailer price matrices (100G/400G/800G DAC, AOC): not systematically gathered — gap.
- White-box/ODM 800G switch systems (Edgecore, Celestica, UfiSpace): not systematically covered — gap.
- 10Gtek official 400G/800G warranty length: unspecified — gap.
- FlexOptix official per-OEM compatibility list: not found — gap.
- Approved Networks warranty exclusions: not found — gap.
- ProLabs public prices and 800G product: not found — gap.
- NVIDIA official third-party policy + ConnectX-8 GA page: not found — gap.
- Dell official third-party support policy: not found — gap.
- 1.6T/OSFP-XD module pricing/sampling/GA: not found — gap.
- Independent regional lead-time datasets: not found — gap.

*End of Wave 17. Phase C complete: 17 waves, append-only, earlier sections untouched.*

---

## Wave 18 — Phase C coverage verification & provenance audit (research date 2026-09-22)

**Method:** editorial audit wave — no new web searches. Maps the original Phase C scope (per file header) to the waves that cover it, and audits provenance-tag discipline.

### 18.1 Scope-to-wave coverage map

| Scope item (file header) | Covered in | Status |
|---|---|---|
| FS.com catalog & pricing | Waves 1, 7, 13 §7, 14 §14.6, 15 §15.4, 16 §16.4, 17 §17.3–17.5 | Covered; prices are dated snapshots |
| Optical transceiver landscape (SFP → OSFP-XD) | Wave 2 | Covered; OSFP-XD has no module evidence (Wave 15 §15.6) |
| 400G/800G standards | Waves 2, 15 §15.6 (802.3dj), 17 §17.7 | Covered; 802.3dj incomplete |
| DSP vs LPO vs CPO/NPO | Wave 3, 15 §15.5 (Marvell CPO demo), 16 §16.5 | Covered |
| DAC vs AOC vs AEC | Wave 6, 14 §14.2 (coding), 16 §16.1, 17 §17.6 | Covered |
| Fiber infrastructure (OS2/OM3–OM5, MPO, polarity) | Waves 4, 8, 16 §16.5 | Covered |
| Copper structured cabling (Cat6A/Cat8) | Wave 5 | Covered |
| Third-party coded optics compatibility | Waves 9–13, 14 (matrix), 16 §16.2–16.3, 17 §17.1–17.2/17.6 | Covered — the promised vendor-compatibility wave is Wave 14 |
| 800G ecosystem readiness | Wave 15 §15.1–15.4, 16 §16.4, 17 §17.3 | Covered |
| Silicon photonics | Wave 15 §15.5 | Covered; direct vendor statements missing |
| 1.6T early signals | Wave 2 §2.4, 15 §15.6 | Covered; no module pricing/GA evidence |
| Regional availability / lead times | Wave 7 §7.7, 15 §15.7 | Covered; independent regional datasets missing |
| FS.com alternatives pricing | Wave 15 §15.4, 16 §16.4, 17 §17.4–17.5 | Covered |
| Counterfeit risks | Wave 14 §14.8 | Covered; no 2026 400G/800G-specific vendor statements |
| Gaps/conflicts/unverified register | Waves 1 §1.11, 2 §2.10, 3 §3.8, 4 §4.9, 5 §5.9, 6 §6.9, 7 §7.8, 8 §8.8, 9–13 registers, 14 §14.9, 15 §15.8, 17 §17.8 | Covered and consolidated |

### 18.2 Provenance-tag audit (spot check, 2026-09-22)

- Tags in use across the file: `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`, plus wave-local qualifiers (`[official-mirrored]`, `[analysis]`, `[standards-based]`) — all defined at first use.
- Standing rules honored in Waves 14–17: the Cisco US$2,440.10 figure is always labeled reseller-stated MSRP, never a verified Cisco list price; all OEM-vs-third-party ratios carry the indicative-snapshot caveat; no SKUs, prices, dates, or URLs were invented; "not found" items are stated, not inferred.
- Name correction applied: "Approved Optics" (task label) → official "Approved Networks", both names recorded (Wave 14 §14.2.1, §14.9).
- File discipline: append-only maintained — Waves 14–18 added after line 2480; no earlier section modified (spot-checked §9 tail at line 2478–2480, intact).

*End of Wave 18. Phase C complete: 18 waves, append-only, earlier sections untouched.*
