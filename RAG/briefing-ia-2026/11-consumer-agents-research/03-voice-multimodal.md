---
id: briefing-ia-2026/11-consumer-agents-research/03-voice-multimodal
title: "Voice and multimodal: the agentic interface race"
domain: consumer-agents-research
role: deep-dive
task: product
actors: ["Google", "Inflection AI", "Meta", "Microsoft", "OpenAI"]
dates: ["2026-07", "2026-07-08", "2026-09", "2026-09-03", "2026-09-15"]
keywords: ["agent", "agentic", "multimodal", "voice", "agents", "benchmark", "consumer", "exploit", "full-duplex", "gemini", "gemini 3.8", "gpt-live"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s11-5"
source_lines: [12044, 12134]
sha256: 35384efee35f4aa6660710b42d6b7dadc906cc51bedd4cda49a065aeb8614f65
---

# Voice and multimodal: the agentic interface race

<a id="s11-5"></a>
### 11.4 Voice and multimodal: the race for voice as the agentic interface

**GPT-Live (OpenAI, 08/07/2026).** OpenAI launched GPT-Live
on 8 July 2026: a full-duplex voice — the user can
interrupt the model as in a human conversation, without
waiting for the end of its turn — powered by the
GPT-Live-1 model billed at $0.05 per minute, with a
cheaper Live-1 mini variant. Full-duplex is the technical
threshold that moves voice from "improved dictation" to
credible conversational interface; per-minute pricing
($0.05) also makes it an API product, not just a consumer
feature. The cheaper mini variant follows the now classic
logic: a premium model for quality, an economical model
for volume.

**Gemini 3.8 Live and Live Extended Thinking (Google,
15/09/2026).** Google responded on 15 September 2026 with
Gemini 3.8 Live, native speech-to-speech models — speech
is processed as speech end to end, with no detour through
text — covering 97 languages with mid-conversation
language switching. Pricing is aggressive: $0.005 to
$0.018 per audio minute, an order of magnitude below the
GPT-Live-1 rate ($0.05/min). The "Live Extended Thinking"
option adds extended reasoning to the voice conversation,
which addresses the classic reproach against real-time
voices: fluidity at the expense of depth. With 97
languages and mid-conversation switching, Google
explicitly targets real multilingual use — mixed
conversations, travel, international support — rather
than the simple monolingual demo.

**MAI-Voice-2 (Microsoft).** Microsoft rounds out the
picture with MAI-Voice-2, a multilingual text-to-speech
(TTS) model incorporating voice cloning, with a Flash
variant to come for low-latency use. Voice cloning is the
feature that turns TTS into a personalization product:
the agent can speak with a chosen or reproduced voice,
which reinforces the relational dimension of personal
agents (cf. Pi Journeys, Scout and its nameable
identity). The announced Flash variant signals that
latency remains the battleground: a voice that thinks
too long before answering breaks the conversational
illusion.

**MAI-Transcribe-2 (Microsoft, 03/09/2026).** On the
comprehension side, MAI-Transcribe-2 claims first place
on the FLEURS benchmark across 60 languages, with a word
error rate (WER) of 5.2%, a speed 10 times faster than
GPT-Transcribe and a price of $0.10 per hour of audio.
Three figures that together tell a story: transcription
has become a fast, cheap commodity ($0.10/hour, 10×
faster), while multilingual quality (#1 FLEURS, 60
languages, 5.2% WER) remains a differentiator. For
agents, reliable, near-free transcription is an
infrastructure brick: meetings, calls, voice notes become
inputs the agent can exploit without friction.

**Voice as the agentic interface.** Taken together, these
four announcements describe a structured race: OpenAI
sets the interaction standard (full-duplex), Google
smashes prices and pushes native multilingual (97
languages, $0.005/min), Microsoft locks down both ends of
the chain (personalized synthesis with MAI-Voice-2,
comprehension with MAI-Transcribe-2). The stake is no
longer "talking to an AI" but making voice the primary
interface of personal agents — hands-free, eyes-free,
which is the existence condition for concepts like
Solara's Badge or the AI glasses planned for Muse.
Pricing (from $0.005 to $0.05 per minute) remains the
parameter that will decide whether voice becomes the
default interface or a premium: at $0.005/min, an hour of
conversation costs $0.30; at $0.05/min, it costs $3 — a
factor of 10 separating occasional use from ambient use.

**TTS and transcription as commodities.** MAI-Voice-2 and
MAI-Transcribe-2 tell another story than the race for
conversational models: that of the commoditization of
bricks. Transcription at $0.10 per hour, ten times faster
than the previous reference, is no longer a product — it
is infrastructure, like storage. Likewise, multilingual
TTS with voice cloning makes voice personalization a
standard option rather than a luxury. The consequence for
agents is direct: when hearing (transcription) and
speaking (synthesis) become near-free and instantaneous,
voice stops being a costly channel to optimize and
becomes the default channel — which strengthens the
thesis of ambient devices (Solara's Badge, Muse's AI
glasses), designed for a world where one talks to agents
continuously.

