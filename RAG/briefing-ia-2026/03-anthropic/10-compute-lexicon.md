---
id: briefing-ia-2026/03-anthropic/10-compute-lexicon
title: "Anthropic: seven gigawatts and the compliance lexicon"
domain: anthropic
role: deep-dive
task: infrastructure
actors: ["AMD", "AWS", "Anthropic", "Glasswing", "Nvidia"]
dates: ["2026-04"]
keywords: ["alignment", "amd", "aws", "benchmark", "claude", "compute", "cyber", "distribution", "fable 5", "forward-commitment", "graviton", "helios"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s03-30"
source_lines: [4485, 4576]
sha256: 76782761aef76391ae9a96f4f80aa135740e087c40c33cd61e794171bb59b98a
---

# Anthropic: seven gigawatts and the compliance lexicon

<a id="s03-30"></a>
### Seven gigawatts: the arithmetic of Anthropic's compute

Let us add up the verified facts: 5 gigawatts secured from AWS (April
20–21 agreement, commitment of more than $100 billion over ten years,
Trainium2 → Trainium4 + Graviton trajectory, AWS as primary cloud) plus
up to 2 gigawatts of MI450-series in Helios racks from AMD (July 22
agreement, first gigawatt by the first half of 2027) — a total capacity
trajectory of 7 gigawatts. Seven gigawatts: the order of magnitude
deserves a pause, because it exceeds the electricity consumption of most
global metropolises, and it is dedicated to a single company and its
models.

The structure of these 7 GW tells the lab's doctrine of independence.
The 5 GW AWS is the foundation: a ten-year commitment of more than $100
billion — the biggest cloud forward-commitment by a frontier lab —
backed by Amazon's silicon roadmap (Trainium2 then Trainium4) and its
Graviton CPUs, with AWS as primary cloud. It is the integration bet:
entrusting a single partner with silicon, networking, and operations
for ten years, in exchange for a scale and predictability no one else
could offer in April 2026. The 2 GW AMD is the diversification: MI450
chips in Helios racks — AMD's system platform — with a first gigawatt
expected less than a year after the announcement (H1 2027), an
aggressive deployment signaling the urgency of the need. Two providers,
two silicon architectures (proprietary Trainium on one side, MI450 on
the other), two integration logics — and, notably, neither is Nvidia,
the historically dominant provider of AI compute, absent from both
agreements.

The deployment timeline is itself information: the first AMD gigawatt by
the first half of 2027, a few months after the end of 2026, indicates
the compute constraint is immediate — you don't deploy 1 GW in under a
year out of caution, but out of necessity. And the loop closes with the
third part of the AMD agreement — "Claude used to optimize ROCm": the
model optimizes the software stack of the provider whose chips will
train it, in a model-infrastructure co-optimization that prefigures the
26% of internal R&D "led" by Claude published in September. Compute is
no longer just the raw material of models: it becomes, via Claude, its
own optimization instrument — and 2026's 7 GW are both the fuel and the
product of this loop.

<a id="s03-31"></a>
### The 2026 compliance lexicon: safeguards, routing, Glasswing, classifier

Anthropic's 2026 produced, in June's urgency then in the summer's
consolidation, a compliance vocabulary worth fixing — because these
terms, invented or institutionalized by the lab, are on their way to
becoming the industry's common lexicon. Four concepts, four functions,
one overall architecture.

The safeguards, first: the guardrails integrated into the model, of
which Fable 5 offers the most accomplished demonstration — two levels,
one for the general public (Fable 5, full safeguards), the other
lightened on cyber for verified partners (Mythos 5). The safeguard is no
longer a binary attribute of the model (safe / dangerous) but a
graduated regime, adjusted to the recipient: safety becomes a function
of the trust granted to the user, not only a property of the system.

The routing, next: the mechanism by which Fable 5 redirects its
cyber/bio/chem requests to Opus 4.8 — an earlier, less capable model
with a better-mastered risk profile. Routing is the operational admission
of a truth the industry was slow to admit: you don't make the most
powerful model harmless through alignment alone; you accept, on at-risk
requests, a targeted degradation of capability by delegating to a
previous-generation model. It is the "capacity firewall" — and the fact
that the recipient is Opus 4.8, at 88.6% on SWE-bench and half the
flagship's price, shows the degradation is entirely relative: the
firewall is itself a cutting-edge model.

Project Glasswing, again: the apparatus for verifying the partners
authorized to access Mythos 5 — then, on September 1, Mythos 5.1.
Glasswing shifts the security burden from the model to the recipient:
the most sensitive capability is not bridled in the absolute, it is
entrusted to organizations whose identity and uses are established and
verifiable. Designed in June for selective distribution, tested during
the crisis as a prototype of geopolitical compliance, institutionalized
in September as the permanent channel of the unbridled versions,
Glasswing is the mechanism that gained the most stature over the year —
from partnership tool to governance standard.

The anti-jailbreak classifier, finally: added July 1 at restoration, it
is the defense-in-depth layer that filters, at the inference level,
attempts to circumvent the safeguards. Distinct from the model itself
and from routing, it is the purest product of the June crisis —
invented in nineteen days, kept after the crisis, inherited by Fable
5.1. Its lesson is general: a frontier model's safety does not rest on a
single mechanism but on a stack — graduated safeguards, routing to an
earlier model, recipient verification, jailbreak filtering — where each
layer compensates for the others' blind spots. It is this stack, more
than any benchmark score, that constitutes Anthropic's true 2026
innovation: the architecture of trust, built under fire, become the
standard against which all labs now measure themselves.
