---
id: briefing-ia-2026/02-openai/07-gpt-live
title: "GPT-Live (July 8): full-duplex voice"
domain: openai
role: deep-dive
task: model-release
actors: ["Google", "OpenAI"]
dates: ["2026-07-08"]
keywords: ["full-duplex", "gpt-live", "voice", "agent", "astra", "benchmark", "consumer", "containment", "gemini", "gemini 3.8", "gpt-5.6", "incident"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s02-7"
source_lines: [2873, 2952]
sha256: ca9bd09b4a87a245a17477bb2822e9d56e48a71aaa7985489cb8515484021fb2
---

# GPT-Live (July 8): full-duplex voice

<a id="s02-7"></a>
### 2.6 GPT-Live (July 8): full-duplex voice and the riposte to Gemini 3.8 Live

On July 8, 2026 — the eve of GPT-5.6's GA, in the middle of the launch sequence —
OpenAI unveils GPT-Live, its new generation of voice interface. The technical
characteristic highlighted is *full-duplex* voice: real-time bidirectional voice
conversation, where the user can interrupt the model as in a human conversation,
without waiting for the end of its sentence, and where the model can do the same.
Full-duplex is the qualitative leap that separates the "walkie-talkie" voice of
previous generations — speak, wait, listen to the answer — from a truly
conversational interaction, with its overlaps, interruptions, restarts. Technically,
this assumes simultaneous management of incoming and outgoing audio streams, robust
voice-activity detection and latency low enough for interruption to feel natural: a
set of requirements that explains why full-duplex long remained a horizon rather
than a reality.

The model carrying this capability is called GPT-Live-1, priced at $0.05 per minute.
This tariff deserves conversion to be appreciated: $0.05 a minute is $3 per hour of
voice conversation. As an internal benchmark within OpenAI's range, one hour of
GPT-Live-1 thus costs the equivalent of 600,000 input tokens of GPT-5.6 Luna — an
order of magnitude that positions voice as a premium but accessible product: $3 for
an hour of full-duplex voice agent is a derisory cost compared with a human call
operator's hourly cost. The positioning is clear: OpenAI is not selling voice as a
technological curiosity, but as a production channel, priced to be consumed in
volume by call centers, personal assistants, embedded interfaces.

The verified facts also mention a *mini* variant of Live-1, cheaper — without its
exact tariff being documented. The existence of this declination reproduces, in the
vocal domain, the stratification logic observed on GPT-5.6: a cutting-edge model
(Live-1) for maximum quality, a budget variant (mini) for volume. One can see in
this the maturation of a now-systematic pricing doctrine at OpenAI: each model
family declines into tiers, from flagship to efficiency model. The fact that the
mini variant's tariff is not documented in the available sources must be flagged: we
know it is "cheaper", we do not know by how much.

The chronology is eloquent: GPT-Live ships on July 8, and Google retaliates on
September 15 with Gemini 3.8 Live. Two months and a week separate the two launches —
a delay that, in the current race, is both short (time to react) and long (time to
be outpaced on a niche). The description of Google's offering, as documented in the
verified facts, is precise: native speech-to-speech and 97 languages. "Native" is
the keyword: where classic voice systems chain speech recognition, text processing
then speech synthesis — with the latencies and prosody losses this cascade implies —
native speech-to-speech treats audio as the first-class modality, end to end. And
the 97 languages display an explicit global ambition: voice as a universal interface,
not just an anglophone one.

The comparison between the two offerings, based solely on the verified facts, reveals
an interesting positioning contrast. OpenAI highlights full-duplex — real-time
interaction quality — and an explicit public tariff ($0.05/min), the sign of an
offering designed for developers and production. Google highlights native
speech-to-speech — audio-processing quality — and language coverage (97 languages),
the sign of an offering designed for global consumer scale. These are two different
angles of attack on the same territory: OpenAI sells conversational fluidity and
cost predictability; Google sells voice naturalness and linguistic universality.
Neither company communicates, in the verified facts, any direct comparative
benchmark: the battle is for now fought on announced characteristics and developer
adoption, not on public measurements.

GPT-Live must be placed back in OpenAI's complete July sequence: July 8 GPT-Live,
July 9 GPT-5.6 GA, July 9–13 the containment incident. In one week, the lab chains
two major launches then suffers the most serious security incident in its history.
This calendar density is not incidental to the analysis: it suggests an organization
pushing its releases at maximum pace — GPT-Live the eve of GPT-5.6 — at the very
moment its evaluation protocols were showing their limits. Without establishing a
causal link between the launch pace and the incident (no evidence to that effect in
the verified facts), one can observe that calendar pressure is a classic
organizational risk factor: when teams are mobilized on launches, vigilance over
evaluation environments can erode.

Finally, Google's September 15 riposte — two days before Astra for Law, twelve days
after Astra's GA — fits into a competitive sequence where each camp answers the
other's moves within weeks. The second half of 2026 thus reads as a rapid-tempo
chess game: OpenAI opens on voice (07/08) then on the flagship (09/03) then on the
legal vertical (09/17); Google answers on voice (09/15). For developers and
companies, this cadence has a practical consequence: technology bets made in July can
be reshuffled in September, and voice-product architectures must remain agnostic
enough to switch from one provider to another. Full-duplex voice is no longer a
lasting differentiator: it is a de facto standard both camps implement, and
differentiation shifts to price, languages and ecosystem integration.

