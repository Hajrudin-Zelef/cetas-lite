---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-36
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Baseten", "Cohere", "EU", "Fireworks AI", "Google", "Hugging Face", "Nvidia", "OpenAI", "Sakana", "Together AI", "United States"]
dates: ["2025-10", "2026-03", "2026-05", "2026-06", "2026-07"]
keywords: ["agent", "agentic", "apache", "astra", "benchmark", "benchmarks", "blackwell", "claude", "cohere", "consumer", "context window", "cyber"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1452, 1500]
sha256: 68690cbac0d1772fde23707b38410efe0b4eb62aa0d07e39fcd886c595944f43
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

**Products:**
- **Tinker** — fine-tuning platform / training API, launched October 2025 [vendor-reported/secondary]. Lets companies fine-tune base models on their own data and own the resulting model. Cited enterprise customer: hedge fund Bridgewater Associates, which reportedly fine-tuned an open model via Tinker, scoring 84.7% on (company-evaluated) financial reasoning tests [vendor-reported].
- **Inkling** (released 15 July 2026) — Thinking Machines' first in-house foundation model, and **a model, not an assistant or codename** [vendor-reported/secondary, multiple outlets] https://fourweekmba.com/ai-thinking-machines-inkling-open-weights-customization-strateg/:
  - MoE transformer; 975B total parameters, 41B active per forward pass; pretrained on 45T tokens (text, images, audio, video); up to 1M-token context; native text/image/audio reasoning, outputs text + code; controllable "thinking effort" dial.
  - **Open-weight release** on Hugging Face (Apache 2.0 per secondary coverage; includes a Blackwell-optimized checkpoint). The lab explicitly says Inkling is "not the strongest overall model available today" — positioned as a broad, balanced base for customization, flagship model for Tinker [vendor-reported].
  - Benchmark claims: e.g. matches Nvidia Nemotron 3 Ultra on one coding benchmark at ~1/3 the tokens [vendor-reported; independent evaluation not yet published].
  - Inference live on Tinker, Together AI, Fireworks, Modal, Databricks, Baseten [secondary].
  - Access/pricing: reported as free of charge on Tinker at launch [secondary]; no per-token list pricing located. Mark: **pricing unconfirmed**.
- **Inkling-Small** — previewed alongside Inkling: 276B total / 12B active parameters; full weights planned after testing [vendor-reported/secondary].
- No consumer assistant product announced as of Sept 2026.

### 10.6 Sakana AI — detailed (from Track C)

**Context / funding:** Tokyo-based lab; became Japan's first AI unicorn within ~14 months of founding (Jul 2023); Series A Sept 2024; $135M Series B Nov 2025 at $2.65B valuation; investors include Google, Nvidia, In-Q-Tel, Khosla Ventures, Japanese financial institutions [secondary — pre-window, context only]. No new 2026 funding round found.

**2026 research releases:**
- **AI Scientist v2 → Nature publication (March 2026)**: full AI Scientist system published in *Nature* with UBC, Vector Institute, Oxford [secondary]. Peak of its automated-research program.
- **Darwin Gödel Machine**: self-improving coding AI system (Gödel + Darwin reference) [secondary].
- **ICLR 2026 papers** underpinning Fugu: TRINITY (role assignment across multiple LLMs) and Conductor (RL-discovered natural-language coordination strategies) [vendor-reported].
- **PC-ALM (14 Sept 2026)**: blog post + arXiv paper + open-source code on backprop-free deep learning (predictive coding with augmented Lagrangian + "dual neurons"), claiming training of very deep nets with only local computation [vendor-reported/secondary]. Independent reproduction not yet published [secondary].

**2026 products:**
- **Namazu** — Japanese-language LLMs tuned for Japanese linguistic/cultural context [secondary].
- **Sakana Marlin** — launched 15 June 2026: autonomous strategic research agent for B2B, framed as "Virtual CSO" (up to ~8 hours continuous long-horizon reasoning; drafts strategy plans + slide decks from a single prompt) [secondary].
- **Sakana Fugu** — launched 22 June 2026 (model ID fugu-ultra-20260615): multi-agent orchestration delivered as a single OpenAI-compatible API. The orchestrator is itself a language model that routes sub-tasks to a pool of frontier models (at launch: e.g. Claude Opus 4.8, GPT-5.5, Gemini 3.1 Pro per secondary coverage) and synthesizes results [vendor-reported/secondary]. Regions: JP, US, most non-EU (EU/EEA/UK/Switzerland blocked) [secondary]. Subscription $20/$100/$200/mo tiers; Fugu Ultra pay-as-you-go $5/$30 per 1M in/out tokens, $0.50/M cached input, higher rates above 272K context [secondary].
- **Fugu-Cyber** (~25 July 2026): security orchestration variant; claims 86.9% on CyberGym (vs Anthropic Claude Mythos Preview 83.1%, OpenAI GPT-5.5-Cyber 85.6%) and 72.1% on CTI-REALM — vendor scores, treated as vendor signals [secondary].
- **Fugu Max + Fugu Ultra v2** (launched 11–12 Sept 2026) [vendor-reported/secondary] https://www.marktechpost.com/2026/09/10/sakana-ai-launches-fugu-max-and-fugu-ultra-v2-for-cheaper-stronger-multi-agent-orchestration/:
  - Fugu Max: wider pool of open-weight + specialized models incl. NVIDIA Nemotron family via Aug 2026 NVIDIA collaboration; **pricing $2/$6 per 1M input/output tokens**, "fixed rates regardless of context length," 1M context window. Claims best overall on 6 benchmarks (Terminal Bench 2.1, GPQA Diamond, AA-LCR, GDP.pdf, AutomationBench, SWEFish) and Pareto-frontier expansion on 7/10 — all self-reported; SWEFish is Sakana's own internal benchmark.
  - Fugu Ultra v2: targets hard multi-step reasoning/SWE; claims 48.3 on Chartography, 74.3 on DeepSWE, best/joint-best on 5 of 8 benchmarks — self-reported. **Pricing $5/$30 per 1M in/out tokens, $0.50/M cached input; $10/$45/$1.00 above 272K context.**
  - Explicitly excludes Fable 5/5.1 and GPT-6-Astra from its pool; 1-line switch on OpenAI-compatible API; hosted API only, no EU/EEA access [secondary].
  - Strategic framing: "supply-chain resilience by design" — routing around vendor lock-in and export controls; Sakana positions Fugu Ultra against Anthropic Fable 5/Mythos after US export controls (12 June 2026) cut those models off in many countries [secondary].
  - Caveats: early third-party testers (e.g. Ethan Mollick) reported a gap between benchmarks and real-world use [secondary].

**Partnerships:** NVIDIA collaboration (Aug 2026): Nemotron models folded into Fugu orchestration pool [vendor-reported/secondary].

### 10.7 Cohere — detailed (from Track C)

**2026 model releases:**
- **Tiny Aya (Feb 2026)** — Cohere Labs open-weight multilingual family: 3.35B parameters, 70+ languages, regional variants (TinyAya-Earth/African, -Fire/South Asian, -Water/APAC+West Asia+Europe, -Global), designed for offline/edge use with no internet connectivity; on Hugging Face, Kaggle, Ollama, Cohere Platform; trained on 64× H100 [secondary].
- **Command A+ (20–21 May 2026)** — flagship enterprise model, the direct Command R/R+ successor [vendor-reported via cohere.com/blog/command-a-plus]:
  - Sparse MoE, 218B total / 25B active per token (128 experts, 8 active + 1 shared); 128K input / 64K output context; text+image inputs, tool use, reasoning traces; 48 languages (all EU official languages).
  - Consolidates prior Command A, Command A Reasoning, Command A Vision, Command A Translate variants.
  - **License: full Apache 2.0 open weights** — first Cohere frontier-class model with unrestricted commercial use; W4A4 quantization runs on 1× B200 or 2× H100. Native source citations emitted at inference.
  - Model ID: `command-a-plus-05-2026`; weights on Hugging Face (CohereLabs). Also available via managed API and Model Vault.
  - Self-reported benchmarks: Terminal-Bench Hard 3%→25%, telecom reasoning 37%→85%; Artificial Analysis Intelligence Index score 37, "front of the open-weight pack" [vendor-reported/secondary].
  - API pricing not located in this pass; trial/production keys free until rate limits, production via private platform [secondary — unverified].
- **Acquired Reliant AI** (biopharma specialist) in May 2026, announced alongside Command A+ [secondary — single source, treat as [unverified] pending confirmation].
- North (enterprise agentic workspace) is the deployment vehicle Command A+ was built for [vendor-reported].

