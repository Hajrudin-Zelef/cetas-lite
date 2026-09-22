---
id: etape4-trackb-local-inference/00-local-inference/10-pricing-consolidated
title: "10. Pricing — consolidated"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: pricing
actors: ["AMD", "DeepSeek", "Nvidia", "OpenAI", "vLLM"]
dates: ["2025-07", "2026-09-22"]
keywords: ["pricing", "agent", "amd", "benchmark", "claude", "copilot", "cost", "deepseek", "funding", "gguf", "gpu", "inference"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [603, 687]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 5e7c90cd70fa7a64bf429dcd82583b00c7278bbdc4693ce00f92c0ff5730e6c0
---

# 10. Pricing — consolidated

## 10. Pricing — consolidated

| Offering | Price | Since / observed | Provenance |
|---|---|---|---|
| Ollama local (self-hosted) | **Free** (MIT open source) | standing | [independent] |
| Cloud Free | $0 (starter credits, starter models, 1 concurrent) | Sep 2026 | [secondary] |
| Cloud Pro | **$20/mo** ($60 usage, 3 concurrent, 50x free) | Sep 2026 (also $20/mo in Jul 2026) | [secondary] |
| Cloud Max | **$100/mo** ($300 usage, 10 concurrent, 250x free) | Sep 2026 | [secondary] |
| Cloud Team | **$500/mo** ($1,000 shared) — conflicts with $25/seat/mo (5-seat min) observed 28 Aug 2026 | Aug–Sep 2026 | [secondary, conflict flagged] |
| Extra usage credits | pay-as-you-go at per-token model rates | Sep 2026 | [secondary] |
| Enterprise | custom | 2026 | [secondary] |
| Off-peak discount | 50% off DeepSeek V4 (peak = 2x, 12:00–18:00 UTC Mon–Fri) | Sep 2026 | [secondary] |

---

## 11. Competitive positioning in 2026 coverage

### The fractured market (2026 consensus)
- "The one-size-fits-all approach to offline AI no longer works" — **LM Studio** owns the desktop/GUI lane, **vLLM** owns server throughput; Ollama remains the developer default [secondary: blazetrends, ~Jul 2026]. Source: https://blazetrends.com/lm-studio-and-vllm-challenge-ollama-for-local-ai-dominance-the-best-alternatives/
- Ranked comparison (Aug 2026, editorial scores): **#1 Ollama** (dev API + broad library; MIT; ~1-min setup), **#2 vLLM** (production concurrency, NVIDIA), **#3 LM Studio** (GUI discovery), #4 llama.cpp (engine-level control), #5 LocalAI [secondary: markaicode]. Source: https://markaicode.com/best/best-local-llm-inference-tools-production-2026/

### Ollama's documented weaknesses
- **Concurrency**: processes requests sequentially by default (llama.cpp heritage) — queues rather than batches; gap vs vLLM "small at one user and grows sharply past ~4–8 concurrent requests" [secondary: markaicode].
- **Throughput claim**: "vLLM pushes up to 2.6x the throughput of Ollama on the same hardware" [secondary: blazetrends — single source, treat as [unverified] magnitude].
- **No built-in auth / multi-tenancy**: "reasonable default for small production workloads… provided you put a gateway in front for rate limiting and auth, since Ollama has no built-in multi-tenant controls" [secondary: markaicode].
- **Memory**: "LM Studio vs Ollama 2026: 5x memory gap [tested]" headline (Sep 2026) claims LM Studio wins on memory efficiency for desktop use [secondary: tech-insider.org — headline-level, methodology not reviewed here].

### LM Studio
- Proprietary; **free for personal and commercial use since July 2025**; paid **enterprise tier with SSO/admin features** [secondary: tech-insider.org, Sep 2026].
- Versions: 0.4.20 (Jul 2026), 1.0.9 (Aug 2026) [secondary: localist news].
- Strengths: polished GUI model browser, chat UI, visual parameter controls, local OpenAI-compatible server; best for non-CLI teams and CI-adjacent desktop testing [secondary].
- Source: https://tech-insider.org/lm-studio-vs-ollama-2026/

### vLLM
- Production serving standard: batching, multi-user throughput, GPU servers; v0.26.0 released 26 Jul 2026 (411 commits, 212 contributors) [secondary: localist news].
- Maker **Inferact** raising as an open-source inference company alongside Ollama's round [independent: TechCrunch].
- Positioning: "Choose vLLM when local experimentation becomes real serving" [secondary: sabbirz].

### Ecosystem notes
- **llama.cpp** is Ollama's engine (direct since v0.30.0) and simultaneously its low-level competitor; 123,903 stars (Aug 2026) [secondary].
- **Open WebUI** (148,767 stars) is the standard chat frontend over Ollama [secondary].
- **OpenClaw** (Jan 2026 breakout) was Ollama's business proving point; bundled web search (OpenClaw provider) shipped in Ollama itself Apr 2026 [vendor-reported/secondary].
- 2026 hardware tailwinds for local AI cited with Ollama: AMD Strix Halo mini PCs (128 GB unified), Mac Mini M4, NVIDIA RTX Spark Windows-on-Arm builds [secondary].

---

## 12. Explicit uncertainties and gaps

1. **"Turbo" branding** — could not confirm a product literally named "Ollama Turbo" in 2026 sources; the cloud product is consistently "Ollama Cloud" with `:cloud` models. Possibly a legacy/coloquial name. [unverified — do not assert].
2. **Team pricing conflict** — $500/mo ($1,000 shared) vs $25/seat/mo (5-seat min), observed ~12 days apart (28 Aug vs 9 Sep 2026). Unresolved; likely a page/packaging change mid-transition to per-token billing. [flagged].
3. **Per-token transition date** — the GPU-time→per-token switch is bracketed (Jul 2026 → Sep 2026) but the exact announcement date was not recovered from an official source. [unverified].
4. **v0.31.x** — versions mentioned with a conflicting date (Apr 22 2026); no release-note content recovered. [unverified].
5. **52M monthly downloads (Q1 2026)** and 135K GGUF models on HF — single-source research notes, methodology unclear. [unverified].
6. **Laguna model provenance** — appears across official release notes (support added/fixed), but the model family's vendor/line was not verified. [unverified].
7. **Desktop app versioning** — no separate desktop-app version numbers found; app updates ship with engine releases. [gap].
8. **vLLM 2.6x throughput claim** — single secondary source, no methodology reviewed. [unverified magnitude].
9. **"5x memory gap" LM Studio vs Ollama** — headline only; methodology not reviewed. [unverified].
10. Off-peak pricing and per-model rates are from third-party docs verified against ollama.com/pricing on 8–9 Sep 2026 — accurate as of then, but the pricing page changes frequently (one source explicitly advises rechecking). [secondary, time-sensitive].

---

## 13. Key sources

- Official releases: https://github.com/ollama/ollama/releases/tag/v0.34.0 ; https://github.com/ollama/ollama/releases/tag/v0.33.3 ; https://github.com/ollama/ollama/releases/tag/v0.34.1 ; https://github.com/ollama/ollama/releases/tag/v0.30.11-rc0 ; https://github.com/ollama/ollama/releases/tag/v0.34.2
- Official docs: https://docs.ollama.com/cloud ; https://github.com/ollama/ollama/blob/HEAD/docs/windows.mdx
- Official library (sample verified page): https://ollama.com/library/deepseek-v4.1-flash
- TechCrunch (funding/adoption): https://techcrunch.com/2026/07/09/popular-open-source-ai-developer-tool-ollama-raises-65m-grows-to-nearly-9m-users/
- Adoption ranking: https://github.com/dappros/state-of-conversational-ai/blob/HEAD/2026-q3/report/report.md
- Pricing (secondary, Sep 2026): https://techtoheart.com/ollama-clouds-off-peak-pricing-turns-scheduling-into-a-cost-lever/ ; https://github.com/mindstone-agent/mindstone-agent/blob/HEAD/OLLAMA_CLOUD_QWEN3.5.md ; https://github.com/carldog/wobblebot/blob/HEAD/docs/implementation/ollama-cloud.md
- Enterprise/team (secondary): https://github.com/rizquuula/coding-plan/blob/HEAD/.claude/skills/provider-ollama/references/plans.md
- Comparisons (secondary): https://markaicode.com/best/best-local-llm-inference-tools-production-2026/ ; https://blazetrends.com/lm-studio-and-vllm-challenge-ollama-for-local-ai-dominance-the-best-alternatives/ ; https://tech-insider.org/lm-studio-vs-ollama-2026/ ; https://www.sabbirz.com/blog/ollama-vs-lm-studio-vs-llama.cpp-vs-vllm
- Release coverage (secondary): https://appselfhost.com/ollama-v0-21-2-rc0-released-bundled-web-search-and-structured-outputs-docs-land-in-new-rc/ ; https://github.com/bokiko/localist/blob/HEAD/news/2026-07.md ; https://github.com/bokiko/localist/blob/HEAD/news/2026-08.md ; http://dev.to/jtorchia/ollama-v0342-vs-v0341-a-targeted-mlx-fix-not-the-same-regression-mkh
- Model roundups (secondary): https://medium.com/@mmoufakkir2/the-best-local-ollama-models-in-2026-without-pretending-one-model-wins-everything-d1d2645737d6 ; https://github.com/grikomsn/ollama-cloud-copilot-chat/blob/HEAD/docs/models.md


---

## Part C — LM Studio + local inference hardware + GGUF quantization

**Research date:** September 22, 2026 (cutoff)
**All deliverables in English per project rule.** Provenance tags used throughout:
`[official]` = the vendor/project's own site, docs, blog, changelog · `[vendor-reported]` = numbers stated by the vendor about its own product · `[independent]` = independent review/benchmark/guide · `[secondary]` = aggregator, blog, or derived summary · `[unverified]` = could not be cross-confirmed; do not treat as fact.

---

