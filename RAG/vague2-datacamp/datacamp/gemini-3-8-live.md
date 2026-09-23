---
id: vague2-datacamp/datacamp/gemini-3-8-live
title: "Gemini 3.8 Live : fonctionnalités, benchmarks, tarifs et accès"
domain: datacamp
role: reference
task: article
actors: ["Google", "OpenAI", "xAI"]
dates: ["2026-09", "2026-09-23"]
keywords: ["benchmark", "benchmarks", "gemini", "gemini 3.8", "agent", "agentic", "agents", "astra", "cost", "gpt-live", "grok", "latency"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/gemini-3-8-live.md
source_anchor: ""
source_lines: [1, 59]
sha256: 5195724f188ede562d5284bbfa920266ae9baac2a4d3d33ad6ba16e05d806551
---

# Gemini 3.8 Live : fonctionnalités, benchmarks, tarifs et accès

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/gemini-3-8-live
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers Google's two new speech-to-speech models, Gemini 3.8 Live and Gemini 3.8 Live Extended Thinking, launched 15 September 2026 to replace Gemini 3.1 Flash Live. On Artificial Analysis's Speech to Speech Index, Extended Thinking takes first place with 82.6 and leads the agentic τ-Voice benchmark at 68.6%. The base 3.8 Live is the cheapest: in Artificial Analysis's cost test it processes one hour of input audio for $0.84, the lowest of all Google-charted models. Both are available in the Gemini API at the same price as their predecessor.

Both are native speech-to-speech models for real-time voice agents, running via the Gemini Live API over WebSocket, accepting audio, video, images, and text, and returning audio. Google positions 3.8 Live as the default for low-latency dialogue and Extended Thinking for complex multi-step tasks. The generational leap concerns handling "non-speech work": tool/API calls now run in the background by default while the model keeps talking, and Extended Thinking adds configurable background reasoning. Google frames this as a replacement for cascade pipelines (ASR + LLM + TTS).

**Five key features**: (1) continuing to speak while tools execute—async function calling is now default, with `behavior: BLOCKING` to force synchronous, and scheduling options `SILENT`, `WHEN_IDLE`, `INTERRUPTED`; Extended Thinking requires non-blocking tools and errors on blocking ones. (2) Reasoning in the background without going silent—Extended Thinking reasons and speaks simultaneously, using early verbal signals ("Let me check"), with `thinking_level` of `low`/`medium`/`high` (`MINIMAL` removed); clients must watch the new `interaction_status` field (`IN_PROGRESS`/`IDLE`) since `turnComplete: true` no longer means idle. (3) Seeing what the user sees—near-real-time video/image processing; but video isn't free, and audio+video sessions are limited to 2 minutes vs 15 minutes for audio-only. (4) Alphanumeric accuracy for dictated confirmation codes, file numbers, and technical identifiers. (5) Mid-sentence language switching across 97 languages, plus incremental content updates (injecting structured data with explicit `user`/`model` role). Two breaking changes: proactive audio is now always on (setting `false` errors), and affective dialog is fully removed.

**Benchmarks** (from Artificial Analysis's public leaderboard, independently measured): τ-Voice—Extended Thinking 68.6%, GPT-Live-1 Astra 67.9%, Grok Voice Think Fast 2.0 56.5%, Gemini 3.1 Flash Live 37.7%, Gemini 3.8 Live base 30.1% (notably lower than its predecessor). τ³-Banking (Sierra, vendor-cited): Extended Thinking 35.1% vs 32.0% GPT-Live-1 Astra, 16.5% Grok Voice Think Fast 2.0, 11.3% Gemini 3.1 Flash Live, 10.3% GPT-Realtime 2. Speech to Speech Index: Extended Thinking 82.6, GPT-Live-1 Astra 81.5, Grok Voice Think Fast 2.0 81.3, base 76.0, Gemini 3.1 Flash Live 71.5. Big Bench Audio: Extended Thinking 97.7%. Cost per hour of input audio: base $0.84, Extended Thinking $3.50, Grok Voice Think Fast 2.0 $4.80, GPT-Live-1 Astra $5.83, Gemini 3.1 Flash Live $1.50–$1.75.

**Pricing/availability**: both share Gemini 3.1 Flash Live Preview pricing (upgrade is free): paid tier audio input $3.00/1M tokens (~$0.005/min), audio output $12.00/1M (~$0.018/min, thinking tokens billed at output rate); text input $0.75, text output $4.50, image/video input $1.00. Free tier covers both (data used to improve Google products). Grounding with Google Search: 5,000 free requests/month then $14/1,000. Model IDs: `gemini-3.8-live` and `gemini-3.8-live-extended-thinking`. Available to developers (GA in Gemini API/AI Studio), enterprises (private preview in Gemini Enterprise), and consumers (Search Live, Gemini Live, Docs/Gmail/Keep). All audio carries SynthID watermarking.

## Key points

- Gemini 3.8 Live and 3.8 Live Extended Thinking replace Gemini 3.1 Flash Live at the same price.
- Extended Thinking tops Artificial Analysis's Speech to Speech Index (82.6) and τ-Voice (68.6%).
- Base 3.8 Live is cheapest at $0.84/hour input audio but weak on agentic tasks (30.1% τ-Voice).
- Async tool calling is now default; Extended Thinking adds configurable background reasoning (`thinking_level`).
- Breaking changes: clients must watch `interaction_status`; proactive audio always on; affective dialog removed.
- Supports video input (2-min limit), 97 languages, alphanumeric accuracy for dictated codes.
- Both share pricing; free tier available; model IDs `gemini-3.8-live` and `gemini-3.8-live-extended-thinking`.

## Technical data / figures

| Benchmark | Gemini 3.8 Live ET | GPT-Live-1 Astra | Grok Voice Think Fast 2.0 | Gemini 3.1 Flash Live | Gemini 3.8 Live (base) |
| --- | --- | --- | --- | --- | --- |
| τ-Voice | 68.6% | 67.9% | 56.5% | 37.7% | 30.1% |
| Speech to Speech Index | 82.6 | 81.5 | 81.3 | 71.5 | 76.0 |
| τ³-Banking (Sierra) | 35.1% | 32.0% | 16.5% | 11.3% | — |
| Cost/hour input audio | $3.50 | $5.83 | $4.80 | $1.50–$1.75 | $0.84 |

| Modality | Paid tier per 1M tokens | Per-minute estimate |
| --- | --- | --- |
| Text input | $0.75 | n/a |
| Audio input | $3.00 | $0.005/min |
| Image/video input | $1.00 | $0.002/min |
| Text output | $4.50 | n/a |
| Audio output (incl. thinking tokens) | $12.00 | $0.018/min |

- Token limits: 131,072 input / 65,536 output.
- Model IDs: `gemini-3.8-live`, `gemini-3.8-live-extended-thinking`.
- Session limits: audio+video 2 min; audio-only 15 min.
- Grounding with Google Search: 5,000 free requests/month, then $14/1,000.
- Big Bench Audio: Extended Thinking 97.7%.

## Why this source matters for the RAG

It provides detailed, current data on Google's speech-to-speech models, including feature changes, independent benchmark rankings, pricing, and migration considerations—useful for voice-agent architecture questions. It also captures the shift to async tool calling and background reasoning in real-time voice agents.
