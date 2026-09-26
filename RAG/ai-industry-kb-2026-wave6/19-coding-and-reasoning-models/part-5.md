---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/part-5
title: "§19. Coding and Reasoning Models (part 5)"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2026-03", "2026-07-09", "2026-09"]
keywords: ["reasoning", "agentic", "apache", "benchmarks", "claude", "cost", "cybersecurity", "deepseek", "distillation", "fable 5", "fp8", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9354, 9372]
section: "§19. Coding and Reasoning Models"
delta_of: ai-industry-kb-2026
sha256: 73852cc1b2189b65cad685a1853f89aa9667b503056b8e2181105f88dbe33ce8
---

# §19. Coding and Reasoning Models (part 5)

- Fable 5 safeguard-routing caveat: requests touching cybersecurity, biology, chemistry, or distillation reportedly fall back to Opus 4.8 — Fable 5's published scores are only achievable in "unsafeguarded" domains [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Kimi K3's reasoning_effort supports only max at launch; thinking cannot be disabled — unlike DeepSeek V4.1 Flash [SECONDARY]. Sources: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3 and https://cirra.ai/articles/en/pdfs/kimi-k3-release-open-weight-models.pdf
- V4.1 Flash vs Kimi K3 (OrcaRouter/AA, Sept 2026): $0.27 vs $2.00 per Intelligence Index task (7.4x); Vals AI agentic coding $0.41 vs $17.59 (~40x); throughput 214.4 vs 34.7 tok/s; cache read $0.003 vs $0.30 per 1M [SECONDARY]. Source: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- DeepSeek V4-Flash (0731) official card: peak cache hit $0.014 / miss $0.44 / output $1.32 per 1M; off-peak half price; peak hours 01:00–04:00 and 06:00–10:00 UTC weekdays — third-party flat cards are not the official table [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- DeepSeek V4.1 Flash card (newer SKU): $0.15/$0.60 off-peak, $0.30/$1.20 peak — do not confuse with the 0731 card [SECONDARY]. Source: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- A September 2026 analysis flags a 36Kr September 8 "V4.1 Flash probe" write-up as unverified, noting the live API id deepseek-v4-flash still resolved to the 0731 checkpoint — SKU identity is date-sensitive [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- Moonshot reports above 90% cache hits in coding workloads for K3 — effective input cost $0.30/M despite the $3.00 list rate [VENDOR/SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- AA blended 7:2:1 pricing: K3 $2.31, GLM-5.2 $0.90, DeepSeek V4 Pro $0.18 per 1M tokens; cost per task $0.94/$0.32/$0.04 [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- Self-hosting VRAM reality: GLM-5.2 (744B) needs >1TB VRAM in BF16 (~8x H200 at FP8); Kimi K3 heaviest with 64+ accelerators recommended — local serving out of reach for most teams despite open weights [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- DeepSeek V4 Pro list per one comparison: $0.435 input / $0.87 output / ~$0.0036 cached input per 1M [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- AA throughput: GLM-5.2 ~168 tok/s vs DeepSeek V4 Pro and Kimi K3 ~62 each [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- GPT-5.6 launch (2026-07-09): Sol/Terra/Luna share 1.05M context and 128K max output; gpt-5.6 is an alias for gpt-5.6-sol; no separate gpt-5.6-pro model ID — Pro is a reasoning mode on Sol [SECONDARY]. Source: https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- GPT-5.4 mini and nano cap at 272,000 input tokens (400K total with 128K output) — they cannot reach the >272K long-context tier and always bill at base price [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- Claude Code (2026): proprietary CLI, per-action permissions, VS Code/JetBrains integration, first-class MCP, multi-file operations [SECONDARY]. Sources: https://intuitionlabs.ai/articles/claude-code-vs-codex-vs-gemini-cli-comparison and https://www.deployhq.com/blog/comparing-claude-code-openai-codex-and-google-gemini-cli-which-ai-coding-assistant-is-right-for-your-deployment-workflow
- Codex CLI: Apache-2.0 client, Suggest/Auto-Edit/Full-Auto modes, sandboxing, MCP, CLI plus GitHub/PR workflows [SECONDARY]. Sources: https://intuitionlabs.ai/articles/claude-code-vs-codex-vs-gemini-cli-comparison and https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli
- Gemini CLI: Apache-2.0, 1M context, trusted-folders/Yolo mode, MCP, Google cloud tooling, free-tier claims — exact limits conflict across sources; cite per dated source [SECONDARY]. Sources: https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli and https://theaicareerlab.com/blog/claude-code-cli-vs-codex-cli-vs-gemini-cli-vs-opencode-cli
- Gemini CLI Plan Mode shipped in v0.34.0 (March 2026) as a read-only planning phase before execution [SECONDARY]. Source: https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli
- Practitioner qualitative evidence: Codex "felt better than Claude" on DeepSWE-style agentic work despite score proximity — attributed to harness UX and tool-loop design, not weights [COMMUNITY]. Source: https://github.com/the-vibe-company/website/blob/HEAD/content/articles/why-codex-felt-better-than-claude-deepswe.md

