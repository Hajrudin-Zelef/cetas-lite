---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/model-router-ferduque-orchestrator-kimi-k3-worker
title: "model-router (ferduque) — orchestrator + Kimi K3 worker"
domain: kimi-and-moonshot-ai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "DeepSeek", "Hugging Face", "Microsoft", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "Together AI", "Z.ai", "vLLM", "xAI"]
dates: ["2026-07", "2026-07-16", "2026-07-26", "2026-07-28", "2026-08-02", "2026-09", "2026-09-07", "2026-09-21"]
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "copilot", "deepseek", "fable 5", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2211, 2264]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: 618dbcdf87d63d237284f5a80ea4519fdd4f0eca5f765045954348d36eb4f255
---

# model-router (ferduque) — orchestrator + Kimi K3 worker

### model-router (ferduque) — orchestrator + Kimi K3 worker
- Agent skill: a frontier model orchestrates while **cheap open-weight workers — Kimi K3, GLM 5.2, DeepSeek V4 — write the code as headless Claude Code sessions via one OpenRouter key** [COMMUNITY] (github.com/ferduque/model-router).
- July 2026 OpenRouter worker pricing table: **Kimi K3 $3.00/$15.00** (hardest tasks, long refactors); **GLM 5.2 $0.97/$3.06** (balanced default, frontend); DeepSeek V4 Pro $0.44/$0.87; V4 Flash $0.10/$0.20 [COMMUNITY].

### personal-pi-setup (atombarel) — Kimi K3 as the advertised OpenRouter model
- Local-first Pi agent setup advertises **Kimi K3 via OpenRouter (`moonshotai/kimi-k3`)** as its default model lane [COMMUNITY] (github.com/atombarel/personal-pi-setup).

### copilot-custom-endpoints (tugudush) — proxy-based Kimi K3 in Copilot
- Kimi K3 listed for VS Code Copilot via Moonshot **with proxy required**; comparison table also lists **MiniMax M3** (no proxy, dedicated extension) and **GLM 5.3 Flash** (no proxy) [COMMUNITY] (github.com/tugudush/copilot-custom-endpoints).

### generate-commit-extension (everton-dgn) — multi-provider AI commits
- Supports **Kimi (Moonshot) at `https://api.moonshot.ai/anthropic`** (default `kimi-k2.6`); **GLM (z.ai)** at `https://api.z.ai/api/anthropic`; **MiniMax** at `https://api.minimax.io/anthropic` (default `MiniMax-M2.5-highspeed`) — evidence of all three labs' Anthropic-compatible endpoints in one tool [COMMUNITY] (github.com/everton-dgn/generate-commit-extension).
- Practical note: **avoid `kimi-k2.7-code` for simple requests** — it rejects requests without thinking (verified 2026-08-02) [COMMUNITY].
- China hosts: **`https://api.moonshot.cn/anthropic`** (Kimi) and **`https://api.minimaxi.com/anthropic`** (MiniMax) [COMMUNITY].
- Claude Code CLI aliases observed: `fable`, `opus`, `sonnet`, `haiku` (verified with `claude --help` 2.1.220) [COMMUNITY].

### Promptfoo provider docs (ibamir) — Moonshot provider parameters
- Community-maintained Moonshot provider docs mirror the OpenAI provider's configuration surface for promptfoo evaluations [COMMUNITY] (github.com/ibamir/promptfoo).


### New verified facts — expansion (continued — Kimi K2 base and K3 release/benchmarks)

### Kimi K2 base model (aibase.com release coverage)
- **1T total / 32B active** parameters; claimed **top open-source performance on SWE Bench Verified, Tau2, and AceBench** at release [VENDOR via secondary] (aibase.com/news/19635).
- **MuonClip** "effectively solved the problem of large attention logits during large-scale training," improving training stability and token efficiency; **15.5T tokens trained stably with no loss spike** [VENDOR via secondary] (aibase.com).
- Two versions open-sourced simultaneously: **Kimi-K2-Base** (no instruction fine-tuning; research/custom scenarios) and **Kimi-K2-Instruct** (general instruction-tuned; Q&A and agent tasks); **model + FP8 weights on Hugging Face, free** [SECONDARY] (aibase.com).
- Inference engines with synchronized day-one support: **vLLM, SGLang, ktransformers** [SECONDARY] (aibase.com).
- API at launch: **up to 128K context**; China billing **4 yuan/M input / 16 yuan/M output** tokens; **compatible with both OpenAI and Anthropic API formats** [SECONDARY] (aibase.com).
- Coding examples cited: front-end code with design sense (particle systems, visualization, 3D scenes); **autonomously building a complete futures-trading interface without specific instructions** [SECONDARY] (aibase.com).
- Agent tool calls: parses complex instructions and decomposes requirements into **standardized, directly executable ToolCall structures**, integrating with Agent/Coding frameworks [SECONDARY] (aibase.com).

### Kimi K3 — release mechanics
- **Released July 16, 2026**; **open weights shipped July 26, 2026** — at 2.8T parameters, the **largest open-weight model ever released** at the time [SECONDARY] (theairankings.com).
- theairankings describes the weights as released **"under the Kimi K3 License"** — supporting the custom-license side of the §5 license conflict [SECONDARY] (theairankings.com).
- Described as the **first Chinese model to land inside the frontier pack on independent testing** [SECONDARY] (theairankings.com).
- Third-party hosts/agents at launch included **Devin** [SECONDARY] (theairankings.com).
- Pricing commentary: at **$3/$15 per M tokens** hosted, K3 costs roughly **triple its predecessor** — framed as "the end of super-cheap Chinese AI" [SECONDARY] (theairankings.com).

### Kimi K3 — vendor benchmark table (Moonshot technical blog; all [VENDOR])
- **Terminal-Bench 2.1: 88.3%** (Kimi Code harness); **FrontierSWE: 81.2%**; **DeepSWE: 67.5%** (Kimi Code; 67.3% on mini-SWE-agent); **SWE Marathon: 42.0%** (Claude Code harness); **Program Bench: 77.8%** (internal, raw pass rate not fully-resolved); **Automation Bench: 30.8%** (internal); **BrowseComp: 91.2%** (context compaction at 300K; 90.4% without) [SECONDARY reporting vendor] (emergent.sh/learn/kimi-k3-benchmark).
- **Harness caveat**: models were tested on different agent harnesses and the harness changes the score; **since the July 27 model card, Moonshot publishes per-benchmark footnotes naming the harness and source for every row** — attributions are now documented rather than inferred [SECONDARY] (emergent.sh).
- **Kimi Code Bench v2** (Moonshot in-house coding-agent benchmark, quarterly refresh, public question set): **K3 leads the public snapshot at 72.9%** (September 21, 2026), followed by Kimi K2.7 Code at 62.0%; BenchLM stores these as **provider-reported, display-only** launch evidence [SECONDARY] (benchlm.ai).

### Kimi K3 — independent benchmark readings
- **AA Intelligence Index v4.1: 57.1 — 4th overall**, behind Claude Fable 5 (59.9) and GPT-5.6 Sol (58.9), **ahead of Claude Opus 4.8 (55.7)** [SECONDARY] (theairankings.com).
- **AA v4.3 rebase (September 7, 2026)**: K3 (max) reads **44 at $2.00/task** — level with Grok 4.6, **a point under GLM-5.3**, three under GPT-5.6 Sol; theairankings keeps older-scale figures as dated and notes placement unchanged [SECONDARY] (theairankings.com).
- **AA-Briefcase: 1543 Elo (2nd, behind Fable 5)**; **1547 on AA's private long-horizon agentic eval — a 732-point jump over K2.6** [SECONDARY] (morphllm.com; artokun/comfyui-mcp).
- **Together AI DeepSWE: 68.5 pass@1**; **Terminal-Bench 2.1 via Vals: 80.9**; **#1 debut on LMArena Frontend Code Arena at 1679 Elo** [SECONDARY] (morphllm.com).
- **GDPval-AA v2: 1687 — 3rd, ahead of Claude Opus 4.8** [SECONDARY] (artokun/comfyui-mcp).
- Presentation quality: **Elo 1471** vs Sol 1660 and Opus 1492 — analytically strong but visually less polished [SECONDARY] (emergent.sh).
- Independent cross-lab table (intelligentliving.co, September 2026; `*` = vendor-reported/cited): K3 — **SWE-bench Multilingual 80.8\*; SWE-bench Pro 63.3\*; DeepSWE 67.5/74.0\*; Terminal-Bench 2.1 88.3/85.7\*; SWE-Marathon 42/44.4\*; CyberGym 80.0; ProgramBench 24.5\*; NL2Repo-Bench 58.0/58.3\*; PostTrainBench V1.1 32.0; SWE Atlas Codebase Q&A 35.2\*; Test Writing 35.6\*; Refactoring 37.4\*** [SECONDARY].
- Same table's GLM-5.3 column (for §4 cross-reference): Multilingual 81.3\*; Pro 64.6\*; DeepSWE 66.9/68.1\*; Terminal-Bench 2.1 88.2/88.3\*; CyberGym 84.5/83.0\*; ProgramBench 18.0\*; SWE-Marathon 42.5/35.6\* [SECONDARY] (intelligentliving.co).
- Community coding benchmark **wave 2 (2026-07-28)**: **K3 v2 score 95** (64.9 min, 13.2M tokens, **$6.14**, Moderato sub) — tied #2 with Claude Opus 5; **Kimi K2.7-Coding 86** ($4.37); **Kimi K2.6 77** ($2.64 via OpenRouter); GLM 5.2 91 ($0 marginal on Z.ai coding plan; $12.05 API-equiv) [COMMUNITY] (akitaonrails/llm-coding-benchmark success_report.v2).

