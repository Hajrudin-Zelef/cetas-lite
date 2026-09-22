---
id: briefing-general-tech-2026/05-networking-optics/03-d7070-muxponder
title: "D7070 800G muxponder"
domain: networking-optics
role: deep-dive
task: networking
actors: ["FS.com", "Telxius"]
dates: ["2026-08", "2026-08-24"]
keywords: ["muxponder", "dci", "ethernet", "gpu", "wavelength"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-3"
source_lines: [5773, 5836]
sha256: 0a3ddbc2efdd634e1e59f347683ac7a0862fa5e22b1e5fc9d1e6d9d7f2ac5e28
---

# D7070 800G muxponder

<a id="g06-3"></a>
### 6.3 D7070 800G muxponder

On August 24, 2026 (via BusinessWire), FS.com confirmed the **D7070**, an
800G muxponder — a device that multiplexes multiple lower-rate client signals
onto coherent line-side wavelengths for transport. The headline figure is
**3.2 Tb/s of aggregated capacity**: the box aggregates traffic onto four
800G coherent line interfaces.

The muxponder sits in the "Scale-Across" half of the portfolio logic from
section 6.1: it is the box that takes the high-rate fabric traffic and puts it
on long-haul fiber between data-center sites. The timing — late August 2026,
just before the September release cluster — fits the broader pattern of vendors
completing their 800G coherent lineups ahead of ECOC 2026, where 1.6T would
steal the spotlight. The D7070 matters in this chapter mainly as a milestone
of completeness: FS's coherent portfolio now spans from pluggable
transceivers to box-level aggregation, which is what a vendor needs to claim
credibly before competing for the large AI-cluster DCI builds of 2027.

#### What a muxponder does in an AI DCI architecture

A muxponder is the aggregation point between the packet world and the
wavelength world. On the client side it takes multiple lower-rate signals —
in the D7070's case, enough 100G/400G client ports to fill 3.2 Tb/s of
aggregate capacity; on the line side it maps them onto four 800G coherent
wavelengths for transport. The arithmetic is the product definition:
**4 × 800G = 3.2 Tb/s**. In an AI-cluster DCI design, the box sits at the
data-center egress: the GPU fabric's east-west traffic, broken out at the
spine layer, is aggregated here and handed to the long-haul fiber plant as
coherent wavelengths — the "Scale-Across" domain of section 6.1 made into a
rack-mountable product.

#### Completing the coherent lineup before ECOC

The August 24 date is strategic, not accidental. ECOC 2026 opened on
September 21 with 1.6T dominating the agenda (sections 6.6–6.8); a vendor
arriving in Malaga without a complete 800G coherent story would have looked
like a 1.6T-only boutique. The D7070 fills the box-level gap above FS's
coherent pluggables: transceivers for the router faceplate, a muxponder for
aggregation, DCI reach up to 500 km. Whether the large AI-cluster DCI builds
of 2027 buy the integrated box or assemble their own from pluggables and
third-party muxponders is an open procurement question — but the product
now exists to be evaluated, which is the precondition for competing.

#### 800G coherent in 2026: the quiet standard

The D7070 is best understood against the year's quieter consensus: 800G
coherent became the default long-haul currency while 1.6T took the
headlines. Three independent data points in this chapter say the same
thing from different angles: FS.com's "Scale-Across" portfolio puts 400G
and 800G coherent at up to 500 km (6.1); Telxius standardized its
terrestrial rollout on 800G ZR/ZR+ pluggables across three continents
(6.12); and the Ethernet Alliance's ECOC demo includes 800G AI-lossless
features alongside the 1.6T showcase (6.8). None of these is a
record-breaking announcement — that is precisely the point. 800G coherent
in 2026 is infrastructure: bought, deployed, and demoed as a matter of
course. The D7070's 3.2 Tb/s of aggregated 800G capacity is a product
built for a market that already exists, not a bet on one that might. For
the reader tracking technology generations, the pattern is the familiar
one: the previous generation (800G) becomes the volume workhorse at the
exact moment the next generation (1.6T) starts shipping — and the vendors
that cover both, like FS with pluggables plus the D7070 muxponder, are
the ones positioned for the 2027 DCI builds.

