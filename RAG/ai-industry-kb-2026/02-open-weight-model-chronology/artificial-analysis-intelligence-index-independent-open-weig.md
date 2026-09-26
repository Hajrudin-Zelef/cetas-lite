---
id: ai-industry-kb-2026/02-open-weight-model-chronology/artificial-analysis-intelligence-index-independent-open-weig
title: "Artificial Analysis Intelligence Index — independent open-weight ranking snapshots"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "Meta", "Mistral", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: ["2025-01", "2025-10", "2026-03", "2026-05-29", "2026-07-17", "2026-07-30", "2026-09", "2026-09-04", "2026-09-15"]
keywords: ["open-weight", "agentic", "agents", "astra", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "cyber", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [791, 819]
section: "2. Open-Weight Model Chronology"
sha256: 7067e634d359dc8dc8db84230f2bce1ebfa5d25458039faefa074a19a8868e6a
---

# Artificial Analysis Intelligence Index — independent open-weight ranking snapshots

- **Epoch AI Capabilities Index:** average open-weight lag ~4 months / 8 ECI points (May 29, 2026 update; 3.5 months / 7 points in October 2025 — the gap widened slightly as GPT-5.x pulled ahead; convergence is directional, not monotonic). Six months under a stricter criterion.
- **Mozilla State of Open Source AI, 2nd ed. (September 15, 2026):** Chinese open-weight models 4.4 months behind the US frontier (read off the Artificial Analysis Intelligence Index; Kimi K3 within 3 index points of Anthropic's Fable 5 at ~30% of the cost). Mozilla's practical conclusion: most organizations should use open models as the default for the majority of work — paying for the frontier buys "a four-month head start at five times the cost."
- **Arena AI crowdsourced Elo (September 2026):** closed–open gap 29 points — Claude Opus 5 Max 1505 vs Kimi K3 trailing ~30; in January 2025 the gap had briefly hit zero (parity). Human-preference Elo rewards polish/instruction-following/reliability, favoring closed models more than task benchmarks do.
- **AISI open-weight gap report (July 17, 2026):** 4–7 months on cyber benchmarks — 4 months on narrow tasks, up to 7 on autonomous multi-step attack ranges. Once weights are distributed they cannot be recalled; DeepSeek V4-Pro's refusals on cyber tasks were overcome by simply retrying failed requests.
- **Where the gap remains vs. where it closed (Sept 2026):** coding benchmarks — closed or open-weight leads (Kimi K3 #1 Frontend Code Arena; GLM-5.3 #1 on Ed-o-meter); agentic computer use — near-parity (Qwen3.8-Max 86.1 vs Fable 5's 85.0 on OSWorld-Verified [VENDOR]); hardest reasoning — closed leads (Fable 5 HLE 53.3 vs Qwen3.8-Max 43.6); long-horizon autonomous tasks — closed leads; video/interleaved multimodality — closed leads; 1M-context coherence — closed leads (Claude claimed 76% at 1M, independently unverified); adversarial robustness — closed leads decisively.

### Artificial Analysis Intelligence Index — independent open-weight ranking snapshots

- **v4.1 scale (summer 2026):** GLM-5.3 60 (9th of 186/187 models, alongside proprietary frontier); Kimi K3 57 (July 30, 2026 — #1 among open-weight models at the time); V4-Flash-0731 50 (3rd among open weights); Muse Spark 1.2 57 (vendor-adjacent) vs 54 (independent evaluation — methodology difference, flag not average); GLM-5.2 51 (at release, open-model lead).
- **v4.3 scale (rebased September 4, 2026 — GPQA-Diamond dropped as solved, held-out evals added, 40% private-test weighting; absolute scores fell and prior-version numbers are no longer comparable):** GLM-5.3 ≈ 45, Kimi K3 ≈ 44, GLM-5.3-Flash 42, Qwen3.8 (2.4T A95B) 40, DeepSeek V4 Pro 0813 36 — vs closed frontier Claude Fable 5.1 = GPT-6 Astra = 53. Qwen3.8-Max is 4th among open weights, not the leader ("dominance" is vendor-reported and benchmark-specific).
- **Size class (v4.3-era, independent):** Qwen3.8-27B 52 (size-class leader), Qwen3.6-27B 38, Muse Glimmer 30B 35.

### Hugging Face distribution scale (2026)

- 2.4M+ public models, 730k+ datasets, ~1M Spaces, 50,000+ organizations, 11M users (Hugging Face's March 2026 DOE filing). The second million models landed in roughly a third of the time the first took.
- Transformers at 113M+ monthly downloads (late 2025); gpt-oss-120b at 4.3M+ downloads — among the most-downloaded open-weight models ever.
- Every major 2026 open-weight release lands on HF within days: Qwen3.8-Max (Aug 12), Kimi K3 (July 27 — "appeared on Hugging Face at 2:00 AM Beijing time"), DeepSeek V4 family, GLM-5.3-Flash (Aug 26).
- Curation signal vs raw count: Epoch tracks 1,340 open-weight models (964 language models); ATOM ~1,500 mainline open language models; 95 of 170 models on the AA Intelligence Index are open-weight — against 2.4M+ total HF models, model *selection* is now the hard problem.

### Key vendor-reported benchmark points (all self-reported unless marked independent)

- **Qwen3.8-Max (Alibaba, self-reported):** Terminal-Bench 2.1 86.6, PaperBench 93.0, GPQA Diamond 92.6, HLE 43.6, OSWorld-Verified 86.1, Agents' Last Exam 52.4, DeepSWE 1.1 56.6 (vs 21.6 for Qwen3.7-Max), ERQA 77.8 (Fable 5: 70.0), MobileWorld 77.8, JobBench 53.4 (vs 31.3).
- **Kimi K3 (Moonshot, self-reported unless marked):** Frontend Code Arena 1,679 Elo — #1 (independent, crowdsourced); Terminal-Bench 2.1 88.3; FrontierSWE 81.2; SWE-Bench Verified 76.8%; GPQA-Diamond 93.5; AIME 2025 96.1%; BrowseComp 91.2; HLE-Full 43.5; GDPval-AA v2 1687 vs Opus 4.8's 1600.
- **DeepSeek V4 family (DeepSeek, self-reported):** LiveCodeBench 93.5, GPQA Diamond 90.1, SWE-Verified ~80, Codeforces 3206; V4-Flash-0731 Terminal-Bench 2.1 82.7 (up from 61.8 for the April preview); DSA attention: ~27% inference compute and 10% KV cache vs V3.2 at the same length.
- **DeepSeek V4.1-Flash (DeepSeek, self-reported):** GPQA Diamond 90.9, Terminal-Bench 2.1 90.6, Codeforces 3471, DeepSWE v1.1 74.2% (harness swing 65.5–74.2%); concedes Terminal-Bench 3.0 30.0 / 4.0 31.2 / HLE 36.8.
- **GLM-5.3 (Z.ai, self-reported):** Terminal-Bench 3.0 28.3 (from 4.6), DeepSWE v1.1 66.9, Z.ai Code Bench 34.5% at ~75K output tokens (vs Opus 4.8's 29.5% at ~120K), CyberGym 84.5%, ExploitBench 54.4% (from 24.4%); 2,436 vulnerabilities found across 269 projects.
- **Muse Glimmer 30B (Meta, self-reported):** SWE-Bench Pro 51.2%, AIME 94.7, GPQA 83.5, MCP Atlas 75.5 (vs 54.2 Gemma4-31B / 62.5 Qwen3.6-27B); trails Qwen3.6-27B on Terminal-Bench 2.1 (51.7 vs 60.7) and OSWorld-Verified.
- **Mistral Medium 3.5 (Mistral, vendor-announced):** SWE-Bench Verified 77.6% (ahead of Qwen3.5 397B and Devstral 2), τ³-Telecom 91.4.

