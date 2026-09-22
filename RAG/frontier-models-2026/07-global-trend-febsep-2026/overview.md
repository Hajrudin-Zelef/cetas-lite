---
id: frontier-models-2026/07-global-trend-febsep-2026/overview
title: "5. Global trend (Feb–Sep 2026)"
domain: global-trend-febsep-2026
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agentic", "benchmark", "benchmarks", "claude", "deepseek", "fable 5", "fp4", "gemini", "glm", "grok", "grok 4"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [384, 390]
section: "5. Global trend (Feb–Sep 2026)"
sha256: 7d78dfc6f8242607a3dcffd7869cea833fe1f30f2593ee54bf31a5524c69c040
---

# 5. Global trend (Feb–Sep 2026)
- **Chinese open-weight models (DeepSeek, Qwen, Kimi, GLM/Z.ai, MiniMax, MiMo) dominate the open-source top tier** in coding, agentic work, and price/performance. The pattern repeats: frontier-class scores at 5–10% of Western API prices, MIT-licensed weights, 1M-token contexts as standard.
- **Closed models (Claude Opus/Fable 5.x, GPT-5.x/6, Gemini 3.x, Grok 4.x) retain leads** on the hardest closed benchmarks and enterprise agentic tooling — but the gap on *long-horizon coding* has narrowed to single-digit points (GLM-5.2 vs Opus 4.8), and DeepSeek V4.1 Flash outright retired its own premium tier.
- **Architectural convergence on inference efficiency:** sparse MoE everywhere, KV-cache compression (FP4, CSA2, IndexShare), conditional memory (Engram), reasoning-effort dials, agent-task synthetic post-training.
- **Hardware sovereignty subplot:** GLM-5.3-Flash serving production traffic on ~100K domestic Chinese accelerators — first credible demonstration that SOTA open models can be trained/served without Nvidia.
- **Benchmark fragmentation:** vendors increasingly report on *different* benchmarks (SWE-bench Verified vs Pro vs DeepSWE vs Terminal-Bench), making like-for-like comparison unreliable — a structural caveat for the RAG.

