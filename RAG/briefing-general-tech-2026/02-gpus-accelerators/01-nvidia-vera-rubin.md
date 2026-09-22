---
id: briefing-general-tech-2026/02-gpus-accelerators/01-nvidia-vera-rubin
title: "Nvidia: Vera Rubin in full production"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "CoreWeave", "Dell", "Huawei", "MLCommons", "Nebius", "Nvidia", "Qualcomm"]
dates: ["2026-01-06", "2026-03-16", "2026-06-01"]
keywords: ["vera rubin", "accelerator", "ascend", "blackwell", "cost per token", "gpu", "gpus", "mlperf", "mlperf v6.1", "rack-scale", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-1"
source_lines: [1442, 1508]
canonical_for: ["nvidia-vera-rubin"]
sha256: b4ae0ee9a4dd5a93f09ca750d58b4b94d487eaf0af2d930f7fb9337807a33a22
---

# Nvidia: Vera Rubin in full production

<a id="g03-1"></a>
### 3.1 Nvidia: Vera Rubin in full production

The Vera Rubin platform — Nvidia's sixth-generation AI platform, pairing Rubin GPUs with Vera CPUs — was the single most choreographed product rollout of 2026. Jensen Huang used his CES keynote on January 6, 2026 to make the headline declaration: **"Vera Rubin is in full production."** The phrasing was deliberate. It answered the market's only real question about Nvidia's 2026 — would the annual cadence hold? — with a three-word verdict, months before any independent observer could verify the claim.

The verification, such as it was, came in stages. The full technical reveal followed at GTC on March 16, 2026, where Nvidia laid out the platform in detail. Then came the first production evidence: on June 1, 2026, the first NVL72 rack entered production, deployed by CoreWeave with Dell hardware in Livingston. Volume shipments were scheduled for the second half of 2026 across eight partners — four hyperscalers and four neoclouds — making Rubin the broadest first-half launch of any Nvidia platform to date.

The cadence behind Rubin matters as much as Rubin itself. Huang's roadmap beyond the launch was explicit: an NVL144 configuration, then Rubin Ultra NVL576 in the second half of 2027, and the Feynman generation in 2028. That is an annual-to-18-month rhythm with no visible slowdown, and it is the tempo every competitor in this chapter is now forced to match or explain away.

#### GTC as a venue: why March mattered

Nvidia's GTC (GPU Technology Conference) is the company's annual platform event — the venue where architectures are revealed and roadmaps set. The March 16, 2026 edition carried an unusual burden: it had to convert January's three-word manufacturing claim ("full production") into a technical story developers and buyers could plan around. The cost-per-token claims (MoE training at 1/4 the GPUs; 10x throughput/watt at 1/10 cost per token — both vendor claims) were the conversion mechanism: they translated silicon into economics, giving buyers the ROI language for Rubin purchases months before volume shipments. GTC's function in the Rubin rollout was therefore commercial as much as technical — it was the event where the platform became *buyable* in the minds of customers, even before it was deliverable.

#### The neocloud phenomenon: 2026's structural shift

One of 2026's underappreciated stories is who deployed Rubin first — not the hyperscalers but the neoclouds. CoreWeave put the first NVL72 rack into production (June, Livingston, with Dell hardware), Nebius submitted Rubin results to MLPerf, and CoreWeave confirmed the first 504-GPU multi-rack cluster (September). The pattern has a structural explanation: neoclouds build greenfield AI datacenters where 190–230 kW liquid-cooled racks were designed in from the foundation, while hyperscalers retrofit general-purpose footprints. The result is a deployment-speed inversion — the smaller, newer clouds bring new platforms to market first, and the hyperscalers follow at volume. Nvidia's eight-partner H2 ramp (four and four) institutionalizes the pattern: the neocloud channel is no longer an experiment but half the launch strategy. For AMD, Huawei, and Qualcomm, the lesson is direct — your fastest path to deployed racks runs through buyers building new datacenters, not through those adapting old ones.

#### The staged rollout, in order

| Milestone | Date | Status |
|---|---|---|
| "Vera Rubin is in full production" (Huang, CES keynote) | January 6, 2026 | Declared |
| Full platform reveal (GTC) | March 16, 2026 | Presented |
| First NVL72 rack in production (CoreWeave/Dell, Livingston) | June 1, 2026 | In production |
| Volume shipments (8 partners: 4 hyperscalers + 4 neoclouds) | H2 2026 | Announced calendar |
| NVL144, Rubin Ultra NVL576 | H2 2027 | Announced roadmap |
| Feynman generation | 2028 | Announced roadmap |

#### January 6: why CES

Huang chose CES — the consumer-electronics show — for a datacenter manufacturing announcement, and the venue choice is part of the message. CES keynotes reach a general audience far beyond the datacenter trade press; declaring "full production" there made Rubin's status a mainstream fact rather than an industry rumor. It also set the year's narrative before any competitor could: by the time AMD's Advancing AI arrived in July, Nvidia's production claim had six months of incumbency. In 2026's compressed news cycle, *when* you announce is a competitive weapon, and January 6 was the earliest possible date to fire it.

#### "In production" vs. "shipping": reading the January claim

A note on reading these milestones, because the press routinely conflated them: "full production" in January and "volume shipments" in the second half are not contradictory. The January statement was a **manufacturing claim** — silicon and racks moving through the line. The second-half schedule was the **delivery calendar**. The distinction matters for two reasons.

First, it is the same staged language AMD used in July ("shipments starting at the end of the third quarter"), and comparing the two vendors' timelines requires comparing like with like: Nvidia's manufacturing claim came in January, AMD's delivery guidance in July. There is a real, six-plus-month gap in platform maturity between them, and it is the single most important fact about the AMD-vs-Nvidia race in 2026.

Second, the partner structure of the H2 ramp — **four hyperscalers and four neoclouds** — shows where Nvidia placed its delivery bets. The hyperscalers are the volume; the neoclouds (CoreWeave first among them, with the June Livingston deployment and the September 504-GPU cluster) are the speed. Neoclouds deploy faster because they are building dedicated AI datacenters rather than retrofitting general-purpose clouds — and Nvidia used them as the tip of the spear for every Rubin milestone that had a date attached.

#### Why the cadence is the product

It is worth stating explicitly what Huang's roadmap implies: Rubin is not a destination but a tick in a cadence. NVL144 doubles the domain; Rubin Ultra NVL576 (H2 2027) scales it by 8x; Feynman (2028) resets the architecture. For buyers, the message is that committing to Rubin in H2 2026 buys into a known upgrade path — the same lock-in logic that made the Hopper-to-Blackwell transition so sticky, now formalized as a public schedule. For competitors, the message is harsher: to compete with Nvidia in 2027, you must compete not with Rubin but with Rubin Ultra, a platform that exists today only as a name and a date. AMD's MI500 and Huawei's Ascend 970 are, in this sense, not responses to Rubin at all — they are responses to Rubin Ultra.

---

#### Nvidia's 2026, month by month

The Rubin rollout is easier to judge as a timeline than as a set of claims. Every dated Rubin event in the verification window, in order:

| Month | Event | Significance |
|---|---|---|
| January | "Full production" declared at CES (Jan 6) | Manufacturing claim; sets the year's tempo |
| March | Full platform reveal at GTC (Mar 16) | Technical disclosure; cost-per-token claims made |
| June | First NVL72 rack in production — CoreWeave/Dell, Livingston (Jun 1) | First dated production evidence |
| H2 | Volume shipments to 8 partners begin | Delivery calendar (announced) |
| September | MLPerf v6.1 "preview submission" (Sept 16) | First measured results |
| September | 504-GPU multi-rack cluster confirmed (Sept 16) | First beyond-single-rack evidence |

The shape of the year is a steady de-risking: claim (January) → disclosure (March) → production evidence (June) → measurement (September) → scale evidence (September). No step in the sequence contradicted an earlier one, which is itself informative — 2026's other accelerator stories (§3.5, §3.12) were still in the claim-and-disclosure phase at the same point in the year.

#### The eight partners: why the mix matters

Nvidia specified the H2 ramp's shape — four hyperscalers, four neoclouds — without naming names beyond the CoreWeave/Dell Livingston deployment. The mix is the strategy. Hyperscalers buy volume and integrate Rubin into general-purpose clouds; neoclouds buy speed, deploying dedicated AI datacenters where the power and cooling envelopes (§3.2) were designed in from the foundation. The verified record shows the neocloud channel delivering the dated milestones: CoreWeave's June rack, Nebius's MLPerf submission, CoreWeave's September 504-GPU cluster. In 2026, "neocloud" stopped being a euphemism for "small cloud" and became the industry's term for the fastest deployment path for power-hungry rack-scale systems — a structural shift in who brings new platforms to market first.

---

