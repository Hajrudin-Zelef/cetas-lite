---
id: briefing-ia-2026/07-infrastructure-compute/05-amd-helios
title: "AMD Advancing AI: Instinct MI400 and the Helios racks"
domain: infrastructure-compute
role: deep-dive
task: infrastructure
actors: ["AMD", "Anthropic", "Meta", "Microsoft", "Nvidia", "OpenAI"]
dates: ["2026-07-23"]
keywords: ["amd", "helios", "agent", "agents", "benchmarks", "compute", "gpu", "gpus", "inference", "memory", "nvidia", "rubin"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s07-6"
source_lines: [8554, 8668]
canonical_for: ["amd-mi400"]
sha256: 36cab2dab82eb7b5b600844b40f3f65c1c8b1a027424927d0eeb5e65ea58f66b
---

# AMD Advancing AI: Instinct MI400 and the Helios racks

<a id="s07-6"></a>
### 7.5 AMD Advancing AI: Instinct MI400 and the Helios racks

On July 23, 2026, AMD held its Advancing AI event in San Francisco.
It was there that the foundry unveiled its most ambitious counteroffensive against Nvidia.
On the program: the Instinct MI400 family, including the MI455X, and the Helios racks.
AMD also presented EPYC "Venice," its sixth generation of server processors.
The event brought Meta executives on stage, a sign of the offensive's seriousness.
It was a show of force, intended to prove AMD is now a full-fledged competitor.

The Instinct MI400 family is the core of AMD's arsenal for 2026.
The MI455X is its spearhead, designed for the most demanding AI workloads.
AMD no longer just follows: it offers an architecture designed for the agent era.
The positioning is clear: large-scale inference, where the bulk of budgets now plays out.
By presenting a whole family rather than a single chip, AMD shows the depth of its roadmap.
It is a message to buyers: there will be a sequel, and it will be coherent.
A supplier's credibility is also measured by the visibility it gives on the future.

The Helios racks are AMD's answer to the dense rack format popularized by the Rubin NVL72.
Each rack packs 72 MI455X GPUs, a figure that speaks to buyers accustomed to this format.
Memory reaches 31 TB of unified HBM4, a considerable memory pool at rack scale.
Announced power is 2.9 exaflops in FP4, the reduced-precision format prized for inference.
These three figures sketch a machine designed to serve models continuously, at scale.
The rack has become the unit of sale for AI compute, and AMD now speaks this language fluently.
Helios is not a chip, it is a complete system, ready to deploy.

The networking choice is one of Helios's most interesting bets.
AMD opted for all-Ethernet networking based on open standards.
It is a head-on positioning against the proprietary approaches that dominate the market.
The argument is twofold: avoid vendor lock-in and benefit from a mature Ethernet ecosystem.
For data center operators, open standards mean more choice and better-controlled costs.
It is also a technological-sovereignty argument that speaks to many buyers.
By making Ethernet a banner, AMD turns a technical constraint into a commercial argument.

The most aggressive claim concerns token economics.
AMD claims to offer "30% more tokens per dollar" than the Nvidia Rubin NVL72.
It is a direct attack on the terrain where Nvidia is most vulnerable: cost of use.
One must however keep in mind that this is a seller's claim, according to available sources.
Commercial benchmarks must always be read with caution, as measurement conditions matter enormously.
But the very choice of metric is revealing: it is no longer raw power that sells, it is yield.
In 2026, the typical customer buys tokens, not teraflops.
And on this terrain, AMD believes it can beat the reference.

The delivery timeline is another key element of the announcement.
Helios rack shipments began at the end of the third quarter of 2026.
It is an important signal: AMD is not selling vaporware, but hardware going to customers.
In a market where delivery lead times make the difference, this punctuality is an argument.
It also lends credibility to the multi-year contracts signed with the big labs.
An announced but never-delivered rack is worth nothing; a rack going into production is worth gold.
The end of the third quarter of 2026 therefore marks the passage from words to deeds.

EPYC "Venice," AMD's sixth generation of server processors, completes the picture.
Its shipments also began at the end of the third quarter of 2026.
The CPU remains an essential component of AI infrastructure, for orchestration as for inference.
And agents demand significant CPU capacity for inference, as the September 21 session showed.
AMD thus covers both fronts, accelerators and processors, with a coherent offering.
It is this completeness that makes a supplier's strength against hyperscalers.
A player able to supply the entire rack has a decisive advantage in negotiations.

AMD's counteroffensive reads as a three-stage strategy.
First, a credible product in the market's format, the Helios rack.
Second, a punchy economic argument, tokens per dollar against the NVL72.
Finally, reference customers validating the bet, from OpenAI to Meta via Microsoft and Oracle.
Each level reinforces the others: without customers, the product remains a promise; without product, customers do not sign.
The July 23 event served to publicly assemble these three levels.
It is a well-oiled machine, which contrasts with the more scattered offensives of the past.

The price/performance positioning is the common thread of the whole strategy.
AMD does not claim to beat Nvidia on all terrains, but on the one that matters to buyers.
The "30% more tokens per dollar" is a slogan, but also a philosophy.
It says: at equal budget, you will serve more users, more agents, more requests.
In a world where inference runs continuously, this argument weighs heavily in investment committees.
It is an asymmetric attack against a leader settled on its margins.
And it is often thus that challengers end up moving the lines.

The question of execution remains, which will make or break this offensive.
Delivering racks at large scale is an industrial challenge as much as a technological one.
The gigawatt contracts signed with Anthropic and OpenAI will put the logistics chain to the test.
The end of the third quarter of 2026 is only a beginning: the pace must be held over several quarters.
The industry's history is full of brilliant announcements followed by disappointing deliveries.
AMD seems aware of this, and multiplies proofs of seriousness.
The coming quarters will tell whether Helios keeps its promises in the field.

The open-Ethernet bet is also an ecosystem bet.
By choosing open standards, AMD lets multiple equipment makers supply network components.
This plurality lowers prices and accelerates innovation through competition.
It also avoids dependence on a single supplier for interconnection, the neuralgic point of clusters.
For operators, it is the promise of an infrastructure where every brick is replaceable.
It is a philosophy that seduces buyers burned by past lock-ins.
And it is an argument Nvidia, with its proprietary approaches, will struggle to turn around.

The FP4 format, at the heart of the announced 2.9 exaflops, deserves an explanation.
It is a reduced-precision compute format, on 4 bits, suited to inference.
The idea is simple: to serve a model, one does not always need maximum precision.
By reducing precision, throughput multiplies and memory consumption divides.
It is a calculated compromise, which preserves perceived quality while exploding performance.
The fact that AMD communicates in FP4 shows where the battle plays out: inference, not training.
Yesterday's exaflops measured raw power; today's measure service efficiency.

The 31 TB of unified HBM4 tell another story, that of memory.
High-bandwidth memory is the limiting factor of large-model inference.
A model must fit in memory to be served quickly, and the biggest models are immense.
A unified 31 TB pool at rack scale can serve colossal models without splitting them.
The qualifier "unified" suggests memory seen as a single space by the system.
It is a major simplification for developers, who no longer have to manage fragmentation.
In token economics, memory is as strategic as pure compute.

The choice of the "72 GPU" format is not a technical accident, it is a commercial choice.
It is the format the market learned to buy with the dominant reference.
By adopting it, AMD speaks the buyers' language and eases comparisons.
A data center designed for one type of dense rack can host the other without rethinking everything.
This format compatibility reduces the cost of switching from one supplier to another.
It is a decision that says: we want to be interchangeable, not exotic.
And it is precisely what makes competition credible in operators' eyes.

