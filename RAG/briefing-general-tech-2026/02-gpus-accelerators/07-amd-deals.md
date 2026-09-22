---
id: briefing-general-tech-2026/02-gpus-accelerators/07-amd-deals
title: "AMD's deals: Anthropic, OpenAI, Meta"
domain: gpus-accelerators
role: deep-dive
task: funding-deals
actors: ["AMD", "Anthropic", "Meta", "Nvidia", "OCP", "OpenAI"]
dates: ["2025-10-06", "2026-02-24", "2026-04", "2026-07-22"]
keywords: ["accelerator", "custom silicon", "gpu", "helios", "mi455x", "nvlink", "valuation", "warrants"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-7"
source_lines: [1870, 1942]
canonical_for: ["amd-deals"]
sha256: 22adf24f6058a125b5eeb342c6bded83e6f3234faff84ebb8dc2ae484f380ef7
---

# AMD's deals: Anthropic, OpenAI, Meta

<a id="g03-7"></a>
### 3.7 AMD's deals: Anthropic, OpenAI, Meta

Hardware claims are cheap; purchase commitments are not. By the end of the verification window, AMD had converted its roadmap into three of the largest accelerator deals in industry history — one per frontier lab — and the structure of the three deals reveals how each lab values AMD's offering.

#### Anthropic — July 22, 2026

Announced on the first day of the Advancing AI event: **2 GW of MI450 Series** capacity, with the **first gigawatt online in the first half of 2027**. The financial structure is notable for what it is not: AMD's investment commitment is **up to $5 billion, milestone-conditioned** — a capped investment tied to deployment milestones, **not** the share-warrant structures of the other two deals. The distinction is material. Warrants give the lab equity upside in AMD; a capped milestone investment is closer to a co-build financing arrangement. Anthropic's deal is structured as a deployment partnership with AMD putting capital at risk against milestones — a sign of how badly AMD wants a second-source anchor tenant, and how carefully Anthropic is managing its own balance sheet.

The timing is also notable: announced the day before Su's keynote, the Anthropic deal was the event's opening proof point — a customer commitment presented before the product pitch, to frame everything that followed as demand-led rather than supply-pushed.

#### Why Katti took the stage

An OpenAI executive keynoting a supplier's launch event is unusual enough to merit explanation. Sachin Katti's July 23 appearance did three things at once: it gave AMD the strongest possible credibility signal (a frontier lab deploying your unshipped platform "at massive scale"), it gave OpenAI leverage over its primary supplier (Nvidia negotiates differently when a credible second source exists), and it disclosed the April pilot timeline — the earliest proof that Helios ran frontier workloads. The quid pro quo is visible in the deal structure: OpenAI holds 160M warrants, so promoting AMD is, in a literal sense, promoting the value of its own holdings. None of this makes the endorsement false — the pilots were real, the deployment guidance was specific — but it explains why the endorsement happened on AMD's stage rather than in a press release.

#### OpenAI — October 6, 2025

The oldest of the three, and the template: **6 GW over five years**, paired with **160 million share warrants**. The warrants are the tell — OpenAI gets equity exposure to AMD's upside, aligning the lab's incentives with AMD's stock performance. (The 2025 date places the agreement before the coverage window; it is included because its 2026 milestones — the Helios deployment ramp — are the story.)

By 2026 the deal's operational content arrived: OpenAI's **Sachin Katti** took the Advancing AI stage on July 23 and said the company would deploy Helios **"at massive scale starting towards the end of this year and then accelerating towards 2027"** — and disclosed that **GPT-class pilot workloads had been running on Helios since around April 2026**. That April pilot date is the earliest verified evidence of MI450-family silicon running a frontier lab's workloads, and it predates the public launch by three months. It is also the strongest single data point for Helios's software readiness: whatever the state of AMD's stack, it was good enough for OpenAI to run GPT-class pilots on it in Q2 2026.

#### Meta — February 24, 2026

**6 GW** of a **custom MI450-based GPU**, again with **160 million warrants**, and with Helios **co-developed through the Open Compute Project**. The custom-silicon element makes this the deepest of the three partnerships: Meta is not buying an off-the-shelf MI455X but a Meta-specific variant, developed in the open-hardware framework of OCP. Sources cited in coverage put the deal's implied valuation in the **$60–100 billion** range — **unofficial, per sources**, and should be treated as market chatter rather than a disclosed figure.

The OCP co-development detail is strategically significant beyond the dollar figures: it means Meta's variant is being designed in an open-hardware process, which lowers the barrier for other OCP members to adopt related designs — a potential ecosystem multiplier for AMD's rack architecture that a purely bilateral custom deal would not provide.

#### OCP co-development: why the venue matters

The Meta deal's "co-developed via OCP" detail is more than a procurement footnote. The Open Compute Project is the industry's open-hardware foundation: designs developed under it are published, reusable, and adoptable by other members. A custom Meta GPU developed through OCP — rather than behind closed doors — means the rack architecture, board designs, and potentially firmware interfaces become part of the commons. For AMD, that is an ecosystem multiplier: every OCP member that adopts related designs expands the Helios-compatible footprint without AMD spending a dollar. For Meta, it is leverage: open designs prevent single-vendor lock-in even while committing 6 GW to one vendor. The OCP path turns a bilateral deal into a potential standard — which is precisely why Nvidia, with its proprietary NVLink fabric, has historically kept its distance from it.

#### The three deals at a glance

| Deal | Date | Scale | Financial structure | Silicon |
|---|---|---|---|---|
| OpenAI | Oct 6, 2025 | 6 GW / 5 years | 160M share warrants | MI450 Series (Helios) |
| Meta | Feb 24, 2026 | 6 GW | 160M share warrants; custom GPU via OCP | Custom MI450-based |
| Anthropic | Jul 22, 2026 | 2 GW; first GW H1 2027 | Up to $5B milestone-conditioned investment (not warrants) | MI450 Series |

#### Warrants in one paragraph: the financial technology of the deals

A share warrant gives its holder the right to buy stock at a fixed price — it is equity upside without upfront equity purchase. In the OpenAI and Meta deals, 160 million AMD warrants each mean the labs profit directly if AMD's stock rises on the partnership's success: the better AMD executes, the more the warrants are worth. The mechanism aligns incentives (the lab wants its supplier to thrive) and compensates the lab for platform risk (betting a roadmap on unshipped silicon deserves upside). The Anthropic structure — up to $5 billion of milestone-conditioned AMD investment, no warrants — inverts the risk: AMD puts its own capital at risk against Anthropic's deployment milestones, paying for commitment rather than granting upside. Both are solutions to the cold-start problem of a challenger platform; the choice between them is a negotiation over who bears the risk of the ramp.

#### Reading the structures

The pattern across the three deals is a menu of alignment mechanisms: warrants (OpenAI, Meta) versus milestone-conditioned investment (Anthropic); off-the-shelf silicon (OpenAI, Anthropic) versus custom silicon via open hardware (Meta). The warrants deals give the labs equity upside — they profit if AMD's stock rises on the partnership's success. The Anthropic structure instead puts AMD's capital at risk against deployment milestones — AMD pays, in effect, for Anthropic's commitment to deploy. Both structures solve the same cold-start problem (why should a lab bet its roadmap on an unproven platform?), but they allocate the risk differently, and the difference reveals negotiating leverage: OpenAI and Meta extracted equity; Anthropic extracted co-investment.

Read together, the three deals are AMD's real 2026 product: not the MI455X's spec sheet but 14 GW of forward-committed demand from the three labs most capable of absorbing it. The open question — the one no verified source answers — is conversion: how much of the committed gigawattage becomes deployed, revenue-generating capacity on schedule. The first tests are OpenAI's "massive scale" ramp "towards the end of this year" and the H1 2027 Anthropic gigawatt. Until then, the deals are commitments, and the dossier records them as such.

---

#### The 14-gigawatt arithmetic

| Lab | Committed | Horizon | Running total |
|---|---|---|---|
| OpenAI | 6 GW | 5 years (from Oct 2025) | 6 GW |
| Meta | 6 GW | (custom GPU program) | 12 GW |
| Anthropic | 2 GW | First GW H1 2027 | 14 GW |

Fourteen gigawatts of forward-committed datacenter power is difficult to intuit, so a scale anchor helps: it is roughly the electricity consumption of several million households, committed to a single vendor's unshipped platform by three buyers. The figure's significance is not the wattage but the *forward* nature — these are not purchase orders for inventory but multi-year capacity commitments that underwrite AMD's manufacturing and financing plans. In that sense the deals function as AMD's balance sheet: the company can invest against 14 GW of contracted demand the way a utility invests against contracted load.

#### Deal timeline: eleven months, three labs

| Date | Deal | Structure |
|---|---|---|
| October 6, 2025 | OpenAI: 6 GW / 5 yrs | 160M warrants |
| February 24, 2026 | Meta: 6 GW custom | 160M warrants, OCP co-development |
| July 22, 2026 | Anthropic: 2 GW | Up to $5B milestone-conditioned (no warrants) |

The eleven-month span from first to third deal shows the template propagating: OpenAI set the warrant-based precedent in 2025, Meta extended it to custom silicon in February, and Anthropic — arriving last, with the most leverage as the most recent negotiator — extracted a different structure (AMD's capital at risk instead of equity upside). Each successive deal was struck with full knowledge of the previous ones, which means the variation in structures is negotiated, not accidental. For future reference: the next large AMD deal's structure will reveal whether the Anthropic model (co-investment) or the OpenAI/Meta model (warrants) has become the market standard.

---

