---
id: labs-hyperscalers-2026/00-labs-hyperscalers/part-32
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 32)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "MiniMax", "Mistral", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-06-09", "2026-06-15", "2026-06-17", "2026-08-18"]
keywords: ["agent", "aws", "bedrock", "claude", "compute", "deepseek", "fable 5", "gemini", "glm", "grok", "grok 4", "guardrails"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1023, 1036]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 8b621e34f407c3ab645e257a0254c900d877764518dbb89e3b1f1ec31dcb2647
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 32)

**Other models onboarded to Bedrock in 2026:**
- **xAI Grok 4.3** — GA **June 15, 2026** (first xAI model on Bedrock); via **Bedrock Mantle** (OpenAI-compatible inference engine); 1M-token context; configurable reasoning effort; **$1.25/M input, $2.50/M output** [secondary; AWS launch blog].
- **xAI Grok 4.6** — launched on Bedrock **August 18, 2026**; 500K-token context; reasoning effort at four levels (low/medium/high/xhigh); available on both bedrock-mantle and bedrock-runtime endpoints; supports Converse API + Chat Completions + Responses [official] https://aws.amazon.com/blogs/machine-learning/xais-grok-4-6-is-now-available-in-amazon-bedrock/.
- **Anthropic Claude line:** Opus 4.8, Sonnet 5, Haiku 4.5, and "Claude Fable 5" **GA June 9, 2026** on Bedrock [secondary][unverified details].
- **DeepSeek:** fully managed DeepSeek-R1 serverless option on Bedrock added 2026; DeepSeek V3.2/V3.1 listed serverless (as of June 15, 2026) [secondary].
- **Chinese open-weight additions (as of June 15, 2026):** Z.AI GLM-5, GLM-4.7, GLM-4.7 Flash; Moonshot Kimi K2.5, K2 Thinking; MiniMax M2.5, M2.1, M2; Alibaba Qwen3 variants [secondary — AWS Summit 2026 slide deck].
- Bedrock scale: **18+ providers, 110+ model variants** as of mid-2026 [secondary]; **Bedrock used by 125,000+ customers**, with most inference running on Trainium [vendor-reported][secondary].

**AgentCore (2026):**
- **Bedrock Guardrails GA June 17, 2026**, integrated into AgentCore: content filters, denied topics, sensitive-info filters, contextual grounding checks [secondary].
- AgentCore runtime compute pricing: **$0.0895 per vCPU-hour + $0.00945 per GB-hour**; session/memory: short-term $0.25/1,000 events; long-term storage $0.75/1,000 (built-in) or $0.25/1,000 (custom); retrieval $0.50/1,000; Gateway routing $0.005/1,000 API invocations, $0.025/1,000 search invocations, $0.02 per 100 tools indexed/month [secondary].
- AgentCore is model-agnostic (OpenAI, Gemini, Claude, Nova, Llama, Mistral); 15 AWS regions for agent runtime [secondary].
- No Bedrock-wide pricing *cuts* found in 2026; competitive posture is per-model undercutting (e.g., Nova 2 Pro at $1.25/$10 vs. rivals) [secondary].

