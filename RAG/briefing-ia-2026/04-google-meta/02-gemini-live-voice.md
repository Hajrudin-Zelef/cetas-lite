---
id: briefing-ia-2026/04-google-meta/02-gemini-live-voice
title: "Gemini 3.8 Live and Live Extended Thinking: voice as a native channel"
domain: google-meta
role: deep-dive
task: product
actors: ["Google", "Meta", "OpenAI"]
dates: ["2026-09-15"]
keywords: ["gemini", "gemini 3.8", "voice", "agent", "agentic", "consumer", "cyber", "distribution", "gpt-live", "multimodal", "muse", "personal agent"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s04-3"
source_lines: [4849, 4955]
sha256: fdf5282ff3b32dcde9f2b76d0051cf54dd568f080f4a04344ea1ec45136e1e37
---

# Gemini 3.8 Live and Live Extended Thinking: voice as a native channel

<a id="s04-3"></a>
### Gemini 3.8 Live and Live Extended Thinking: voice becomes a native channel

On September 15, 2026, Google added two voice variants to the
3.8 family: Gemini 3.8 Live and Live Extended Thinking. The spec
sheet is compact but telling: native speech-to-speech, 97
languages supported with the ability to switch languages
mid-conversation, and per-audio-minute pricing — $0.005 per
minute in, $0.018 per minute out. Three facts, three strategic
signals. Native speech-to-speech means voice is no longer a
layer on top — transcription, then text reasoning, then
synthesis — but a first-class processing mode. Mid-conversation
language switching answers a real use case, that of multilingual
conversations where you hop from one language to another without
warning. And per-minute pricing aligns the price with the unit
that telephony and call-center buyers understand, not the token,
an engineer's unit.

The distinction between Live and Live Extended Thinking deserves
emphasis. The second name suggests a variant that reasons longer
before answering — the voice equivalent of what "thinking" modes
bring to text. It is a segmentation by latency and depth: a
reactive version for fluid conversation, a reflective version
for tasks demanding rigor. The same segmentation logic seen with
Flash's Cyber variant reappears: Google differentiates its
models by use case rather than offering one do-everything model.
For developers, that means choosing an explicit tradeoff between
speed, cost, and reasoning quality — a choice that, in voice, is
even more sensitive than in text, because latency is immediately
felt.

The price — half a cent per minute in, 1.8 cents out — puts
conversational voice in the zone of costs that are negligible at
the scale of one call and significant at the scale of millions
of calls. Do the math: an hour of conversation costs 30 cents in
and $1.08 out, on the order of $1.40 an hour for a continuous
dialogue. That is an order of magnitude that makes economically
thinkable uses previously reserved for demos: always-on voice
assistants, automated phone receptionists, conversational
companions. Per-minute rather than per-token pricing is also a
commercial gesture: it speaks the language of telecom budgets
and customer-relations departments, the natural buyers of these
uses.

Support for 97 languages, with mid-conversation switching, is
the other trump card. The number itself — 97 — signals an
ambition of near-universal coverage, far beyond the few dozen
languages consumer products ordinarily target. But it is the
mid-conversation switch that is the real usage differentiator:
it treats multilingualism not as a configuration option but as
a property of the conversation. In practice, humans constantly
mix languages — a Franco-German executive switching to English
for a technical term, an airport agent greeting in Spanish then
concluding in French. A system that requires choosing the
language up front breaks that naturalness; a system that follows
the speaker preserves it. It is an engineering detail that
changes the experience.

The question of voice competition remains. This dossier's
editorial framing pits Gemini 3.8 Live against OpenAI's GPT-Live,
in what it describes as a race for voice. Let us be precise
about what the verified facts allow us to say: we have Gemini
3.8 Live's spec sheet — native speech-to-speech, 97 languages,
in-conversation switching, per-minute pricing — and we know
OpenAI fields a competing voice product. This dossier, however,
does not contain GPT-Live's pricing or language-coverage sheet;
any quantified comparison would therefore be an invention. What
can be analyzed, by contrast, is the nature of the competition:
it will not be fought only on the quality of voice synthesis,
but on three fronts — per-minute price, which decides large-scale
deployments; real language coverage, which decides addressable
markets; and ecosystem integration, which decides distribution.

On this third front, Google starts with a structural advantage:
the Gemini app, with its one billion monthly users (see the next
subsection), is a native distribution channel for voice
features. A voice feature landing in an app already installed by
a billion people does not have the same fate as one that must
conquer its audience. OpenAI, for its part, has the advantage of
earlier brand association with conversational voice. The "race"
evoked by the framing is therefore less a technology sprint than
an adoption battle: who will make voice the default interaction
mode, and at what price.

Live must also be placed within Google's agentic trajectory. A
credible personal agent, of the kind Meta launched with Meta
Muse on September 8 (see below), will have to speak as much as
write: negotiating a trip by phone, calling a restaurant,
handling a dispute with a customer service. Native
speech-to-speech is the building block that makes such scenarios
conceivable without breaking the experience. By launching Live a
week after Meta Muse, Google signals it will not cede the
voice-agent field to the competition. The week of September 8–15,
2026 may go down as the one in which the personal agent became
multimodal for good.

Finally, a word on timing: Live arrives thirteen days after
Gemini 3.8 Flash. The proximity of the two launches is not a
scheduling accident; it illustrates the industrial cadence
described above. Google no longer launches models; it feeds a
lineup — and each variant strengthens the others. Flash brings
volume and low price, Live brings the voice channel, the app
brings the audience. Taken in isolation, each of these launches
is a product news item; taken together, they sketch a strategy
of encirclement of the conversational-agent market, from text to
voice, from developer to the general public.

