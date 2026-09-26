---
id: ai-industry-kb-2026-wave6/07-minimax/minimax-m3-vendor-benchmark-table-all-vendor
title: "MiniMax-M3 vendor benchmark table (all [VENDOR])"
domain: minimax
role: deep-dive
task: benchmark
actors: ["AMD", "Alibaba", "EU", "MiniMax", "Nvidia", "OpenAI", "United States"]
dates: ["2026-01-09"]
keywords: ["benchmark", "amd", "attribution", "blackwell", "ipo", "license", "mcp", "moe", "nvfp4", "pricing", "research", "revenue"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3366, 3400]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: c9eda847fb61f8b6a3aa7da7c7afec2c53f493412e47d6f15d9c68f68b04d852
---

# MiniMax-M3 vendor benchmark table (all [VENDOR])

### MiniMax-M3 vendor benchmark table (all [VENDOR])
| Benchmark | M3 score |
|---|---|
| SWE-bench Verified | 80.5 |
| SWE-bench Pro | 59.0 |
| Terminal-Bench 2.1 | 66.0 |
| MCP Atlas | 74.2 |
| OSWorld-Verified | 70.06 |
| BrowseComp | 83.5 |
| KernelBench Hard | 28.8 |

[SECONDARY reporting vendor] (inferencex-app; codingfleet.com; pasqualepillitteri.it).
- Cross-model hygiene: M3's Terminal-Bench **2.1** (66.0) must not be compared with GPT-5.5's Terminal-Bench **2.0** figure — a comparison source makes exactly this error and flags it [SECONDARY] (codingfleet.com).
- SWE-bench Verified 80.5 vs Pro 59.0: different suites, never merged [DIRECTIONAL].

### MiniMax-M3 architecture figures
- 60 layers (3 dense + 57 MoE); 128 routed experts, top-4/token, 1 shared; sparse Top-16 blocks × 128 block size, 4 index heads, 128 index dim; 7 MTP modules; RoPE on half of each 128-dim head; swigluoai (α 1.702, limit 7.0); Gemma-style RMSNorm [SECONDARY] (inferencex-app config reading).
- ~428B total / ~23B active; 1M context (512K guaranteed-usable floor [UNVERIFIED, single source]); MXFP8 ~440 GB (Blackwell/AMD), BF16 (Hopper/H200) [SECONDARY].

### MiniMax-M3 pricing
- Standard: **$0.60/M input / $2.40/M output / $0.15/M cached input**; launch promo **$0.30/$1.20** (promo cached input $0.06) [SECONDARY] (aitoolgrade.com; codingfleet.com).

### MSA speedup figure sets (setup-labelled)
- Production diagram: **9.7× / 15.6×**; launch card (rounded): **9× / 15×**; 109B research model: **14.2× / 7.6×** [SECONDARY] (inferencex-app; morphllm.com; techtimes.com).

### MiniMax H3 figures
- 33B params; 4–15s clips @ up to 2K/24fps + native stereo audio; 6 aspect ratios [SECONDARY].
- API: $0.13/s 2K ($7.80/min); $0.09/s 768p (closed beta); 768p = 8/13 of 2K rate; ref audio free; 5 ref images free then $0.04; ≤9 images / ≤3 videos (2–15s, ≤15s total) / ≤3 audio / 12 files / prompt ≤7,000 chars [SECONDARY].
- Independent price comparison: ~$7.80/min vs ~$20–22/min Kling 3.0 / Seedance 2.0 @1080p [SECONDARY] (justbeingresourceful.com).
- Local perf: ~175s per 10s clip (NVFP4, RTX 5090 32GB); community RTX 3060 renders [SECONDARY] (runaihome.com).
- License: MiniMax Community License; commercial use under **$20M revenue** + prominent attribution; **US/EU/UK/SK local-deployment exclusion reported** [SECONDARY, single source for territories] (runaihome.com; huggingface.co/blog/ResterChed).

### MiniMax IPO figures
- Up to **HK$4.19B** from **25.4M shares @ HK$151–165**; debut **2026-01-09**; ~**$6.5B** implied valuation; cornerstone **Alibaba + ADIA** [SECONDARY] (reuters.com).

