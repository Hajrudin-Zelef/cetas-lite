---
id: briefing-ia-2026/03-anthropic/08-research-science
title: "Anthropic research: R&D share, protein design, Fermat, Zhipu accusation"
domain: anthropic
role: deep-dive
task: research
actors: ["Anthropic", "CISA", "China", "Glasswing", "Kevin Buzzard", "Z.ai"]
dates: ["2026-08-20", "2026-09-04", "2026-09-11"]
keywords: ["fermat", "protein", "research", "advisory", "claude", "compute", "consumer", "cyber", "cybersecurity", "distillation", "export controls", "fable 5"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s03-24"
source_lines: [4212, 4370]
canonical_for: ["anthropic-research"]
sha256: 11c10fdd5b87d54354e056ec10f29816207a12c35ae8b3e0b208d5b0cc556b71
---

# Anthropic research: R&D share, protein design, Fermat, Zhipu accusation

<a id="s03-24"></a>
### 26% of internal R&D "led" by Claude: the September tipping point

On September 17–18, 2026, Anthropic publishes August data according to
which Claude "leads" — pilots, directs — 26% of the lab's internal R&D,
versus less than 1% in February, and 0% in fully autonomous mode. Every
element of this sentence is a bombshell. 26%: more than a quarter of
the research of one of the world's most advanced AI labs is now piloted
by its own models — experiment design, results analysis, iterations, in
proportions that make it a first-rank contributor, not a supporting
assistant. Less than 1% in February: in six months, the share was
multiplied by more than twenty-six — an internal adoption curve steeper
still than the run-rate's. 0% fully autonomous: the precision is
capital, and Anthropic gives it itself — none of this work is conducted
end-to-end without human supervision; the "lead" is piloting under
control, not autonomy let loose in the wild.

The distinction between "led" and "fully autonomous" is the key to
reading the announcement: Anthropic claims deep automation of its
research while explicitly keeping the human in the loop — a position
that answers in advance concerns about the risks of self-accelerating
R&D without supervision. It is also a demonstration of confidence in its
own evaluation systems: entrusting 26% of its R&D to Claude assumes
proven internal safeguards, the very ones whose development the June
crisis accelerated. The lab applies to itself the doctrine it preaches:
maximum capability under maximum control.

The strategic implication is vertiginous and fits in a loop: if Claude
accelerates Anthropic's R&D, and if that R&D produces better Claudes,
then the lab enters a regime of self-improvement — limited, supervised,
partial (26%, not 100%; 0% full autonomy), but real and measured. It is
the first time a frontier lab has published a figure on the automation
of its own research, and this figure — a quarter, in six months, from
less than 1% — gives a timescale to what theorists called slow-motion
"intelligence explosion": not an overnight singularity, but a
composition of gains which, at this pace, recomposes the lab within a
few years. The announcements that follow — proteins, Fermat — are the
proofs by example of what this 26% concretely produces.

<a id="s03-25"></a>
### Protein design: 354 minibinders validated in wet-lab (~August 20)

Around August 20, 2026, Anthropic unveils protein-design results
obtained with Opus 4.8 and Mythos Preview: 1,320 designs generated, 354
minibinders validated in wet-lab — that is, by real biological
experiments, not simulations — on 14 of the 15 targeted targets, with
hit rates of 22.6% to 35.1%. The verified fact is dense; let us unfold
it. 1,320 designs: the scale is that of massive computational screening,
where the model proposes thousands of candidates. 354 validated in
wet-lab: experimental validation — protein synthesis, binding tests in
the wet laboratory — confirms that more than a quarter of the designs
effectively bind their target. 14/15 targets: the success is not a lucky
shot on an easy target, it is a near-systematic success on a
diversified panel. 22.6–35.1% hit rates: experimental success rates
which, in drug discovery, are considerable — the classic pharmaceutical
industry dreams of such screening yields.

The joint use of Opus 4.8 and "Mythos Preview" is a detail that says a
lot: the preliminary version of the Mythos family — the one whose
cyber/bio safeguards are lightened for verified uses — is mobilized for
biological research, exactly the domain where Fable 5 routes its
consumer requests to Opus 4.8. June's doctrine applies here to the
letter: maximum biological capability is reserved for a verified use
(Anthropic's own internal research), while the general public accesses
it only via the bridled model. The lab is its own first Glasswing
"verified partner."

The mention of "dual-use risks" remains, present in the verified facts:
minibinders capable of binding 14 targets out of 15 with such hit rates
are a therapeutic feat — and, by construction, a dual-use capability,
since the same design machinery can serve hostile ends. Anthropic does
not hide this duality: publishing it with the results is assuming that
the boundary between legitimate research and risk is now internal to
the model's capabilities, and that the answer can only be governance —
user verification, usage traceability, safeguards — not research
abstention. It is, applied to biology, the same lesson as June's for
cyber: the capability exists, it is immense, and we must learn to
distribute it without spreading it.

<a id="s03-26"></a>
### Fermat in 11 days: 13 million lines of Lean validated by Kevin Buzzard (September 4)

On September 4, 2026, Anthropic announces the formalization of Fermat's
Last Theorem: 13 million lines of Lean code produced in 11 days,
29,500 theorems formalized, the whole validated by Kevin Buzzard —
with an essential precision provided by the verified facts: this is the
formalization of Wiles's proof, not a new proof. This precision takes
nothing away from the feat, it situates it: the model did not prove
Fermat — Wiles did that in the 1990s — it translated that proof, one of
the most complex in the history of mathematics, into a
machine-verifiable formalism, at an unprecedented scale and speed.

The figures give the measure: 13 million lines of Lean in 11 days is
more than a million lines per day — a throughput exceeding by several
orders of magnitude what a human formalization team can produce, and
which had made the human mathematicians' Fermat formalization project a
long-haul enterprise. 29,500 theorems: the granularity of the
formalization — each lemma, each intermediate step made explicit and
verified — is itself a result, because it is what makes the proof
mechanically checkable. And validation by Kevin Buzzard — one of the
world's foremost specialists in Lean formalization — brings the
indispensable human endorsement: it is not Anthropic congratulating
itself, it is a reference third party attesting that the formalization
is correct.

The feat fits into the September sequence as the demonstration by
absurdity of Opus 5's "thinking by default" (July 24) and of the 26% of
"led" R&D: a model that reasons in depth by default, piloting research
under human supervision, can accomplish in 11 days what belonged to the
decadal project. And it prefigures a transformation of mathematics
itself: if exhaustive formalization becomes cheap, the demand for
mechanically verified rigor could extend to whole swaths of the
discipline — with, in the background, the question of what becomes of
mathematical intuition when verification is delegated to the machine.
Anthropic, as a good logician of its own doctrine, supplies the feat
and leaves the question open — but the brute fact remains: on
September 4, 2026, an AI formalized Wiles in 11 days, and Kevin Buzzard
said it was correct.

<a id="s03-27"></a>
### The accusation against Zhipu (September 11) and the September 8 advisory

On September 11, 2026, Anthropic publishes a report accusing Zhipu — a
Chinese AI lab — of "massive distillation" of Claude. The accusation is
grave and precise in its alleged mechanism: distillation, in machine
learning, consists of training a "student" model on the outputs of a
"teacher" model to transfer its capabilities without replicating its
training cost — a form of appropriation of others' compute and data
investments, proscribed by the terms of use of all frontier labs.
Qualified as "massive," the alleged distillation would not be an
isolated abusive use but a systematic campaign of capability extraction.

The context, provided by the verified facts, illuminates the
accusation's scope: on September 8, three days earlier, a joint advisory
from CISA, the NSA, and the FBI cited Z.AI — the structure associated
with Zhipu — already placing the Chinese actor in the crosshairs of
American cybersecurity and intelligence agencies. The sequence — agency
advisory on the 8th, Anthropic report on the 11th — suggests a
convergence, deliberate or fortuitous, between the American state's
security reading and the lab's competitive reading: the same actor is
designated, three days apart, as a cyber threat by the agencies and as
a model thief by Anthropic. Whether this convergence is coordinated or
not does not appear in the verified facts and will not be settled here;
the brute fact is the calendar coincidence and the common target.

The Zhipu affair closes the loop opened in June: the Fable 5 saga had
established that frontier capability is a strategic asset subject to
export controls; the September accusation establishes that this
capability, once deployed via API, can be siphoned by actors who have
not paid its cost — and that protecting models also means detecting and
denouncing distillation. Between the two, Anthropic built a whole
arsenal — two-tier safeguards, routing, Glasswing, anti-jailbreak
classifier — whose ultimate function is precisely this: ensuring that
the most advanced capability diffuses only to verified recipients, and
not to clandestine students. The September 11 report is thus the first
offensive use of June's defensive doctrine: no longer merely protecting
itself, but publicly naming those who would attempt to circumvent the
protection.

