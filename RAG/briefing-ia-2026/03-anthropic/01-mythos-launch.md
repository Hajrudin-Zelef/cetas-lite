---
id: briefing-ia-2026/03-anthropic/01-mythos-launch
title: "Anthropic: Mythos-class launch, safeguards, Glasswing, export control"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["AMD", "AWS", "Anthropic", "Glasswing", "United States", "Z.ai"]
dates: ["2026-06-09", "2026-06-12"]
keywords: ["export control", "safeguards", "agentic", "alignment", "amd", "aws", "claude", "compute", "consumer", "context window", "copyright", "cyber"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s03"
source_lines: [3363, 3550]
canonical_for: ["anthropic-mythos-saga"]
sha256: a83adceacf63503a5b0b3f4367de702895bde596b88255bfc3652572dfd10e33
---

# Anthropic: Mythos-class launch, safeguards, Glasswing, export control

<a id="s03"></a>
## 3. Anthropic in depth

2026 is, for Anthropic, the year of the tipping point: a technical
tipping point, with the release of a "Mythos-class" family of models
that redefines the top of the frontier-model hierarchy; a geopolitical
tipping point, with the first country-scale suspension of a frontier
model, followed by a nineteen-day restoration that urgently invents
compliance mechanisms that have become references; a financial tipping
point, with two of the largest fundraises in tech history,
infrastructure commitments exceeding $100 billion and a valuation
nearing $1,000 billion; and finally a scientific tipping point, with
Claude now driving more than a quarter of the lab's internal R&D and
producing published, validated wet-lab and mathematical results. This
section traces that trajectory in six parts: the Fable 5 / Mythos 5
saga (launch, suspension, restoration and their lessons), the September
iterations, the Opus / Sonnet lineup and its pricing logic, the
financial trajectory (raises, run-rate, AWS, AMD, IPO), the $1.5 billion
copyright settlement, then Claude's breakthrough in R&D and science —
before the Zhipu affair and its geopolitical reading.

The originality of Anthropic's position in 2026 lies in an embraced
contradiction: the lab is both the self-proclaimed champion of AI
safety — two-tier safeguard architecture, routing of sensitive requests,
closed partnerships via Project Glasswing — and the one that pushed the
capability cursor the furthest, to the point of triggering the first
export control that, in practice, targets by name a frontier model in
production. This tension between capability and control structures the
entire narrative of the year: every capability advance calls for a new
compliance mechanism, and every compliance mechanism, invented in an
emergency, then becomes a standard the lab exports to its partners. The
reader should therefore take away less a succession of launches than an
overall logic: the frontier AI company learns, under the joint pressure
of regulators and its own evaluation systems, to build the governance
infrastructure of its models at the same time as the models themselves.

<a id="s03-2"></a>
### June 9, 2026: the "Mythos-class" family takes the stage

On June 9, 2026, Anthropic simultaneously unveiled Claude Fable 5 and
Claude Mythos 5, inaugurating a new internal designation — the
"Mythos" class — which signals, in the lab's nomenclature, a capability
tier distinct from everything that preceded it. Fable 5 is the consumer
variant: a model available to all users, with full safeguards, priced at
10 dollars per million input tokens and 50 dollars per million output
tokens, with a one-million-token context window. This pricing
positioning — five times the price of Opus 4.8 released the previous
month — clearly shows its hand: Fable 5 is not one more general-purpose
model, it is the absolute flagship, whose price reflects both the cost
of compute and the claimed level of capability.

Mythos 5, for its part, is not a consumer product. It is the version
whose cyber safeguards are lifted, reserved for verified partners via a
program called Project Glasswing. The distinction between the two models
is architectural from the outset, not cosmetic: these are not two
marketing skins on a single system, but two safety regimes applied to a
single capability class. This duality installs a new doctrine in the
industry: frontier capability is no longer distributed binarily
(published or unpublished), but along a spectrum of trust, where the
level of safeguards decreases as the recipient's level of verification
increases. The general public receives Fable 5 bridled; verified
partners receive Mythos 5 unbridled on cyber — and between the two, an
internal routing mechanism sorts the most sensitive requests.

The one-million-token context window, common to both models, confirms
the generalization of what was still a distinctive selling point a few
months earlier. Opus 4.8 (May) and Sonnet 5 (June) already displayed
this million tokens: Fable 5 therefore distinguishes itself not by
context length, but by what the model does with that context — agentic
capabilities, reasoning, and precisely the cyber/bio/chem capabilities
that justify the very existence of the two safeguard levels. The price,
for its part, tells another story: at $10/$50 per million tokens,
Fable 5 costs exactly double Opus 5 ($5/$25), released July 24 with an
explicit promise — "quasi-Fable 5 for agentic use at half the price."
The lineup thus reads as a deliberate scale: Sonnet 5 ($2/$10) for value
for money, Opus 4.8 then Opus 5 ($5/$25) for agentic performance,
Fable 5 ($10/$50) for the capability summit.

<a id="s03-3"></a>
### Two levels of safeguards: cyber/bio/chem routing

The Fable 5 safety apparatus rests on a principle simple to state and
complex to operate: every request in the cyber, biological, or chemical
domains is routed to Opus 4.8 rather than handled by Fable 5 itself. In
other words, the flagship does not answer the most sensitive questions
directly; it delegates them to an earlier, less capable model whose risk
profile is better mastered. This architectural choice says a great deal
about Anthropic's doctrine in 2026: rather than trying to make the most
powerful model harmless through alignment alone, the lab accepts a
targeted degradation of capability on at-risk requests, relying on a
previous-generation model as a "capacity firewall."

The routing to Opus 4.8 is no accident of timing: Opus 4.8, released in
May with 88.6% on SWE-bench Verified, is at this date the lab's
highest-performing model outside Fable 5, and its behavior is
documented, tested, and has been deployed for several weeks. In choosing
it as the recipient of sensitive requests, Anthropic favors the known
over the powerful — a form of operational conservatism that contrasts
with the image of a lab rushing headlong toward maximum capability. The
implication is clear: the frontier is no longer just a race for
capability, it is a race for governable capability, and routing is the
instrument of that governance at the inference level.

Mythos 5 constitutes the other side of the apparatus: its cyber
safeguards are lifted, but its access is closed. The lifting of
safeguards is therefore not deregulation, it is a transfer of the
security burden from the model to the recipient — the verified partner.
This logic, which prefigures Project Glasswing, rests on a strong
hypothesis: certain capabilities (notably cyber) are legitimate in
verified hands and dangerous in anonymous hands, and the right answer is
not universal prohibition but selective distribution. This is a break
with the industry's earlier doctrine, under which a model was either
safe enough to be published or held back internally. Anthropic is
inventing here a third regime: published subject to verification.

<a id="s03-4"></a>
### Project Glasswing: the closed distribution of Mythos 5

Project Glasswing is the name of the apparatus through which Anthropic
verifies the partners authorized to access Mythos 5, the version with
lifted cyber safeguards. The verified fact is sober — "reserved for
verified partners via Project Glasswing" — but its implications are
considerable: for the first time, a frontier lab is putting in place a
dedicated infrastructure for verifying a model's recipients, distinct
from the model's own security infrastructure. Trust is no longer merely
a property of the system (alignment, safeguards, routing), it is also a
property of the recipient, established through a verification process of
which Glasswing is the instrument.

This apparatus takes on its full meaning in light of what follows: when
the US Department of Commerce imposes, on June 12, an export control
prohibiting access to any foreign national, it is precisely the
existence of a verified distribution circuit that will allow Anthropic,
on July 1, to restore access by distinguishing users according to their
status. Glasswing, designed to manage the selective distribution of
Mythos 5, thus becomes, by force of circumstance, the prototype of a
geopolitical compliance infrastructure. The urgency of June transforms a
partnership mechanism into a compliance mechanism — and it is this
conversion, carried out under pressure, that will make Glasswing a
standard.

Also note the symmetry with Mythos 5.1, released September 1 and
"reserved for verified organizations via Project Glasswing": the
apparatus survives the June crisis and becomes institutionalized as the
permanent channel for distributing the unbridled versions. What was an
experiment in selective distribution becomes a durable architecture:
each future Mythos generation will now have its bridled consumer
variant and its unbridled verified variant, with Glasswing acting as
customs between the two worlds.

<a id="s03-5"></a>
### June 12: the export control and the suspension

Three days after launch, on June 12, 2026, the United States Department
of Commerce imposes an export control prohibiting access to Fable 5 and
Mythos 5 to any foreign national — including, notably, foreign
employees. The scope of the measure is vertiginous: it is not a matter
of banning export to certain adversary countries, in the usual grammar
of export controls, but of excluding a category of persons —
non-Americans — from access to a digital service, wherever they are,
whoever their employer. A foreign employee of an American company,
working on American soil, is thus barred from accessing the model.

Anthropic's reaction is immediate and radical: the lab shuts down both
models — Fable 5 and Mythos 5 — for everyone, including American users
who were not targeted by the measure. This choice deserves attention,
because it was not the only technically conceivable option: Anthropic
could have attempted nationality-based filtering, approximate
geolocation, or partial suspension. By choosing total shutdown, the lab
makes a legible bet: rather than implementing, under deadline pressure
of a few hours, discriminatory filtering by nationality — a legally
perilous, technically fragile, and morally contestable operation — it
prefers the cost of universal, temporary unavailability. It is a refusal
to build, under a deadline of a few hours, an infrastructure for sorting
users by passport.

The event is historic on two counts. First, it is the first time an
export control has directly struck access to a frontier model in
production, rather than the hardware (GPUs), design software, or
technology transfers that produced it: regulation climbs the value chain
all the way to the service itself. Second, it is the first time a lab
has voluntarily shut down its flagship for its entire user base in
response to a regulatory measure — a precedent that redefines the
balance of power between labs and states: the state can prohibit, but
it is the lab that chooses how to obey, and that choice carries a
political cost for the regulator itself, faced with American users
deprived of service by its decision.

