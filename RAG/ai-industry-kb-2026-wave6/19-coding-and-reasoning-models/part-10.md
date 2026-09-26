---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/part-10
title: "§19. Coding and Reasoning Models (part 10)"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Anthropic", "DeepSeek", "Google", "Meta", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-02", "2026-05", "2026-05-19", "2026-07-30", "2026-08-28", "2026-09", "2026-09-03", "2026-09-21"]
keywords: ["reasoning", "agent", "agentic", "agents", "astra", "benchmark", "benchmarks", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9582, 9599]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 46cdf39a0de298c516622bfe77fccec0f52789b652fefc39852cb6387f36693f
---

# §19. Coding and Reasoning Models (part 10)

- Terminal-Bench version history: 2.1 (May 2026) ran 89 tasks and topped near 83%; 3.0 (2026-07-30) replaced the set with 74 harder tasks across 7 domains under the working name Frontier-Bench and the top score fell to 34.4%; 4.0 (2026-08-28) removed 8 tasks and fixed 19, leaving 66 tasks [SECONDARY]. Source: https://capitalandcompute.net/ai-benchmarks/terminal-bench/
- Terminal-Bench 4.0 (September 2026, official tbench.ai runs): Claude Fable 5.1 57.9% ±3.8 (#1), Opus 5 51.8% ±3.4, Fable 5 44.5%, GLM-5.3 41.8%, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1% ±3.4 (mini-SWE-agent, Sep 2) — TB 4.0 scores are NOT comparable with TB 2.1 numbers [SECONDARY]. Source: https://codingfleet.com/blog/terminal-bench-leaderboard-2026/
- Anthropic's own TB 4.0 figures differ from the official board: Mythos 5.1 60.9%, Fable 5.1 55.8%, Opus 5 52.3%, Fable 5 42.0% — vendor scaffold versus official runs [VENDOR]. Source: https://codingfleet.com/blog/terminal-bench-leaderboard-2026/
- GPT-6 Astra entered TB 4.0 on its 2026-09-03 launch day and took the top score at 58.18% in Codex at max effort; Fable 5.1 and two lower Astra effort settings tied at 57.88% just behind; the board grew from 10 to 18 entries that week and to 27 by late September [SECONDARY]. Source: https://capitalandcompute.net/ai-benchmarks/terminal-bench/
- Grok 4.7 entered TB 4.0 on its 2026-09-21 launch day at rank 17 with 37.58%, run in xAI's own Grok Build harness and flagged partial at 324 of 330 trials [SECONDARY]. Source: https://capitalandcompute.net/ai-benchmarks/terminal-bench/
- Semantic-versioning policy (from 3.0 onward): major releases change the environment or task set and force full re-runs; minor releases change verifiers and re-grade saved artifacts; patches change nothing score-moving; 4.1 (tamper-resistant verifiers) and 5.0 (new tasks) are open and unshipped as of September 2026 [SECONDARY]. Source: https://capitalandcompute.net/ai-benchmarks/terminal-bench/
- Harness-bound spread on TB 2.1: Artificial Analysis reports Fable 5.1 at 91.4% (Terminus 2, max effort), vals.ai reports the same model under the same harness at 85.02%, and the official board has no Fable 5.1 entry with Fable 5 in Claude Code at 83.8% on top — 6+ points of spread from sandbox, effort, and run count alone [SECONDARY]. Source: https://hackernoon.com/your-coding-agents-leaderboard-score-isnt-a-production-guarantee
- Official TB 2.1 verified top-of-board: Claude Code/Fable 5 83.8% ±1.2% (Jun 7, 2026); Codex/GPT-5.5 83.1% ±1.1% (May 1); Terminus 2/Fable 5 80.4% ±1.2% (Jun 5); Cursor CLI/Grok 4.5 79.3% ±1.5% (Jul 9); Claude Code/Opus 4.8 78.9% ±1.3% (Jul 9); Codex/GPT-5.6 Terra 78.4% ±1.3% (Jul 11); Terminus 2/GPT-5.5 78.0% ±1.2%; mini-SWE-agent/Muse Spark 1.1 76.2% ±1.2% (Jul 9); Codex/GPT-5.6 Luna 75.7% ±1.3% (Jul 11); Claude Code/Sonnet 5 74.6% ±1.6% (Jul 9); Terminus 2/Gemini 3 Pro 73.9% ±1.3%; Claude Code/Opus 4.7 68.9% ±1.4%; Gemini CLI/Gemini 3.1 Pro 65.8% ±1.7%; Claude Code/GLM-5.1 58.7% ±1.2% — 17 submissions [SECONDARY]. Source: https://github.com/kamil1721/coding-agent/blob/HEAD/docs/research/01-model-billing-and-architecture.sources.md
- Same-model harness delta: Fable 5 scores 83.8% in Claude Code versus 80.4% in Terminus 2 (3.4pt); Opus 4.7 68.9% versus 66.1% (2.8pt); Gemini 3.1 Pro 65.8% versus 65.6% (0.2pt) — published scores are earned inside vendor-tuned harnesses [SECONDARY]. Source: https://github.com/kamil1721/coding-agent/blob/HEAD/docs/research/01-model-billing-and-architecture.sources.md
- Contradiction ledger: codingfleet (Sep 11, 2026) claims DeepSeek V4.1 Flash takes TB 2.1 #1 at 90.6% (mini-SWE 90.3%, Claude Code 88.0%) at $0.15/$0.60 with MIT weights, but the official tbench.ai 2.1 board's verified top is Fable 5 at 83.8% — the 90.6% figure's verification status is unresolved [SECONDARY — contradiction]. Sources: https://codingfleet.com/blog/terminal-bench-leaderboard-2026/ versus https://github.com/kamil1721/coding-agent/blob/HEAD/docs/research/01-model-billing-and-architecture.sources.md
- Early-preview contradiction: Crypto Briefing's June 26 report had GPT-5.6 Sol at 88.8% versus Opus 4.8 at 78.9% with a "Sol Ultra" run at 91.9%, but the official verified board never listed Sol at those figures — treat preview reports as superseded by verified runs [SECONDARY]. Source: https://windowsforum.com/news/gpt-5-6-sol-leads-terminalbench-2-1-agentic-coding-beats-claude-for-enterprises.434773/
- Gemini 3.5 Flash (announced Google I/O, May 19, 2026): TB 2.1 76.2%, MCP Atlas 83.6%, GDPval-AA 1656 Elo, CharXiv Reasoning 84.2%, 4x faster output tokens/sec per Google's measurement — beats Gemini 3.1 Pro on coding/agentic at Flash latency [VENDOR]. Source: https://blog.imseankim.com/gemini-3-5-flash-vs-3-1-pro-terminal-bench-mcp-atlas-agentic-coding-may-2026/
- Benchmark-validity data point: OpenAI's February 2026 audit found 59.4% of the failed tasks it examined on SWE-bench Verified have flawed tests — the benchmark penalizes correct work as well as rewarding incorrect work [SECONDARY]. Source: https://hackernoon.com/your-coding-agents-leaderboard-score-isnt-a-production-guarantee
- τ²-bench (September 21, 2026): 151 models in the public snapshot; GLM-5.2 highest published at 99.1% but coverage does not establish a market leader; top-10 range is 1.4 points; the benchmark is displayed for reference but excluded from BenchLM's scoring formula (Agentic category, 22% weight) [SECONDARY]. Source: https://benchlm.ai/benchmarks/tau2-bench
- Artificial Analysis methodology note: TB 2.1 runs use Terminus 2 (E2B sandbox), 3 repeats, pass@1 averaged; TB 4.0 uses mini-SWE-agent v2.4.6 per-task verifier containers; quantisation arms (FP16-KV, NVFP4) inherit base-model scores — quantisation loss on this benchmark is not measured [SECONDARY]. Source: https://github.com/t0msilver/working-set/blob/HEAD/research/terminal_bench.md

**Additional facts, seventh tranche (DeepSWE methodology):**

