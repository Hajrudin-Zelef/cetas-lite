---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/grok-4-1-released-2025-11-17-details-beyond-base-section
title: "Grok 4.1 (released 2025-11-17; details beyond base section)"
domain: xai-and-grok
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "EU", "Google", "Microsoft", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: ["2025-11-17", "2026-02-01", "2026-02-17", "2026-04", "2026-06", "2026-06-27", "2026-07", "2026-07-08", "2026-08", "2026-08-12", "2026-08-18"]
keywords: ["grok", "grok 4", "acquisition", "agent", "agentic", "agents", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt", "claude"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5223, 5295]
section: "§11. xAI and Grok"
sha256: 4541f7725dda460b2201a95b41d8c34ad6be56b0d69257a38edd8e1ea70168f4
---

# Grok 4.1 (released 2025-11-17; details beyond base section)

### Grok 4.1 (released 2025-11-17; details beyond base section)
- Grok 4.1 was released 2025-11-17 in two variants: Grok 4.1 and Grok 4.1 Thinking. [SECONDARY, S4][SECONDARY, S1]
- Headline changes: emotional-intelligence upgrade and hallucination reduction. [SECONDARY, S4][SECONDARY, S1]
- Vendor-reported hallucination rate fell from 12.09% (Grok 4 Fast) to 4.22% (non-reasoning mode) — a ~65% reduction; FActScore error rate fell from 9.89% to 2.97%. [VENDOR, S1][SECONDARY, S48][SECONDARY, S49]
- Vendor-reported LMArena score: 1483 Elo for Grok 4.1 Thinking (#1 on Text Arena, 31 points ahead of Gemini 2.5 Pro); the standard Grok 4.1 took #2. [VENDOR, S1][SECONDARY, S48][SECONDARY, S51]
- xAI claimed Grok 4.1 outperformed Gemini 2.5 Pro and Claude 4.5 Sonnet on certain internal metrics; coverage also noted it beat ChatGPT on EQ-Bench and Creative Writing v3 (#2 behind GPT 5.1). [VENDOR, S4][SECONDARY, S48][SECONDARY, S51]
- Both variants were free to all users at launch; paid users retained higher usage ceilings. [SECONDARY, S4][SECONDARY, S72][SECONDARY, S73]

### Grok 4.20 Beta (released 2026-02-17)
- Grok 4.20 public beta launched 2026-02-17, announced by Elon Musk posting that it was live rather than through a formal launch event. [SECONDARY, S2 — single source]
- Its headline architecture feature is "rapid learning": weekly capability improvements driven by public-use feedback, with release notes published alongside every update; the model is described as a council of 4 specialized agent replicas (named in demos as Grok + Harper + Benjamin + Lucas) running think → debate → consensus with fact-checking loops. [SECONDARY, S2][SECONDARY, S50]
- Users had to manually select "Grok 4.2" in the model menu to access the 4.20 beta. [SECONDARY, S2 — single source]
- Headline capabilities: medical document analysis via photo upload and improved engineering reasoning. [SECONDARY, S2 — single source]
- Musk's stated goal: Grok 4.20 would become "an order of magnitude smarter and faster" than Grok 4 by the end of the beta. [SECONDARY, S2 — single source]
- It was described as a 4-agent system in model timelines. [SECONDARY, S3 — single source]
- API behavior note: for `grok-4.20-multi-agent`, the `reasoning_effort` parameter controls collaborator count rather than reasoning depth. [SECONDARY, S29 — single source]
- Grok 4.20's public status (beta vs GA) was still listed as beta in timelines as of mid-2026. [SECONDARY, S3 — single source]

### Grok 4.2 — unresolved public status (documented, not asserted)
- One secondary timeline lists Grok 4.2 as "expected Nov–Dec 2025" as a polished 4.x release with Grok Imagine video. [SECONDARY, S1 — single source]
- A second timeline omits Grok 4.2 entirely, jumping from 4.1 to 4.20 Beta. [SECONDARY, S3]
- Whether Grok 4.2 ever shipped publicly is unresolved in the sources found; the expansion does not assert a release date. [UNVERIFIED]
- The "no Grok 4.4" numbering gap is consistently supported: no source documents a 4.4 release. [SECONDARY, S3][SECONDARY, S1]

### Grok 4.3 (~April 2026; details beyond base section)
- Grok 4.3 shipped around April 2026 as a lower-cost long-context option with up to 1M tokens. [SECONDARY, S3][SECONDARY, S12]
- It supports native video input. [SECONDARY, S12 — single source]
- It can generate PDFs, spreadsheets, and slide decks directly. [SECONDARY, S12 — single source]
- Artificial Analysis Intelligence Index: 38 at release (methodology v4.1, per the reporting source). [SECONDARY, S12 — single source]
- Reasoning-effort levels on the API: `off` / `minimal` / `low` / `medium` / `high` — the only Grok 4.x model documented with a true `off` setting. [SECONDARY, S29 — single source]

### Grok 4.5 (released 2026-07-08; details beyond base section)
- Grok 4.5 was released to the public 2026-07-08, eleven days after a private beta at SpaceX and Tesla (2026-06-27). [SECONDARY, S8][SECONDARY, S10]
- It is built on V9, xAI's ninth-generation architecture, reported at 1.5 trillion parameters. [SECONDARY, S8][SECONDARY, S10][SECONDARY, S11]
- It was jointly developed with Cursor (Anysphere): xAI folded real Cursor developer session data — debugging traces, multi-file diffs, user corrections — into training. [SECONDARY, S8][SECONDARY, S10]
- SpaceX announced the acquisition of Cursor's parent Anysphere in mid-June 2026 at a $60B valuation, expected to close Q3 2026; the joint model shipped weeks before close. [SECONDARY, S10 — single source]
- Pricing: $2 per million input tokens, $6 per million output tokens. [SECONDARY, S8][SECONDARY, S10][SECONDARY, S11]
- A faster premium variant runs at $4/$18. [SECONDARY, S10 — single source]
- Context window: 500K tokens. [SECONDARY, S9][SECONDARY, S12]
- Reported inference speed: roughly 80 tokens/second. [SECONDARY, S11 — single source]
- Vendor benchmarks vs Claude Opus 4.8: DeepSWE 1.0 win, Terminal-Bench 2.1 win (83.3%), DeepSWE 1.1 loss by 6 points, SWE-Bench Pro loss by 4.5 points (64.7%). [VENDOR, S8]
- Vendor-reported token efficiency: ~4.2× fewer tokens than Opus 4.8 on SWE-Bench Pro (15,954 vs 67,020 average output tokens). [VENDOR, S11][SECONDARY, S8]
- SWE Marathon pass@1: 29.0% vs Opus 4.8 26% and Fable 5 24% (vendor). [VENDOR, S10]
- On Cursor's CursorBench platform it handled automated coding tasks at $1.51 per task. [SECONDARY, S11 — single source]
- Available at launch in Grok Build, in Cursor on all plans, and through the SpaceXAI console; free for a limited period in Grok Build and Cursor. [SECONDARY, S8][SECONDARY, S10]
- Cursor stated Grok 4.5 and its own Composer 2.5 are different weight classes and both remain available. [SECONDARY, S8 — single source]
- EU availability was expected mid-July 2026, not at launch. [SECONDARY, S9 — single source]
- Supports configurable reasoning effort. [SECONDARY, S9 — single source]
- Artificial Analysis Intelligence Index: 54 (v4.1 methodology, high reasoning), #4 overall at release — up from Grok 4.3's 38. AWS's Bedrock launch coverage cites xAI's launch figure of 61 for Grok 4.6 High (v4.1.1 per AWS) — version-level discrepancy, see ledger. [SECONDARY, S12][SECONDARY, S77][SECONDARY, S78]
- It was the first model released after the Cursor acquisition announcement and the SpaceXAI rebrand. [SECONDARY, S12 — single source]
- HN commenters flagged the benchmark figures as vendor-reported and noted the $2/$6 price against $5/$30 (GPT-5.5) and $5/$25 (Opus 4.8). [COMMUNITY, S8 — single source]

### Grok 4.6 (released 2026-08-12; details beyond base section)
- Grok 4.6 was released 2026-08-12 as a SpaceXAI reasoning model for coding, long-running agents, interactive/visual projects, and knowledge work. [SECONDARY, S25][SECONDARY, S26]
- API identifier: `grok-4.6`; text and image input, text output. [SECONDARY, S25][SECONDARY, S26]
- Context window: 500,000 tokens; prompts at or above 200,000 tokens use long-context pricing. [SECONDARY, S25][SECONDARY, S26]
- Four configurable reasoning-effort levels: `low`, `medium`, `high` (default), `xhigh`; reasoning cannot be disabled. [SECONDARY, S25][SECONDARY, S26]
- The `xhigh` level requires xai-sdk 1.18.0 or later (older SDKs reject the value client-side). [SECONDARY, S27 — single source]
- Knowledge cutoff: 2026-02-01. [SECONDARY, S26 — single source]
- Pricing: $2/M input, $6/M output; cached input $0.50/M. [VENDOR, S30 — single source]
- Vendor evaluations: CursorBench v3.2 69.9%; DeepSWE v1.1 65.9%; FrontierCode v1.1 Extended 61.3%; APEX-Agents 57.5%. [VENDOR, S25 — single source]
- Training: a longer supplemental training run than Grok 4.5 received — not a larger base model — using curated model-generated data, high-quality engineering data, and an improved optimizer and training recipe. [SECONDARY, S31 — single source]
- Grok 4.5 was used to regenerate supervised fine-tuning trajectories across reasoning-effort levels, agent harnesses, and domains (STEM, software engineering, knowledge work), with problematic traces filtered by model-based checks. [SECONDARY, S31 — single source]
- Reinforcement learning followed in agentic environments covering knowledge work, general coding, web development, computer-aided design, and kernel optimization. [SECONDARY, S25][SECONDARY, S31]
- Reported behavioral effect: more self-testing and verification on longer trajectories (vendor observation from internal testing, not independently measured). [VENDOR, S31 — single source]
- No parameter count was published for Grok 4.6. [SECONDARY, S31 — single source]
- The official SpaceXAI models overview recommends Grok 4.6 for code and chat and calls it the most intelligent and fastest model xAI has built. [SECONDARY, S29 — single source]
- Generally available on Amazon Bedrock (Bedrock model card dated 2026-08-18); xAI's announcement frames Bedrock as a first-class channel. Bedrock coverage describes Grok 4.6 as an MoE design of roughly 1.5T parameters with a 500K context window and four selectable reasoning levels. [VENDOR, S30][SECONDARY, S77][SECONDARY, S78]
- Later in August 2026 it was extended to GitHub Copilot, the Gemini Enterprise Agent Platform, and Microsoft Foundry. [SECONDARY, S37 — single source]
- Live in Cursor, Grok Build 1.0, and Grok Bot agents at launch; distributed through OpenRouter, Vercel, and Cloudflare. [SECONDARY, S26 — single source]
- xAI documentation recommends the `prompt_cache_key` parameter to keep multi-turn request routing stable and cache hits reliable. [SECONDARY, S26 — single source]
- Third-party provider Pi maps its `minimal` thinking level onto xAI's `low` for Grok 4.6. [SECONDARY, S28 — single source]

