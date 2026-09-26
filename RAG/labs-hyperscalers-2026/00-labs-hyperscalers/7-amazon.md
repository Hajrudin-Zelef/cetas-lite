---
id: labs-hyperscalers-2026/00-labs-hyperscalers/7-amazon
title: "§7 — AMAZON"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "EU", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "United States"]
dates: ["2025-03", "2025-12", "2026-04", "2026-04-28", "2026-06", "2026-06-01", "2026-07-28", "2026-07-31", "2026-09", "2026-09-14", "2026-09-30"]
keywords: ["acquisition", "agent", "agentic", "agents", "agi", "aws", "bedrock", "benchmark", "capex", "claude", "compute", "copilot"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [975, 1022]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 5f55af13fa6ec3e01dd246bf8f34082d57b6559b0e98aa19b453fafd1fa13588
---

# §7 — AMAZON

## §7 — AMAZON

> **2026 at a glance — Amazon:** $50B into OpenAI (announced Feb 27, completed Jul 31) — largest single corporate AI check ever, with AWS as exclusive third-party cloud for OpenAI Frontier and 2 GW of Trainium committed; Nova line wound down to KTLO (EOL Sep 14–30) as the Frontier Model Research division under Pieter Abbeel takes over; Alexa+ GA (Feb 4).

### 7.1 Nova consolidation
- **Amazon Nova** model family — 2026 consolidation: Nova became AWS's flagship in-house model line (text, multimodal, image/video generation) competing with Bedrock-hosted third-party models. [secondary: Track B]
- Nova positioning: cost-optimized frontier for AWS enterprise customers. [secondary]

### 7.2 The $50B OpenAI investment — announced Feb 27, 2026; completed July 31, 2026 (full record in §7.6)
- Amazon anchored OpenAI's $122B round with **$50B** — **$15B funded at close; $35B contingent on IPO or AGI by end-2028** (announced Feb 27, 2026; completed July 31, 2026 per SEC filing). Largest single corporate AI investment on record. [secondary] https://tech-insider.org/openai-122-billion-funding-round-852-billion-valuation-2026/ ; https://abhs.in/blog/openai-122-billion-852-billion-valuation-amazon-agi-clause-ipo-2026
- The contingent structure was widely read as establishing a firm IPO timeline for OpenAI. [secondary]

### 7.3 Bedrock & cloud
- **OpenAI models on Amazon Bedrock from 28 April 2026** (GPT-5.5, Codex, agents) — landmark, given the Microsoft exclusivity era had just ended. [secondary] https://www.vaasblock.com/research/microsoft-openai-exclusivity-end-copilot-moat-aws-bedrock-2026/
- **Project Rainier**: reported **$38B, 7-year AWS–OpenAI cloud deal** (hundreds of thousands of GPUs, fully deployed by end-2026); separate reporting claims a **$100B AWS deal** tied to the Amazon investment — [unverified]. [secondary] https://www.coinlive.com/news/openai-raises-122-billion-in-record-breaking-funding-round-at-852 (tweet citation)
- Bedrock continued hosting Anthropic (Claude), Meta, Mistral, and Amazon Nova models. [secondary]

### 7.4 Silicon & capex
- **Trainium3** — Amazon's third-gen AI training chip, ramped in 2026; positioned against Nvidia GPUs for internal + AWS workloads. [secondary: Track B]
- **~$220B capex** — Amazon's guided 2026 capital expenditure, the highest among hyperscalers, driven by AI infrastructure. [secondary: Track B]
- **Alexa+** — the LLM-rebuilt Alexa (launched 2025) continued rollout through 2026 with Nova-class models. [secondary]

### 7.5 Partnerships & regulatory
- **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — AWS signatory. [secondary]
- EU AI Act / DSA compliance posture for AWS AI services (Track B).

---
### 7.6 Nova model family — detailed (from Track B)

**Nova 2 generation (launched Dec 2, 2025 at re:Invent — context for 2026 status):**
- **Nova 2 Lite** (GA Dec 2, 2025): fast, cost-effective reasoning model; 1M-token context, 64K max output; extended thinking with 3-level intensity control; multimodal (text/image/video); prompt caching. Model IDs: `amazon.nova-2-lite-v1:0` (regional + global variants) [official docs] https://docs.aws.amazon.com/nova/latest/nova2-userguide/whats-new.html.
- **Nova 2 Pro** (preview at launch): most intelligent reasoning model; text/image/video/speech → text; positioned for agentic coding, long-range planning; teacher for knowledge distillation. AWS CEO Matt Garman claimed better absolute benchmark results vs. **GPT-5.1, Google Gemini 3 Pro, Claude Sonnet 4.5** [vendor-reported][secondary]. **Pricing: $1.25/M input tokens, $10/M output tokens** [secondary] https://www.crn.com/news/ai/2025/aws-nova-2-ai-models-launched-at-reinvent-2025-as-ceo-touts-new-innovation.
- **Nova 2 Sonic** (GA): second-generation speech-to-speech; unified text+speech understanding/generation; expanded multilingual support, expressive voices, 1M-token context; seamless voice/text switching [official docs]. Independent evaluation (Artificial Analysis, via Loka/AWS case study June 2026): **87.0–87.1%** on Big Bench Audio (1,000 complex audio questions) — 2nd behind Gemini 2.5 Flash Native Audio Thinking, ahead of GPT Realtime (83.0). **Time-to-first-audio: 1.39s** [independent][secondary].
- **Nova 2 Omni** (preview Dec 2025): unified multimodal model — text, image, video, speech input; text and image output. Initially accessible to Nova Forge customers [official docs].
- **Nova Multimodal Embeddings** (Oct 2025): text, documents, images, video, audio → unified embedding space for RAG/semantic search [official docs].
- **Nova Forge** (announced Dec 2025): service to build custom models from Nova training checkpoints ("open training"); includes company data during training (not just fine-tuning), RL gyms, distillation. Early adopters: **Reddit** (built a moderation model reflecting platform dynamics, per CTO Chris Slowe), Booking.com, Hertz [secondary].
- **Nova Act** (March 2025 agent/SDK → **GA as AWS service Dec 2025**): browser and UI automation; AWS claimed **90% reliability** in early customer workflows [vendor-reported][secondary].

**September 2026: Nova portfolio consolidation (Business Insider, July 28, 2026):** Amazon is winding down **Nova Premier, Nova 2 Omni, Nova Reel, and Nova Canvas** to a "KTLO" (keep the lights on) mode — maintenance only, no new features — and moving engineers/compute to a new **Frontier Model Research (FMR)** division led by **Pieter Abbeel** (joined Amazon via the 2024 Covariant acquisition). A new flagship model, possibly retaining the Nova brand, is expected to debut at **re:Invent late 2026** (per Reuters + BI) [secondary][unverified details]. **Official lifecycle dates (AWS Bedrock lifecycle):** Nova Premier **EOL September 14, 2026**; Nova Canvas and two Nova Reel versions **EOL September 30, 2026**. New customers cannot access them; existing customers may lose access after 15 days of inactivity; no automatic migration [secondary, citing official schedule]. Context: former AI chief **Rohit Prasad** departed end of 2025; **Peter DeSantis** took over the consolidated AI group in December 2025; staff cuts in the general AI division; AGI Lab research team reportedly shut down [secondary] https://www.eweek.com/news/amazon-nova-ai-overhaul/.

### 7.7 Bedrock updates (2026) — detailed

**OpenAI models on Bedrock:**
- **Preview began April 28, 2026** (day after Microsoft-OpenAI restructuring); **generally available June 1, 2026**: **GPT-5.5** (most advanced), **GPT-5.4**, and **Codex** coding agent on Bedrock, in US regions (GPT-5.5 in one US region at launch, GPT-5.4 in two) and GovCloud [official] https://aws.amazon.com/blogs/machine-learning/openai-models-and-codex-on-amazon-bedrock-are-now-generally-available/.
- Pricing: GPT-5.5/5.4 at **the same per-token rates as OpenAI direct**, no additional AWS fees; Codex pay-per-token; all inference through Bedrock infrastructure; usage counts toward existing AWS commitments [official].
- Access via **Responses API**; IAM, VPC/PrivateLink isolation, KMS encryption, CloudTrail audit; prompts/responses not used for training [official].
- **Bedrock Managed Agents** powered by OpenAI entered **limited preview** (April 28, 2026) [secondary].

