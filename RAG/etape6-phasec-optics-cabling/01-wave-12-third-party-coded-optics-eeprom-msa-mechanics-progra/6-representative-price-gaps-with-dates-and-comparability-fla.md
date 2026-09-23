---
id: etape6-phasec-optics-cabling/01-wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra/6-representative-price-gaps-with-dates-and-comparability-fla
title: "6. Representative price gaps (with dates and comparability flags)"
domain: wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Cohere", "Meta", "Microsoft", "Nvidia", "OpenAI", "United States"]
dates: ["2021-01-28", "2025-05-28", "2026-01-09", "2026-03-12", "2026-09-22"]
keywords: ["advisory", "amd", "antitrust", "asic", "attribution", "compute", "consumer", "datacenter", "hyperscaler", "incident", "nvidia", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2203, 2279]
section: "Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal"
sha256: 4e5c52b4078d459d222058fa3dfafebc0a76f5a3cf56307ce85bb9d4c0714549
---

# 6. Representative price gaps (with dates and comparability flags)

## 6. Representative price gaps (with dates and comparability flags)

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

- Facebook-inspired TIP Voyager: open packet-DWDM whitebox architecture with twelve 100G QSFP28 client ports, four 200G coherent line modules, a Broadcom Tomahawk ASIC, and hardware/software separation through open specifications `[independent]` — https://www.lightwaveonline.com/optical-tech/transport/article/16654294/telecom-infra-project-intros-voyager-white-box-open-packet-dwdm-transponder
- Optical Compute Interconnect / Optical Scale-up Consortium announced **2026-03-12**: founding members AMD, Broadcom, Meta, Microsoft, NVIDIA, and OpenAI; goal of an open, multi-vendor supply chain for AI scale-up optics; NRZ plus WDM; silicon-centric model; supports pluggable, onboard, and eventually co-packaged approaches `[independent]` — https://www.businesswire.com/news/home/20260312254951/en/Optical-Scale-up-Consortium-Established-to-Create-an-Open-Specification-for-AI-Infrastructure-Led-by-Founding-Members-AMD-Broadcom-Meta-Microsoft-NVIDIA-and-OpenAI and https://www.lightwaveonline.com/home/article/55365387/ofc-2026-optical-scale-up-consortium-sets-path-for-an-open-ai-infrastructure-specification
- Hyperscale procurement pattern (with caution): hyperscalers pursue multi-source/open specifications and direct qualification; their optical suppliers are module manufacturers (Coherent, Lumentum, Innolight, Eoptolink, Accelink, HG Genuine) rather than retail recoding brands `[secondary]` — https://growthmarketreports.com/report/optical-transceiver-market-global-industry-analysis
- One investment-research note estimated Innolight and Eoptolink supply roughly 60% of NVIDIA's 800G volume, and NVIDIA deployed >$6.5B into photonics suppliers YTD 2026 including $2B each into Coherent and Lumentum with multiyear purchase commitments — treat as analyst/secondary claims, not verified procurement facts `[secondary]` — https://github.com/47tzp4ydc9-cmyk/life-os/blob/HEAD/investment-os/narrative/research/silicon-photonics.md and https://umele-pestovani.eu/umele/files/The-company-with-the-highest-proportion-of-high-end-optical-module-products_Sun-22-Sep-2024-13386.pdf
- **Do not assert** "hyperscalers use retail third-party coded optics" — no evidence for that was found; the verified pattern is direct manufacturer relationships and open specs `[unverified]`.
- Evidence for specific direct hyperscaler contracts with InnoLight/Eoptolink/Coherent/Lumentum/Source Photonics was not collected beyond the analyst notes above `[unverified]`.
- Transceiver industry history note (secondary/educational): equipment vendors mostly rebrand modules built by the large module manufacturers — "an optic with a vendor's name on it and one without are frequently the same component with different firmware" — widely repeated but keep as `[secondary]` — https://github.com/ronutz/arsenal/blob/HEAD/src/content/learn/en/transceiver-family-history.mdx

## 9. Conflicts, unverified claims, and gaps register

- **FS warranty conflict:** current 5-year statements vs older blog "lifetime warranty" language — unresolved within FS's own material `[vendor-reported]`.
- **ProLabs warranty wording:** "Lifetime ProLabs Replacement Warranty" vs "5 Years limited Lifetime Advance Replacement Warranty" in historical brochure — internally conflicting `[secondary]`.
- **FS 100G price non-comparability:** QSFP-ESR4-100G (US$139) is 300 m ESR4, not 100 m SR4 — do not use as the SR4-compatible price `[vendor-reported]`.
- **Axiom price non-comparability:** TAA vs non-TAA listings with different listed MPNs — separate series `[secondary]`.
- **400G DR4 gap:** no Cisco official list price found — no OEM-discount ratio can be stated `[unverified]`.
- **Encryption claims:** "encrypted coding" for Cisco/Arista/Extreme is vendor-reported (10Gtek) without OEM confirmation `[unverified]`.
- **Skylane failure rates:** <0.03%/<0.02% figures in a 2024 deck cannot be attributed to Skylane due to ambiguous slide layout `[unverified]`.
- **InterOptic "1 billion hours" MTBF:** extraordinary vendor claim with no published method `[vendor-reported]` — flag.
- **Netceed reliability figures** (99.98%, <0.02% failure, 0% DOA): vendor-reported with no published method `[vendor-reported]` — flag.
- **Magnuson-Moss vendor invocations:** marketing claims, not jurisdiction-specific legal analysis `[vendor-reported]`.
- **"Same as OEM"/"100% compatible":** vendor-reported marketing across FS/AddOn/Approved Networks/Netceed/ATGBICS unless independently validated — flag each instance `[vendor-reported]`.
- **Missing in this pass:** EDGE Optic profile; FluxLight company profile/pricing; current 800G launches/acquisitions 2024–2026 beyond Integra (2025-05-28), FiberMall (2026-01-09), Halo/Skylane (2021-01-28), AddOn/Amphenol; comparable OEM-vs-compatible 40G SR4 pair; direct hyperscaler procurement contracts `[unverified]`.

## 10. Source ledger (key URLs, by category)

- **Standards:** SFF-8472 mirrors (f-tone.com, gigalight.com), data-fields PDF (git.wertzui.xyz), py-sfp-eeprom (github.com/better-internet-ltd), SFF-8636 Rev 2.9 (gigalight.com), Rev 2.11 (iommu.com), SFF-8436 (newnets.ru), CMIS Rev 4.0/5.0/5.1 (qsfp-dd.com), CMIS Rev 3.0 (fluxlight.com), QSFP-DD HW spec (fluxlight.com), SONiC CMIS doc (github.com/sonic-net).
- **OEM policy/docs:** Cisco IOS XE troubleshooting (cisco.com), Cisco GBIC/SFP support note (static-cisco.com), Cisco Catalyst 9200 fiber troubleshooting (cisco.com), Cisco NX-OS transceiver firmware (cisco.com), Cisco warranty (audentia-gestion.fr), Juniper QFX5100 HW guide (gotomojo.com), Juniper brand protection (manuals.plus), ArubaOS-Switch unsupported-transceiver guide (higherlogicdownload/HPE), Aruba Instant On (media.bechtle.com), Dell C9000 non-qualified transceivers (dell.com), Dell ProSupport terms (dell.com), NetVisor/Arista note (techdocassets.pluribusnetworks.com), Arista EOS um-eos (arista.com), Extreme policy (extreme-networks.my.site.com), Extreme pluggable guide (documentation.extremenetworks.com).
- **Programmers:** FS BOX V4 product/datasheet/quick-start (fs.com, resource.fs.com, img-en.fs.com), Flexoptix (flexoptix.net), FLEXBOX reseller/pricing (ausoptic.com.au), Solid Optics Multi Fiber Tool (fibre-systems.com), Solid Optics brochure (storage.googleapis.com).
- **Vendor blogs/products:** FS 800G article, FS SFP guide, FS 10G comparison blog, FS product pages (SKUs 11552, 128242, 147611), AddOn catalog/multicode/objections (addonnetworks.com), ProLabs (prolabs.com, itweb.africa, slideshare), 10Gtek (cn.10gtek.com), Axiom (axiomupgrades.com), Champion ONE (lightwaveonline.com, telecomramblings newswire), Skylane/Halo (prnewswire.co.uk, datacenter-forum.com, digitec.ch), Approved Networks (approvednetworks.com), Netceed (netceed.com), ATGBICS (atgbics.com), InterOptic (interoptic.com), Integra (telecomramblings newswire), FiberMall (globenewswire.com), FS Box compatibility blog (fs.com).
- **Reseller evidence:** SHI/publicsector.shidirect (Axiom, AddOn, StarTech, Legrand listings), hssl.us (AddOn), LINK-PP (l-p.com), LightOptics (lightoptics.co.uk), allhdd.com (Cisco QDD-400G-DR4-S), seesuo.com / friendsoffice (Cisco QSFP-100G-SR4-S pricing).
- **Legal:** lawstreetmedia.com (Cisco/Ciena v. Wolon), aphnetworks.com (Cisco injunctions), ciplawyer.com (Dexon v. Cisco).
- **Hyperscale:** lightwaveonline.com (Voyager; Optical Scale-up Consortium), businesswire.com (consortium announcement), growthmarketreports.com (market), github equity-research notes (life-os, equity-watch, arsenal).
- **Incident/anecdote:** community.juniper.net (SFP-T inventory anecdote), network-switch.com (IOS XE 17.9.6 / CSCwm57734), theregister.com (Cisco clock-component advisory 2017, background), optcore.net (Magnuson-Moss vendor claim), FSL/MC-Custom-Worx PDFs (third-party explainers), dknconsulting.co.za (Cisco optical-module certification article).

## 11. Coverage check against the eight requested areas

1. **Coding mechanics:** §1 (SFF-8472 fields/checksums, SFF-8636 map/pages/passwords, CMIS layout, OUI validation, encryption-claim caveats, recoding limits) — complete.
2. **Vendor ecosystem (12 named):** §3.1–§3.12 — EDGE Optic (§3.7) and FluxLight company profile (§3.10) not found; noted as gaps. All others covered with products, dates, prices, URLs.
3. **FS BOX V4 and comparable programmers:** §2 — FS BOX V4, FLEXBOX, Solid Optics Multi Fiber Tool.
4. **Warranty models:** §4 — matrix, lifetime claims, advance replacement, exclusions, FS 5-year vs lifetime conflict, MTBF gap.
5. **Lock-in behavior (Cisco, Juniper, HPE/Aruba, Dell, Arista):** §5.1–§5.6 plus Extreme; support implications synthesized §5.7.
6. **Price gaps:** §6 — comparable 10G pair, within-third-party spreads, 100G series with non-comparability flags, 400G third-party prices with the Cisco-list gap flagged.
7. **Legal/gray areas:** §7 — counterfeiting vs compatible, Wolon/Cisco-Ciena litigation, Dexon, Magnuson-Moss/antitrust claim labels, nominative-use labeling.
8. **Enterprise/hyperscale practice:** §8 — Voyager, Optical Scale-up Consortium (2026-03-12), direct-manufacturer qualification pattern, explicit non-assertion on retail-brand use by hyperscalers.

**Final caveat:** every quantitative claim above is traceable to its cited URL; where a number, date, or attribution could not be verified in this pass it is labeled `[unverified]` or `[secondary]` and, where load-bearing, repeated in §9.

---
---

