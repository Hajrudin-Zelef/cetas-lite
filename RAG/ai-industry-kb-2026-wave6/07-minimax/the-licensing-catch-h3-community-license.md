---
id: ai-industry-kb-2026-wave6/07-minimax/the-licensing-catch-h3-community-license
title: "The licensing catch (H3 Community License)"
domain: minimax
role: deep-dive
task: licenses
actors: ["Anthropic", "ByteDance", "China", "EU", "Google", "Hugging Face", "MiniMax", "OpenAI", "United States"]
dates: ["2025-09", "2026-05-26", "2026-06", "2026-06-01", "2026-06-07", "2026-08-03"]
keywords: ["license", "agent", "attention", "benchmark", "claude", "compute", "context window", "copyright", "cost", "gemini", "inference", "lawsuit"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3142, 3190]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 501858bf2b6e1b221e928ec7695d68a67d6e41871dbec4642aa8001d22e605c1
---

# The licensing catch (H3 Community License)

### The licensing catch (H3 Community License)
- Weights published to Hugging Face **August 3, 2026** — what shipped is **H3-Base, a 33.1B-parameter model generating natively at 768p** [SECONDARY].
- The license defines an **"Applicable Territory" for local deployment that specifically excludes the United States, the European Union, the United Kingdom, and South Korea** — users there are not licensed to run, modify, or deploy outputs from locally hosted H3 weights [SECONDARY].
- **Ryan Lee, MiniMax's Head of Developer Relations, confirmed publicly** that the US exclusion ties directly to the company's active copyright litigation [SECONDARY].
- Second stated reason: video models sit in a messier regulatory spot (EU AI Act, evolving UK/South Korean rules, unsettled US landscape); once weights are public the company cannot enforce safeguards downstream, so license restriction became the alternative to delaying release [SECONDARY].
- **The 2K output in every promotional clip comes from H3-Regenerate-2K — a separate module that was NOT open-sourced and remains API-only**; even a fully licensed self-hosted setup must call MiniMax's servers for the marketed resolution [SECONDARY].
- Free for organizations under **$20M annual revenue** (with **"MiniMax H3" in the product interface**), subject to individual licensing for excluded territories; **bans using H3 outputs to train or improve a competing model** — a restriction that applies everywhere [SECONDARY].

### The copyright lawsuit underneath
- **Disney, Universal, and Warner Bros. Discovery jointly sued MiniMax and its Hailuo platform in September 2025**, alleging the company trained its video/image models on **unauthorized copies of their characters** and that Hailuo could generate recognizable **Spider-Man, Darth Vader, and Shrek** from simple text prompts [SECONDARY].
- MiniMax sought dismissal on jurisdictional grounds (US courts have no authority over a Chinese company); on **May 26, 2026, a federal judge rejected that argument** and let the case proceed into **full discovery** [SECONDARY].
- **H3 runs on the same Hailuo infrastructure named in the lawsuit** [SECONDARY].

### Pricing and market reaction
- MiniMax's claim: **2K generation costs less than a third of mainstream competitors; 768p costs less than half of what competitors charge at 720p** (competitors unnamed) [VENDOR via secondary].
- Independent breakdown: **H3 2K ~$7.80/min vs Kling 3.0 ~$20.16/min (1080p) vs ByteDance Seedance 2.0 ~$22.45/min (1080p)** — "lines up reasonably well" with MiniMax's claim; treat exact dollars as directional (one explainer, not a controlled benchmark) [SECONDARY].
- AA independent scores: **H3 ranked 1st globally in video editing at launch**; **top 3 in text-to-video and image-to-video**; trailed **Google's Gemini Omni Flash** on T2V and **Gemini Omni Flash + Seedance 2.0** on I2V specifically [SECONDARY].
- Stock: MiniMax shares climbed **roughly 13% intraday on heavy turnover** on the H3 news; **ByteDance rolled out Seedance 2.5 the same week** [SECONDARY].

### Limitations (all [SECONDARY])
- **15 seconds is a hard ceiling** per clip; Seedance 2.5's 30-second clips are a real advantage for narrative work.
- 2K requires an API call even for local open-weights users.
- US/EU/UK/SK users need an **individual license** for legal self-hosting.
- Every API call routes through **infrastructure in China**, subject to China's National Intelligence Law.
- Vendor pricing claims/demos are promotional; **AA scores reflect the hosted API pipeline, not a locally run H3-Base render**.
- **No independent security audit of the released weights**; training-data provenance not independently verified.


### New verified facts — expansion (continued — M3 launch day, stock story, Kilo audit, MSA internals, MiniMax Agent)

### Launch day and the stock market (pasqualepillitteri.it, June 2026)
- **June 1, 2026**: MiniMax stock **climbed more than 5%** on the Hong Kong Stock Exchange, touching **907.5 Hong Kong dollars**, then **collapsed to a 12.38% loss within a few hours** on turnover of **1.49 billion Hong Kong dollars** — classic "sell the news" behavior [SECONDARY] (pasqualepillitteri.it).
- Same-day headline: M3 billed as the **first open-weights system to combine frontier coding, a 1M-token context window, and native multimodality** [SECONDARY] (pasqualepillitteri.it).
- Vendor claim at launch: **59.0% SWE-Bench Pro** — per MiniMax above **GPT-5.5 and Gemini 3.1 Pro**, close to **Claude Opus 4.7** [VENDOR via secondary] (pasqualepillitteri.it).
- **Open-weights caveat at launch**: at announcement the model was accessible via **API and MiniMax Agent**; actual weights were **not yet public** — MiniMax promised weights plus the full technical report **within roughly 10 days** of launch. The article frames it as a statement of intent, not a delivered fact [SECONDARY] (pasqualepillitteri.it).

### Kilo independent audit (June 7, 2026)
- The Kilo team gave **the same code and the same prompt** to **MiniMax M3 and Claude Opus 4.8**: find **17 real bugs** deliberately planted in a webhook delivery service (TypeScript, Bun, SQLite) [SECONDARY] (pasqualepillitteri.it).
- Result: **M3 found 13/17 bugs at $0.07** in 5m 03s; **Opus 4.8 (medium) found 13/17 at $1.30** (3m 53s); **Opus 4.8 (high) 13/17 at $1.93** (4m 34s); **Opus 4.8 (xhigh) 15/17 at $2.03** (7m 26s); **Opus 4.8 (max) 15/17 at $3.39** (9m 24s) [SECONDARY].
- Cost ratios: Opus at mid-tier costs **18–27× more** for the same 13 bugs; at max it wins by 2 bugs at **up to 48× the per-run cost** [SECONDARY].
- What M3 missed (caught by Opus at higher reasoning levels): **invalid JSON returning a 500 error, database setup running at import time, async callback inside a synchronous transaction** — subtle bugs separating quick review from exhaustive review [SECONDARY].
- The article's practical read: for high-volume/repeated audits M3 delivers the same harvest at a fraction of the price; for maximum coverage with unconstrained budget, Opus 4.8 at full reasoning stays ahead [SECONDARY].

### MSA internals ("KV outer gather Q")
- MSA = **MiniMax Sparse Attention**; each token attends only to the tokens that matter; compute per token at 1M context **drops to 1/20th of the previous generation** [VENDOR via secondary] (pasqualepillitteri.it).
- Decoding **up to 15.6× faster**, prefill **more than 9× faster** than M2 on very long contexts [VENDOR via secondary] (pasqualepillitteri.it).
- **Why MSA instead of lightning attention**: MiniMax's previous models used "lightning attention" (a linear variant); M3 returns to sparse attention with a new implementation called **"KV outer gather Q"**, in which each memory block is **read only once** and accesses are **contiguous** — claimed at **more than 4× the gain of Flash-Sparse-Attention**, a reference implementation [VENDOR via secondary] (pasqualepillitteri.it).
- Inference efficiency as the control point: "whoever controls this layer controls the costs" — the same lesson the article ties to antirez's local-inference work [SECONDARY analysis] (pasqualepillitteri.it).

