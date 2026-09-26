---
id: ai-industry-kb-2026-wave6/05-kimi-and-moonshot-ai/k3-license-evidence-now-resolves-the-5-conflict
title: "K3 license — evidence now resolves the §5 conflict"
domain: kimi-and-moonshot-ai
role: deep-dive
task: licenses
actors: ["Alibaba", "Anthropic", "DeepSeek", "Fireworks AI", "Hugging Face", "Microsoft", "Moonshot", "Nebius", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-07-25", "2026-07-27", "2026-08-02"]
keywords: ["license", "sol", "agent", "agentic", "apache", "claude", "consumer", "copilot", "cost", "decode", "deepseek", "distribution"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2162, 2210]
section: "§5. Kimi and Moonshot AI"
delta_of: ai-industry-kb-2026
sha256: 89eed05d194d0f6466e20937afa284e577f61e184f653aed3417b38073c2e0ba
---

# K3 license — evidence now resolves the §5 conflict

### K3 license — evidence now resolves the §5 conflict
- The **mattrobenolt pi-configs research note** (updated ~5 days before this writing) inspects the primary text and reports: weights use the **bespoke Kimi K3 License, not MIT or Apache**. Terms: ordinary use, modification, distribution, fine-tuning, and commercial products allowed; a **MaaS business whose group revenue exceeds $20M over any consecutive 12 months needs a separate Moonshot agreement**; products exceeding **100M MAU or $20M monthly revenue must display "Kimi K3" prominently**; conditions **do not apply to internal use or access through Moonshot/certified inference partners**. Primary text: `https://raw.githubusercontent.com/MoonshotAI/Kimi-K3/main/LICENSE` [SECONDARY] (github.com/mattrobenolt/pi-configs).
- Corroborated by three more sources: **analyticsindiamag** ("custom Kimi K3 License ... blends open access with targeted commercial restrictions"); **felloai** ("custom Kimi K3 License rather than MIT"); **techtimes** ("Kimi K3 License") [SECONDARY].
- The "Modified MIT" label persists in two sources (**orcarouter.ai**, **zeronoise.ai**) — but orcarouter's description ("adds branding terms above 100M monthly users or $20M monthly revenue") matches the bespoke license's substance, so this is a **naming difference, not a substance difference** [SECONDARY].
- K2-line context (techtimes): the **K2 line used a Modified MIT license** with the same 100M-MAU/$20M-revenue display term — K3 appears to have moved from Modified MIT (K2) to the bespoke Kimi K3 License (K3). Treat the §5 license conflict as **largely resolved in favor of the bespoke custom license** [DIRECTIONAL].
- analyticsindiamag's comparison framing: unlike **DeepSeek's fully permissive MIT** and unlike **Qwen's commercial paywalls at high traffic**, K3 lets consumer services over 100M MAU/$20M monthly revenue use the model freely **provided they prominently display "Kimi K3" branding**; **internal enterprise deployments remain completely free and exempt from revenue caps regardless of company size**, as long as outputs are not resold as an external API [SECONDARY].

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

