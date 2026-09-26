---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/sources-and-urls
title: "Sources and URLs"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Baidu", "ByteDance", "China", "Cohere", "DeepSeek", "Falcon", "Groq", "MiniMax", "Moonshot", "Nvidia", "Poolside", "SGLang", "StepFun", "Z.ai"]
dates: ["2024-10", "2025-12", "2026-07"]
keywords: ["antitrust", "apache", "attribution", "benchmark", "benchmarks", "cohere", "cost", "deepseek", "distribution", "fine-tuning", "funding", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8627, 8662]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: 3fd9a04e7a6c5ecf1df2491abe033300ef0c129443a66649c2ea43d3585a7ba5
---

# Sources and URLs

1. 01.AI's pivot validates the post-DeepSeek economics: training foundation models is reserved for "bottomless balance sheets" while the money moves to fine-tuning, data infrastructure and sovereign deployments — the same sovereign pitch as Cohere's Aleph Alpha combination in §16 [DIRECTIONAL]. (S1)
2. The 01.AI figures (250M yuan audited revenue, 1.5B+ yuan orders, ~half recurring) are the strongest audited-figure set among the "AI Tigers" second tier in this file — but they describe an enterprise-software business, not a model-lab business [SECONDARY]. (S1)
3. Hong Kong's specialist-technology listing rules (pre-profit allowed) are now the structural enabler for the 2026–2027 Chinese AI IPO wave: Z.AI and MiniMax listed in 2026, with 01.AI, Moonshot and StepFun preparing [SECONDARY]. (S2, S3)
4. Baichuan-M3's vendor benchmarks (HealthBench 44.4, 3.5% hallucination) must not be compared against benchmarks run under different Artificial Analysis Index versions or vendor protocols — benchmark hygiene from the base document applies [DIRECTIONAL]. (S9)
5. The M3 training-data refusal (licensed/private corpora) plus the October 2024 inherited cutoff create a reproducibility ceiling: open weights, closed data, stale-ish knowledge — buyers should weigh those three facts together [DIRECTIONAL]. (S13, S9)
6. dots3-note's 512K context at 16B active makes it the longest-context open MoE in this file — but the preview label, open PRs for Transformers/SGLang integrations, and absent production API mean it is a research artefact, not a product [DIRECTIONAL]. (S15, S17)
7. Doubao's 180-trillion daily token volume and 49.5% MaaS share describe a distribution moat: ByteDance competes on API price (6/30 yuan) and workplace integration (Doubao Work), not on model licensing [SECONDARY]. (S23)
8. Step 3.7 Flash's independent test (1.5x cheaper per task vs 9x cheaper per token) is the cleanest illustration in this wave of why per-token price ≠ per-task cost — verbose outputs erase sticker-price advantages [SECONDARY]. (S31)
9. Hunyuan3D's split (open line stalled at 2.1; 2.5/3.0/3.1 API-only) is a cautionary pattern: "open-source commitment" claims in press copy do not map to weight releases [DIRECTIONAL]. (S34, S35)
10. Community research notes (HY3D-Bench, Hunyuan3D lineage tracking) are doing the versioning work Tencent's own documentation does not publish — cite them as [COMMUNITY], not as vendor fact . (S35)
11. Keep the 2025-vs-2026 boundary hard: Hunyuan3D 3.0 (Sept 2025), ERNIE-4.5-VL-28B-A3B-Thinking (2025) and Command A Reasoning (2025) are 2025 artefacts; presenting them as 2026 launches is a factual error [DIRECTIONAL]. (S32; §16 claim 26 covers the parallel 2025-boundary case)
12. Contradictions to preserve: Step 3.7 Flash release May 28 vs 29, 2026; 01.AI unicorn timeline 6 vs 8 months; dots3-note Terminal-Bench 75.1 is second-hand SemiAnalysis attribution, not verified independently [DIRECTIONAL]. (S27, S31, S8, S1, S16)


13. The community self-host bar for 500B+ models is dropping: AngelSlim's ~214-GiB GGUF builds for Hy4 Preview and community NVFP4 CUDA builds of openPangu Flash at ~56.9 GB resident — permissive licenses now attach to models whose weights were previously considered undeployable outside datacenters [DIRECTIONAL]. (S42, S78)
14. The Poolside deal's non-exclusive structure — $6B license + $1B investment + 109 team offers — plus NVIDIA's concurrent Groq ($20B) and Enfabrica (>$900M) license plays is becoming NVIDIA's standard expansion playbook; the "fiction of competition" antitrust critique (Bernstein, via secondary coverage) applies equally here [SECONDARY]. (S51)
15. The MiniMax/Zhipu IPOs (HK$8.54B combined, 1,837× oversubscription) prove the Hong Kong specialist-technology listing path is the AI Tigers' funding rail — the same rail 01.AI (S4 in this file), Moonshot (confidential, unverified) and StepFun (S68) are lining up for [SECONDARY]. (S69, S71, S68)
16. The Ling-3.0-Flash MIT attribution (community sources) superseding the base corpus's license-[UNVERIFIED] tag demonstrates the expansion-file discipline:  tags must be re-checked against newer evidence rather than carried forward blindly . (S59, S60)
17. ERNIE's three-way pricing conflict ($0.85/$3.40 vs $0.59/$2.65 vs $3/$12) plus no open weights means Baidu's numbers are not independently checkable — cite all three figures, never a blended one [DIRECTIONAL]. (S55, S56, S57)
18. Falcon-H1 Arabic's vendor-reported OALL figures (3B beating 70B-class systems) are leaderboard-specific; sovereign-language benchmark claims need independent replication before any procurement reading [DIRECTIONAL]. (S48)
19. Step 5's verbose-reasoning caveat (160M vs 92M median output tokens) repeats Step 3.7 Flash's independent-test lesson: per-token price ≠ per-task cost — cost analysis must use per-task spend [SECONDARY]. (S64, S31)
20. Qwen Image 2.1's same-day launch note (moved from Apache to a more restrictive license) is a reminder that the permissive-license trend is not monotonic — read each release's terms [SECONDARY]. (S65)
21. The Poolside-M.1 lag (announced December 2025, public release July 2026) and openPangu Pro's one-month weight delay show frontier-labs announcements now routinely trail staged releases — date every claim to the release it describes [DIRECTIONAL]. (S51, S44)

## Sources and URLs
- https://www.marktechpost.com/2026/09/20/stepfun-launches-step-5-preview/
- https://earlyterms.com/term/hy4-preview
- https://www.ai-all.info/en/ai-models/openpangu-2-0
- https://www.globenewswire.com/news-release/2026/07/21/3330818/0/en/Poolside-releases-Laguna-S-2-1-the-West-s-most-capable-open-weight-model.html
- https://www.tii.ae/news/tii-launches-falcon-perception-new-multimodal-ai-model-helps-machines-see-and-understand-world
- https://the-decoder.com/perplexity-ai-launches-new-ultra-fast-ai-search-model-sonar/
- https://www.intelligentliving.co/step-5-preview-cheapest-frontier-ai/


### New sources — expansion

