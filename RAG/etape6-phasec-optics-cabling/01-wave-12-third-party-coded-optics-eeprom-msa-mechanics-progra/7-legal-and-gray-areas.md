---
id: etape6-phasec-optics-cabling/01-wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra/7-legal-and-gray-areas
title: "7. Legal and gray areas"
domain: wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2026-09-22"]
keywords: ["consumer", "optics", "pricing"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2205, 2229]
section: "Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal"
sha256: fda95df9ad0325a10a10bd125628949a9ae24c6066be795a4afd97a89432fa2e
---

# 7. Legal and gray areas

- **10G SFP+ pair (most comparable):** Cisco SFP-10G-SR list $995 (older FS-cited Cisco price) vs FS Cisco-compatible SFP-10G-SR US$25.00 (SKU 11552, "1.3M Sold", 5-year warranty) — same form factor, reach, and rate; FS price observed 2026-09-22, Cisco list price figure from older vendor marketing `[vendor-reported]` — https://www.fs.com/blog/a-comprehensive-understanding-of-cisco-10g-sfp-8890.html and https://www.FS.COM/products/11552.html?attribute=95058&id=4246741
- AddOn Cisco SFP-10G-SR compatible: $217.00 (MSRP $330.00) on SHI — illustrates the spread *within* third-party brands for the same Cisco-coded part `[secondary]` — https://www.shi.com/product/29319327/AddOn-Cisco-SFP-10G-SR-Compatible-SFP-Transceiver
- StarTech.com Cisco SFP-10G-SR compatible: $55.00 (MSRP $60.46), limited lifetime warranty — further within-third-party spread `[secondary]` — https://www.shi.com/Product/26893618/StarTech.com-Cisco-SFP-10G-SR-Compatible-SFP-Module
- **100G QSFP28 pair:** Cisco QSFP-100G-SR4-S list price $1,995 (secondary historical) `[secondary]` — https://www.seesuo.com/article/70350.html; secondary reseller MSRP $2,440.10 / sale $1,936.61 `[secondary]` — https://techstore.friendsoffice.com/cisco-qsfp-100g-sr4-s-network-transceiver-module-fiber-optic-100000-mbit-s
- FS Cisco QSFP-100G-SR4-S-compatible QSFP-ESR4-100G (SKU 147611): US$139 — **flag NON-COMPARABLE**: it is ESR4/300 m on OM4, not the same 100 m SR4 specification `[vendor-reported]` — https://www.fs.com/products/147611.html
- Legrand Cisco-compatible QSFP-100G-SR4-S: $109 (MSRP $106.25), lifetime warranty — secondary reseller; warranty/pricing asymmetry vs MSRP is unexplained, flag `[secondary]` — https://www.shi.com/product/47010799/CISCO-QSFP-100G-SR4
- Axiom Cisco QSFP-100G-SR4-S-AX: $418 (MSRP $481.54) `[secondary]` — https://www.publicsector.shidirect.com/Product/42470309/Axiom-Cisco-QSFP-100G-SR4-S=-Compatible; TAA variant $464 (MSRP $1,096.67) — **non-comparable to the $418 listing** (different TAA status and listed MPN) `[secondary]` — https://www.shi.com/product/42468608/Axiom-Cisco-QSFP-100G-SR4-S=-Compatible
- AddOn Cisco QSFP-100G-SR4-S equivalent: $787 (MSRP $1,200) `[secondary]` — https://www.shi.com/product/31696350/AddOn-QSFP28-transceiver-module-(equivalent-to:-Cisco-QSFP-100G-SR4-S); $720.72 with 4–6+ week ETA at another reseller `[secondary]` — hssl.us listing.
- **400G DR4:** FS Cisco QDD-400G-DR4-S compatible (SKU 128242): US$749.00, "1K Sold" `[vendor-reported]` — https://www.fs.com/products/128242%20.html; LINK-PP Cisco QDD-400G-DR4-S compatible: $368.94–429/pc `[secondary]` — https://www.l-p.com/store-26044-100-200-400-800g-transceiver-modules.htm; LightOptics Cisco QDD-400G-DR4 compatible: $519.50 `[secondary]` — https://www.lightoptics.co.uk/collections/400g-transceiver
- **No Cisco official list price for QDD-400G-DR4-S was found in this pass** — do not present the 400G third-party prices as a quantified OEM discount `[unverified]`.
- Cross-cutting caveat: retail prices vary with TAA status, warranty terms, condition (new/refurbished), geography, channel, and inventory; never present a clean apples-to-apples OEM discount without the comparability notes above `[secondary]` — synthesis of collected listings.

## 7. Legal and gray areas

- **Core distinction:** a lawful "compatible" product carries its *own* brand while describing OEM equivalence; a counterfeit falsely bears OEM marks (Cisco/Ciena logos, forged labels) — this distinction separates the third-party coding industry from counterfeiting `[independent]` — derived from case records below.
- Cisco and Ciena sued Wuhan Wolon entities in 2021 over allegedly counterfeit-marked transceivers; the proposed consent judgment included a permanent injunction and potential **$10,000 per counterfeit item** in liquidated damages `[independent]` — https://lawstreetmedia.com/news/tech/cisco-systems-and-ciena-corporation-resolve-counterfeit-transceiver-suit-with-chinese-firms/
- Cisco secured injunctions against alleged counterfeiters including Shenzhen Tianheng, Gezhi Photonics, Shenzhen Sourcelight, and Dariocom `[independent]` — https://aphnetworks.com/news/21688-cisco-secures-injunction-against-chinese-counterfeiters
- Dexon dispute: Dexon alleged Cisco spread FUD and made anticompetitive threats; Cisco characterized Dexon as an unauthorized reseller selling counterfeit goods — case `Dexon Computer Inc v. Cisco Systems Inc`, E.D. Texas, No. `5:22-cv-00053-RWS-JBB` `[independent]` — https://www.ciplawyer.com/articles/152656.html
- **Magnuson-Moss claims:** multiple third-party vendors/resellers (FS, Optcore, ATGBICS) assert the US Magnuson-Moss Warranty Act protects against warranty voiding merely for using third-party optics `[vendor-reported]` — https://www.fs.com/blog/deep-analysis-on-optical-transceiver-module-6543.html; https://www.optcore.net/article43032/; https://atgbics.com/blogs/tech-talk/compatible-optical-transceivers-are-as-reliable-as-original-and-save-you-money-too
- Do not state that Magnuson-Moss universally protects enterprise service contracts — the act's consumer-product scope and contractual context require caution; vendor invocations are marketing, not legal analysis `[unverified]`.
- Vendor claims invoking the Sherman Act, TFEU, or UK Competition Act must be labeled `[vendor-reported]`; they are not a substitute for jurisdiction-specific legal analysis `[unverified]`.
- Compatible-labeling rule of thumb from collected practice: nominative description ("compatible with Cisco QSFP-100G-SR4-S") under the seller's own brand is the industry's standard lawful posture; reproducing OEM trademarks as product identity crosses into counterfeiting risk `[secondary]` — synthesis of labeling practice across FS/AddOn/Axiom listings.

## 8. Enterprise/hyperscale practice

