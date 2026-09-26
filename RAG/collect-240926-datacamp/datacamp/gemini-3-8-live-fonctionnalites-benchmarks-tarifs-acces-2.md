---
id: collect-240926-datacamp/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces-2
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: datacamp
role: reference
task: reference
actors: ["Google", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["benchmark", "gemini", "agent", "agents", "astra", "cost", "gemini 3.8", "gpt-live", "grok", "jailbreak", "latency", "leaderboard"]
source: docs/RAG/clean_en/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [75, 150]
sha256: df2d38ffe1c60ca041f621091921c17fdb7020824ea2db741dd88a493494d3ef
---

# gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces

The surprising figure is that the base model's τ-Voice score is lower than its own predecessor's. Google is clear: 3.8 Live is designed for scale and cost efficiency rather than complex workflows, but a gap of more than 38 points between the two variants means the tier choice isn't a nuance. If your agent chains tools together, the base model isn't the right default.

On Sierra's τ³-Banking leaderboard, a harder customer service benchmark focused on banking workflows, Extended Thinking scores 35.1% versus 32.0% for GPT-Live-1 Astra, 16.5% for SpaceXAI's Grok Voice Think Fast 2.0, 11.3% for Gemini 3.1 Flash Live, and 10.3% for GPT-Realtime 2. Every model in that table fails most of the time, a sign of how far there still is to go before an autonomous voice agent in banking, but tripling the previous Gemini generation's score is real progress.

### Voice quality and reasoning

On the Speech to Speech Index, Extended Thinking also leads the competition:

- **Gemini 3.8 Live Extended Thinking:** 82.6
- **GPT-Live-1 Astra:** 81.5
- **Grok Voice Think Fast 2.0:** 81.3
- **Gemini 3.8 Live (base):** 76.0
- **Gemini 3.1 Flash Live:** 71.5

A 1.1-point lead over OpenAI is still a lead, but I wouldn't base a purchasing decision on it alone.

For pure reasoning on spoken input, Google reports 97.7% on Big Bench Audio for Extended Thinking. Google also says both models push the Pareto frontier on ServiceNow's EVA-Bench for voice agents, balancing accuracy and conversational quality, though that test was run via the Live API through the Gemini Enterprise Agent Platform rather than the public API.

### Cost per hour of audio

This is the chart Google most wants to show you. In Artificial Analysis's cost test on a subset of Big Bench Audio, Gemini 3.8 Live processes an hour of input audio for $0.84.

Extended Thinking costs $3.50 for the same hour, Grok Voice Think Fast 2.0 costs $4.80, and GPT-Live-1 Astra $5.83. Gemini 3.1 Flash Live was at $1.50 and $1.75 depending on the "thinking" level.

Reading the two Gemini bars together, the product strategy is obvious. The base model crushes every other one on price, and the reasoning model that rivals OpenAI on quality still costs 40% less per hour of input. Note that the reasoning model consumes more tokens at higher thinking levels, hence a cost roughly 4 times that of its twin despite a shared pricing grid.

## Which variant should you choose?

Gemini 3.8 Live is the right default for most voice agents, and Extended Thinking only justifies its higher token consumption when the agent needs to plan, compare, or wait on slow tools. Google's recommendations draw the line at tool latency: if your functions respond in milliseconds, stay on the base model.

Both variants share the same pricing grid (see below), token limits of 131,072 input and 65,536 output, and the same WebSocket endpoint. What distinguishes them is the reasoning architecture.

Gemini 3.8 Live uses interleaved reasoning with fixed latency and no `thinking_level` setting. Extended Thinking exposes background reasoning levels `low`, `medium`, and `high`, streams conversational "fillers" while it works, and requires asynchronous tools. These thinking levels are the only effort lever in this version.

| Use case | Choice | Why |
|---|---|---|
| Customer service triage, voice search, language practice | Gemini 3.8 Live | Immediate exchange matters more than depth, and each user turn calls for a single response |
| Smart home control, sensor reading | Gemini 3.8 Live | Tools respond in milliseconds, nothing to mask |
| Technical support on logs, error codes, and configuration checks | Extended Thinking | A multi-step diagnosis requires background reasoning between utterances |
| Travel and booking querying flights and hotels in parallel | Extended Thinking | Parallel asynchronous calls with spoken progress updates rather than silence |
| STEM and coding assistance | Extended Thinking | The model verifies formulas or debugs code before stating an explanation |

One last point: client-side complexity. Switching to Extended Thinking means rewriting your state management around `interaction_status`; if you have a working 3.1 Flash Live app, the base model is a like-for-like upgrade, and the reasoning model requires a small refactor.

## Pricing and availability of Gemini 3.8 Live

Both models share the pricing grid of Gemini 3.1 Flash Live Preview, so the upgrade is free. On the paid tier, input audio costs $3.00 per million tokens (Google estimates it at $0.005 per minute) and output audio $12.00 per million tokens, thinking tokens included (about $0.018 per minute). Thinking tokens are billed at the output rate, which explains the higher runtime cost of Extended Thinking.

| Modality | Paid tier, per 1M tokens | Estimate per minute |
|---|---|---|
| Input text | $0.75 | n/a |
| Input audio | $3.00 | $0.005/min |
| Input image and video | $1.00 | $0.002/min |
| Output text | $4.50 | n/a |
| Output audio (thinking tokens included) | $12.00 | $0.018/min |

The free tier covers both models at no charge for input and output, with the usual caveat: Google uses free-tier data to improve its products; paid-tier data is not used for that purpose.

Grounding with Google Search is supported on both tiers, with 5,000 free search queries per month shared across all Gemini 3.x models, then $14 per 1,000 queries. Google has not published rate limits specific to these models; check the Gemini API limits page for your tier before planning a launch.

Availability comes in three parts:

- **Developers:** both models are generally available in the Gemini API and Google AI Studio as of September 15.
- **Enterprises:** available in private preview in Gemini Enterprise and soon in Gemini Enterprise for Customer Experience, with Extended Thinking also coming for Google Workspace business customers.
- **General public:** Gemini 3.8 Live powers Search Live, while Extended Thinking is coming to Gemini Live, as well as Docs for Google AI Pro and Ultra subscribers, and Gmail and Keep for all Google AI subscribers.

All audio output from both models carries a SynthID watermark, and Google's model card is transparent about the limitations: models can still hallucinate, jailbreak resistance is improving, and there may be occasional slowness or timeouts. Worth reading before putting a model in front of customers.

## How to get access to Gemini 3.8 Live

The model IDs are `gemini-3.8-live` and `gemini-3.8-live-extended-thinking`, two stable strings with no preview suffix.

You can access them via the Gemini Live API with the google-genai SDK for Python or JavaScript, via raw WebSockets, or interactively in the Stream view of Google AI Studio. Google also lists Agora, Fishjam, LiveKit, LangChain, Pipecat, Vercel, and Vision Agents as integration partners handling the media streaming layer, and a streaming path via the Agent Development Kit if you are already building on ADK.

The minimal Python session below opens a connection, sends a text turn, and enables an output transcription to read what the model said:

