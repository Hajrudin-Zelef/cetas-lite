---
id: etape4-trackb-local-inference/00-local-inference/4-model-library-notable-2026-additions
title: "4. Model library — notable 2026 additions"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Cohere", "DeepSeek", "Google", "Meta", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai"]
dates: ["2026-07-31", "2026-09-06"]
keywords: ["agent", "agentic", "agents", "claude", "compute", "copilot", "decode", "deepseek", "embedding", "embeddings", "fp4", "glm"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [418, 491]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 7b802e1ee9fdf34108cdc6a909acb22e07b56c0dc23d10c55ee7dd5c44df54aa
---

# 4. Model library — notable 2026 additions

## 4. Model library — notable 2026 additions

> Names verified via official `ollama.com/library/*` pages or official release notes — none guessed. Cloud models carry a `:cloud` suffix in the CLI/app (`gemma4:cloud`) and are suffix-less via the direct Cloud API (`gemma4:31b`) [official: docs.ollama.com/cloud].

### Cloud library (verified names, Sep 2026 snapshot)
Per community-maintained snapshots of the official catalog [secondary: grikomsn/ollama-cloud-copilot-chat (snapshot 2026-09-06), kirklasalle/prism]:

| Model (cloud name) | Params | Notes |
|---|---|---|
| `gpt-oss:120b`, `gpt-oss:20b` | 120B / 20B | OpenAI open-weight models; 131K context; thinking Low/Med/High [secondary] |
| `deepseek-v3.1:671b` | 671B | Reasoning, hybrid thinking [secondary] |
| `deepseek-v4-flash`, `deepseek-v4-pro`, `deepseek-v4.1-flash` | — | **Official library page verified**: `deepseek-v4.1-flash` is a multimodal MoE, **552B backbone, up to 1M-token context**, Causal Encoder-Decoder design, 8B active params/token prefill, 16B decode, FP4 main KV cache (890 bytes/token) [official: https://ollama.com/library/deepseek-v4.1-flash] |
| `kimi-k2:1t`, `kimi-k2-thinking` | 1T | Moonshot; agentic/coding [secondary] |
| `kimi-k3` | — | 1.049M context, images, tools; **requires Pro/Max subscription + extra usage credits** [secondary: official library https://ollama.com/library/kimi-k3 cited in third-party doc] |
| `kimi-k2.6`, `kimi-k2.7-code` | — | 262K context, images, tools [secondary] |
| `qwen3-coder:480b` | 480B | Agentic coding [secondary] |
| `qwen3.5:397b` (aka Qwen 3.5 397B) | 397B | 262K context, images, tools [secondary] |
| `glm-5.1`, `glm-5.2`, `glm-5.3`, `glm-5.3-flash` | — | Zhipu; 1.049M context (Flash adds images) [secondary] |
| `minimax-m2.7`, `minimax-m3` | — | 197K/512K context [secondary] |
| `gemma4:31b` | 31B | Google; 262K, images, tools; images+audio on MLX since v0.33.3 [official/secondary] |
| `mistral-large-3:675b` | 675B | 262K, images, tools [secondary] |
| `nemotron-3` (nano 30B / super / ultra) | — | NVIDIA; 262K [secondary] |

### Local library — 2026 additions
- **muse-glimmer** (30B, Meta Superintelligence Labs) — first model from Meta's new Superintelligence Labs; multimodal, agent-purpose-built; MLX engine with DFlash + image input; tags include `muse-glimmer:30b-mlx` [secondary: v0.32.7/v0.32.8 release notes].
- **gpt-oss:20b** — recommended 2026 local reasoning option, OpenAI-style [secondary: Medium Sep 2026 roundup].
- **qwen3-coder:30b** — best 2026 local coding model per roundups; 30B MoE, ~19 GB package, 256K context, needs ~24 GB VRAM [secondary].
- **qwen3.6:27b**, **qwen3.6:35b-a3b**, **qwen3:4b-instruct** [secondary].
- **devstral-small-2:24b** (Mistral coding-agent alternative) [secondary].
- **qwen3-vl:235b** (vision-language), **glm-4.6:14b** (agentic), **minimax-m2** (efficient coding) — listed in Jan 2026 official library pulse [secondary].
- **LFM2 / LFM2.5** (Liquid AI) — parser/render support added v0.30.9; community MLX ports note DFlash support [secondary].
- **Laguna** family (`laguna-xs.2` referenced in v0.30.0 known issues; NVFP4 fixes v0.32.4/v0.32.5) [secondary]. *Provenance of the Laguna model family itself (vendor) was not verified in this research — flagged [unverified].*
- **Cohere2Moe** architecture support added v0.30.9 [secondary].
- **deepseek-r1:14b** — compact reasoning pick in Sep 2026 roundups [secondary].

### Retirements (cloud)
- **Kimi K2.5 and MiniMax M2.5 retired 2026-07-31**; untagged `deepseek-v4-pro` / `deepseek-v4-flash` aliases removed from `/api/tags`, only dated `0731`/`0813` tags kept [secondary: grikomsn doc, citing official https://docs.ollama.com/cloud#retirements].
- Official docs confirm usage settings show upcoming retirements; downloaded local models are never affected [official: https://docs.ollama.com/cloud].

---

## 5. API features

### Native API (localhost:11434)
- `POST /api/generate`, `POST /api/chat`, `POST /api/pull`, `POST /api/create`, `GET /api/tags`, `DELETE /api/delete`, `POST /api/show`, `POST /api/embed`, `GET /api/ps` [secondary: whisperjav research doc].
- **`/api/tags` performance**: 3.1 s → 294 ms cold on large model libraries (v0.34.1) [official].
- **Web search (cloud)**: `POST https://ollama.com/api/web_search` with `query` (required), `max_results` (default 5, max 10); returns title/url/content snippets; API-key auth [secondary: uwuclxdy docs].
- **Bundled web search at launch**: OpenClaw web search integrated into Ollama startup since v0.21.2-rc0 (Apr 2026) [secondary].
- **Responses API**: OpenAI-compatible `/v1/responses` gained **web search support** in v0.32.11 (Aug 2026) [secondary].

### OpenAI compatibility
- Endpoints: `POST /v1/chat/completions`, `GET /v1/models`, `POST /v1/embeddings` [secondary]; any OpenAI-SDK code works with a base-URL change (dummy API key accepted) [secondary: Medium practical guide, Apr 2026].
- **Structured outputs** (JSON-schema enforcement), **tool calling / function calling**, streaming supported via the OpenAI-compatible surface [secondary: Medium guide].
- **Tool search** and **response compaction** for OpenAI-compatible clients added in v0.34.0 (Sep 2026) [official].
- Improved **structured-output performance on Apple Silicon** in v0.34.0 [official].

### Tool calling / agents
- `ollama launch` integrations (auto-install when missing, since v0.30.11): **claude** (Claude Code), **codex**, **pi**, **openclaw**, **hermes**, **dsh** (DeepSeek Harness, v0.32.11), **muse** (Muse Code, v0.32.11) [official/secondary].
- v0.32.0 added an **agent UI** (`cmd`) and warnings before old agent models [secondary].
- **Thinking/reasoning**: "max think level" documented (v0.30.11); Anthropic thinking streams fixed (v0.32.3); thinking-capability detection for opencode (v0.30.11); models expose thinking modes (e.g. Kimi/DeepSeek/GPT-OSS Low/High/Max) [official/secondary].

### Vision & audio & embeddings
- Vision: image input supported (e.g. gemma4 images+audio on MLX since v0.33.3 [official]; Muse Glimmer image input on MLX [secondary]; qwen3-vl, kimi-k3, deepseek-v4.1-flash multimodal [official/secondary]).
- Embeddings: `/v1/embeddings` + native embedding support; `nomic-embed-text` (lowercase behavior change in v0.30.0) commonly used [secondary].
- Direct Cloud API: point any client at `https://ollama.com` with `Authorization: Bearer <API key>`; supports OpenAI and Anthropic client shapes (subsets) [official: docs.ollama.com/cloud].

### Cloud access modes (two)
1. **Cloud-local hybrid**: `ollama signin` locally → pull `-cloud` models → local Ollama proxies to cloud [secondary].
2. **Direct Cloud API**: client pointed at `https://ollama.com` with API key; cloud model names omit the `-cloud` suffix (e.g. `gpt-oss:120b`) [secondary].
- Compute hosted primarily in the **US**, routable to **Europe and Singapore** for capacity [secondary: mindstone doc citing ollama.com/cloud, 9 Sep 2026].
- Privacy: "Ollama processes cloud prompts and responses to answer your requests. **We do not use them to train models**" [official: docs.ollama.com/cloud]; zero data retention claimed for cloud tiers [secondary].

---

