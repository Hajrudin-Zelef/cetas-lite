---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/k3-weights-and-spec-felloai
title: "K3 weights and spec (felloai)"
domain: kimi-and-moonshot-ai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "DeepSeek", "Fireworks AI", "Hugging Face", "Microsoft", "MiniMax", "Moonshot", "Nebius", "OpenAI", "OpenRouter", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-07", "2026-07-25", "2026-07-27", "2026-08-02"]
keywords: ["agent", "agentic", "attention", "benchmarks", "claude", "copilot", "cost", "decode", "deepseek", "fable 5", "fine-tuning", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2169, 2241]
section: "§5. Kimi and Moonshot AI"
sha256: 831b23b83b4cce57aa61e3d42fe7ed94d01790992c9239557e1167de5950bf5a
---

# K3 weights and spec (felloai)

### K3 weights and spec (felloai)
- Weights published **July 27, 2026** (11 days after the Kimi-app launch), **96 safetensors shards totaling ~1.56 TB** on Hugging Face [SECONDARY] (felloai.com).
- **2.8T total / 104B active per token**; **1,048,576-token context**; **text, image, and video input**; always-on reasoning with **low/high/max effort** settings [SECONDARY] (felloai.com).
- **MXFP4 weights via quantization-aware training** [SECONDARY] (zeronoise.ai).
- AA Intelligence Index **57 — top open-weight result**, with GLM-5.2 next at 51 [SECONDARY] (felloai.com).

### K3 serving economics vs DeepSeek V4.1 Flash (orcarouter.ai)
- **Cost per AA Intelligence Index task: $0.27 (V4.1 Flash) vs $2.00 (K3) — 7.4×**, described as "the number that survives every version of this comparison" because token counts don't depend on harness [SECONDARY].
- On Vals AI's agentic coding bench the gap widens to **~40×: $0.41 vs $17.59**, with the Flash run finishing in **~15 minutes vs ~90 minutes** — the more agentic the workload, the more caching and output volume dominate the bill [SECONDARY].
- Rate: Flash **$0.15/$0.60 off-peak** ($0.30/$1.20 peak, weekday 01:00–04:00 and 06:00–10:00 UTC) vs K3 **$3.00/$15.00 flat** [SECONDARY].
- Cache discount: **98% vs 90%**; cache read **$0.003 vs $0.30 per 1M** [SECONDARY].
- Context: **1,048,576 both, neither applies a long-context surcharge** [SECONDARY].
- Throughput: **214.4 vs 34.7 tok/s (6.2×)**; verbosity: **250M vs 160M output tokens** across the index [SECONDARY].
- Parameters: Flash **552B total, 8B active on prefill / 16B on decode** vs K3 **2.8T, 16 of 896 experts active** [SECONDARY].
- Sampling: Flash exposes full temperature/top_p/seed; **K3 exposes reasoning_effort only and thinking cannot be disabled** [SECONDARY].

### K3 — known limitations (techtimes, July 25, 2026)
- Speed: **~32.6 tok/s vs 70.9 median** for comparable models; **TTFT 161.17s vs 2.87s median** [SECONDARY].
- **vLLM requires updates incorporating KDA support** before K3 serves efficiently; **standard prefix caching does not map to KDA's architecture** [SECONDARY].
- **Full technical report with training details not yet published** (as of July 25); conventional toolchain integrations through **LangChain and LlamaIndex expected but not validated**; **production-stable local inference likely a Q4 2026 story** for most teams [SECONDARY].

### K3 — day-0 ecosystem and arena results
- Day-0 support from **Together, Fireworks, Nebius, vLLM, SGLang, Cursor, and Devin**; third-party blended pricing **$5.40 per M tokens** — matching Kimi's direct pricing [SECONDARY] (zeronoise.ai).
- **K3 (Max) #1 among open-weight models in Agent Arena (+9.75% net improvement)** and **#1 overall in Frontend Code Arena** [SECONDARY] (zeronoise.ai).
- **Frontend Code Arena first place at roughly a third of Fable 5's price** ($3/$15 vs Fable 5 rate) [SECONDARY] (api-evangelist/hashnode).
- Moonshot's own launch presentation is **more candid than K2.7 Code's**: it explicitly says K3 **still trails Fable 5 and GPT-5.6 Sol overall** [SECONDARY] (mattrobenolt/pi-configs).
- Harness critique (same source): **SWE-Marathon used an H20-calibrated branch** rather than the final stock environment, and **Fable 5 fell back on 35% of tasks** — vendor margins are directional, not clean rankings; but unlike K2.7 Code, K3's central coding claim **survives an independent common-harness rerun** [SECONDARY].
- analyticsindiamag frames K3 alongside Moonshot **building its own chip** — the headline claims in-house silicon; details were not verified in the retrieved material, so treat as [UNVERIFIED].


### New verified facts — expansion (continued — Kimi developer ecosystem integrations)

### Kimi Code in VS Code Copilot Chat (rafled-com/kimi-openrouter-chat-provider)
- Community extension puts the **Kimi Code subscription** (Moonshot AI's **Kimi K2.7 Code and Kimi K3**) directly in **VS Code Copilot Chat** — agent mode with tools, without the official Kimi Code extension [COMMUNITY] (github.com/rafled-com/kimi-openrouter-chat-provider).
- Talks to Kimi Code's **OpenAI-compatible API (`https://api.kimi.com/coding/v1`)** with a Kimi Code API key — the same setup Moonshot documents for Claude Code and Roo Code [COMMUNITY].
- **K3 effort variants exposed as separate picker entries** (`K3 · Low / High / Max`) — the only community integration surfacing K3's effort tiers in-editor [COMMUNITY].
- Thinking shown as a *Thinking* block and **preserved across turns** (Kimi's preserved-thinking models expect it during multi-turn tool calling); quota (5-hour and weekly windows) shown in the status bar; command-safety verdicts (green/amber/red) above proposed terminal commands; **secret protection** withholding `.env`/keys from the model [COMMUNITY].

### copilot-api (caozhiyuan) — Kimi as a Copilot provider lane
- `copilot-api auth login --provider kimi` adds Kimi with **baseUrl `https://api.kimi.com/coding`** — noted as serving **both the Anthropic and OpenAI-compatible endpoints** from the same base URL [COMMUNITY] (github.com/caozhiyuan/copilot-api).
- Cross-check: generate-commit-extension independently documents **`/v1/messages` and `/v1/models` working on `api.kimi.com/coding`** (verified 2026-08-02) — two independent community sources for the endpoint [COMMUNITY].

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

