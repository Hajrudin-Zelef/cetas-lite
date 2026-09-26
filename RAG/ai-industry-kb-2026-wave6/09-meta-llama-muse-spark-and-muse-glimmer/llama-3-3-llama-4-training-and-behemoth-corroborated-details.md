---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/llama-3-3-llama-4-training-and-behemoth-corroborated-details
title: "Llama 3.3, Llama 4 training, and Behemoth — corroborated details (2024-12 → 2026-04)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: training
actors: ["Anthropic", "EU", "Meta", "Microsoft", "OpenAI", "United States"]
dates: ["2023-12", "2026-04", "2026-07-09", "2026-08-05", "2026-09-02"]
keywords: ["llama", "training", "agent", "agentic", "benchmark", "benchmarks", "claude", "compute", "cost", "latency", "leaderboard", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4315, 4356]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: eb01a02972798ef035aa12e5b5b62b83cea37b1d7406d4968680681353464d73
---

# Llama 3.3, Llama 4 training, and Behemoth — corroborated details (2024-12 → 2026-04)

### Llama 3.3, Llama 4 training, and Behemoth — corroborated details (2024-12 → 2026-04)
- Llama 3.3 full vendor benchmark table: MMLU Pro 68.9; IFEval 92.1; GPQA Diamond 50.5; MBPP EvalPlus (base) 87.6; MATH 77.0; BFCL v2 77.3; MGSM 91.1. [VENDOR, S34][VENDOR, S35]
- Llama 3.3 pretraining data cutoff: December 2023. [VENDOR, S35 — single source]
- Llama 3.3 API pricing at the time: $0.1 per million input tokens, $0.4 per million output tokens. [SECONDARY, S43 — single source]
- Llama 4 (Scout/Maverick) parameter counts: Maverick 17B active / 400B total across 128 experts, 1M context (needs H100 DGX-class); Scout 17B active / 109B total across 16 experts, 10M context (runs on a single H100). [SECONDARY, S40][SECONDARY, S63][SECONDARY, S64][SECONDARY, S65][COMMUNITY, S66]
- Llama 4 training carbon: Maverick and Scout training emitted 2,000 tons of CO2 (Behemoth's figure undisclosed). [SECONDARY, S40 — single source]
- Maverick operating cost estimate: 19–49 cents per million input/output tokens. [SECONDARY, S40 — single source]
- Llama 4 long-context retrieval: needle-in-haystack tests show retrieval up to 1M tokens for Maverick and 10M for Scout (retrieval only, not full-context problem solving). [SECONDARY, S36 — single source]
- LMArena aftermath: 2,000+ head-to-head battle results released publicly; the experimental Maverick produced longer, more formatted, emoji-heavy responses than the public build; LMArena updated its leaderboard policies and added the HF version of Maverick. [SECONDARY, S37][SECONDARY, S38]
- Behemoth scale: ~2T total parameters with ~288B active per token across 16 experts. [SECONDARY, S39][SECONDARY, S63][SECONDARY, S65]
- Meta's stated intent was an open-weight Behemoth under the same Llama 4 community license (restricting companies with 700M+ MAU, and EU-domiciled companies barred without explicit approval) — no release has confirmed it. [SECONDARY, S39][SECONDARY, S63][SECONDARY, S65]
- Research-team attrition: 11 of the 14 researchers behind the original Llama model had left Meta by late 2025. [SECONDARY, S41 — single source]
- Meta's planned annual capital expenditure at the time: up to $72B, much of it for AI development. [SECONDARY, S41 — single source]


### Muse Spark 1.3 launch, Contributor tier, and Meta Muse agent (2026-09)
- Muse Spark 1.3 was released 2026-09-02 via Muse Code and the Meta Model API; Meta AI, Instagram, and Facebook access expected later. [SECONDARY, S47][SECONDARY, S44]
- Standard pricing: $1.25 per million input tokens, $4.25 per million output tokens, $0.15 cached input; multimodal text+image input; proprietary license. [SECONDARY, S44][SECONDARY, S47]
- Contributor tier: $0.10/M input, $0.20/M output, $0.002/M cached input — 12.5×, 21×, and 75× discounts vs standard — in exchange for consent described as "permission to use your prompts and completions to train future Meta models"; retention period, human review, and deletion/revocation are not specified in the documentation. [SECONDARY, S45][SECONDARY, S48]
- Vercel AI Gateway serves both `meta/muse-spark-1.3` (standard rates) and `meta/muse-spark-1.3-contributor` (90% cheaper with training-data contribution). [SECONDARY, S52 — single source]
- Meta's technical approach: "thought compression" — an RL reward signal plus a penalty on thinking-token count; the model first reasons more carefully/concisely, then undergoes a phase transition to compressed reasoning chains extended from a new efficient baseline. [SECONDARY, S46][SECONDARY, S67]
- Token intensity: the 1.3 xhigh variant generated 100M output tokens across Artificial Analysis's nine-evaluation Intelligence Index suite, vs 58M for the original Muse Spark; average cost per evaluation task rose $0.40 (1.2) to $0.55 (1.3) as input tokens per task rose ~57%. [SECONDARY, S46][SECONDARY, S67][SECONDARY, S68]
- Launch messaging: Alexandr Wang called it Meta's "biggest jump so far on model performance"; Zuckerberg declared "frontier performance almost too cheap to meter"; internal coding workflows claimed −25% tokens and −20% tool calls — with the caveat that 1.3 was evaluated at max reasoning vs 1.2 at xhigh, so part of the gain reflects settings, not just the model. [SECONDARY, S47][SECONDARY, S68][SECONDARY, S69][SECONDARY, S71]
- Technical report published at research.meta.ai/static/muse-spark-1-3-multimodal-evaluation-methodology. [SECONDARY, S44 — single source]
- Observed latency: p95 time-to-first-token 8.60 seconds via Meta Model API over the trailing 7 days. [SECONDARY, S44 — single source]
- Meta launched "Muse," a personal AI agent powered by Muse Spark 1.3 running on a dedicated VM in Meta's cloud — free for up to 100M tokens per week, with $20 and $100 monthly tiers for more compute; initially US-only, with AI glasses support planned. [SECONDARY, S49][SECONDARY, S71]
- Subscription-tier details: Instagram premium perks include unlimited audience lists, a non-follower list, and anonymous Story viewing; Meta plans to fold in Manus (the AI agent acquired for a reported $2B) and freemium Vibes video creation. [SECONDARY, S50][SECONDARY, S51]


### Muse Spark 1.2 agentic benchmarks and methodology debate (2026-08)
- Reuters/Zuckerberg internal comparison: TerminalBench 2.1 — Spark 1.2 82.9 (in Muse Code), Codex 81.8 (OpenAI), Claude Code 86.7 (Anthropic); DeepSWE 1.1 — Spark 1.2 59.3 vs Claude Code 65.0 vs Codex 64.8. [SECONDARY, S54 — single source]
- Artificial Analysis (2026-08-05) 1.1→1.2: Intelligence Index 51→54 (+3); GDPval-AA v2 Elo 1371→1631 (+260); τ³-Banking 25%→27%; SciCode 58%→56%; Humanity's Last Exam 45%→44%; CritPt 15%→18%; cost per Intelligence Index task $0.29→$0.40 as input/output tokens rose ~53%/36%. [SECONDARY, S58 — single source]
- Methodology critique: Meta's Terminal-Bench 2.1 chart ran each model in a different agent harness (1.2 in Muse Code, 1.1 in mini-swe-agent at 76.2, Opus 5 in Claude Code, Codex in Codex, etc.); a constant-harness analysis found ~2 points of real gain vs the 6.7 points Meta's chart implied. [SECONDARY, S59 — single source]
- Muse Code installs via `curl -fsSL https://dev.meta.ai/install.sh | bash`; installer defaults to the contributor tier. A developer-preview SDK extends the agent beyond the CLI; subagents run in isolated git worktrees with inter-session messaging. [SECONDARY, S55][COMMUNITY, S73][SECONDARY, S76]


### Muse Spark 1.1 launch (2026-07-09; beyond base section)
- Muse Spark 1.1 was released 2026-07-09 with developer access in US public preview; sign-ups got $20 in free credits, then pay-as-you-go at $1.25/M input and $4.25/M output. [SECONDARY, S61][SECONDARY, S63]
- Meta called it its most capable model for real-world coding and agentic tasks, part of the "personal superintelligence" mission; it shipped in Thinking mode in the Meta AI app and was expected to replace Llama-powered chatbots on WhatsApp, Instagram, Facebook, and smart glasses. [SECONDARY, S61][SECONDARY, S63]
- The original Muse Spark debuted in April 2026 as Meta Superintelligence Labs' first text-and-reasoning model, tested with partners in private preview. [SECONDARY, S61 — single source]
- At the August 5 1.2 launch: no deprecation and no repricing of 1.1; 1.1 remained on the Meta Model API at $1.25/$4.25 (model ID meta/muse-spark-1.1); the contributor tier is 1.2-only (muse-spark-1.2-contributor) with a 60-requests-per-minute cap. [SECONDARY, S62 — single source]

