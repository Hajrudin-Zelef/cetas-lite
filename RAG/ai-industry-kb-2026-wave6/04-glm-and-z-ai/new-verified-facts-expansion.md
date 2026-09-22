---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/new-verified-facts-expansion
title: "New verified facts — expansion"
domain: glm-and-z-ai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Baseten", "China", "DeepSeek", "Fireworks AI", "Google", "Huawei", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "StepFun", "United States", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-01-08", "2026-02-14", "2026-04-01", "2026-04-07", "2026-06", "2026-06-13", "2026-07", "2026-07-02", "2026-08", "2026-08-14", "2026-08-18", "2026-08-26", "2026-08-27", "2026-08-28", "2026-08-29", "2026-09", "2026-09-01", "2026-09-07", "2026-09-09", "2026-09-20"]
keywords: ["agent", "agentic", "ascend", "attention", "benchmark", "benchmarks", "claude", "compute", "context window", "cost", "cyber", "cybersecurity"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1755, 1920]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: 79d104c7c2b21ea1455ad8abfc3b52263f206b730aeda283aae1957d1186c12c
---

# New verified facts — expansion

### New verified facts — expansion

### GLM-5.3-Flash — release, license, identity
- Z.ai released GLM-5.3-Flash on **August 26, 2026**, after testing it anonymously as **"Ox Alpha"** (Chinese: "Niu Lai") on OpenRouter and OpenCode; Z.ai confirmed the identity on **August 27, 2026** [SECONDARY] (MarkTechPost, 2026-08-26; industry.co.id/MEN, 2026-08).
- GLM-5.3-Flash shipped with **weights on Hugging Face on day one** under an **MIT license** — a deliberately different open posture from the delayed flagship release [SECONDARY] (MarkTechPost, 2026-08-26; startupfortune.com; industry.co.id/MEN). The official repository is `zai-org/GLM-5.3-Flash`; a BF16 variant `zai-org/GLM-5.3-Flash-BF16` is referenced as the conversion source by community quantizers [COMMUNITY] (RadixArk NVFP4 model card).
- The stealth-test strategy was deliberate: the lab collected real-world usage data at massive scale under the anonymous codename before committing to a public launch, generating organic developer attention [SECONDARY] (industry.co.id/MEN, 2026-08).
- During its anonymous week, Ox Alpha **topped leaderboards on OpenRouter and OpenCode**, prompting community speculation it was a Gemini competitor or another frontier lab's model [SECONDARY] (industry.co.id/MEN, 2026-08).
- The stealth preview was served **entirely on domestically produced Chinese AI chips using a custom SGLang-based inference stack** [SECONDARY] (MarkTechPost, 2026-08-26; YouTube/AISeeKing breakdown citing Z.ai confirmation).
- GLM-5.3-Flash is the **first natively multimodal model in the GLM-5 series**: text, image, and video input; text output; first GLM-family model with multimodal inputs on **Cloudflare's Workers AI** platform [SECONDARY] (MarkTechPost, 2026-08-26; industry.co.id/MEN, 2026-08).

### GLM-5.3-Flash — architecture and serving spec
- 320B total parameters / **18B active per token** (A18B); 1,048,576-token context window (exactly 2^20) [SECONDARY] (MarkTechPost, 2026-08-26; startupfortune.com; RadixArk model card).
- **45 decoder layers: 3 dense MLP layers + 42 MoE layers; 288 routed experts per MoE layer; one shared expert** [SECONDARY] (RadixArk/GLM-5.3-Flash-NVFP4 model card, describing the upstream BF16 checkpoint).
- Hybrid attention stack: **34 Kimi Delta Attention (KDA) linear-attention layers + 11 KPool-indexed DeepSeek Sparse Attention (DSA) layers**, per NVIDIA NeMo AutoModel documentation; only the sparse layers carry a growing KV cache — the 34 KDA layers use a fixed-size recurrent state, which is what makes long context cheaper to serve [SECONDARY] (startupfortune.com, citing NeMo AutoModel docs).
- Architecture additionally includes **MLA, manifold-constrained hyper-connections, a vision encoder, and a native MTP/NextN draft layer** [SECONDARY] (RadixArk model card reading of the upstream checkpoint).
- Routing top-k count for Flash was not disclosed in the sources reviewed; do not quote one [DIRECTIONAL].
- **Pretraining: ~30T multimodal tokens** [SECONDARY] (model-guide research note, github.com/mattrobenolt pi-configs; startupfortune.com context — treat as [UNVERIFIED] pending the vendor model card, single-source chain).
- Deployment support listed on the model card: **vLLM, SGLang, Transformers, Unsloth, KTransformers** [SECONDARY] (startupfortune.com; MarkTechPost, 2026-08-26).
- Self-hosting floor: default **FP8 checkpoint ≈ 306 GiB of weights before KV cache**; current vLLM path supports **NVIDIA Hopper and newer only**; realistic minimum is an **8-GPU node (or a GB200 tray at TP4)**; everyone below that consumes Flash as an API [SECONDARY] (MarkTechPost, 2026-08-26 — single source, mark [UNVERIFIED] for procurement decisions).
- Community quantization: **RadixArk/GLM-5.3-Flash-NVFP4** released on Hugging Face **August 28, 2026**, pinned to source revision `a6c167b6` of `zai-org/GLM-5.3-Flash-BF16`; converted to mixed-precision **NVFP4 W4A4 using NVIDIA Model Optimizer** with plain abs-max scaling and 256 tensor-scale normalization; card declares **MIT license** and **global deployment geography** [COMMUNITY] (RadixArk model card). A matching RadixArk NVFP4 quant of the GLM-5.3 flagship exists at `RadixArk/GLM-5.3-NVFP4` [COMMUNITY] (Hugging Face listing).

### GLM-5.3 flagship — staged weights release and custom license
- At the **August 14, 2026** launch, GLM-5.3 was **not yet downloadable**: weights and license terms were both unpublished, so "open" described a roadmap, not a file [SECONDARY] (emergent.sh guide, 2026-08).
- Z.ai said it would delay the public weights release by **about two weeks** for additional security checks and strengthened safeguards; the most sensitive cybersecurity functions were initially restricted to a **"trusted access" program** for selected partners and verified users [SECONDARY] (kr-asia.com; deeplearning.ai/The Batch, 2026-08).
- As of **August 28, 2026** the flagship's weights remained unpublished — the two-week window closed with the MIT-licensed Flash shipping instead [SECONDARY] (ainvest.com, 2026-08; emergent.sh).
- When the flagship weights did release, they came under a **new custom license (not MIT)**: companies with **more than $10 billion in annual revenue over 12 months must pass a Z.ai security review** before using the model or derivative works commercially [SECONDARY] (Techmeme summary of Frederic Lardinois/The New Stack, 2026-08-28; techbooky.com citing the same reporting). GLM-5.2's MIT license does not carry this gate [SECONDARY] (deeplearning.ai).
- The revenue gate follows the Llama-style pattern of keeping the largest rivals from free commercial advantage while preserving ecosystem distribution [DIRECTIONAL] (techbooky.com analysis).

### GLM-5.3 flagship — serving spec and cyber capability framing
- Text in up to **1M tokens**; text out up to **128,000 tokens**; API throughput claimed at **90 tokens/second** [SECONDARY] (deeplearning.ai/The Batch, 2026-08).
- Architecture: MoE transformer, **753B total parameters / 40B active per token** — consistent with the 753B tensor-class figure in existing §4 [SECONDARY] (deeplearning.ai).
- Features: adjustable reasoning levels (**low / high / max**), tool calling, structured output, streaming, context caching [SECONDARY] (deeplearning.ai).
- Z.ai's framing: GLM-5.3 was **not trained as a dedicated cybersecurity model**; cyber capabilities are described as **"emergent"** from expanded reinforcement learning and training in longer, more varied software-development environments [VENDOR via secondary] (kr-asia.com, 2026-08).
- Z.ai acknowledged the model is currently **better at early vulnerability stages — code review, discovery, verification — than at deeper exploitation or complete offensive/defensive operations** [VENDOR via secondary] (kr-asia.com, 2026-08).
- Training recipe details reported: **single-rollout asynchronous optimization** (trains on attempts one at a time instead of full batches); long agent-attempt records are split into **compacted segments** so the model learns from long-running tasks, not only short ones [SECONDARY] (deeplearning.ai).
- On Artificial Analysis' index the model scored **60 points and effectively ties Kimi K3 as open-weights leader**; it took the **best score among all models on CyberGym** (exploit-detection benchmark) in Z.ai's tests [SECONDARY reporting vendor tests] (deeplearning.ai — vendor numbers, label [VENDOR] for the CyberGym claim).

### GLM-5V-Turbo — closed commercial counterpart
- GLM-5V-Turbo (April 1, 2026) is **closed-source and API-only**: no downloadable weights, no MIT license, no self-hosting or fine-tuning [SECONDARY] (techtimes.com, 2026-07-02).
- Vision handled by the **CogViT** vision encoder for native image/video/document-layout processing [SECONDARY] (techtimes.com; MarkTechPost 2026-04-01; github.com/kyegomez/cogvit reference implementation).
- Reported spec: **200K context window, up to 128K max output**; **$1.20/M input / $4.00/M output** tokens [SECONDARY] (cometapi.com review, 2026).
- The open/closed split is deliberate product architecture: Z.ai kept vision capability in a closed commercial product while the open-weight flagship stayed text-only — a point of friction with the self-hosting community that adopted GLM-5.2 for its MIT license [SECONDARY] (techtimes.com, 2026-07-02; emergent.sh GLM-5.3 guide notes the same community demand for vision in an open checkpoint).

### Access routes and pricing-table mechanics (August 2026)
- GLM-5.3 per-token API access opened **August 18, 2026** at **$1.40 input / $4.40 output / $0.26 cached input per 1M tokens — identical to the GLM-5.2 rate** [SECONDARY] (emergent.sh, 2026-08).
- As of August 2026, Z.ai's **public pricing page still listed GLM-5.2 as its top row**; the confirmed 5.3 rate came through the API and gateways rather than a dedicated pricing-table line — re-verify before relying on it [SECONDARY] (emergent.sh, 2026-08).
- GLM Coding Plan: subscription for use inside Claude Code and Z.ai's **ZCode** development environment; **from $18/month on a points-based quota with an off-peak discount** [SECONDARY] (emergent.sh, 2026-08).
- deeplearning.ai/The Batch describes the Coding Plan range as **"$18 to $168 per month"**, which differs from the $18/$72/$160 tier figures in other sources — record as a **pricing-source conflict**, not a price change [SECONDARY] (deeplearning.ai vs llm-coding-benchmark pricing doc and pilot-shell blog).

### Z.ai / Zhipu capital-market facts
- Hong Kong IPO: **January 8, 2026**, priced at **HK$116.20/share**; ~**$558M raised** at ~**HK$51B / $6.5–6.7B** IPO valuation [SECONDARY] (caproasia.com, 2026-02-14; existing §4).
- First-day close **HK$131.50**, implying roughly **$7.4B** market value [SECONDARY] (caproasia.com, 2026-02-14).
- By **February 14, 2026**, shares were **+317.3% from the IPO price**; the company was founded in **2019**; February reporting also noted **Shanghai IPO plans** after the Hong Kong listing [SECONDARY] (caproasia.com, 2026-02-14).
- By **late June 2026** shares had risen **more than 2,000% from the January listing** as GLM-5.2 drew attention for approaching leading US models in coding and agent tasks [SECONDARY] (kr-asia.com, 2026-08).
- On **August 14, 2026** (GLM-5.3 launch day) the Hong Kong-listed shares **closed down 3.6%** — no launch rally, contrasting with the GLM-5.2 summer enthusiasm [SECONDARY] (kr-asia.com, 2026-08).
- Late-August 2026 market snapshot: ticker **2513**; market cap roughly **HK$500B (~$64B)** against about **HK$3.2B of consensus 2026 revenue and no earnings**; ainvest frames the valuation as resting on the claim that Z.ai leads the open-weights frontier [SECONDARY] (ainvest.com, 2026-08 — single source for the exact figures, mark [UNVERIFIED] for financial use).
- The company bills itself as the **world's first listed foundation-model company** — company self-description, not an independent classification [SECONDARY] (ainvest.com, 2026-08).

### Competitive board snapshot (Artificial Analysis, September 2026)
- On the independent board, **Step 5 Preview** (StepFun, announced September 20, 2026) and **GLM-5.3** sat one point apart: **44 vs 45 on AA Intelligence Index v4.3.2**, GLM-5.3 scored as "max" effort and Step 5 Preview carrying no effort label [SECONDARY] (orcarouter.ai, 2026-09-20).
- StepFun's Hugging Face repository contained only a **.gitattributes** file at announcement; the BF16 checkpoint was slated for **October 15** — described as a "top-three open-source result" but not yet open source at announcement, unlike GLM-5.3's downloadable weights [SECONDARY] (orcarouter.ai, 2026-09-20).
- Do not mix this AA **v4.3.2** reading with the v4.3 readings elsewhere in §4 [DIRECTIONAL].

## Figures and metrics

| Model | Date | Size (total/active) | Context | License | Provenance |
|---|---|---|---|---|---|
| GLM-5V-Turbo | 2026-04-01 | 744B / ~40B | 200K | closed / API-only | [VENDOR] |
| GLM-5.1 | 2026-04-07 | 744B / 40B | — | MIT | [VENDOR] |
| GLM-5.2 | 2026-06-13 (API); 06-16/17 (weights) | 744B-class | 1M | MIT | [VENDOR] |
| GLM-5.3-Flash | 2026 (stealth "Ox Alpha" → Flash) | same base as 5.2 | 1M | MIT | [SECONDARY] |
| GLM-5.3 | 2026 | identical base to 5.2 | 1M | GLM-5.3 License (bespoke) | [VENDOR] |

| Benchmark/price | Figure | Date | Provenance |
|---|---|---|---|
| GLM-5.3, AA Intelligence Index | ≈44.9 | v4.3, 2026-09-07 | [SECONDARY] |
| GLM-5.3-Flash, AA Intelligence Index | 42 | v4.3, 2026-09-07 | [SECONDARY] |
| GLM-5.3 API | $1.40/$4.40 per M in/out, 81% cache discount | 2026-08-18 | [VENDOR] |
| GLM-5.3-Flash API | $0.15/M input | 2026-08-26 | [VENDOR] |
| GLM Coding Plan | from $18/mo | 2026-06 | [VENDOR] |
| GLM-5V-Turbo benchmarks | vendor-only (no independent replication) | 2026-04 | [VENDOR] |
| Terminal-Bench 4.0, GLM-5.3 | 41.8% | official, 2026-09-01/02 | [SECONDARY] |
| TB 4.0 official leader (Fable 5.1) | 57.9%±3.8 | official, 2026-09-01/02 | [SECONDARY] |
| TB 4.0 vendor-reported (Mythos 5.1) | 60.9% | Anthropic self-report | [VENDOR] |

**No cross-version comparison:** GLM-5.3's 44.9 is v4.3-pinned; do not compare against v4.1.1/v4.2-era scores. TB 4.0 (41.8%) is not comparable with TB 2.1-era scores. [DIRECTIONAL]


### New verified metrics — expansion (continued — Z.ai capital figures)

- IPO: HK$4.35B ($558M); HK$116.2 offer → HK$131.5 close (+13.1%); ~$7.4B day-1 value [SECONDARY].
- Peak: HK$2,980; >HK$1T (~$128B) cap; +2,000% since IPO [SECONDARY].
- Follow-on: ~$4B via 19.8M shares @ HK$1,588–1,698; CICC sole coordinator [SECONDARY].
- 2024 revenue: 312.4M yuan (~$46M) [SECONDARY].


### New verified metrics — expansion (continued — GLM-5V-Turbo figures)

- 744B-A40B (753B on HF cards); CogViT; 200K ctx / 128K out; 30+ RL tasks; Design2Code 94.8; WebVoyager/AndroidWorld #1; SpeedBench 221.2 tok/s (#5); $1.20/$4.00; 100K Ascend 910B [VENDOR via secondary / SECONDARY].


### New verified metrics — expansion (continued — GLM-5.2 benchmark tables)

### Vendor/independent coding benchmarks (provenance-labelled)
- **SWE-bench Pro 62.1** vs GPT-5.5 58.6 vs Claude Opus 4.8 69.2 [SECONDARY] (techtimes.com; cryptobriefing.com).
- **Terminal-Bench 2.1: 81.0** — best open-source result at the time, near Claude Opus 4.8's 85.0 [SECONDARY] (cryptobriefing.com; toknow.ai).
- **FrontierSWE 74.4** vs Claude Opus 4.8 75.1 vs GPT-5.5 72.6 — evaluated by **independent third-party firm Proximal**, not Z.ai [SECONDARY] (techtimes.com).
- **MCP-Atlas 77.0** — nearly on par with Claude Opus 4.8's 77.8 [SECONDARY] (techtimes.com).
- **Artificial Analysis Intelligence Index: 51** — ranked **first among open-weights models** in AA's 9-eval composite; AA clocked **168.8 output tokens/sec** but flagged it as **token-hungry (~43K output tokens per task)**, inflating real cost above sticker price [SECONDARY] (pilot-shell blog).
- **AA-Briefcase (agentic knowledge work): Elo 1266 at $2.40/task** — between GPT-5.5 and Opus 4.8 (1356 at $10.40); Claude Fable 5 far ahead at 1587 [SECONDARY] (pilot-shell blog).
- **Semgrep IDOR cyber benchmark: 39% F1** (prompt-only, Pydantic-AI) — edging Claude Code on Opus 4.6 (37%) and Opus 4.8 (28%) at **~$0.17 per vulnerability**; Semgrep's caveat: "one task, one dataset, one run"; Sonnet 5 not tested [SECONDARY] (pilot-shell blog).
- **Code Arena: 2nd globally** on front-end web development, trailing only Claude Fable 5; **beat Claude Fable 5 on the crowdsourced Design Arena** design-task benchmark with **Elo 1360** [SECONDARY] (indianexpress.com).
- **Databricks July 2026 enterprise test**: multi-million-line internal coding test, GLM-5.2 on par with Claude Opus 4.8 at **$1.28/task vs $1.94 — ~34% cheaper** for equivalent quality [SECONDARY] (cryptobriefing.com).
- Price-to-performance: about **1/6 the cost of GPT-5.5** while surpassing it on several benchmarks [SECONDARY] (medium/@wenmingtech).
- GLM-5.3 vs 5.2 on SWE-bench Pro (independent table, starred = vendor/cited): **5.3 at 64.6 vs 5.2 at 62.1** — post-training-only improvement direction [SECONDARY] (intelligentliving.co, 2026-09).

### GLM-5.3-Flash — deep spec from the pi-configs evidence brief (2026-08-29)
- **A new base model, not a 5.3 post-training refresh** — this distinction matters for §4's model-lineage accuracy [COMMUNITY] (pi-configs evidence brief, citing z.ai blog and arXiv 2602.15763).
- Hybrid sparse+linear attention: **linear attention for local dependencies via state modeling, sparse indexer for global retrieval**; **IndexPool compresses 4 indexer key vectors into 1** [COMMUNITY] (pi-configs brief).
- Efficiency vs GLM-5.3: **3.0× less attention compute, 4.4× smaller KV cache** [COMMUNITY] (pi-configs brief).
- **mHC (Manifold-Constrained Hyper-Connections)**; **30T-token multimodal pre-train**; paper **arXiv 2602.15763**; **45 layers vs GLM-4.5's 92** [COMMUNITY] (pi-configs brief).
- Modalities: **vision + text**; confirmed live on Fireworks with `image_url` blocks accepted (alertmanager screenshot read correctly in testing) [COMMUNITY] (pi-configs brief, live probe 2026-08-29).
- Context enforcement on the Fireworks route: **prompt + max_tokens ≤ 1,048,576** — a 917,505-token prompt with max_tokens 131072 fails with a 400 that looks like a model failure [COMMUNITY] (pi-configs brief).
- **Thinking is always on**; Fireworks enum `low/medium/high/xhigh/max`; `none` rejected ("reasoning cannot be disabled"); no `minimal` [COMMUNITY] (pi-configs brief, probed 2026-08-29).
- Weights note: brief dated 2026-08-29 says **"MIT promised, not shipped at Fireworks launch"** — in tension with day-one-HF-weights reporting; record as a **source difference on weights timing**, not a contradiction of the license [COMMUNITY vs SECONDARY].
- Live on Fireworks **2026-08-29** as `fireworks/accounts/fireworks/models/glm-5p3-flash`; brief status: **"default-seat challenger"** [COMMUNITY] (pi-configs brief).
- Vendor claims in the brief (all [VENDOR]): **AA Intelligence Index 57 at $0.045/task** ("intelligence previously only available at roughly 10× the cost"); **outperforms GLM-5.2 across benchmarks at one-tenth the price**; **approaching Claude Opus 4.8** on coding/agentic benchmarks; **GLM Coding Plan quota: 3× the points of 5.3-proper**.
- Independent local eval (2026-08-29, Fireworks, max effort): routine pack **12/12 vs DeepSeek Flash 0731** ($0.0199 vs $0.0240, 19s vs 25s, 10.2K vs 22.4K output tokens); hard pack **4/4** ($0.0087 vs $0.0162); **needle ~925K: 2/2 PASS at $0.139 each** — 5.2 does this at ~$1.30, **5.3-proper cannot (hangs)**; vision 2/2 PASS; guess-vs-abstain **5/5 with 0 fabrications**; strict json_schema + thinking **10/10** [COMMUNITY] (pi-configs brief).
- Throughput (controlled stream, 2026-08-29): **~50–64 tok/s steady-state**, TTFT 0.13s warm / 1–4s cold; agentic wall-clock 36.5 tok/s (vs DeepSeek 81, 5.3-proper 22) [COMMUNITY] (pi-configs brief).
- OpenRouter field data: **50 tps / 1.03s TTFT / 99.84% completion** over the observed window [COMMUNITY] (pi-configs brief).
- Z.ai's own same-harness scorecard: **Flash DeepSWE 63.4 vs 5.3-proper 66.9; Terminal-Bench 2.1 84.3 vs 88.2; AutomationBench 48.8; Toolathlon 78.4; GDPval-AA 1773** — beating Opus 4.8 (1582), K3 (1685), Sol (1728); only Opus 5 (1852) higher [VENDOR via community brief].
- Per-success economics: Flash ~$0.15–0.25/task at 63.4% vs proper ~$1.50–2/task at 66.9% — the ~9× token premium buys ~3.5–4 pass points, favoring Flash **~8–10× per success** [COMMUNITY analysis] (pi-configs brief).
- Cross-provider throughput spread on identical weights — **Baseten 109 / Friendli 77 / Fireworks 50 tps** — proves throughput is **deployment-bound, not model-bound** [COMMUNITY] (pi-configs brief).
- Prefill already fast: **~23K tok/s on the 925K needle** [COMMUNITY] (pi-configs brief).
- Competitive notes: cheaper than **DeepSeek V4 Flash 0731** on both axes after its 2026-08-29 hike to $0.22/$0.66; DeepSeek retains a **384K output cap (vs 64K)** and **$0.007 cache-read (vs $0.029)** [COMMUNITY] (pi-configs brief).
- Positioning: the **only sub-$1 model on the route serving ~925K context** — takes 5.2's "marathon" role at ~1/9th the cost where <64K output suffices [COMMUNITY] (pi-configs brief).


### New verified metrics — expansion

### GLM-5.3-Flash — API pricing and launch promotion
- Standard API pricing: **$0.15/M input / $0.029/M cached input / $0.50/M output** tokens [SECONDARY] (model-guide research note; memeburn.com).
- Launch promotion: **$0.075/M input / $0.25/M output** through **September 9, 2026** [SECONDARY] (model-guide research note; startupfortune.com pricing discussion).
- Z.ai's positioning: Flash **beats GLM-5.2 across benchmarks and real workloads at roughly one-tenth the price** [VENDOR] (MarkTechPost, 2026-08-26, reporting Z.ai).
- Flash lands **within half a point of Claude Opus 4.8** on Z.ai's internal coding benchmark [VENDOR] (MarkTechPost, 2026-08-26; Medium/noahkenji283, 2026-08-28 — vendor's internal benchmark, not independently rerun).

### GLM-5.3-Flash — independent/community benchmark note
- On **KingBench**, GLM-5.3-Flash scored **63/80**, slightly **below its anonymous Ox Alpha score** — evidence that anonymous-preview scores and branded-release scores can diverge [COMMUNITY] (YouTube/AISeeKing breakdown, 2026-08).
- Strengths noted in community testing: reasoning, math, agentic coding, and local fine-tuning tasks [COMMUNITY] (YouTube/AISeeKing).

### GLM-5.3 flagship — serving and benchmark figures
- 753B total / 40B active; 1M-token input; 128K-token output; 90 tokens/s API throughput [SECONDARY] (deeplearning.ai/The Batch).
- AA Intelligence Index: **60 points**, effectively tied with Kimi K3 as open-weights leader [SECONDARY] (deeplearning.ai).
- **CyberGym: best score among all models** in Z.ai's own tests [VENDOR] (deeplearning.ai — vendor number).
- GLM-5.3 per-token API: **$1.40 / $0.26 cached / $4.40** per 1M in/cached/out (same as GLM-5.2), opened August 18, 2026 [SECONDARY] (emergent.sh).

### GLM-5V-Turbo — price and vendor benchmark figures
- API: **$1.20/M input / $4.00/M output**; 200K context / up to 128K output [SECONDARY] (cometapi.com, 2026).
- **Design2Code 94.8** [VENDOR] (techtimes.com reporting Z.ai; MarkTechPost 2026-04-01).
- **BridgeBench SpeedBench 221.2 tokens/s** [VENDOR] (cometapi.com reporting Z.ai figures — vendor number, not independently measured).
- Integrations claimed: **OpenClaw** and **Claude Code** agent frameworks [SECONDARY] (MarkTechPost, 2026-04-01).

### GLM Coding Plan — recorded price conflict
- Tier figures in circulation: **$18 / $72 / $160 per month** (llm-coding-benchmark pricing doc; pilot-shell blog) vs **$18 to $168 per month** range (deeplearning.ai/The Batch) vs promotional descriptions of **$3 or $10 Lite / $30 Pro / ~$80 Max** (older/promotional coverage).
- Treat as a **date-and-promotion conflict**, not a single current price: record the figure with its source date; the only consistently repeated current floor is **from $18/month** [DIRECTIONAL].

