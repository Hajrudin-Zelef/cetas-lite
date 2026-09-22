---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/kimi-k3-release-mechanics
title: "Kimi K3 — release mechanics"
domain: kimi-and-moonshot-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "China", "Fireworks AI", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "Together AI", "Z.ai", "xAI"]
dates: ["2026-01-27", "2026-04", "2026-04-20", "2026-07-16", "2026-07-26", "2026-07-28", "2026-08", "2026-09", "2026-09-07", "2026-09-21"]
keywords: ["kimi", "agent", "agentic", "agents", "agi", "attention", "aws", "bedrock", "benchmark", "claude", "copilot", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2242, 2312]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: ea971e65999e7cee2bf68f70aa7f94a91c762f3305e892ae1612c4d60d3cc73b
---

# Kimi K3 — release mechanics

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

### K2.6 — swarm scale and deprecation mechanics
- Agent Swarm scale: **up to 300 sub-agents and 4,000 coordinated steps** for long-horizon coding [SECONDARY] (marktechpost.com, 2026-04-20).
- Launch API pricing at April 2026 release: **$0.60 input / $2.50 output per 1M tokens** — roughly 8.3×/10× cheaper than Claude Opus 4.7 at the time [SECONDARY] (medium.com/@tentenco review, 2026-04). Current pay-as-you-go list price is **$0.95 input / $0.16 cached / $4.00 output** — record as price evolution, not a contradiction [SECONDARY] (saygm.com/blog/kimi-api-pricing; gemmaclaw docs).
- Deprecation mechanics on Moonshot's platform: model list and sunset notices live at the platform docs models page; pricing numbers live in per-model pages (`pricing/chat-k3.md`, `chat-k27-code.md`, `chat-k26.md`, `chat-k25.md`, `chat-v1.md`) — **a pricing page can outlive its model card** (K2.5's card was removed while `chat-k25.md` still served) [COMMUNITY] (enricoros/big-agi llms notes).
- **`kimi-k2` was retired 2026/05/25** with migration guidance to `kimi-k2.6` [COMMUNITY] (benedictking/ccx kimi.md — community mirror of vendor notice).

### K2.7 Code — spec and vendor benchmark table
- Same K2.5/K2.6 architecture; **native INT4** weights; **MoonViT 400M** vision encoder with image/video support marked **experimental** [SECONDARY] (bytedance-iaas sglang cookbook; byteiota.com).
- **Mandatory/preserved thinking** with **~30% fewer thinking tokens than K2.6** [VENDOR] (byteiota.com reporting vendor).
- Vendor benchmark table (all [VENDOR] unless independently rerun): **Kimi Code Bench v2 62.0; Program Bench 53.6; MLS Bench Lite 35.1; Kimi Claw 24/7 Bench 46.9; MCP Atlas 76.0; MCP Mark Verified 81.1** [SECONDARY reporting vendor] (byteiota.com).
- API list prices: **K2.7 Code $0.95 / $0.19 cached / $4.00**; **K2.7 Code Highspeed $1.90 / $0.38 / $8.00** per 1M input/cached/output; Highspeed is ~5–6× faster output with higher usage, gated to Allegretto+ membership [SECONDARY] (saygm.com; benedictking/ccx kimi.md [COMMUNITY] for the speedup factor).

### Kimi K3 — architecture spec (full)
- **2.8T total / 104B active** parameters [SECONDARY] (siml1169/kimi-copilot-provider_k3; ai-stack.ai; felloai.com; glows.ai).
- **93 layers: 69 KDA (Kimi Delta Attention) + 24 Gated MLA; one dense layer** [SECONDARY] (siml1169/kimi-copilot-provider_k3; ai-stack.ai).
- **896 routed experts, 16 selected per token, 2 shared experts** [SECONDARY] (siml1169/kimi-copilot-provider_k3; felloai.com).
- Vision: **MoonViT-V2 ≈ 401M** parameters [SECONDARY] (ai-stack.ai; glows.ai).
- Quantization format: **MXFP4 weights / MXFP8 activations** [SECONDARY] (ai-stack.ai; glows.ai).
- **1,048,576-token context**; text/image/video input [SECONDARY] (felloai.com; ai-stack.ai).
- Infrastructure released alongside the weights: **MoonEP, FlashKDA, AgentEnv** [SECONDARY] (ai-stack.ai; glows.ai).

