---
id: labs-hyperscalers-2026/00-labs-hyperscalers/13-inference-hyperscalers-cerebras-sambanova-together-ai-fir
title: "§13 — INFERENCE HYPERSCALERS: CEREBRAS, SAMBANOVA, TOGETHER AI, FIREWORKS AI, NEBIUS, COREWEAVE"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Baseten", "Cerebras", "Cohere", "CoreWeave", "EU", "Fireworks AI", "Google", "Groq", "Meta", "Microsoft", "Mistral", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "Sakana", "Stripe", "Together AI", "United States", "xAI"]
dates: ["2026-05", "2026-05-13", "2026-07-01", "2026-07-08", "2026-07-16"]
keywords: ["hyperscaler", "inference", "acquisition", "agent", "agents", "aws", "backlog", "bedrock", "capex", "claude", "cohere", "copilot"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1692, 1799]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 0ec10f9fb8bb5af887a9ee8a6e0ff90cc94b6b60a552cdc1bb9f8fb6f46cb3b9
---

# §13 — INFERENCE HYPERSCALERS: CEREBRAS, SAMBANOVA, TOGETHER AI, FIREWORKS AI, NEBIUS, COREWEAVE

## §13 — INFERENCE HYPERSCALERS: CEREBRAS, SAMBANOVA, TOGETHER AI, FIREWORKS AI, NEBIUS, COREWEAVE

> **2026 at a glance — inference hyperscalers:** Cerebras IPO'd (May 13); SambaNova ($1B), Together AI ($800M) and Fireworks ($1.505B) raised in a nine-day July window; Nebius built out with reported Meta orders; CoreWeave levered its backlog into September converts; Together/Fireworks/Baseten/Modal/Databricks all served Inkling at launch.

### 13.1 Cerebras
- **Cerebras IPO** — the wafer-scale AI chip company's **initial public offering** took place during the window (Track D) — landmark liquidity event for the AI-silicon sector. [secondary]
- **WSE (Wafer-Scale Engine)** lineage: CS-3 systems; wafer-scale architecture as the differentiator vs. GPU clusters. [secondary]
- **Cerebras Inference** — cloud inference API on wafer-scale hardware; 2026 expansion. [secondary]
- Post-IPO partnerships and enterprise traction (Track D).

### 13.2 SambaNova
- **Series F** financing round during the window (Track D) — continued capitalization of the reconfigurable-dataflow (RDU) AI platform. [secondary]
- **SambaNova Cloud** inference offerings; sovereign-AI and enterprise positioning. [secondary]

### 13.3 Together AI
- **Series C** financing during the window (Track D). [secondary]
- Inference + fine-tuning API platform; open-model hosting leadership. [secondary]
- **Together AI is a separate company from Cerebras and Groq** — independent inference API provider. [research note]

### 13.4 Fireworks AI
- **Series D** financing during the window (Track D). [secondary]
- Fast inference API for open models; function-calling / agent infrastructure. [secondary]

### 13.5 Nebius
- **Capital stack** activity in 2026 (Track D): continued financing of the Nebius AI-cloud buildout (post-Volozh/Yandex split). [secondary]
- **Meta orders**: Nebius reported **Meta as a customer / order flow** in 2026 (see §4.4). [secondary]
- Positioned as the "neocloud" challenger alongside CoreWeave. [secondary]

### 13.6 CoreWeave
- **Debt financing and backlog** growth through 2026 (Track D): CoreWeave continued raising debt against contracted GPU revenue; reported backlog expansion. [secondary]
- Post-IPO (2025) execution: data-center buildout for hyperscaler and lab customers. [secondary]
- Customer concentration dynamics (Microsoft, OpenAI, Meta) watched through 2026. [secondary]

### 13.7 Inference price landscape (2026 snapshot — dated, [secondary])
- Inference pricing across Groq, Cerebras, Together, Fireworks, and hyperscaler APIs compressed through 2026; open-model inference (e.g., Llama/Muse/Qwen-class) fell to fractions of a cent per 1M tokens on some providers (Track D — specific dated rates in §16 pricing tables).
- **Long-context premiums** and **reasoning-token billing** became standard differentiators (cf. OpenAI's >272K 2x/1.5x billing, §2.3). [secondary]

---

### 13.7 Inference providers — detailed (from Track D)

**Cerebras — IPO May 13, 2026:**
- Went public **May 13, 2026**; reported pricing **$185/share**, raising **~$5.55B** — figures appear in secondary roundups but conflict across sources (some cite different raise/share counts); **keep flagged as conflicting** [secondary, conflicting].
- Technology: wafer-scale WSE-3 chips; inference service (Cerebras Inference) known for very high tokens/sec on Llama-class models [secondary].
- Post-IPO: continued DC expansion; inference API pricing aggressive vs GPU clouds [secondary — thin 2026 sourcing, follow-up recommended].

**SambaNova — Series F July 8, 2026:**
- Reported **$1B raise at $11B post-money** [secondary — single-source lineage, treat as reported not confirmed].
- Full-stack (chips + cloud + models); Samba-1 / Composition of Experts platform [secondary].

**Together AI — Series C July 1, 2026:**
- Reported **$800M at $8.3B valuation** [secondary]. One report described an April tranche structure — **conflicting tranche reporting flagged**; do not present a single clean figure without the caveat [secondary, conflicting].
- Serves open models incl. Thinking Machines' Inkling at launch (Jul 15, 2026) [secondary].

**Fireworks AI — Series D July 16, 2026:**
- Reported **$1.505B at $17.5B valuation** [secondary — single-source lineage].
- Serves Inkling at launch; function-calling / compound-AI focus [secondary].

**Nebius (ex-Yandex N.V.):**
- 2026: continued GPU-cloud buildout across US/EU; reported large-scale GPU supply agreements; **reported Meta orders** for inference capacity [secondary — thin sourcing, flag unverified].
- Public company (NASDAQ: NBIS); capex-heavy expansion financed via equity/debt [secondary].

**CoreWeave:**
- Backlog/debt: large contracted backlog against debt-financed GPU fleet; **Sept 17, 2026: convertible notes issuance** reported [secondary — terms thin, flag for follow-up].
- Public company (NASDAQ: CRWV, IPO Mar 2025); key OpenAI/Microsoft-adjacent capacity supplier [secondary].

**Baseten / Modal / Databricks:** served Thinking Machines' Inkling at launch (Jul 15, 2026) alongside Together/Fireworks [secondary].

### 13.8 Pricing snapshot — third-party inference (2026, verify live)

> Third-party inference pricing moves frequently. Figures below are secondary snapshots; verify against provider pricing pages before citing.

| Provider | Representative price point (2026) | Notes |
|---|---|---|
| Groq (GroqCloud) | ~$0.35/$0.79 per 1M in/out (Llama 3.3 70B class) | LPU; free tier (rate-limited) |
| Cerebras Inference | Aggressive per-token pricing on Llama-class models | Wafer-scale; verify live |
| Together AI | Open-model serverless; varies by model | Served Inkling at launch |
| Fireworks AI | Function-calling optimized; varies | Served Inkling at launch |
| Nebius | GPU-cloud + serverless | Verify live |
| CoreWeave | Reserved GPU capacity | Verify live |

### 13.9 Key sources — inference providers
- Cerebras IPO coverage: secondary roundups, May 2026 (figures conflict — verify)
- Artificial Analysis provider leaderboards: https://artificialanalysis.ai

### 13.10 Provider model-catalog matrix (Sep 2026 snapshot)

> Which flagship models each inference route serves. "1P" = provider's own models; catalogs change frequently — verify live.

| Provider | Serves (representative) | Notes |
|---|---|---|
| GroqCloud (LPU) | Llama 3.3 70B-class, Qwen, Mistral, Whisper large-v3-turbo | Open-weight catalog; deterministic latency |
| Cerebras Inference | Llama-class open models | Wafer-scale throughput pitch |
| Together AI | Broad open catalog incl. **Inkling** (from Jul 15, 2026) | Serverless GPU |
| Fireworks AI | Broad open catalog incl. **Inkling**; function-calling optimized | Compound-AI focus |
| Baseten / Modal / Databricks | **Inkling** at launch; general serverless | Enterprise serverless |
| NVIDIA NIM | Nemotron 3 family; 1P containers | Self-host / DGX Cloud / Build credits |
| OpenRouter (→ Stripe) | Multi-provider routing incl. **MAI models** (from Jun 2, 2026) | Aggregator; acquisition announced Aug 19 |
| Amazon Bedrock | **OpenAI GPT-5.5, Codex, agents** (from Apr 28, 2026); Nova (EOL); Anthropic Claude | AWS-managed |
| Microsoft Foundry | **MAI family**; OpenAI models; **Grok 4.6+**; **Mistral Medium 3.5 + OCR 4** (from Jul 21, 2026) | Multi-model router |
| Google Vertex / AI Studio | Gemini 3.x family | 1P |
| Gemini Enterprise Agent Platform | Gemini + **Grok 4.6+** (from Aug 2026) | Enterprise agents |
| GitHub Copilot | **MAI-Code-1-Flash** (from Jun 2); **Grok 4.6+** | Coding |
| xAI API | Grok 4.x | 1P |
| Mistral La Plateforme | Mistral family | 1P (+ Azure Local route) |
| Cohere Platform | Command A+; Tiny Aya | 1P |
| Sakana API | Fugu orchestration (routes to Claude/GPT/Gemini/Nemotron pools) | Meta-router |

