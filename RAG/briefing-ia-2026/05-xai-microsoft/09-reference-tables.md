---
id: briefing-ia-2026/05-xai-microsoft/09-reference-tables
title: "Reference: glossary, Grok and MAI comparison tables"
domain: xai-microsoft
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["grok", "mai", "agent", "agentic", "agents", "aws", "bedrock", "benchmark", "context window", "copilot", "distillation", "distribution"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s05-38"
source_lines: [6680, 6752]
sha256: c1e5c57e7dc3741242caa89d7dc5c0ff5b5520109b69aba6afbd3b4968070cb7
---

# Reference: glossary, Grok and MAI comparison tables

<a id="s05-38"></a>
### Glossary of the dossier's technical terms

Post-training: all techniques applied after a model's pre-training — supervised fine-tuning, reinforcement learning, behavior optimization.
The Grok 4.5, 4.6, and 4.7 upgrades are post-trainings on the V9 base, not new pre-trainings.
Foundation model: a model pre-trained from zero on immense corpora, serving as a base for adaptations.
Grok 4.8, with its announced 2.5T parameters, is xAI's candidate in this category.
MoE (mixture of experts): architecture where only certain "experts" (sub-networks) are activated per token — MAI-Thinking-1 activates 35B parameters per token.
Context window: the amount of text the model can take into account at once — 500K tokens for Grok 4.5, 256K for MAI-Thinking-1.
SWE-Bench Pro: benchmark for solving real software-engineering tasks — MAI-Thinking-1 matches Opus 4.6 there.
FLEURS: multilingual speech-recognition benchmark — MAI-Transcribe-2 is #1 across 60 languages.
WER (word error rate): word error rate in transcription — 5.2% for MAI-Transcribe-2.
Diarization: identifying "who speaks when" in a recording — an MAI-Transcribe-2 feature, indispensable in meetings.
Arena (LMArena): community leaderboards by blind human votes — MAI-Image-2.5 is #3 in text-to-image and #2 in editing.
Distillation: training a small model by making it imitate an existing large model — the practice Suleyman claims not to have used for the MAIs ("from scratch, no distillation").
Neuralese: the hypothetical internal-reasoning language proper to networks, inaccessible to inspection — the Code of Conduct forbids it as hidden reasoning.
Model welfare: the idea that models' well-being would be a moral consideration — defended in Anthropic's approach, explicitly rejected by Microsoft's Code.
Legal personhood: recognizing legal personality for an AI — set aside by the Code.
Cross-region (AWS Bedrock): deploying a model across multiple cloud regions — Grok 4.6 is available there from August 19 at the same $2/$6 tariff.
Always-on: an assistant active continuously, not only on request — Microsoft Scout's paradigm.
Work IQ: the orchestration and work-knowledge layer powering Copilot Cowork.
All-stock: a merger paid entirely in shares, no cash — the SpaceX–xAI case ($1.25 trillion, February 2).
AOSP: Android's open-source base — MDEP, Project Solara's lightweight OS, rests on it (not Windows).
Entra: Microsoft's identity and access platform (ex-Azure AD) — it governs Scout's identity.

<a id="s05-39"></a>
### Comparison table: Grok 4.5, 4.6, and 4.7 at a glance

| Criterion | Grok 4.5 (July 8) | Grok 4.6 (August 12) | Grok 4.7 (September 21) |
|---|---|---|---|
| Nature | Post-training upgrade, V9 base | Post-training upgrade, V9 base | Post-training upgrade (V9 base per the lineage's logic) |
| Context | 500K ctx | Unspecified (facts) | Unspecified (facts) |
| Price | $2 / $6 per M tokens | $2 / $6 per M tokens | $2 / $6 per M tokens, unchanged |
| Co-development | With Cursor | Not documented | Not documented |
| Positioning | "Opus-class" coding/agents | Not explicitly renewed | Not explicitly renewed |
| Distribution | xAI API | + AWS Bedrock from August 19 (cross-region); GitHub Copilot in August (with 4.5) | Inherits existing channels |
| Schedule | Release with no documented drama | Release with no documented drama | 5 delays; Musk publicly downgrades it before release |

This table invites several readings.
First, the stability of method and price contrasts with the growing instability of the schedule: one month between 4.5 and 4.6, then more than a month and five delays for 4.7.
Second, distribution is the real differentiator between versions: 4.5 proves (Cursor, Opus-class), 4.6 distributes (Bedrock, GitHub Copilot), 4.7 suffers (delays, downgrade).
It is a three-act trajectory: proof, expansion, trial.
Third, the facts' silence on 4.6 and 4.7's context and positioning must not be filled with suppositions: the anti-fabrication rule forbids it, and it is also information — xAI communicates less about specs as the lineage matures, and more about distribution and uses.
Finally, the 4.7 column shows that even a light-upgrade strategy hits walls: five delays on a mastered base (V9) is a sign that post-training is no formality.
The table also suggests a working hypothesis for what follows — to handle with care: if 4.7 cost five delays, it may be that the V9 base is reaching the end of what post-training can extract from it, and that 4.8 (2.5T parameters, in-house C++ stack) is the answer to that exhaustion.
That would be the cycle's logic: iterate fast while the base allows, then refound when it no longer does.
The facts do not explicitly confirm it — but they do not contradict it either, and the calendar proximity (4.8 announcement September 13, 4.7 release September 21) makes the reading plausible.
Either way, this table will remain the compact reference for the 2026 Grok 4 lineage: three upgrades, one price, one method, and a chaotic release to finish.

<a id="s05-40"></a>
### Table: the MAI family at a glance

| Model | Segment | Verified characteristics | Targeted reference |
|---|---|---|---|
| MAI-Thinking-1 | Reasoning (flagship) | 35B active MoE, 256K ctx; preferred over Sonnet 4.6 (humans); = Opus 4.6 on SWE-Bench Pro | Anthropic (Sonnet/Opus 4.6) |
| MAI-Code-1-Flash | Light code | 5B; GitHub Copilot, VS Code | Daily developer use |
| MAI-Image-2.5 / Flash | Image | #3 Arena text-to-image; #2 image-to-image editing | Arena leaders |
| MAI-Transcribe-1.5 | Transcription | 43 languages; $0.36/hour | Installed base |
| MAI-Transcribe-2 | Transcription | September 3; #1 FLEURS 60 languages (WER 5.2%); 10× GPT-Transcribe; $0.10/h promo through December 31 | OpenAI (GPT-Transcribe) |
| MAI-Voice-2 (+ Flash coming) | Voice / TTS | Multilingual, voice cloning | Complete voice chain |

This table shows a range designed as a system, not a collection.
Each model has its role in the agentic chain: Thinking-1 reasons, Code-1-Flash assists the developer, Image-2.5 creates visuals, Transcribe-2 listens, Voice-2 speaks.
Together, they form a complete agent — the one running in Scout, Cowork, and tomorrow the Solara gadgets.
The "targeted reference" column is the most political: Microsoft does not merely publish models, it designates its adversaries — Anthropic in reasoning and code, OpenAI in transcription, the Arena leaders in image.
It is a declaration of war benchmark by benchmark, segment by segment.
Also note the Flash-variant logic: Code, Image, Voice — wherever latency makes the experience, Microsoft declines a fast version.
Only Thinking-1 (heavy reasoning, where quality rules) and Transcribe (where version 2 crushes 1.5 on all fronts) escape the Flash treatment.
The Transcribe-1.5 → Transcribe-2 dynamic is a textbook case of assumed cannibalization: better in everything (60 languages versus 43, #1 on FLEURS, 10× faster) and 3.6× cheaper on promo.
Microsoft does not hesitate to kill its own model — the sign of a company that privileges conquest over rent.
Finally, this table is incomplete by construction: seven models announced at Build, six rows here — the seventh is not detailed in the verified facts beyond the announcement.
That is a limit of the dossier, not of the range: the MAI family is vaster than the facts allow describing finely.

