---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/k2-6-swarm-scale-and-deprecation-mechanics
title: "K2.6 — swarm scale and deprecation mechanics"
domain: kimi-and-moonshot-ai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Moonshot"]
dates: ["2026-04", "2026-04-20"]
keywords: ["agent", "agents", "agi", "attention", "benchmark", "claude", "copilot", "int4", "kimi", "mcp", "mxfp4", "opus 4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2292, 2312]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: 1d648ad5e4f27936f4dad8a83526b6448daa9cf1a44bff40302aa2f29d944af3
---

# K2.6 — swarm scale and deprecation mechanics

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

