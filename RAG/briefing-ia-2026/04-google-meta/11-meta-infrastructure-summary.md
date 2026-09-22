---
id: briefing-ia-2026/04-google-meta/11-meta-infrastructure-summary
title: "Meta infrastructure: Graviton5 and the Vera Rubin pact, plus summary"
domain: google-meta
role: deep-dive
task: analysis
actors: ["AWS", "Google", "Irregular", "Meta", "Nvidia"]
dates: ["2026-06-10"]
keywords: ["rubin", "agent", "agents", "aws", "compute", "distribution", "gemini", "gemini 4", "incident", "inference", "multimodal", "muse"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s04-12"
source_lines: [5673, 5788]
sha256: eede8e034809b70e31951900f83bdf5b1d3a63323a664f52a581f1b55f728b8e
---

# Meta infrastructure: Graviton5 and the Vera Rubin pact, plus summary

<a id="s04-12"></a>
### The infrastructure: Graviton5 by the tens of millions of cores, and the Vera Rubin pact

On June 10, 2026, AWS moved Graviton5 to general availability:
etched at 3 nanometers, 192 cores per chip. It is the fifth
generation of the ARM processors Amazon designs for its own data
centers, and its characteristics — etching fineness, core
density — make it a machine built for massively parallel
workloads. The striking fact for this dossier is not the spec
sheet, though, but the customer: Meta is deploying tens of
millions of Graviton5 cores, alongside Uber and Snowflake. Tens
of millions of cores: the order of magnitude is dizzying, and
it says the essential — 2026's AI is first a question of
infrastructure, and Meta buys that infrastructure by the
container-load.

One must unpack what "tens of millions of cores" means. At 192
cores per chip, ten million cores represent on the order of
fifty thousand chips; tens of millions, therefore, several
hundred thousand chips — a figure corroborated, from another
angle, by the information reported by Bloomberg: hundreds of
thousands of rented chips. The convergence of the two
formulations is reassuring: we are indeed talking about a
deployment on the order of half a million chips, rented on the
AWS cloud rather than bought. It is a choice worth emphasizing:
Meta, which owns its own gigantic data centers, nevertheless
chooses to rent compute massively from AWS. Flexibility wins
over ownership — in a market where demand fluctuates and chip
generations follow one another fast, renting allows trimming the
sails without immobilizing billions.

The presence of Uber and Snowflake alongside Meta in the first
big deployments is not trivial either. It shows Graviton5 is
not a niche product for generative AI, but a generalist compute
platform adopted by companies with very different needs —
transport, data cloud, social networks, and AI. For AWS, it is
validation of its silicon strategy: designing its own chips to
reduce dependence on external suppliers and optimize its cloud's
performance-cost ratio. For Meta, it is assurance of abundant,
economical generalist compute capacity, complementing
specialized AI accelerators.

For the year's second infrastructure chapter is just as
structuring: Meta becomes an official NVIDIA NCP partner, with
Vera Rubin chips in its data centers. The NCP program — NVIDIA
Cloud Partner — designates NVIDIA's privileged partners for
deploying its infrastructure; entering it, for Meta, seals an
alliance at the top of the AI-compute food chain. And the chips
in question are no small matter: Vera Rubin is the generation
of AI accelerators succeeding Blackwell, the one on which
NVIDIA is building its empire's future. Installing them in its
data centers is, for Meta, securing access to the market's best
chips, direct from the manufacturer, with the status to match.

The combination of the two moves — massive Graviton5 rental
from AWS, NCP partnership with NVIDIA's Rubins in-house —
sketches a two-legged infrastructure strategy. The cloud leg,
flexible and generalist: tens of millions of CPU cores for
elastic workloads, services, classical compute. The proprietary
leg, specialized and sovereign: the market's best AI
accelerators in its own data centers for model training and
large-scale inference. It is a mature enterprise architecture,
not putting all its eggs in one basket and playing each
supplier to its strengths.

There is finally a geopolitical reading of these deployments.
Tens of millions of cores, hundreds of thousands of chips, a
partnership with NVIDIA on its most advanced generation: these
are figures placing Meta in the very closed circle of companies
able to mobilize compute at continental scale. In a world where
compute has become a strategic resource — coveted by states,
rationed by exports, scrutinized by regulators — belonging to
that circle is an asset in itself. It gives Meta weight in
industrial and political negotiations, and it constitutes a
barrier to entry for anyone wanting to compete: you can copy a
model, you cannot copy half a million chips.

In sum, 2026 was for Meta the year infrastructure came out of
the shadows. Long confined to technical press releases, it
became a front-page subject — because personal agents,
multimodal models, and billions of users have a physical cost,
measured in megawatts and chips. The Vera Rubin pact and the
tens of millions of Graviton5 cores are the two faces of one
conviction: in the age of agents, compute is the terrain, and
whoever controls the terrain controls the game.

<a id="s04-13"></a>
### In summary: two mirror strategies

Google and Meta spent 2026 answering the same question — how to
win the agent era — with opposite methods. Google chose
industrial continuity: a breakneck release cadence, constantly
falling prices, a billion-user distribution, and communicating
caution that teases without promising. Meta chose avowed
rupture: a brutal pivot to Muse Spark, a personal agent launched
with great fanfare, a return to open source, and massive
infrastructure bets. One optimizes the existing machine, the
other rebuilds it in flight.

The points of convergence are nonetheless numerous. Both
companies made agent efficiency their obsession — fewer tokens,
fewer tool calls, falling prices. Both understood voice is the
next battlefield — Live at Google, glasses in Meta's sights.
Both had their security incident — Irregular in May, revealed
in September — and drew opposite yet complementary lessons from
it: isolate evaluations, reserve sensitive capabilities for
verified actors. And both invested in infrastructure as never
before — because a billion users and omnipresent agents do not
run on promises.

The 2027 unknown remains: will Gemini 4 come out, and what will
it be worth against the Muse Spark lineage? Will LLaMA return
from its fallow period? Will personal agents become habits or
remain gadgets? This dossier stops at 2026's verified facts; the
answers belong to next year. But one thing is certain: neither
Google nor Meta intends to let the other win by forfeit.
