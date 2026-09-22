---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/muse-spark-1-3-launch-contributor-tier-and-meta-muse-agent-2
title: "Muse Spark 1.3 launch, Contributor tier, and Meta Muse agent (2026-09)"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: agents
actors: ["AMD", "Anthropic", "Hugging Face", "Meta", "Microsoft", "Nvidia", "OpenAI", "SGLang", "United States", "vLLM"]
dates: ["2025-06", "2026-01-26", "2026-02-05", "2026-04", "2026-07-09", "2026-08-05", "2026-08-07", "2026-08-10", "2026-09-02"]
keywords: ["agent", "muse", "muse spark", "acquisition", "agentic", "amd", "apache", "benchmark", "benchmarks", "capex", "claude", "compute"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4330, 4441]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: 1123e1a0cf2a3184ec3c763e4a853b32d4691f61680c7024d6ab5cd037eac240
---

# Muse Spark 1.3 launch, Contributor tier, and Meta Muse agent (2026-09)

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

### Muse Glimmer additional corroborated details (2026-08-10)
- First open-weight model from Meta Superintelligence Labs since the $14.3B Scale AI acquihire of June 2025; distilled from the flagship Muse Spark teacher model (logit distillation + long-context agent data + RL). Zuckerberg confirmed plans to open-source Muse Spark 1.2 weights in the near future. [SECONDARY, S18][SECONDARY, S80][SECONDARY, S82]
- Meta's K-Quant-17GB build targets 24GB VRAM at 1.0% average degradation across 15 benchmarks; K-Quant-Dynamic targets 32GB at 0.2%; quantization compresses 55GB to ~4-bit under 20GB with minimal degradation on agentic tasks. [SECONDARY, S67][SECONDARY, S78][SECONDARY, S81]
- Model weights are on Hugging Face with adapted builds for llama.cpp, MLX, ExecuTorch, Ollama, LM Studio, vLLM, and SGLang; Ollama called it the first MSL release on its platform, with NVIDIA/AMD/LM Studio launch-day support. [SECONDARY, S65][SECONDARY, S80][SECONDARY, S82]
- Zuckerberg published a 14-page essay alongside the release arguing for distributed, open AI development over centralized systems; Meta frames open weights as vital for American competitiveness and against regulatory capture. [SECONDARY, S64][SECONDARY, S77][SECONDARY, S81]
- Deployment caveat: the 17GB figure describes one quantized weight file only; adding the perception encoder, DFlash drafter, KV cache, runtime, and app leaves much less headroom on 24GB cards than the headline suggests. [SECONDARY, S68 — single source]

## Figures and metrics
| Model | Release | Weights | License | Context | Price $/1M in/out |
|---|---|---|---|---|---|
| Llama 4 (family) | — | Open | Llama 4 Community License | — | — |
| Muse Spark 1.2 | 2026-08-05 | Closed | — | 1M | $1.25/$4.25; contributor $0.10/$0.20 [SECONDARY] |
| Muse Glimmer 30B | 2026-08-10 | Open | Apache 2.0 | — | — [SECONDARY] |
| Muse Spark 1.3 | 2026-09-02 | Closed | — | — | $1.25/$4.25 [VENDOR] |

- Llama 4 Community License MAU cap: 700M [SECONDARY].
- Spark 1.1 + mini-SWE-agent: 61.5% on SWE-bench Pro official board [SECONDARY].
- Spark 1.3: AA Index v4.3 = 48 [SECONDARY] (v4.3 only; no cross-version comparison).
- Version-pinning rule for Meta figures: Spark 1.3's 48 is an AA **v4.3** (Sept 7, 2026) number — the v4.1.1→v4.2→v4.3 churn week means any undated Meta benchmark citation is suspect; re-pin before reuse [SECONDARY/DIRECTIONAL].
- Spark 1.2 contributor tier ($0.10/$0.20) vs standard ($1.25/$4.25): a 12.5×/21× input/output spread between tiers on the same model — cite the tier or the price is meaningless [SECONDARY].


### New verified metrics — expansion

### Spark 1.1 / Glimmer additional figures
- Spark 1.1: released 2026-07-09; $20 free credits; $1.25/$4.25; no deprecation at 1.2 launch; contributor tier 1.2-only, 60 req/min cap. [SECONDARY, S61][SECONDARY, S62]
- Glimmer: AIME 2026 94.7%; SWE-Bench Pro 51.2%; K-Quant-17GB 1.0% degradation / Dynamic 0.2%. [SECONDARY, S66][SECONDARY, S67]

### Spark 1.2 agentic figures
- TerminalBench 2.1 (Meta chart, mixed harnesses): 1.2 82.9; Codex 81.8; Claude Code 86.7; 1.1 76.2. [SECONDARY, S54][SECONDARY, S59]
- DeepSWE 1.1: 1.2 59.3; Claude Code 65.0; Codex 64.8. [SECONDARY, S54]
- AA 2026-08-05: Intelligence Index 51→54; GDPval-AA v2 1371→1631; cost/task $0.29→$0.40. [SECONDARY, S58]

### Spark 1.3 / Muse / subscriptions additional figures
- Spark 1.3 standard: $1.25/$4.25 per 1M; cached $0.15/M; released 2026-09-02. [SECONDARY, S44][SECONDARY, S47]
- Contributor tier: $0.10/$0.20 per 1M; cached $0.002/M. [SECONDARY, S45]
- 1.3 xhigh: 100M output tokens per AA suite (vs 58M for original Spark). [SECONDARY, S46]
- p95 TTFT: 8.60s via Meta Model API. [SECONDARY, S44]
- Muse agent: free 100M tokens/week; $20/$100 monthly tiers. [SECONDARY, S49]
- Manus acquisition: reported $2B. [SECONDARY, S50]

### Llama 3.3 / Llama 4 / Behemoth additional figures
- Llama 3.3: MMLU Pro 68.9; IFEval 92.1; GPQA Diamond 50.5; MBPP EvalPlus 87.6; MATH 77.0; BFCL v2 77.3; MGSM 91.1; cutoff Dec 2023; pricing $0.1/$0.4 per 1M. [VENDOR, S34][VENDOR, S35][SECONDARY, S43]
- Maverick: 17B/400B; 19–49¢ per M input/output; LMArena Elo 1417 (experimental). [SECONDARY, S40][SECONDARY, S36]
- Scout: 17B/109B; retrieval to 10M tokens. [SECONDARY, S40][SECONDARY, S36]
- Llama 4 training CO2: 2,000 tons (Maverick + Scout). [SECONDARY, S40]
- Behemoth: ~2T total / 288B active / 16 experts; 30T+ multimodal tokens. [SECONDARY, S39]
- Meta: $72B annual capex (planned); 11/14 original Llama researchers departed. [SECONDARY, S41]


### Llama 3.x family figures
- Llama 3: 8B/70B params; 8K context; ~15T training tokens. [SECONDARY, S1]
- Llama 3.1: 8B/70B/405B; 128K context; 32/80/126 layers. [SECONDARY, S1]
- Llama 3.2: 1B/3B text + 11B/90B vision; 128K context; ~9T tokens. [SECONDARY, S1]
- Llama 3.3 70B (vendor): MMLU 86.0; HumanEval 88.4; vs 405B at 88.6/89.0. [VENDOR, S2]
- Llama 3.3: 8 supported languages; 128K context. [SECONDARY, S2][SECONDARY, S4]

### Llama 4 figures
- Scout: 109B total / 17B active; 16 experts; 10M context. [SECONDARY, S5]
- Maverick: ~400B total / 17B active; 128 experts; 1M context. [SECONDARY, S5]
- Behemoth (preview): ~2T total / 288B active; 16 experts; 30T+ tokens; unreleased. [SECONDARY, S9][SECONDARY, S11]
- Scout INT4: fits one H100. [SECONDARY, S5]
- Launch distribution: 40 countries across WhatsApp/Messenger/Instagram. [SECONDARY, S6]

### Muse Spark 1.2 figures
- Context: 1M; vendor Terminal-Bench 2.1 82.9%; DeepSWE v1.1 59.3%. [VENDOR, S13]
- Public-board check 2026-08-07: entries not listed. [SECONDARY, S12]

### Muse Glimmer 30B figures
- 29.6B dense LM + 1.8B ViT-G/14; 120K+ context; 100+ languages. [SECONDARY, S17][SECONDARY, S19]
- Quantized build <20GB; targets 24–32GB GPUs. [SECONDARY, S17][SECONDARY, S18]
- RTX 5090 throughput: 74.9 → 233.4 tok/s (3.1×); M5 Max 1.9×; M4 Max 1.6×. [VENDOR, S17]
- Vendor benchmarks: MCP Atlas 75.5; TB 2.1 51.7; OSWorld-Verified 65.9. [VENDOR, S17]

### Muse Spark 1.3 figures
- Context 1,048,576; max output 943,718. [SECONDARY, S22]
- Efficiency vs 1.2: −20% tool calls, −25% tokens (vendor). [VENDOR, S23]
- Cached input: $0.15/M (standard tier). [SECONDARY, S22]
- Vendor card: DeepSWE v1.1 75.4; TB 2.1 88.8; SWE-Atlas QnA 59.4; MRCR v2 98.5 (256K–512K) / 98.1 (512K–1M); OSWorld 2.0 66.9 (max) / 57.2 (xhigh); GDPval-AA v2 1754 (max) / 1709 (xhigh); JobBench 64.9 (max) / 61.2 (xhigh). [VENDOR, S22]

### Meta platform figures
- 1.2B cumulative Llama downloads at LlamaCon 2025 (vendor). [VENDOR, S26]
- Vibes testing start: 2026-02-05; initial markets Brazil and Mexico. [SECONDARY, S30]
- Premium subscription testing announced 2026-01-26/27 (IG/FB/WhatsApp). [SECONDARY, S32]

