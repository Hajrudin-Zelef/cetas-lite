---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/timeline-and-context
title: "Timeline and context"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-02-03", "2026-02-05", "2026-02-17", "2026-04-16", "2026-05", "2026-05-01", "2026-05-28", "2026-06-07", "2026-06-09", "2026-06-12", "2026-06-30", "2026-07", "2026-07-09", "2026-07-22", "2026-07-24", "2026-08-06", "2026-08-10", "2026-08-21", "2026-08-26", "2026-08-28", "2026-08-31", "2026-09-01", "2026-09-10", "2026-09-21", "2026-09-22"]
keywords: ["agentic", "apache", "astra", "benchmark", "benchmarks", "chatgpt", "claude", "cost", "cyber", "deepseek", "distillation", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9426, 9479]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 398da2c969783b44b7de930f5f2fb91f4456e29254f50588c381dda3400be336
---

# Timeline and context

## Timeline and context
- **2026-02-03/04** — Qwen3-Coder-Next released [VENDOR].
- **2026-02-05** — Claude Opus 4.6; GPT-5.3-Codex [SECONDARY].
- **2026-02-17** — Claude Sonnet 4.6 [SECONDARY].
- **2026-04-16** — Claude Opus 4.7 [SECONDARY].
- **2026-05-28** — Claude Opus 4.8 [SECONDARY].
- **2026-06-09** — Claude Fable 5 / Mythos 5 [SECONDARY].
- **2026-06-12** — Kimi K2.7 Code released [VENDOR].
- **2026-06** — Windsurf → Devin Desktop rebrand [SECONDARY].
- **2026-06-30** — Claude Sonnet 5 [SECONDARY].
- **2026-07-09** — GPT-5.6 family GA [SECONDARY].
- **2026-07-22** — Cursor Router shipped [COMMUNITY].
- **2026-07-24** — Claude Opus 5 [SECONDARY].
- **2026-08-06** — GPT-5.6 Luna free-tier default [SECONDARY].
- **2026-08-10** — GPT-5.6-Cyber [SECONDARY].
- **2026-08-28** — Terminal-Bench 4.0 announced [SECONDARY].
- **2026-08-31** — GPT-5.4/5.4-mini removed from ChatGPT-plan Codex [SECONDARY].
- **2026-09-01** — Claude Fable 5.1 / Mythos 5.1; TB 4.0 leaderboard circulation [SECONDARY].


### New verified timeline entries — expansion

- 2026-03: Gemini CLI Plan Mode ships in v0.34.0 (read-only planning phase) [SECONDARY]. Source: https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli
- 2026-05-01: GPT-5.5 + Codex (xhigh) TB 2.1 entry at 83.1%±1.1 — earliest dated row on the transcribed board [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- May 2026: DeepSWE launches; GPT-5.5 leads at 70%±4; Opus-loophole story is the launch coverage [SECONDARY]. Sources: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole and http://gigazine.net/gsc_news/en/20260528-deepswe-ai-coding-benchmark/
- 2026-06-07: Fable 5 + Claude Code (xhigh) TB 2.1 entry at 83.8%±1.2 — transcribed board leader [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- 2026-07-09: GPT-5.6 launches (Sol/Terra/Luna, 1.05M context) [SECONDARY]. Source: https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- 2026-08-21/22: "ox-alpha" (GLM-5.3-Flash) stealth-tested on OpenRouter/OpenCode — community forensics [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- 2026-08-26: GLM-5.3-Flash MIT launch; Zhipu HK shares +12% [SECONDARY]. Sources: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com and https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- 2026-09-01: Steel.dev Verified board updated; Opus 5 Vals 97.00%±0.76 top third-party figure [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- 2026-09-10: OSWorld-Verified snapshot and generic LCB ledger dated [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- 2026-09-21: LiveCodeBench v6 snapshot [SECONDARY]. Source: http://benchlm.ai/benchmarks/livecodebench-v6
- 2026-09-22: DeepSWE board observed — four-way 74/73% tie; GPT-6 Astra and Gemini 3.8 Flash appear [VENDOR]. Source: https://deepswe.datacurve.ai/

## Implications
1. SWE-bench Pro numbers must always carry their class label (public-731 / held-out-858 / commercial-276 / vendor aggregate); unqualified "SWE-bench Pro %" claims are unusable [DIRECTIONAL].
2. The June–July 2026 benchmark-trust episode (AA removal, OpenAI's ~30%-broken finding, the contested Datacurve audit) makes benchmark dates and harnesses first-class citation elements, not footnotes [DIRECTIONAL].
3. Terminal-Bench 4.0's 66-task reset and TB 2.1 are different populations; any V4.1-Flash 90.6 vs TB-4.0 comparison is invalid [DIRECTIONAL].
4. The open coding-model story in 2026 runs through Qwen3-Coder-Next (Apache 2.0, 44.3% vendor SWE-bench Pro) and Kimi K2.7 Code (Modified MIT) — but K2.7's boards are all vendor-invented, so public-board evidence favors Qwen [DIRECTIONAL].
5. Full Claude/GPT coding-line pricing and the DeepSeek R2 non-release detail live in §§5/12/13; per-model one-liners here only (delta discipline) [DIRECTIONAL].


### New verified implications — expansion

- TB 2.1's June board: $552–$2,059/run for sub-84% scores, while the Ante harness reaches 83.9% (V4.1 Flash) for ~$18 total — the same nominal benchmark number can cost 100x more or less by harness and model; published TB 2.1 figures are purchasing decisions, not model constants [DIRECTIONAL].
- DeepSWE September's $2.36–$13.41 per-run spread at equal accuracy makes benchmark leadership a cost question first — procurement tables omitting per-run cost misrank models [DIRECTIONAL].
- Opus 5's 10x cost-per-task range across effort ($1.78→$17.79) plus accuracy declining at max effort makes "max by default" a spend-more-for-worse configuration — effort routing belongs in cost models [SECONDARY].
- GPQA's ~94% cluster means the benchmark no longer discriminates at the frontier — new claims need unsaturated suites (DeepSWE, Frontier-Bench v0.1, OSWorld-Verified) [DIRECTIONAL].
- X-Coder's 7-point v5→v6 drop and issue #99 confirm: any KB row lacking a benchmark version pin is unverifiable — dated snapshots must stay in separate tables [SECONDARY].
- Fable 5's safeguard fallback to Opus 4.8 in cyber/bio/chem/distillation means published Fable 5 scores describe a subset of real traffic [SECONDARY].
- The absent GPT-5.6 Sol / Opus 5 from the 17-row TB 2.1 transcription versus the September community file's 89.5%/89.1% shows the "official board" is a moving, partially-reported artifact — the contradiction ledger stands [SECONDARY].
- GLM-5.3-Flash's ox-alpha stealth test shows open-weight launches are now A/B-tested under aliases before the license is announced — the "launch date" is the marketing date, not the first-public-exposure date [SECONDARY].
- A practitioner finding Codex "felt better than Claude" on agentic work despite score proximity is qualitative evidence that harness UX and tool-loop design dominate perceived quality — the model is necessary but not sufficient [COMMUNITY].

