---
id: briefing-ia-2026/11-consumer-agents-research/06-paper2agent
title: "Paper2Agent: from publication to working agent (Nature)"
domain: consumer-agents-research
role: deep-dive
task: research
actors: ["Anthropic", "James Zou"]
dates: ["2026-09-16"]
keywords: ["agent", "agentic", "agents", "claude", "compute", "mcp"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s11-8"
source_lines: [12242, 12322]
sha256: 85d279a19d28e89cb5fe0c650b8b81e70626b7adedbde707781b7b1c48dfb6a0
---

# Paper2Agent: from publication to working agent (Nature)

<a id="s11-8"></a>
### 11.7 Paper2Agent: from scientific publication to functional agent (Nature, ~16/09/2026)

**The principle.** Published in Nature around 16 September
2026 by James Zou's team (Stanford), Paper2Agent converts
a scientific paper — along with its code and data — into
a functional AI agent, via an MCP (Model Context
Protocol) server, with auto-generated and validated
tools. The idea is simple and radical: rather than having
the researcher read the paper then reimplement the code,
an agent is automatically generated that encapsulates the
paper's method as callable, verified tools. The MCP
protocol serves as the exposure standard: the produced
agent is interoperable with the existing agent ecosystem
instead of being an ad hoc prototype.

**The AlphaGenome test.** The flagship demonstration
concerns AlphaGenome: an agent was assembled in about 45
minutes for about $14 of compute, with 22 MCP tools
generated and validated without human intervention. The
scores are spectacular: 98.7% accuracy on "tutorial"
queries (reproducing documented uses) and 100% on novel
queries (generalization beyond the tutorial). The
baselines give the measure of the leap: Claude with
simple repository access reaches 82.7% / 78.7%, and
Biomni — a biomedical agent framework — tops out at 37.3%
/ 56%. In other words, the automatically generated agent
outperforms both direct use of the frontier model on the
source code and the domain's specialized agentic
framework. The cost ($14) and time (45 minutes) make the
operation near-free at laboratory scale.

**Scaling up.** The team applied the method to 100
bioRxiv papers: 74 were successfully converted into
agents, and 593 of 599 tools were validated — a tool
validation rate near 99%. This second result is as
important as the first: it shows the method is not a lucky
hit on a favorable paper, but a robust pipeline (74%
conversion, ~99% valid tools). The 26 unconverted papers
nonetheless recall the limits: unavailable code, exotic
dependencies or non-computational methods remain
obstacles.

**Implications: automating scientific
reproduction.** Paper2Agent tackles one of the costliest
problems of modern science: reproducing results. Today,
reproducing a computational paper takes days or weeks of
engineering; tomorrow, a validated agent could be
available in under an hour for a few tens of dollars.
Eventually, one can imagine "paper-agent libraries" where
each computational publication comes with its executable
twin — which would change the very nature of scientific
publication, from static document to executable artifact.
The open question is validation: who certifies that the
agent faithfully reproduces the paper? The 22
"intervention-free validated" AlphaGenome tools suggest
automatic validation, whose criteria and blind spots
(edge cases, implicit code assumptions) will need
scrutiny. Still, on the published figures, Paper2Agent
is one of 2026's most convincing demonstrations of the
agent as scientific instrument.

**MCP as the interoperability standard.** Choosing the
MCP protocol to expose the generated agent is no
implementation detail: it is what makes Paper2Agent an
infrastructure brick rather than a demo. An agent exposed
via a standard MCP server can be invoked by other agents,
composed into workflows, chained with existing tools —
it becomes a citizen of the agentic ecosystem instead of
an ad hoc script. If the "paper-agent" practice
generalizes, MCP (or its successor) could become for
computational science what the PDF was for publication:
the default exchange format. Automatic tool validation
(22 of 22 for AlphaGenome, 593 of 599 at scale) is the
other pillar: without it, each generated agent would
require human review that would cancel the time savings.
The criteria of this validation remain to be audited — a
tool "validated" against nominal cases can still fail on
edge cases, and that is precisely where scientific
reproduction most often fails.

