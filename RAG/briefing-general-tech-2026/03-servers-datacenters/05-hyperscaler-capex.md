---
id: briefing-general-tech-2026/03-servers-datacenters/05-hyperscaler-capex
title: "Hyperscaler capex: ~$725–730 billion in 2026"
domain: servers-datacenters
role: deep-dive
task: funding-deals
actors: ["Amazon", "Google", "Meta", "Micron", "Microsoft", "Nvidia"]
dates: ["2026-09"]
keywords: ["capex", "hyperscaler", "800v dc", "accelerator", "dram", "ethernet", "gpu", "memory shortage", "nvlink", "run-rate"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g04-5"
source_lines: [3499, 3630]
canonical_for: ["hyperscaler-capex"]
sha256: 47ece66cbe475f3c2fcd5e9f7a9dfb07465dafb03881fb87b8dd2bb300c6729e
---

# Hyperscaler capex: ~$725–730 billion in 2026

<a id="g04-5"></a>
### 4.5 Hyperscaler capex: ~$725–730 billion in 2026

Against the backdrop of record server prices and a historic memory shortage —
against every signal that building datacenters had never been more expensive —
the four American hyperscalers did not slow down. They accelerated. Combined
capital expenditure for **Amazon, Alphabet, Meta and Microsoft** in 2026 is
estimated at approximately **$725–730 billion**, against roughly **$410 billion
in 2025**. That is a year-over-year increase of nearly 80%, and a figure that
would, on its own, rank among the largest private investment programs ever
undertaken in a single year — comparable to the GDP of a mid-sized country,
deployed as concrete, steel, silicon and copper.

The individual commitments, as verified:

| Hyperscaler | 2026 capex (approx.) |
|---|---|
| Amazon | ~$220B |
| Alphabet | $180–190B |
| Microsoft | ~$190B |
| Meta | $125–145B |
| **Big 4 total** | **~$725–730B** (vs ~$410B in 2025) |

Three things are worth saying about these numbers, because each is easy to misread.

First, **they are dominated by AI infrastructure but are not pure AI numbers**.
Each company's capex includes non-AI items — fulfillment centers and logistics
for Amazon, offices and undersea cables across the board, non-AI cloud capacity.
The verified figures do not split AI from non-AI capex, and this chapter does
not invent a split. The ~$725–730B should be read as the *envelope within which
the AI buildout happens*, not as a meter reading on AI alone. That said, the
year-over-year *increase* — roughly $315–320 billion — is overwhelmingly
attributable to AI datacenter investment, because nothing else in these
companies' capital plans grew at anything approaching that rate. The delta is
the AI story; the total is the envelope.

Second, **the growth rate is the real story, not the level**. From ~$410B to
~$725–730B in one year is not incremental capacity planning or prudent
provisioning. It is a strategic bet, taken independently by four companies
within months of each other, that compute demand will absorb everything built —
a coordination without coordination that tells you how uniformly the industry's
largest buyers read the demand signals. Each of the four had its own forecasting
apparatus, its own board, its own shareholders to answer to; each reached the
same conclusion: *build, at a pace that would have seemed delusional three years
ago.* When four independent actors with four independent datasets converge on
the same extreme, the convergence itself is information.

Third, **this spending is the demand-side mirror of the memory crisis** — and
seeing the two together is essential. When four buyers commit three-quarters of
a trillion dollars to infrastructure in a single year, and a large share of that
flows into memory-hungry accelerator systems, the Wall Street Journal's figure —
AI datacenters absorbing 70% of high-end DRAM in 2026 — stops looking surprising
and starts looking mechanical. The capex *is* the memory demand. The shortage
*is* the capex, priced through the only channel available. Sections 4.3 and 4.4
describe what the suppliers saw; this section describes what the buyers did.
They are the same phenomenon viewed from opposite ends of the purchase order.

The spending also reframes the server-market record in 4.1 in an important way.
The $166.3B quarterly server market is, to a significant degree, these four
companies' procurement budgets flowing through ODMs, OEMs and direct-
manufacturer relationships. When the buyers are this concentrated and this
committed, the "market" behaves less like a textbook market and more like a
small number of very large bilateral relationships — which is why *allocation*,
not price discovery, became the industry's central activity in 2026. Micron
meeting only 50–66% of key demand (4.4) is what allocation looks like from the
supplier's side; $725–730B of capex chasing that supply is what it looks like
from the buyers' side. The price mechanism still operated — triple-digit memory
inflation proves it — but at these volumes, the binding activity was securing
supply at all, at whatever price cleared.

There is a distributional detail worth noting in the table. Amazon's ~$220B
leads, which surprises no one familiar with AWS's scale and the company's
historical willingness to outspend. Alphabet's $180–190B and Microsoft's ~$190B
are effectively tied — the two companies' cloud businesses (Google Cloud, Azure)
growing into capex footprints that rival Amazon's. Meta's $125–145B is the
smallest of the four in absolute terms but arguably the most aggressive relative
to revenue: Meta has no cloud business to monetize the capacity externally, so
its capex is a pure bet on its own AI workloads — advertising, recommendations,
and the open-model ecosystem — absorbing the spend. Four different business
models, one identical conclusion.

The open question — unanswerable from the verified facts, but unavoidable, and
the most consequential unknown in technology infrastructure — is **duration**.
Capex at this level is predicated on AI revenues, eventually, justifying it:
cloud AI services, advertising uplift, enterprise contracts, consumer
subscriptions. Through September 2026, all four companies' statements supported
continuation, and none revised downward. But ~$725–730B is not a run-rate any
company sustains casually; it is roughly double the 2025 level, and doubling
again in 2027 would take the Big 4 past $1.4 trillion — a number at which even
the hyperscalers' balance sheets start to feel the weight. Whether 2027
sustains, grows, or retrenches from the 2026 level determines everything
downstream: memory demand (4.3–4.4), server volumes (4.1–4.2), the pace of
custom-silicon ramps (4.6), and the urgency of the power and interconnect
transitions (4.9–4.10). Every forecast in this chapter is downstream of that
single unknown. The dossier records the 2026 fact; the 2027 answer belongs to a
future edition.

**What the capex buys — the physical anatomy of $725 billion.** The verified
figures do not split the Big 4's capex into line items, but the industry's
physical requirements make the rough anatomy uncontroversial as context (not
as reported fact). The largest share flows into *accelerators and servers* —
the GPU systems of 4.2 at ~$170,200 ASP, multiplied by millions of units. The
second share is *construction*: datacenter shells, increasingly built
speculatively ahead of demand, on timelines where a large campus takes
eighteen to thirty months from groundbreaking to energized. The third is
*power infrastructure*: substations, generators, and increasingly the
purpose-built electrical plants that feed AI campuses — the demand side of
the 800V DC transition in 4.9. The fourth is *networking*, both scale-out
(Ethernet/InfiniBand fabrics) and the NVLink-scale-up domains of 4.10. And a
growing fifth is *land and power contracts* themselves — options on future
capacity, because in 2026 the binding constraint on a new datacenter was often
not capital but megawatts: utilities' interconnection queues stretched for
years in the hottest markets. Three-quarters of a trillion dollars, in other
words, buys the entire physical stack this chapter describes — which is why
every section in it is downstream of this one.

**The 2027 question — asked, not answered.** The dossier's discipline requires
stating plainly what is not known. The verified record contains 2026 capex
commitments and no verified 2027 figures; anything written here about 2027 is
framing, not fact. The framing that matters: ~$725–730B is roughly double the
2025 level, and the three trajectories from here — *sustain* (another ~$730B
year), *grow* (toward or past $1T), *retrench* (a pullback as digestion sets
in) — imply radically different worlds for memory demand, server volumes, and
every supplier in this chapter. History offers precedents for all three:
telecom capex retrenched violently after the 2001 overbuild; cloud capex
sustained and grew through every cycle since 2010. Which precedent 2027
follows depends on a variable this chapter cannot measure — whether AI
revenues begin visibly justifying AI capex. Through September 2026, the
hyperscalers' behavior (no downward revisions, continued pre-commitments like
the 14 GW Iris target in 4.6) indicated they believed the answer was yes. The
dossier records the bet; the payoff belongs to a future edition.

