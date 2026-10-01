---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-40
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Baseten", "Cerebras", "Cohere", "CoreWeave", "EU", "Fireworks AI", "Google", "Groq", "Meta", "Microsoft", "Mistral", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "Sakana", "Stripe", "Together AI", "United States", "xAI"]
dates: ["2026-05", "2026-05-13", "2026-07-01", "2026-07-08", "2026-07-16"]
keywords: ["hyperscaler", "acquisition", "agent", "agents", "aws", "backlog", "bedrock", "capex", "claude", "cohere", "copilot", "foundry"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1720, 1820]
sha256: 82a50aa3eba3ecdfa8379ded1be494deed2891fdb692c160cfbf6d39009bdf54
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

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

**Routing-layer trend (2026):** the value migrates from single-model APIs to routers/orchestrators — OpenRouter (acquired by Stripe), Sakana Fugu (multi-agent orchestration), Microsoft Foundry, Mistral Workflows, AWS Bedrock AgentCore. Inference is becoming a routing problem, not just a serving problem. [research finding]

## §14 — "CABRERAS" DISAMBIGUATION + FREELLMAPI

> **2026 at a glance — this section:** "Cabreras" is not an AI entity (likely a Cerebras misspelling — research finding); FreeLLMAPI is an open-source router over 34 providers' free tiers (repo-reported 7.4B monthly tokens), not an inference provider — resolve every model claim to its underlying provider.

### 14.1 "Cabreras" — not found as an AI company or product
- **Research finding: "Cabreras" was not found as an AI company, model, or product** in any source consulted during Feb–Sep 2026 collection. [research finding]
- **Treated as a likely misspelling of Cerebras** (the wafer-scale AI chip company, §13.1). All "Cabreras" references in source material should be read as Cerebras unless new evidence emerges. [research finding]
- This disambiguation is preserved here so downstream RAG retrieval does not invent a "Cabreras" entity. [research note]

### 14.2 FreeLLMAPI — what it actually is
- **FreeLLMAPI is NOT a first-party inference provider.** It is an **open-source, self-hosted router/aggregator** that pools the **free tiers of third-party providers** behind a single OpenAI-compatible API surface. [secondary: Track D]
- **Terms-of-Service risk**: routing traffic through providers' free tiers via an aggregator may violate those providers' ToS; users should verify each upstream provider's terms. This risk is retained explicitly. [research note]
- Use cases: development, prototyping, low-volume experimentation — not production workloads. [secondary]

---
### 14.5 FreeLLMAPI — detailed (from Track D)

**What it is:** FreeLLMAPI is an **open-source router/aggregator**, not a first-party inference provider. It presents a unified OpenAI-compatible API over a pool of providers' free tiers / free models [secondary — repo documentation].

