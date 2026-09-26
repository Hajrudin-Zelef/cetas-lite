---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/part-5
title: "§17. Other Chinese Labs (part 5)"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Ant", "DeepSeek", "Google", "MiniMax", "Moonshot", "OpenAI", "StepFun", "Z.ai"]
dates: ["2026-08", "2026-09-08"]
keywords: ["benchmark", "consumer", "cost", "deepseek", "fp8", "funding", "gemini", "gguf", "glm", "gpt-5.6", "int4", "ipo"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8415, 8435]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: c689b070e2b06839b28bc653a423c81942bdd22b8f011b0a680d07a2959995e0
---

# §17. Other Chinese Labs (part 5)

89. Ling-3.0-Flash engineering detail beyond the base specs: cluster-level hierarchical caching cuts TTFT 60–80% on long conversations; built-in MTP layer; context scalable to 1M per the community catalog; reasoning enabled by default with a per-request toggle [COMMUNITY]. (S59, S60)
90. Artificial Analysis Intelligence Index: Ling-3.0-flash scores 38 — "smartest open model under 124B total parameters", at parity with Qwen3.6 27B and trailing DeepSeek V4 Flash (52); AA Omniscience hallucination rate improved from 97% to 44% (refusal behavior improved) [SECONDARY]. (S58 — single secondary coverage)
91. First-party API pricing: $0.075/M input, $0.22/M output (inclusionAI API and DeepInfra); checkpoints published in BF16, FP8 (E4M3), INT4 (W4A16) and MXFP4 [COMMUNITY]. (S59, S60)
92. September 8, 2026 — Ant open-sourced Ling-3.0-flash-Fin for real-world financial workflows (retrieval, research reasoning, valuation modeling, report generation) alongside the FinFIRST benchmark built with CICC: 123 expert-level tasks, 701 atomic evaluation criteria, 12,300 rubric points; plus Ling-3.0-tiny (7.9B total / 1.3B active) that runs fully locally without cloud dependency [VENDOR]. (S61 — single vendor source coverage)
93. Community NVFP4 GGUF builds of Ling-3.0-flash (72.3 GB): the AD-NVFP4 variant reports mean KLD 0.05363 / 94.86% top-1 match versus stock 0.05602 / 94.72% — same format, encoder-only difference [COMMUNITY]. (S62 — single community source coverage)
94. Step 5 Preview new detail: 600B/27B implies ~4.5% active density; the API is live at platform.stepfun.ai [SECONDARY]. (S63 — single secondary coverage)
95. CONTRADICTION PRESERVED — Step 5 input modalities: most secondary sources report text+image, while MarkTechPost adds video input; the conflict is preserved [SECONDARY]. (S63, S64)
96. Artificial Analysis Intelligence Index 44 (v4.3.2): tied with Kimi K3 (max) at ~65% lower cost per task; ~99.8 tok/s output; $0.71 cost per Index task; ranked 27th of 653 on one board and 24th of 200 on another — different boards, both figures preserved [SECONDARY]. (S63, S67)
97. techarcade comparison: Step 5 Preview at AA 44 versus GPT-5.6 Sol at 47 max effort (44 when Sol's reasoning effort is dialed down to match Step 5's speed), at ~7.4× cheaper output ($2.70 vs $20/M) [SECONDARY]. (S66 — single secondary coverage)
98. Verbose-reasoning caveat: MarkTechPost notes Step 5 generated 160M output tokens on the index run versus a 92M median — verbosity eats part of the per-token savings; the article was already cited in the base §17 sources [SECONDARY]. (S64 — single secondary coverage)
99. Bad Signal's filing discipline: treat October 15 as StepFun's schedule, not a completed release; Artificial Analysis lists the model as proprietary today [SECONDARY]. (S65 — single secondary coverage)
100. StepFun reportedly raised $3.2B across two rounds in 2026 — single-source (easternherald/Pandaily syndication); treat as an unverified funding claim [UNVERIFIED]. (S68)
101. orcarouter comparison: Step 5 Preview (AA 44) versus Gemini 3.1 Pro Preview (AA 30); price $1.00/$2.70 vs $2.00/$12.00; TTFT 2.96 s vs 35.72 s; Step 5 open weights scheduled October 15, Gemini proprietary throughout [SECONDARY]. (S67 — single secondary coverage)
102. January 8–9, 2026 — Zhipu AI debuted Thursday on HKEX (2513), MiniMax Friday (HKEX 0100): Zhipu raised HK$4.35B (~$552M) at ~$6.6–7.4B valuation; MiniMax raised HK$4.8–5.54B (~$619–620M) at HK$165 per share [SECONDARY]. (S69, S70)
103. CONTRADICTION PRESERVED: MiniMax closed day one at HK$345 (+109.1% per one source, +109.9% per another), valuing it at HK$103–106.7B (~$13.2–13.7B); public offer oversubscribed 1,837×, international tranche 37×; raised figure stated as HK$4.8B by Reuters and HK$5.54B by Global Times depending on whether the offer extension is included — preserve both, do not average [SECONDARY]. (S70, S71)
104. Zhipu opened +3.3% and closed +13% on day one; together the two IPOs raised HK$8.54B (~$1.09–1.2B) [SECONDARY]. (S74, S69)
105. Business profiles: MiniMax (founded early 2022 by ex-SenseTime Yan Junjie) is consumer-facing — Hailuo AI video and Talkie apps; 200M+ users worldwide; ~70% of revenue overseas; 2024 revenue $30.5M. Zhipu (Tsinghua roots) is enterprise/government — flagged by OpenAI as gaining government contracts; a $3/month coding assistant; 2024 revenue $44.7M at a loss [SECONDARY]. (S70, S74)
106. MiniMax claims the four-year founding-to-listing interval as a global AI-industry record [SECONDARY]. (S71 — single secondary coverage)
107. August 2026 — MiniMax became Stock-Connect eligible (August 6, +22.8% to HK$311.60 that day); mainland investors poured HK$10.6B (~$1.4B) into the stock in August, building a 9.7% stake — outpacing net buying of Alibaba and Tencent [SECONDARY]. (S73 — single secondary coverage)
108. MiniMax is studying a Shanghai STAR dual listing (A+H) and has hired Citic Securities; Zhipu plans to allocate 70% of IPO proceeds to GLM-4.7 R&D despite H1 2025 R&D of $228M against roughly 8× its revenue — Reuters-sourced [SECONDARY]. (S72, S71)

