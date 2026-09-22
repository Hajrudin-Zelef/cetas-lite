---
id: briefing-general-tech-2026/03-servers-datacenters/10-ualink
title: "UALink: the open coalition against NVLink"
domain: servers-datacenters
role: deep-dive
task: interconnect
actors: ["AMD", "Amazon", "Apple", "Google", "Intel", "Meta", "Micron", "Microsoft", "Nvidia", "UALink"]
dates: ["2025-04", "2025-04-08", "2025-08", "2026-04-07", "2026-05-14", "2026-09"]
keywords: ["nvlink", "ualink", "accelerator", "asic", "benchmark", "blackwell", "capex", "chiplet", "custom silicon", "ethernet", "gpu", "gpus"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g04-10"
source_lines: [4082, 4228]
canonical_for: ["ualink"]
sha256: 387e9c9c4a6217ae8165f3cf2e8a03e81efc6a95d1f56001f99bb8f5f012b7d5
---

# UALink: the open coalition against NVLink

<a id="g04-10"></a>
### 4.10 UALink: the open coalition against NVLink

If power was the year's quiet revolution, interconnects were its loudest
political fight — and its most consequential unfinished business. **UALink**
(Ultra Accelerator Link) spent 2026 positioning itself as the open, multi-vendor
alternative to **NVLink**, Nvidia's proprietary scale-up fabric. The Register's
April 7, 2026 headline captured the framing with characteristic bluntness: the
**"No-Nvidia interconnect club."** Behind the quip is the deepest structural
question in AI infrastructure: *who owns the wires that let thousands of
accelerators act as one computer?*

A brief primer, because the distinction carries the whole section. AI scale-up
has two networking problems. *Scale-out* — connecting racks to each other across
the datacenter — is a solved, standardized, Ethernet-and-InfiniBand world with
multiple vendors. *Scale-up* — connecting accelerators *within* a rack or pod so
they share memory coherently and act as a single massive compute domain — is
where the physics gets brutal (terabytes per second of bandwidth, single-digit
microsecond latency) and where, in 2026, exactly one shipping solution existed:
NVLink, proprietary to Nvidia, available only on Nvidia's terms, on Nvidia's
roadmap, at Nvidia's margin. Every alternative accelerator — AMD's, Intel's,
every custom ASIC in 4.6 — faced the same ceiling: brilliant chips that could
not be wired together at datacenter scale without the one fabric they could not
license. UALink is the industry's collective answer to that ceiling.

**The coalition.** UALink's promoter members read like a roll call of everyone
with a strategic interest in not depending on Nvidia: **AMD, Intel, Google,
Microsoft, Meta, AWS and Alibaba** — plus **Apple, Astera Labs, Cisco, HPE and
Synopsys** among the promoters, with **Nvidia conspicuously absent**. The
absentee is the point: this is a consortium defined as much by who is not in the
room as by who is. Total membership stands at **85+ companies** — but the date
on that figure matters, and this dossier does not launder it: it is the
**official figure dated April 2025**, and **no official 2026 figure has been
published**. The coalition is large, it is visibly growing — the 2026
specification work drew in contributors across the stack — but the verified
count remains the April 2025 one, and we report it as such rather than inflating
it to sound current.

**The specification.** The technical base is real, published, and progressing on
schedule — which is itself an achievement for a multi-vendor consortium:

- **UALink 200G 1.0**, ratified **April 8, 2025**: **200 Gb/s per lane**, scale-up to **1,024 accelerators per pod**, PHY based on **IEEE P802.3dj** (the Ethernet physical-layer standard the industry converged on for 200G signaling). The specification was made available for **public download in August 2025** — open in the meaningful sense, not merely "available to members."
- **UALink 2.0**, published **April 7, 2026**: splits the **Data Link and Physical layers at 200G** (cleaner layering, easier multi-vendor interop), adds **In-Network Compute** (processing inside the network fabric itself — the direction large-scale AI collectives are heading), **confidential compute** support (encrypted, attested accelerator-to-accelerator traffic — the datacenter counterpart to the privacy architecture Apple built for PCC in 4.7), **multi-path** capabilities (resilience and load-balancing across fabric paths), a **UCIe 3.0-aligned chiplet specification** (letting UALink ride the industry's chiplet interconnect standard rather than inventing a parallel one), and a PHY path from **200G to 400G**.

On paper, this is a complete, modern scale-up fabric specification — and the 2.0
additions show the consortium designing for where AI infrastructure is *going*
(in-network compute, confidential fabrics, chiplets), not merely cloning where
NVLink *is*. As specification work, it is impressive by any standard.

**The reality check.** And here the story gets honest, because the same April 7,
2026 Register piece that announced 2.0 carried the subtext that defines UALink's
2026: **"UALink delivers 2.0 spec before v1.0 silicon ships."** As of 2026,
there were **no shipping UALink products**. Not one switch, not one accelerator,
not one deployed pod. The specification was running ahead of the silicon — a
consortium can ratify standards faster than its members can tape out chips,
qualify them, and build systems around them, and in 2026 that gap was the entire
story. A standard without silicon is a promise; UALink in 2026 was a well-
engineered, broadly backed promise.

The silicon is coming, and the milestones are verifiable, dated, and checkable —
which is why the promise deserves to be taken seriously:

- **AMD's MI400**, with **UALink 1.0** support, was **in qualification in mid-September 2026** — the first major accelerator designed for the open fabric, moving through the bring-up and validation phase that precedes volume shipment.
- **Astera Labs** was developing **Scorpio X-Series switches** with **up to 320 lanes** for UALink fabrics — the switching silicon a real pod-scale deployment requires, from the connectivity vendor that already ships at hyperscaler scale.
- **First deployments expected in 2027** (per Zacks, May 14, 2026), with **AWS/Trainium volume ramp in 2027** as the other anchor tenant — Amazon's custom silicon (4.6) doubling as the open fabric's first high-volume vehicle.

These are credible milestones toward 2027, not 2026 facts, and the distinction
is load-bearing. Qualification is not shipment; expected deployments are not
deployed pods; a volume ramp dated 2027 is a plan. The dossier records the plan
because the plan is real and the participants are serious — but it does not book
2027's revenue in 2026's column.

The competitive reality of 2026, stated as plainly as it can be: **NVLink
remained the only deployed scale-up fabric** — on existing Blackwell systems,
and with **NVLink 6 arriving on Rubin in H2 2026**, extending the lead even as
the challenger organized. Every large-scale AI training cluster operating
anywhere in the world in 2026 ran on Nvidia's interconnect. UALink in 2026 was a
specification, a coalition, and a roadmap — potentially a very important one,
but not yet a product anyone could buy or a pod anyone could benchmark.

**Why it matters anyway.** Interconnects are where platform lock-in runs
deepest, and understanding why explains the coalition's urgency. GPUs can be
second-sourced — AMD exists, custom ASICs are ramping (4.6). Memory is fungible
— a byte is a byte, whoever fabs it. But the fabric that lets a thousand
accelerators share memory coherently determines whose systems can actually
*scale*, and in 2026 that fabric was proprietary to a single vendor whose
roadmap, pricing and allocation terms the entire industry depended on. Breaking
that dependence is not a technical nice-to-have; it is the precondition for
every other diversification effort. A custom ASIC without a scale-up fabric is a
fast chip that cannot join a large cluster; an open fabric without custom ASICs
is a highway with no cars. UALink and the 4.6 silicon programs are two halves of
the same strategy, and 2027 is the year both halves have to arrive together.

UALink's significance, then, is not what shipped in 2026 but what it represents:
the first time the rest of the industry — hyperscalers, chipmakers, system
vendors, IP suppliers — has coordinated on a *credible* open alternative, with
specifications published on schedule, silicon in qualification, switch chips in
development, and deployments dated. If the 2027 ramps materialize — MI400
shipping, Scorpio switching, Trainium volumes on the fabric — the interconnect
monopoly breaks, and the industry's most durable moat gets its first real
breach. If they slip, Nvidia's moat gets another year deeper, NVLink 6 extends
the technical lead, and the coalition faces the familiar graveyard of open
standards that arrived late. Either way, the outcome will be decided in 2027 —
which is why this chapter ends here: with the coalition assembled, the spec
published, the silicon in the lab, and the most important race in infrastructure
hardware still to be run.

**NVLink vs. UALink at a glance (verified facts only).**

| Dimension | NVLink (Nvidia) | UALink (consortium) |
|---|---|---|
| Model | Proprietary, single vendor | Open, multi-vendor consortium |
| Status in 2026 | Only deployed scale-up fabric | Specification published; no shipping silicon |
| Current generation | Blackwell shipping; NVLink 6 on Rubin H2 2026 | 200G 1.0 (ratified 8/04/2025); 2.0 (7/04/2026) |
| Scale | Production pods in operation | Up to 1,024 accelerators/pod (spec) |
| Silicon | Shipping | MI400 in qualification (mid-Sep 2026); Astera Labs Scorpio switches in development |
| First open-fabric deployments | — | Expected 2027 (projection, not fact) |

The table is the chapter's competitive situation in miniature: a proprietary
incumbent that ships against an open challenger that specifies. Every
strategic question in AI infrastructure for 2027 — custom ASIC viability
(4.6), Apple's interconnect choice (4.8), hyperscaler negotiating leverage —
is a bet on which column gains or loses rows.

**Chapter synthesis: the physical stack, repriced.** The ten sections above
describe a single system under stress, and the stress had a single shape:
exponential demand meeting the lead times of the material world. The server
market doubled in value on the back of GPU systems whose prices were inflated
by memory (4.1–4.2); memory was short because fabs take years and wafers are
short into 2030 (4.3–4.4); buyers responded by committing $725–730 billion
and designing their own silicon (4.5–4.6); Apple ran its own parallel stack on
its own cadence (4.7–4.8); and the industry rebuilt power delivery and
interconnects for densities the old designs could not carry (4.9–4.10). No
section stands alone; each is the next section's cause or consequence.

Three threads carry into 2027, where this dossier's cutoff leaves them. *The
memory horizon*: whether "gradual improvement in 2028" (Micron) arrives on
schedule determines whether the cost-push regime of 2026 persists or breaks.
*The capex trajectory*: sustain, grow, or retrench — the single variable every
other forecast is downstream of. *The open challengers*: UALink silicon and
the custom-ASIC ramps either materialize in 2027 or they do not, and the
industry's structure for the rest of the decade depends on the answer. This
chapter recorded 2026's facts; 2027 will report the verdicts.

---

*End of Chapter 4 — Servers and datacenters. Chapter 5: [to be written].*
