---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/kimi-k3-efficiency-mechanisms-and-serving-economics
title: "Kimi K3 — efficiency mechanisms and serving economics"
domain: kimi-and-moonshot-ai
role: deep-dive
task: finance
actors: ["AWS", "Anthropic", "Fireworks AI", "Hugging Face", "Moonshot", "OpenRouter", "Together AI"]
dates: ["2026-01-27", "2026-08", "2026-09"]
keywords: ["kimi", "agent", "attention", "aws", "bedrock", "benchmark", "claude", "gpu", "inference", "int4", "mcp", "memory"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2265, 2291]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: 6476fbe1f023f1e84092aadbaa308d7e6f3d4eb711e75f29e7c0c0706faa6c38
---

# Kimi K3 — efficiency mechanisms and serving economics

### Kimi K3 — efficiency mechanisms and serving economics
- **Kimi Delta Attention: 3:1 linear-to-full ratio, up to 75% less KV-cache memory**; 1M context **priced flat across the whole window** [SECONDARY] (morphllm.com).
- **Attention Residuals**: second efficiency mechanism; Moonshot credits KDA + Attention Residuals with **up to 6.3× faster decoding** [SECONDARY] (theairankings.com, citing kimi.com).
- **Kimi-Vendor-Verifier**: Moonshot-published harness for checking third-party provider fidelity [SECONDARY] (morphllm.com).
- Third-party serving: **Morph serves K3 at $2.50/$14 per M tokens with full 1M context and $0.29/M cached input** — under Moonshot's $3/$15/$0.30 list; **Anthropic Messages API natively**, so Claude Code works against it with two env vars; **100B+ tokens/day** on private deployments [SECONDARY] (morphllm.com).
- Open-weights escalation read: the open-to-closed gap has compressed from **6–9 months to ~3–5 months**, with K3 as the proof point [SECONDARY] (interconnects.ai via artokun/comfyui-mcp).

### Kimi K2.5 — third-party provider integration matrix (deployment/integrations)
- Provider model IDs observed for Kimi K2.5 (evidence of broad third-party hosting): **AWS Bedrock** (us-east-1, us-east-2, us-west-2, eu-north-1, ap-northeast-1, ap-south-1, ap-southeast-3, sa-east-1); **Fireworks** (`accounts/fireworks/models/kimi-k2p5`); **Azure AI** (`azure_ai/kimi-k2.5`, `azure_ai/FW-Kimi-K2.5`); **DeepInfra**; **Together AI**; **OpenRouter**; **Novita**; **W&B**; Cloudflare Workers (`@cf/moonshotai/kimi-k2.5`); Hugging Face Inference (`huggingface-llm-kimi-k2-5`) [SECONDARY] (cloudprice.net model-ID listing).
- CloudPrice hardware estimate: K2.5 needs **~2541 GB GPU memory at FP16** (estimated from parameter count) — no on-demand GPU instance fits at FP16; INT4 required for self-hosting [SECONDARY] (cloudprice.net — estimate, mark [UNVERIFIED] for capacity planning).


### New verified facts — expansion

### Kimi K2.5 / K2.6 / K2.7-Code — shared architecture spec
- **61 layers including one dense layer; 384 routed experts with top-8 routing per token; one shared expert** [SECONDARY] (MoonshotAI/Kimi-K2.5 GitHub; marktechpost.com, 2026-01-27; sglang cookbook for K2.7-Code — architecture shared across the K2.5/2.6/2.7 family).
- Attention: **hidden dimension 7,168; per-expert MoE hidden dimension 2,048; 64 attention heads** [SECONDARY] (MoonshotAI/Kimi-K2.5 GitHub; wilsonwu-ai/scaling-open-models).
- **160K vocabulary**; **MLA + SwiGLU** attention/activation design [SECONDARY] (MoonshotAI/Kimi-K2.5 GitHub).
- Vision: **MoonViT 400M** vision encoder from K2.5 onward (native image/video understanding added at K2.5) [SECONDARY] (MoonshotAI/Kimi-K2.5 GitHub; marktechpost.com, 2026-01-27).
- K2.5 continual pretraining used **~15T mixed visual/text tokens** on top of the K2 Base checkpoint [SECONDARY] (marktechpost.com, 2026-01-27 — single source, treat as [UNVERIFIED] until vendor confirms).
- K2.6 summarized as **1T total / 32B active** parameters in third-party provider docs [COMMUNITY] (benedictking/ccx kimi.md).

### K2.5 — vendor benchmark set and swarm claims
- Vendor-reported K2.5 scores: **SWE-bench Verified 76.8; MMMU-Pro 78.5; VideoMMMU 86.6; HLE Full with tools 50.2; BrowseComp 74.9** — all [VENDOR] unless independently rerun (marktechpost.com, 2026-01-27).
- Agent Swarm vendor claim: **~4.5× faster on wide research tasks** via native multi-agent parallel execution [VENDOR] (aibase.com; MoonshotAI GitHub).
- Existing §5 records hosted K2.5 retirement May 20 — community API docs add that **kimi-k2.5 and the earlier moonshot-v1 series were sunset at the end of August 2026**; as of September 2026 guides recommend checking the live list-models endpoint before pinning a model string in production [COMMUNITY] (dev.to, 2026-09).

