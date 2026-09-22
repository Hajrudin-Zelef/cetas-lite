---
id: briefing-ia-2026/11-consumer-agents-research/09-concept-steering-breakthroughs
title: "Concept steering and applied breakthroughs (proteins, mathematics)"
domain: consumer-agents-research
role: deep-dive
task: research
actors: ["Anthropic", "James Zou", "Kevin Buzzard", "OpenAI"]
dates: ["2026-06", "2026-08", "2026-08-20", "2026-09", "2026-09-04"]
keywords: ["protein", "agent", "agentic", "agents", "benchmark", "chatgpt", "claude", "exploit", "fermat", "formalization", "lean", "opus 4"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s11-11"
source_lines: [12463, 12567]
sha256: c719389b27f468bad459c53d17efbda2728019bd7e150a96216efe717b375b20
---

# Concept steering and applied breakthroughs (proteins, mathematics)

<a id="s11-11"></a>
### 11.10 Concept steering: an active research theme, with no single dated breakthrough

**State of play.** Concept steering in LLMs — the art of
guiding a model's behavior by intervening on its internal
concept representations — is in 2026 an active research
theme, but no single dated breakthrough has been
identified for this year. This precision deserves
emphasis in a reference dossier: the absence of a "big
paper" does not mean absence of activity, but it forbids
dating a rupture that does not exist in the verified
facts.

**Related work.** One may cite, as related work, "Agentic
Chain-of-Thought Steering for Efficient and Controllable
LLM Reasoning" (June 2026), which applies steering to
agents' chain of thought: the goal is reasoning that is
both efficient and controllable — efficient (fewer wasted
tokens), controllable (behavior stays within the wanted
bounds). This is a typical example of how steering,
originating in interpretability research, migrates
toward agent engineering: controlling the reasoning of
an agent acting for hours (cf. ChatGPT Work, Cowork) is a
security requirement as much as a performance one.

**Why it matters despite the absence of a
breakthrough.** Concept steering touches two major 2026
stakes: the governability of long-running agents (how to
guarantee a multi-hour agent stays aligned with its
mission?) and fine personalization (how to durably steer
a personal agent's style and priorities?). The fact that
the domain remains an "active theme" rather than a
stabilized technology suggests that 2026's commercial
agents' safeguards still rest largely on methods external
to the model (system prompts, tool validation, human
approvals — like the approval cards of agent redesigns)
rather than on internal control of representations. This
is a limit to keep in mind when evaluating autonomy
promises.

<a id="s11-12"></a>
### 11.11 Applied breakthroughs: proteins and mathematics

**Claude designs proteins (~20/08/2026).** Around 20
August 2026, a demonstration established Claude as a
protein designer: with Opus 4.8 and Mythos Preview, 1,320
designs were generated, of which 354 minibinders were
validated in wet-lab (in the laboratory, on real
proteins), covering 14 of 15 targets, with hit rates
(success rates) of 22.6% to 35.1%. These figures deserve
unpacking: 354 experimentally validated minibinders is
not a computer simulation, it is real biological
validation — the most demanding filter there is. A hit
rate of 22.6 to 35.1% means about a quarter to a third of
proposed designs work: in the context of protein design,
where the space of possibilities is astronomical, this is
a remarkable rate. The 14/15 target ratio shows
generalization across diverse targets, not overfitting to
a favorable case. This demonstration places generalist
LLMs on terrain historically reserved for specialized
computational-biology tools — and it echoes Paper2Agent:
biology is in 2026 the domain where agentic AI shows the
most tangible experimental results.

**Why the wet-lab changes everything.** In evaluating
scientific AI, there is an implicit hierarchy of
evidence: simulation (the model predicts), computational
reproduction (the code runs), and experimental validation
(nature answers). The 354 wet-lab-validated minibinders
place the protein demonstration at the top of this
hierarchy: we are no longer talking about benchmark
scores, but about molecules that actually bind their
targets in the laboratory. This is also what sets this
announcement apart from mere "performances": a hit rate
of 22.6 to 35.1% is not an abstract figure, it is an
operational success rate for a laboratory that would
order syntheses. At 14 of 15 targets covered, the method
moreover shows a robustness that forbids reducing it to a
textbook case. For the dossier, this is proof that 2026's
agentic AI no longer merely talks about science: it
produces artifacts that science validates.

**Claude formalizes Fermat's last theorem
(04/09/2026).** On 4 September 2026, another demonstration
made waves: the formalization of Fermat's last theorem in
Lean — 13 million lines of Lean code, produced in 11 days,
covering 29,500 theorems, with validation by Kevin
Buzzard, a mathematician specializing in formalization.
One essential precision, not to be distorted: this is the
formalization of Wiles's proof, not a new proof of the
theorem. The distinction is crucial: the exploit is not
mathematical (the theorem has been proven since 1994),
it is logistical and formal — translating into
machine-checkable code one of the most complex proofs in
the history of mathematics, in 11 days. The 13 million
lines give the scale of the enterprise: no human, no
human team, could have produced that volume in that time.
Validation by Kevin Buzzard brings the formalizers'
community's seal of approval. With Paper2Agent and
ScientistTwo, this formalization completes 2026's
triptych of scientific AI: reproduce (Paper2Agent),
discover (ScientistTwo), certify (Fermat/Lean) — three
complementary functions of a science increasingly
assisted, even driven, by machines.

