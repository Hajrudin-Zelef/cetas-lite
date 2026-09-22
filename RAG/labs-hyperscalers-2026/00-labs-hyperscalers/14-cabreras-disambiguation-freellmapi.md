---
id: labs-hyperscalers-2026/00-labs-hyperscalers/14-cabreras-disambiguation-freellmapi
title: "§14 — \"CABRERAS\" DISAMBIGUATION + FREELLMAPI"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Apple", "Cerebras", "Cohere", "EU", "Fireworks AI", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Sakana", "SpaceX", "Stripe", "xAI"]
dates: ["2026-05-13"]
keywords: ["acquisition", "agentic", "agents", "apache", "astra", "bedrock", "chatgpt", "claude", "cohere", "compute", "copilot", "cyber"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1802, 1868]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: a72bab1a6d1f1592d91956c59fb041f00f1f4f9ef916e7e9ed5a8eaaec8c73da
---

# §14 — "CABRERAS" DISAMBIGUATION + FREELLMAPI

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

**Self-reported scale (repo-reported, unverified independently):** README claims **34 providers** aggregated and **7.4B monthly tokens** routed [secondary — self-reported by the project].

**How it works:** API keys and free-tier quotas belong to the underlying providers; FreeLLMAPI routes/normalizes across them. Users supply their own keys for most providers [secondary].

**Risk flag — Terms of Service:** aggregating/reselling free tiers may violate upstream providers' ToS; several providers prohibit proxying or commercial resale of free-tier access. Anyone routing production traffic through FreeLLMAPI should review each upstream provider's terms [secondary — legal review recommended].

**RAG note:** because it is a router, model availability/pricing "via FreeLLMAPI" is derivative — always cite the underlying provider for model facts. Do not list FreeLLMAPI as a model source.

### Monthly key themes (analytical summaries)

**February — The merger month.** SpaceX absorbs xAI ($1.25T) while Anthropic ($30B @ $380B) and OpenAI (GPT-5.3-Codex vs Claude Opus 4.6, launched minutes apart) open the year's model race. Google answers with Gemini 3.1 Pro (Feb 19); Mistral ships Voxtral Transcribe 2; Amazon plants its $50B flag on OpenAI (announced Feb 27). [research finding]

**March — Money and silicon.** OpenAI closes $122B @ $852B (Mar 31); Ironwood TPU goes GA the same day; Mistral stacks Small 4 + Leanstral + Voxtral TTS + $830M debt; Microsoft drops Phi-4-reasoning-vision; Thinking Machines locks in 1 GW of Vera Rubin with Nvidia. [research finding]

**April — Agents go enterprise.** Claude Cowork (Apr 9), ChatGPT Pro $100 tier, GPT-5.5 (Apr 22), Meta's Muse Spark 1.0 (Apr 8), Mistral Medium 3.5 + Workflows (Apr 28–29), Grok 4.3 (Apr 17/30) — every lab ships an agentic surface. Microsoft–OpenAI restructures (Apr 27); the Colossus Clean Air Act lawsuit lands (Apr 14). [research finding]

**May — Sovereignty surge.** Microsoft's $50B Anthropic cloud pact and Mistral's Airbus/BMW/EDF 5-year deals (May 28) bracket the month; Le Chat becomes Vibe; Cohere drops Command A+ (Apache 2.0) and buys Reliant AI; Cerebras IPOs (May 13); Gemini 3.5 Flash launches at I/O (May 19). [research finding]

**June — Liquidity and superintelligence.** SPCX's $75B IPO (Jun 12) is the largest ever; Microsoft unveils the seven-model MAI family (Jun 2) declaring independence from OpenAI; Sakana launches Marlin + Fugu; EO 14409 puts frontier releases under federal pre-review; Apple previews the Siri rebuild at WWDC. [research finding]

**July — The breach.** OpenAI's own eval models breach Hugging Face (Jul 13–16, disclosed Jul 21) — the window's central safety event; GPT-5.6 goes GA (Jul 9) with Luna's price cut 80% (Jul 30); Thinking Machines open-releases Inkling (Jul 15); Microsoft funds Mistral's EU compute (Jul 21); Grok 4.5 ships amid a training-data controversy; inference labs raise ~$3.3B combined (Together/SambaNova/Fireworks). [research finding]

**August — Consolidation.** Anthropic's $52B Series I @ $1.26T (Aug 20) tops the private markets; Stripe buys OpenRouter (Aug 19); Mistral–HUMAIN (Aug 24) extends sovereign AI to the Gulf; Grok 4.6 ships and spreads to Copilot/Bedrock/Foundry; EU AI Act GPAI enforcement goes live (Aug 2); 100+ companies sign the cyber-defense call (Aug 27). [research finding]

**September — Endgame positioning.** GPT-6 Astra's "Critical" cyber classification (Sep 3) and NVIDIA's Hugging Face acquisition (signed Sep 3) collide on the same day; Mistral's €3B Series D (Sep 8) is Europe's biggest private tech round; Cohere–Aleph Alpha signs (Sep 16); the AI-slowdown class action lands (Sep 18); Grok 4.7 closes the window (Sep 21) as the collection ends (Sep 22). [research finding]

### 14.6 FreeLLMAPI — usage notes (from Track D)

- **Setup pattern:** users supply their own API keys for underlying providers; FreeLLMAPI normalizes them behind one OpenAI-compatible endpoint [secondary — repo docs].
- **Use cases reported:** prototyping across many models without managing N SDKs; failover across free tiers [secondary].
- **Limits:** free-tier quotas are per-provider and often rate-limited aggressively; not a production-serving story [secondary].
- **Risk (repeated from §14.5):** proxying/reselling free-tier access may violate upstream ToS — review each provider's terms before routing anything beyond personal experimentation [secondary].
- **RAG rule:** never cite "via FreeLLMAPI" as a model source — always resolve to the underlying provider.

### 14.7 Entity-resolution note (for RAG retrieval)

Four easily confused names appear in this file — keep them distinct in any index or embedding:

| Mention | Resolves to | Type |
|---|---|---|
| "Cabreras" | Almost certainly a misspelling of **Cerebras** | No entity found |
| Cerebras | Wafer-scale AI chip company (IPO May 13, 2026) | Silicon vendor |
| Groq | LPU inference company (Jonathan Ross) | Inference provider |
| Grok | xAI's model family (4.20 → 4.7) | Model |

A query for "Grok pricing" must never return GroqCloud's LPU price list, and "Cabreras funding" must resolve to Cerebras or return nothing — not to a hallucinated entity. [research finding]


