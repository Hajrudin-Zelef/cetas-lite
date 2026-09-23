---
id: etape6-phaseb-smb-networking/08-supplementary-complementary-research-pass-3-september-22-202/overview
title: "Supplementary / Complementary Research Pass #3 — September 22, 2026"
domain: supplementary-complementary-research-pass-3-september-22-202
role: deep-dive
task: reference
actors: ["China", "Huawei", "United States"]
dates: ["2025-08", "2025-10", "2025-10-24", "2025-12-23", "2026-09-22"]
keywords: ["research", "consumer", "cyber", "disclosure", "license", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1403, 1456]
section: "Supplementary / Complementary Research Pass #3 — September 22, 2026"
sha256: 50412cd05ca75993a4eecd5cf25f7401960de6267a2959267742a989c80429b0
---

# Supplementary / Complementary Research Pass #3 — September 22, 2026

This pass adds only material not already present in the base report or Passes #1/#2 (checked September 22, 2026). Existing sections untouched.

## U1. TP-Link regulatory exposure — supplemental detail (developing story, [unverified] beyond cited secondary sources)

Pass #2 T6 already notes the US FCC action against Chinese routers; the following timeline detail is new. Treat as **[unverified]** pending primary FCC/Commerce/DOJ confirmation. It concerns **consumer routers**, not confirmed to extend to Omada business switches/APs:

- August 2025: US DOJ + Commerce Department probe of TP-Link over Chinese-government ties (announced by Congressman John Moolenaar) [secondary].
- October 2025: Texas Attorney General sued TP-Link under the state's Deceptive Trade Practices-Consumer Protection Act, citing "vulnerable, insecure, and unreliable" products; alleged security flaws patched only 24–72 hours after Chinese regulators ordered disclosure, plus a 2023 complaint about a firmware update that bricked devices; state said it received 20+ consumer complaints [secondary, via Texas OAG].
- October 24, 2025: TP-Link Systems designated by the US federal cyber office as covered by Section 5949 of the FY2023 NDAA, citing risks to national security (PRC-owned/controlled entity) [secondary].
- December 23, 2025: US House bill to mandate disclosure of compromised network equipment in federal agencies; named TP-Link and Huawei; alleged TP-Link routers used by Volt Typhoon in infrastructure attacks [secondary].
- 2026: US Commerce Department rule barring TP-Link routers in the US; additional proposed US bills [secondary].
- Market context [secondary]: TP-Link holds 60%+ of the US home-router market; Morgan Stanley estimate — US router prices could rise 15–25% as TP-Link channel stock (3–6 months) drains; Netgear (42% US share at 29% gross margins) is a named beneficiary; enterprise procurement (NIST 2024 supply-chain guidance, NDAA/TAA) favors US/allied vendors.
- Sources observed September 22, 2026:
  - https://en.wikipedia.org/wiki/TP-Link (regulatory section)
  - https://www.webpronews.com/fccs-ban-on-chinese-made-routers-reshapes-the-networking-hardware-market-and-creates-clear-winners/ (already cited in T6)
- **Open question for the RAG corpus:** whether any US action touches Omada business switches/APs. As of September 22, 2026, no source retrieved says so — Omada pricing research remains valid but flagged [unverified on regulatory continuity].

## U2. TP-Link Omada Cloud Standard license pricing (new, closes a gap)

TP-Link's official Omada cloud-controller comparison page states **Cloud Standard** requires a device license fee, while Cloud Essential is free [official]:

- `LIC-OCC-1YR` (1 device / 1 year): **$13.99 USD** at CDW [vendor-reported via reseller].
- `LIC-OCC-5YR` (1 device / 5 years): **CAD 63** at DirectDial [vendor-reported via reseller].
- Australian examples: 3-year license **AUD 66.40**; 5-year license **AUD 95.49** [vendor-reported via resellers].
- Sources observed September 22, 2026:
  - https://www.tp-link.com/dk/business-networking/omada-sdn-controller/omada-cloud-based-controller/ (comparison table)
  - https://www.cdw.com/product/tp-link-omada-cloud-based-controller-license-1-device/8343487
  - https://www.directdial.com/ca/item/tp-link-omada-cloud-based-controller/lic-occ-5yr
  - https://ezypcsales.com.au/product/tp-link-omada-cloud-based-controller-3-year-license-one-device-cbc/
  - https://www.itspot.com.au/tp-link-omada-cloud-based-controller-5-year-license-fee-for-p1369371.html

## U3. TP-Link EAP783 Wi-Fi 7 (new product data point)

- Retail listing: **BE19000** tri-band, 12 streams, two 10GbE ports, PoE++, 39W, 320 MHz channels, MLO [vendor-reported via reseller].
- Price observed September 22, 2026: **AUD 815.30 ex GST / AUD 896.83 incl. GST**; listing showed out of stock [vendor-reported via reseller].
- **Flag:** some listings inconsistently label it "BE22000"; naming discrepancy unresolved — do not merge with the EAP772/EAP772-Outdoor (BE11000) line without official confirmation [unverified].
- Source: https://www.pbtech.com/au/product/NAPTPL7831/TP-Link-Omada-EAP783-BE19000-Tri-Band-12-Stream-Wi

## U4. Zyxel XS1935 — port-map corrections and pricing status

Correction vs earlier file wording: the XS1935 family is **multi-gig + SFP+ hybrid**, not all-fiber:

- **XS1935-10**: 8× multi-gig 10G RJ45 + 2× SFP+ [official datasheet].
- **XS1935-12HP**: 8× 10G multi-gig PoE++ RJ45 + 2× non-PoE multi-gig RJ45 + 2× SFP+, 400W PoE budget [official datasheet].
- **XS1935-12F**: 10× SFP+ + 2× multi-gig RJ45 (10 fiber + 2 copper — not "all-fiber") [official datasheet].
- **Pricing/availability:** no regional pricing or detailed market-by-market availability had been announced as observed September 22, 2026 [unverified]; T7 gap item 3 retained.
- Also new from the same Chinese-market datasheet: previously unspotted **Zyxel Wi-Fi 7 AP models NWA210BE (BE12300)** and **WBE660S (BE22000, triple-radio)** [unverified — model names only, no specs retrieved].
- Sources observed September 22, 2026:
  - https://thetechrevolutionist.com/2026/09/zyxel-xs1935-10gbe-poe-switches-smb-networks.html
  - https://download.zyxel.com/XS1935-10/datasheet/XS1935-10_R_1076.pdf
  - https://download.zyxel.com/XS1935-12HP/datasheet/XS1935-12HP_R_1076.pdf (Chinese-market datasheet)

