---
id: collect-240926-datacamp/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: datacamp
role: reference
task: reference
actors: ["Google", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-09-15"]
keywords: ["benchmark", "benchmarks", "gemini", "agent", "agentic", "agents", "astra", "consumer", "cost", "gemini 3.8", "gpt-6", "gpt-live"]
source: docs/RAG/clean_en/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [1, 201]
sha256: d367e958cc44a88794cb10ddbf49fa1897fe3bbc41fea46a76bfd07837e714c7
---

# gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces

<!-- source: https://www.datacamp.com/fr/blog/gemini-3-8-live -->

Course

Voice is the playground of the moment. On Artificial Analysis's Speech to Speech Index, OpenAI's GPT-Live-1 Astra and SpaceXAI's Grok Voice Think Fast 2.0 were neck and neck at a 0.2-point gap, and OpenAI launched GPT-6 Astra in early September. On September 15, Google responded with two new speech-to-speech models: Gemini 3.8 Live and Gemini 3.8 Live Extended Thinking.

Extended Thinking takes the top spot on that index with a score of 82.6 and leads the agentic benchmark τ-Voice at 68.6%. The base 3.8 Live model is the cheapest: in Artificial Analysis's cost test, it processes an hour of input audio for $0.84, the lowest of all models charted by Google. Both are available today in the Gemini API, at the same price as their predecessor, Gemini 3.1 Flash Live.

In this article, we review all the new features of Gemini 3.8 Live: capabilities, benchmarks, differences between variants, and running costs. Also check out our guides to get started with the Gemini Live API and build a voice assistant with GPT-Live-1.

## In brief

- Gemini 3.8 Live and 3.8 Live Extended Thinking are Google's new speech-to-speech models for voice agents, replacing Gemini 3.1 Flash Live.
- Extended Thinking reasons and calls tools in the background while speaking, and now tops Artificial Analysis's speech-to-speech and τ-Voice rankings.
- The base model is fast and inexpensive but weak on multi-step agentic tasks: choose Extended Thinking as soon as tools take more than an instant to respond.
- Pricing is unchanged from the previous generation: migrating costs nothing beyond changing the model ID.
- If you're on Gemini 3.1 Flash Live, move to the new version. If you're using OpenAI's real-time stack, the price gap alone justifies an afternoon's trial.

## What is Gemini 3.8 Live?

Gemini 3.8 Live is Google's native speech-to-speech model for real-time voice agents, released on September 15, 2026 alongside a twin with stronger reasoning, Gemini 3.8 Live Extended Thinking. Both run via the Gemini Live API over WebSocket, accept audio, video, images, and text as input, and return audio. Google positions 3.8 Live as the default choice for low-latency dialogue and Extended Thinking for complex, multi-step tasks.

The generational leap over Gemini 3.1 Flash Live concerns the handling of "work that isn't speech." Tool and API calls now run in the background by default while the model continues the conversation, and Extended Thinking adds configurable background reasoning on top. Google presents this duo as a break from its previous live models and as a replacement for cascaded pipelines that chain speech recognition, an LLM, and speech synthesis.

## Key features of Gemini 3.8 Live

Google highlights five capabilities for the new models, and the two that matter most to anyone building a voice agent are the ones that change the conversational loop: asynchronous tool calls and background reasoning.

### Keep talking while tools execute

Both models execute function calls and API requests in the background while streaming audio to the user. On Gemini 3.8 Live, asynchronous execution is now the default mode for function calling. You can still force the old synchronous behavior with `behavior: BLOCKING` on a tool declaration, and the model supports function scheduling with the `SILENT`, `WHEN_IDLE`, and `INTERRUPTED` options that determine when a result is spoken.

Extended Thinking goes further: it requires non-blocking tools and returns an error if you declare a blocking one. Concretely, a caller requesting a reservation change first hears an acknowledgment, then a progress update, then the result, instead of the silence a cascaded pipeline produces while waiting for an API. On Gemini 3.1 Flash Live, function calling was only sequential, and the model wouldn't start responding until you sent back the tool result.

### Reason in the background without going silent

Gemini 3.8 Live Extended Thinking reasons and speaks at the same time. Google's documentation describes the model using early verbal signals like "Let me check" to acknowledge, then narrating progress as multi-step background work executes. You control depth with `thinking_level` set to `low`, `medium`, or `high`; the `MINIMAL` level present on 3.1 Flash Live disappears.

This changes the protocol, and in my view it's the most important migration point in this release. Because the model can speak multiple times during a single request, `turnComplete: true` no longer means it's at rest.

Your client must monitor a new `interaction_status` field, which reports `IN_PROGRESS` while reasoning or tools are running and `IDLE` when the task is fully complete. Without this, your interface will switch back to "listening" mode while the model is still hard at work.

### See what the user sees

Gemini 3.8 Live anchors dialogue in live visual input, processing video frames near real time so an agent reacts to what the user is looking at as much as to what they're saying. Google's demos include a live chess game and an employee onboarding where the model answers questions about what's on screen.

The video isn't free, though. Turning on a session now sends audio activity plus all video frames to the model by default; if your application doesn't need vision, send frames only when they're useful, both for the context budget and for cost. Audio+video sessions are also limited to 2 minutes before you need session management to extend them, versus 15 minutes for audio-only sessions.

### No more getting codes and case numbers wrong

Google calls this "alphanumeric accuracy": correctly parsing strings like confirmation codes, case numbers, or technical identifiers dictated out loud. Anyone who has ever built a voice agent for support or logistics knows this is where cascading stacks stumble: a speech recognition error at the start is unrecoverable downstream. A native voice model that hears the full utterance in context has a structural advantage here.

### Switching languages mid-sentence

Gemini 3.8 Live detects and switches between 97 supported languages during a conversation, with what Google calls accent consistency. Both models also support what Google calls incremental content updates: you can inject structured data into the session at any time, with an explicit `user` or `model` role, and the model merges it with the live audio. This is how you inject a customer record or an order status into an ongoing call without breaking the flow.

Two small changes affect existing code. Proactive audio, where the model can decide not to respond to speech not directed at it, is now permanently enabled, and setting it to `false` returns an error. Affective dialog has been entirely removed from the API; any `enable_affective_dialog` configuration must be removed.

## How does Gemini 3.8 Live perform on benchmarks?

Extended Thinking leads every quality and agentic table published by Google, but only slightly ahead of OpenAI's GPT-Live-1 Astra, while the base 3.8 Live model wins big on cost but loses on agentic tasks.

The index, τ-Voice, Big Bench Audio, and the cost figures below all come from Artificial Analysis's public leaderboard, so independently measured and not vendor-reported. The τ³-Banking figures come from Google's launch chart citing Sierra: treat that line as vendor-reported until Sierra's table reflects the same results.

### Agentic task completion

Gemini 3.8 Live Extended Thinking leads on τ-Voice, Sierra's benchmark measuring whether a voice agent completes multi-step tasks with tools, as measured by Artificial Analysis. GPT-Live-1 Astra follows closely, and the rest of the pack falls away:

- **Gemini 3.8 Live Extended Thinking:** 68.6%
- **GPT-Live-1 Astra:** 67.9%
- **Grok Voice Think Fast 2.0**: 56.5%
- **Gemini 3.1 Flash Live:** 37.7%
- **Gemini 3.8 Live (base version):** 30.1%

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

```
import asyncio
from google import genai
client = genai.Client()
config = {"response_modalities": ["AUDIO"], "output_audio_transcription": {}}
async def main():
    async with client.aio.live.connect(model="gemini-3.8-live", config=config) as session:
        await session.send_client_content(
            turns={"role": "user", "parts": [{"text": "Read back the claim number 7Q-4418-B."}]},
            turn_complete=True,
        )
        async for response in session.receive():
            if response.server_content and response.server_content.output_transcription:
                print(response.server_content.output_transcription.text)
asyncio.run(main())
```
If you're migrating from `gemini-3.1-flash-live-preview`, change the model ID and remove any `thinking_level` or `thinking_config` from your configuration; the turn cycle otherwise stays the same. The Live API is server-to-server by default; browser and mobile clients require ephemeral tokens.

For a complete step-by-step on audio capture, playback, and session management, see our Gemini Live API tutorial, and for the text side of the same family, our Gemini 3.8 Flash API tutorial covers thinking levels and function calling in Python.

## Final thoughts

Google is betting on price rather than raw quality. Extended Thinking beats GPT-Live-1 Astra by one or two points on every table, but Gemini 3.8 Live at $0.84 per hour of input audio costs nearly 7 times less than OpenAI's model, and that number is what steers voice traffic in production.

My take: if you're running anything on Gemini 3.1 Flash Live, update this week: the price is the same and asynchronous tool calling alone is worth the switch. If you're on OpenAI's realtime stack, the base model deserves a try for sorting and search flows, and Extended Thinking deserves one wherever your tools take seconds. The base model's 30.1% on τ-Voice is the reason not to use it for agentic work.

To properly build the tool-calling part of a voice agent, we recommend our Building AI Agents with Google ADK course, which connects Gemini to a customer support agent with tools, guardrails, and delegation.

## FAQ

### What is the difference between Gemini 3.8 Live and Gemini 3.8 Live Extended Thinking?

Gemini 3.8 Live is Google's low-latency speech-to-speech model for direct voice tasks, with interleaved reasoning and no thinking level setting. Gemini 3.8 Live Extended Thinking adds configurable background reasoning (low, medium, or high) and announces progress updates while it runs asynchronous tools. Extended Thinking scores 68.6% on τ-Voice versus 30.1% for the base model: choose it for multi-step agentic work.

### How much does Gemini 3.8 Live cost?

Both models share a single price grid on the paid offering: $3.00 per 1 million input audio tokens (about $0.005 per minute) and $12.00 per 1 million output audio tokens, thinking tokens included (about $0.018 per minute). Text costs $0.75 per 1 million input tokens and $4.50 per 1 million output tokens. A free tier covers both models, with free-tier data used to improve Google's products.

### Where can you access Gemini 3.8 Live?

Developers can use `gemini-3.8-live` and `gemini-3.8-live-extended-thinking` via the Gemini Live API and in Google AI Studio, where both are generally available. Enterprises access them in private preview in Gemini Enterprise. On the consumer side, Gemini 3.8 Live appears in Search Live and Extended Thinking in Gemini Live, as well as in Docs, Gmail, and Keep for Google AI subscribers.

### How does Gemini 3.8 Live compare to Gemini 3.1 Flash Live?

Gemini 3.8 Live replaces the 3.1 Flash Live preview at the same price, with asynchronous function calling enabled by default and better scores on Google's tables: 76.0 versus 71.5 on Artificial Analysis's Speech to Speech Index. Migrating means changing the model ID and removing any `thinking_level` or `thinking_config` from the session configuration. Google recommends that all 3.1 Flash Live users move to 3.8 Live.

### Does Gemini 3.8 Live support function calling and video input?

Yes. Both models support function calling, and tool calls run in the background while the model keeps talking. Gemini 3.8 Live also accepts video frames and images in addition to audio and text, so the agent can react to what the user is looking at. Audio+video sessions are limited to 2 minutes and audio-only sessions to 15 minutes, unless you use session management to extend them.

**Data Science Editor-in-Chief at DataCamp |** **I am passionate about forecasting and development using APIs.**
