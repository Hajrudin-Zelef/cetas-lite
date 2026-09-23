---
id: vague2-nerdykings/nerdykings/jev-typesafe-system-one-models
title: "Jev : Comment Fonctionne Ce Nouveau Modèle IA ?"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "chatgpt", "claude", "cost", "fable 5", "funding", "gemini", "inference", "latency", "pricing", "reasoning"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/jev-typesafe-system-one-models.md
source_anchor: ""
source_lines: [1, 64]
sha256: 7216532b2c1630071dc9bbe8d3cac6d7ecc20d341c666dc6558714d6e5a21ee5
---

# Jev : Comment Fonctionne Ce Nouveau Modèle IA ?

## Metadata

- **Source** : https://www.nerdykings.com/blog/jev-typesafe-system-one-models.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

A new model called Jev is generating significant buzz: responses in 70–500 milliseconds, up to 200x faster than some LLMs, and so cheap that its creators don't even charge for output. Many now ask whether an LLM is really needed for every AI task. Jev works nothing like ChatGPT — it isn't built to chat or write text, but to make huge numbers of tiny decisions extremely fast.

The problem Jev addresses: models like GPT, Claude, or Gemini generate responses token by token, which is powerful but overkill for simple classification. With 100,000 emails to route to billing, support, or sales, an LLM is an enormous machine used just to output a category; at millions of decisions, cost and latency become real problems.

Jev comes from TypeSafe AI, a San Francisco startup emerging after two years of stealth. Founded by Diogo Almeida (former OpenAI researcher who worked on RLHF and Instruct-GPT), with Eric Gafni and Sasa Sheng, it raised $40M to build a new model category: **System One Models**, named after Daniel Kahneman's System 1 (fast, intuitive) vs System 2 (slow, deliberate). Big models resemble System 2; Jev aims for the opposite — decide fast, not think long.

Concretely, instead of generating text, Jev is given the possible choices (support, billing, sales, other) and assigns a probability to each (e.g., 94% billing, 3% support, 2% sales, 1% other). Since answers are predefined, it cannot invent a fifth category. TypeSafe offers three decision types: **choice** (pick among options), **score** (rating on a scale), and **nul** (yes/no). Jev can process multiple questions over the same context in parallel (e.g., department, customer anger level, urgency on the same email).

To make probabilities trustworthy, TypeSafe uses **RLCD** (Reinforcement Learning for Calibrated Decisions): Jev must not only answer correctly but estimate how much it can be trusted. If it says 70% across many similar decisions, it should be right ~7 times out of 10 (calibrated probability). This enables routing: above 95% execute directly, 70–95% ask a larger model to verify, below that send to a human — Jev is essentially "a smart if." Code handles the predictable, Jev makes context-dependent small decisions, and big LLMs intervene only when real reasoning or generation is needed.

Performance: responses typically 70–500 ms (median ~100 ms); ~$0.042 per million input tokens with output not billed; TypeSafe claims up to 200x faster and 400x cheaper, though on tasks particularly suited to Jev. An external test (finding hidden issues in a document) had Jev answer in median 0.35 s vs 8.83 s for Fable 5.1 (~25x faster), but Jev found 6 of 7 problems vs 7 of 7 — much faster/cheaper, not necessarily smarter.

Jev isn't a direct competitor to GPT or Claude. It shines when a small decision must be repeated massively: document creation, content moderation, request routing, evaluating LLM responses, or checking agent actions. Developers used it to control agents in Minecraft, Doom, Subway Surfer, and a drone simulator at ~10 decisions/second. Classification without text generation isn't new (BERT predates GPT), but Jev packages it in very fast infrastructure with a standardized API and RLCD-calibrated probabilities. Crucially, "no hallucination" is true only in a narrow sense: Jev can't invent a category (structural hallucinations disappear) but can still choose the wrong one, even confidently. Zero structural hallucination ≠ zero error.

The author concludes Jev deserves the hype but not for all the stated reasons: it didn't invent classification, isn't smarter than GPT/Claude, and "zero hallucination" needs caveats. Its real value is the inverse movement — a very specialized, extremely fast, nearly free model. If it scales, it may herald a generation of small specialized models, each handling one decision type while big LLMs handle complex tasks. Reliability at scale remains unproven beyond TypeSafe's numbers and one independent test.

## Key points

- Jev is a "System One" model built for fast, repeated small decisions, not chat or text generation.
- Created by TypeSafe AI (SF), founded by ex-OpenAI researcher Diogo Almeida; raised $40M.
- Assigns probabilities to predefined choices, preventing structural hallucinations but not errors.
- Three decision types: choice, score, and nul (yes/no); parallel questions on one context.
- RLCD makes confidence calibrated, enabling automatic/verify/human routing.
- 70–500 ms latency (median ~100 ms); ~$0.042/M input tokens, output free; claims up to 200x faster and 400x cheaper.
- External test: ~25x faster than Fable 5.1 but 6/7 vs 7/7 on hidden issues.
- Classification without generation isn't new (BERT); Jev's novelty is speed + API + calibrated probabilities.

## Technical data / figures

| Metric | Value |
|---|---|
| Latency | 70–500 ms (median ~100 ms) |
| Input price | ~$0.042 / million tokens |
| Output price | Not billed (free) |
| Claimed speed gain | Up to 200x |
| Claimed cost gain | Up to 400x |
| External test (hidden issues) | Jev 0.35 s median, 6/7 found |
| Fable 5.1 (same test) | 8.83 s median, 7/7 found |
| Funding | $40M |
| Agent control rate | ~10 decisions/second |

- Decision types: **choice**, **score**, **nul**
- Calibration method: **RLCD** (Reinforcement Learning for Calibrated Decisions)

## Why this source matters for the RAG

This article clearly explains an emerging model category (System One Models) that diverges from the "bigger is better" LLM trend, with concrete architecture, pricing, and calibration concepts. It is valuable for a RAG knowledge base covering AI architectures, inference economics, and hybrid routing between small fast models and large LLMs.

## Source URL

https://www.nerdykings.com/blog/jev-typesafe-system-one-models.html
