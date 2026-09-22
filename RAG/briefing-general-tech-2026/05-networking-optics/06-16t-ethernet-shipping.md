---
id: briefing-general-tech-2026/05-networking-optics/06-16t-ethernet-shipping
title: "1.6T Ethernet is shipping now"
domain: networking-optics
role: deep-dive
task: networking
actors: ["China", "Credo", "Google", "MACOM", "Meta", "Nvidia", "OIF"]
dates: ["2026-09", "2026-09-07", "2026-09-12", "2026-09-17"]
keywords: ["ethernet", "cpo", "gpu", "hyperscaler", "lpo", "npo", "optics", "tpu", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-6"
source_lines: [5945, 6022]
canonical_for: ["16t-ethernet"]
sha256: 75c8f8408d02be3fcf75eed8bf8306b7f90cf60f52de57d66db275861826c18a
---

# 1.6T Ethernet is shipping now

<a id="g06-6"></a>
### 6.6 1.6T Ethernet is shipping now

The single most important networking fact of 2026: **1.6-terabit Ethernet
crossed from sampling into shipment.** The formulation comes from
LightCounting, reported via EETimes around September 16–17, 2026 — "1.6T is
shipping now" — and it is worth unpacking who is shipping and to whom.
Nvidia and Google were already integrating 1.6T into their infrastructure in
2026; Meta and Oracle were expected to follow in 2027. The sequence maps
neatly onto the hyperscaler AI-cluster roadmaps: the builders with the largest
frontier-training footprints move first, the rest follow a year later.

The financial side confirmed the inflection on September 7, 2026, when
Goldman Sachs **raised its 1.6T shipment forecasts** — an analyst revision
that followed the vendor reports of accelerating order books rather than
preceding them. And at CIOE China on September 12, 2026, the supply-chain
signal was even blunter: 1.6T was described as **ramping to volume, "sold
out," with capacity being reserved** — while 3.2T/6.4T generations in NPO
(near-packaged optics) and CPO (co-packaged optics) form factors were already
in sampling. The phrase "sold out, capacity being reserved" is the language of
a seller's market: allocation, not price, is the binding constraint.

| Milestone | Date | Source |
|---|---|---|
| Goldman Sachs raises 1.6T shipment forecasts | 7/09/2026 | Goldman Sachs |
| CIOE China: 1.6T ramping to volume, "sold out, capacity being reserved"; 3.2T/6.4T NPO/CPO in sampling | 12/09/2026 | CIOE China reports |
| LightCounting: "1.6T is shipping now" (Nvidia, Google integrating; Meta, Oracle 2027) | ~16-17/09/2026 | EETimes via LightCounting |

The 2026 inflection has a structural consequence the reader should keep in
mind for everything that follows in this chapter: when a generation moves to
volume, the ecosystem pivots from proving the physics to proving
interoperability and scaling supply. The OIF 1600ZR agreement (6.7), the
Ethernet Alliance demo (6.8), the LPO acceleration (6.9) and Credo's ZeroFlap
ramp (6.11) are all downstream of this moment.

#### The staggered adoption: why Nvidia and Google first

The LightCounting sequencing — Nvidia and Google integrating in 2026, Meta
and Oracle in 2027 — is not a ranking of enthusiasm but a map of cluster
architecture. The first movers are the builders running the largest
frontier-training fabrics, where the bandwidth density of 1.6T pays off
immediately: Nvidia's own GPU clusters (the Quantum-X800 / B300-GB300
fabrics from section 6.2) and Google's TPU pods. Meta and Oracle, with
large but differently-architected fleets, follow a year later — the normal
one-generation lag of a technology transition, and the same lag that
separated early 800G adopters from the mainstream in 2023–2024. For the
supply chain, the stagger is good news: it stretches the ramp over two
years instead of concentrating it in one, which is exactly what a
capacity-constrained market (see below) needs.

#### "Sold out, capacity being reserved": the seller's market

The CIOE China phrasing from September 12 deserves to be quoted precisely
because it is the language vendors use when demand exceeds supply:
**1.6T ramping to volume, "sold out," capacity being reserved.** In a
seller's market the binding constraint is allocation, not price — buyers
reserve fab and assembly capacity months ahead, and the negotiation is over
delivery slots. This is the context in which to read every September 2026
vendor announcement in this chapter: FS's portfolio release, Credo's
ZeroFlap launch, MACOM's 448G-per-lane demos are all, among other things,
bids for a share of reserved capacity in 2027. It is also the context for
the forward-looking part of the CIOE signal: **3.2T/6.4T in NPO and CPO
form factors already in sampling** — the supply chain is sampling the next
generation while struggling to deliver the current one, the classic
overlapping-generations pattern of optical networking.

#### The analyst layer: Goldman Sachs raises the forecast

Goldman Sachs raising its 1.6T shipment forecasts on September 7 is a
second-order signal — an analyst revision, not a shipment — but a
meaningful one. Sell-side forecast raises in a ramp year typically follow
vendor order-book disclosures and channel checks; they lag the physical
reality by weeks, and their function is to reset investor expectations
rather than to discover demand. The dossier includes it for that reason:
not as proof of shipments (the LightCounting and CIOE signals are the
proof), but as evidence that the financial narrative had caught up with the
physical one by early September 2026.

