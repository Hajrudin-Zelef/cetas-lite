---
id: briefing-ia-2026/11-consumer-agents-research/07-scientisttwo
title: "ScientistTwo: the autonomous discovery cycle (arXiv)"
domain: consumer-agents-research
role: deep-dive
task: research
actors: ["James Zou"]
dates: ["2026-09"]
keywords: ["agent", "agents", "inference", "research"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s11-9"
source_lines: [12323, 12399]
sha256: d8230652e291f22815623df03d679de90a20b566367d33d246145ee0ad57063c
---

# ScientistTwo: the autonomous discovery cycle (arXiv)

<a id="s11-9"></a>
### 11.8 ScientistTwo: the full discovery cycle, autonomous (arXiv 2609.19644, September 2026)

**The framework.** ScientistTwo, published on arXiv in
September 2026 (reference 2609.19644), is a multi-agent
framework that autonomously executes the full scientific
discovery cycle: holistic idea evaluation, hypothesis
refinement by ablation (testing which components of a
hypothesis are truly necessary), and peer review with
closed-loop rebuttal — i.e. the system simulates the peer
evaluation process, responds to criticism, and iterates.
Where Paper2Agent automates the reproduction of existing
work, ScientistTwo automates the production of new work:
they are two complementary halves of automated science.

**The quantified results.** On 107 research challenges
from the ICLR, ICML and NeurIPS conferences, the system
solves 80.4% of advanced problems, with a 25.2% relative
improvement over human SOTA — the best level achieved by
human researchers serving as the reference. More
striking still: the produced manuscripts reach the
acceptance thresholds of top-tier conferences, meaning
the system does not merely solve exercises, it produces
research judged publishable by today's most demanding
standards. This is a change of nature: scientific AI
moves from assistant status (suggest, code, review) to
that of a potentially autonomous author.

**The identified limits.** Two caveats are explicit.
First the cost: about $3,800 per task, which makes
intensive use prohibitive for most laboratories —
compared with Paper2Agent's $14, the factor is on the
order of 270. Second, and perhaps most important: no
"spotlight" breakthroughs yet, i.e. no major discoveries
recognized as such by the community. The system excels at
producing solid research at publication standards, but
has not yet produced the landmark result. This limit
traces the current frontier: automating "good advanced
routine scientific work" is demonstrated; automating
scientific genius is not.

**Cross-reading with Paper2Agent.** The two papers,
published days apart (Nature ~16/09, arXiv September
2026), together tell 2026's story of automated science:
on one side, reproduction becomes near-free and
instantaneous (Paper2Agent: 45 min, $14); on the other,
original production reaches publishable level but stays
expensive (ScientistTwo: $3,800/task). The logical
trajectory would be their convergence: researcher-agents
leaning on paper-agents to instantly reproduce the state
of the art before innovating — a division of labor where
reproduction, turned commodity, feeds discovery.
ScientistTwo's cost is the variable to watch: if it
follows a decline curve comparable to inference costs,
autonomous research could move from demonstrator to
laboratory routine within a few years.

**The closed loop as methodological novelty.** The most
original aspect of ScientistTwo is not so much the scores
as the process architecture: peer review and rebuttal in
a closed loop. The system does not merely generate
hypotheses and test them by ablation — it submits itself
to structured criticism, responds to it, and iterates.
This is the internal simulation of science's social
process: objection, defense, revision. From an
epistemological standpoint, it is both promising and
dizzying: promising, because criticism is the central
mechanism of scientific reliability; dizzying, because a
system that is at once author, reviewer and judge of its
own production can converge toward internal coherence
with no external anchoring — the risk of a "closed-vessel
science", impeccable in form and disconnected from
reality. The 107 challenges from real conferences (ICLR,
ICML, NeurIPS) fortunately provide an anchor: the
problems come from the real world of research, not from
the system itself.

