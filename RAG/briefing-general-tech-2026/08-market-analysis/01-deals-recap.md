---
id: briefing-general-tech-2026/08-market-analysis/01-deals-recap
title: "Deals and funding recap table"
domain: market-analysis
role: deep-dive
task: funding-deals
actors: ["AMD", "Anthropic", "Broadcom", "Meta", "Nvidia", "OCP", "OpenAI", "Positron", "Qualcomm"]
dates: ["2025-10", "2025-10-06", "2025-10-27", "2026-02", "2026-02-24", "2026-07", "2026-07-22", "2026-09", "2026-09-11", "2026-09-21"]
keywords: ["funding", "accelerator", "agentic", "ai200", "capex", "gpu", "helios", "hyperscaler", "inference", "lpddr", "lpddr5x", "mtia"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g09-1"
source_lines: [9158, 9240]
sha256: bcaf103bbed2cd6fbe6ac0de151000829807c53cffdb00927ff4e2c9897804f8
---

# Deals and funding recap table

<a id="g09-1"></a>
### 9.1 Deals and funding recap table

| Date | Deal | Parties | Terms (verified) |
|---|---|---|---|
| 06/10/2025 | AMD–OpenAI compute deal | AMD, OpenAI | 6 GW of AMD compute deployed over 5 years; 160 million share warrants to OpenAI |
| 27/10/2025 | Qualcomm AI200 launch | Qualcomm | Inference-class accelerator with 768 GB of LPDDR memory per card; first customer HUMAIN with a 200 MW commitment |
| 24/02/2026 | AMD–Meta compute deal | AMD, Meta | 6 GW of AMD compute; custom MI450-based GPU co-developed as part of the Helios platform, jointly via the Open Compute Project; 160 million share warrants to Meta |
| 22/07/2026 | AMD–Anthropic compute deal | AMD, Anthropic | 2 GW of MI450 Series systems; first gigawatt scheduled to come online in H1 2027; AMD equity investment of up to $5 billion, capped and conditioned on milestones |
| 11/09/2026 | Positron Series C | Positron | $875 million raised; inference-only chips built on LPDDR5X memory |
| 21/09/2026 | AMD crosses $1 trillion | Market event | +9.6% to $613.50 per share, ≈$1,005B market cap (Reuters); sector-wide AI rally |

Two context numbers frame the table without deserving rows of their own. First,
hyperscaler 2026 capital expenditure sat at roughly $725–730 billion for the year — the
demand pool into which every one of these deals draws. Second, Meta's Iris MTIA
accelerator reached mass production in September 2026, the quiet counterweight to the
loud AMD deals: the hyperscalers are not only buying outside compute, they are
manufacturing their own.

Read the three AMD deals together and a pattern emerges that is worth spelling out
because it is genuinely new in the semiconductor industry. October 2025 (OpenAI),
February 2026 (Meta), July 2026 (Anthropic): each deal pairs multi-gigawatt forward
commitments with large warrant grants — 160 million shares in two of the three — and
each lands further up the value stack. The OpenAI deal was a bulk supply agreement over
five years; the Meta deal added a custom MI450-class part co-developed through OCP; the
Anthropic deal folded in a direct, milestone-conditioned equity investment of up to $5
billion alongside 2 GW of Helios-class MI450 systems. The warrant structure is doing
double duty. It subsidizes the customer's cost of commit in a way that is politically
easier than a discount — OpenAI, Meta and Anthropic each get to participate in the
equity upside of the supplier whose volumes they guarantee — and it binds the customer
to AMD's roadmap for the better part of a decade. Nvidia never had to offer such
instruments when it held an uncontested position; their appearance is a signal of a
maturing, competitive supply market.

The Positron raise deserves separate attention because of what it validates rather than
its size alone. $875 million for a Series C is large but not remarkable in this cycle;
what is remarkable is the thesis it funds — inference-only accelerators built on
LPDDR5X, the memory class that the rest of this chapter will show is the bottleneck of
the industry. Positron is betting that the dominant workload of the late 2020s is
agentic inference with huge context windows, and that the winning architecture is
memory-bandwidth-per-dollar rather than peak FP16 throughput. Qualcomm's AI200 said the
same thing a year earlier with 768 GB of LPDDR on a single card and a 200 MW first
customer in HUMAIN. Two different companies, two different continents, one shared
reading of the memory constraint.

#### The warrant pattern and the capex behind it

Step back from the individual rows and the AMD sequence of 2025–2026 reads as a
single commercial invention, repeated three times with increasing ambition. The
October 2025 OpenAI deal established the template: multi-gigawatt forward
commitment plus a large warrant grant — 160 million shares — that turns the
customer into a quasi-shareholder. The February 2026 Meta deal kept the warrant
count identical but moved the product up the stack: not bulk accelerators, a
custom MI450-class part co-developed through the Open Compute Project, embedded
in Meta's own Helios platform. The July 2026 Anthropic deal kept the warrant
logic, added a direct AMD equity investment of up to $5 billion — capped,
conditioned on milestones — and attached it to 2 GW of Helios-class MI450
systems with the first gigawatt due online in H1 2027. Three deals, one grammar:
commit the gigawatts, share the equity upside, co-develop the silicon.

The grammar exists because the demand pool is large enough to underwrite it.
Hyperscaler capital expenditure for 2026 sat at roughly $725–730 billion for
the year — a number so large it reframes every row in the table. The AMD deals
are not bets on a single customer's solvency; they are slices of a
three-quarter-trillion-dollar build-out, which is why warrant structures could
be offered without the supplier taking existential customer-concentration risk.
And the quiet counterweight to all of it is Meta's Iris MTIA reaching mass
production in September 2026: the same quarter Meta signed for 6 GW of AMD
compute, it was also manufacturing its own inference silicon at scale. The
hyperscalers are not choosing between buying and building. They are doing both,
hedging the merchant market against the in-house roadmap, and the deals in the
table are the merchant-market leg of that hedge.

One more pattern worth naming: the sequencing of the announcements tracks the
industry's migration from training to inference. The OpenAI deal of October
2025 was still legible as training-adjacent bulk compute. By the Meta and
Anthropic deals of 2026, the language is inference: custom parts, Helios racks,
gigawatts scheduled to come online in H1 2027 for serving. The buyers were
securing the machines that would serve the agentic workloads of section 9.4,
and the memory economics of section 9.2 were already shaping what those
machines look like. The table is a compute story on its face; underneath, it is
a memory story and an inference story.

