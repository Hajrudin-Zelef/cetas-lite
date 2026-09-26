---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/qwen3-7-max-may-2026-1m-context-35-hour-agent-tasks-secondar
title: "Qwen3.7-Max — May 2026, 1M context, 35-hour agent tasks [SECONDARY]"
domain: qwen-and-alibaba
role: deep-dive
task: agents
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-05", "2026-05-18", "2026-05-20"]
keywords: ["agent", "qwen", "benchmarks", "claude", "deepseek", "gemini", "glm", "leaderboard", "mcp", "open weights", "opus 4", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1159, 1172]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: b7a697e9162d32ee71c2c14438aa794a90b96e8b9771e9ceae4a3c1ad5d39d1a
---

# Qwen3.7-Max — May 2026, 1M context, 35-hour agent tasks [SECONDARY]

### Qwen3.7-Max — May 2026, 1M context, 35-hour agent tasks [SECONDARY]
- **2026-05-18** — Qwen3.7-Max-Preview and Qwen3.7-Plus-Preview land on Arena's leaderboard (Qwen's own X announcement: Alibaba now **#6 lab in text, #5 in vision**). **2026-05-20** — **Qwen3.7-Max** officially released at the **2026 Apsara Conference**. [SECONDARY]
- Agent positioning: designed for the agent era; claimed ability to **fully autonomously complete ultra-long-duration complex agent tasks lasting up to 35 hours**. [SECONDARY]
- Vendor-claimed benchmarks (marktechpost/wedoany compilations, label as vendor claims): **SWE-Verified 80.4** (vs Claude Opus-4.6 Max 80.8, DeepSeek-v4-Pro Max 80.6), **Terminal-Bench 2.0-Terminus 69.7** (vs DeepSeek-v4-pro-Max 67.9), **GPQA Diamond 92.4** (vs Opus-4.6 91.3), **HLE 41.4** (vs Opus-4.6 40.0), **MCP-Mark 60.8** (vs GLM-5.1 57.5), **MCP-Atlas 76.4** (vs Opus-4.6 75.8), **SpreadSheetBench-v1 87** (top tier), WMT24++ **85.8**, MAXIFE **89.2**. [SECONDARY]
- Independent: **Artificial Analysis Intelligence Index v4.0: 56.6** — +4.8 over Qwen3.6-Max-Preview (51.8), ahead of Gemini 3.5 Flash (55.3), behind GPT-5.5 (60.2), Claude Opus 4.7 (57.3), Gemini 3.1 Pro Preview (57.2). [SECONDARY]
- Text Arena: **#13 overall, Elo 1475**; category ranks #7 Math, #9 Expert Prompts, #9 Software/IT, #10 Coding. One careful-reading flag: AA-Omniscience raw accuracy dropped 7.6 points (37.7→30.1%) while hallucination rate fell 21.3 points (44.2→22.9%) — the model refuses more (attempt rate 67.3→48.0%, lowest among frontier models compared) rather than recalling more. [SECONDARY]
- YC-Bench (simulated year of startup operations): Qwen3.7-Max generated **$2.08M simulated revenue** vs $1.05M (Qwen3.6-Plus) vs $352K (Qwen3.5-Plus) — 2× and 5.9× generational jumps. Treat as vendor-selected simulation, not enterprise ROI. [SECONDARY]
- Licensing posture: **Qwen3.7-Plus to be open-sourced; Qwen3.7-Max stays proprietary** — the same two-tier playbook as 3.8. [SECONDARY]
- **Alibaba Cloud Summit 2026 (May 20, Hangzhou)**: Qwen3.7-Max launched alongside the **Zhenwu M890 AI chip** (3× predecessor performance) and the **Panjiu AL128 Supernode Server** (128 accelerators, PB/s internal bandwidth) — a full-stack reveal from silicon to model to cloud. Liu Weiguang (SVP Alibaba Cloud): "What we're building is China's AI factory." [SECONDARY]
- **The 35-hour agent demo**: Qwen3.7-Max was given a task brief on a **Zhenwu M890 chip it had never encountered in training**; working without human intervention, it ran **35 consecutive hours**, executed **>1,000 tool calls**, and delivered a **production-grade AI computing kernel outperforming the chip manufacturer's official version by 10×**. Optimized for OpenClaw, Hermes Agent, Claude Code, Qwen Paw, Qoder. [SECONDARY]
- **Pricing** (May 2026): Model Studio **¥12/¥36 per 1M** (~$1.71/$5.14); OpenRouter **$2.50/$7.50**; cache input **90% off** (¥1.2 / $0.25). **1M context** (doubled from 256K), **65,536 max output**. OpenAI-compatible endpoint (`qwen3.7-max`) + **native Anthropic Messages Protocol** (Claude Code & OpenClaw via `ANTHROPIC_BASE_URL` swap). Two API key types: `sk-` and `sk-sp-` (Token Plan). [SECONDARY]
- **Qianwen App integration** (May 22): Qwen3.7-Max in Qianwen App v6.9.7+, PC, Web — **free for all users**. [SECONDARY]
- **Policy shift context**: Qwen3.6-Max-Preview (April) and Qwen3.7-Max (May) are **API-only, no open weights** — Alibaba's first sustained turn toward the "frontier-closed, smaller-open" strategy. Alibaba killed the **free tier of Qwen Code** the prior month. [SECONDARY]

