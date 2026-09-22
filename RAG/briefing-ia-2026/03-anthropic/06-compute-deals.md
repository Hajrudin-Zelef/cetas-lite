---
id: briefing-ia-2026/03-anthropic/06-compute-deals
title: "Anthropic compute deals: AWS, AMD, architecture comparison"
domain: anthropic
role: deep-dive
task: funding-deals
actors: ["AMD", "AWS", "Anthropic", "Nvidia"]
dates: ["2026-07-22"]
keywords: ["amd", "aws", "compute", "alignment", "claude", "datacenter", "forward-commitment", "gpus", "graviton", "helios", "inference", "nvidia"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s03-19"
source_lines: [4020, 4142]
canonical_for: ["anthropic-compute-deals"]
sha256: 3446b0bba7bb605fba4f20bbd155a0a04b8591676398dedd7ba0c79d274aa3d9
---

# Anthropic compute deals: AWS, AMD, architecture comparison

<a id="s03-19"></a>
### The AWS deal (April 20–21): the biggest cloud forward-commitment by a frontier lab

On April 20–21, 2026, Anthropic and Amazon announce an agreement of
unprecedented scale: Anthropic commits to spending more than $100
billion on AWS over ten years — 5 gigawatts of capacity, with a
Trainium2-to-Trainium4 chip trajectory, complemented by Graviton — and
makes AWS its primary cloud; in return, Amazon invests $5 billion
immediately, plus up to $20 billion conditional, after about $8 billion
already invested. Each clause of this agreement deserves unpacking,
because it redefines the relationship between frontier labs and
hyperscalers.

On Anthropic's side, the commitment of more than $100 billion over ten
years is, per the verified facts, "the biggest cloud forward-commitment
by a frontier lab" — the phrase is explicit and worth underscoring: no
AI lab had ever committed such future amounts to a cloud provider. The
5 gigawatts give the physical scale: it is the electricity consumption
of a small nation, dedicated to training and inference of Claude models
for the coming decade. The mention "Trainium2 → Trainium4" reveals the
nature of the technological bet: Anthropic commits not to generic
Nvidia GPUs, but to Amazon's proprietary silicon roadmap — Trainium2
today, Trainium4 tomorrow — complemented by Graviton CPUs. It is an
architecture bet: the lab ties its compute destiny to Amazon's ability
to deliver its chip roadmap against Nvidia.

On Amazon's side, the investment reads in three layers: about $8
billion already invested (the relationship's history), $5 billion
immediate (the new agreement's entry ticket), and up to $20 billion
conditional (long-term alignment of interests, whose exact conditions do
not appear in the verified facts). At the potential total — $33
billion — Amazon becomes one of Anthropic's very first economic
shareholders, and Anthropic becomes Amazon's most strategic cloud
customer. The symmetry is perfect: one brings compute and capital, the
other brings guaranteed demand and the models.

The designation of AWS as "primary cloud" is not boilerplate: it means
the training of future generations and a growing share of inference
will transit through Amazon's infrastructure, with the
operational-sovereignty implications that entails. And the timing —
April 20–21, between Opus 4.6 (mid-April) and Opus 4.7 (April 16) — is
not neutral: the agreement is signed at the precise moment Anthropic
demonstrates, models in hand, that it deserves a commitment of this
size. Proven capability (April's temporary leaderships) buys future
capability (the 5 GW).

<a id="s03-20"></a>
### The AMD deal (July 22): up to 2 GW of MI450 in Helios racks

On July 22, 2026, Anthropic announces a partnership with AMD of a
different nature but comparable ambition: up to 2 gigawatts of
MI450-series chips, deployed in Helios racks, with the first gigawatt
expected as early as the first half of 2027; AMD invests up to $5
billion in Anthropic; and Claude will be used to optimize ROCm, AMD's
software stack. Three parts, three logics.

The capacity part first: 2 GW of MI450 in Helios racks is a second
silicon pillar for Anthropic, after AWS's Trainium pillar. The timeline
— first gigawatt by H1 2027, less than a year after the announcement —
indicates aggressive deployment, and the choice of Helios racks (AMD's
system platform for large-scale AI) signals integration at the
datacenter level, not a simple chip purchase. By securing 2 GW at AMD
on top of the 5 GW AWS, Anthropic brings its capacity trajectory to 7
GW — and above all, it gives itself two independent silicon sources,
reducing its dependence on a single provider's roadmap.

The financial part next: AMD invests up to $5 billion in Anthropic —
the same order of magnitude as Amazon's immediate ticket ($5 billion),
but at a chipmaker rather than a hyperscaler. The symmetry of amounts
is no coincidence: it signals that AMD pays, like Amazon, the entry
price to Anthropic's strategic-partner circle — with the difference that
AMD also buys, in passing, a reference customer for its MI450 range
against Nvidia's dominance.

The software part finally — "Claude used to optimize ROCm" — is perhaps
the most revealing of the era: the AI model becomes the optimization
tool for the software stack it runs on. The loop is closed: Claude
improves ROCm, ROCm accelerates the MI450s, the MI450s train future
Claudes. This is the first documented instance, in this dossier's
verified facts, of model-infrastructure co-optimization driven by the
model itself — and it prefigures the September announcements on the 26%
of internal R&D "led" by Claude, of which infrastructure optimization
is one of the most tangible forms.

<a id="s03-21"></a>
### AWS vs AMD: two partnership architectures compared

The AWS (April) and AMD (July) deals compare term by term, and their
comparison illuminates Anthropic's strategy better than each agreement
taken in isolation. On infrastructure amounts: AWS means more than $100
billion in cloud spending over ten years and 5 GW; AMD means up to 2 GW
of MI450 with no spending commitment quantified in the verified facts —
the asymmetry is massive, and it tells the hierarchy: AWS is the
foundation, AMD is the diversification. On capital investments: Amazon
puts $5 billion immediate plus up to $20 billion conditional (after
~$8 billion historic); AMD puts up to $5 billion — again, the
hyperscaler pays the high price of primary-cloud status, the chipmaker
pays the entry ticket of the second pillar.

On the technological nature, the difference runs deeper: the AWS
agreement binds Anthropic to a proprietary roadmap (Trainium2 →
Trainium4 + Graviton) within a cloud operated by Amazon — it is a
vertical-integration bet, where the lab relies on its partner for
silicon, networking, and operations. The AMD agreement, for its part,
covers chips (MI450) deployed in Helios racks — it is an
open-architecture bet, where Anthropic keeps its hand on system
integration and gains negotiating leverage against AWS as against
Nvidia. One buys ten years of operational peace of mind; the other buys
the strategic option and competition between providers.

On timing, finally: the AWS agreement is signed in April, at the heart
of Opus 4.6/4.7's temporary leaderships — it monetizes demonstrated
capability. The AMD agreement is signed July 22, two days after the
final approval of the $1.5 billion settlement (July 20) and two days
before the release of Opus 5 (July 24) — it belongs to the summer's
consolidation sequence, where each week brings its stone: legal peace,
second silicon pillar, democratization of the flagship. Taken together,
the two deals sketch a doctrine: depend on no single compute provider —
neither Nvidia (absent from both agreements), nor Amazon alone — while
making providers pay for the privilege of the reverse dependence. It is
the position of strength of the lab whose models have become the most
coveted asset in the compute industry.

