---
id: briefing-ia-2026/03-anthropic/04-model-lineup
title: "Anthropic 2026 model lineup: Opus, Sonnet, Fable, comparison"
domain: anthropic
role: deep-dive
task: model-release
actors: ["Anthropic"]
dates: ["2026-04", "2026-05", "2026-06", "2026-07-24"]
keywords: ["agentic", "agents", "benchmark", "claude", "copyright", "cyber", "distillation", "fable 5", "fermat", "formalization", "opus 4", "opus 5"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s03-11"
source_lines: [3752, 3921]
sha256: 06b80713f0cd0c2611f389fb6f2e2ef9de84630ca4f856fe71a6169bfef435a4
---

# Anthropic 2026 model lineup: Opus, Sonnet, Fable, comparison

<a id="s03-11"></a>
### Opus 4.6 and Opus 4.7: the temporary leaderships of April

Mid-April 2026 sees the release of Claude Opus 4.6, credited with a
"temporary retaking of the leadership." Two days to two weeks later —
on April 16 — Claude Opus 4.7 is in turn released and takes the same
"temporary retaking of the leadership." This tight sequence, two major
versions in one month with the same cautious qualifier, tells the story
of spring 2026 in frontier AI: a summit so contested that each launch is
only worth the time until the next competitor's riposte, and where the
word "temporary" is not coquetry but the honest description of a market
state.

One must measure what this cadence implies organizationally. Releasing
two Opus generations days apart, each capable of retaking — even
temporarily — the leadership, implies a training and evaluation machine
running at full tilt, with parallel rather than sequential pipelines.
The most economical hypothesis, compatible with the verified facts
alone, is that Opus 4.6 and 4.7 correspond to development branches run
in parallel whose release windows telescoped — or to a 4.7 already ready
that 4.6's short reign led them to fire off earlier than planned. In
either case, the strategic signal is identical: Anthropic refuses to let
a competitor occupy the summit even for a few weeks, even at the cost of
cannibalizing its own previous generation.

The contrast with the second half is striking: where April lined up
"temporary" leaderships, September crowns Fable 5.1 "new No. 1" without
qualification. Between the two came Opus 4.8 (May), Fable 5 (June),
Opus 5 (July) — and above all a regime change: the Mythos family has dug
a gap that competing iterations no longer close in a few days.
April 2026 will thus remain the last moment when the summit was truly
contestable with each release; after that, it is the era of installed
leadership.

<a id="s03-12"></a>
### Opus 4.8: 88.6% on SWE-bench Verified, the May reference

In May 2026 comes Claude Opus 4.8: 88.6% on SWE-bench Verified, a
one-million-token window, priced at 5 dollars per million input tokens
and 25 dollars output. Each of these three figures deserves a pause.
88.6% on SWE-bench Verified is, at this date, the highest score ever
published on the reference benchmark for agentic software engineering —
the measure, on real tasks drawn from open-source repositories, of a
model's ability to diagnose, code, and validate fixes. Such a score
signals not merely a good coding model: it signals a model capable of
executing long, reliable chains of action, the condition of any serious
agentic use.

The million-token context, next, trivializes what was still a
differentiator: Opus 4.8 displays it in May, Sonnet 5 in June, Fable 5
in June as well. The generalization of the million tokens across the
whole lineup in the space of a few weeks indicates that context length
has ceased to be an axis of differentiation and become a category
standard — exactly as high definition had been for screens. Competition
moves elsewhere: toward agentic scores, toward price, toward safeguards.

The price, finally — $5/$25 — halves that of Fable 5 ($10/$50), which
will come out a month later, and it is this relative positioning that
gives Opus 4.8 its architectural role in the lineup: it is the
"reasonable" model of the summit, the one to which Fable 5 routes its
sensitive cyber/bio/chem requests. This routing is not merely a safety
measure: it is also an implicit recognition that Opus 4.8, at half the
flagship's price, suffices for the great majority of demanding tasks.
The most expensive model delegates to the model half as expensive — an
admission that the price-performance curve, at the summit, flattens
faster than the price curve suggests.

<a id="s03-13"></a>
### Sonnet 5: the value-for-money champion (June)

In June 2026, between the release of Fable 5 (June 9) and its
suspension (June 12), Anthropic launches Claude Sonnet 5: 85.2% on
SWE-bench Verified, 2 dollars per million input tokens, 10 dollars
output, a one-million-token window. The positioning is crystal-clear:
3.4 SWE-bench points below Opus 4.8 (88.6%), Sonnet 5 costs two and a
half times less on input ($2 vs $5) and two and a half times less on
output ($10 vs $25). It is, in the 2026 lineup, the rational option par
excellence — the model you choose when you've done the math.

The release of Sonnet 5 in June, in Fable 5's media shadow, illustrates
Anthropic's lineup strategy: each capability level at the summit is
doubled with an economical option that captures its essence at a
fraction of the price. Sonnet 5 is to Opus 4.8 what Opus 5 will be to
Fable 5 in July ("quasi-Fable 5 for agentic use at half the price"):
the rapid democratization of the summit. The delay between the flagship
and its affordable variant is counted in weeks, not generations — a sign
that the internal distillation of capabilities (in the technical sense:
transferring performance from the large model to the small one) has
become a mastered industrial process.

The million-token context at $2/$10 completes Sonnet 5's role as the
year's workhorse: for high-volume uses — codebase indexing, massive
document processing, long-loop agents — the cost of long context becomes
the limiting factor, and Sonnet 5 makes it affordable where Fable 5 made
it prohibitive. The overall pricing logic becomes clearer: Fable 5
($10/$50) monetizes prestige and extreme tasks; Opus 4.8 then Opus 5
($5/$25) monetize everyday agentic performance; Sonnet 5 ($2/$10)
monetizes volume. Three tiers, a million tokens everywhere, and a price
scale that hugs the curve of real uses.

<a id="s03-14"></a>
### Opus 5: quasi-Fable 5 agentic at half price (July 24)

On July 24, 2026 comes Claude Opus 5: priced at $5/$25 per million
tokens — identical to Opus 4.8 — with an explicit promise in a single
formula: "quasi-Fable 5 for agentic use at half the price," and
thinking enabled by default. Every term of this product sheet is a
strategic declaration. "Quasi-Fable 5 for agentic use": the model does
not claim to equal the flagship on every axis — the cyber/bio/chem
capabilities that justify Fable 5's safeguards are presumably excluded
— but it equals it where most economic value is at stake: the autonomous
execution of complex tasks, code, research, tool manipulation. "At half
the price": $5/$25 is exactly half of $10/$50, Fable 5's price — the
arithmetic is displayed, not suggested.

"Thinking by default" is perhaps the year's most significant product
innovation: explicit reasoning, until then an advanced option the user
had to enable, becomes the standard behavior. This shift says that
reasoning is no longer a luxury for difficult tasks but the base
infrastructure of agentic use — and it prepares the ground, two months
later, for the September R&D announcements (26% of internal research
"led" by Claude, formalization of Fermat in 11 days), which all assume
models reasoning in depth by default. Opus 5 is thus the link between
June's raw prowess and September's scientific productivity: the model
that makes reasoning ordinary.

The release of Opus 5 on July 24, four days after the final approval of
the $1.5 billion settlement (July 20, Judge Martinez-Olguin), also
places the launch in a pacified legal context: the largest copyright
case in AI history has just been settled, and the lab can roll out its
product roadmap without the sword of Damocles of a trial. Anthropic's
2026 calendar decidedly reads like a score in which each movement —
raises, launches, regulatory crisis, settlement, scientific breakthroughs
— follows without pause.

<a id="s03-15"></a>
### Comparison: Opus 4.6, 4.7, 4.8, 5 — the trajectory of a lineup

Taken together, the four 2026 Opus iterations tell of an acceleration
then a stabilization. April: 4.6 (mid-April) then 4.7 (April 16), two
"temporary retakings of the leadership" days apart — the time of
contested leaderships, when each version reigns only until the next.
May: 4.8, 88.6% on SWE-bench Verified, $5/$25, 1M context — the time of
the quantified reference, when leadership expresses itself in benchmark
score rather than ephemeral crown, and when the model becomes the
standard to which the flagship will route its sensitive requests. July:
Opus 5, same $5/$25 price, "quasi-Fable 5 agentic at half price,"
thinking by default — the time of democratization, when yesterday's
summit becomes today's standard.

The price stability between 4.8 and 5 ($5/$25 in both cases) is the most
telling indicator: in two months, at constant price, the lab delivers a
"quasi-Fable 5" model where it delivered an 88.6% SWE-bench model —
meaning performance per dollar made a leap without the price moving.
And meanwhile, Sonnet 5 (June, $2/$10, 85.2% SWE-bench) occupies the
lower tier with only 3.4 points less than Opus 4.8 for a price divided by
2.5. The 2026 lineup thus verifies an empirical law: each price tier
sees its performance caught up by the tier below within weeks, and the
summit's premium is justified only on extreme tasks — or, in Fable 5's
case, on the sensitive capabilities that alone justify the safeguards'
existence.

The overall picture — Sonnet 5 at $2/$10, Opus 4.8 and Opus 5 at $5/$25,
Fable 5 at $10/$50, all at 1M context — sketches a lineup of rare
legibility in model history: three prices, one context, and one promise
per tier (volume, agentic, extreme). It is a lineup designed to be read
at a glance by a CTO — and it is also, let us not forget, the lineup
that carried the official run-rate beyond $47 billion in May, before
Opus 5's release.

