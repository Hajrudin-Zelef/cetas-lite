---
id: collect-240926-datacamp/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces-1
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: datacamp
role: reference
task: reference
actors: ["Google", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-09-15"]
keywords: ["benchmark", "benchmarks", "gemini", "agent", "agentic", "agents", "astra", "cost", "gemini 3.8", "gpt-6", "gpt-live", "grok"]
source: docs/RAG/clean_en/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [1, 74]
sha256: e8b20e69451f78be9550dc2d9321f63e5aeeac94fd1277954c9d16c008cb9750
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

