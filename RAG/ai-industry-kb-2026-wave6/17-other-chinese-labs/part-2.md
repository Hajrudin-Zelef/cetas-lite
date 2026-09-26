---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/part-2
title: "§17. Other Chinese Labs (part 2)"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "Hugging Face", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "StepFun", "Z.ai", "vLLM"]
dates: ["2024-10", "2026-05", "2026-06-23", "2026-08-14"]
keywords: ["agent", "apache", "attention", "benchmarks", "claude", "consumer", "cost", "deepseek", "distillation", "fp8", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8326, 8358]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: 9037781d4ccdc5d40dfc5c57935c20800e0a2140bff01d957ac61c98bcc88dde
---

# §17. Other Chinese Labs (part 2)

1. 01.AI abandoned the frontier-model race: Kai-Fu Lee told Bloomberg at the World AI Conference in Shanghai that only "a handful of companies with essentially bottomless balance sheets could still justify the cost of building models from scratch," and 01.AI now builds enterprise software on top of open models from others [SECONDARY]. (S1, S2)
2. 01.AI now fine-tunes or customizes Chinese open-weight models such as DeepSeek, Alibaba's Qwen and Z.AI's GLM rather than training foundation models from scratch [SECONDARY]. (S1, S3)
3. 01.AI's main product, Boss AI, creates isolated data pools for clients supporting instant querying and visualization; buyers in regulated industries want on-premises deployment, driven by data security [SECONDARY]. (S1 — single secondary coverage)
4. Lee describes the approach as building "the Palantir of China" — top-down sales of customized AI systems to governments and large enterprises instead of consumer apps or model benchmarks [SECONDARY]. (S1, S3)
5. 01.AI reported audited revenue of 250 million yuan (about $34 million) for 2025; contract orders exceeded 1.5 billion yuan (roughly $207 million) as of May 2026, with management setting a 2 billion yuan target for the full year [SECONDARY]. (S1, S6)
6. Recurring subscription revenue already accounts for nearly half the 2026 contract volume; Lee said about half of 01.AI's business comes from outside China [SECONDARY]. (S1, S3)
7. 01.AI targets a Hong Kong IPO in 2027, after the fiscal year closes in December; Lee disclosed a pre-IPO financing round expected to conclude around the time 01.AI publishes its first annual results [SECONDARY]. (S2, S3)
8. 01.AI is unwinding its offshore holding structure — a prerequisite for a smoother overseas listing; Moonshot dismantled a similar structure in May 2026 to clear its own path to a Hong Kong debut [SECONDARY]. (S2, S3)
9. Hong Kong's specialist technology listing rules allow AI companies to go public before becoming profitable, subject to operating history, R&D spending, sophisticated-investor backing and a credible path to commercialization [SECONDARY]. (S3 — single secondary coverage)
10. Lee founded 01.AI in 2023; sources differ on whether it reached a $1B+ valuation within six months (caproasia) or eight months (webpronews/techinasia) — preserved as a contradiction; Alibaba's cloud unit was a backer [SECONDARY]. (S8, S1)
11. Among the half-dozen "AI Tigers," Z.AI (Zhipu) and MiniMax listed on the Hong Kong exchange in 2026 while StepFun, DeepSeek and Moonshot are preparing their own debuts [SECONDARY]. (S2 — single secondary coverage)
12. One source claims Lee said 01.AI aims to become profitable in 2026 and will issue 20 million additional stock options to staff — single-source via KR Asia/Huxiu [UNVERIFIED]. (S6)
13. Baichuan-M3-235B is built on the Qwen3-235B-A22B base — Baichuan Intelligent Technology's own staff confirmed on Hugging Face discussions that M3 inherits Qwen3's October 2024 pretraining knowledge cutoff [VENDOR]. (S9, S10)
14. SPAR stands for Segmented Pipeline Reinforcement Learning — technically "Step-Penalized Advantage with Relative baseline": it decomposes clinical workflows into four stages (history taking, differential diagnosis, laboratory testing, final diagnosis), each with independent rewards plus process-level rewards for credit assignment [VENDOR]. (S11, S10)
15. Baichuan's Fact-Aware RL integrates factual verification directly into the RL loop with an online hallucination detection module validating medical claims against authoritative evidence in real time, plus dynamic reward aggregation balancing task learning against factual constraints [VENDOR]. (S11, S12)
16. Baichuan claims a 3.5% hallucination rate for M3 — "industry's lowest" — and claims it beats GPT-5.2 even in a tool-free setting; these are vendor claims, not independent measurements [VENDOR]. (S9, S12)
17. M3 scores HealthBench-Hard 44.4 and HealthBench Total 65.1 per vendor reporting, ranking first and 28 points above M2, and SCAN-bench first across all three dimensions with +12.4 points over second place in Clinical Inquiry — vendor-reported figures [VENDOR]. (S9, S13)
18. M3 training uses a three-stage multi-expert fusion paradigm (Domain-Specific RL → Offline Distillation → MOPD), Gated Eagle3 speculative decoding (claimed 96% speedup) and W4 quantization (26% memory) [VENDOR]. (S11 — single vendor source coverage)
19. W4 quantization is claimed to allow consumer-GPU deployment on 2× RTX 4090 [SECONDARY]. (S9 — single secondary coverage)
20. The HF discussion thread shows community requests to release M3's training data; Baichuan staff declined, citing licensed sources, curated private corpora and human-created annotations that make open-sourcing legally and ethically difficult [VENDOR]. (S13 — single vendor source coverage)
21. HF's open-source team invited Baichuan to host SCAN-bench and HealthBench-Hallu datasets on the Hub after the paper was featured as a daily paper (https://huggingface.co/papers/2602.06570); Baichuan's README said SCAN-bench "will be open-sourced soon" [VENDOR]. (S14, S12)
22. A community-built Medical-Reasoning-SFT-Baichuan-M3-235B dataset on HF corroborates the base (Qwen3-235B-A22B), Apache 2.0 license, 44.4% HealthBench-Hard and 3.5% hallucination rate — community evidence, not independent of vendor [COMMUNITY]. (S10 — single community source coverage)
23. dots.studio (Xiaohongshu's AI division) released dots3-note preview on August 14, 2026 — announcement timestamp 02:00:31 UTC — as the first open-weight model in the dots3 family under Apache 2.0 [SECONDARY]. (S15, S16)
24. dots3-note preview is a multimodal MoE: 280B total parameters, 16B active, 512K-token context; inputs are text, image, video and audio, output is text [VENDOR]. (S17, S15)
25. Architecture detail from the official card: 1 dense + 45 MoE layers, hidden size 5120, 256 routed + 1 shared experts with top-8 routing, MTP 1 shared layer 1.13B, 13 DSA + 33 SWA attention layers, vocab 152K [VENDOR]. (S17 — single vendor source coverage)
26. The vision encoder is an MoE ViT with 7B total / 1.2B activated parameters; the audio encoder is a dense 800M model; supported precision BF16 and FP8 [VENDOR]. (S17 — single vendor source coverage)
27. The recommended FP8 deployment uses one eight-GPU node via vLLM (available on main); Transformers and SGLang source integrations were tied to open pull requests at release [SECONDARY]. (S15 — single secondary coverage)
28. One source reports a Terminal-Bench 2.1 score of 75.1 attributed to SemiAnalysis data, 4.9 points ahead of leading U.S. open-weight models, and cites a "TEMPO" RL method — second-hand reporting, treat as [DIRECTIONAL]. (S16, S18)
29. OpenRouter lists dots3-note preview as free ($0 prompt/completion) with rate limits; its FAQ confirms text+image input (not video/audio at the API level), tool calling and structured outputs [SECONDARY]. (S19 — single secondary coverage)
30. dots3-note is described as the "most lightweight member" of the dots3 family, implying larger family members are planned [VENDOR]. (S17 — single vendor source coverage)
31. ByteDance's Volcano Engine released Doubao-Seed-2.1 Pro and Doubao-Seed-2.1 Turbo at the June 23, 2026 Volcano Engine FORCE conference, with APIs live on VolcanoArk the same day [SECONDARY]. (S20, S21)
32. Doubao 2.1 Pro pricing: 6 yuan per million input tokens, 30 yuan per million output tokens, cache-hit pricing as low as 1.2 yuan; Volcano Engine claims total cost of ownership nearly 80% lower than Claude Opus 4.6; Turbo costs half the Pro price [SECONDARY]. (S22, S23)
33. ByteDance claims Doubao 2.1 Pro leads Terminal Bench 2.1, SWE-Pro and SciCode for code, and OSWorld, MobileWorld and MMMU-Pro for agent/multimodal tasks, outperforming Claude Opus 4.6 — vendor claims, not independent [VENDOR]. (S21, S22)
