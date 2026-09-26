---
id: etape6-phasec-optics-cabling/01-wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra/9-conflicts-unverified-claims-and-gaps-register
title: "9. Conflicts, unverified claims, and gaps register"
domain: wave-12-third-party-coded-optics-eeprom-msa-mechanics-progra
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Cohere", "Meta", "Microsoft", "Nvidia", "OpenAI", "United States"]
dates: ["2021-01-28", "2025-05-28", "2026-01-09", "2026-03-12"]
keywords: ["amd", "asic", "compute", "hyperscaler", "nvidia", "optics", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2230, 2252]
section: "Wave 12 — Third-party coded optics: EEPROM/MSA mechanics, programmers, vendor ecosystem, lock-in, warranty, legal"
sha256: ae31c5287bb13b0cb00066756b24621878df008f2ca294b12770097a2290c143
---

# 9. Conflicts, unverified claims, and gaps register

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

