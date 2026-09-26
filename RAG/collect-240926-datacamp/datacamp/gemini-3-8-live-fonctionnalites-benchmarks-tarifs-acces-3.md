---
id: collect-240926-datacamp/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces-3
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: datacamp
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["gemini", "agent", "agentic", "agents", "astra", "consumer", "cost", "gemini 3.8", "gpt-live", "guardrails", "latency", "reasoning"]
source: docs/RAG/clean_en/datacamp/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [151, 201]
sha256: 06d01f46b1dfa1126c65be7209807f36c74f51a6cd74f1526845c5995899b360
---

# gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces

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
