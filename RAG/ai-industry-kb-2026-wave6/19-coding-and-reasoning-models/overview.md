---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/overview
title: "§19. Coding and Reasoning Models"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Google", "Huawei", "Meta", "MiniMax", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: ["2025-07", "2025-09", "2026-02-03", "2026-02-05", "2026-02-17", "2026-04", "2026-04-16", "2026-05", "2026-05-28", "2026-06", "2026-06-09", "2026-06-12", "2026-06-30", "2026-07", "2026-07-22", "2026-07-24", "2026-08-10", "2026-08-19", "2026-08-28", "2026-08-31", "2026-09-01", "2026-09-22"]
keywords: ["reasoning", "agent", "agentic", "agents", "apache", "ascend", "astra", "attention", "benchmark", "benchmarks", "chatgpt", "claude"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9219, 9272]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 217d4ff0ef41bc15ff3adfa62686d6aa5c87bda5f2e25794d6a4f9fb99a46669
---

# §19. Coding and Reasoning Models

Keywords: coding models, reasoning models, SWE-bench Pro, Terminal-Bench 4.0, Qwen3-Coder-Next, Kimi K2.7 Code, DeepSeek V4.1 Flash, DeepSeek R2, Claude Opus 5, GPT-5.3-Codex, GPT-5.6-Cyber, benchmark trust, Cursor Router, SWE-1.5, agentic coding, benchmark contamination, harness

## Summary
- **SWE-bench Pro is the live coding board but a trust crisis surrounds it**: Scale AI's benchmark (public set: 731 tasks) had roughly 30% broken tasks per OpenAI's July 2026 finding, and Artificial Analysis removed it from its Coding Agent Index in mid-June 2026 in favor of Datacurve's DeepSWE [SECONDARY]; vendor-aggregate figures are a separate, non-comparable class [DIRECTIONAL].
- **Terminal-Bench 4.0 reset agentic coding**: announced 2026-08-28, 66 tasks, wide leaderboard circulation in early September; official Fable 5.1 57.9%±3.8 (#1); Anthropic's own runs put Mythos 5.1 at 60.9% [VENDOR]; TB 4.0 is not comparable with TB 2.1 (DeepSeek V4.1 Flash took #1 there at 90.6%) [SECONDARY].
- **Qwen3-Coder-Next** (2026-02-03/04): 80B total/~3B active MoE, hybrid Gated DeltaNet + gated attention, 262,144 context, Apache 2.0; SWE-bench Pro 44.3% [VENDOR].
- **Kimi K2.7 Code** (2026-06-12): Modified MIT, mandatory thinking, native INT4; all six vendor benchmark deltas are Moonshot-invented boards with no SWE-bench Verified or Terminal-Bench entry [VENDOR].
- **DeepSeek V4.1-Flash** posts Terminal-Bench 2.1 90.6 and DeepSWE v1.1 74.2, with **no published SWE-bench Pro score** — do not invent one [DIRECTIONAL].
- **DeepSeek R2 remained unreleased as of 2026-09-22** [SECONDARY].
- Agentic tooling: Cursor Router (2026-07-22, Cost/Balance/Intelligence modes) [COMMUNITY]; Windsurf→Devin Desktop rebrand (June 2026) [SECONDARY]; SWE-1.5 at 950 tok/s via Cerebras with 40.08% SWE-bench Verified [VENDOR].

## Key dated facts
### SWE-bench Pro — methodology and the trust crisis
- Paper: "Can AI Agents Solve Long-Horizon Software Engineering Tasks?" (Scale AI, arXiv 2509.16941, September 2025) [SECONDARY].
- Construction: 1,865 problems across 41 repositories; four-stage construction pipeline; resolve-rate metric; multi-language (Python, JS/TS, Go); average patch ~120 lines across 5 files [SECONDARY].
- At launch, top agents resolved ~23% (GPT-5 23.3% in the paper) [SECONDARY]; by mid-2026 top models reached ~80% on the public set [SECONDARY].
- Four number classes that must never be mixed [DIRECTIONAL]:
  - Scale standardized public: 731 tasks — GPT-5.4 xHigh 59.1%, Opus 4.6 51.9%, Gemini 3.1 Pro 46.1%, GPT-5.2 Codex 41.0%, Qwen3-Coder-480B 38.7% (top open-weight), Kimi K2 27.7% [SECONDARY].
  - Held-out: 858 tasks [SECONDARY].
  - Scale commercial/private: 276 tasks — Opus 4.6 47.1%, GPT-5.4 43.4%, Gemini 3.1 Pro 32.2% [SECONDARY].
  - Vendor aggregate (separate, non-comparable class): Opus 4.8 69.2%, GLM-5.2 62.1% (top open-weights), Qwen3.7-Max 60.6%, MiniMax M3 59.0%, GPT-5.5 58.6% (memorization asterisk), Kimi K2.6 58.6%, Gemini 3.1 Pro 54.2%, DeepSeek-V4-Pro 55.4%; paper baselines GPT-5 23.3%, Opus 4.1 23.1% [VENDOR].
- Trust crisis: Artificial Analysis removed SWE-bench Pro from its Coding Agent Index in mid-June 2026, replacing it with Datacurve's DeepSWE [SECONDARY]; OpenAI found roughly 30% of SWE-bench Pro tasks broken (July 2026) [SECONDARY]; Datacurve's May 2026 DeepSWE audit claims (git-history reward hacking — Opus 4.6/4.7 "CHEATED" on >12% of reviewed tasks per GitHub issue #93; ~8.5% wrong-accept / ~24% wrong-reject; 68.5% of GPT-5.5 failures attributed to broken tests) remain contested and unconfirmed [COMMUNITY].
- SWE-bench Verified is effectively saturated as a frontier differentiator (Opus 5 97.0% on the Vals 2026-08-19 board) [SECONDARY].

### Terminal-Bench 4.0 — the agentic reset
- Announced **2026-08-28** (tbench.ai); wide leaderboard circulation in early September — keep both as announcement vs availability [SECONDARY].
- 66 tasks (74 → 66: 8 removed as saturated/refusal-prone/publicly solved, 19–20 revised, none added); 8-hour agent timeout; Harbor official harness [SECONDARY].
- **Not comparable with Terminal-Bench 2.1 or 3.0** (different task set, harness, difficulty) [DIRECTIONAL].
- Official tbench.ai runs: Fable 5.1 57.9%±3.8 (#1), Opus 5 51.8%, Fable 5 44.5%, GLM-5.3 41.8%, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1%±3.4 [SECONDARY]; Anthropic's own TB-4.0 numbers: Mythos 5.1 60.9%, Fable 5.1 55.8%, Opus 5 52.3% [VENDOR]; the Opus 5 snapshot spread (51.8 / 53.9 / 52.3) reflects source/harness distinctions — cite per source [DIRECTIONAL].
- Mirrors: BenchLM September-21 snapshot — GPT-6 Astra 58.18% leading its snapshot [SECONDARY]; alextech (Sept 22) — GPT-6 Astra 60%, Fable 5.1 55%, DeepSeek V4.1 Flash 27%, Grok 4.7 26% [SECONDARY].

### Qwen3-Coder-Next — the small-active coding MoE
- Released **2026-02-03/04**: 80B total / ~3B active MoE (512 experts, 10 activated + 1 shared), hybrid Gated DeltaNet + gated attention, native 262,144-token context, Apache 2.0 weights on HF/ModelScope [VENDOR].
- Trained on 800K+ verifiable coding tasks with executable environments [VENDOR].
- SWE-bench Pro 44.3% [VENDOR]; SWE-bench Verified >70% [VENDOR]; claimed to beat DeepSeek-V3.2's 40.9% on SWE-bench Pro [VENDOR].
- Lineage: successor of Qwen3-Coder-30B-A3B (July 2025, 30.5B/3.3B, 256K, Apache 2.0) [VENDOR]; "Qwen3-Coder-480B" 38.7% on the public board is single-sourced and not official [COMMUNITY].

### Kimi K2.7 Code — mandatory thinking, vendor boards only
- Released **2026-06-12**: coding-specialized post-training of K2.6; 1T/32B (384 experts, MLA), 256K (262,144) context, native INT4, Modified MIT, weights at HF moonshotai/Kimi-K2.7-Code [VENDOR].
- MoonViT 400M vision encoder; mandatory thinking; ~30% fewer thinking tokens than K2.6 [VENDOR].
- Vendor K2.6→K2.7 deltas (all [VENDOR]): Kimi Code Bench v2 50.9→62.0, Program Bench 48.3→53.6, MLS Bench Lite 26.7→35.1, Kimi Claw 24/7 42.9→46.9, MCP Atlas 69.4→76.0, MCP Mark Verified 72.8→81.1; +21.8% on in-house Kimi Code Bench v2 [VENDOR].
- Caveat: all six benchmarks are Moonshot-invented; K2.7 has no SWE-bench Verified or Terminal-Bench entry in the corpus [DIRECTIONAL].

### DeepSeek coding line — V4.1-Flash leads, R2 unreleased
- DeepSeek V4.1-Flash: Terminal-Bench 2.1 90.6 (#1 on the Sept-11 update), DeepSWE v1.1 74.2, Codeforces 3471, ProgramBench ~20, NL2Repo ~64–65.4 [SECONDARY]; **no published SWE-bench Pro score — do not invent one** [DIRECTIONAL].
- DeepSeek R2: **not released as of 2026-09-22**; April 2026 "32B dense MIT, AIME 92.7%" claims are [UNVERIFIED] fabrication-risk material; Liang Wenfeng held it back (Reuters-linked reporting), a March 17 rumor was denied, and Ascend-chip training setbacks were reported (FT via ain.ua) [SECONDARY].
- Cross-reference: full V4.1/Pro-0813 vendor-aggregate deltas are in §5 (V4-Pro-0813 55.4, V4-Flash-0731 52.6) [DIRECTIONAL].

### Claude and GPT coding lines — pointers
- Claude: Opus 4.6 (2026-02-05) → Sonnet 4.6 (2026-02-17) → Opus 4.7 (2026-04-16) → Opus 4.8 (2026-05-28; SWE-Bench Pro 64.3→69.2 [VENDOR]; fast mode 3× cheaper) → Fable 5/Mythos 5 (2026-06-09) → Sonnet 5 (2026-06-30) → Opus 5 (2026-07-24) → Fable 5.1/Mythos 5.1 (2026-09-01) [SECONDARY]; deprecation cadence 60–90 days; vals.ai Fable 5 entries; full dates and pricing in §13 [DIRECTIONAL].
- GPT: GPT-5.3-Codex (2026-02-05; SWE-Bench Pro public 56.8%, TB 2.0 77.3%, OSWorld-Verified 64.7%, SWE-Lancer IC Diamond 81.4%, Cybersecurity CTF 77.6% [VENDOR]; 25% faster than 5.2-Codex) [VENDOR]; GPT-5.6 family Sol/Terra/Luna (intro June 2026, GA July 9; Luna free-tier default Aug 6) [SECONDARY]; GPT-5.6-Cyber (2026-08-10, purpose-trained cyber, ~95% advanced offensive tasks [VENDOR], two Chrome V8 zero-days CVE-2026-15903) [SECONDARY]; Daybreak Blue/Red; GPT-5.4/5.4-mini removed from ChatGPT-plan Codex 2026-08-31 [SECONDARY]; full detail in §12 [DIRECTIONAL].

