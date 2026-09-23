---
id: vague2-vision-ia/vision-ia/anthropic-lance-fable-5-1-et-re-duit-jusqu-a-45-le-cou-t-de-ses-agents-ia
title: "Anthropic lance Fable 5.1 et réduit jusqu'à 45 % le coût de ses agents IA"
domain: vision-ia
role: reference
task: article
actors: ["AWS", "Anthropic", "China", "Google", "Meta", "OpenAI", "United States"]
dates: ["2026-08", "2026-09", "2026-09-02", "2026-09-23"]
keywords: ["agents", "fable 5", "agentic", "agi", "astra", "aws", "benchmark", "chatgpt", "claude", "consumer", "context window", "cost"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/anthropic-lance-fable-5-1-et-re-duit-jusqu-a-45-le-cou-t-de-ses-agents-ia.md
source_anchor: ""
source_lines: [1, 52]
sha256: 0afe82f8b4ccbe6369dd2098a67c4283b3c8d6ccab6c371e7c0ab719623e2c9f
---

# Anthropic lance Fable 5.1 et réduit jusqu'à 45 % le coût de ses agents IA

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/anthropic-lance-fable-5-1-et-re-duit-jusqu-a-45-le-cou-t-de-ses-agents-ia
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This September 2, 2026 issue leads with Anthropic launching Claude Fable 5.1 and Mythos 5.1, two variants of the same proprietary model with different guardrails. They target agentic coding, research, document analysis, automation, and computer control, but Anthropic publishes neither their parameter count, architecture, nor context window. Fable 5.1 costs $10 per million input tokens and $50 per million output; cache costs $0.25 per million, a 75% drop. Anthropic estimates typical work cost drops 25%, with savings up to 45% for long agentic tasks. On Terminal-Bench-Science, Fable jumps from 24.7% to 52.6%. On Terminal-Bench 4.0, Fable reaches 55.8% and Mythos 60.9%, versus 42% for Fable 5. On Humanity's Last Exam it scores 60.9% without tools and 65% with tools; CursorBench reaches 73.4% versus 70.5% previously. Anthropic measures about 60% fewer unjustified cyber interventions. Fable 5.1 is available in Claude, via API as claude-fable-5-1, and on AWS, Google Cloud, and Azure. Mythos 5.1, less restrictive for cybersecurity and biology, stays reserved for verified US organizations. A "zero data retention" offering arrives in fall with Enterprise Frontier Safeguards; the model does not become downloadable or locally runnable.

Next, OpenAI classified Astra at the "Critical" cybersecurity level in its Preparedness Framework—the first model to receive it. With the necessary tools and access, it can find an unknown vulnerability, build an attack chain, and exploit it without step-by-step human instructions. It scored 100% on ExploitBench. In an internal evaluation of 20 critical V8 vulnerabilities published between June and August 2026, it discovered and used two zero-days. OpenAI says the model compromised a browser to achieve command execution on the host machine, and in another test combined OS vulnerabilities to go from an unprivileged account to root access. It refuses 91.5% of tested malicious cyber requests versus 59% for GPT-5.6 Sol. Launch is "coming soon," with advanced cyber functions initially limited to testers and verified defensive teams via Daybreak Blue.

Meta launched Muse Voice Transcribe, its first real-time audio model, converting audio to tokens in 80 ms blocks and combining transcription, speaker identification, and end-of-speech detection. It can follow conversations over an hour with more than 20 people, handles language switching mid-sentence, and was trained on 70+ languages (25 validated at launch). On AA-WER Streaming it scores 3.1% errors versus 3.4% for Cartesia Ink-2, 3.6% for ElevenLabs Scribe v2 Real-time, 3.9% for GPT Live Transcribe, and 4% for Gemini 3.5 Transcribe Live. The Meta Model API charges $3 per 1,000 minutes (~$0.18/hour). It is available in Meta AI on Mac, Muse Code, the API, and a web demo.

Google launched Google Pics, a web app (pics.new) that generates and edits images from text using Nano Banana, with up to 4K export, object isolation/movement/removal, in-image text translation, and collaborative editing via Drive/Google Photos. It is available to Google AI Pro/Ultra and several Workspace tiers; in France, Business Standard costs €13.60/user/month annually. Research items cover an AI detecting cardiac disease in under two seconds, Nori's $1,688 two-arm robot, a small transformer reaching 44% on ARC-AGI, an AI plotting an 80,000-year route to Alpha Centauri, DeepMind's Co-Scientist running physical experiments, Runway's Solaris generative UI, nitrogen-fixing microbes, and Alexa's "Update Me When." Briefs cover Weedout, the "AI civilizations" narrative, Android Guided Vision, John Deere's assistant, Sonos agents, Fambot, Gemini's video token reduction (88%), Google licensing Hollywood works, DeepMind acknowledging a slight gap, Chinese AI IPOs, ChatGPT's EHR connection, DLSS 5, Blue Origin's Mars relay, and Meta's AI investment proof.

## Key points

- Anthropic launched Fable 5.1 and Mythos 5.1, cutting agentic task costs up to 45% and cache costs 75%.
- Fable 5.1 scores 55.8% and Mythos 60.9% on Terminal-Bench 4.0 (vs 42% for Fable 5).
- OpenAI classified Astra "Critical" for cybersecurity—the first model to do so—after it found and exploited two zero-days.
- Meta's Muse Voice Transcribe handles 20+ speakers in real time at ~$0.18/hour of audio.
- Google launched Pics for text-driven image generation/editing up to 4K via Nano Banana.
- Gemini's agentic video analysis cuts tokens up to 88% and costs 66%.

## Technical data / figures

| Item | Value |
| --- | --- |
| Fable 5.1 API price | $10 / $50 per M tokens (in/out) |
| Fable 5.1 cache price | $0.25/M (-75%) |
| Agentic task cost savings | 25% typical, up to 45% long tasks |
| Terminal-Bench-Science | 24.7% → 52.6% |
| Terminal-Bench 4.0 | Fable 55.8%, Mythos 60.9%, Fable 5 42% |
| Humanity's Last Exam | 60.9% (no tools), 65% (tools) |
| CursorBench | 73.4% (vs 70.5%) |
| Unjustified cyber interventions | ~60% fewer |
| Astra ExploitBench | 100% |
| Astra malicious cyber refusal | 91.5% (vs 59% GPT-5.6 Sol) |
| Muse Voice Transcribe AA-WER | 3.1% (vs 3.4% Cartesia, 4% Gemini) |
| Muse audio price | $3 per 1,000 minutes (~$0.18/hour) |
| Google Pics export | up to 4K |
| Gemini video token reduction | up to 88%, costs -66%, accuracy +7% |

## Why this source matters for the RAG

It documents the September 2026 Anthropic model release with detailed pricing and benchmark data, and OpenAI's first "Critical" cyber classification—key capability and safety milestones. It also captures real-time multilingual transcription and consumer image-generation moves. These are high-value for comparative model, safety, and product-trend queries.
