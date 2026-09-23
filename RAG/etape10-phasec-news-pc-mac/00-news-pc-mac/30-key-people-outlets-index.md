---
id: etape10-phasec-news-pc-mac/00-news-pc-mac/30-key-people-outlets-index
title: "30. Key People & Outlets Index"
domain: step-10-phase-c-pc-mac-news-february-september-2026
role: deep-dive
task: reference
actors: ["AMD", "Apple", "Intel", "Nvidia", "TSMC"]
dates: []
keywords: ["18a", "2nm", "agentic", "agents", "ai pc", "amd", "crescent island", "gpu", "gpus", "intel", "lpddr5x", "memory"]
source: docs/RAG/etape10_phaseC_news_pc_mac.md
source_anchor: ""
source_lines: [597, 648]
section: "Step 10 Phase C — PC & Mac News (February → September 2026)"
sha256: e45f70505e54c6749575846beaabdf80703438f0da174f1788343952fb7fd2f6
---

# 30. Key People & Outlets Index

## 30. Key People & Outlets Index

| Name | Role in 2026 PC/Mac news | Provenance |
|---|---|---|
| Lip-Bu Tan | Intel CEO; confirmed Nova Lake end-of-2026 on investor call | [secondary] |
| Pavan Davuluri | Head of Windows + Surface; confirmed no Windows 12, teased N1X hardware | [secondary] |
| Strauss Zelnick | Take-Two CEO; held GTA VI Nov 19, 2026 line through earnings | [secondary] |
| Sri Santhanam | Apple VP Silicon Engineering; M6 messaging | [secondary] |
| Zac Bowden | Windows Central; debunked Windows 12 rumor | [secondary] |
| kopite7kimi | NVIDIA leaker; RTX 60 → 2028 | [unverified] |
| Jaykihn (@jaykihn0) | Intel leaker; Celestial cancellation, Druid roadmap | [unverified] |
| Moore's Law Is Dead | YouTube; "Gaming Rubin" H1 2027 claim (Sept 14, 2026) | [unverified] |
| MEGAsizeGPU | Leaker; disputed H1 2027 RTX 60 | [unverified] |

---

## 31. Methodology Note

- All dates are event/publication dates as reported by cited outlets; "Last Crawl" metadata was used only to sequence undated items.
- Leak-sourced claims (GPU roadmaps, cancellations, unreleased specs) are tagged `[unverified]` even when reported by multiple outlets, since they trace to single anonymous sources.
- Vendor performance claims (TOPS, "4.3x AI", "1.8x graphics") are tagged `[official]`/`[vendor-reported]`, never presented as independent measurements; independent measurements (The Verge, Engadget, NeoTeo/Linux) are tagged `[independent]`.
- Sections marked "gap" in §16 were not researched further to avoid fabricating coverage; they are honest boundaries of this file.

---

## 32. Deep Dive — Intel in 2026 (Panther Lake → Nova Lake)

- **The 18A bet:** Panther Lake is the first high-volume Intel 18A product and the payoff of the "five nodes in four years" campaign; Intel framed it as reclaiming "the silicon crown" [secondary].
- **Panther Lake SKU stack (14 SKUs):** tops at 16 cores / 5.1 GHz; Core Ultra X9 388H-class parts (seen in Framework Laptop 13 Pro) pair Xe3 "Arc B390" graphics with a 50 TOPS NPU and 180 TOPS platform [secondary]/[independent].
- **Linux support:** Linux 7.3 added Panther Lake optimization worth ~3% performance, verified on a Framework 13 Pro [independent].
- **Agentic AI PCs:** Intel's 2026 framing pushed local autonomous agents running on the 50 TOPS NPU [secondary].
- **Nova Lake (Core Ultra 400, end of 2026):** new LGA-1954 socket; hybrid of Coyote Cove P-cores and Arctic Wolf E-cores; desktop parts up to 52 cores (16P+32E+4LP-E); bLLC cache on select SKUs; Xe3 graphics with Xe3P on high-end iGPUs; TDPs from 125W to 150W for flagships [official via secondary for timing; details unverified].
- **What Nova Lake replaces:** Arrow Lake (Core Ultra 200S) on LGA-1851; motherboard change means a new-platform upgrade cycle for enthusiasts [secondary].
- **The GPU question:** Battlemage (Arc B580 $249) crossed 1% discrete share per Jon Peddie Research in late 2025 [secondary]; beyond that, no Intel gaming GPU in 2026 [secondary]; Xe3P Celestial dGPUs canceled per leaker [unverified]; Xe3P silicon redirected to Crescent Island AI GPUs sampling H2 2026 (160GB LPDDR5X, OAM, PCIe 6.0) [secondary]; Xe4 Druid "up in the air," one leak pointing to late 2027 [unverified].
- **Intel's own AI PC hardware:** not in research — the company's 2026 story was the Panther Lake platform and OEM designs, not Intel-branded devices — gap noted.

---

## 33. Deep Dive — AMD in 2026 (AM5 Longevity → Zen 6)

- **AM5 through 2029:** AMD confirmed at Computex 2026 (via TechTimes) that the AM5 platform will be supported through 2029 [secondary] — one of the longest desktop socket commitments in the industry.
- **Zen 6 roadmap (official portal data, Aug 27, 2026):** Zen 6 CPU families confirmed: Olympic Ridge (AM5 desktop, H1 2027), Gator Range (mobile), EPK, Medusa, Venice [secondary].
- **Olympic Ridge specs (leak consensus):** TSMC N2P (2nm-class); up to 24 cores/48 threads on desktop AM5; 12-core CCXs with 48MB L3; ~10%+ IPC uplift target over Zen 5 [secondary].
- **Medusa (mobile halo):** up to 24C/48T, 96MB L3, RDNA 5 iGPU, LPDDR6 — positioned against Apple-class mobile performance [unverified].
- **Board partners ready early:** Biostar exhibited Zen 6-ready motherboards at Computex 2026 [secondary].
- **2026 mobile:** Ryzen AI 400 "Gorgon Point" (Zen 5, 4nm) in 8C/16T and 16C configs, XDNA 2 NPU, RDNA 3.5 iGPU, shipping 1H 2026 — AMD's direct Panther Lake rival [secondary].
- **RDNA 5 (GFX13):** taped out on TSMC N3P; AT0–AT4 chip stack with AT0 flagship; memory controller "MMHub" redesign for better bandwidth; Project Amethyst collaboration with Sony [secondary]; board-partner consensus: late 2027–early 2028 [secondary]; HDMI 2.2 support rumor [unverified].
- **UDNA unification:** AMD merging RDNA/CDNA into UDNA (Gfx13) for gaming + AI, per TechPowerUp [secondary].
- **The value question:** AMD's 2026 positioning was platform longevity (AM5→2029) and mid-range value (RX 9070 XT class) while the flagship race paused [secondary].

---

