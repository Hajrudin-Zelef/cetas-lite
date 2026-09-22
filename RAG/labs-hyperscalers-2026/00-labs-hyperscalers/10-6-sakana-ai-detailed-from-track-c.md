---
id: labs-hyperscalers-2026/00-labs-hyperscalers/10-6-sakana-ai-detailed-from-track-c
title: "10.6 Sakana AI — detailed (from Track C)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "Anthropic", "Cohere", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "Nvidia", "OpenAI", "Oracle", "Sakana", "United States"]
dates: ["2026-03", "2026-04", "2026-05", "2026-06", "2026-07"]
keywords: ["sakana", "agent", "agentic", "amd", "apache", "arr", "astra", "benchmark", "benchmarks", "claude", "cohere", "context window"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1463, 1508]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 8633033093bb7e4acb7ac7713e7c7edbfd83a103f8b43f6dbac2adae49d9c5b3
---

# 10.6 Sakana AI — detailed (from Track C)

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

**Enterprise products / deployments 2026:**
- **North**: agentic AI workspace; sector deployments: **North for Banking with Royal Bank of Canada** [secondary]; **S&P Global** (June 2026) — S&P financial data embedded in North for citation-backed agentic workflows [secondary]; **Saab AB** (March 2026) — Cohere tech supporting the GlobalEye surveillance aircraft programme [secondary]; **Hanwha Ocean** (Jan 2026) — generative AI for ship design/procurement [secondary].
- Named partners: Fujitsu (Japanese-language models), LG CNS (South Korea), Dell (on-prem), AMD (hardware optimization), Oracle & Microsoft clouds [secondary].

**Funding / valuation:** $500M round (Aug 2025, context) — after which President & COO Martin Kon stepped back from day-to-day [secondary]. Merger deal: first disclosed **April 2026** at ~$20B (€17.34B) combined valuation; **definitive agreement signed 16 Sept 2026**; Schwarz Group investing €500M; **Series E of $2–3B** reportedly in progress [independent — Reuters; secondary]. Combined entity ARR run-rate: Cohere reported $240M ARR "this year" per Reuters [independent]; IPO (mid-2027) base-case estimate ~$17B in one analyst note [secondary/unverified]. No separate 2026 funding round found outside the merger/Series E context.

**People:** Sara Hooker (Cohere Labs head) exited Aug 2025; **Joëlle Pineau** (ex-Meta VP AI Research) joined as Chief AI Officer Aug 2025; **Phil Blunsom** (ex-DeepMind) promoted to CTO mid-2025; **François Chadwick** (ex-KPMG/ex-Uber) joined as CFO; Ilhan Scheer (Aleph Alpha co-chief) to become COO of combined entity post-merger [secondary/independent].

