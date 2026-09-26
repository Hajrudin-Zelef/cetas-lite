---
id: etape6-phasec-optics-cabling/00-front-matter/3-2-latency-power-trade-offs
title: "3.2 Latency & power trade-offs"
domain: front-matter
role: reference
task: reference
actors: ["Meta", "Nvidia"]
dates: []
keywords: ["latency", "asic", "cpo", "dpo", "dsp", "lpo", "npo", "nvidia", "optics", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [368, 392]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 30adf000a2d14250f45b8e5e872f5747b4ddf84c0d8561eb9de3f28f3e052b10
---

# 3.2 Latency & power trade-offs

#### 3.1.2 Linear-drive Pluggable Optics (LPO)
- LPO removes digital processing units (DSP and CDR) from the module, creating a purely analog "linear direct-drive" optical link [secondary: https://naddod.medium.com/optical-interconnect-technology-analysis-lpo-npo-cpo-bd9b3488fb10]. Concept first proposed by MACOM and NVIDIA in 2022 [secondary: same].
- Architecture: TX side uses a high-linearity driver chip driving the optical modulator; RX side uses a high-linearity transimpedance amplifier (TIA) (+ optionally a linear equalizer / CTLE for input equalization). No DSP, no retiming, no clock recovery. Signal equalization and compensation are done by the host switch/xPU SerDes [vendor-reported: https://blog.semtech.com/ai-date-center-basics-what-is-linear-pluggable-optics-lpo; official: https://www.oiforum.com/wp-content/uploads/OIF_PLL_Demo_Eoptolink_OFC2024.pdf].
- Eoptolink (OFC 2023 launch): "800G LPOs are designed without DSPs or CDRs"; relies on the host ASIC's native 112G PAM4 SerDes equalization [official: https://www.eoptolink.com/news?start=15].
- LPO is NOT passive: it still contains active analog components (linear driver + TIA). It is also NOT a universal replacement for retimed modules — high-loss electrical channels, strong reflections, or deployments needing maximum TX FIR flexibility may still favor fully retimed or hybrid approaches [vendor-reported: https://blog.semtech.com/ai-date-center-basics-what-is-linear-pluggable-optics-lpo].
- Related variants:
  - **LRO (Linear Receive Optics), also "half-retimed"/RTLR**: keeps a (simplified) DSP on the transmit side, linear receive path only. Semtech reports RTLR/LRO at ~16W today (for 200G-class links — see §3.2 ambiguity flag) [vendor-reported: https://blog.semtech.com/ai-date-center-basics-what-is-linear-pluggable-optics-lpo]. AscentOptics puts 800G LRO at ~9–12W, ~25% lower than DSP baseline [secondary: https://ascentoptics.com/blog/800g-power-consumption/].
  - **ACC-MSA**: Feb 2026, a separate new MSA (MACOM/Semtech co-chaired) for linear active copper cables with integrated linear equalizers — same "linear, no-DSP" philosophy extended to copper [secondary: https://convergedigest.com/new-industry-msa-targets-low-power-copper-interconnect-for-800g-and-1-6t/].

#### 3.1.3 NPO (Near-Packaged Optics) and CPO (Co-Packaged Optics)
- **CPO**: integrates the optical engine directly onto the same package/substrate as the switch ASIC, shrinking electrical trace length from 10–30 cm (pluggable) to <1 cm, and in NVIDIA's implementation from ~10 cm to <1 cm, cutting signal loss from ~22 dB to ~4 dB [secondary: https://dev.to/lsolink/everything-you-need-to-know-about-800g16t-optical-transceiver-and-co-package-module-1m6o]. Shorter copper trace = less attenuation = less compensation circuitry = less power [secondary: https://momoview.com/blog/en/posts/co-packaged-optics-cpo-silicon-photonics-industry-analysis-ai-interconnect-bottleneck-2026/].
- **NPO**: optical engine placed next to (but not inside) the ASIC package on the same substrate — a transitional architecture preserving pluggable-like serviceability while improving power efficiency [secondary: https://momoview.com/blog/en/posts/co-packaged-optics-cpo-silicon-photonics-industry-analysis-ai-interconnect-bottleneck-2026/; https://finance.biggo.com/news/uei-yp4BOLsyMWM01Hme].
- Where the "DSP function" lives in each architecture:
  - DSP module: in-module DSP does CDR/equalization/retiming (both directions).
  - LPO/LRO: in-module DSP removed (LPO) or halved (LRO); host SerDes performs equalization/FEC; module is analog-only [official/vendor: OIF demo PDF; Semtech blog].
  - CPO/NPO: very short electrical channel reduces need for heavy equalization; CPO engines typically use linear or lightly-DSP'd drive adjacent to the switch SerDes; per LightCounting, "removing the DSPs saves power, but more complex SerDes are needed to make direct drive possible" [independent: https://www.lightwaveonline.com/home/article/55141192/lpo-msa-achieves-multi-vendor-interoperability].
- **Optical Scale-Up Consortium (OFC 2026)**: new MSA defining an open AI scale-up infrastructure spec supporting pluggable, on-board, AND co-packaged optics; OCI GEN1 (4λ × 50G NRZ, 200G/dir), GEN2 (400G/dir BiDi), roadmap to 3.2T/fiber. Meta is a participant [secondary: https://www.lightwaveonline.com/home/article/55365387/ofc-2026-optical-scale-up-consortium-sets-path-for-an-open-ai-infrastructure-specification].

### 3.2 Latency & power trade-offs

#### 3.2.1 Latency figures
- DSP-based: end-to-end optical module latency ~100 ns in the "traditional DPO solution" [secondary: https://medium.com/@aicplight888/from-400g-to-1-6t-lpo-technology-gains-traction-in-optical-transceivers-84d11ca8181e — vendor-adjacent Medium post, treat as [vendor-reported/marketing]].
- LPO: end-to-end delay "less than 10 ns" (AICPLIGHT) [secondary/vendor-reported, same source]; "<2 ns" for 800G-FR4-class LPO [secondary: https://dev.to/lsolink/everything-you-need-to-know-about-800g16t-optical-transceiver-and-co-package-module-1m6o — blog, weak credibility]; Eoptolink 800G LPO family: "low latency less than 1 ns" [vendor-reported: https://www.oiforum.com/wp-content/uploads/OIF_PLL_Demo_Eoptolink_OFC2024.pdf]. **Conflict note: published LPO latency claims range from <10 ns to <1 ns across sources — no single independent measured figure found; treat all as vendor/marketing claims.**
- FS 800G LPO: "significantly reduced latency" vs DSP modules (no numeric value given) [vendor-reported: https://aithority.com/machine-learning/fs-launches-800g-lpo-module-a-power-efficiency-and-latency-optimized-solution-for-ai-hpc-data-centers/].

