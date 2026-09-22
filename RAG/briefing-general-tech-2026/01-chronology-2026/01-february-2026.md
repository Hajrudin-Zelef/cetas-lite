---
id: briefing-general-tech-2026/01-chronology-2026/01-february-2026
title: "February 2026"
domain: chronology
role: timeline
task: chronology
actors: ["AMD", "Amazon", "Anthropic", "Apple", "Gartner", "Google", "Meta", "Micron", "Nvidia", "OpenAI", "SK Hynix", "TSMC", "Taalas"]
dates: ["2026-02"]
keywords: ["accelerator", "blackwell", "cowos", "custom silicon", "dram", "gpu", "hbm", "helios", "hyperscaler", "inference", "nand", "packaging"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g02-3"
source_lines: [608, 702]
sha256: 91e0cd28004137938fc70e2b83753989eba7f2efa9c730adc9d1918e863e49ef
---

# February 2026

<a id="g02-3"></a>
### 2.3 February 2026

February's news arrived as a set of apparently unrelated items that, in
retrospect, described the same pressure from four directions. Apple was
redesigning its inference infrastructure around its own silicon; Meta was
committing gigawatt-scale capital to AMD; a startup was shipping an inference
chip that deliberately avoided the two scarcest inputs in the industry (HBM and
CoWoS); and the memory market was publishing the numbers that explained why
everyone was doing all of this. The month's logic is clearer in hindsight:
inference was becoming the dominant workload, and every February move was a
hedge against the same memory bottleneck.

The Apple item came first, on February 17, via 9to5Mac: Apple was moving Private
Cloud Compute — its privacy-preserving cloud inference tier for Apple
Intelligence — to M5-based hardware, identified by the internal hardware
reference J226C, with iOS 26.4 as the software vehicle. The significance was
architectural as much as commercial. Private Cloud Compute had been built to
extend Apple's on-device privacy model into the cloud: requests that cannot be
handled on-device are routed to PCC nodes that Apple claims cannot retain or
inspect user data, with the whole design auditable by researchers. Moving that
tier to M5 meant the company was treating server inference as a first-generation
workload for its newest silicon rather than as a commodity task to rent from
someone else's datacenter.

Equally telling was what Apple was skipping: the M3 Ultra and M4 generations
were not being adopted for PCC. That skip is a quiet statement about cadence —
Apple would rather wait for the M5 generation than retrofit an inference cloud
to interim silicon — and it foreshadowed the pattern of the whole year: buyers
holding out for the next generation because each generation changed the
inference economics materially. The same logic would show up in AMD's July pitch
(buy Helios, skip the wait) and in Nvidia's Rubin ramp (the generation everyone
had been waiting for since Blackwell). February's PCC story was the first
instance of a 2026 motif: the generation gap as a strategic decision.

A week later, on February 24, Meta and AMD announced a deal that reset the scale
of custom-silicon commitments. The headline terms: a 6-gigawatt commitment, a
custom GPU built on AMD's MI450 architecture, and 160 million AMD warrants
issued to Meta. To put 6 GW in context, it is roughly the power budget of a
small country's datacenter fleet committed by a single customer to a single
accelerator family — and it was, at the time, the largest custom-silicon
arrangement AMD had ever signed. The warrant structure matters because it aligns
incentives beyond a purchase order: Meta gets bespoke silicon tuned to its
workloads, AMD gets a committed buyer large enough to justify the engineering
investment, and both sides share in the equity upside if the partnership
succeeds.

The deal also signaled that the frontier labs' custom-silicon playbook was
becoming the hyperscalers' playbook too. For years, "custom silicon" meant
Google's TPUs and the occasional Amazon Trainium deployment — in-house efforts
by companies with the scale to justify them. Meta–AMD was something different: a
hyperscaler co-designing with a merchant silicon vendor, at gigawatt scale, with
warrants. It was a template, and templates get copied: five months later, AMD
would sign the mirror-image deal with Anthropic (2 GW, July 22), and OpenAI's
Sachin Katti would commit to Helios at massive scale on the same July stage. The
two announcements together — February's and July's — marked 2026 as the year the
custom-accelerator market stopped being a sideshow and became the main event for
anyone buying compute at scale.

Between those two, on February 20, a smaller announcement illustrated the same
logic from the opposite end of the market. Taalas launched HC1, an inference
chip notable for what it did not use: no HBM, no CoWoS advanced packaging. That
is a deliberate design choice aimed squarely at the industry's two scarcest and
most expensive inputs — high-bandwidth memory, whose supply was being absorbed
by datacenter buyers, and TSMC's chip-on-wafer-on-substrate packaging capacity,
the advanced-packaging bottleneck that constrained every leading-edge
accelerator. The bet is that inference economics can be won without them, by
architectural efficiency rather than by stacking scarce components. Whether HC1
succeeds as a product is a separate question (and one this dossier cannot answer
from the February announcement alone); what matters for the narrative is that a
startup's entire value proposition in February 2026 was "inference without the
bottlenecks," which tells you how binding those bottlenecks had become. Taalas
was designing around the constraints; Meta was buying its way through them with
gigawatts and warrants. Same constraints, opposite strategies.

And then the numbers arrived. Gartner's February 2026 data put hard figures on
the memory squeeze: a combined +130% surge in DRAM and SSD pricing, feeding
through into PC prices up ~17% and smartphone prices up ~13%. Read carefully,
this is not a report about memory — it is a report about who pays for AI demand.
Datacenter buyers were absorbing HBM and high-end NAND capacity, and the
spillover was landing on consumer device prices: the laptop and the phone got
more expensive because the datacenter got hungrier. The +130% figure became the
reference point for the rest of the year's pricing story: SK Group's chairman
would warn in March of a wafer shortage stretching into 2030, Micron's CEO would
warn in June that shortages would take "considerable time" to improve, and Apple
would raise Mac, iPad and Vision Pro prices in June citing memory costs.
February was when the memory wall became a published fact; the rest of the year
was its consequences, quarter by quarter, price tag by price tag.

| Date | Event | Significance |
|------|-------|--------------|
| 17/02 | 9to5Mac: Apple PCC moving to M5 (J226C, iOS 26.4); M3 Ultra/M4 skipped | Inference cloud on latest-gen Apple silicon; the generation-skip motif |
| 20/02 | Taalas HC1 inference chip (no HBM, no CoWoS) | Architectural bet against the two scarcest inputs |
| 24/02 | Meta–AMD: 6 GW, custom MI450 GPU, 160M warrants | Largest custom-silicon deal of the year (to date); the template |
| Feb | Gartner: DRAM+SSD +130% → PCs +17%, smartphones +13% | Memory squeeze quantified; consumer prices moving |
