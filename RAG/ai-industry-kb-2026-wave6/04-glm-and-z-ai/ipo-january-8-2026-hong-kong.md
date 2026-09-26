---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/ipo-january-8-2026-hong-kong
title: "IPO: January 8, 2026, Hong Kong"
domain: glm-and-z-ai
role: deep-dive
task: finance
actors: ["Alibaba", "Anthropic", "China", "Google", "Huawei", "Hugging Face", "Meituan", "MiniMax", "OpenAI", "OpenRouter", "United States", "Xiaomi", "Z.ai"]
dates: ["2026-01-08", "2026-04", "2026-06"]
keywords: ["ipo", "agent", "agentic", "agents", "ascend", "benchmarks", "claude", "context window", "cost", "fable 5", "gemini", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1673, 1723]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: e6d0864641babb6d35058cac425d2dbfbd52d05f09a48e0502b93c986dad206a
---

# IPO: January 8, 2026, Hong Kong

### IPO: January 8, 2026, Hong Kong
- **Zhipu AI (Z.ai) went public on the Hong Kong Stock Exchange on January 8, 2026** — the **first of China's "six tigers"/"AI tigers"** to list, a day before rival **MiniMax's own market debut** (January 9) [SECONDARY] (business.inquirer.net; techxplore.com; techinasia.com).
- Oversubscribed IPO raised **HK$4.35 billion (US$558 million)**; shares rallied **~11.8–12%** in early trade on debut [SECONDARY] (business.inquirer.net; techxplore.com).
- Offering price **HK$116.2**, day-1 close **HK$131.5 (+13.1%)**, implying roughly **$7.4 billion market value** [SECONDARY] (caproasia.com).
- Company chairman: **Liu Debing**; founded **2019** (Tsinghua University professors **Tang Jie & Li Juanzi** per caproasia); key controlling shareholders Tang Jie & Liu Debing [SECONDARY] (techxplore.com; caproasia.com).
- Investor roster per caproasia: **Alibaba, Ant Group, Tencent, Meituan, Xiaomi, HongShan, Saudi Aramco & Prosperity7 Ventures** [SECONDARY] (caproasia.com).
- IPO proceeds earmarked for **developing general-purpose large AI models, including key algorithms and system infrastructure** [SECONDARY] (techxplore.com).
- Prospectus showed **2024 revenue of 312.4 million yuan (~US$46 million)** [SECONDARY] (techinasia.com).
- Liu Debing at the listing: "Zhipu is honored to stand at this historic juncture as a representative of China's large model sector"; told Bloomberg TV the company sees a trend of **computing costs for AI development "gradually decreasing"** [SECONDARY] (techxplore.com).
- Analyst read: Hello China Tech's **Poe Zhao** told AFP the two IPOs "demonstrate both the revenue potential and the fundamental challenges facing this new generation of LLM companies" [SECONDARY] (techxplore.com).
- Market context: China's LLM market estimated at **101.1 billion yuan (US$14.5 billion) by 2030** (Frost & Sullivan) [SECONDARY] (techxplore.com).

### Post-IPO valuation arc
- Stock surged **more than 2,000% since its January listing** (Seoul Economic Daily, ~June 2026); offering price HK$116.2 → stock as high as **HK$2,980**; market cap pushed **above HK$1 trillion (~US$128 billion / ~197 trillion won)** [SECONDARY] (en.sedaily.com).
- Bloomberg (via sedaily): Z.ai weighing a **multibillion-dollar share placement** ahead of the **six-month lock-up expiry (July 8)**; amount could "significantly exceed" the $558M IPO [SECONDARY] (en.sedaily.com).
- TechInAsia: Zhipu seeking **~US$4 billion via 19.8 million shares** at **HK$1,588–1,698** each (up to 13% discount to close); stock **+1,500% since listing**; sole coordinator **China International Capital Corp**; proceeds for **R&D, business expansion, investments, M&A, capital structure optimization, working capital** [SECONDARY] (techinasia.com).
- The GLM-5.2 rally: stock up **~50% in five days** amid praise for GLM-5.2 — AA Intelligence Index **51 points, 4th** after Fable 5, Opus 4.8, GPT-5.5; **first model from China — and among all open-source models — to surpass the 50-point mark** [SECONDARY] (en.sedaily.com).
- Cross-read: the HK$1T cap was reached **within a week of GLM-5.2's release** (per Tom's Hardware, cited by aqalion) — model performance and equity value are directly coupled in this market [SECONDARY] (aqalion.com).


### New verified facts — expansion (continued — GLM-5V-Turbo (native multimodal coding model))

### Release and positioning
- **Released April 1–2, 2026** (coverage dates; Z.ai announcement April 1 via marktechpost; launch write-ups April 2) [SECONDARY] (marktechpost.com; the-decoder.com; ubos.tech).
- **Z.AI's first multimodal coding foundation model**, purpose-built for vision-based coding: design-to-code, GUI agents, native vision for images/videos/designs [SECONDARY] (agentnativedev.medium.com; ubos.tech).
- **Closed-source, API-only** model (not part of the MIT open-weights line) [SECONDARY] (aqalion.com).
- Deep **OpenClaw and Claude Code integration** is a headline feature — the "perceive → plan → execute" loop for autonomous environment interaction [SECONDARY] (marktechpost.com; cometapi.com; ubos.tech).

### Architecture and spec
- **744B-A40B MoE backbone**: ~744B total parameters, ~40B active per token [SECONDARY] (aqalion.com; marktechpost.com).
- **Parameter-count note**: the aqalion review explicitly acknowledges the same 744B-vs-753B tension recorded in §4's contradictions — "Hugging Face model cards for the family list the confirmed total at 753B; treat 744B as Z.ai's official designation" [SECONDARY] (aqalion.com).
- **CogViT** vision encoder — a new encoder built by Z.ai for this model; images and text are processed together **from the start of training**, not as a bolted-on module [VENDOR via secondary] (the-decoder.com; marktechpost.com).
- **MTP**: marktechpost defines it as **Multi-Token Prediction** (multiple tokens predicted per inference step, speeding up output); ubos.tech defines it as **Multimodal-Task-Prompt** — record the expansion conflict [SECONDARY contradiction].
- **200K context window** (202,752 on OpenRouter); **128K max output** [SECONDARY] (cometapi.com; the-decoder.com; aqalion.com).
- **30+ Task Joint Reinforcement Learning** across vision-language-coding tasks (image captioning, VQA, code synthesis, STEM, grounding, video, GUI agents, coding agents) — maintains programming logic while scaling visual perception [VENDOR via secondary] (marktechpost.com; the-decoder.com; ubos.tech).
- Features: **thinking mode, streaming output, function calling, context caching** [SECONDARY] (the-decoder.com; aqalion.com).
- **Training hardware**: 100,000 Huawei Ascend 910B (MindSpore) — the same domestic-chip training claim as GLM-5.2 [SECONDARY] (aqalion.com).
- **Multimodal toolchain** (box drawing, screenshots, website reading incl. image understanding) completes the perception–planning–execution loop; **agentic meta-skills baked into pre-training** to strengthen action prediction early; a **multi-level, controllable, verifiable data system** was built to address the shortage of agent training data [VENDOR via secondary] (the-decoder.com).

### Benchmarks
- **Design2Code: 94.8** (headline, native design-to-code; cf. Opus 4.6 77.3 in the cometapi comparison table) [VENDOR via secondary] (cometapi.com; aqalion.com).
- SOTA on specialized agentic leaderboards: **CC-Bench-V2** (coding/repo exploration), **ZClawBench** (GUI agent interaction); #1 **WebVoyager / AndroidWorld** GUI-agent performance per the cometapi comparison [SECONDARY] (marktechpost.com; cometapi.com; ubos.tech).
- **#5 BridgeBench SpeedBench at 221.2 tokens/sec** — faster than Gemini 3.1 Pro, Claude Sonnet/Opus, and GPT 5.4 [SECONDARY] (agentnativedev.medium.com).
- **"Beats Opus 4.6 on multimodal benchmarks"** is a community/media headline, not an independently verified claim; one FAQ review explicitly cautions the Design2Code numbers are self-reported and unverified [SECONDARY/UNVERIFIED] (agentnativedev.medium.com; wavespeedai medium).

### Pricing
- **$1.20 input / $4.00 output per 1M tokens** [SECONDARY] (aqalion.com; cometapi.com; wavespeedai medium).
- Cost comparison (April 2026): **GPT-4o $2.50/$10.00** — ~2× cheaper input, 2.5× cheaper output; estimated **~$0.004 vs ~$0.027 per typical UI-coding task** [SECONDARY] (wavespeedai medium).
- Native **video input** (short clips alongside images/text in one context) — a stated differentiator over GPT-4o's frame-by-frame workaround [SECONDARY] (wavespeedai medium).
- Access: Z.ai native API (OpenAI-compatible) or **OpenRouter** [SECONDARY] (wavespeedai medium).

