---
id: briefing-general-tech-2026/03-servers-datacenters/04-dram-nand-prices
title: "DRAM and NAND: prices, quotes, horizons"
domain: servers-datacenters
role: deep-dive
task: memory-crisis
actors: ["Gartner", "IDC", "Micron", "SK Hynix", "Samsung", "TrendForce"]
dates: ["2026-03", "2026-06", "2026-06-24", "2026-07", "2026-08", "2026-09", "2026-09-17"]
keywords: ["dram", "nand", "accelerator", "hbm", "hbm4", "packaging", "quantization", "wafer"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g04-4"
source_lines: [3337, 3498]
canonical_for: ["memory-crisis"]
sha256: f206ab3fd791eaa012dad8a0d8802e28b5a79b6e663760301f341559be60baf0
---

# DRAM and NAND: prices, quotes, horizons

<a id="g04-4"></a>
### 4.4 DRAM and NAND: prices, quotes, horizons

This section collects the price data and the on-the-record statements that
together define the memory crisis — the numbers first, then the voices, then the
question everyone in the industry was asking by September 2026: *when does it
end?*

**The price record**

The table below gathers the verified price measurements. They come from
different analysts, cover different segments and different windows, and should
be read as a mosaic, not a single series. The dossier's convention applies
throughout: each figure is scoped to its source, segment and date, and no figure
is stretched beyond them.

| Source / date | Segment | Measurement |
|---|---|---|
| TrendForce | Server DRAM contract prices | ×3.5 by September 2026 (+90–95% in Q1, +58–63% in Q2) |
| TrendForce | NAND contract prices | ~×3 |
| Gartner (Feb 2026) | DRAM + SSD basket | +130% → PCs +17%, smartphones +13%; shipments −10.4% / −8.4% |
| Gartner | Overall memory prices | ~200% of early-2026 levels; shortage at least into H1 2027 |
| CSIS, 17/09/2026 | DDR4 | +700–800% since early 2025 (+$575–765 per 128 GB) |
| ainvest, Aug 2026 | Mainstream DDR5 (consumer SKUs) | ×3.5–4 between late 2025 and Q1 2026 |
| Jefferies, late June 2026 | DRAM contract (projection) | +40–50% in Q3, then +30–40% in Q4 |

Two patterns stand out across the mosaic, and both are worth naming. First,
**the increases compound**. A 90–95% rise in Q1 followed by a 58–63% rise in Q2
does not add to roughly 150% — it multiplies to roughly ×3.5, because each
quarter's shortage reprices from the previous quarter's already-elevated level.
Compounding is what turns a bad year into a historic one: by September, buyers
were not paying triple the January price of a stable market, they were paying
triple the price of a market that had already nearly doubled. Anyone modeling
the crisis as a one-time step-up missed its shape entirely; it was a staircase,
and every quarter added a step.

Second, **the pain was broadest at the extremes of the stack**. The newest high-
end memory — HBM for accelerators, high-speed DDR5 for AI servers — was consumed
by datacenters at the 70% rate the Journal reported. The oldest mainstream
memory — DDR4 — spiked hardest in percentage terms as buyers cascaded down
looking for anything available (4.3). The middle of the market, mainstream DDR5
for PCs, roughly tripled to quadrupled depending on the window measured. There
was no segment in which a buyer could plausibly claim to have sat out the
crisis.

**The voices: what the memory makers said**

The most striking feature of the 2026 memory crisis is how openly the suppliers
described it — not as a temporary dislocation but as a structural regime, and
not with the usual corporate optimism but with timelines measured in years.
Three companies effectively control the DRAM market — Samsung, SK Hynix and
Micron — and their leadership, directly or through reported remarks, converged
on the same message: *relief is years away; plan accordingly.* When the parties
with every incentive to promise imminent relief instead tell customers to plan
for years of scarcity, that guidance is the single most important fact in this
section.

**Micron — CEO Sanjay Mehrotra, Q3 FY2026 earnings call, June 24, 2026.** The
key quote, worth reproducing in full for the care of its construction: *"supply
shortages… will take considerable time to improve, even as we expect industry
supply to improve gradually in 2028."* Note what the sentence does not say. It
does not say shortages end in 2028; it says supply improves *gradually* in 2028
— the beginning of easing, two years out, not resolution. A CEO choosing the
word "gradually" about a date two years distant is telling you, as plainly as
earnings-call language permits, that 2027 is spoken for. On the same call,
Micron disclosed the operational facts behind the language: its **HBM output for
2026 was fully sold**; it was meeting only **50–66% of demand** on key products
— turning away a third to a half of what customers wanted to buy; and **HBM3E
and HBM4 capacity was 100% booked through the end of 2027**, meaning its two
most advanced high-bandwidth memory generations were sold out eighteen months
forward. A memory maker rationing a third to half of demand, with its flagship
products allocated into the following year, is not experiencing a blip. It is
administering scarcity.

**SK Hynix — CEO Kwak, July 2026.** Kwak warned of the **worst shortage possible
in 2027**, with demand running **beyond the company's capacity into the end of
the decade**. If Micron's horizon was 2028, SK Hynix's was longer still: not a
cycle to be ridden out but a structural deficit between AI-driven memory demand
and the industry's ability to build fabs — a deficit measured against the end of
the *decade*, not the end of next year. Coming from the world's second-largest
DRAM maker and the leader in HBM, the statement carried the weight of someone
describing their own order book, not forecasting others'.

**SK Group — Chairman Chey Tae-won, March 2026 (via Reuters).** Chey pushed the
constraint one layer deeper into the supply chain: a **wafer shortage extending
into 2030**. Wafers — the blank silicon discs on which memory is fabricated —
are the feedstock of the entire semiconductor industry, and a wafer constraint
into 2030 means that even the fabs under construction face a ceiling on what
they can produce. It is the kind of statement that reframes every other horizon
in this section: if the raw material is short into 2030, then "gradual
improvement in 2028" describes finished-memory supply catching up only as far as
the wafer supply permits.

**TrendForce — analyst Wu.** The channel-side view matched the makers: **new
production capacity would have no notable effect before 2028**, given fab
construction and qualification lead times. When the analysts who track equipment
orders and fab ramps agree with the CEOs, the timeline is not posturing — it is
physics plus procurement.

**Dan Ives — via CNBC, August 2026.** The **~15:1 demand-to-supply ratio** cited
in 4.3: sentiment, not metrology, but eloquent sentiment.

**When does it end? The horizons, side by side**

| Source | Relief horizon stated |
|---|---|
| Gartner | Shortage at least into H1 2027 |
| Micron (Mehrotra, Jun 2026) | Gradual supply improvement in 2028 |
| TrendForce (Wu) | New production: no notable effect before 2028 |
| SK Hynix (Kwak, Jul 2026) | Demand beyond capacity into end of decade |
| SK Group (Chey, Mar 2026) | Wafer shortage into 2030 |
| IDC | Worst crisis in 15 years (severity assessment, not a horizon) |

The honest reading of the table is that **nobody with knowledge of the supply
side expected balance before 2028**, and the wafer constraint could push true
normalization into the 2030s. Gartner's "at least into H1 2027" is the most
optimistic entry and reads, against the others, as a lower bound on the
shortage's duration rather than a forecast of its end. The industry's planning
assumption for 2026–2027 was therefore not "when do prices come back down" but
"how do we build through the shortage" — which is exactly what the remaining
sections describe: spending anyway (4.5), designing around scarcity with custom
silicon (4.6), and rebuilding the physical stack to use every watt (4.9) and
every byte more efficiently.

There is a final, forward-looking implication worth stating explicitly because
it reframes the whole chapter. If memory remains the binding constraint into
2028 and possibly beyond, then the competitive advantage in AI infrastructure
shifts — at the margin — from *who can buy the most accelerators* to *who can do
the most with the memory they can get*. Architectural efficiency (memory
hierarchies, sparsity, quantization, custom ASICs co-designed around available
memory) becomes relatively more valuable than raw accelerator count. The
shortage, in other words, does not just tax the buildout; it changes what kind
of buildout wins. That is a 2027 story, but its premises were all established in
2026.
**HBM: the memory inside the crisis.** The quotes above name HBM3E and HBM4 —
high-bandwidth memory, the stacked DRAM that sits beside every AI accelerator
and feeds it at terabytes per second. A brief explainer for readers meeting
the acronym here: HBM stacks DRAM dies vertically and bonds them to the
accelerator package with thousands of microscopic interconnects, achieving
bandwidth an order of magnitude beyond conventional DDR at the cost of far
greater manufacturing complexity. Each generation — HBM3E (extended), then
HBM4 — raises bandwidth and capacity; each is harder to yield than the last.
That complexity is why HBM became the sharpest edge of the shortage: it is
the hardest memory to make, it goes exclusively into AI accelerators, and
demand for it is a direct function of accelerator shipments. When Micron says
HBM3E and HBM4 are 100% booked through end-2027, it is saying the most
technically demanding memory product in the world has an eighteen-month order
backlog — the backlog being, in effect, the AI industry's build schedule.

**What "100% booked" means operationally.** It is worth translating the
booking language into factory reality, because "fully booked" undersells it.
Memory sold eighteen months forward is not a waiting list; it is *allocated
production* — fab capacity, wafer starts, packaging and test slots reserved
against specific customers' forecasts, with contractual commitments on both
sides. For the buyer, it means supply certainty at whatever price was agreed —
valuable beyond measure in a ×3.5 market. For everyone without a booking, it
means the merchant market for leading-edge HBM effectively did not exist in
2026: there was nothing to buy at any price because everything coming off the
line already had a name on it. That is the mechanism behind the 15:1
sentiment in 4.3 — not a metaphorical tightness but a literal absence of
uncommitted supply.

